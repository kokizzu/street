<script>
  export let options = [];
  export let selected = '';
</script>

<div class="option_container">
  {#each options as option}
    <label class={selected === option ? 'option clicked': 'option'} for={option}>
      <input type="radio" on:click={() => (selected = option)} id={option} value={option} />
      {option}
    </label>
  {/each}
</div>

<style>
 .option_container {
    width: 100%;
    display: grid;
    grid-template-rows: 1fr 1fr;
    gap: 10px;
    text-transform: capitalize;
  }
 .option_container .option {
    margin: 0;
    padding: 10px 12px;
    border-radius: 8px;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    font-weight: 500;
    text-align: center;
    cursor: pointer;
  }
 .option_container .option:hover {
    border: 1px solid #f97316;
    color: #f97316;
  }
 .option_container .option.clicked {
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