import React from 'react';

function AlertPanel({ alerts }) {
  const getAlertStyle = (level) => {
    switch (level) {
      case 'High':
        return { backgroundColor: '#ffcccc', borderLeft: '5px solid #ff0000' };
      case 'Medium':
        return { backgroundColor: '#fff2cc', borderLeft: '5px solid #ffa500' };
      case 'Low':
        return { backgroundColor: '#e6ffe6', borderLeft: '5px solid #00cc44' };
      default:
        return { backgroundColor: '#f2f2f2', borderLeft: '5px solid #808080' };
    }
  };

  if (alerts.length === 0) {
    return (
      <div className="no-alerts">
        <p>Şu anda aktif uyarı bulunmuyor.</p>
      </div>
    );
  }

  return (
    <div className="alerts-list">
      {alerts.map((alert, index) => (
        <div
          key={index}
          className="alert-item"
          style={{
            ...getAlertStyle(alert.anomaly_level),
            padding: '10px',
            margin: '10px 0',
            borderRadius: '4px'
          }}
        >
          <h3>{alert.pollutant} - {alert.anomaly_level} Seviye Uyarı</h3>
          <p>Değer: {alert.value} (Eşik: {alert.threshold_value})</p>
          <p>Sağlık Riski: {alert.health_risk_level}</p>
          <p>Zaman: {new Date(alert.timestamp).toLocaleString()}</p>
        </div>
      ))}
    </div>
  );
}

export default AlertPanel;