<script>
  import { Loader } from '@googlemaps/js-api-loader';
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let mapId  = /** @type {string} */ ('STREET_MAP')
  export let lat    = /** @type {number} */ (23.775418);
  export let lng    = /** @type {number} */ (120.6665624);
  export let zoom   = /** @type {number} */ (8);

  let map;

  const loader = /** @type {Loader} */ (new Loader({
    apiKey: 'AIzaSyAB19JBmmsKCtZqOZaMhUew9V5RPgmkSXY',
    version: 'weekly',
    libraries: ['places', 'marker', 'maps']
  }));

  let mapElement = /** @type {HTMLDivElement} */ (null);

  loader.importLibrary('maps').then(({ Map }) => {
    map = new Map(mapElement, {
      center: {
        lat: lat,
        lng: lng
      },
      zoom: zoom,
      mapId: mapId
    })
  }).catch((e) => {
    console.log('failed to load google map: ', e)
  }).finally(() => {
    dispatch('ready');
  });

  /**
   * @description Set map's marker
   * @param latitude {number}
   * @param longitude {number}
   * @param title {string}
   * @returns void
   */
  export function SetMarker(latitude, longitude, title) {
    loader.importLibrary('marker')
    .then(( {AdvancedMarkerElement } ) => {
      new AdvancedMarkerElement({
        map: map,
        position: {
          lat: latitude,
          lng: longitude
        },
        title: title
      })
    }).catch((e) => {
      console.log('failed to load marker: ', e)
    });
  }
</script>

<div
  bind:this={mapElement}
  class="map"
>
</div>

<style>
  .map {
    height : 100%;
    width  : 100%;
  }
</style>