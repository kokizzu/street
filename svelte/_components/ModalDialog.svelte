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
  <div class="backdrop">
    <div class="modal_dialog">
      <div class="modal_content">
        {#if title}
          <h3>{title}</h3>
        {/if}
        <slot name="content" />
      </div>
      <div class="buttons">
        <button on:click|preventDefault={() => hideModal()} class="cancel_button">{close_label} </button>
        <slot name="buttons" />
      </div>
    </div>
  </div>
{/if}

<style>
  .backdrop {
    position: fixed;
    z-index: 101;
    top: 0;
    left: 0;
    bottom: 0;
    background: rgba(0 0 0 / 40%);
    overflow-y: auto;
    width: 100%;
    display: flex;
    justify-content: center;
  }
  /* Hide scrollbar to not make it 2 in right side */
  .backdrop::-webkit-scrollbar-thumb {
    background: transparent;
  }
  .backdrop::-webkit-scrollbar {
    width: 0;
  }
  .backdrop::-webkit-scrollbar-track {
    background-color: transparent;
  }

  .modal_dialog {
    background-color: white;
    width: 55%;
    height: fit-content;
    margin: 50px 0 0 0;
    padding: 20px;
    filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    border-radius: 15px;
    display: flex;
    flex-direction: column;
    gap: 20px;
    color: #334155;
  }

  .modal_dialog .modal_content {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }
  .modal_dialog .modal_content h3 {
    font-weight: 600;
    margin: 0;
    font-size: 18px;
    width: 100%;
    text-align: center;
  }

  .modal_dialog .buttons {
    display: flex;
    flex-direction: row;
    align-items: stretch;
    font-weight: 500;
    width: 100%;
  }

  .modal_dialog .buttons button {
    border-radius: 8px;
    border: none;
    padding: 10px;
    cursor: pointer;
    width: 100%;
    text-transform: capitalize;
  }

  .modal_dialog .buttons .cancel_button {
    margin-right: 10px;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    color: #f97316;
  }

  .modal_dialog .buttons .cancel_button:hover {
    border: 1px solid #f97316;
  }
</style>