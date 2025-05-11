import React, { useEffect, useRef, useState } from 'react';
import L from 'leaflet';
import 'leaflet/dist/leaflet.css';
import 'leaflet.heat';

delete L.Icon.Default.prototype._getIconUrl;
L.Icon.Default.mergeOptions({
  iconRetinaUrl: 'https://unpkg.com/leaflet@1.7.1/dist/images/marker-icon-2x.png',
  iconUrl: 'https://unpkg.com/leaflet@1.7.1/dist/images/marker-icon.png',
  shadowUrl: 'https://unpkg.com/leaflet@1.7.1/dist/images/marker-shadow.png',
});

const INITIAL_VIEW = [41.0082, 28.9784];
const INITIAL_ZOOM = 10;
const HEAT_CONFIG = {
  radius: 25,
  blur: 15,
  maxZoom: 5,
  minOpacity: 0.4,
  max: 1.0,
  gradient: {
    0.0: 'green',
    0.3: 'yellow',
    0.7: 'orange',
    1.0: 'red'
  }
};

function Map({ pollutionData }) {
  const mapRef = useRef(null);
  const mapInstanceRef = useRef(null);
  const markersRef = useRef([]);
  const heatLayerRef = useRef(null);
  const [showMarkers, setShowMarkers] = useState(true);
  const dataPointsRef = useRef({});

  const getRiskData = (riskLevel) => {
    switch (riskLevel) {
      case 'Hazardous':
        return { color: 'red', intensity: 1.0 };
      case 'Unhealthy':
        return { color: 'orange', intensity: 0.7 };
      case 'Moderate':
        return { color: 'yellow', intensity: 0.4 };
      default:
        return { color: 'blue', intensity: 0.2 };
    }
  };

  useEffect(() => {
    if (mapInstanceRef.current) return;
    
    mapInstanceRef.current = L.map(mapRef.current).setView(INITIAL_VIEW, INITIAL_ZOOM);
    
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: '&copy; OpenStreetMap contributors'
    }).addTo(mapInstanceRef.current);
    
    heatLayerRef.current = L.heatLayer([], HEAT_CONFIG).addTo(mapInstanceRef.current);
    
    mapInstanceRef.current.on('zoomend', adjustHeatRadius);
    
    addToggleMarkersButton();
    
    return () => cleanupMap();
  }, []);

  const adjustHeatRadius = () => {
    const currentZoom = mapInstanceRef.current.getZoom();
    const baseRadius = 25;
    
    let newRadius = baseRadius;
    if (currentZoom <= 5) {
      newRadius = baseRadius * 0.8;
    } else if (currentZoom > 8) {
      newRadius = baseRadius * 0.6;
    }
    
    if (heatLayerRef.current) {
      heatLayerRef.current.setOptions({ radius: newRadius });
    }
  };

const addToggleMarkersButton = () => {
  const toggleButton = L.control({ position: 'topright' });
  
  toggleButton.onAdd = () => {
    const button = L.DomUtil.create('button', 'toggle-markers-button');
    button.id = 'toggle-map-markers-btn';
    button.innerHTML = 'Raptiyeleri Gizle';
    button.style.cssText = 
      'background-color: #fff; padding: 8px 10px; border: 2px solid rgba(0,0,0,0.2); border-radius: 4px; cursor: pointer;';
    
    L.DomEvent.on(button, 'click', () => {
      const isMarkersVisible = markersRef.current.length > 0 && 
        markersRef.current[0]._map != null;
      
      if (isMarkersVisible) {
        hideMarkers();
        button.innerHTML = 'Raptiyeleri Göster';
      } else {
        updateMarkers(pollutionData);
        button.innerHTML = 'Raptiyeleri Gizle';
      }
      
      setShowMarkers(!isMarkersVisible);
    });
    
    return button;
  };
  
  toggleButton.addTo(mapInstanceRef.current);
};

  const hideMarkers = () => {
    markersRef.current.forEach(marker => {
      mapInstanceRef.current.removeLayer(marker);
    });
  };

  const cleanupMap = () => {
    if (mapInstanceRef.current) {
      mapInstanceRef.current.remove();
      mapInstanceRef.current = null;
    }
  };

  const updateMarkers = (data) => {
    markersRef.current.forEach(marker => marker.remove());
    markersRef.current = [];
    
    data.forEach(dataPoint => {
      if (dataPoint.is_anomalous) {
        const { color } = getRiskData(dataPoint.health_risk_level);
        
        const markerIcon = new L.Icon({
          iconUrl: `https://raw.githubusercontent.com/pointhi/leaflet-color-markers/master/img/marker-icon-2x-${color}.png`,
          shadowUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/0.7.7/images/marker-shadow.png',
          iconSize: [25, 41],
          iconAnchor: [12, 41],
          popupAnchor: [1, -34],
          shadowSize: [41, 41]
        });

        const marker = L.marker(
          [dataPoint.latitude, dataPoint.longitude], 
          { icon: markerIcon }
        ).bindPopup(createPopupContent(dataPoint));
        
        marker.addTo(mapInstanceRef.current);
        markersRef.current.push(marker);
      }
    });
  };

  const createPopupContent = (data) => {
    return `
      <strong>${data.pollutant}</strong><br>
      Değer: ${data.value}<br>
      Eşik: ${data.threshold_value}<br>
      Risk: ${data.health_risk_level}<br>
      Zaman: ${new Date(data.timestamp).toLocaleString()}
    `;
  };

  useEffect(() => {
    if (!mapInstanceRef.current || pollutionData.length === 0) return;
    
    updateDataPoints();
    
    if (showMarkers) {
      updateMarkers(pollutionData);
    }
  }, [pollutionData, showMarkers]);

  const updateDataPoints = () => {
    dataPointsRef.current = {};
    const heatPoints = [];
    
    pollutionData.forEach(data => {
      if (data.latitude && data.longitude) {
        const latLngKey = `${data.latitude.toFixed(6)},${data.longitude.toFixed(6)}`;
        dataPointsRef.current[latLngKey] = data;
        
        const { intensity } = getRiskData(data.health_risk_level);
        const heatIntensity = data.is_anomalous ? intensity : 0.1;
        
        heatPoints.push([data.latitude, data.longitude, heatIntensity]);
      }
    });
    
    updateHeatLayer(heatPoints);
  };

  const updateHeatLayer = (heatPoints) => {
    if (heatLayerRef.current) {
      if (heatLayerRef.current._map) {
        heatLayerRef.current._map.off('click', handleHeatLayerClick);
      }
      
      mapInstanceRef.current.removeLayer(heatLayerRef.current);
      heatLayerRef.current = L.heatLayer(heatPoints, HEAT_CONFIG).addTo(mapInstanceRef.current);
      
      mapInstanceRef.current.on('click', handleHeatLayerClick);
    }
  };

  const handleHeatLayerClick = (e) => {
    const nearestPoint = findNearestPoint(e.latlng);
    
    if (nearestPoint) {
      L.popup()
        .setLatLng(e.latlng)
        .setContent(createPopupContent(nearestPoint))
        .openOn(mapInstanceRef.current);
    }
  };

  const findNearestPoint = (latlng) => {
    let nearestPoint = null;
    let minDistance = Infinity;
    const MAX_DISTANCE = 0.05; // ~5km
    
    Object.keys(dataPointsRef.current).forEach(key => {
      const [pointLat, pointLng] = key.split(',').map(parseFloat);
      const distance = Math.sqrt(
        Math.pow((pointLat - latlng.lat), 2) + 
        Math.pow((pointLng - latlng.lng), 2)
      );
      
      if (distance < minDistance && distance < MAX_DISTANCE) {
        minDistance = distance;
        nearestPoint = dataPointsRef.current[key];
      }
    });
    
    return nearestPoint;
  };

  return <div ref={mapRef} style={{ width: '100%', height: '400px' }}></div>;
}

export default Map;