<script>
  import { createEventDispatcher } from 'svelte';
  import { datetime } from './formatter.js';
  
  export let fields = []; // array of field object
  export let rows = []; // 2 dimension array
  export let pager = {}; // pagination
  
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
  
  function editRow( row ) {
    dispatch( 'editRow', row );
  }
  
  $: allowPrevPage = pager.page>1;
  $: allowNextPage = pager.page<pager.pages;
</script>
<table>
  <thead>
  <tr>
    {#each fields as field}
      {#if field.name==='id'}
        <th>Action</th>
      {:else}
        <th>{field.label}</th>
      {/if}
    {/each}
  </tr>
  </thead>
  <tbody>
  {#each rows as row, no}
    <tr>
      {#each fields as field, i}
        {#if field.name==='id'}
          <td>
            <button class='action' on:click={() => editRow(row[i])}>
              <i class='gg-pen' />
            </button>
            {no + 1}
          
          </td>
        {:else if field.inputType==='checkbox'}
          <td>{!!row[ i ]}</td>
        {:else if field.inputType==='datetime'}
          <td>{datetime( row[ i ] )}</td>
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

    button.action {
        height : 2em;
    }
</style>