<script>
  import { createEventDispatcher } from 'svelte';
  import GoogleSdk from "./GoogleSdk.svelte";
  
  const dispatch = createEventDispatcher();
  let mapElement;
  let map;
  
  export let options = {}
  export function setCentre(location) {
    map.setCenter(location)
  }
  async function initialise () {
    const {Map} = await google.maps.importLibrary( 'maps' );
    map = new Map( mapElement, options);
    map.addListener("dragend", () => {
      dispatch("mapDragged", map);
    })
    dispatch('ready');
  }
</script>

<GoogleSdk on:ready={initialise} />
<div class="map" bind:this={mapElement}>
<!--	Maps goes here, will be render automatically -->
</div>

<style>
  .map {
    height: 100%;
    width: 100%;
  }
</style>