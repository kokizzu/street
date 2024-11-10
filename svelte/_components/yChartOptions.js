/** @typedef {import('chart.js').ChartConfiguration} ChartConfiguration */
/** @typedef {import('chart.js').ChartDataset} ChartDataset */

/**
 * @returns {ChartConfiguration}
 * @param {string[]} labels
 * @param {ChartDataset[]} datasets
 */
function getRequestStatsChartOptions(labels, datasets) {
  return {    
    type: 'line',
    data: {
      labels: labels,
      datasets: datasets,
    },
    options: {
      plugins: {
        legend: {
          rtl: true,
          labels: {
            usePointStyle: true,
            color: 'var(--gray-008)',
            textAlign: 'right',
          }
        }
      },
      maintainAspectRatio: false,
      responsive: true,
      scales: {
        x: {
          ticks: {
            color: 'var(--gray-008)',
          },
          grid: {
            display: false
          }
        },
        y: {
          beginAtZero: true,
          ticks: {
            stepSize: 50,
            color: 'var(--gray-008)',
          },
          border: {
            dash: [2],
          },
          grid: {
            tickColor: 'var(--orange-005)',
            tickBorderDash: [2],
            color: 'var(--orange-005)'
          }
        }
      }
    }
  }
}

/**
 * @returns {ChartConfiguration}
 * @param {string[]} labels
 * @param {ChartDataset[]} datasets
 */
function getRequestActionChartOptions(labels, datasets) {
  return {
    type: 'bar',
    data: {
      labels: labels,
      datasets: datasets,
    },
    options: {
      plugins: {
        legend: {
          position: 'right',
          labels: {
            color: '#475569',
          },
        },
      },
      maintainAspectRatio: false,
      responsive: true,
      hover: {
        mode: 'nearest',
        intersect: true,
      },
      scales: {
        x: {
          stacked: true,
          ticks: {
            color: '#475569',
          },
          grid: {
            tickColor: 'transparent',
            color: 'transparent',
          },
        },
        y: {
          stacked: true,
          ticks: {
            color: '#475569',
          },
          border: {
            dash: [4],
          },
          grid: {
            tickColor: '#D1D5DB',
            tickBorderDash: [4],
            color: '#D1D5DB',
          },
        },
      },
    },
  }
}

module.exports = {
  getRequestStatsChartOptions, getRequestActionChartOptions
}