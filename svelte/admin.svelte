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
  } );

  // init chart
  onMount(async () => {
    const statsChart = document.getElementById("stats-chart");
    const actionChart = document.getElementById("action-stats")

    new Chart(statsChart, {
      type: "line",
      data: {
        labels: formattedDates,
        datasets: [
          {
            label: 'Requests',
            backgroundColor: "#818cf8",
            borderColor: "#818cf8",
            data: data_requestsPerDate,
            fill: false
          }, {
            label: 'Unique IP',
            backgroundColor: "#fb923c",
            borderColor: "#fb923c",
            data: data_uniqueIpPerDate,
            fill: false
          }, {
            label: 'Unique User',
            backgroundColor: "#2dd4bf",
            borderColor: "#2dd4bf",
            data: data_uniqueUserPerDate,
            fill: false
          },
        ]
      },
      options: {
        maintainAspectRatio: false,
        responsive: true,
        title: {
          display: false,
          text: "Sales Charts",
          fontColor: "white",
        },
        legend: {
          labels: {
            fontColor: "white",
          },
          align: "end",
          position: "bottom",
        },
        tooltips: {
          mode: "index",
          intersect: false,
        },
        hover: {
          mode: "nearest",
          intersect: true,
        },
        scales: {
          xAxes: [
            {
              ticks: {
                fontColor: "rgba(255,255,255,.7)",
              },
              display: true,
              scaleLabel: {
                display: false,
                labelString: "Month",
                fontColor: "white",
              },
              gridLines: {
                display: false,
                borderDash: [2],
                borderDashOffset: [2],
                color: "rgba(33, 37, 41, 0.3)",
                zeroLineColor: "rgba(0, 0, 0, 0)",
                zeroLineBorderDash: [2],
                zeroLineBorderDashOffset: [2],
              },
            },
          ],
          yAxes: [
            {
              ticks: {
                fontColor: "rgba(255,255,255,.7)",
              },
              display: true,
              scaleLabel: {
                display: false,
                labelString: "Value",
                fontColor: "white",
              },
              gridLines: {
                borderDash: [3],
                borderDashOffset: [3],
                drawBorder: false,
                color: "rgba(255, 255, 255, 0.15)",
                zeroLineColor: "rgba(33, 37, 41, 0)",
                zeroLineBorderDash: [2],
                zeroLineBorderDashOffset: [2],
              },
            },
          ],
        },
      },
    });

    new Chart(actionChart, {
      type: "bar",
      data: {
        labels: [
          "January",
          "February",
          "March",
          "April",
          "May",
          "June",
          "July",
        ],
        datasets: [
          {
            label: new Date().getFullYear(),
            backgroundColor: "#ed64a6",
            borderColor: "#ed64a6",
            data: [30, 78, 56, 34, 100, 45, 13],
            fill: false,
            barThickness: 8
          },
          {
            label: new Date().getFullYear() - 1,
            fill: false,
            backgroundColor: "#818cf8",
            borderColor: "#818cf8",
            data: [27, 68, 86, 74, 10, 4, 87],
            barThickness: 8
          }
        ]
      },
      options: {
        maintainAspectRatio: false,
        responsive: true,
        title: {
          display: false,
          text: "Orders Chart",
        },
        tooltips: {
          mode: "index",
          intersect: false,
        },
        hover: {
          mode: "nearest",
          intersect: true,
        },
        legend: {
          labels: {
            fontColor: "rgba(0,0,0,.4)",
          },
          align: "end",
          position: "bottom",
        },
        scales: {
          xAxes: [
            {
              display: false,
              scaleLabel: {
                display: true,
                labelString: "Month",
              },
              gridLines: {
                borderDash: [2],
                borderDashOffset: [2],
                color: "rgba(33, 37, 41, 0.3)",
                zeroLineColor: "rgba(33, 37, 41, 0.3)",
                zeroLineBorderDash: [2],
                zeroLineBorderDashOffset: [2],
              },
            },
          ],
          yAxes: [
            {
              display: true,
              scaleLabel: {
                display: false,
                labelString: "Value",
              },
              gridLines: {
                borderDash: [2],
                drawBorder: false,
                borderDashOffset: [2],
                color: "rgba(33, 37, 41, 0.2)",
                zeroLineColor: "rgba(33, 37, 41, 0.15)",
                zeroLineBorderDash: [2],
                zeroLineBorderDashOffset: [2],
              },
            },
          ],
        },
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
          <canvas id='action-stats'></canvas>
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
    height: 300px;
    box-shadow: 0px 4px 24px 0px rgba(0, 0, 0, 0.25);
    background-color: #334155;
    border-radius: 8px;
    padding: 16px 16px 25px 16px;
  }
  .chart_container .statistics header {
    margin-bottom: 20px;
  }
  .chart_container .statistics header h3 {
    font-size: 20px;
    color: white;
    margin: 0;
  }
  .chart_container .statistics .stats {
    height: 87%;
    width: 100%;
  }

  .chart_container .actions {
    width: 100%;
    margin-top: 40px;
    box-shadow: 0px 4px 24px 0px rgba(0, 0, 0, 0.25);
    background-color: white;
    border-radius: 8px;
    height: 350px;
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