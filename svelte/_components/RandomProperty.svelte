<script>
  import { UserSearchProp } from '../jsApi.GEN';
  import { onMount } from 'svelte';
  
  let props = [];
  
  onMount( async () => {
    await UserSearchProp( {}, async ( res ) => {
      console.log( res );
      props = res.properties;
    } );
  } );
</script>

<h3>Random Property</h3>

<ul>
  {#each props as prop}
    <li>{JSON.stringify( prop )}</li>
  {/each}
</ul>