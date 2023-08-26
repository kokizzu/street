<script>
    // @ts-nocheck
    import { createEventDispatcher } from 'svelte';
    import { datetime } from './formatter.js';

    import Icon from 'svelte-icons-pack/Icon.svelte';
    import HiOutlinePencil from 'svelte-icons-pack/hi/HiOutlinePencil';
    import FaSolidAngleRight from 'svelte-icons-pack/fa/FaSolidAngleRight';
    import FaSolidAngleLeft from 'svelte-icons-pack/fa/FaSolidAngleLeft';
    import FaSolidAngleDoubleRight from 'svelte-icons-pack/fa/FaSolidAngleDoubleRight';
    import FaSolidAngleDoubleLeft from 'svelte-icons-pack/fa/FaSolidAngleDoubleLeft';

    export let fields = []; // array of field object
    export let rows = []; // 2 dimension array
    export let pager = {}; // pagination
    export let extraActions = []; // array of object {icon, label, onClick}

    $: deletedAtIdx = (function() {
        for (let z = 0; z < fields.length; z++) {
            let field = fields[z];
            if (field.name === 'deletedAt') {
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
                        <th class='table_header'>{field.label}</th>
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
                                <button class='action' title='Edit user' on:click={() => editRow(row[i])}>
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

    .table_users .table_header,
    .table_users .table_data {
        min-width : 180px;
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
        text-decoration       : line-through;
        text-decoration-color : #EF4444;
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
        font-size     : 14px;
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
        border     : 1px solid #CBD5E1 !important;
        background : none !important;
        color      : #5C646F;
        filter     : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    }

    .pagination button:disabled:hover {
        background : none;
    }
</style>
