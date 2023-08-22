<script>
  export let visible = false;
  export function showModal() {
    visible = true;
  }
  export function hideModal() {
    visible = false;
  }
  export let room_type = '';
  // unit mode
  const sqft = 'SqFt';
  const m2 = 'M2';
  export let unit_mode = m2;
  export let room_size = 0;
  let sqft_size = 0;
  export let m2_size = 0;
  function toggleUnitMode() {
    if (unit_mode == sqft) {
      if (room_size !== 0) {
        sqft_size = room_size;
        m2_size = sqftToM2(room_size);
        unit_mode = m2;
        room_size = m2_size;
      } else {
        unit_mode = m2;
      }
    } else if (unit_mode == m2) {
      if (room_size !== 0) {
        m2_size = room_size;
        sqft_size = m2ToSqft(room_size);
        unit_mode = sqft;
        room_size = sqft_size;
      } else {
        unit_mode = sqft;
      }
    }
  }
  export function sqftToM2(sqft) {
    var value = sqft * 0.09290304;
    var minifiedValue = value.toFixed(2);
    return parseFloat(minifiedValue);
  }
  function m2ToSqft(m2) {
    var value = m2 / 0.09290304;
    var minifiedValue = value.toFixed(2);
    return parseFloat(minifiedValue);
  }
</script>

{#if visible}
<div class='backdrop'>
  <div class='add_room_dialog'>
    <div class='add_room_content'>
      <h3>Add Room</h3>
      <div class="room_type">
        <label for="room_type">House Type</label>
        <div class="option_container">
          <label class={room_type === 'bedroom' ? 'option clicked': 'option'} for="bedroom">
            <input type="radio" on:click={() => (room_type = 'bedroom')} id="bedroom" value="bedroom" />
            bedroom
          </label>
          <label class={room_type === 'bathroom' ? 'option clicked': 'option'} for="bathroom">
            <input type="radio" on:click={() => (room_type = 'bathroom')} id="bathroom" value="bathroom" />
            bathroom
          </label>
          <label class={room_type === 'living room' ? 'option clicked': 'option'} for="living_room">
            <input type="radio" on:click={() => (room_type = 'living room')} id="living_room" value="living room" />
            living room
          </label>
        </div>
      </div>
      <div class='room_size'>
        <label for='room_size'>Size</label>
        <div class='room_input_size_box'>
          <input bind:value={room_size} min='0' step='0.01' type='number' name='room_size' id='room_size'>
          <div class='unit_toggle'>
            <p>{unit_mode}</p>
            <button on:click={toggleUnitMode}>
              <i class='gg-arrows-exchange'></i>
            </button>
          </div>
        </div>
      </div>
    </div>
    <div class='buttons'>
      <button
        on:click|preventDefault={() => hideModal()}
        class='cancel_button'>Cancel</button>
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
  .add_room_dialog {
    color: #334155;
    width: 400px;
    height: 400px;
    background-color: white;
    padding: 20px;
    border-radius: 8px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }
  .add_room_dialog h3 {
    font-weight: 600;
    margin: 0;
    font-size: 18px;
    width: 100%;
    text-align: center;
  }
  .add_room_dialog .buttons {
    display: flex;
    flex-direction: row;
    align-items: stretch;
    font-weight: 500;
    margin-top: 20px;
    width: 100%;
  }
  .add_room_dialog .buttons button {
    border-radius: 8px;
    border: none;
    padding: 10px;
    cursor: pointer;
    width: 100%;
    text-transform: capitalize;
  }
  .add_room_dialog .buttons .cancel_button {
    margin-right: 10px;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    color: #f97316;
  }
  .add_room_dialog .buttons .cancel_button:hover {
    border: 1px solid #f97316;
  }
  .add_room_content {
    display: flex;
    flex-direction: column;
    gap: 15px;
  }
  .add_room_content .room_type {
    position: relative;
    display: flex;
    flex-direction: column;
    width: 100%;
    height: fit-content;
  }
  .add_room_content .room_type label {
    font-size     : 13px;
    font-weight   : 700;
    margin-left   : 10px;
    margin-bottom : 8px;
  }
  .add_room_content .room_type .option_container {
    width: 100%;
    display: grid;
    grid-template-rows: 1fr 1fr;
    gap: 10px;
  }
  .add_room_content .room_type .option_container .option {
    margin: 0;
    padding: 10px 12px;
    border-radius: 8px;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    font-weight: 500;
    text-align: center;
    cursor: pointer;
  }
  .add_room_content .room_type .option_container .option:hover {
    border: 1px solid #f97316;
    color: #f97316;
  }
  .add_room_content .room_type .option_container .option.clicked {
    background-color: #f97316;
    color: white;
    border: none;
  }
  .option input[type='radio'] {
    position       : absolute;
    opacity        : 0;
    pointer-events : none;
  }
  .add_room_content .room_size {
    display: flex;
    flex-direction: column;
    width: 100%;
    text-transform: capitalize;
  }
  .add_room_content .room_size label {
    font-size: 13px;
    font-weight: 700;
    margin-left: 10px;
    margin-bottom: 8px;
  }
  .add_room_content .room_size .room_input_size_box {
    display: flex;
    flex-direction: row;
    align-items: center;
  }
  .add_room_content .room_size .room_input_size_box input {
    flex-grow: 1;
    width: auto;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    border-radius: 8px;
    padding: 10px 12px;
    text-align: left;
    text-transform: capitalize;
  }
  .add_room_content .room_size .room_input_size_box .unit_toggle {
    display: flex;
    flex-direction: row;
    align-items: center;
    margin-left: 8px;
    font-weight: 500;
  }
  .add_room_content .room_size .room_input_size_box .unit_toggle button {
    border: none;
    background: none;
    padding: 9px 8px;
    border-radius: 50%;
    margin-left: 8px;
    cursor: pointer;
    color: #f97316;
  }
  .add_room_content .room_size .room_input_size_box .unit_toggle button:hover {
    background-color: rgb( 0 0 0 / 0.06);
  }
</style>