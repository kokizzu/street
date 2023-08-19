<script>
  //@ts-nocheck
  let verified = '#{verified}';
  let error = '#{error}';
  let email = '#{email}';

</script>

<h1>/*! title */</h1>

{#if verified==='true'}
  <h2>{email} is verified</h2>
{/if}

{#if error}
  <h2 class='error'>{error}</h2>
{/if}

<hr />
<a href='/'>Back to main page</a>