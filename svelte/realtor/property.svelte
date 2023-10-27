<script>
  // @ts-nocheck
  import Menu from '../_components/Menu.svelte';
  import ProfileHeader from '../_components/ProfileHeader.svelte';
  import Footer from '../_components/Footer.svelte';
  import CreateProperty from "../_components/CreateProperty.svelte";
  import EditProperty from "../_components/EditProperty.svelte";
  
  // Use property from backend if backend is ready
  let property = {/* property */};
  let user = {/* user */};
  let segments = {/* segments */}
  let countries = [/* countries */];

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

<section class='dashboard'>
	<Menu access={segments}/>
	<div class='dashboard_main_content'>
		<ProfileHeader></ProfileHeader>
		{#if property.id}
			<EditProperty {property} {countries} />
		{:else}
			<CreateProperty {property} {user} {countries}/>
		{/if}
		<Footer></Footer>
	</div>
</section>

