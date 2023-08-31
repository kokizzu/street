<script>
  // @ts-nocheck
  import { onMount } from 'svelte';

  import Menu from './_components/Menu.svelte';
  import ProfileHeader from './_components/ProfileHeader.svelte';
  import Footer from './_components/Footer.svelte';
  import GoogleMap from './_components/GoogleMap/GoogleMap.svelte'

  onMount( async () => {
    await initMap();
  } );
  
  let user = {/* user */};
  let segments = {/* segments */};

  let map;
  let map_container;

  async function initMap() {
    const myLatLng = {lat: -34.397, lng: 150.644};
    // TODO: each coordinate is location of house/apartment for sale
    let markers = [
      {lat: -34.03360401120961, lng: 149.86401361846924},
      {lat: -34.40606480160426, lng: 149.94091791534424},
      {lat: -34.53286884912313, lng: 149.08398432159424},
      {lat: -34.242748228904865, lng: 149.32568354034424},
      {lat: -34.87157170169384, lng: 150.00683588409424},
      {lat: -34.18369473718617, lng: 150.24853510284424},
    ];
    const {Map} = await google.maps.importLibrary( 'maps' );
    map = new Map( map_container, {
      center: myLatLng,
      zoom: 8,
      mapTypeId: 'roadmap',
      mapId: 'street_project',
    } );

    markers.forEach((marker) => {
      new google.maps.Marker({
        position: { lat: marker.lat, lng: marker.lng },
        map: map,
        draggable: true,
      });
    });
  }
</script>

<svelte:head>
  <!-- Google Map SDK -->
  <script>
    (g => {
      var h, a, k, p = 'The Google Maps JavaScript API', c = 'google', l = 'importLibrary', q = '__ib__', m = document, b = window;
      b = b[ c ] || (b[ c ] = {});
      var d = b.maps || (b.maps = {}), r = new Set, e = new URLSearchParams, u = () => h || (h = new Promise( async ( f, n ) => {
        await (a = m.createElement( 'script' ));
        e.set( 'libraries', [...r] + '' );
        for( k in g ) e.set( k.replace( /[A-Z]/g, t => '_' + t[ 0 ].toLowerCase() ), g[ k ] );
        e.set( 'callback', c + '.maps.' + q );
        a.src = `https://maps.${c}apis.com/maps/api/js?` + e;
        d[ q ] = f;
        a.onerror = () => h = n( Error( p + ' could not load.' ) );
        a.nonce = m.querySelector( 'script[nonce]' )?.nonce || '';
        m.head.append( a );
      } ));
      d[ l ] ? console.warn( p + ' only loads once. Ignoring:', g ) : d[ l ] = ( f, ...n ) => r.add( f ) && u().then( () => d[ l ]( f, ...n ) );
    })( {
      key: 'AIzaSyBKF5w6NExgYbmNMvlbMqF6sH2X4dFvMBg',
      v: 'weekly',
      // Use the 'v' parameter to indicate the version to use (weekly, beta, alpha, etc.).
      // Add other bootstrap parameters as needed, using camel case.
    } );
  </script>
</svelte:head>

<section class="dashboard">
   <Menu
      access={segments}
   />
   <div class="dashboard_main_content">
      <ProfileHeader></ProfileHeader>
      <div class="content">
        <div class='map_container' bind:this={map_container}>
          <!-- Map goes here, rendered automatically -->
        </div>
         <GoogleMap apiKey="AIzaSyBKF5w6NExgYbmNMvlbMqF6sH2X4dFvMBg"/>
      </div>
      <Footer></Footer>
   </div>
</section>

<style>
  .map_container {
    width: 70%;
    height        : 430px;
    position: relative;
    margin-top: -40px;
    margin-left: auto;
    margin-right: auto;
    border-radius: 8px;
  }
</style>