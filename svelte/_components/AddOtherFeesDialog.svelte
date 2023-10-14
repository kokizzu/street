<script>
  export let visible = false;
  
  export function showModal() {
    visible = true;
  }
  
  export function hideModal() {
    visible = false;
  }
  
  export let name = '';
  export let fee = 0;
</script>

{#if visible}
	<div class='backdrop'>
		<div class='add_fee_dialog'>
			<div class='add_fee_content'>
				<h3>Add Floor</h3>
				<div class='input_box'>
					<label for='fee_name'>Title</label>
					<input id='fee_name' type='text' placeholder='Management Fee' bind:value={name}/>
				</div>
				<div class='input_box'>
					<label for='total_fee'>Fee</label>
					<input id='total_fee' type='number' min='0' bind:value={fee}/>
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

    .add_fee_dialog {
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

    .add_fee_dialog h3 {
        font-weight : 600;
        margin      : 0;
        font-size   : 18px;
        width       : 100%;
        text-align  : center;
    }

    .add_fee_dialog .add_fee_content {
        display        : flex;
        flex-direction : column;
        gap            : 15px;
    }

    .add_fee_dialog .add_fee_content .input_box {
        display        : flex;
        flex-direction : column;
        gap            : 8px;
        width          : 100%;
    }

    .add_fee_dialog .add_fee_content .input_box label {
        font-size   : 13px;
        font-weight : 700;
        margin-left : 10px;
    }

    .add_fee_dialog .add_fee_content .input_box input {
        width            : 100%;
        border           : 1px solid #CBD5E1;
        background-color : #F1F5F9;
        border-radius    : 8px;
        padding          : 12px;
    }

    .add_fee_dialog .add_fee_content .input_box input:focus {
        border-color : #3B82F6;
        outline      : 1px solid #3B82F6;
    }

    .add_fee_dialog .buttons {
        display        : flex;
        flex-direction : row;
        align-items    : stretch;
        font-weight    : 500;
        width          : 100%;
    }

    .add_fee_dialog .buttons button {
        border-radius  : 8px;
        border         : none;
        padding        : 10px;
        cursor         : pointer;
        width          : 100%;
        text-transform : capitalize;
    }

    .add_fee_dialog .buttons .cancel_button {
        margin-right     : 10px;
        border           : 1px solid #CBD5E1;
        background-color : #F1F5F9;
        color            : #F97316;
    }

    .add_fee_dialog .buttons .cancel_button:hover {
        border : 1px solid #F97316;
    }
</style>