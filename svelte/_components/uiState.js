// store

import { writable } from 'svelte/store';

export let isSideMenuOpen = writable( false );

export let stackPageRealtor = writable([
   {
      subroute: 'main',
      attrs: {
         // attribute goes here
         prop1: 'haha',
         prop2: 'hehe'
      }
   }
]);