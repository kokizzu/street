import { writable } from 'svelte/store'

export const mapsLoaded = writable(false);
export const mapsLoading = writable(false);
export let mapComponent = writable(null);