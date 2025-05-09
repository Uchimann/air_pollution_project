import React, { useEffect, useRef } from 'react';
import L from 'leaflet';
import 'leaflet/dist/leaflet.css';
import 'leaflet.heat';

delete L.Icon.Default.prototype._getIconUrl;
L.Icon.Default.mergeOptions({
  iconRetinaUrl: 'https://unpkg.com/leaflet@1.7.1/dist/images/marker-icon-2x.png',
  iconUrl: 'https://unpkg.com/leaflet@1.7.1/dist/images/marker-icon.png',
  shadowUrl: 'https://unpkg.com/leaflet@1.7.1/dist/images/marker-shadow.png',
});

function Map({ pollutionData }) {
  const mapRef = useRef(null);
  const mapInstanceRef = useRef(null);
  const markersRef = useRef([]);
  const heatLayerRef = useRef(null);

  useEffect(() => {
    if (mapInstanceRef.current) return;
    
    mapInstanceRef.current = L.map(mapRef.current).setView([41.0082, 28.9784], 10);
    
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: '&copy; OpenStreetMap contributors'
    }).addTo(mapInstanceRef.current);
    
  heatLayerRef.current = L.heatLayer([], { 
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
  }).addTo(mapInstanceRef.current);
    
  mapInstanceRef.current.on('zoomend', () => {
  const currentZoom = mapInstanceRef.current.getZoom();
  const baseRadius = 25;
  
  let newRadius;
  if (currentZoom <= 5) {
    newRadius = baseRadius * 0.8;
  } else if (currentZoom <= 8) {
    newRadius = baseRadius;
  } else {
    newRadius = baseRadius * 0.6;
  }
  
  if (heatLayerRef.current) {
    heatLayerRef.current.setOptions({ radius: newRadius });
  }
});
  
  return () => {
    if (mapInstanceRef.current) {
      mapInstanceRef.current.remove();
      mapInstanceRef.current = null;
    }
  };
}, []);

  useEffect(() => {
    if (!mapInstanceRef.current || pollutionData.length === 0) return;
    
    markersRef.current.forEach(marker => marker.remove());
    markersRef.current = [];
    
    const heatPoints = [];
    
    pollutionData.forEach(data => {
      if (data.latitude && data.longitude) {
        const intensity = getIntensityFromRisk(data.health_risk_level, data.anomaly_level, data.is_anomalous);
        
        heatPoints.push([data.latitude, data.longitude, intensity]);
        
        if (data.is_anomalous && data.health_risk_level === 'Hazardous') {
          const marker = L.marker([data.latitude, data.longitude])
            .bindPopup(`
              <strong>${data.pollutant}</strong><br>
              Değer: ${data.value}<br>
              Eşik: ${data.threshold_value}<br>
              Risk: ${data.health_risk_level}<br>
              Zaman: ${new Date(data.timestamp).toLocaleString()}
            `);
          marker.addTo(mapInstanceRef.current);
          markersRef.current.push(marker);
        }
      }
    });
    
    if (heatLayerRef.current && heatPoints.length > 0) {
      heatLayerRef.current.setLatLngs(heatPoints);
    }
  }, [pollutionData]);

  const getIntensityFromRisk = (healthRiskLevel, anomalyLevel, isAnomalous) => {
    if (!isAnomalous) return 0.1;
    
    switch (healthRiskLevel) {
      case 'Hazardous':
        return 1.0;
      case 'Unhealthy':
        return 0.7;
      case 'Moderate':
        return 0.4;
      case 'Safe':
      default:
        return 0.2;
    }
  };

  return <div ref={mapRef} style={{ width: '100%', height: '400px' }}></div>;
}

export default Map;