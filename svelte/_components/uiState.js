import { writable } from 'svelte/store';

export let langOptions = {
  en: 'EN',
  jp: 'JP',
  tw: 'TW' // add more languages here
}
export let isSideMenuOpen = writable( false ); // Side Menu
export let currentLang = writable('EN');