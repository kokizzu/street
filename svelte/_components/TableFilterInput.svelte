<script>
  import { run } from 'svelte/legacy';

  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { RiSystemFilterLine } from '../node_modules/svelte-icons-pack/dist/ri';

  /**
   * @typedef {Object} Props
   * @property {string} [label]
   * @property {string} [value]
   */

  /** @type {Props} */
  let { label = '', value = $bindable('') } = $props();

  let isShowInput = /** @type {boolean} */ ($state(false));
  let isFiltering = /** @type {boolean} */ ($state(false));

  run(() => {
    if (value !== '') {
      isFiltering = true;
    } else {
      isFiltering = false;
    }
  });
</script>

<div class="header" class:active={isFiltering}>
  <span>{label}</span>
  <button class="btn" onclick={() => (isShowInput = !isShowInput)}>
    <Icon
      src={RiSystemFilterLine}
      size="17"
      color="var(--gray-008)"
    />
  </button>
  <div class="input-box" class:active={isShowInput}>
    <input type="text" bind:value title="Separate with pipe for multiple values, for example:
    >=100|<50|61|72 will show values greater equal to 100, OR less than 50, OR exactly 61 OR 72
      filtering with greater or less than will only show correct result if data is saved as number
      currently price and size NOT stored as number
    <>abc* will show values NOT started with abc*
    abc*def|ghi will show values started with abc ends with def OR exactly ghi
    *jkl* will show values containing jkl substring
  multiple filter from other fields will do AND operation"
    placeholder="Input filter here"
  />
  </div>
</div>

<style>
  .header {
    position: relative;
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 5px;
    padding: 5px 10px;
    text-wrap: nowrap;
  }

  .header.active {
    background-color: var(--orange-transparent);
    color: var(--orange-005);
    border-radius: 5px;
  }

  :global(.header.active svg) {
    fill: var(--orange-005);
  }

  .header .btn {
    border: none;
    color: var(--gray-008);
    background-color: transparent;
    width: fit-content;
    padding: 8px;
    border-radius: 8px;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
  }

  .header .btn:hover {
    background-color: var(--orange-transparent);
    color: var(--orange-005);
  }

  .header .btn:disabled {
    cursor: not-allowed;
    background-color: var(--gray-002);
  }

  :global(.header .btn:hover svg) {
    fill: var(--orange-005);
  }

  :global(.header .btn:disabled svg) {
    fill: var(--gray-007);
  }

  .header .input-box {
    position: absolute;
    top: 100%;
    left: -5px;
    width: 200px;
    z-index: 100;
    background-color: #fff;
    border: 1px solid var(--gray-002);
    border-radius: 8px;
    padding: 5px;
    display: none;
    height: fit-content;
  }

  .header .input-box.active {
    display: flex;
  }

  .header .input-box input {
    width: 100%;
    border: 1px solid var(--gray-003);
    outline: none;
    border-radius: 8px;
    padding: 10px 12px;
    font-size: 13px;
    background-color: transparent;
  }

  .header .input-box input:focus {
    border-color: var(--orange-005);
    outline: 1px solid var(--orange-005);
  }
</style>