<script>
  /**
   * @typedef {Object} Props
   * @property {string} [err]
   */

  /** @type {Props} */
  let { err = '' } = $props();
</script>
{#if err}
	<div class="error">{err}</div>
{/if}
<style>
    .error {
        color       : red;
        font-size   : 0.8em;
        padding     : 0.5em;
        font-weight : bold;
    }
</style>