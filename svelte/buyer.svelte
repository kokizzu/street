<script>
	/** @typedef {import('./_types/user').User} User */
	/** @typedef {import('./_types/master').Access} Access */

  import Main from './_layouts/Main.svelte';

	const user = /** @type {User} */ ({/* user */});
	const access = /** @type {Access} */ ({/* segments */});
</script>

<Main {user} {access}>
	<div>
		<span>Buyer</span>
	</div>
</Main>