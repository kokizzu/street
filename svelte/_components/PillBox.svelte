<script>
	export let label 		= 'label';
	export let content	= '';
</script>

{#if content}
	<span class="pill-box">
			<span class="label">{label}</span>
			<span class="content">{content}</span>
	</span>
{/if}

<style>
	.pill-box {
		line-height : 25px;
		display     : inline-flex;
	}

	.pill-box .label,
	.pill-box .content {
		border: 1px solid var(--blue-005);
		border-radius : 5px;
		padding: 2px;
		margin-bottom: 1em;
	}

	.pill-box .label {
		background-color : var(--blue-005);
		border-bottom-right-radius: 0;
		border-top-right-radius: 0;
		padding-left: 6px;
		padding-right: 6px;
		text-align: center;
		color: #FFF;
		text-wrap: nowrap;
		padding: auto;
		display: inline-flex;
		align-items: center;
	}

	.pill-box .content {
		background-color: #FFF;
		border-bottom-left-radius: 0;
		border-top-left-radius: 0;
		border-left: 0;
		padding-left: 5px;
		padding-right: 5px;
		margin-right: 1em;
	}
</style>