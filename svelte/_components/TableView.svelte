<script>
  import { createEventDispatcher } from 'svelte';
  import { datetime } from './formatter.js';
  
  export let fields = []; // array of field object
  export let rows = []; // 2 dimension array
  export let pager = {}; // pagination
  
  $: deletedAtIdx = (function() {
    for( let z = 0; z<fields.length; z++ ) {
      let field = fields[ z ];
      if( field.name=='deletedAt' ) {
        return z;
      }
    }
    return -1;
  })();
  
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

<section>
  <slot></slot>
  <div class='pagination' style='float:right; display: inline-block'>
    <button title='Go to first page' disabled={!allowPrevPage} on:click={() => gotoPage(1)}>
      <i class='gg-push-chevron-left' />
    </button>
    <button title='Go to previous page' disabled={!allowPrevPage} on:click={() => gotoPage(pager.page-1)}>
      <i class='gg-chevron-left' />
    </button>
    <button title='Go to next page' disabled={!allowNextPage} on:click={() => gotoPage(pager.page+1)}>
      <i class='gg-chevron-right' />
    </button>
    <button title='Go to last page' disabled={!allowNextPage} on:click={() => gotoPage(pager.pages)}>
      <i class='gg-push-chevron-right' />
    </button>
  </div>
  <table class='table_users'>
    <thead>
    <tr>
      {#each fields as field}
        {#if field.name==='id'}
          <th class='table_header'>Action</th>
        {:else}
          <th class='table_header'>{field.label}</th>
        {/if}
      {/each}
    </tr>
    </thead>
    <tbody>
    {#each rows as row, no}
      <tr class:deleted={row[deletedAtIdx]>0}>
        {#each fields as field, i}
          {#if field.name==='id'}
            <td class='table_data' title='Edit user'>
              <button class='action' on:click={() => editRow(row[i])}>
                <i class='gg-pen' />
              </button>
              <!-- {no + 1} -->
            </td>
          {:else if field.inputType==='checkbox'}
            <td class='table_data'>{!!row[ i ]}</td>
          {:else if field.inputType==='datetime'}
            <td class='table_data'>{datetime( row[ i ] )}</td>
          {:else if field.inputType==='number'}
            <td>{(row[ i ] || 0).toLocaleString()}</td>
          {:else}
            <td class='table_data'>{row[ i ]}</td>
          {/if}
        {/each}
      </tr>
    {/each}
    </tbody>
  </table>
  
  <div class='pages_set'>
    <div class='page_and_rows_count'>
      <span>Page {pager.page} of {pager.pages},</span>
      <input id='perPage' class='perPage' type='number' min='0' bind:value={pager.perPage} on:change={() => changePerPage(pager.perPage)} />
      <span>rows per page.</span>
    </div>
    
    <p>Total: {pager.countResult}</p>
    
    <div class='pagination'>
      <button title='Go to first page' disabled={!allowPrevPage} on:click={() => gotoPage(1)}>
        <i class='gg-push-chevron-left' />
      </button>
      <button title='Go to previous page' disabled={!allowPrevPage} on:click={() => gotoPage(pager.page-1)}>
        <i class='gg-chevron-left' />
      </button>
      <button title='Go to next page' disabled={!allowNextPage} on:click={() => gotoPage(pager.page+1)}>
        <i class='gg-chevron-right' />
      </button>
      <button title='Go to last page' disabled={!allowNextPage} on:click={() => gotoPage(pager.pages)}>
        <i class='gg-push-chevron-right' />
      </button>
    </div>
  </div>
</section>

<style>
    .table_users {
        margin-top       : 10px;
        background-color : white;
        border-collapse  : collapse;
        font-size        : 14px;
        width            : 100%;
        color            : #475569;
    }

    .table_users .table_header {
        text-align : left !important;
        color      : #6366F1 !important;
    }

    .table_users .table_header {
        border-bottom : 1px solid #CBD5E1;
    }

    .table_users .table_header, .table_users .table_data {
        padding : 12px 5px;
    }

    .table_users .table_data .action {
        border     : none;
        background : none;
        color      : #475569;
        text-align : center;
        margin     : auto;
        cursor     : pointer;
    }

    .table_users .table_data .action:hover {
        color : #EF4444;
    }

    .table_users .deleted td {
        text-decoration       : line-through;
        text-decoration-color : #EF4444;
    }

    .pages_set {
        margin-top      : 5px;
        padding-top     : 10px;
        border-top      : 1px solid #CBD5E1;
        display         : flex;
        flex-direction  : row;
        align-items     : center;
        justify-content : space-between;
        align-content   : center;
        font-size       : 15px;
        color           : #161616;
    }

    .pages_set .page_and_rows_count {
        display        : flex;
        flex-direction : row;
        align-content  : center;
        align-items    : center;
    }

    .pages_set .page_and_rows_count .perPage {
        margin        : auto 5px;
        width         : 4em;
        border        : 1px solid #CBD5E1;
        border-radius : 4px;
        padding       : 5px;
        font-size     : 14px;
        text-align    : center;
        color         : #161616;
        outline-color : #6366F1;
    }

    .pages_set .page_and_rows_count .perPage::-webkit-inner-spin-button, .pages_set .page_and_rows_count .perPage::-webkit-outer-spin-button {
        opacity : 1;
    }

    .pagination button {
        color            : white;
        background-color : #6366F1;
        padding          : 3px 7px;
        border-radius    : 5px;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
        margin-left      : 4px;
        cursor           : pointer;
        border           : none;
    }

    .pagination button:hover {
        background-color : #7E80F1;
    }

    .pagination button:disabled {
        border     : 1px solid #CBD5E1 !important;
        background : none !important;
        color      : #5C646F;
        filter     : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    }

    .pagination button:disabled:hover {
        background : none;
    }
</style>
