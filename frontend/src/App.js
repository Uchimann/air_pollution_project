import React, { useState, useEffect } from 'react';
import Map from './components/Map';
import PollutionChart from './components/PollutionChart';
import AlertPanel from './components/AlertPanel';
import './App.css';

function App() {
  // localStorage'dan verileri yüklemeyi deneyin
  const [pollutionData, setPollutionData] = useState(() => {
    const savedData = localStorage.getItem('pollutionData');
    return savedData ? JSON.parse(savedData) : [];
  });
  
  const [alerts, setAlerts] = useState(() => {
    const savedAlerts = localStorage.getItem('alerts');
    return savedAlerts ? JSON.parse(savedAlerts) : [];
  });

  useEffect(() => {
    // pollutionData güncellendiğinde localStorage'a kaydet
    localStorage.setItem('pollutionData', JSON.stringify(pollutionData));
  }, [pollutionData]);

  useEffect(() => {
    // alerts güncellendiğinde localStorage'a kaydet
    localStorage.setItem('alerts', JSON.stringify(alerts));
  }, [alerts]);
  
  useEffect(() => {
    const eventSource = new EventSource('http://localhost:8081/events');
    
    eventSource.onmessage = (event) => {
      const newData = JSON.parse(event.data);
      
      setPollutionData(prevData => {
        const updatedData = [...prevData, newData].slice(-100);
        return updatedData;
      });
      
      if (newData.is_anomalous) {
        setAlerts(prevAlerts => {
          return [newData, ...prevAlerts].slice(0, 10); 
        });
      }
    };
    
    return () => {
      eventSource.close();
    };
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <h1>Hava Anomalileri İzleme Sistemi</h1>
      </header>
      
      <div className="dashboard">
        <div className="map-container">
          <h2>Kirlilik Haritası</h2>
          <Map pollutionData={pollutionData} />
        </div>
        
        <div className="chart-container">
          <h2>Kirlilik Grafikleri</h2>
          <PollutionChart data={pollutionData} />
        </div>
        
        <div className="alert-container">
          <h2>Uyarılar</h2>
          <AlertPanel alerts={alerts} />
        </div>
      </div>
    </div>
  );
}

export default App;