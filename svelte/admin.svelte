<script>
  import Menu from './_components/Menu.svelte';
  import AdminSubMenu from './_components/AdminSubMenu.svelte';
  import ProfileHeader from './_components/ProfileHeader.svelte';
  import Footer from './_components/Footer.svelte';
  import { onMount } from 'svelte';
  
  let user = {/* user */};
  let segments = {/* segments */};
  
  let uniqueIpPerDate = {/* uniqueIpPerDate */};
  let requestsPerDate = {/* requestsPerDate */};
  let uniqueUserPerDate = {/* uniqueUserPerDate */};
  let registeredUserTotal = +'#{registeredUserTotal}';
  let countPerActionsPerDate = {/* countPerActionsPerDate */};
  
  let sortedDate = [];
  
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
    console.log( sortedDate, uniqueIpPerDate, requestsPerDate, uniqueUserPerDate, uniqueDate, registeredUserTotal );
  } );
</script>

<section class='dashboard'>
  <Menu access={segments} />
  <div class='dashboard_main_content'>
    <ProfileHeader />
    <AdminSubMenu></AdminSubMenu>
    <div class='content'>
      <table class='table_stats'>
        <tr class='table_row'>
          <th class='table_header'>Stats</th>
          {#each sortedDate as date}
            <!-- <th>{(date + '').substring( 5, 10 ).replace( '-', ':' )}</th> -->
            <th class='table_header'>{formatDate( date )}</th>
          {/each}
        </tr>
        <tr class='table_row'>
          <td class='table_data'>Requests</td>
          {#each sortedDate as date}
            <td class='table_data'>{requestsPerDate[ date ] || '0'}</td>
          {/each}
        </tr>
        <tr class='table_row'>
          <td class='table_data'>UniqueIP</td>
          {#each sortedDate as date}
            <td class='table_data'>{uniqueIpPerDate[ date ] || '0'}</td>
          {/each}
        </tr>
        <tr class='table_row'>
          <td class='table_data'>Unique User</td>
          {#each sortedDate as date}
            <td class='table_data'>{uniqueUserPerDate[ date ] || '0'}</td>
          {/each}
        </tr>
      </table>
      
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
    .table_stats {
        position   : relative;
        margin-top : -40px;
    }

    .table_actions {
        margin-top : 30px;
    }

    .table_stats, .table_actions {
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