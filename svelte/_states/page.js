import { writable } from 'svelte/store';

export const isShrinkMenu = /** @type {import('svelte/store').Writable<boolean>} */ (writable(false));