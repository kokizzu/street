<script>
  export let visible = false;
  export function showModal() {
    visible = true;
  }
  export function hideModal() {
    visible = false;
  }
  export let room_type = '';
  const room_type_lists = [
    'bedroom',
    'bathroom',
  ]
  let showRoomTypeOption = false;
  function handleRoomTypeOption() {
    showRoomTypeOption = !showRoomTypeOption;
  }
  export let room_total = 0;
</script>

{#if visible}
<div class='backdrop'>
  <div class='add_room_dialog'>
    <div class='add_room_content'>
      <h3>Add Room</h3>
      <div class='room_type'>
        <label for='room_type'>Type</label>
        <button on:click={handleRoomTypeOption}>{room_type ? room_type : 'select'}</button>
        {#if showRoomTypeOption}
          <div class='option_container'>
            {#each room_type_lists as room}
              <label class='option' for={room}>
                <input
                  type='radio'
                  bind:group={room_type}
                  on:change={() => handleRoomTypeOption()}
                  id={room}
                  value={room}
                />
                {room}
              </label>
            {/each}
          </div>
        {/if}
      </div>
      <div class='room_total'>
        <label for='room_total'>Total</label>
        <input bind:value={room_total} type='number' name='room_total' id='room_total'>
      </div>
    </div>
    <div class='buttons'>
      <button
        on:click|preventDefault={() => {
          showRoomTypeOption = false;
          hideModal();
        }}
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
  .add_room_content .room_type {
    display: flex;
    flex-direction: column;
    width: 100%;
    position: relative;
    text-transform: capitalize;
    margin-top: 15px;
  }
  .add_room_content .room_type label {
    font-size: 13px;
    font-weight: 700;
    margin-left: 10px;
    margin-bottom: 8px;
  }
  .add_room_content .room_type button {
    width: auto;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    border-radius: 8px;
    padding: 10px 12px;
    text-align: left;
    cursor: pointer;
    text-transform: capitalize;
  }
  .add_room_content .room_type button:hover {
    border: 1px solid #f97316;
    outline: 1px solid #f97316;
  }
  .add_room_content .option_container {
    width: 100%;
    position: absolute;
    top: 70px;
    z-index: 3;
    display: flex;
    flex-direction: column;
    border-radius: 8px;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
  }
  .add_room_content .option_container .option {
    margin: 0;
    padding: 10px 12px;
    font-weight: 500;
    cursor: pointer;
  }
  .add_room_content .option_container .option:hover {
    color: #f97316;
  }
  .option input[type='radio'] {
    position: absolute;
    opacity: 0;
    pointer-events: none;
  }
  .add_room_content .room_total {
    display: flex;
    flex-direction: column;
    width: 100%;
    text-transform: capitalize;
    margin-top: 15px;
  }
  .add_room_content .room_total label {
    font-size: 13px;
    font-weight: 700;
    margin-left: 10px;
    margin-bottom: 8px;
  }
  .add_room_content .room_total input {
    width: auto;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    border-radius: 8px;
    padding: 10px 12px;
    text-align: left;
    text-transform: capitalize;
  }
</style>