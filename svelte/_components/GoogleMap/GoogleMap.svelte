<script>
	import { createEventDispatcher } from 'svelte';
	import GoogleSdk from './GoogleSdk.svelte';
	import { mapComponent } from './stores';

	const dispatch = createEventDispatcher();

	let mapElement;
	let map;

	export let options = {};
	export let bounds = {}; // bound of map from googlemap (top-left and bottom-right)

	export function setCentre(location) {
		map.setCenter(location);
	}

	export async function createMarker(latitude, longitude) {
		// @ts-ignore
		const markerLib = await google.maps.importLibrary("marker");
		const pin = new markerLib.PinElement({
			scale: 1.5,
	});
		const marker = new markerLib.AdvancedMarkerElement({
			map: map,
			position: {
				lat: latitude,
				lng: longitude
			},
			content: pin.element,
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
			`<h3 style='margin: 0 !important; padding: 15px 5px 10px 0px !important; border-bottom: 1px solid #CBD5E1 !important;'>${name}</h3>` +
			`<p style='margin: 6px 0;'>${address}</p>` +
			`<p style='margin: 6px 0;'><b>Type:</b> ${type}</p>`;
		// @ts-ignore
		const infowindow = new google.maps.InfoWindow({
			content: contentString,
		});
		
		return infowindow;
	}

	export function clearInfoWindow(infoWindows) {
		infoWindows = null;
		return infoWindows;
	}

	async function initialise() { // @ts-ignore
		const {Map} = await google.maps.importLibrary('maps');
		map = new Map(mapElement, options);
		map.addListener('dragend', () => {
			dispatch('mapDragged', map);
		});

		map.addListener('zoom_changed', () => {
			dispatch('zoomChanged', map);
		});

		map.addListener('idle', () => {
			bounds = map.getBounds();
		});

		$mapComponent = map;
		dispatch('ready');
	}
</script>

<GoogleSdk on:ready={initialise} />
<div bind:this={mapElement} class='map'>
    <!--	Maps goes here, will be render automatically -->
</div>

<style>
    .map {
        height : 100%;
        width  : 100%;
    }
</style>