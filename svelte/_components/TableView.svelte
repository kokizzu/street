<script>
  /** @typedef {import('../_types/master.js').Field} Field */
  /** @typedef {import('../_types/master.js').PagerIn} PagerIn */
  /** @typedef {import('../_types/master.js').PagerOut} PagerOut */
  /** @typedef {import('../_types/master.js').ExtendedAction} ExtendedAction */

  import { onMount } from 'svelte';
  import { datetime } from './formatter.js';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import {
    RiSystemRefreshLine, RiSystemFilterLine, RiDesignPencilLine,
    RiArrowsArrowRightSLine, RiArrowsArrowRightDoubleFill,
    RiArrowsArrowLeftSLine, RiArrowsArrowLeftDoubleFill
  } from '../node_modules/svelte-icons-pack/dist/ri';
  import TableFilterInput from './TableFilterInput.svelte';
  
  export let renderFuncs  = /** @type {Record<string, Function>} */ ({});
  export let arrayOfArray = /** @type {boolean} */ (true);
  export let fields       = /** @type {Field[]} */ ([]);
  export let rows         = /** @type {any[] | Record<string, any>[]} */ ([]);
  export let pager        = /** @type {PagerOut} */ ({});
  export let extraActions = /** @type {ExtendedAction[]} */ ([]);

  export let onRefreshTableView = function(/** @type {PagerIn} */ pager ) {
    console.log( 'TableView.onRefreshTableView', pager );
  };
  export let onEditRow = function(/** @type {number | string} */ id, /** @type {any | any[]} */ row ) {
    console.log( 'TableView.onEditRow', id, row );
  };
  
  // Index of deletedAt field
  let deletedAtIdx = /** @type {number} */ (-1);

  onMount( () => {
    console.log( 'onMount.TableView' );
    console.log( 'fields=', fields );
    console.log( 'rows=', rows );
    console.log( 'pager=', pager );
    console.log( 'extraActions=', extraActions );
    
    for( let z = 0; z < fields.length; z++ ) {
      let field = fields[ z ];
      if( field.name==='deletedAt' ) {
        deletedAtIdx = z;
      }

      // empty all filter at beginning
      filtersMap[ field.name ] = '';
    }

    oldFilterStr = JSON.stringify( filtersMap );
  } );
  
  let oldFilterStr = '{}';
  let newFilterStr = '';

  $: newFilterStr = JSON.stringify( filtersMap );
  
  let filtersMap = /** @type {Record<string, string>} */ ({});
  
  // @deprecated
  // function filterKeyDown(/** @type {KeyboardEvent} */ event ) {
  //   if( event.key === 'Enter' ) applyFilter();
  // }
  
  function applyFilter() {
    let filters = /** @type {Record<string, string[]>}*/ ({});
    for( let key in filtersMap ) {
      let value = filtersMap[ key ];
      if( value ) {
        filters[ key ] = value.split( '|' );
      }
    }

    onRefreshTableView( {
      page: pager.page,
      perPage: pager.perPage,
      order: pager.order,
      filters: filters,
    } );

    oldFilterStr = newFilterStr;
  }
  
  function gotoPage(/** @type {number} */ page ) {
    onRefreshTableView({ ...pager, page });
  }
  
  function changePerPage(/** @type {number} */ perPage ) {
    onRefreshTableView({ ...pager, perPage});
  }
  
  function cell( row, i, field ) {
    if( arrayOfArray ) return row[ i ] || '';
    return row[ field.name ] || '';
  }
  
  $: allowPrevPage = pager.page>1;
  $: allowNextPage = pager.page<pager.pages;

 
</script>

<section class="table-root">
  <div class="actions-container">
    <div class="left">
      <div class="actions-button">
        <slot />
        <button class="btn"
          disabled={oldFilterStr===newFilterStr} on:click={applyFilter}
          title="Apply Filter"
        >
          <Icon
            color="var(--gray-008)"
            size="17"
            src={RiSystemFilterLine}
          />
        </button>
        <button class="btn"
          on:click={() => gotoPage(pager.page)}
          title="Refresh Table"
        >
          <Icon
            color="var(--gray-008)"
            size={17}
            src={RiSystemRefreshLine}
          />
        </button>
      </div>
    </div>
  </div>
  <div class="table-container">
    <table>
      <thead>
        <tr>
          <th class="no">No</th>
          {#each (fields || []) as field}
            {#if field.name==='id'}
              <th class='a-row'>Action</th>
            {:else}
              <th
                class="
                {field.inputType === 'textarea' ? 'textarea' : ''}
                {field.inputType === 'datetime' ? 'datetime' : ''}
                {field.name === 'fullName' ? 'full-name' : ''}
                {field.name === 'userAgent' ? 'user-agent' : ''}
              ">
               <TableFilterInput
                label={field.label}
                bind:value={filtersMap[ field.name ]}
               />
              </th>
            {/if}
          {/each}
        </tr>
      </thead>
      <tbody>
        {#if rows && rows.length > 0}
          {#each (rows || []) as row, idx}
            <tr class:deleted={row[deletedAtIdx] > 0}>
              <td class="num-row">{(pager.page -1) * pager.perPage + idx + 1}</td>
              {#each fields as field, i}
                {#if field.name === 'id'}
                  <td class="a-row">
                    <div class="actions">
                      <button class="btn" title="Edit" on:click={() => onEditRow(cell(row,i,field), row)}>
                        <Icon
                          src={RiDesignPencilLine}
                          size="17"
                          color="var(--gray-008)"
                        />
                      </button>
                      {#each extraActions as action}
                        {#if action.link}
                          <a href={action.link(row)} class="btn" target="_blank" title={action.label || ''}>
                            <Icon
                              src={action.icon}
                              size="17"
                              color="var(--gray-008)"
                            />
                          </a>
                        {:else}
                          <button class="btn" title={action.label || ""}
                            on:click={() => action.onClick(row)}
                          >
                            <Icon
                              src={action.icon}
                              size="17"
                              color="var(--gray-008)"
                            />
                          </button>
                        {/if}
                      {/each}
                    </div>
                  </td>
                {:else if renderFuncs[ field.name ]}
                  <td>{renderFuncs[ field.name ]( cell( row, i, field ) ) }</td>
                {:else if field.inputType==='checkbox'}
                  <td>{!!cell( row, i, field )}</td>
                {:else if field.inputType==='datetime' || field.name==='deletedAt'}
                  <td>{datetime( cell( row, i, field ) )}</td>
                {:else if field.inputType==='number'}
                  <td>{(cell( row, i, field ) || 0).toLocaleString()}</td>
                {:else}
                  <td>{cell( row, i, field )}</td>
                {/if}
              {/each}
            </tr>
          {/each}
        {:else}
          <tr>
            <td class="num-row">0</td>
            <td colspan={fields.length - 2}>No data</td>
          </tr>
        {/if}
      </tbody>
    </table>
  </div>
  <div class="pagination-container">
    <div class="pager-info">
      <span>Page {pager.page} of {pager.pages}</span>
      <input
        bind:value={pager.perPage}
        class="per-page"
        id="perPage"
        min="0"
        on:change={() => changePerPage(pager.perPage)}
        type="number"
      />
      <span>rows per page.</span>
    </div>
    <div class="total">
      <p>Total: {pager.countResult | 0}</p>
    </div>
    <div class="pagination">
      <button
        class="btn"
        disabled={!allowPrevPage}
        on:click={() => gotoPage(1)}
        title="Go to first page"
      >
        <Icon
          color="var(--gray-008)"
          size="18"
          src={RiArrowsArrowLeftDoubleFill}
        />
      </button>
      <button
        class="btn"
        disabled={!allowPrevPage}
        on:click={() => gotoPage(pager.page - 1)}
        title="Go to previous page"
      >
        <Icon
          color="var(--gray-008)"
          size="18"
          src={RiArrowsArrowLeftSLine}
        />
      </button>
      <button
        class="btn"
        disabled={!allowNextPage}
        on:click={() => gotoPage(pager.page + 1)}
        title="Go to next page"
      >
        <Icon
          color="var(--gray-008)"
          size="18"
          src={RiArrowsArrowRightSLine}
        />
      </button>
      <button
        class="btn"
        disabled={!allowNextPage}
        on:click={() => gotoPage(pager.pages)}
        title="Go to last page"
      >
        <Icon
          color="var(--gray-008)"
          size="18"
          src={RiArrowsArrowRightDoubleFill}
        />
      </button>
    </div>
  </div>
</section>

<style>
  .table-root {
    display: flex;
    flex-direction: column;
    background-color: #fff;
    border-radius: 10px;
    border: 1px solid var(--gray-003);
    padding: 0 0 20px 0;
    overflow: hidden;
  }

  .table-root p {
    margin: 0;
  }

  .table-root .actions-container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 10px 15px;
    background-color: #fff;
  }

  .table-root .actions-container .left {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
  }

  .table-root .actions-container .left .debug .btn {
    border: none;
    background-color: var(--violet-006);
    color: #fff;
    width: fit-content;
    padding: 4px 10px;
    border-radius: 9999px;
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    gap: 3px;
    cursor: pointer;
  }

  .table-root .actions-container .left .debug .btn:hover {
    background-color: var(--violet-005);
  }

  :global(.table-root .actions-container .right .search_handler .search_btn:hover svg) {
    fill: var(--violet-005);
  }

  .table-root .actions-container .actions-button {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 5px;
  }

  .table-root .table-container {
    overflow-x: auto;
    scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: thin;
  }

  .table-root .table-container table {
    width: 100%;
    background: #fff;
    border-top: 1px solid var(--gray-003);
    border-bottom: 1px solid var(--gray-003);
    box-shadow: none;
    text-align: left;
    border-collapse: separate;
    border-spacing: 0;
    overflow: hidden;
    font-size: 13px;
  }

  .table-root .table-container table thead {
    box-shadow: none;
    border-bottom: 1px solid var(--gray-003);
  }

  .table-root .table-container table thead tr th {
    padding: 5px;
		background-color: var(--gray-001);
		text-transform: capitalize;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-003);
		min-width: fit-content;
		width: auto;
    text-wrap: nowrap;
  }

  .table-root .table-container table thead tr th:nth-child(1),
  .table-root .table-container table thead tr th:nth-child(2) {
    padding: 5px 12px !important;
  }

  .table-root .table-container table thead tr th.textarea {
    min-width: 280px !important;
  }

  .table-root .table-container table thead tr th.datetime,
  .table-root .table-container table thead tr th.full-name {
    min-width: 140px !important;
  }

  .table-root .table-container table thead tr th.user-agent {
    min-width: 300px !important;
  }

  .table-root .table-container table tbody tr.deleted {
    color: var(--red-005);
  }

  .table-root .table-container table thead tr th.no {
    width: 30px;
  }

  .table-root .table-container table thead tr th.a-row {
    max-width: fit-content;
    min-width: fit-content;
    width: fit-content;
  }

  .table-root .table-container table thead tr th:last-child {
    border-right: none;
  }

  .table-root .table-container table tbody tr td {
    padding: 8px 12px;
    white-space: nowrap;
    padding: 8px 12px;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-004);
  }

	.table-root .table-container table tbody tr:last-child td,
	.table-root .table-container table tbody tr:last-child th {
		border-bottom: none !important;
	}

  .table-root .table-container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

	.table-root .table-container table tbody tr td.num-row {
		border-right: 1px solid var(--gray-003);
		font-weight: 600;
		text-align: center;
	}

  .table-root .table-container table tbody tr:last-child td,
  .table-root .table-container table tbody tr:last-child th {
    border-bottom: none !important;
  }

  .table-root .table-container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .table-root .table-container table tbody tr td:last-child {
    border-right: none !important;
  }

  .table-root .table-container table tbody tr th {
    text-align: center;
    border-right: 1px solid var(--gray-004);
    border-bottom: 1px solid var(--gray-004);
  }

  .table-root .table-container table tbody tr td .actions {
    display: flex;
    flex-direction: row;
  }

  .table-root .table-container table tbody tr td .actions .btn {
    border: none;
    padding: 6px;
    border-radius: 8px;
    background-color: transparent;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
    text-decoration: none;
  }

  .table-root .table-container table tbody tr td .actions .btn:hover {
    background-color: var(--orange-transparent);
  }

  :global(.table-root .table-container table tbody tr td .actions .btn:hover svg) {
    fill: var(--orange-005);
  }

  .table-root .pagination-container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 15px 15px 0 15px;
  }

  .table-root .pagination-container .pager-info {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 5px;
  }

  .table-root .pagination-container .pager-info input.per-page {
    padding: 5px 8px;
    border: 1px solid var(--gray-003);
    border-radius: 8px;
    width: 55px;
  }

  .table-root .pagination-container .pager-info input.per-page:focus {
    border-color: var(--orange-005);
    outline: 1px solid var(--orange-005);
  }

  .table-root .pagination-container .pagination {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 5px;
    overflow: hidden;
  }

  .table-root .pagination-container .pagination .btn {
    border: none;
    background-color: transparent;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 5px 8px;
    border-radius: 8px;
    cursor: pointer;
    border: 1px solid transparent;
    text-decoration: none;
  }

  .table-root .pagination-container .pagination .btn:hover {
    background-color: var(--orange-transparent);
    color: var(--orange-005);
  }

  :global(.table-root .pagination-container .pagination .btn:hover svg) {
    fill: var(--orange-005);
    outline: var(--orange-005);
  }

  :global(.table-root .pagination-container .pagination .btn:disabled:hover svg) {
    fill: var(--gray-008);
    outline: var(--gray-008);
  }

  .table-root .pagination-container .pagination .btn:disabled {
    background-color: var(--gray-002);
    font-weight: 600;
    cursor: default;
  }
</style>
