// store

import { writable } from 'svelte/store';

export let isSideMenuOpen = writable( false );

export let stackPageCount = writable(1);

// export let stackPageRealtor = writable([
//    {
//       subroute: 'main',
//       attrs: {
//          components: [],
//          prop1: 'haha',
//          prop2: 'hehe'
//       }
//    }
// ]);