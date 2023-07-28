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
  // Each data to be display
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
    // data_countPerActionsPerDate = Object.values(countPerActionsPerDate);
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

    // console.log(data_countPerActionsPerDate)
    // console.log(data_actionLists)
    // console.log(countPerActionsPerDate)
    // console.log(sortedDate)
  } );

  // init chart
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
            backgroundColor: '#ed64a6',
            borderColor: '#ed64a6',
            data: data_countPerActionsPerDate[0],
            fill: false,
            barThickness: 5
          }, {
            label: data_actionLists[1],
            backgroundColor: '#ed64a6',
            borderColor: '#ed64a6',
            data: data_countPerActionsPerDate[1],
            fill: false,
            barThickness: 5
          }, {
            label: data_actionLists[2],
            backgroundColor: '#ed64a6',
            borderColor: '#ed64a6',
            data: data_countPerActionsPerDate[2],
            fill: false,
            barThickness: 5
          }, {
            label: data_actionLists[3],
            backgroundColor: '#ed64a6',
            borderColor: '#ed64a6',
            data: data_countPerActionsPerDate[3],
            fill: false,
            barThickness: 5
          }, {
            label: data_actionLists[4],
            backgroundColor: '#ed64a6',
            borderColor: '#ed64a6',
            data: data_countPerActionsPerDate[4],
            fill: false,
            barThickness: 5
          }, {
            label: data_actionLists[5],
            backgroundColor: '#ed64a6',
            borderColor: '#ed64a6',
            data: data_countPerActionsPerDate[5],
            fill: false,
            barThickness: 5
          }, /* TODO: Display Datasets to 11*/
        ]
      },
      options: {
        plugins: {
          legend: {
            position: 'right'
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
          x: {},
          y: {

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
      
      <table class='table_actions'>
        <tr class='table_row'>
          <th class='table_header'>Actions</th>
          {#each sortedDate as date}
            <th class='table_header'>{formatDate( date )}</th>
          {/each}
        </tr>
        {#each Object.keys( countPerActionsPerDate ) as actionsPerDate}
          <tr class='table_row'>
            <td class='table_data'>{actionsPerDate}</td>
            {#each sortedDate as date}
              <td class='table_data'>
                {countPerActionsPerDate[ actionsPerDate ][ date ] || '0'}
              </td>
            {/each}
          </tr>
        {/each}
      </table>
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


    .table_actions {
        margin-top : 30px;
    }
    
    .table_actions{
        margin-left      : auto;
        margin-right     : auto;
        border-radius    : 8px;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
        padding          : 20px;
        background-color : white;
        width            : 88%;
        height           : fit-content;
        color            : #475569;
        font-size        : 16px;
    }

    .table_header {
        text-align : left !important;
        color      : #6366F1 !important;
    }

    .table_header, .table_data {
        padding-top    : 7px;
        padding-bottom : 7px;
    }
</style>