<script>
  export let visible = false;
  export function showModal() {
    visible = true;
  }
  export function hideModal() {
    visible = false;
  }
  export let floor_type = '';
</script>

{#if visible}
<div class='backdrop'>
  <div class='add_floor_dialog'>
    <div class='add_floor_content'>
      <h3>Add Floor</h3>
      <div class="floor_type">
        <label for="floor_type">House Type</label>
        <div class="option_container">
          <label class={floor_type === 'floor' ? 'option clicked': 'option'} for="floor">
            <input type="radio" on:click={() => (floor_type = 'floor')} id="floor" value="floor" />
            floor
          </label>
          <label class={floor_type === 'basement' ? 'option clicked': 'option'} for="basement">
            <input
              type="radio"
              on:click={() => (floor_type = 'basement')}
              id="basement"
              value="basement"
            />
            Basement
          </label>
        </div>
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
  .add_floor_content .floor_type .option_container {
    width: 100%;
    display: grid;
    grid-template-rows: 1fr 1fr;
    gap: 10px;
  }
  .add_floor_content .floor_type .option_container .option {
    margin: 0;
    padding: 10px 12px;
    border-radius: 8px;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    font-weight: 500;
    text-align: center;
    cursor: pointer;
  }
  .add_floor_content .floor_type .option_container .option:hover {
    border: 1px solid #f97316;
    color: #f97316;
  }
  .add_floor_content .floor_type .option_container .option.clicked {
    background-color: #f97316;
    color: white;
    border: none;
  }
  .option input[type='radio'] {
    position       : absolute;
    opacity        : 0;
    pointer-events : none;
  }
</style>