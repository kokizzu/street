<script>
  import { Loader } from '@googlemaps/js-api-loader';
  import { createEventDispatcher } from 'svelte';
  import { gmap, mapLoader } from './stores';

  const dispatch = createEventDispatcher();

  /**
   * @typedef {Object} Props
   * @property {string} [mapId]
   * @property {number} [lat]
   * @property {number} [lng]
   * @property {number} [zoom]
   */

  /** @type {Props} */
  let {
    mapId = 'STREET_MAP',
    lat = 23.775418,
    lng = 120.6665624,
    zoom = 8
  } = $props();

  let map = /** @type {google.maps.Map} */ (null);
  const loader = /** @type {Loader} */ (new Loader({
    apiKey: 'AIzaSyAB19JBmmsKCtZqOZaMhUew9V5RPgmkSXY',
    version: 'weekly',
    libraries: ['places', 'marker', 'maps']
  }));
  
  mapLoader.set(loader);

  let mapElement = /** @type {HTMLDivElement} */ ($state(null));

  loader.importLibrary('maps').then(({ Map }) => {
    map = new Map(mapElement, {
      center: {
        lat: lat,
        lng: lng
      },
      zoom: zoom,
      mapId: mapId
    });

    gmap.set(map);
  }).catch((e) => {
    console.error('failed to load google map: ', e)
  }).finally(() => {
    dispatch('ready');
  });

  /**
   * @description Set map's marker
   * @param latitude {number}
   * @param longitude {number}
   * @param title {string}
   * @returns {google.maps.marker.AdvancedMarkerElement}
   */
  export function SetMarker(latitude, longitude, title) {
    let marker = /** @type {google.maps.marker.AdvancedMarkerElement} */ (null);

    loader.importLibrary('marker')
    .then(( {AdvancedMarkerElement } ) => {
      marker = new AdvancedMarkerElement({
        map: map,
        position: {
          lat: latitude,
          lng: longitude
        },
        title: title
      })
    }).catch((e) => {
      console.log('failed to load marker: ', e);
    });

    return marker;
  }

  /**
   * @description Delete Marker(s) from array and map
   * @param markers {google.maps.marker.AdvancedMarkerElement[]}
   * @returns {google.maps.marker.AdvancedMarkerElement[]}
   */
  export function DeleteMarkers(markers) {
    (markers || []).forEach(m => m.map = null);
    markers = [];

    return markers;
  }

  /**
   * @description Set map's center by lat, lng
   * @param lat {number}
   * @param lng {number}
   * @returns {void}
   */
  export function SetCenter(lat, lng) {
    map.setCenter({ lat, lng });
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