<script>
  import { createEventDispatcher } from 'svelte';
  import { datetime } from './formatter.js';

  export let fields = []; // array of field object
  export let rows = []; // 2 dimension array
  export let pager = {}; // pagination

  $: deletedAtIdx = (function () {
    for (let z = 0; z < fields.length; z++) {
      let field = fields[z];
      if (field.name == 'deletedAt') {
        return z;
      }
    }
    return -1;
  })();

  const dispatch = createEventDispatcher();

  function gotoPage(page) {
    dispatch('refreshTableView', {
      ...pager,
      page,
    });
  }

  function changePerPage(perPage) {
    dispatch('refreshTableView', {
      ...pager,
      perPage,
    });
  }

  function editRow(row) {
    dispatch('editRow', row);
  }

  $: allowPrevPage = pager.page > 1;
  $: allowNextPage = pager.page < pager.pages;
</script>

<section>
  <slot />
  <div class="pagination" style="float:right; display: inline-block">
    <button title="Go to first page" disabled={!allowPrevPage} on:click={() => gotoPage(1)}>
      <i class="gg-push-chevron-left" />
    </button>
    <button title="Go to previous page" disabled={!allowPrevPage} on:click={() => gotoPage(pager.page - 1)}>
      <i class="gg-chevron-left" />
    </button>
    <button title="Go to next page" disabled={!allowNextPage} on:click={() => gotoPage(pager.page + 1)}>
      <i class="gg-chevron-right" />
    </button>
    <button title="Go to last page" disabled={!allowNextPage} on:click={() => gotoPage(pager.pages)}>
      <i class="gg-push-chevron-right" />
    </button>
  </div>
  <div class="table_container">
    <table class="table_users">
      <thead>
        <tr>
          {#each fields as field}
            {#if field.name === 'id'}
              <th class="col_action">Action</th>
            {:else}
              <th class="table_header">{field.label}</th>
            {/if}
          {/each}
        </tr>
      </thead>
      <tbody>
        {#each rows as row, no}
          <tr class:deleted={row[deletedAtIdx] > 0}>
            {#each fields as field, i}
              {#if field.name === 'id'}
                <td class="col_action" title="Edit user">
                  <button class="action" on:click={() => editRow(row[i])}>
                    <i class="gg-pen" />
                  </button>
                  <!-- {no + 1} -->
                </td>
              {:else if field.inputType === 'checkbox'}
                <td class="table_data">{!!row[i]}</td>
              {:else if field.inputType === 'datetime'}
                <td class="table_data">{datetime(row[i])}</td>
              {:else if field.inputType === 'number'}
                <td>{(row[i] || 0).toLocaleString()}</td>
              {:else}
                <td class="table_data">{row[i]}</td>
              {/if}
            {/each}
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
  <div class="pages_set">
    <div class="page_and_rows_count">
      <span>Page {pager.page} of {pager.pages},</span>
      <input
        id="perPage"
        class="perPage"
        type="number"
        min="0"
        bind:value={pager.perPage}
        on:change={() => changePerPage(pager.perPage)}
      />
      <span>rows per page.</span>
    </div>

    <p>Total: {pager.countResult}</p>

    <div class="pagination">
      <button title="Go to first page" disabled={!allowPrevPage} on:click={() => gotoPage(1)}>
        <i class="gg-push-chevron-left" />
      </button>
      <button title="Go to previous page" disabled={!allowPrevPage} on:click={() => gotoPage(pager.page - 1)}>
        <i class="gg-chevron-left" />
      </button>
      <button title="Go to next page" disabled={!allowNextPage} on:click={() => gotoPage(pager.page + 1)}>
        <i class="gg-chevron-right" />
      </button>
      <button title="Go to last page" disabled={!allowNextPage} on:click={() => gotoPage(pager.pages)}>
        <i class="gg-push-chevron-right" />
      </button>
    </div>
  </div>
</section>

<style>
  .table_container {
    overflow-x: auto;
  }
  .table_container::-webkit-scrollbar-thumb {
    background-color : #EF4444;
    border-radius    : 4px;
  }
  .table_container::-webkit-scrollbar-thumb:hover {
    background-color: #ec6262;
  }

  .table_container::-webkit-scrollbar {
    height : 10px;
  }

  .table_container::-webkit-scrollbar-track {
    background-color : transparent;
  }
  .table_container .table_users {
    margin: 10px 1px 10px 0;
    background-color: white;
    border-collapse: collapse;
    font-size: 14px;
    width: 100%;
    color: #475569;
  }
  .table_container .table_users th {
    color: #6366f1;
    border: 1px solid #cbd5e1;
    padding: 12px;
  }
  .table_container .table_users td {
    border: 1px solid #cbd5e1;
    padding: 12px;
  }

  .table_users .table_header {
    text-align: left;
    white-space: nowrap;
  }

  .table_users .col_action {
    width: fit-content;
    max-width: fit-content;
  }

  .table_users .table_header,
  .table_users .table_data {
    min-width: 180px;
  }

  .table_users .col_action .action {
    border: none;
    background: none;
    color: #475569;
    text-align: center;
    margin: auto;
    cursor: pointer;
    padding: 18px 15px;
    border-radius: 8px;
  }

  .table_users .col_action .action:hover {
    color: #ef4444;
    background-color: #f0f0f0;
  }

  .table_users .deleted td {
    text-decoration: line-through;
    text-decoration-color: #ef4444;
  }

  .pages_set {
    padding-top: 10px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    align-content: center;
    font-size: 15px;
    color: #161616;
  }

  .pages_set .page_and_rows_count {
    display: flex;
    flex-direction: row;
    align-content: center;
    align-items: center;
  }

  .pages_set .page_and_rows_count .perPage {
    margin: auto 5px;
    width: 4em;
    border: 1px solid #cbd5e1;
    border-radius: 4px;
    padding: 5px;
    font-size: 14px;
    text-align: center;
    color: #161616;
    outline-color: #6366f1;
  }

  .pages_set .page_and_rows_count .perPage::-webkit-inner-spin-button,
  .pages_set .page_and_rows_count .perPage::-webkit-outer-spin-button {
    opacity: 1;
  }

  .pagination button {
    color: white;
    background-color: #6366f1;
    padding: 3px 7px;
    border-radius: 5px;
    filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    margin-left: 4px;
    cursor: pointer;
    border: none;
  }

  .pagination button:hover {
    background-color: #7e80f1;
  }

  .pagination button:disabled {
    border: 1px solid #cbd5e1 !important;
    background: none !important;
    color: #5c646f;
    filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
  }

  .pagination button:disabled:hover {
    background: none;
  }
</style>
