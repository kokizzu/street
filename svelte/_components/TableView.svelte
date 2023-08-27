<script>
    // @ts-nocheck
    import { onMount } from 'svelte';
    import { datetime } from './formatter.js';

    import Icon from 'svelte-icons-pack/Icon.svelte';
    import HiOutlinePencil from 'svelte-icons-pack/hi/HiOutlinePencil';
    import FaSolidAngleRight from 'svelte-icons-pack/fa/FaSolidAngleRight';
    import FaSolidAngleLeft from 'svelte-icons-pack/fa/FaSolidAngleLeft';
    import FaSolidAngleDoubleRight from 'svelte-icons-pack/fa/FaSolidAngleDoubleRight';
    import FaSolidAngleDoubleLeft from 'svelte-icons-pack/fa/FaSolidAngleDoubleLeft';
    import FaSolidFilter from 'svelte-icons-pack/fa/FaSolidFilter';

    export let fields = []; // array of field object
    export let rows = []; // 2 dimension array
    export let pager = {}; // pagination
    export let extraActions = []; // array of object {icon, label, onClick}
    export let onRefreshTableView = function(pager) {
        console.log('TableView.onRefreshTableView', pager);
    };
    export let onEditRow = function(id, row) {
        console.log('TableView.onEditRow', id, row);
    };
    export let widths = {}; // array of key and css width

    let deletedAtIdx = -1;
    onMount(() => {
        console.log('fields=', fields);
        console.log('rows=', rows);
        console.log('pager=', pager);
        console.log('extraActions=', extraActions);

        for (let z = 0; z < fields.length; z++) {
            let field = fields[z];
            // find deletedAt index
            if (field.name === 'deletedAt') {
                console.log('deletedAtIdx', z);
                deletedAtIdx = z;
            }
            // empty all filter at beginning
            filtersMap[field.name] = '';
        }
        oldFilterStr = JSON.stringify(filtersMap);
    });

    let oldFilterStr = '{}';
    let newFilterStr = '';
    $: newFilterStr = JSON.stringify(filtersMap);

    let filtersMap = {};

    function filterKeyDown(event) {
        if (event.key === 'Enter') applyFilter();
    }

    function applyFilter() {
        let filters = {};
        for (let key in filtersMap) {
            let value = filtersMap[key];
            if (value) {
                filters[key] = value.split('|');
            }
        }
        console.log('filters=', filters);
        onRefreshTableView({
            ...pager,
            filters: filters,
        });
        oldFilterStr = newFilterStr;
    }

    function gotoPage(page) {
        onRefreshTableView({
            ...pager,
            page,
        });
    }

    function changePerPage(perPage) {
        onRefreshTableView({
            ...pager,
            perPage,
        });
    }

    $: allowPrevPage = pager.page > 1;
    $: allowNextPage = pager.page < pager.pages;
</script>

<section>
    <slot />
    <button class='apply_filter_button' disabled={oldFilterStr===newFilterStr} onclick={applyFilter}>
        <Icon size={18} color={oldFilterStr === newFilterStr ? '#5C646F' : '#FFFF'} src={FaSolidFilter} />
        <span>Apply Filter</span>
    </button>
    <div class='pagination' style='float:right; display: inline-block'>
        <button title='Go to first page' disabled={!allowPrevPage} on:click={() => gotoPage(1)}>
            <Icon size={18} color={!allowPrevPage ? '#5C646F' : '#FFFF'} src={FaSolidAngleDoubleLeft} />
        </button>
        <button title='Go to previous page' disabled={!allowPrevPage} on:click={() => gotoPage(pager.page - 1)}>
            <Icon size={18} color={!allowPrevPage ? '#5C646F' : '#FFFF'} src={FaSolidAngleLeft} />
        </button>
        <button title='Go to next page' disabled={!allowNextPage} on:click={() => gotoPage(pager.page + 1)}>
            <Icon size={18} color={!allowNextPage ? '#5C646F' : '#FFFF'} src={FaSolidAngleRight} />
        </button>
        <button title='Go to last page' disabled={!allowNextPage} on:click={() => gotoPage(pager.pages)}>
            <Icon size={18} color={!allowNextPage ? '#5C646F' : '#FFFF'} src={FaSolidAngleDoubleRight} />
        </button>
    </div>
    <div class='table_container'>
        <table class='table_users'>
            <thead>
            <tr>
                {#each fields as field}
                    {#if field.name === 'id'}
                        <th class='col_action'>Action</th>
                    {:else}
                        <th class='table_header'
                        style='{widths[field.name] ? "min-width: "+widths[field.name] : ""}'>
                            <label for='th_{field.name}'>{field.label}</label><br />
                            <input id='th_{field.name}'
                                   title='separate with pipe for multiple values, for example:
  >=100|<50|61|72 will show values greater equal to 100, OR less than 50, OR exactly 61 OR 72
    filtering with greater or less than will only show correct result if data is saved as number
    currently price and size NOT stored as number
  <>abc* will show values NOT started with abc*
  abc*def|ghi will show values started with abc ends with def OR exactly ghi
  *jkl* will show values containing jkl substring
multiple filter from other fields will do AND operation'
                                   type='text'
                                   style='width: 0; min-width: 100%; box-sizing: border-box;'
                                   bind:value={filtersMap[field.name]}
                                   on:keydown={filterKeyDown}
                            />
                        </th>
                    {/if}
                {/each}
            </tr>
            </thead>
            <tbody>
            {#each rows as row, no}
                <tr class:deleted={row[deletedAtIdx] > 0}>
                    {#each fields as field, i}
                        {#if field.name === 'id'}
                            <td class='col_action'>
                                <button class='action' title='Edit user' on:click={() => onEditRow(row[i], row)}>
                                    <Icon src={HiOutlinePencil} />
                                </button>
                                {#each extraActions as action}
                                    <button class='action' title='{action.label || ""}' on:click={() => action.onClick(row)}>
                                        <Icon src={action.icon} />
                                    </button>
                                {/each}
                            </td>
                        {:else if field.inputType === 'checkbox'}
                            <td class='table_data'>{!!row[i]}</td>
                        {:else if field.inputType === 'datetime'}
                            <td class='table_data'>{datetime(row[i])}</td>
                        {:else if field.inputType === 'number'}
                            <td>{(row[i] || 0).toLocaleString()}</td>
                        {:else}
                            <td class='table_data'>{row[i]}</td>
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
                    id='perPage'
                    class='perPage'
                    type='number'
                    min='0'
                    bind:value={pager.perPage}
                    on:change={() => changePerPage(pager.perPage)}
            />
            <span>rows per page.</span>
        </div>

        <p>Total: {pager.countResult}</p>

        <div class='pagination'>
            <button title='Go to first page' disabled={!allowPrevPage} on:click={() => gotoPage(1)}>
                <Icon size={18} color={!allowPrevPage ? '#5C646F' : '#FFF'} src={FaSolidAngleDoubleLeft} />
            </button>
            <button title='Go to previous page' disabled={!allowPrevPage} on:click={() => gotoPage(pager.page - 1)}>
                <Icon size={18} color={!allowPrevPage ? '#5C646F' : '#FFF'} src={FaSolidAngleLeft} />
            </button>
            <button title='Go to next page' disabled={!allowNextPage} on:click={() => gotoPage(pager.page + 1)}>
                <Icon size={18} color={!allowNextPage ? '#5C646F' : '#FFF'} src={FaSolidAngleRight} />
            </button>
            <button title='Go to last page' disabled={!allowNextPage} on:click={() => gotoPage(pager.pages)}>
                <Icon size={18} color={!allowNextPage ? '#5C646F' : '#FFF'} src={FaSolidAngleDoubleRight} />
            </button>
        </div>
    </div>
</section>

<style>
    .table_container {
        overflow-x : auto;
    }

    .table_container::-webkit-scrollbar-thumb {
        background-color : #EF4444;
        border-radius    : 4px;
    }

    .table_container::-webkit-scrollbar-thumb:hover {
        background-color : #EC6262;
    }

    .table_container::-webkit-scrollbar {
        height : 10px;
    }

    .table_container::-webkit-scrollbar-track {
        background-color : transparent;
    }

    .table_container .table_users {
        margin           : 10px 1px 10px 0;
        background-color : white;
        border-collapse  : collapse;
        font-size        : 14px;
        width            : 100%;
        color            : #475569;
    }

    .table_container .table_users th {
        color   : #6366F1;
        border  : 1px solid #CBD5E1;
        padding : 10px 7px;
    }

    .table_container .table_users td {
        border  : 1px solid #CBD5E1;
        padding : 2px 3px 2px 3px;
    }

    .table_users .table_header {
        text-align  : left;
        white-space : nowrap;
    }

    .table_users .col_action {
        text-align : center;
    }

    .table_users td.col_action {
        text-align : center;
        display    : table-cell;
        padding    : 0
    }

    tr, td {
        height: 2em;
    }

    .table_users .col_action .action {
        background    : none;
        color         : #475569;
        text-align    : center;
        cursor        : pointer;
        height        : 2.5em;
        width         : 2em;
        border        : none;
        display       : inline;
        border-radius : 5px;
    }

    .table_users .col_action .action:hover {
        color            : #EF4444;
        background-color : #F0F0F0;
    }

    .table_users .deleted td {
        text-decoration           : line-through;
        text-decoration-color     : rgba(239, 68, 68, 0.5);
        text-decoration-thickness : 5px;
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

    .pagination button, .apply_filter_button {
        color            : white;
        background-color : #6366F1;
        padding          : 8px;
        border-radius    : 5px;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
        margin-left      : 4px;
        cursor           : pointer;
        border           : none;
    }

    .pagination button:hover, .apply_filter_button:hover {
        background-color : #7E80F1;
    }

    .apply_filter_button {
        padding          : 8px 20px;
        font-size        : 14pt;
        font-weight      : bold;
        display          : inline-flex;
        flex-direction   : row;
        align-items      : center;
        justify-content  : center;
        border           : none;
        background-color : #6366F1;
        border-radius    : 8px;
        color            : white;
        cursor           : pointer;
        gap              : 8px;
    }

    .pagination button:disabled, .apply_filter_button:disabled {
        border     : 1px solid #CBD5E1 !important;
        background : none !important;
        color      : #5C646F;
        filter     : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    }

    .pagination button:disabled:hover {
        background : none;
    }
</style>
