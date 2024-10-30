<script>
  import { onMount } from 'svelte';
  import { datetime } from './formatter.js';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { BiPencil } from '../node_modules/svelte-icons-pack/dist/bi';
  import {
    FaSolidAngleRight, FaSolidAngleLeft,
    FaSolidAnglesRight, FaSolidFilter, FaSolidArrowsRotate,
  } from '../node_modules/svelte-icons-pack/dist/fa';
  
  export let renderFuncs = {};
  export let arrayOfArray = true;
  export let fields = []; // array of field object
  export let rows = []; // 2 dimension array or array of object if arrayOfArray = false
  export let pager = {}; // pagination
  export let extraActions = []; // array of object {icon, label, onClick}
  export let onRefreshTableView = function( pager ) {
    console.log( 'TableView.onRefreshTableView', pager );
  };
  export let onEditRow = function( id, row ) {
    console.log( 'TableView.onEditRow', id, row );
  };
  export let widths = {}; // array of key and css width
  
  let deletedAtIdx = -1;
  onMount( () => {
    console.log( 'onMount.TableView' );
    console.log( 'fields=', fields );
    console.log( 'rows=', rows );
    console.log( 'pager=', pager );
    console.log( 'extraActions=', extraActions );
    
    for( let z = 0; z<fields.length; z++ ) {
      let field = fields[ z ];
      // find deletedAt index
      if( field.name==='deletedAt' ) {
        console.log( 'deletedAtIdx', z );
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
  
  let filtersMap = {};
  
  function filterKeyDown( event ) {
    if( event.key==='Enter' ) applyFilter();
  }
  
  function applyFilter() {
    let filters = {};
    for( let key in filtersMap ) {
      let value = filtersMap[ key ];
      if( value ) {
        filters[ key ] = value.split( '|' );
      }
    }
    console.log( 'filters=', filters );
    onRefreshTableView( {
      ...pager,
      filters: filters,
    } );
    oldFilterStr = newFilterStr;
  }
  
  function gotoPage( page ) {
    onRefreshTableView( {
      ...pager,
      page,
    } );
  }
  
  function changePerPage( perPage ) {
    onRefreshTableView( {
      ...pager,
      perPage,
    } );
  }
  
  function cell( row, i, field ) {
    if( arrayOfArray ) return row[ i ] || '';
    return row[ field.name ] || '';
  }
  
  $: allowPrevPage = pager.page>1;
  $: allowNextPage = pager.page<pager.pages;

  const tableTitle = `separate with pipe for multiple values, for example:
  >=100|<50|61|72 will show values greater equal to 100, OR less than 50, OR exactly 61 OR 72
    filtering with greater or less than will only show correct result if data is saved as number
    currently price and size NOT stored as number
  <>abc* will show values NOT started with abc*
  abc*def|ghi will show values started with abc ends with def OR exactly ghi
  *jkl* will show values containing jkl substring
multiple filter from other fields will do AND operation`
</script>

<section class="table-root">
  <div class="actions-container">
    <div class="left">
      <div class="actions-button">
        <slot />
        <button class="btn" disabled={oldFilterStr===newFilterStr} on:click={applyFilter}>
          <Icon
            color={oldFilterStr === newFilterStr ? '#5C646F' : '#FFF'}
            size="17"
            src={FaSolidFilter}
          />
          <span>Apply Filter</span>
        </button>
        <button class="btn" on:click={() => gotoPage(pager.page)}>
          <Icon
            color="#FFF"
            size={17}
            src={FaSolidArrowsRotate}
          />
          <span>Refresh</span>
        </button>
      </div>
    </div>
  </div>
  <div class="table-container">
    <table>
      <thead>
        <tr>
          {#each (fields || []) as field}
            {#if field.name==='id'}
              <th class='a-row'>Action</th>
            {:else}
              <th
                class="
                {field.inputType === 'textarea' ? 'textarea' : ''}
                {field.inputType === 'datetime' ? 'datetime' : ''}
              ">{field.label}
              </th>
            {/if}
          {/each}
        </tr>
      </thead>
      <tbody>
        {#each (rows || []) as row, idx}
          <tr class:deleted={row[deletedAtIdx] > 0}>
            {#each fields as field, i}
              {#if field.name === 'id'}
                <td class="a-row">
                  <div class="actions">
                    <button class="btn" title="Edit" on:click={() => onEditRow(cell(row,i,field), row)}>
                      <Icon src={BiPencil} size="15" />
                    </button>
                    {#each extraActions as action}
                      {#if action.link}
                        <a href='{action.link(row)}' class="action" target="_blank" title={action.label || ''}>
                          <Icon src={action.icon} />
                        </a>
                      {:else}
                        <button class="action" title='{action.label || ""}' on:click={() => action.onClick(row)}>
                          <Icon src={action.icon} />
                        </button>
                      {/if}
                    {/each}
                  </div>
                </td>
              {:else if renderFuncs[ field.name ]}
                <td class='table-data'>{renderFuncs[ field.name ]( cell( row, i, field ) ) }</td>
              {:else if field.inputType==='checkbox'}
                <td class='table-data'>{!!cell( row, i, field )}</td>
              {:else if field.inputType==='datetime' || field.name==='deletedAt'}
                <td class='table-data'>{datetime( cell( row, i, field ) )}</td>
              {:else if field.inputType==='number'}
                <td>{(cell( row, i, field ) || 0).toLocaleString()}</td>
              {:else}
                <td class='table-data'>{cell( row, i, field )}</td>
              {/if}
            {/each}
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
  <div class='pages_set'>
    <div class='page_and_rows_count'>
      <span>Page {pager.page} of {pager.pages},</span>
      <input
        bind:value={pager.perPage}
        class='perPage'
        id='perPage'
        min='0'
        on:change={() => changePerPage(pager.perPage)}
        type='number'
      />
      <span>rows per page.</span>
    </div>
    
    <p>Total: {pager.countResult | 0}</p>
    
    <div class='pagination'>
      <button disabled={!allowPrevPage} on:click={() => gotoPage(1)} title='Go to first page'>
        <Icon color={!allowPrevPage ? '#5C646F' : '#FFF'} size={18} src={FaSolidAnglesRight} />
      </button>
      <button disabled={!allowPrevPage} on:click={() => gotoPage(pager.page - 1)} title='Go to previous page'>
        <Icon color={!allowPrevPage ? '#5C646F' : '#FFF'} size={18} src={FaSolidAngleLeft} />
      </button>
      <button disabled={!allowNextPage} on:click={() => gotoPage(pager.page + 1)} title='Go to next page'>
        <Icon color={!allowNextPage ? '#5C646F' : '#FFF'} size={18} src={FaSolidAngleRight} />
      </button>
      <button disabled={!allowNextPage} on:click={() => gotoPage(pager.pages)} title='Go to last page'>
        <Icon color={!allowNextPage ? '#5C646F' : '#FFF'} size={18} src={FaSolidAnglesRight} />
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

  .table-root .text-violet {
    color: var(--violet-005);
    font-weight: 600;
    padding: 5px;
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

  .table-root .actions-container .left,
  .table-root .actions-container .right {
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

  .table-root .actions-container .right .search_handler {
    display: flex;
    flex-direction: row;
    width: fit-content;
    height: fit-content;
    position: relative;
  }

  .table-root .actions-container .right .search_handler input.search {
    padding: 12px 40px 12px 15px;
    border-radius: 8px;
    border: none;
    background-color: var(--gray-001);
    width: 370px;
  }

  .table-root .actions-container .right .search_handler input.search:focus {
    border-color: none;
    outline: 1px solid var(--gray-003);
    box-shadow: var(--shadow-md);
  }

  .table-root .actions-container .right .search_handler .search_btn {
    position: absolute;
    background-color: transparent;
    padding: 8px;
    display: flex;
    justify-content: center;
    align-items: center;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    right: 5px;
    top: 3px;
  }

  .table-root .actions-container .right .search_handler .search_btn:hover {
    background-color: var(--violet-transparent);
  }

  :global(.table-root .actions-container .right .search_handler .search_btn:hover svg) {
    fill: var(--violet-005);
  }

  .table-root .actions-container .actions-button {
    display: flex;
    flex-direction: row;
    align-items: center;
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
  }

  .table-root .table-container table thead {
    box-shadow: none;
    border-bottom: 1px solid var(--gray-003);
  }

  .table-root .table-container table thead tr th {
    padding: 12px;
		background-color: var(--gray-001);
		text-transform: capitalize;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-003);
		min-width: fit-content;
		width: auto;
    text-wrap: nowrap;
  }

  .table-root .table-container table thead tr th.textarea {
    min-width: 280px !important;
  }

  .table-root .table-container table thead tr th.datetime {
    min-width: 140px !important;
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
  }

	.table-root .table-container table tbody tr td {
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
  }

  .table-root .table-container table tbody tr td .actions .btn:hover {
    background-color: var(--violet-transparent);
  }

  :global(.table-root .table-container table tbody tr td .actions .btn:hover svg) {
    fill: var(--violet-005);
  }

  .table-root .pagination-container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 15px 15px 0 15px;
  }

  .table-root .pagination-container .filter {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 8px;
  }

  .table-root .pagination-container .filter .row_to_show {
    position: relative;
    width: fit-content;
    height: fit-content;
  }

  .table-root .pagination-container .filter .row_to_show .btn {
    border: none;
    background-color: var(--violet-transparent);
    color: var(--violet-005);
    width: fit-content;
    padding: 3px 3px 3px 6px;
    font-weight: 600;
    border: 1px solid var(--violet-004);
    border-radius: 9999px;
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    gap: 1px;
    cursor: pointer;
  }

  .table-root .pagination-container .filter .row_to_show .btn:hover {
    background-color: var(--violet-002);
  }

  .table-root .pagination-container .filter .row_to_show .rows {
    display: flex;
    flex-direction: column-reverse;
    position: absolute;
    width: 100%;
    top: -180px;
    border-radius: 5px;
    border: 1px solid var(--gray-004);
    background-color: #fff;
  }

  .table-root .pagination-container .filter .row_to_show .rows button {
    border: none;
    background-color: transparent;
    padding: 5px;
    cursor: pointer;
    color: var(--gray-007);
  }

  .table-root .pagination-container .filter .row_to_show .rows button:hover {
    background-color: var(--violet-transparent);
    color: var(--violet-007);
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
    flex-direction: row;
    justify-content: center;
    align-items: center;
    padding: 6px 10px;
    border-radius: 9999px;
    cursor: pointer;
    gap: 5px;
    color: var(--gray-007);
    border: 1px solid transparent;
  }

  .table-root .pagination-container .pagination .btn:hover {
    border: 1px solid var(--gray-004);
  }

  .table-root .pagination-container .pagination .btn.active {
    background-color: var(--violet-transparent);
    color: var(--violet-006);
    font-weight: 600;
    border: 1px solid var(--violet-004);
  }

  .table-root .pagination-container .pagination .btn.to {
    background-color: var(--violet-006);
    color: #fff;
    font-weight: 600;
    border: none;
  }

  .table-root .pagination-container .pagination .btn.to:hover {
    background-color: var(--violet-005);
  }

  .table-root .pagination-container .pagination .btn.to:disabled {
    background-color: var(--gray-002);
    color: var(--gray-006);
    font-weight: 600;
    border: 1px solid var(--gray-004);
    cursor: not-allowed;
  }

  .pages_set {
    padding-top     : 10px;
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
    padding       : 5px;
    font-size     : 14pt;
    font-weight   : bold;
    text-align    : center;
    color         : #161616;
    outline-color : #6366F1;
  }

  .pages_set .page_and_rows_count .perPage::-webkit-inner-spin-button,
  .pages_set .page_and_rows_count .perPage::-webkit-outer-spin-button {
    opacity : 1;
  }

  .pagination button {
    color            : white;
    background-color : #6366F1;
    padding          : 8px;
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
    cursor     : not-allowed;
    border     : 1px solid #CBD5E1 !important;
    background : none !important;
    color      : #5C646F;
    filter     : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
  }

  .pagination button:disabled:hover {
    background : none;
  }

  input.input_filter {
    width      : 0;
    min-width  : 100%;
    box-sizing : border-box;
  }
</style>
