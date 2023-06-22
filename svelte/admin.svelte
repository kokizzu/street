<script>
  import Menu from './_components/Menu.svelte';
  import AdminSubMenu from './admin/_adminSubMenu.svelte';
  import { onMount } from 'svelte';
  
  let user = {/* user */};
  let segments = {/* segments */};
  
  let uniqueIpPerDate = {/* uniqueIpPerDate */};
  let requestsPerDate = {/* requestsPerDate */};
  let uniqueUserPerDate = {/* uniqueUserPerDate */};
  let registeredUserTotal = +'#{registeredUserTotal}';
  let countPerActionsPerDate = {/* countPerActionsPerDate */};
  
  let sortedDate = [];
  onMount( () => {
    let uniqueDate = {};
    for( let i in uniqueIpPerDate ) uniqueDate[ i ] = true;
    for( let i in requestsPerDate ) uniqueDate[ i ] = true;
    for( let i in uniqueUserPerDate ) uniqueDate[ i ] = true;
    sortedDate = Object.keys( uniqueDate ).sort();
    console.log( sortedDate, uniqueIpPerDate, requestsPerDate, uniqueUserPerDate, uniqueDate, registeredUserTotal );
  } );

</script>

<Menu access={segments} />
<AdminSubMenu></AdminSubMenu>
<table>
  <tr>
    <th>Stat</th>
    {#each sortedDate as date}
      <th>{(date + '').substring( 5, 10 ).replace( '-', ' ' )}</th>
    {/each}
  </tr>
  <tr>
    <td>Requests</td>
    {#each sortedDate as date}
      <td>{requestsPerDate[ date ] || ''}</td>
    {/each}
  </tr>
  <tr>
    <td>UniqueIP</td>
    {#each sortedDate as date}
      <td>{uniqueIpPerDate[ date ] || ''}</td>
    {/each}
  </tr>
  <tr>
    <td>Unique User</td>
    {#each sortedDate as date}
      <td>{uniqueUserPerDate[ date ] || ''}</td>
    {/each}
  </tr>
  <tr>
    <th>Actions</th>
  </tr>
  {#each Object.keys( countPerActionsPerDate ) as actionsPerDate}
    <tr>
      <td>{actionsPerDate}</td>
      {#each sortedDate as date}
        <td>{countPerActionsPerDate[ actionsPerDate ][ date ] || ''}</td>
      {/each}
    </tr>
  {/each}
</table>
<style>
    table, td, th {
        border : 1px solid black;
    }

    table {
        border-collapse : collapse;
        width           : 100%;
    }
</style>