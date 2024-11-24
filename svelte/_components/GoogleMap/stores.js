import { writable } from 'svelte/store'

export const mapsLoaded = writable(false);
export const mapsLoading = writable(false);
export let mapComponent = writable(null);

/**
 * @typedef {google.maps.Map} Map
 */
export const gmap = /** @type {import('svelte/store').Writable<Map>} */ (writable(null));

/**
 * @typedef {import('@googlemaps/js-api-loader').Loader} Loader
 */
export const mapLoader = /** @type {import('svelte/store').Writable<Loader>} */ (writable(null));