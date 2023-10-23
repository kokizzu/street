<script>
  // note: this is admin dashboard
  
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
  let totalRegisteredUserToday = +'#{registeredUserToday}';
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
    console.log('onMount.admin')
    let uniqueDate = {};
    for( let i in uniqueIpPerDate ) uniqueDate[ i ] = true;
    for( let i in requestsPerDate ) uniqueDate[ i ] = true;
    for( let i in uniqueUserPerDate ) uniqueDate[ i ] = true;
    sortedDate = Object.keys( uniqueDate ).sort();
    formattedDates = sortedDate.map( date => formatDate( date ) );
    data_requestsPerDate = sortedDate.map( date => requestsPerDate[ date ] );
    data_uniqueIpPerDate = sortedDate.map( date => uniqueIpPerDate[ date ] );
    data_uniqueUserPerDate = sortedDate.map( date => uniqueUserPerDate[ date ] );
    let datasets = [];
    data_actionLists = Object.keys( countPerActionsPerDate );
    for( let action in countPerActionsPerDate ) {
      let action_data = [];
      for( let i = 0; i<sortedDate.length; i++ ) {
        let date = sortedDate[ i ];
        action_data.push( countPerActionsPerDate[ action ][ date ] || 0 );
      }
      data_countPerActionsPerDate.push( action_data );
    }
    
    // fill the data
    let total = data_actionLists.length;
    
    for( let idx = 0; idx<total; ++idx ) {
      const degree = Math.floor( (360 * idx) / (total + 1) ); // since 360 in hsl = 0
      const color = 'hsl(' + degree + ', 100%, 47%)';
      const borderColor = 'hsl(' + degree + ', 100%, 60%)';
      datasets.push( {
        label: data_actionLists[ idx ],
        backgroundColor: color,
        borderColor: borderColor,
        data: data_countPerActionsPerDate[ idx ],
        fill: false,
        barThickness: 20,
      } );
    }
    
    // init Chart.js
    const statsChart = document.getElementById( 'stats-chart' );
    const actionChart = document.getElementById( 'action-chart' );
    
    new Chart( statsChart, {
      type: 'line',
      data: {
        labels: formattedDates,
        datasets: [
          {
            label: 'Requests',
            pointRadius: 4,
            pointBackgroundColor: '#A5B4FC',
            backgroundColor: '#818CF8',
            borderColor: '#818CF8',
            data: data_requestsPerDate,
            fill: false,
          },
          {
            label: 'Unique IP',
            pointRadius: 4,
            pointBackgroundColor: '#FDBA74',
            backgroundColor: '#FB923C',
            borderColor: '#FB923C',
            data: data_uniqueIpPerDate,
            fill: false,
          },
          {
            label: 'Unique User',
            pointRadius: 4,
            pointBackgroundColor: '#5EEAD4',
            backgroundColor: '#2DD4BF',
            borderColor: '#2DD4BF',
            data: data_uniqueUserPerDate,
            fill: false,
          },
        ],
      },
      options: {
        plugins: {
          legend: {
            rtl: true,
            labels: {
              color: '#FFF',
              textAlign: 'right',
            },
          },
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
              color: '#FFF',
            },
            grid: {
              tickColor: 'transparent',
              color: 'transparent',
            },
          },
          y: {
            beginAtZero: true,
            ticks: {
              stepSize: 100,
              color: '#FFF',
            },
            border: {
              dash: [4],
            },
            grid: {
              tickColor: 'rgb(255, 255, 255, 0.2)',
              tickBorderDash: [4],
              color: 'rgb(255, 255, 255, 0.2)',
            },
          },
        },
      },
    } );
    
    new Chart( actionChart, {
      type: 'bar',
      data: {
        labels: formattedDates,
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
    } );
  } );
</script>

<section class='dashboard'>
  <Menu access={segments} />
  <div class='dashboard_main_content'>
    <ProfileHeader />
    <AdminSubMenu />
    <div class='content'>
      <div class='info_card'>
        <div class='total_container'>
          <strong>Registered User Total</strong>: {registeredUserTotal}
        </div>
        <div class='total_container'>
          <strong>Registered User Today</strong>: {totalRegisteredUserToday}
        </div>
      </div>
      <div class='chart_container'>
        <!-- Statistics -->
        <div class='statistics'>
          <header>
            <h3>Last 30 days User Statistics</h3>
          </header>
          <div class='stats'>
            <canvas id='stats-chart' />
          </div>
        </div>
        <!-- Actions [ Still Dummy Chart]-->
        <div class='actions'>
          <header>
            <h3>Last 30 days Actions</h3>
          </header>
          <div class='action'>
            <canvas id='action-chart' />
          </div>
        </div>
      </div>
    </div>
    <Footer />
  </div>
</section>

<style>
  .info_card {
    display          : flex;
    width            : 88%;
    margin           : -20px auto 20px auto;
    position         : relative;
    background-color : antiquewhite;
    padding          : 1em;
    border-radius    : 10px;
    flex-direction   : row;
  }

  .total_container {
    flex : auto;
  }

  .chart_container {
    position       : relative;
    width          : 88%;
    margin         : 20px auto 20px auto;
    display        : flex;
    flex-direction : column;
  }

  .chart_container .statistics {
    width            : 100%;
    height           : 400px;
    box-shadow       : 0px 4px 24px 0px rgba(0, 0, 0, 0.25);
    background-color : #334155;
    border-radius    : 8px;
    padding          : 16px 16px 25px 16px;
  }

  .chart_container .statistics header h3 {
    font-size : 20px;
    color     : white;
    margin    : 0;
  }

  .chart_container .statistics .stats {
    height : 92%;
    width  : 100%;
  }

  .chart_container .actions {
    width            : 100%;
    margin-top       : 30px;
    box-shadow       : 0px 4px 24px 0px rgba(0, 0, 0, 0.25);
    background-color : white;
    border-radius    : 8px;
    height           : 500px;
    padding          : 16px 16px 25px 16px;
  }

  .chart_container .actions header h3 {
    font-size : 20px;
    margin    : 0 0 20px 0;
  }

  .chart_container .actions .action {
    height : 90%;
    width  : 100%;
  }
</style>
