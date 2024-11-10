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
  export let autocomplete = /** @type {('on' | 'off')} */ ('on');
  export let valuesObj    = /** @type {Record<string, string>} */ ({});
  export let valuesArr    = /** @type {any[]} */ ([]);

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
  <div class="input-box {type == 'password' ? 'with-password' : ''}">
    {#if type === 'textarea'}
      <label class="label" for={id}>{label}</label>
      <textarea bind:value={value} {id} {placeholder}></textarea>
    {:else if type === 'text'}
      <label class="label" for={id}>{label}</label>
      <input type="text" bind:value={value} {id} {placeholder} {autocomplete} />
    {:else if type === 'email'}
      <label class="label" for={id}>{label}</label>
      <input type="email" bind:value={value} {id} {placeholder} {autocomplete} />
    {:else if type === 'date'}
      <label class="label" for={id}>{label}</label>
      <input type="date" bind:value={value} {id} {placeholder}/>
    {:else if type === 'password'}
      <label class="label" for={id}>{label}</label>
      <input bind:value={value} {id} bind:this={inputElm} {placeholder} />
      {#if type === 'password'}
        <button class="eye" on:click={toggleShowPassword} title={isShowPassword ? 'Hide Password' : 'Show Password'}>
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
    {:else if type === 'combobox-obj'}
      <label class="label" for={id}>{label}</label>
      <select bind:value={value} {id} {placeholder}>
        {#each Object.entries(valuesObj) as [key, value]}
          <option value={key} selected={value === value}>{value}</option>
        {/each}
      </select>
    {:else if type === 'combobox-arr'}
      <label class="label" for={id}>{label}</label>
      <select bind:value={value} {id} {placeholder}>
        {#each valuesArr as value}
          <option value={value} selected={value === value}>{value}</option>
        {/each}
      </select>
    {:else if type === 'date'}
      <label class="label" for={id}>{label}</label>
      <input type="date" bind:value={value} {id} {placeholder}/>
    {:else if type === 'month'}
      {#if label || label !== ''}
        <label class="label" for={id}>{label}</label>
      {/if}
      <input type="month" bind:value={value} {id} {placeholder}/>
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

  .input-box {
    position: relative;
    display: flex;
    flex-direction: column;
    width: 100%;
    color: var(--gray-007);
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .input-box.with-password input{
    padding-right: 40px !important;
  }

  .input-box.bool {
    width: fit-content;
  }

  .input-box .label {
    font-size: var(--font-base);
    margin-left: 10px;
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    line-clamp: 1;
  }

  .input-box input,
  .input-box textarea,
  .input-box select {
    width: 100%;
    border: 1px solid var(--gray-003);
    border-radius: 8px;
    background-color: transparent;
    padding: 12px 12px;
  }

  .input-box input:focus,
  .input-box textarea:focus,
  .input-box select:focus {
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

  .input-box textarea {
    resize: vertical;
    height: 90px;
    min-height: 50px;
    max-height: 300px;
  }

  .input-box .eye {
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

  .input-box .eye:focus {
    border: none;
    outline: none;
  }

  :global(.input-box .eye:hover svg) {
    fill: var(--orange-005);
  }

  .input_percentage {
    display: flex;
    position: relative;
  }

  .input_percentage input {
    padding-right: 30px !important;
  }
</style>