<script>
    export let visible = false;
    export let close_label = 'Close';
    export let title = '';

    export function showModal() {
        visible = true;
    }

    export function hideModal() {
        visible = false;
    }
</script>

{#if visible}
    <div class='backdrop'>
        <div class='modal_dialog'>
            <div class='modal_content'>
                {#if title}
                    <h3>{title}</h3>
                {/if}
                <slot name='content' />
            </div>
            <div class='buttons'>
                <button
                        on:click|preventDefault={() => hideModal()}
                        class='cancel_button'>{close_label}
                </button>
                <slot name='buttons' />
            </div>
        </div>
    </div>
{/if}

<style>
    .backdrop {
        position        : absolute;
        z-index         : 40;
        background      : rgba(41, 41, 41, 0.5);
        width           : 100%;
        left            : 0;
        top             : 0;
        bottom          : 0;
        height          : auto;
        border-radius   : 8px;
        display         : flex;
        justify-content : center;
        align-items     : center;
    }

    .modal_dialog {
        color            : #334155;
        width            : 400px;
        height           : fit-content;
        gap              : 40px;
        background-color : white;
        padding          : 20px;
        border-radius    : 8px;
        display          : flex;
        flex-direction   : column;
        justify-content  : space-between;
    }

    .modal_dialog h3 {
        font-weight : 600;
        margin      : 0;
        font-size   : 18px;
        width       : 100%;
        text-align  : center;
    }

    .modal_dialog .buttons {
        display        : flex;
        flex-direction : row;
        align-items    : stretch;
        font-weight    : 500;
        width          : 100%;
    }

    .modal_dialog .buttons button {
        border-radius  : 8px;
        border         : none;
        padding        : 10px;
        cursor         : pointer;
        width          : 100%;
        text-transform : capitalize;
    }

    .modal_dialog .buttons .cancel_button {
        margin-right     : 10px;
        border           : 1px solid #CBD5E1;
        background-color : #F1F5F9;
        color            : #F97316;
    }

    .modal_dialog .buttons .cancel_button:hover {
        border : 1px solid #F97316;
    }
</style>