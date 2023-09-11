<script>
  import { createEventDispatcher } from 'svelte';
  import GoogleSdk from "./GoogleSdk.svelte";
  import { mapComponent } from "./stores";
  
  const dispatch = createEventDispatcher();
  let mapElement, map;
  export let options = {}
  export function setCentre(location) {
    map.setCenter(location);
  }
  export function createMarker(latitude, longitude, iconPath, iconSize, title) {
    let marker = new google.maps.Marker({
      map,
      icon: {
        url: iconPath, // URL to your custom icon image
        scaledSize: new google.maps.Size(iconSize, iconSize),
      },
      position: {
        lat: latitude,
        lng: longitude
      },
      title: title || "",
      draggable: false,
    });
    return marker;
  }
  export function clearMarkers(markers) {
    markers.forEach((marker) => {
      marker.setMap(null);
      if (marker.listenerHandle && 'function' === typeof marker.listenerHandle.remove) {
        marker.listenerHandle.remove();
      }
    });
    markers.length = 0;
    markers = [];
    return markers;
  }
  export function infoWindow(name, address, type) {
    const contentString =
      `<h3 id="firstHeading" class="firstHeading">${name}</h3>` +
      `<p>${address}</p>` +
      `<p><b>Type:</b> ${type}</p>`;
    const infowindow = new google.maps.InfoWindow({
      content: contentString
    });
    return infowindow;
  }
  export function clearInfoWindow( infoWindows ) {
    infoWindows = null;
    return infoWindows;
  }
  async function initialise () {
    const {Map} = await google.maps.importLibrary( 'maps' );
    map = new Map( mapElement, options);
    map.addListener("dragend", () => {
      dispatch("mapDragged", map);
    });
    $mapComponent = map;
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