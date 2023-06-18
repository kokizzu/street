<script>
  import { createEventDispatcher } from 'svelte';
  
  export let fields = []; // array of field object
  export let rows = []; // 2 dimension array
  export let pager = {}; // pagination
  
  function date( unixSec ) {
    if( !unixSec ) return '';
    const date = new Date( unixSec * 1000 );
    return date.toISOString().substring( 0, 10 );
  }
  
  const dispatch = createEventDispatcher();
  
  function gotoPage( page ) {
    dispatch( 'refreshTableView', {
      ...pager,
      page,
    } );
  }
  
  function changePerPage( perPage ) {
    dispatch( 'refreshTableView', {
      ...pager,
      perPage,
    } );
  }
  
  $: allowPrevPage = pager.page>1;
  $: allowNextPage = pager.page<pager.pages;
</script>
<table>
  <thead>
  <tr>
    {#each fields as field}
      <th>{field.label}</th>
    {/each}
  </tr>
  </thead>
  <tbody>
  {#each rows as row}
    <tr>
      {#each fields as field, i}
        {#if field.inputType==='checkbox'}
          <td>{!!row[ i ]}</td>
        {:else if field.inputType==='datetime'}
          <td>{date( row[ i ] )}</td>
        {:else}
          <td>{row[ i ]}</td>
        {/if}
      {/each}
    </tr>
  {/each}
  </tbody>
  <tfoot>
  <tr>
    <th colspan={fields.length}>
      <span class='left'>
      Page {pager.page} of {pager.pages},
      <input class='perPage' type='number' bind:value={pager.perPage} on:change={() => changePerPage(pager.perPage)} /> rows per page.
        </span>
      Total: {pager.countResult}
      <span class='right'>
        <button disabled={!allowPrevPage} on:click={() => gotoPage(1)}>|&lt;</button>
        <button disabled={!allowPrevPage} on:click={() => gotoPage(pager.page-1)}>&lt;</button>
        <button disabled={!allowNextPage} on:click={() => gotoPage(pager.page+1)}>&gt;</button>
        <button disabled={!allowNextPage} on:click={() => gotoPage(pager.pages)}>&gt;|</button>
      </span>
    </th>
  </tr>
  </tfoot>
</table>

<style>
    table {
        border-collapse : collapse;
    }

    table, th, td {
        border : 1px solid black;
    }

    input.perPage {
        width : 4em;
    }

    span.left {
        float : left;
    }

    span.right {
        float : right;
    }
</style>