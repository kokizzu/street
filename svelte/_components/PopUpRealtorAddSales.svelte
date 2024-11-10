<script>
  /** @typedef {import('../_types/business').Sales} Sales */

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { FiLoader } from '../node_modules/svelte-icons-pack/dist/fi';
  import { IoClose } from '../node_modules/svelte-icons-pack/dist/io';
  import InputBox from './InputBox.svelte';
  import { dateISOFormat } from './formatter';

  let isShow = /** @type {boolean} */ (false);

  export let heading      = /** @type {string} */ ('Add Sales');
  export let isSubmitted  = /** @type {boolean} */ (false);

  /**
   * @description Submit add sales
   * @type {Function}
   * @param salesObj {Sales}
   * @param propKey {string}
   * @returns {Promise<void>}
   */
  export let OnSubmit = async function(salesObj, propKey) {
    console.log('Sales =', salesObj);
    console.log('Property Key =', propKey);
  };

  let propertyCountry = /** @type {string} */ ('Default');
  let propertyKey = /** @type {string} */ ('');
  let buyerEmail = /** @type {string} */ ('');
  let price = /** @type {number|string|any} */ ('0');
  let salesDate = /** @type {string|Date|any} */ (dateISOFormat(0));

  export const Reset = () => {
    propertyCountry = /** @type {string} */ ('Default');
    propertyKey = /** @type {string} */ ('');
    buyerEmail = /** @type {string} */ ('');
    price = /** @type {number|string|any} */ ('0');
    salesDate = /** @type {string|Date|any} */ (dateISOFormat(0));
  }
  export const Show = () => isShow = true;
  export const Hide = () => {
    isShow = false;
    isSubmitted = false;
  }
  const cancel = () => {
    isShow = false;
    isSubmitted = false;
    Reset();
  }

  function submitAddSales() {
    isSubmitted = true;
    const salesObj = /** @type {Sales|any} */ ({
      propertyCountry: propertyCountry,
      buyerEmail: buyerEmail,
      price: String(price),
      salesDate: String(salesDate)
    });
    OnSubmit(salesObj, propertyKey);
  }
</script>

<div class={`popup_container ${isShow ? 'show' : ''}`}>
  <div class="popup">
    <header class="header">
      <h2>{heading}</h2>
      <button on:click={Hide}>
        <Icon size="22" color="var(--red-005)" src={IoClose}/>
      </button>
    </header>
    <div class="forms">
      <InputBox
        label="Buyer Email"
        type="email"
        id="buyerEmail"
        bind:value={buyerEmail}
        autocomplete="on"
        placeholder="johndoe@example.com"
      />
      <InputBox
        label="Property Key"
        type="text"
        id="propertyKey"
        bind:value={propertyKey}
        autocomplete="on"
        placeholder="#US78193"
      />
      <InputBox
        label="Property Country"
        type="combobox-arr"
        id="propertyCountry"
        bind:value={propertyCountry}
        valuesArr={[
          'Default', 'US', 'TW'
        ]}
      />
      <InputBox
        label="Price"
        type="number"
        id="price"
        bind:value={price}
        autocomplete="on"
        placeholder="Price"
      />
      <InputBox
        label="Sales Date"
        type="date"
        id="salesDate"
        bind:value={salesDate}
        autocomplete="on"
        placeholder="Sales Date"
      />
    </div>
    <div class="foot">
      <div class="left">
      </div>
      <div class="right">
        <button class="cancel" on:click|preventDefault={cancel}>Cancel</button>
        <button class="ok" on:click|preventDefault={submitAddSales} disabled={isSubmitted}>
          {#if !isSubmitted}
            <span>Submit</span>
          {/if}
          {#if isSubmitted}
            <span>Submitting...</span>
            <Icon className="spin" color="#FFF" size="14" src={FiLoader} />
          {/if}
        </button>
      </div>
    </div>
  </div>
</div>

<style>
  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }

  :global(.spin) {
    animation: spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
  }

  .popup_container {
    display: none;
		position: fixed;
		width: 100%;
		height: 100%;
		top: 0;
		left: 0;
		bottom: 0;
		right: 0;
		z-index: 2000;
		background-color: rgba(0 0 0 / 40%);
		backdrop-filter: blur(1px);
		justify-content: center;
		padding: 50px;
    overflow: auto;
	}

  .popup_container.show {
    display: flex;
  }

	.popup_container .popup {
		border-radius: 8px;
		background-color: #FFF;
		height: fit-content;
		width: 500px;
		display: flex;
		flex-direction: column;
	}

  .popup_container .popup header {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		padding: 10px 20px;
		border-bottom: 1px solid var(--gray-004);
	}

	.popup_container .popup header h2 {
		margin: 0;
	}

	.popup_container .popup header button {
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 5px;
		border-radius: 50%;
		border: none;
		background-color: transparent;
		cursor: pointer;
	}

	.popup_container .popup header button:hover {
		background-color: #ef444420;
	}

	.popup_container .popup header button:active {
		background-color: #ef444430;
	}

	.popup_container .popup .forms {
		padding: 20px;
		display: flex;
		flex-direction: column;
		gap: 10px;
	}

	.popup_container .popup .foot {
		display: flex;
		flex-direction: row;
    justify-content: space-between;
		gap: 10px;
		align-items: center;
		padding: 10px 20px;
		border-top: 1px solid var(--gray-004);
	}

  .popup_container .popup .foot .right {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
  }

	.popup_container .popup .foot button {
		padding: 12px 15px;
		border-radius: 9999px;
		border: none;
		color: #FFF;
		cursor: pointer;
		font-weight: 600;
	}

	.popup_container .popup .foot button.ok {
		background-color: var(--green-006);
    display: flex;
    flex-direction: row;
    gap: 10px;
    justify-content: center;
    align-items: center;
	}

	.popup_container .popup .foot button.ok:hover {
		background-color: var(--green-005);
	}

	.popup_container .popup .foot button.ok:disabled {
		cursor: not-allowed;
		background-color: var(--green-005);
	}

	.popup_container .popup .foot button.cancel {
		background-color: #fbbf2430;
		color: var(--amber-005);
	}

	.popup_container .popup .foot button.cancel:hover {
		background-color: #fbbf2450;
	}

  @media only screen and (max-width : 768px) {
    .popup_container {
      padding: 10px;
    }
  }
</style>