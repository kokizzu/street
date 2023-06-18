<script>
  import { datetime } from './formatter.js';
  
  export let fields = []; // list of editable properties
  export let visible = false;
  export let loading = false;
  export let row = {};
  export let rowType = 'Row';
  export let onConfirm = function( action, row ) {
    // action = upsert, delete, restore
    console.log( 'ModalForm.onConfirm', action, row );
  };
  
  let originalRowJson = '';
  
  export function showModal( fetchedRow ) {
    visible = true;
    originalRowJson = JSON.stringify( fetchedRow );
    row = fetchedRow;
    loading = false;
    console.log( 'ModalForm.showModal', fetchedRow );
  }
  
  export function hideModal() {
    visible = false;
  }
  
  export function setLoading( to ) {
    loading = !!to;
  }
  
  function haveChanges() {
    return originalRowJson!==JSON.stringify( row );
  }
  
  function deletePressed() {
    if( !haveChanges() ) return cancelPressed();
    // TODO: change confirm below to show diff before save
    if( confirm( 'are you sure you want to delete this row?' ) ) {
      loading = true;
      onConfirm( 'delete', row );
    }
  }
  
  function restorePressed() {
    if( !haveChanges() ) return cancelPressed();
    // TODO: change confirm below to show diff before save
    if( confirm( 'are you sure you want to restore this row?' ) ) {
      loading = true;
      onConfirm( 'restore', row );
    }
  }
  
  function savePressed() {
    if( !haveChanges() ) return cancelPressed();
    // TODO: change confirm below to show diff before save
    if( confirm( 'are you sure you want to save this row?' ) ) {
      loading = true;
      onConfirm( 'upsert', row );
    }
  }
  
  function cancelPressed() {
    if( haveChanges() ) {
      if( confirm( 'are you sure you want to cancel? changes will be lost.' ) )
        visible = false;
    } else {
      visible = false;
    }
    
    // brb eating
  }

</script>
{#if visible}
  
  <div class='backdrop'></div>
  <div class='modal'>
    <div class='modal_background'></div>
    <div class='modal_content box -small' style='outline: none;'>
      <button tabindex='0' aria-label='close' title='Close' type='button' class='close' on:click={cancelPressed}>
        <i class='gg-close' />
      </button>
      <div class='title'>Edit {rowType} ID: {row.id}</div>
      <p>
        {#each fields as field}
          {#if field.inputType==='hidden'}
            <input type='hidden' bind:value={row[field.name]} />
          {:else}
            <label for='modalForm__{field.name}'>{field.label}</label>
            {#if field.readOnly}
              {#if field.inputType==='datetime'}
                <span>{datetime( row[ field.name ] )}</span>
              {:else}
                <span>{row[ field.name ]}</span>
              {/if}
            {:else}
              {#if field.inputType==='textarea'}
                <textarea id='modalForm__{field.name}' bind:value={row[field.name]}></textarea>
              {:else if field.inputType==='number'}
                <input id='modalForm__{field.name}' type='number' bind:value={row[field.name]} />
              {:else}
                <input id='modalForm__{field.name}' type='text' bind:value={row[field.name]} />
              {/if}
            {/if}
            <br />
          {/if}
        {/each}
        <slot />
      </p>
      <div class='layout-h'>
        <button tabindex='0' class='left button secondary' on:click={cancelPressed}>Cancel</button>
        {#if loading}
          <span class='right'>Saving..</span>
        {:else}
          <button tabindex='0' class='right button primary' on:click={savePressed}>Save</button>
          {#if row.deletedAt>0}
            <button tabindex='0' class='right button orange' on:click={restorePressed}>Restore</button>
          {:else}
            <button tabindex='0' class='right button danger' on:click={deletePressed}>Delete</button>
          {/if}
        {/if}
      </div>
    </div>
  </div>
{/if}
<style>
    .left {
        float : left
    }

    .right {
        float : right
    }

    .backdrop {
        position   : fixed;
        top        : 0;
        left       : 0;
        background : grey;
        width      : 100%;
        height     : 100%;
        opacity    : .8;
    }

    .modal {
        background-color : white;
        position         : fixed;
        top              : 20%;
        left             : 50%;
        width            : 600px;
        margin-left      : -300px;
        margin-top       : 0;
        overflow-y       : auto;
        box-shadow       : .25rem .25rem .5rem #DADADA, -.25rem -.25rem .5rem #FFF;
        border-radius    : 1em;
        padding          : 1em;
        opacity          : 1;
    }

    .title {
        font-weight : bold;
        font-size   : 1.2em;
    }

    button.close {
        float  : right;
    }

    label {
        width        : 12em;
        display      : inline-block;
        text-align   : right;
        margin-right : 1em;
    }

</style>