<script>
  
  
  let error = '#{error}';
  let email = '#{email}';
  const createdAt = new Date( +'#{createdAt}' );
  const now = new Date()
  const recentlyRegistered = now - createdAt < 1000 * 60 * 10; // 10 minutes

</script>

<h1>/*! title */</h1>

{#if error}
	<h2 class="error">{error}</h2>
{:else}
	<h2>{email} { recentlyRegistered ? 'registered and' : ''} logged in successfully</h2>
{/if}

<hr/>
<a href="/">Back to main page</a>