<script>
    // @ts-nocheck
    import Icon from 'svelte-icons-pack/Icon.svelte';
    import FaSolidInfoCircle from 'svelte-icons-pack/fa/FaSolidInfoCircle'; // info
    import FaSolidCheckCircle from 'svelte-icons-pack/fa/FaSolidCheckCircle'; // success
    import FaSolidExclamationTriangle from 'svelte-icons-pack/fa/FaSolidExclamationTriangle'; // warning
    import FaSolidTimesCircle from 'svelte-icons-pack/fa/FaSolidTimesCircle'; // error

    let message = 'This is growl, XD';
    let growlType = 'info'; // info, warning, error, success

    let isShow = false;

    export function showInfo(msg) {
        growlType = 'info';
        isShow = true;
        message = msg;
        setTimeout(() => {
            isShow = false;
        }, 3000);
    }

    export function showWarning(msg) {
        growlType = 'warning';
        isShow = true;
        message = msg;
        setTimeout(() => {
            isShow = false;
        }, 3000);
    }

    export function showError(msg) {
        growlType = 'error';
        isShow = true;
        message = msg;
        setTimeout(() => {
            isShow = false;
        }, 3000);
    }

    export function showSuccess(msg) {
        growlType = 'success';
        isShow = true;
        message = msg;
        setTimeout(() => {
            isShow = false;
        }, 3000);
    }
</script>

{#if isShow === true}
    <div class={`growl ${growlType}`}>
        {#if growlType === 'info'}
            <Icon className="icon_growl" size={20} color='#FFF' src={FaSolidInfoCircle}/>
        {/if}
        {#if growlType === 'success'}
            <Icon className="icon_growl" size={20} color='#FFF' src={FaSolidCheckCircle}/>
        {/if}
        {#if growlType === 'error'}
            <Icon className="icon_growl" size={20} color='#FFF' src={FaSolidTimesCircle}/>
        {/if}
        {#if growlType === 'warning'}
            <Icon className="icon_growl" size={19} color='#FFF' src={FaSolidExclamationTriangle}/>
        {/if}
        <span>{message}</span>
    </div>
{/if}

<style>
    .growl {
        position       : fixed;
        display        : flex;
        flex-direction : row;
        align-items    : center;
        gap            : 8px;
        font-size      : 14px;
        bottom         : 20px;
        right          : 20px;
        padding        : 10px 20px;
        border-radius  : 3px;
        box-shadow     : 0px 4px 24px 0px rgba(0, 0, 0, 0.25);
        z-index        : 9999;
        min-width      : 120px;
        max-width      : 350px;
        height         : fit-content;
    }

    :global(.icon_growl) {
        flex-shrink : 0;
    }

    .growl span {
        flex-shrink : 1;
    }

    .growl.info {
        background-color : #1080E8;
        color            : #FFF;
    }

    .growl.success {
        background-color : #059669;
        color            : #FFF;
    }

    .growl.error {
        background-color : #EF4444;
        color            : #FFF;
    }

    .growl.warning {
        background-color : #D97706;
        color            : #FFF;
    }
</style>
