import React, { useEffect, useState } from 'react';
import { Chart as ChartJS, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend } from 'chart.js';
import { Line } from 'react-chartjs-2';

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

function PollutionChart({ data }) {
  const [chartData, setChartData] = useState({
    labels: [],
    datasets: []
  });

  useEffect(() => {
    if (data.length === 0) return;

    const pollutantGroups = {};
    data.forEach(item => {
      if (!pollutantGroups[item.pollutant]) {
        pollutantGroups[item.pollutant] = [];
      }
      pollutantGroups[item.pollutant].push(item);
    });

    const timestamps = data
      .slice(-20)
      .map(item => new Date(item.timestamp).toLocaleTimeString());

    const datasets = Object.entries(pollutantGroups).map(([pollutant, items], index) => {
      const colors = ['#FF6384', '#36A2EB', '#FFCE56', '#4BC0C0', '#9966FF'];
      const color = colors[index % colors.length];

      const values = items.slice(-20).map(item => item.value);

      return {
        label: pollutant,
        data: values,
        fill: false,
        borderColor: color,
        backgroundColor: color,
        tension: 0.1
      };
    });

    setChartData({
      labels: timestamps,
      datasets
    });
  }, [data]);

  const options = {
    responsive: true,
    plugins: {
      legend: {
        position: 'top',
      },
      title: {
        display: true,
        text: 'Kirlilik Seviyelerinin Zaman İçindeki Değişimi',
      },
    },
    scales: {
      y: {
        beginAtZero: true
      }
    }
  };

  return <Line data={chartData} options={options} />;
}

export default PollutionChart;