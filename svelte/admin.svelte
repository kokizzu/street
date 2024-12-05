<script>
  /** @typedef {import('./_types/user').User} User */
  /** @typedef {import('./_types/master').Access} Access */
  /** @typedef {import('chart.js').ChartConfiguration} ChartConfiguration */

  import Chart from 'chart.js/auto';
  import { onMount } from 'svelte';
  import Main from './_layouts/Main.svelte';
  import AdminSubMenu from './_components/AdminSubMenu.svelte';
  import { getRequestStatsChartOptions, getRequestActionChartOptions } from './_components/yChartOptions';
  
  let user                      = /** @type {User} */ ({/* user */});
  let access                    = /** @type {Access} */ ({/* segments */});
  let uniqueIpPerDate           = /** @type {Record<string, number>} */ ({/* uniqueIpPerDate */});
  let requestsPerDate           = /** @type {Record<string, number>} */ ({/* requestsPerDate */});
  let uniqueUserPerDate         = /** @type {Record<string, number>} */ ({/* uniqueUserPerDate */});
  let registeredUserTotal       = /** @type {number} */ (parseInt('#{registeredUserTotal}') || 0);
  let totalRegisteredUserToday  = /** @type {number} */ (parseInt('#{registeredUserToday}') || 0);
  let countPerActionsPerDate    = /** @type {Record<string, Record<string, number>>} */ ({/* countPerActionsPerDate */});

  let sortedDate = /** @type {string[]}*/ ([]);

  // Data to be display on the Charts
  let formattedDates              = [];
  let data_requestsPerDate        = [];
  let data_uniqueIpPerDate        = [];
  let data_uniqueUserPerDate      = [];
  let data_actionLists            = [];
  let data_countPerActionsPerDate = [];
  
  function formatDate(/** @type {string} */ dateString ) {
    const date = new Date( dateString );
    return date.toLocaleDateString('en-US', {
      day: 'numeric',
      month: 'long'
    });
  }
  
  onMount( () => {
    let uniqueDate = {};
    for( let i in uniqueIpPerDate ) uniqueDate[ i ] = true;
    for( let i in requestsPerDate ) uniqueDate[ i ] = true;
    for( let i in uniqueUserPerDate ) uniqueDate[ i ] = true;

    sortedDate = Object.keys( uniqueDate ).sort();

    formattedDates = sortedDate.map( date => formatDate( date ) );
    data_requestsPerDate = sortedDate.map( date => requestsPerDate[ date ] );
    data_uniqueIpPerDate = sortedDate.map( date => uniqueIpPerDate[ date ] );
    data_uniqueUserPerDate = sortedDate.map( date => uniqueUserPerDate[ date ] );

    let DatasetsReqActions = /** @type {import('chart.js').ChartDataset[]} */ ([]);

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
      DatasetsReqActions.push( {
        label: '/'+data_actionLists[ idx ],
        backgroundColor: color,
        borderColor: borderColor,
        data: data_countPerActionsPerDate[ idx ],
        fill: false,
        barThickness: 20,
      } );
    }
    
    const ElmStatsChart = /** @type {HTMLCanvasElement} */ (document.getElementById('stats-chart'));
    const ElmActionChart = /** @type {HTMLCanvasElement} */ (document.getElementById( 'action-chart' ));
    
    const DatasetsStatsChart = /** @type {import('chart.js').ChartDataset[]} */ ([
      {
        label: 'Requests',
        pointRadius: 4,
        pointBackgroundColor: '#8b5cf6',
        data: data_requestsPerDate,
        borderColor: '#8b5cf6',
        backgroundColor: '#8b5cf620',
        fill: 'start',
        tension: 0.4
      },
      {
        label: 'Unique IP',
        pointRadius: 4,
        pointBackgroundColor: '#eab308',
        backgroundColor: '#eab30830',
        borderColor: '#eab308',
        data: data_uniqueIpPerDate,
        fill: 'start',
        tension: 0.4
      },
      {
        label: 'Unique User',
        pointRadius: 4,
        pointBackgroundColor: '#22c55e',
        backgroundColor: '#22c55e30',
        borderColor: '#22c55e',
        data: data_uniqueUserPerDate,
        fill: 'start',
        tension: 0.4
      }
    ])
    new Chart(ElmStatsChart, getRequestStatsChartOptions(formattedDates, DatasetsStatsChart));
    
    new Chart( ElmActionChart, getRequestActionChartOptions(formattedDates, DatasetsReqActions));
  } );
</script>

<Main {user} {access}>
  <div class="admin-container">
    <AdminSubMenu />
    <div class="admin-content">
      <div class="info-card">
        <div class="pill-box">
          Registered User Total: {registeredUserTotal}
        </div>
        <div class="pill-box">
          Registered User Today: {totalRegisteredUserToday}
        </div>
      </div>
      <div class="charts-container">
        <div class="statistics">
          <header>
            <h3>Last 30 days User Statistics</h3>
          </header>
          <div class="stats">
            <canvas id="stats-chart" />
          </div>
        </div>
        <div class="actions">
          <header>
            <h3>Last 30 days Actions</h3>
          </header>
          <div class="action">
            <canvas id="action-chart" />
          </div>
        </div>
      </div>
    </div>
  </div>
</Main>

<style>
  .admin-container {
    display: flex;
    flex-direction: column;
    width: 100%;
    gap: 20px;
    padding: 10px 20px 20px;
  }

  .admin-container .admin-content {
    display: flex;
    flex-direction: column;
    width: 100%;
    gap: 20px;
  }
  
  .admin-container .admin-content .info-card {
    display: flex;
    flex-direction: row;
    width: 100%;
    gap: 10px;
    flex-wrap: wrap;
  }

  .admin-container .admin-content .info-card .pill-box {
    width: fit-content;
    padding: 10px 20px;
    border: 1px solid var(--gray-003);
    border-radius: 8px;
  }

  .admin-container .admin-content .charts-container {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  .admin-container .admin-content .charts-container header h3 {
    font-size: 20px;
    margin: 0;
  }

  .admin-container .admin-content .charts-container .statistics {
    width: 100%;
    height: 450px;
    border-radius: 8px;
    padding: 16px 16px 40px 16px;
    border: 1px solid var(--gray-003);
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .admin-container .admin-content .charts-container .statistics .stats {
    height: 92%;
    width: 100%;
  }

  .admin-container .admin-content .charts-container .actions {
    width: 100%;
    border-radius: 8px;
    height: 500px;
    padding: 16px 16px 25px 16px;
    border: 1px solid var(--gray-003);
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .charts-container .actions .action {
    height: 90%;
    width: 100%;
  }
</style>