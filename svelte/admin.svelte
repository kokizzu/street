<script>
  // @ts-nocheck
  import Chart from 'chart.js/auto';
  import { onMount } from 'svelte';

  import Menu from './_components/Menu.svelte';
  import AdminSubMenu from './_components/AdminSubMenu.svelte';
  import ProfileHeader from './_components/ProfileHeader.svelte';
  import Footer from './_components/Footer.svelte';
  
  let user = {/* user */};
  let segments = {/* segments */};
  
  let uniqueIpPerDate = {/* uniqueIpPerDate */};
  let requestsPerDate = {/* requestsPerDate */};
  let uniqueUserPerDate = {/* uniqueUserPerDate */};
  let registeredUserTotal = +'#{registeredUserTotal}';
  let countPerActionsPerDate = {/* countPerActionsPerDate */};
  let sortedDate = [];
  // Data to be display on the Charts
  let formattedDates = [];
  let data_requestsPerDate = [];
  let data_uniqueIpPerDate = [];
  let data_uniqueUserPerDate = [];
  let data_actionLists = [];
  let data_countPerActionsPerDate = [];

  function formatDate( dateString ) {
    const options = {day: 'numeric', month: 'long'};
    const date = new Date( dateString );
    return date.toLocaleDateString( 'en-US', options );
  }
  
  onMount( () => {
    let uniqueDate = {};
    for( let i in uniqueIpPerDate ) uniqueDate[ i ] = true;
    for( let i in requestsPerDate ) uniqueDate[ i ] = true;
    for( let i in uniqueUserPerDate ) uniqueDate[ i ] = true;
    sortedDate = Object.keys( uniqueDate ).sort();
    formattedDates = sortedDate.map(date => formatDate(date));
    data_requestsPerDate = sortedDate.map(date => requestsPerDate[date]);
    data_uniqueIpPerDate = sortedDate.map(date => uniqueIpPerDate[date]);
    data_uniqueUserPerDate = sortedDate.map(date => uniqueUserPerDate[date]);
    data_actionLists = Object.keys(countPerActionsPerDate);
    for (let action in countPerActionsPerDate) {
      let action_data = []
      for (let i = 0; i < sortedDate.length; i++) {
        date = sortedDate[i];
        if (countPerActionsPerDate[action][date]) {
          action_data.push(countPerActionsPerDate[action][date])
        } else {
          action_data.push(0)
        }
      }
      data_countPerActionsPerDate.push(action_data)
    }
  } );

  // init Chart.js
  onMount(async () => {
    const statsChart = document.getElementById('stats-chart');
    const actionChart = document.getElementById('action-chart');

    new Chart(statsChart, {
      type: 'line',
      data: {
        labels: formattedDates,
        datasets: [
          {
            label: 'Requests',
            pointRadius: 4,
            pointBackgroundColor: '#a5b4fc',
            backgroundColor: '#818cf8',
            borderColor: '#818cf8',
            data: data_requestsPerDate,
            fill: false
          }, {
            label: 'Unique IP',
            pointRadius: 4,
            pointBackgroundColor: '#fdba74',
            backgroundColor: '#fb923c',
            borderColor: '#fb923c',
            data: data_uniqueIpPerDate,
            fill: false
          }, {
            label: 'Unique User',
            pointRadius: 4,
            pointBackgroundColor: '#5eead4',
            backgroundColor: '#2dd4bf',
            borderColor: '#2dd4bf',
            data: data_uniqueUserPerDate,
            fill: false
          },
        ]
      },
      options: {
        plugins: {
          legend: {
            rtl: true,
            labels: {
              color: '#ffff',
              textAlign: 'right'
            }
          }
        },
        maintainAspectRatio: false,
        responsive: true,
        title: {
          display: false,
          text: 'Sales Charts',
          fontColor: 'white',
        },
        tooltips: {
          mode: 'index',
          intersect: false,
        },
        hover: {
          mode: 'nearest',
          intersect: true,
        },
        scales: {
          x: {
            ticks: {
              color: '#ffff'
            },
            grid: {
              tickColor: 'transparent',
              color: 'transparent'
            }
          },
          y: {
            beginAtZero: true,
            ticks: {
              stepSize: 100,
              color: '#ffff'
            },
            border: {
              dash: [4]
            },
            grid: {
              tickColor: 'rgb(255, 255, 255, 0.2)',
              tickBorderDash: [4],
              color: 'rgb(255, 255, 255, 0.2)'
            }
          }
        },
      },
    });

    new Chart(actionChart, {
      type: 'bar',
      data: {
        labels: formattedDates,
        datasets: [
          {
            label: data_actionLists[0],
            backgroundColor: '#ef4444',
            borderColor: '#ef4444',
            data: data_countPerActionsPerDate[0],
            fill: false,
            barThickness: 20
          }, {
            label: data_actionLists[1],
            backgroundColor: '#ed64a6',
            borderColor: '#ed64a6',
            data: data_countPerActionsPerDate[1],
            fill: false,
            barThickness: 20
          }, {
            label: data_actionLists[2],
            backgroundColor: '#06b6d4',
            borderColor: '#06b6d4',
            data: data_countPerActionsPerDate[2],
            fill: false,
            barThickness: 20
          }, {
            label: data_actionLists[3],
            backgroundColor: '#84cc16',
            borderColor: '#84cc16',
            data: data_countPerActionsPerDate[3],
            fill: false,
            barThickness: 20
          }, {
            label: data_actionLists[4],
            backgroundColor: '#78716c',
            borderColor: '#78716c',
            data: data_countPerActionsPerDate[4],
            fill: false,
            barThickness: 20
          }, {
            label: data_actionLists[5],
            backgroundColor: '#d946ef',
            borderColor: '#d946ef',
            data: data_countPerActionsPerDate[5],
            fill: false,
            barThickness: 20
          }, {
            label: data_actionLists[6],
            backgroundColor: '#14b8a6',
            borderColor: '#14b8a6',
            data: data_countPerActionsPerDate[6],
            fill: false,
            barThickness: 20,
          }, {
            label: data_actionLists[7],
            backgroundColor: '#f59e0b',
            borderColor: '#f59e0b',
            data: data_countPerActionsPerDate[7],
            fill: false,
            barThickness: 20
          }, {
            label: data_actionLists[8],
            backgroundColor: '#0ea5e9',
            borderColor: '#0ea5e9',
            data: data_countPerActionsPerDate[8],
            fill: false,
            barThickness: 20
          }, {
            label: data_actionLists[9],
            backgroundColor: '#eab308',
            borderColor: '#eab308',
            data: data_countPerActionsPerDate[9],
            fill: false,
            barThickness: 20
          }, {
            label: data_actionLists[10],
            backgroundColor: '#22c55e',
            borderColor: '#22c55e',
            data: data_countPerActionsPerDate[10],
            fill: false,
            barThickness: 20
          },
        ]
      },
      options: {
        plugins: {
          legend: {
            position: 'right',
            labels: {
              color: '#475569'
            }
          }
        },
        maintainAspectRatio: false,
        responsive: true,
        title: {
          display: false,
          text: 'Orders Chart',
        },
        tooltips: {
          mode: 'index',
          intersect: false,
        },
        hover: {
          mode: 'nearest',
          intersect: true,
        },
        scales: {
          x: {
            stacked: true,
            ticks: {
              color: '#475569'
            },
            grid: {
              tickColor: 'transparent',
              color: 'transparent'
            }
          },
          y: {
            stacked: true,
            ticks: {
              color: '#475569'
            },
            border: {
              dash: [4]
            },
            grid: {
              tickColor: '#d1d5db',
              tickBorderDash: [4],
              color: '#d1d5db'
            }
          }
        }
      },
    })
    
  });
</script>

<section class='dashboard'>
  <Menu access={segments} />
  <div class='dashboard_main_content'>
    <ProfileHeader />
    <AdminSubMenu></AdminSubMenu>
    <div class='content'>
      <div class='chart_container'>
        <!-- Statistics -->
        <div class='statistics'>
          <header>
            <h3>Statistics</h3>
          </header>
          <div class='stats'>
            <canvas id='stats-chart'></canvas>
          </div>
        </div>
        <!-- Actions [ Still Dummy Chart]-->
        <div class='actions'>
          <header>
            <h3>Actions</h3>
          </header>
          <div class='action'>
            <canvas id='action-chart'></canvas>
          </div>
        </div>
      </div>
    </div>
    <Footer></Footer>
  </div>
</section>

<style>
  .chart_container {
    position: relative;
    width: 88%;
    margin: -40px auto 20px auto;
    display: flex;
    flex-direction: column;
  }
  .chart_container .statistics {
    width: 100%;
    height: 400px;
    box-shadow: 0px 4px 24px 0px rgba(0, 0, 0, 0.25);
    background-color: #334155;
    border-radius: 8px;
    padding: 16px 16px 25px 16px;
  }
  .chart_container .statistics header h3 {
    font-size: 20px;
    color: white;
    margin: 0;
  }
  .chart_container .statistics .stats {
    height: 92%;
    width: 100%;
  }

  .chart_container .actions {
    width: 100%;
    margin-top: 30px;
    box-shadow: 0px 4px 24px 0px rgba(0, 0, 0, 0.25);
    background-color: white;
    border-radius: 8px;
    height: 500px;
    padding: 16px 16px 25px 16px;
  }
  .chart_container .actions header h3 {
    font-size: 20px;
    margin: 0 0 20px 0;
  }
  .chart_container .actions .action {
    height: 90%;
    width: 100%;
  }
</style>