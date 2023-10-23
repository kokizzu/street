<script>
  // @ts-nocheck
  import {onMount} from 'svelte';
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
  let payload = {}
  let countryCurrency = 'TWD';
  onMount( () => {
    console.log( 'onMount.property' );
    console.log( 'property = ', property );
    console.log( 'Country data = ', countries );
    const defaultLat = 23.6978, defaultLng = 120.9605;
    if( Object.keys( property ).length===0 ) {
      property = {
        countryCode: user.countryCode,
        city: '', // new
        countyName: '', // new
        address: '',
        district: '',
        street: '', // new
        street2: '', // new
        floors: 0,
        formattedAddress: '',
        lat: defaultLat,
        lng: defaultLng,
        coord: [defaultLat, defaultLng],
        elevation: 0, // new
        
        uniqPropKey: '1_12449819078726277117',
        serialNumber: '',
        
        houseType: '',
        bedroom: 0,
        bathroom: 0,
        livingroom: 0, // new
        sizeM2: 0,
        purpose: 'sell',
        agencyFeePercent: 0,
        parking: '0', // new
        depositFee: 0, // new
        minimumDurationYear: 0, // new
        otherFee: [], // new
        lastPrice: 0,
        
        images: [],
        imageLabels: [], // new
        createdAt: 1692641835,
        updatedAt: 1692641835,
        
        // floorList: [], // no longer used
        // mainUse: '', // no longer used
        // note: '', // no longer used
        
        //"id": '1234'
        // mainBuildingMaterial: '',
        // constructCompletedDate: '',
        // 'numberOfFloors': '0',
        // buildingLamination: '',
        // createdBy: '0',
        // 'updatedBy': '0',
        // 'deletedAt': 0,
        // priceHistoriesSell: [],
        // priceHistoriesRent: [],
      };
      for( let i = 0; i<countries.length; i++ ) {
        if( countries[ i ].iso_2===property.countryCode ) {
          countryCurrency = countries[ i ].currency.code;
          console.log( 'Country currency =', countryCurrency );
        }
      }
    } else {
      console.log( property.lastPrice );
      if( property.coord && property.coord.length ) {
        property.lat = property.coord[ 0 ];
        property.lng = property.coord[ 1 ];
      }
      property.lastPrice = +property.lastPrice || 0;
      property.agencyFeePercent = +property.agencyFeePercent;
      property.floorList = property.floorList || [];
      property.images = property.images || [];
      property.parking = property.parking.toString();
      for( let i = 0; i<countries.length; i++ ) {
        if( countries[ i ].iso_2===property.countryCode ) {
          countryCurrency = countries[ i ].currency.code;
          console.log( 'Country currency =', countryCurrency );
        }
      }
    }
    console.log( 'Property =', property);
  } );

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
			<EditProperty />
		{:else}
			<CreateProperty {property} {countryCurrency} {payload} {countries}/>
		{/if}
		<Footer></Footer>
	</div>
</section>

