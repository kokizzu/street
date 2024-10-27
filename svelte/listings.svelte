<script>
	/** @typedef {import('./_types/user').User} User */

  import Main from './_layouts/Main.svelte';
  import PropertyLocation from './_components/PropertyLocation.svelte';

	const user = /** @type {User} */ ({/* user */});
</script>

<Main {user} >
	<div>
		<PropertyLocation
      initialLatLong={[/* initialLatLong */]}
      randomProps={[/* randomProps */]}
    />
	</div>
</Main>