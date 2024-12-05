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
          labels: {
            usePointStyle: true,
            color: '#343a40'
          }
        }
      },
      maintainAspectRatio: false,
      responsive: true,
      scales: {
        x: {
          ticks: {
            color: '#343a40',
          },
          grid: {
            display: false
          },
          border: {
            dash: [2],
            color: '#343a40'
          }
        },
        y: {
          beginAtZero: true,
          ticks: {
            stepSize: 50,
            color: '#343a40',
          },
          border: {
            dash: [2],
          },
          grid: {
            tickBorderDash: [2],
            color: '#343a40'
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
          position: 'top',
          align: 'start',
          labels: {
            usePointStyle: true,
            color: '#343a40',
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
            color: '#343a40',
          },
          grid: {
            tickColor: 'transparent',
            color: 'transparent',
          },
        },
        y: {
          stacked: true,
          ticks: {
            color: '#343a40',
          },
          border: {
            dash: [2],
          },
          grid: {
            tickBorderDash: [2],
            color: '#343a40',
          },
        },
      },
    },
  }
}

module.exports = {
  getRequestStatsChartOptions, getRequestActionChartOptions
}