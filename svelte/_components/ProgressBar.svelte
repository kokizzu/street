<script>
  import FloatingNotification from "./FloatingNotification.svelte";
  
  export let modifier = '';
  export let tooltip = '';
  export let value = 0; // 0 to 1
  
  // local state
  let showProgressTooltip = false;
</script>
<div class="progressBar {modifier}"
     on:mouseenter={() => showProgressTooltip = true}
     on:mouseleave={() => showProgressTooltip = false}>
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