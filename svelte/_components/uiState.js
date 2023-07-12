// store

import { writable } from 'svelte/store';

export let isSideMenuOpen = writable( false );

export let stackPageCount = writable(1);

// export let stackPageRealtor = writable([
//    {
//       subroute: 'location',
//       attrs: {
//          coordinate: (coordinate),
//       }
//    },
//    { TODO, each object will send or push from each subpage}
// ]);