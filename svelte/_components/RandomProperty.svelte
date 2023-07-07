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

<div class="random_property">
   <h3>Random Property</h3>
   <ul>
      {#each props as prop}
         <li>{JSON.stringify( prop )}</li>
      {/each}
   </ul>
</div>

<style>
   .random_property {
      position: relative;
      margin-top: -40px;
      margin-left: auto;
      margin-right: auto;
      border-radius: 8px;
      filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
      padding: 25px;
      background-color: white;
      width: 88%;
      height: fit-content;
   }
</style>