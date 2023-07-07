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
      // TODO: change confirm below to show diff before save
      if( confirm( 'are you sure you want to delete this row?' ) ) {
         loading = true;
         onConfirm( 'delete', row );
      }
   }
  
   function restorePressed() {
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
   <div class='modal_container'>
      <header>
         {#if row.id}
            <h2>Edit {rowType} ID: {row.id}</h2>
         {:else}
            <h2>New {rowType}</h2>
         {/if}
         <button aria-label='close' title='Close' type='button' on:click={cancelPressed}>
            <i class='gg-close' />
         </button>
      </header>
      <div class="input_container">
         {#each fields as field}
            {#if field.inputType==='hidden'}
               <input type='hidden' bind:value={row[field.name]} />
            {:else} 
               {#if field.readOnly}
                  <div class="form_info">
                     <label for='modalForm__{field.name}'>{field.label}</label>
                     {#if field.inputType==='datetime'}
                        <span>{datetime( row[ field.name ] )}</span>
                     {:else}
                        <span>{row[ field.name ]}</span>
                     {/if}
                  </div>
               {:else}
                  <div class="input_box">
                     <label for='modalForm__{field.name}'>{field.label}</label>
                     {#if field.inputType==='textarea'}
                        <textarea id='modalForm__{field.name}' bind:value={row[field.name]}></textarea>
                     {:else if field.inputType==='number'}
                        <input id='modalForm__{field.name}' type='number' bind:value={row[field.name]} />
                     {:else}
                        <input id='modalForm__{field.name}' type='text' bind:value={row[field.name]} />
                     {/if}
                  </div>
               {/if}
            {/if}
         {/each}
         <slot />
      </div>
      <div class='button_container'>
         <button tabindex='0' class='cancel' on:click={cancelPressed}>
            Cancel
         </button>
         {#if loading}
            <span class='right'>Saving..</span>
         {:else}
            <button tabindex='0' class='save' on:click={savePressed}>
               Save
            </button>
            {#if row.id}
               {#if row.deletedAt>0}
                  <button tabindex='0' class='restore' on:click={restorePressed}>
                     Restore
                  </button>
               {:else}
                  <button tabindex='0' class='delete' on:click={deletePressed}>
                     Delete
                  </button>
               {/if}
            {/if}
         {/if}
      </div>
   </div>
{/if}

<style>
   .backdrop {
      position: fixed;
      z-index: 40;
      top: 0;
      left: 0;
      background: grey;
      width: 100%;
      height: 100%;
      opacity: .8;
   }

   .modal_container {
      background-color: white;
      position: fixed;
      z-index: 50;
      top: 15%;
      left: 50%;
      width: 500px;
      height: fit-content;
      margin-left: -300px;
      margin-top: 0;
      padding: 20px;
      filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
      border-radius: 15px;
      display: flex;
      flex-direction: column;
      color: #334155;
   }
   .modal_container header {
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      align-items: center;
   }
   .modal_container header h2 {
      font-size: 20px;
      line-height: 1.5rem;
      padding: 0;
      margin: 0;
   }
   .modal_container header button {
      padding: 5px;
		border: none;
		background: none;
		border-radius: 5px;
		font-size: 14px;
		cursor: pointer;
   }
   .modal_container header button:hover {
      background-color: rgb(0 0 0 / 0.07);
      color: #EF4444;
   }

   .input_container {
      margin: 0 0 20px 0;
      width: 100%;
      display: flex;
      flex-direction: column;
   }
   .input_container .input_box {
      display: flex;
      flex-direction: column;
      width: 100%;
      margin-top: 10px;
   }
   .input_container .input_box label {
      font-size: 13px;
      font-weight: 700;
      margin-left: 10px;
   }
   .input_container .input_box input {
      width: 100%;
      border: 1px solid #CBD5E1;
      background-color: #F1F5F9;
   	border-radius: 8px;
   	padding: 8px;
      margin-top: 5px;
   }
   .input_container .input_box input:focus {
      border-color: #3b82f6;
      outline: 1px solid #3b82f6;
   }
   .input_container .form_info {
      display: inline-flex;
      align-items: center;
      font-size: 13px;
      margin: 5px 0 0 10px;
   }

   .input_container .form_info label {
      font-weight: 700;
      margin-right: 10px;
   }

   .button_container {
      display: flex;
      flex-direction: row;
      justify-content: flex-end;
      align-items: center;
      align-content: stretch;
   }
   .button_container button {
      margin-left: 10px;
      padding: 8px 10px;
      font-size: 13px;
      font-weight: 600;
      border-radius: 5px;
      color: white;
      border: none;
      cursor: pointer;
      filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
   }
   .button_container .cancel {
      background-color: #EF4444;
   }
   .button_container .save {
      background-color: #3b82f6;
   }
   .button_container .delete {
      background-color: #ec4899;
   }
   .button_container .restore {
      background-color: #f97316;
   }
</style>