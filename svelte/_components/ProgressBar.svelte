<script>
  import FloatingNotification from "./FloatingNotification.svelte";
  
  /**
   * @typedef {Object} Props
   * @property {string} [modifier]
   * @property {string} [tooltip]
   * @property {number} [value] - 0 to 1
   */

  /** @type {Props} */
  let { modifier = '', tooltip = '', value = 0 } = $props();
  
  // local state
  let showProgressTooltip = $state(false);
</script>
<div class="progressBar {modifier}"
     onmouseenter={() => showProgressTooltip = true}
     onmouseleave={() => showProgressTooltip = false}>
	<div class="progressBar_value {modifier}" style="width: {Math.floor(value*100)}%;"></div>
</div>
{#if showProgressTooltip}
	<FloatingNotification text={tooltip}/>
{/if}
<style>
    .progressBar {
        height           : 0.4rem;
        padding-right    : 1em;
        border-radius    : 0.2em;
        background-color : lightgrey;
    }

    .progressBar_value {
        height        : 0.4rem;
        border-radius : 0.2em;
    }

    div.progressBar_value {
        background-color : forestgreen;
    }
</style>