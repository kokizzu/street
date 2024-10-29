<script>
  /** @typedef {import('../_types/master').InputType} InputType */

  import { onMount } from 'svelte';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { AiOutlineEye, AiOutlineEyeInvisible } from '../node_modules/svelte-icons-pack/dist/ai';

  export let className    = /** @type {string} */ ('');
  export let id           = /** @type {string} */ ('');
  export let label        = /** @type {string} */ ('');
  export let value        = /** @type {string} */ ('');
  export let placeholder  = /** @type {string} */ ('');
  export let type         = /** @type {InputType | string} */ ('text');

  let isShowPassword  = /** @type {boolean} */ (false);
  let inputElm        = /** @type {HTMLInputElement} */ (null);
  
  onMount(() => {
    if (type === 'password') inputElm.type = type;
  });

  function toggleShowPassword() {
    isShowPassword = !isShowPassword;
    if (isShowPassword) inputElm.type = 'text';
    else inputElm.type = 'password';
  }
</script>

<div class={className}>
  <div class="input_box {type == 'password' ? 'with_password' : ''}">
    {#if type === 'textarea'}
      <label class="label" for={id}>{label}</label>
      <textarea bind:value={value} {id} {placeholder}></textarea>
    {:else if type === 'text'}
      <label class="label" for={id}>{label}</label>
      <input type="text" bind:value={value} {id} {placeholder} autocomplete="off" />
    {:else if type === 'email'}
      <label class="label" for={id}>{label}</label>
      <input type="email" bind:value={value} {id} {placeholder}/>
    {:else if type === 'date'}
      <label class="label" for={id}>{label}</label>
      <input type="date" bind:value={value} {id} {placeholder}/>
    {:else if type === 'password'}
      <label class="label" for={id}>{label}</label>
      <input bind:value={value} {id} bind:this={inputElm} {placeholder}/>
      {#if type === 'password'}
        <button class="eye" on:click={toggleShowPassword}>
          {#if !isShowPassword}
            <Icon color="#495057" size="20" src={AiOutlineEye}/>
          {/if}
          {#if isShowPassword}
            <Icon color="#495057" size="20" src={AiOutlineEyeInvisible}/>
          {/if}
        </button>
      {/if}
    {:else if type === 'color'}
      <label class="label" for={id}>{label}</label>
      <div class="color_box">
        <input type="color" bind:value={value} {id} class="color-input"/>
      </div>
    {:else}
      <label class="label" for={id}>{label}</label>
      <input type="text" bind:value={value} {id} {placeholder}/>
    {/if}
  </div>
</div>

<style>
  .show {
    display: block;
  }

  .hidden {
    display: none;
  }

  .input_box {
    position: relative;
    display: flex;
    flex-direction: column;
    width: 100%;
    color: var(--gray-007);
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .input_box.with_password input{
    padding-right: 40px !important;
  }

  .input_box.bool {
    width: fit-content;
  }

  .input_box .label {
    font-size: var(--font-base);
    margin-left: 10px;
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    line-clamp: 1;
  }

  .input_box input,
  .input_box textarea {
    width: 100%;
    border: 1px solid var(--gray-003);
    border-radius: 5px;
    background-color: transparent;
    padding: 12px 12px;
  }

  .input_box input:focus,
  .input_box textarea:focus {
    border-color: var(--orange-005);
    outline: 1px solid var(--orange-005);
  }

  .color_box {
    display: flex;
    flex-direction: column;
    justify-content: start;
  }

  .color-input {
    -webkit-appearance: none;
    -moz-appearance: none;
    appearance: none;
    width: 100px !important;
    height: 30px;
    background-color: transparent;
    border: none;
    cursor: pointer;
    padding: 0 !important;
  }
  .color-input::-webkit-color-swatch {
    border-radius: 5px;
    border: none;
  }
  .color-input::-moz-color-swatch {
    border-radius: 5px;
    border: none;
  }

  .input_box textarea {
    resize: vertical;
    height: 90px;
    min-height: 50px;
    max-height: 300px;
  }

  .input_box .eye {
    position: absolute;
    height: fit-content;
    width: fit-content;
    background-color: transparent;
    padding: 0;
    top: 30px;
    bottom: auto;
    right: 10px;
    border: none;
    cursor: pointer;
  }

  .input_box .eye:focus {
    border: none;
    outline: none;
  }

  :global(.input_box .eye:hover svg) {
    fill: var(--orange-005);
  }

  .input_percentage {
    display: flex;
    position: relative;
  }

  .input_percentage input {
    padding-right: 30px !important;
  }

  .input_box .options_container .input_container input {
    padding-right: 30px !important;
  }
</style>