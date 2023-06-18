<script>
  export let shares = []
</script>

{#each shares as share}
	<span style="display:inline-block; background-color: {share.color}"> &nbsp; </span> {share.label} <br/>
{/each}