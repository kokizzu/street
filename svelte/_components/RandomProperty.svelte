<script>
  // @ts-nocheck
  import { UserSearchProp } from '../jsApi.GEN';
  import { onMount } from 'svelte';

  let props = [];

  onMount(async () => {
    console.log('onMount.RandomProperty')
    await UserSearchProp({}, async res => {
      console.log(res);
      props = res.properties;
    });
  });
</script>

<div class="random_property">
  <h3>Random Property</h3>
  <p>{JSON.stringify(props[0])}</p>
  <!-- {#each props as prop}
    <li>{JSON.stringify( prop )}</li>
  {/each} -->
</div>

<style>
  .random_property {
    width: 100%;
    display: flex;
  }
</style>
