<script>
  import OptionButtons from "./OptionButtons.svelte";

  export let visible = false;
  export function showModal() {
    visible = true;
  }
  export function hideModal() {
    visible = false;
  }
  export let floor_type = '';
  let floor_types = [
    "basement", "floor"
  ]
</script>

{#if visible}
<div class='backdrop'>
  <div class='add_floor_dialog'>
    <div class='add_floor_content'>
      <h3>Add Floor</h3>
      <div class="floor_type">
        <label for="floor_type">House Type</label>
        <OptionButtons selected={floor_type} options={floor_types}/>
      </div>
    </div>
    <div class='buttons'>
      <button on:click|preventDefault={() => hideModal()} class='cancel_button'>Cancel</button>
      <slot></slot>
    </div>
  </div>
</div>
{/if}

<style>
  .backdrop {
    position: absolute;
    z-index: 40;
    background: rgba(41, 41, 41, 0.5);
    width: 100%;
    left: 0;
    top: 0;
    bottom: 0;
    height: auto;
    border-radius: 8px;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .add_floor_dialog {
    color: #334155;
    width: 400px;
    height: fit-content;
    gap: 40px;
    background-color: white;
    padding: 20px;
    border-radius: 8px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }
  .add_floor_dialog h3 {
    font-weight: 600;
    margin: 0;
    font-size: 18px;
    width: 100%;
    text-align: center;
  }
  .add_floor_dialog .buttons {
    display: flex;
    flex-direction: row;
    align-items: stretch;
    font-weight: 500;
    width: 100%;
  }
  .add_floor_dialog .buttons button {
    border-radius: 8px;
    border: none;
    padding: 10px;
    cursor: pointer;
    width: 100%;
    text-transform: capitalize;
  }
  .add_floor_dialog .buttons .cancel_button {
    margin-right: 10px;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    color: #f97316;
  }
  .add_floor_dialog .buttons .cancel_button:hover {
    border: 1px solid #f97316;
  }
  .add_floor_content .floor_type {
    position: relative;
    display: flex;
    flex-direction: column;
    width: 100%;
    height: fit-content;
  }
  .add_floor_content .floor_type label {
    font-size     : 13px;
    font-weight   : 700;
    margin-left   : 10px;
    margin-bottom : 8px;
  }
  
</style>