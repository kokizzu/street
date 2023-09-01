<script>
  import GoogleSdk from './GoogleSdk.svelte'
  import { mapStore} from "./stores";

  let mapElement
  let map
  export let options = {}
  export function setCentre(location) {
    map.setCenter(location)
  }
  async function initialise () {
    const {Map} = await google.maps.importLibrary( 'maps' );
    map = new Map( mapElement, options);
    mapStore.set(map);
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