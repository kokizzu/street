<script>
  // @ts-nocheck
  import {onMount} from 'svelte';
  import Menu from '../_components/Menu.svelte';
  import ProfileHeader from '../_components/ProfileHeader.svelte';
  import Footer from '../_components/Footer.svelte';
  import Growl from '../_components/Growl.svelte';
  import AddOtherFeeDialog from '../_components/AddOtherFeeDialog.svelte';
  import {formatPrice} from '../_components/formatter';
  import {RealtorUpsertProperty} from '../jsApi.GEN';
  import {StreetView} from "../_components/GoogleMap/components";
  
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidAngleLeft from 'svelte-icons-pack/fa/FaSolidAngleLeft';
  import FaSolidMapMarkerAlt from 'svelte-icons-pack/fa/FaSolidMapMarkerAlt';
  import FaSolidFlagUsa from 'svelte-icons-pack/fa/FaSolidFlagUsa';
  import FaSolidHome from 'svelte-icons-pack/fa/FaSolidHome';
  import FaCheckCircle from 'svelte-icons-pack/fa/FaCheckCircle';
  import FaSolidBorderStyle from 'svelte-icons-pack/fa/FaSolidBorderStyle';
  import FaSolidBed from 'svelte-icons-pack/fa/FaSolidBed';
  import FaSolidExchangeAlt from 'svelte-icons-pack/fa/FaSolidExchangeAlt';
  import FaSolidBath from 'svelte-icons-pack/fa/FaSolidBath';
  import FaSolidChair from 'svelte-icons-pack/fa/FaSolidChair';
  import FaSolidCamera from 'svelte-icons-pack/fa/FaSolidCamera';
  import FaSolidTimes from 'svelte-icons-pack/fa/FaSolidTimes';
  import FaSolidCircleNotch from "svelte-icons-pack/fa/FaSolidCircleNotch";
  
  // Use property from backend if backend is ready
  let property = {/* property */};
  let user = {/* user */};
  let segments = {/* segments */};
  let countries = [/* countries */];
  let currentPage = 0, isPropertySubmitted = false;
  let cards = [{}, {}, {}, {}];
  let showGrowl = false, gMsg = '', gType = ''; // Growl
  let countryCurrency = 'TWD', payload = {};
  
  onMount( () => {
    console.log( 'property = ', property )
    console.log( 'country data =', countries )
    const defaultLat = 23.6978, defaultLng = 120.9605;
    if( Object.keys( property ).length===0 ) {
      property = {
        country: user.country,
        city: '', // new
        address: '',
        district: '',
        street1: '', // new
        street2: '', // new
        floors: 0,
        formattedAddress: '',
        lat: defaultLat,
        lng: defaultLng,
        coord: [defaultLat, defaultLng],
        elevation: 0, // new
        resolution: 0.1, // new
        
        uniqPropKey: '1_12449819078726277117',
        serialNumber: '',
        
        houseType: '',
        bedroom: 0,
        bathroom: 0,
        livingroom: 0, // new
        sizeM2: 0,
        purpose: 'sell',
        agencyFeePercent: 0,
        parking: 'false', // new
        depositFee: 0, // new
        minimumDurationYear: 0, // new
        otherFee: [], // new
        lastPrice: 0,
        
        images: [],
        imageDescriptions: [], // new
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
      }
    } else {
      console.log( property.lastPrice );
      property.lat = property.coord[ 0 ];
      property.lng = property.coord[ 1 ];
      property.lastPrice = +property.lastPrice || 0;
      property.agencyFeePercent = +property.agencyFeePercent;
      property.floorList = property.floorList || [];
      property.images = property.images || [];
    }
  } );
  
  $: {
    let id = '0'
    if( property.id>0 ) id = '' + property.id;
    payload = {
      // new
      country: property.country,
      city: property.city,
      district: property.district,
      street1: property.street1,
      street2: property.street2 || '', // this means optional
      livingroom: property.livingroom,
      parking: property.parking,
      depositFee: property.depositFee,
      minimumDurationYear: property.minimumDurationYear,
      otherFee: property.otherFee || [],
      imageDescriptions: property.imageDescriptions || [],
      elevation: property.elevation,
      resolution: property.resolution,
      
      // old
      id: id,
      formattedAddress: property.formattedAddress,
      coord: [property.lat, property.lng],
      houseType: property.houseType,
      purpose: property.purpose,
      images: property.images || [],
      bedroom: property.bedroom,
      bathroom: property.bathroom,
      sizeM2: '' + property.sizeM2, // have to be string because of taiwan data
      mainUse: property.mainUse,
      note: property.note,
      lastPrice: '' + property.lastPrice,
      agencyFeePercent: property.agencyFeePercent,
      numberOfFloors: '' + property.floors
    }
  }
  
  function useGrowl( type, msg ) {
    showGrowl = true;
    gMsg = msg;
    gType = type;
    setTimeout( () => {
      showGrowl = false;
    }, 3000 );
  }
  
  function progressDotHandler( toPage ) {
    let card = cards[ toPage ];
    card.scrollIntoView( {behavior: 'smooth'} );
  }
  
  function nextPage() {
    if( currentPage<3 ) {
      currentPage++;
      let card = cards[ currentPage ];
      card.scrollIntoView( {behavior: 'smooth'} );
    }
  }
  
  function backPage() {
    if( currentPage>0 ) {
      currentPage--;
      let card = cards[ currentPage ];
      card.scrollIntoView( {behavior: 'smooth'} );
    }
  }
  
  // +=============| Location and Signage |=============+ //
  const LOC_ADDR = 'Address';
  const LOC_MAP = 'Put the pin on your house location by clicking the map';
  const LOC_STREETVIEW = 'Put signage on your house location';
  let modeLocationCount = 0, countryName = 'Country', countryIso2 = user.country;
  const modeLocationLists = [
    {mode: LOC_ADDR},
    {mode: LOC_MAP},
    {mode: LOC_STREETVIEW},
  ];
  let modeLocation = modeLocationLists[ modeLocationCount ].mode;
  let map, map_container, input_address
  
  async function initMap() {
    const {Map} = await google.maps.importLibrary( 'maps' );
    const geocoder = new google.maps.Geocoder();
    map = new Map( map_container, {
      center: {
        lat: property.lat,
        lng: property.lng,
      },
      zoom: 8,
      mapTypeId: 'roadmap',
      mapId: 'street_project',
    } );
    const {SearchBox} = await google.maps.importLibrary( 'places' );
    let searchBox = new SearchBox( input_address );
    map.controls[ google.maps.ControlPosition.TOP_LEFT ].push( input_address );
    let markers = [];
    const markerEventHandler = ( event ) => {
      property.coord[ 0 ] = event.latLng.lat();
      property.coord[ 1 ] = event.latLng.lng();
      getAddress( event.latLng );
    };
    const createMarker = ( map, latLng ) => {
      let marker = new google.maps.Marker( {
        map,
        position: latLng,
        draggable: true,
      } );
      marker.listenerHandle = google.maps.event.addListener( marker, 'dragend', markerEventHandler );
      return marker;
    };
    const clearMarkers = ( markers ) => {
      markers.forEach( ( marker ) => {
        marker.setMap( null );
        if( marker.listenerHandle && 'function'=== typeof marker.listenerHandle.remove ) marker.listenerHandle.remove();
      } );
      markers.length = 0;
    };
    markers.push( createMarker( map, {
      lat: property.coord[ 0 ] || 0,
      lng: property.coord[ 1 ] || 0,
    } ) );
    const getAddress = ( latLng ) => {
      geocoder.geocode( {location: latLng}, ( results, status ) => {
        if( status===google.maps.GeocoderStatus.OK && results.length>0 ) {
          property.formattedAddress = results[ 0 ].formatted_address;
          for( let i = 0; i<results[ 0 ].address_components.length; i++ ) {
            if( results[ 0 ].address_components[ i ].types.indexOf( 'country' )!== -1 ) {
              property.country = results[ 0 ].address_components[ i ].short_name;
              countryName = results[ 0 ].address_components[ i ].long_name;
            }
          }
        } else {
          useGrowl( 'error', 'Address not found' );
          property.formattedAddresss = '';
          property.country = '';
        }
      } );
    };
    map.addListener( 'click', ( event ) => {
      clearMarkers( markers );
      const latLong = event.latLng;
      markers = [createMarker( map, latLong )];
      // Update data structure
      property.coord[ 0 ] = latLong.lat();
      property.coord[ 1 ] = latLong.lng();
      getAddress( latLong );
      // Callback for dragend event
    } );
    // Searchable map
    map.addListener( 'bounds_changed', () => {
      searchBox.setBounds( map.getBounds() );
    } );
    searchBox.addListener( 'places_changed', () => {
      const places = searchBox.getPlaces();
      if( places.length==0 ) {
        return;
      }
      // Fill formatted_address, latitude, and longitude as JSON values
      property.lng = places[ 0 ].geometry.location.lng();
      property.lat = places[ 0 ].geometry.location.lat();
      property.formattedAddress = places[ 0 ].formatted_address;
      for( let i = 0; i<places[ 0 ].address_components.length; i++ ) {
        if( places[ 0 ].address_components[ i ].types.indexOf( 'country' )!== -1 ) {
          property.country = places[ 0 ].address_components[ i ].short_name;
          countryName = places[ 0 ].address_components[ i ].long_name;
        }
      }
      // Clear out the old markers
      clearMarkers( markers );
      // For each place, get the icon, name and location.
      const bounds = new google.maps.LatLngBounds();
      places.forEach( ( place ) => {
        if( !place.geometry || !place.geometry.location ) {
          console.log( 'Returned place contains no geometry' );
          return;
        }
        const icon = {
          url: place.icon,
          size: new google.maps.Size( 71, 71 ),
          origin: new google.maps.Point( 0, 0 ),
          anchor: new google.maps.Point( 17, 34 ),
          scaledSize: new google.maps.Size( 25, 25 ),
        };
        // create marker for each place
        markers.push(
          new google.maps.Marker( {
            map,
            icon,
            title: place.name,
            position: place.geometry.location,
          } ),
        );
        if( place.geometry.viewport ) {
          // Only geocodes have viewport
          bounds.union( place.geometry.viewport );
        } else {
          bounds.extend( place.geometry.location );
        }
        console.log( markers );
      } );
      map.fitBounds( bounds );
    } );
  }
  
  async function handleBackLocation() {
    if( modeLocationCount<modeLocationLists.length ) {
      modeLocationCount -= 1;
      modeLocation = modeLocationLists[ modeLocationCount ].mode;
      if( modeLocation===LOC_MAP ) {
        await initMap();
      }
    }
  }
  
  const handleNextLocation = {
    'LOC_ADDR': async () => {
      property.country = countryIso2;
      if( property.city==='' || property.street1==='' || property.floors===0 ) {
        useGrowl( 'error', 'Please fill required form' );
        return;
      }
      countries.forEach( ( c ) => {
        if( c.iso_2===property.country ) {
          countryName = c.country
          property.lat = parseInt( c.coordinate.lat );
          property.lng = parseInt( c.coordinate.lng );
        }
      } );
      modeLocationCount += 1;
      modeLocation = modeLocationLists[ modeLocationCount ].mode;
      await initMap();
    },
    'LOC_MAP': () => {
      if( property.formattedAddress==='' ) {
        useGrowl( 'error', 'Please mark location on map' );
        return;
      }
      modeLocationCount += 1;
      modeLocation = modeLocationLists[ modeLocationCount ].mode;
    },
    'LOC_STREETVIEW': () => {
      nextPage();
    },
  };
  
  // +================| Info |=================+ //
  const INFO_FEAT = 'Feature', INFO_PRICE = 'Price';
  const m2 = 'M2', ping = 'Ping';
  let modeInfoCount = 0, infoUnitMode = m2, houseSize = 0, houseSizeM2 = 0, houseSizePing = 0;
  const modeInfoLists = [
    {mode: INFO_FEAT},
    {mode: INFO_PRICE},
  ];
  let modeInfo = modeInfoLists[ modeInfoCount ].mode;
  let infoObj = {
    agencyFee: 'false',
    deposit: 'false',
    parking: property.parking,
    houseType: property.houseType
  };
  let otherFeeObj = {
    name: '',
    fee: 0,
  };
  let addOtherFeeDialog = AddOtherFeeDialog;
  let houseTypeLists = [
    'house', 'land', 'apartment', 'townhouse', 'condo', 'villa', 'factory', 'parking', 'other'
  ];
  
  function addOtherFee() {
    property.otherFee = [...property.otherFee, otherFeeObj];
    otherFeeObj = {
      name: '',
      fee: 0,
    };
    addOtherFeeDialog.hideModal();
  }
  
  const handleInfoUnitMode = {
    'toggle': () => {
      if( infoUnitMode===ping ) {
        if( houseSize!==0 ) {
          houseSizePing = houseSize;
          houseSizeM2 = handleInfoUnitMode.pingToM2( houseSize );
          infoUnitMode = m2;
          houseSize = houseSizeM2;
        } else {
          infoUnitMode = m2;
        }
      } else if( infoUnitMode===m2 ) {
        if( houseSize!==0 ) {
          houseSizeM2 = houseSize;
          houseSizePing = handleInfoUnitMode.m2ToPing( houseSize );
          infoUnitMode = ping;
          houseSize = houseSizePing;
        } else {
          infoUnitMode = ping;
        }
      }
    },
    'm2ToPing': ( m2 ) => {
      const value = m2 / 3.30579;
      const minifiedValue = value.toFixed( 2 );
      return parseFloat( minifiedValue );
    },
    'pingToM2': ( ping ) => {
      const value = ping * 3.30579;
      const minifiedValue = value.toFixed( 2 );
      return parseFloat( minifiedValue );
    },
  };
  
  function handleBackInfo() {
    if( modeInfoCount>0 ) {
      modeInfoCount -= 1;
      modeInfo = modeInfoLists[ modeInfoCount ].mode;
      return;
    }
    backPage();
  }
  
  const handleNextInfo = {
    'INFO_FEAT': () => {
      property.parking = infoObj.parking;
      property.houseType = infoObj.houseType;
      if( infoUnitMode===ping ) {
        property.sizeM2 = handleInfoUnitMode.pingToM2( houseSize );
      }
      if( infoUnitMode===m2 ) {
        property.sizeM2 = houseSize;
      }
      if( property.sizeM2===0 ) return useGrowl( 'error', 'Please fill required form' );
      modeInfoCount += 1;
      modeInfo = modeInfoLists[ modeInfoCount ].mode;
    },
    'INFO_PRICE': () => {
      if( property.lastPrice===0 ) return useGrowl( 'error', 'Price cannot be 0' );
      nextPage();
    },
  };
  
  // +================| Picture |=================+ //
  let imageHouseInput;
  let houseImgUploading = false;
  let uploadHouseStatus = '';
  let uploadHousePercent = 0;
  
  function handlerHouseImage() {
    if( !imageHouseInput ) return;
    const file = imageHouseInput.files[ 0 ];
    if( file ) {
      let formData = new FormData();
      formData.append( 'rawFile', file );
      formData.append( 'purpose', 'property' ); // property or floorPlan
      let ajax = new XMLHttpRequest();
      ajax.addEventListener( 'progress', function( event ) {
        houseImgUploading = true;
        let percent = (event.loaded / event.total) * 100;
        uploadHousePercent = Math.round( percent );
        uploadHouseStatus = `${uploadHousePercent}% uploaded... please wait`;
      } );
      ajax.addEventListener( 'load', function( event ) {
        houseImgUploading = false;
        if( ajax.status===200 ) {
          const out = JSON.parse( event.target.responseText );
          if( !out.error ) {
            property.images = [...property.images, out.urlPattern]; // push house image url to array
            property.imageDescriptions = [...property.imageDescriptions, ''];
          }
          useGrowl( 'success', 'Image uploaded' );
        } else if( ajax.status===413 ) {
          useGrowl( 'error', 'Image too large' );
        } else {
          useGrowl( 'error', `Error: ${ajax.status}  ${ajax.statusText}` );
        }
      } );
      ajax.addEventListener( 'error', function( event ) {
        useGrowl( 'error', 'Network error' );
      } );
      ajax.addEventListener( 'abort', function( event ) {
        useGrowl( 'error', 'Upload aborted' );
      }, false );
      ajax.open( 'POST', '/user/uploadFile' );
      ajax.send( formData );
    }
    imageHouseInput.value = null;
  }
  
  function removeImage( index ) {
    property.images = property.images.filter( ( _, i ) => i!==index );
    property.imageDescriptions = property.imageDescriptions.filter( ( _, i ) => i!==index );
  }
  
  // SUBMIT =====================+
  let res_propId, submitLoading = false;
  
  async function handleSubmit() {
    submitLoading = true;
    console.log( 'property=', property, 'payload=', payload );
    const prop = {property: payload};
    await RealtorUpsertProperty( prop, function( res ) {
      if( res.error ) {
        submitLoading = false;
        useGrowl( 'error', res.error );
        return
      }
      res_propId = res.property.id;
      isPropertySubmitted = true;
      submitLoading = false;
    } );
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
{#if showGrowl}
	<Growl message={gMsg} growlType={gType}/>
{/if}
<section class='dashboard'>
	<Menu access={segments}/>
	<div class='dashboard_main_content'>
		<ProfileHeader></ProfileHeader>
		<div class='realtor_step_progress_bar'>
			<div class='container'>
				<a class='back_button' href='/realtor'>
					<Icon className='iconBack' color='#475569' size={18} src={FaSolidAngleLeft}/>
				</a>
				<div class='step_wrapper'>
					<div class={currentPage >= 0 ? 'step_item completed' : 'step_item active'}>
						<button on:click={() => progressDotHandler(0)}></button>
						<p>Location</p>
					</div>
					<div class={currentPage === 1 ? 'step_item active' : 'step_item' && currentPage > 1 ? 'step_item completed' : 'step_item'}>
						<button on:click={() => progressDotHandler(1)}></button>
						<p>Info</p>
					</div>
					<div class={currentPage === 2 ? 'step_item active' : 'step_item' && currentPage > 2 ? 'step_item completed' : 'step_item'}>
						<button on:click={() => progressDotHandler(2)}></button>
						<p>Picture</p>
					</div>
					<div class={currentPage === 3 ? 'step_item active' : 'step_item'}>
						<button on:click={() => progressDotHandler(3)}></button>
						<p>Preview</p>
					</div>
				</div>
			</div>
		</div>
		<div class='content'>
			<div class='realtor_subpage_container'>
				<section bind:this={cards[0]} class='location' id='subpage_1'>
					{#if modeLocation!==LOC_ADDR}
						<button
							class='back_button'
							on:click={() => {
                        if (currentPage === 0) {
                        handleBackLocation();
                     } else {
                        backPage()
                     }
						}}>
							<Icon className='iconBack' color='#475569' size={18} src={FaSolidAngleLeft}/>
						</button>
					{/if}
					<div class='subpage_content'>
						<h3>{ modeLocation }</h3>
						{#if modeLocation===LOC_ADDR}
							<div class='address'>
								<div class='row'>
									<div class='input_box'>
										<label for='country'>Country or Region <span class='asterisk'>*</span></label>
										<select id='country' name='country' bind:value={countryIso2}>
											{#each countries as country}
												<option value={country.iso_2}>{country.country}</option>
											{/each}
										</select>
									</div>
									<div class='input_box'>
										<label for='city'>City or County <span class='asterisk'>*</span></label>
										<input id='city' type='text' placeholder='Required' bind:value={property.city}/>
									</div>
								</div>
								<div class='row'>
									<div class='input_box'>
										<label for='street1'>Street 1 <span class='asterisk'>*</span></label>
										<input id='street1' type='text' placeholder='Required' bind:value={property.street1}/>
									</div>
									<div class='input_box'>
										<label for='street2'>Street 2</label>
										<input id='street2' type='text' placeholder='Optional' bind:value={property.street2}/>
									</div>
								</div>
								<div class='row'>
									<div class='input_box'>
										<label for='floors'>Floors <span class='asterisk'>*</span></label>
										<input id='floors' type='number' placeholder='10' min='0' bind:value={property.floors}/>
									</div>
								</div>
							</div>
						{/if}
						{#if modeLocation===LOC_MAP}
							<div class='location_map'>
								<div class='input_box'>
									<label for='input_address'></label>
									<input bind:this={input_address} id='input_address' type='text'/>
								</div>
								<div class='address_country_info'>
									<div class='address'>
										<Icon color='#f97316' size={18} src={FaSolidMapMarkerAlt}/>
										<p>{property.formattedAddress || 'Address'}</p>
									</div>
									<div class='country'>
										<Icon color='#f97316' size={15} src={FaSolidFlagUsa}/>
										<p>{countryName}</p>
									</div>
								</div>
								<div bind:this={map_container} class='map_container'>
									<!-- Map goes here, rendered automatically -->
								</div>
							</div>
						{/if}
						{#if modeLocation===LOC_STREETVIEW}
							<div class='location_streetview'>
								<p class='description'>Let buyers find your house on camera.</p>
								<div class='streetview_container'>
									<!-- TODO: render streetview here -->
									<StreetView bind:elevation={property.elevation} bind:resolution={property.resolution} />
								</div>
							</div>
						{/if}
					</div>
					<button
						class='next_button'
						on:click|preventDefault={() => {
                     if (modeLocationCount === 0) {
                       handleNextLocation.LOC_ADDR()
                     } else if (modeLocationCount === 1) {
                       handleNextLocation.LOC_MAP()
                     } else if (modeLocationCount === 2) {
                       handleNextLocation.LOC_STREETVIEW()
                     }
                  }}>
						<span>NEXT</span>
					</button>
				</section>
				<section bind:this={cards[1]} class='info' id='subpage_2'>
					<AddOtherFeeDialog
						bind:fee={otherFeeObj.fee}
						bind:name={otherFeeObj.name}
						bind:this={addOtherFeeDialog}
					>
						<button class='add_fee_btn' on:click={addOtherFee}>
							Add
						</button>
					</AddOtherFeeDialog>
					<button class='back_button' on:click={handleBackInfo}>
						<Icon className='iconBack' color='#475569' size={18} src={FaSolidAngleLeft}/>
					</button>
					<div class='subpage_content'>
						<h3>{ modeInfo }</h3>
						{#if modeInfo===INFO_FEAT}
							<div class='feature'>
								<div class='inputs'>
									<div class='row'>
										<div class='input_box'>
											<label for='houseType'>House type <span class='asterisk'>*</span></label>
											<select id='houseType' name='houseType' bind:value={infoObj.houseType}>
												{#each houseTypeLists as t}
													<option value={t}>{t}</option>
												{/each}
											</select>
										</div>
										<div class='input_box'>
											<label for='parking'>Parking <span class='asterisk'>*</span></label>
											<select id='parking' name='parking' bind:value={infoObj.parking}>
												<option value='true'>Yes</option>
												<option value='false'>No</option>
											</select>
										</div>
									</div>
									<div class='room_area'>
										{#if property.houseType!=='land'}
											<div class='input_box beds'>
												<label for='beds'>
													<Icon color='#475569' size={16} src={FaSolidBed}/>
													<span>Beds</span>
												</label>
												<input id='beds' type='number' min='0' bind:value={property.bedroom}/>
											</div>
											<div class='input_box baths'>
												<label for='baths'>
													<Icon color='#475569' size={13} src={FaSolidBath}/>
													<span>Baths</span>
												</label>
												<input id='baths' type='number' min='0' bind:value={property.bathroom}/>
											</div>
											<div class='input_box livings'>
												<label for='livings'>
													<Icon color='#475569' size={13} src={FaSolidChair}/>
													<span>Livings</span>
												</label>
												<input id='livings' type='number' min='0' bind:value={property.livingroom}/>
											</div>
										{/if}
										<div class='input_box area'>
											<label for='area'>
												<Icon color='#475569' size={13} src={FaSolidBorderStyle}/>
												<span>{infoUnitMode} <span class='asterisk'>*</span></span>
												<button class='unit_toggle' on:click={handleInfoUnitMode.toggle}>
													<span class='bg'></span>
													<Icon color='#F97316' size={13} src={FaSolidExchangeAlt}/>
												</button>
											</label>
											<input id='area' type='number' min='0' bind:value={houseSize}/>
										</div>
									</div>
								</div>
							</div>
						{/if}
						{#if modeInfo===INFO_PRICE}
							<div class='price'>
								<div class='rent_or_sell'>
									<label class={property.purpose === 'sell' ? 'option clicked': 'option'} for='sell'>
										<input
											type='radio'
											name='rent_or_sell'
											on:click={() => (property.purpose = 'sell')}
											id='sell'
											value='sell'
										/>
										Sell
									</label>
									<label class={property.purpose === 'rent' ? 'option clicked': 'option'} for='rent'>
										<input
											type='radio'
											name='rent_or_sell'
											on:click={() => (property.purpose = 'rent')}
											id='rent'
											value='rent'
										/>
										Rent
									</label>
								</div>
								<div class='row'>
									<div class='input_box prop_price'>
										<label for='price'>{property.purpose==='sell' ? 'Property Price' : 'Rent'}</label>
										<input id='price' type='number' min='0' bind:value={property.lastPrice}/>
										<span>$</span>
									</div>
									{#if property.purpose==='rent'}
										<p class='permonth'>/month</p>
									{/if}
								</div>
								<div class='row'>
									<div class='input_box'>
										<label for='agency_fee'>Agency Fee</label>
										<select id='agency_fee' bind:value={infoObj.agencyFee}>
											<option value='true'>Yes</option>
											<option value='false'>No</option>
										</select>
									</div>
									{#if infoObj.agencyFee==='true'}
										<div class='input_box agency_fee'>
											<label for='agency_fee_percent'>Charge to Buyer</label>
											<input id='agency_fee_percent' type='number' min='0' max='99' bind:value={property.agencyFeePercent}/>
											<span>%</span>
										</div>
									{/if}
								</div>
								{#if property.purpose==='rent'}
									<div class='row'>
										<div class='input_box'>
											<label for='deposit'>Deposit Fee</label>
											<select id='deposit' bind:value={infoObj.deposit}>
												<option value='true'>Yes</option>
												<option value='false'>No</option>
											</select>
										</div>
										{#if infoObj.deposit==='true'}
											<div class='input_box deposit_fee'>
												<label for='deposit_fee'>Fee</label>
												<input id='deposit_fee' type='number' min='0' bind:value={property.depositFee}/>
												<span>$</span> <!-- TODO: Use current country currency sign -->
											</div>
										{/if}
									</div>
									<div class='row'>
										<div class='input_box min_duration'>
											<label for='minimum_duration'>Minimum Duration</label>
											<input id='minimum_duration' type='number' min='1' bind:value={property.minimumDurationYear}/>
											<span>Year</span>
										</div>
									</div>
									<div class='other_fee'>
										<header>
											<h4>Other Fee</h4>
											<button class='add_fee' on:click={addOtherFeeDialog.showModal()}>
												Add
											</button>
										</header>
										<div class='other_fee_lists'>
											{#if property.otherFee && property.otherFee.length}
												{#each property.otherFee as otherFee}
													<div class='fee'>
														<span>{otherFee.name}</span>
														<b>{formatPrice( otherFee.fee, countryCurrency )}/mo</b>
													</div>
												{/each}
											{:else}
												<div class='fee'>
													<span>Example Fee #1</span>
													<span>$$$</span>
												</div>
												<div class='fee'>
													<span>Example Fee #1</span>
													<span>$$$</span>
												</div>
											{/if}
										</div>
									</div>
								{/if}
							</div>
						{/if}
					</div>
					<button
						class='next_button'
						on:click|preventDefault={() => {
                     if (modeInfoCount === 0) {
                       handleNextInfo.INFO_FEAT();
                     } else if (modeInfoCount === 1) {
                       handleNextInfo.INFO_PRICE();
                     }
                  }}>
						<span>NEXT</span>
					</button>
				</section>
				<section bind:this={cards[2]} class='picture' id='subpage_3'>
					<button class='back_button' on:click={backPage}>
						<Icon className='iconBack' color='#475569' size={18} src={FaSolidAngleLeft}/>
					</button>
					<div class='subpage_content'>
						<h3>Upload house photo</h3>
						<div class='upload_picture'>
							<div class='upload_container'>
								<label class='image_upload_button' for='upload_image'>
									{#if !houseImgUploading}
										<input
											bind:this={imageHouseInput}
											on:change={handlerHouseImage}
											type='file'
											accept='image/*'
											id='upload_image'
										/>
										<Icon size={35} className='upload_icon' color='#475569' src={FaSolidCamera}/>
										<p>Select file to Upload</p>
									{:else}
										<progress value={uploadHousePercent} max='100'></progress>
										<p>{uploadHouseStatus}</p>
									{/if}
								</label>
							</div>
							<div class='image_lists'>
								{#if property.images && property.images.length}
									{#each property.images as img, idx}
										<div class='image_card'>
											<div class='image_container'>
												<img alt='' src={img}/>
											</div>
											<div class='image_description'>
												{#if property.imageDescriptions && property.imageDescriptions.length}
													<input placeholder='Description' type='text' bind:value={property.imageDescriptions[idx]}/>
												{:else}
													<input placeholder='Description' type='text'/>
												{/if}
											</div>
											<button
												class='remove_image'
												title='remove this image'
												on:click|preventDefault={() => removeImage(idx)}
											>
												<Icon color='#FFF' size={12} src={FaSolidTimes}/>
											</button>
										</div>
									{/each}
								{/if}
							</div>
						</div>
					</div>
					<button class='next_button' on:click={nextPage}>
						<span>NEXT</span>
					</button>
				</section>
				<section bind:this={cards[3]} class='preview' id='subpage_4'>
					{#if isPropertySubmitted===false}
						<button class='back_button' on:click={backPage}>
							<Icon className='iconBack' color='#475569' size={18} src={FaSolidAngleLeft}/>
						</button>
						<div class='subpage_content'>
							<div class='preview_content'>
								<h3>Preview your property</h3>
								<div class='streetview_container'>
									<!-- TODO: render streetview here -->
									<div class='img_container'>
										<img alt='' src='/assets/img/street-view.jpeg'/>
									</div>
								</div>
								<h4>Property Detail</h4>
								{#if property.images && property.images.length}
									<div class='image_properties'>
										<!-- TODO: render property images as slide? -->
										<div class='img_container'>
											<!-- TODO: click image to zoom -->
											<img alt='' src={property.images[0]}/>
										</div>
									</div>
								{/if}
								<div class='preview_details'>
									<div class='main_details'>
									<span class={property.purpose === 'rent' ? `label_rent purpose` : `label_sale purpose`}>
										{property.purpose==='rent' ? `For ${property.purpose}` : `On Sale`}
									</span>
										<div class='price_house'>
											<div class='left'>
												<div class='price'>
													<h5>{formatPrice( property.lastPrice, countryCurrency )}</h5>
													{ #if property.purpose==='rent'}
														<span>/mo</span>
													{/if}
												</div>
												<span class='agency_fee'>
												Agency Fee: {property.agencyFeePercent}%
											</span>
											</div>
											<div class='right'>
												<div class='house_type'>
													<Icon color='#FFF' size={18} src={FaSolidHome}/>
													<span>{property.houseType}</span>
												</div>
											</div>
										</div>
									</div>
									<div class='feature_details'>
										<div class='feature'>
											<b>{property.bedroom}</b>
											<div>
												<Icon color='#475569' size={16} src={FaSolidBed}/>
												<span>Beds</span>
											</div>
										</div>
										<div class='feature'>
											<b>{property.bathroom}</b>
											<div>
												<Icon color='#475569' size={13} src={FaSolidBath}/>
												<span>Baths</span>
											</div>
										</div>
										<div class='feature'>
											<b>{property.livingroom}</b>
											<div>
												<Icon color='#475569' size={13} src={FaSolidChair}/>
												<span>Livings</span>
											</div>
										</div>
										<div class='feature'>
											<b>{houseSize}</b>
											<div>
												<Icon color='#475569' size={13} src={FaSolidBorderStyle}/>
												<span>{infoUnitMode}</span>
												<button class='unit_toggle' on:click|preventDefault={handleInfoUnitMode.toggle}>
													<span class='bg'></span>
													<Icon color='#F97316' size={13} src={FaSolidExchangeAlt}/>
												</button>
											</div>
										</div>
									</div>
									<div class='other_details'>
										{#if property.purpose==='rent'}
											<div class='details'>
												<h5>Rent Detail</h5>
												<div class='item'>
													<span>Deposit Fee</span>
													<b>{formatPrice( property.depositFee, countryCurrency )}</b>
												</div>
												<div class='item'>
													<span>Minimum Duration</span>
													<b>{property.minimumDurationYear} Year</b>
												</div>
											</div>
											{#if infoObj.otherFee && infoObj.otherFee.length}
												<div class='details'>
													<h5>Other Fee</h5>
													{#each infoObj.otherFee as otherfee}
														<div class='item'>
															<span>{otherfee.name}</span>
															<b>{formatPrice( otherfee.fee, countryCurrency )}/mo</b>
														</div>
													{/each}
												</div>
											{/if}
										{/if}
										<div class='details'>
											<h5>Parking</h5>
											<div class='item'>
												<span>Parking</span>
												<b>{property.parking==='true' ? 'Yes' : 'No'}</b>
											</div>
										</div>
									</div>
								</div>
							</div>
						</div>
						<button class='next_button' on:click|preventDefault={handleSubmit}>
							{#if submitLoading===false}
								<span>SUBMIT</span>
							{/if}
							{#if submitLoading===true}
								<Icon className="spin" color='#FFF' size={15} src={FaSolidCircleNotch}/>
							{/if}
						</button>
					{/if}
					{#if isPropertySubmitted===true}
						<button class='back_button' on:click={() => isPropertySubmitted = false}>
							<Icon className='iconBack' color='#475569' size={18} src={FaSolidAngleLeft}/>
						</button>
						<div class='property_submitted_container'>
							<div class='property_submitted'>
								<div class='icon_submitted'>
									<Icon size={110} color='#059669' src={FaCheckCircle}/>
								</div>
								<div class='message_submmitted'>
									<b>We will review it soon</b>
									<p>Thanks you for submitting your property.</p>
								</div>
							</div>
							<div class='actions'>
								<button class='new' on:click={() =>{window.location.href = "/realtor/property"}}>Create New one</button>
								<button class='see' on:click={() => {window.location.href = `/realtor/ownedProperty/${res_propId}`}}>See the property</button>
							</div>
						</div>
					{/if}
				</section>
			</div>
		</div>
		<Footer></Footer>
	</div>
</section>

<style>
    /* General purpose selector */
    :global(.back_button:hover .iconBack) {
        fill : #EF4444;
    }

    @keyframes spin { /* TODO: use it for loading */
        from {
            transform : rotate(0deg);
        }
        to {
            transform : rotate(360deg);
        }
    }

    :global(.spin) {
        animation : spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
    }

    .realtor_subpage_container section .back_button {
        padding          : 8px;
        border           : none;
        background-color : rgb(0 0 0 / 0.06);
        border-radius    : 5px;
        font-size        : 14px;
        cursor           : pointer;
        display          : flex;
        flex-direction   : row;
        align-items      : center;
        gap              : 8px;
        position         : absolute;
        left             : 20px;
        top              : 20px;
    }

    .realtor_subpage_container section .back_button:hover {
        background-color : rgb(0 0 0 / 0.05);
        color            : #EF4444;
    }

    .realtor_subpage_container section .subpage_content h3 {
        margin     : 8px 0;
        text-align : center;
        font-size  : 20px;
    }

    .input_box {
        display        : flex;
        flex-direction : column;
        gap            : 8px;
        width          : 100%;
    }

    .input_box label {
        font-size   : 13px;
        font-weight : 700;
        margin-left : 10px;
    }

    .input_box input,
    .input_box select {
        width            : 100%;
        border           : 1px solid #CBD5E1;
        background-color : #F1F5F9;
        border-radius    : 8px;
        padding          : 12px;
        text-transform   : capitalize;
    }

    .input_box input:focus,
    .input_box select:focus {
        border-color : #3B82F6;
        outline      : 1px solid #3B82F6;
    }

    .next_button {
        background-color : #F97316;
        color            : white;
        display          : flex;
        justify-content  : center;
        border-radius    : 8px;
        border           : none;
        padding          : 10px;
        cursor           : pointer;
        width            : 100%;
        font-weight      : 600;
    }

    .next_button:disabled,
    .next_button:hover {
        background-color : #F58433;
    }

    .asterisk {
        color      : #EF4444;
        font-size  : 14px;
        margin-top : -3px;
    }

    /* =========================*/

    .realtor_step_progress_bar {
        position         : relative;
        margin-top       : -40px;
        margin-left      : auto;
        margin-right     : auto;
        display          : flex;
        justify-content  : center;
        background-color : #EF4444;
        padding-bottom   : 70px;
    }

    .realtor_step_progress_bar .container {
        width            : 940px;
        margin-left      : auto;
        margin-right     : auto;
        display          : flex;
        flex-direction   : row;
        justify-content  : center;
        align-items      : center;
        background-color : white;
        border-radius    : 10px;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
        padding          : 13px 5%;
    }

    .realtor_step_progress_bar .container .back_button {
        padding          : 8px;
        background-color : rgb(0 0 0 / 0.06);
        border-radius    : 5px;
        font-size        : 14px;
        cursor           : pointer;
        border           : none;
        text-decoration  : none;
        height           : fit-content;
        color            : #334155;
    }

    .realtor_step_progress_bar .container .back_button:hover {
        color : #EF4444;
    }

    .realtor_step_progress_bar .step_wrapper {
        width           : 500px;
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
        padding-top     : 5px;
    }

    .realtor_step_progress_bar .step_wrapper .step_item {
        position        : relative;
        display         : flex;
        flex-direction  : column;
        justify-content : center;
        align-items     : center;
        flex            : 1;
        font-size       : 12px;
        color           : #334155;
    }

    .realtor_step_progress_bar .step_wrapper .step_item::before {
        position      : absolute;
        content       : '';
        border-bottom : 2px solid #CBD5E1;
        width         : 100%;
        top           : 5px;
        left          : -50%;
        z-index       : 2;
    }

    .realtor_step_progress_bar .step_wrapper .step_item::after {
        position      : absolute;
        content       : '';
        border-bottom : 2px solid #CBD5E1;
        width         : 100%;
        top           : 5px;
        left          : 50%;
        z-index       : 2;
    }

    .realtor_step_progress_bar .step_wrapper .step_item:first-child::before {
        content : none;
    }

    .realtor_step_progress_bar .step_wrapper .step_item:last-child::after {
        content : none;
    }

    .realtor_step_progress_bar .step_wrapper .step_item button {
        width            : 13px;
        height           : 13px;
        background-color : #CBD5E1;
        border-radius    : 100%;
        border           : none;
        z-index          : 4;
        cursor           : pointer;
    }

    .realtor_step_progress_bar .step_wrapper .step_item button:hover {
        outline : 5px solid rgb(0 0 0 / 0.09);
    }

    .realtor_step_progress_bar .step_wrapper .step_item p {
        margin : 8px 0 0 0;
    }

    .realtor_step_progress_bar .step_wrapper .step_item.active button {
        width            : 11px;
        height           : 11px;
        background-color : white;
        outline          : 3px solid #F97316;
        cursor           : pointer;
    }

    .step_item.completed::after {
        position      : absolute !important;
        content       : '' !important;
        border-bottom : 2px solid #F97316 !important;
        width         : 100% !important;
        top           : 5px !important;
        left          : 50% !important;
        z-index       : 3 !important;
    }

    .step_item.completed button {
        background-color : #F97316 !important;
    }

    /* Main Container Subpage */
    .realtor_subpage_container {
        color            : #475569;
        position         : relative;
        margin-top       : -40px;
        margin-left      : auto;
        margin-right     : auto;
        display          : flex;
        flex-direction   : row;
        gap              : 20px;
        border-radius    : 10px;
        width            : 940px;
        min-height       : 700px;
        overflow-x       : scroll;
        scroll-snap-type : x mandatory;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    }

    .realtor_subpage_container::-webkit-scrollbar-thumb {
        background-color : #EF4444;
        border-radius    : 4px;
    }

    .realtor_subpage_container::-webkit-scrollbar {
        height : 0;
        width  : 10px;
    }

    .realtor_subpage_container::-webkit-scrollbar-track {
        background-color : transparent;
    }

    .realtor_subpage_container section {
        padding           : 20px;
        background-color  : white;
        border-radius     : 10px;
        min-height        : 700px;
        flex              : 0 0 940px;
        scroll-snap-align : start;
        position          : relative;
        display           : flex;
        gap               : 20px;
        flex-direction    : column;
        justify-content   : space-between;
    }

    /* +============| SUBPAGE LOCATION |===========+ */
    .realtor_subpage_container section.location .subpage_content {
        display        : flex;
        flex-direction : column;
        flex-grow      : 1;
    }

    .realtor_subpage_container section.location .subpage_content .address {
        margin-top     : 20px;
        display        : flex;
        flex-direction : column;
        gap            : 15px;
    }

    .realtor_subpage_container section.location .subpage_content .address .row {
        display               : grid;
        grid-template-columns : 1fr 1fr;
        gap                   : 20px;
    }

    .realtor_subpage_container section.location .location_map {
        height         : 100%;
        width          : 100%;
        display        : flex;
        flex-direction : column;
        gap            : 10px;
        margin-top     : 20px;
    }

    .realtor_subpage_container section.location .location_map #input_address {
        margin-top    : 10px;
        padding       : 12px 15px;
        border-radius : 2px;
        border        : none;
        width         : 400px;
        filter        : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    }

    .realtor_subpage_container section.location .location_map #input_address:focus {
        border-color : #F97316;
        outline      : 2px solid #F97316;
    }

    .realtor_subpage_container section.location .location_map .address_country_info {
        display         : flex;
        flex-direction  : row;
        flex-wrap       : wrap;
        justify-content : space-between;
        gap             : 10px;
    }

    .realtor_subpage_container section.location .location_map .address_country_info .country,
    .realtor_subpage_container section.location .location_map .address_country_info .address {
        display        : flex;
        flex-direction : row;
        gap            : 10px;
        align-items    : center;
        margin         : 0 !important;
    }

    .realtor_subpage_container section.location .location_map .address_country_info p {
        margin    : 0;
        font-size : 15px;
    }

    .realtor_subpage_container section.location .location_map .map_container {
        border        : 1px solid #CBD5E1;
        display       : block;
        position      : relative;
        width         : 100%;
        height        : 100%;
        flex-grow     : 1;
        border-radius : 8px;
        overflow      : hidden;
    }

    .realtor_subpage_container section.location .location_streetview {
        display        : flex;
        flex-direction : column;
        gap            : 20px;
        flex-grow      : 1;
    }

    .realtor_subpage_container section.location .location_streetview .description {
        font-size   : 17px;
        text-align  : center;
        color       : #636C77;
        line-height : 1em;
        margin      : 0;
    }

    .realtor_subpage_container section.location .location_streetview .streetview_container {
        flex-grow : 1;
	     position: relative;
    }
    
    /* +============| SUBPAGE INFO |===========+ */

    .realtor_subpage_container section.info .add_fee_btn {
        background-color : #F97316;
        color            : white;
        border-radius    : 8px;
        border           : none;
        padding          : 10px;
        cursor           : pointer;
        width            : 100%;
    }

    .realtor_subpage_container section.info .subpage_content .feature {
        margin-top : 50px;
    }

    .realtor_subpage_container section.info .subpage_content .feature .inputs {
        display        : flex;
        flex-direction : column;
        gap            : 30px;
        margin-top     : 20px;
    }

    .realtor_subpage_container section.info .subpage_content .feature .inputs .row {
        display               : grid;
        grid-template-columns : 1fr 1fr;
        gap                   : 20px;
    }

    .realtor_subpage_container section.info .subpage_content .feature .inputs .room_area {
        display        : flex;
        flex-direction : row;
        gap            : 20px;
    }

    .realtor_subpage_container section.info .subpage_content .feature .inputs .room_area .beds,
    .realtor_subpage_container section.info .subpage_content .feature .inputs .room_area .baths,
    .realtor_subpage_container section.info .subpage_content .feature .inputs .room_area .livings {
        flex-basis : 20%;
    }

    .realtor_subpage_container section.info .subpage_content .feature .inputs .room_area .beds label,
    .realtor_subpage_container section.info .subpage_content .feature .inputs .room_area .baths label,
    .realtor_subpage_container section.info .subpage_content .feature .inputs .room_area .livings label,
    .realtor_subpage_container section.info .subpage_content .feature .inputs .room_area .area label {
        display         : flex !important;
        flex-direction  : row !important;
        justify-content : left !important;
        align-items     : center !important;
        gap             : 8px !important;
    }

    .realtor_subpage_container section.info .subpage_content .feature .inputs .room_area .area {
        flex-basis : 40%;
    }

    .realtor_subpage_container section .unit_toggle {
        border     : none;
        background : transparent;
        position   : relative;
        cursor     : pointer;
    }

    .realtor_subpage_container section .unit_toggle .bg {
        width            : 0;
        height           : 0;
        border-radius    : 50%;
        background-color : rgb(0 0 0 / 0.06);
        z-index          : 1;
        position         : absolute;
        top              : -4px;
        left             : 1px;
    }

    .realtor_subpage_container section .unit_toggle:hover .bg {
        width  : 24px;
        height : 24px;
    }

    .realtor_subpage_container section.info .subpage_content .price {
        margin-top     : 30px;
        display        : flex;
        flex-direction : column;
        gap            : 20px;
    }

    .realtor_subpage_container section.info .subpage_content .price .permonth {
        line-height : 1em;
        margin      : 20px 0 0;
        font-size   : 17px;
        font-weight : 500
    }

    .realtor_subpage_container section.info .subpage_content .price .rent_or_sell {
        width                 : 100%;
        display               : grid;
        grid-template-columns : 1fr 1fr;
        border-collapse       : collapse;
    }

    .realtor_subpage_container section.info .subpage_content .price .rent_or_sell .option {
        margin           : 0;
        padding          : 10px 12px;
        border           : 1px solid #CBD5E1;
        background-color : #F1F5F9;
        font-weight      : 500;
        text-align       : center;
        cursor           : pointer;
    }

    .realtor_subpage_container section.info .subpage_content .price .rent_or_sell .option:nth-child(1) {
        border-top-left-radius    : 8px;
        border-bottom-left-radius : 8px;
    }

    .realtor_subpage_container section.info .subpage_content .price .rent_or_sell .option:nth-child(2) {
        border-top-right-radius    : 8px;
        border-bottom-right-radius : 8px;
    }

    .realtor_subpage_container section.info .subpage_content .price .rent_or_sell .option:hover {
        border : 1px solid #F97316;
        color  : #F97316;
    }

    .realtor_subpage_container section.info .subpage_content .price .rent_or_sell .option.clicked {
        background-color : #F97316;
        color            : white;
        border           : none;
    }

    .realtor_subpage_container section.info .subpage_content .price .rent_or_sell .option input[type='radio'] {
        position       : absolute;
        opacity        : 0;
        pointer-events : none;
    }

    .realtor_subpage_container section.info .subpage_content .price .row {
        display               : grid;
        grid-template-columns : 1fr 1fr;
        gap                   : 20px;
        align-items           : center;
    }

    .realtor_subpage_container section.info .subpage_content .price .row .prop_price,
    .realtor_subpage_container section.info .subpage_content .price .row .agency_fee,
    .realtor_subpage_container section.info .subpage_content .price .row .deposit_fee,
    .realtor_subpage_container section.info .subpage_content .price .row .min_duration {
        position : relative;
    }

    .realtor_subpage_container section.info .subpage_content .price .row .agency_fee input {
        padding : 12px 25px 12px 12px;
    }

    .realtor_subpage_container section.info .subpage_content .price .row .agency_fee span {
        position  : absolute;
        right     : 10px;
        top       : 35px;
        font-size : 15px;
        color     : black;
    }

    .realtor_subpage_container section.info .subpage_content .price .row .prop_price input,
    .realtor_subpage_container section.info .subpage_content .price .row .deposit_fee input {
        padding : 12px 12px 12px 25px;
    }

    .realtor_subpage_container section.info .subpage_content .price .row .prop_price span,
    .realtor_subpage_container section.info .subpage_content .price .row .deposit_fee span {
        position  : absolute;
        left      : 10px;
        top       : 35px;
        font-size : 14px;
        color     : black;
    }

    .realtor_subpage_container section.info .subpage_content .price .row .min_duration input {
        padding : 12px 42px 12px 12px;
    }

    .realtor_subpage_container section.info .subpage_content .price .row .min_duration span {
        position  : absolute;
        right     : 10px;
        top       : 35px;
        font-size : 15px;
        color     : #F97316;
    }

    .realtor_subpage_container section.info .subpage_content .price .other_fee {
        display        : flex;
        flex-direction : column;
        gap            : 15px;
    }

    .realtor_subpage_container section.info .subpage_content .price .other_fee header {
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
        align-items     : center;
        margin-top      : 20px;
    }

    .realtor_subpage_container section.info .subpage_content .price .other_fee header h4 {
        margin    : 0;
        font-size : 18px;
    }

    .realtor_subpage_container section.info .subpage_content .price .other_fee header .add_fee {
        color            : #F97316;
        border           : none;
        padding          : 10px 15px;
        font-weight      : 600;
        font-size        : 15px;
        border-radius    : 8px;
        cursor           : pointer;
        background-color : transparent;
    }

    .realtor_subpage_container section.info .subpage_content .price .other_fee header .add_fee:hover {
        background-color : #F1F5F9;
    }

    .realtor_subpage_container section.info .subpage_content .price .other_fee .other_fee_lists {
        padding          : 15px;
        background-color : #F1F5F9;
        border-radius    : 8px;
        display          : flex;
        flex-direction   : column;
        min-height       : 200px;
    }

    .realtor_subpage_container section.info .subpage_content .price .other_fee .other_fee_lists .fee {
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
        padding         : 10px 0;
        border-bottom   : 1px solid #CBD5E1;
        font-size       : 15px;
        font-weight     : 500;
    }

    /* +============| SUBPAGE PICTURE |===========+ */

    .realtor_subpage_container section.picture .subpage_content .upload_picture {
        margin-top     : 30px;
        display        : flex;
        flex-direction : column;
        gap            : 20px;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .upload_container {
        width  : 100%;
        height : fit-content;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .upload_container .image_upload_button {
        display          : flex;
        flex-direction   : column;
        justify-content  : center;
        align-items      : center;
        border           : 1px solid #CBD5E1;
        background-color : #F1F5F9;
        border-radius    : 8px;
        gap              : 8px;
        width            : 100%;
        height           : 110px;
        cursor           : pointer;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .upload_container .image_upload_button:hover {
        border : 1px solid #F97316;
        color  : #F97316;
    }

    :global(.realtor_subpage_container section.picture .subpage_content .upload_picture .upload_container .image_upload_button:hover .upload_icon) {
        fill : #F97316;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .upload_container .image_upload_button input {
        position       : absolute;
        opacity        : 0;
        pointer-events : none;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .upload_container .image_upload_button p {
        font-size : 16px;
        margin    : 0;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .upload_container .image_upload_button progress {
        appearance    : none;
        border-radius : 8px;
        height        : 13px;
        overflow      : hidden;
        margin-bottom : 8px;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .upload_container .image_upload_button progress::-webkit-progress-bar {
        background-color   : aliceblue;
        box-shadow         : -1px 1px 10px 0 rgba(0, 0, 0, 0.3) inset;
        -webkit-box-shadow : -1px 1px 10px 0 rgba(0, 0, 0, 0.3) inset;
        -moz-box-shadow    : -1px 1px 10px 0 rgba(0, 0, 0, 0.3) inset;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .upload_container .image_upload_button progress::-webkit-progress-value {
        background-color : #F97316;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .upload_container .image_upload_button progress::-moz-progress-bar {
        background-color : #F97316;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .image_lists {
        display               : grid;
        gap                   : 20px;
        grid-auto-columns     : 1fr 1fr;
        grid-auto-rows        : 1fr 1fr;
        grid-template-columns : 1fr 1fr;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .image_lists .image_card {
        display          : flex;
        position         : relative;
        flex-direction   : column;
        padding          : 15px;
        background-color : #F1F5F9;
        border-radius    : 10px;
        gap              : 15px;
        border           : 1px solid #CBD5E1;
        height           : fit-content;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .image_lists .image_card .image_container {
        width         : 100%;
        height        : 250px;
        overflow      : hidden;
        border-radius : 8px;
        border        : 1px solid #CBD5E1;
        cursor        : pointer;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .image_lists .image_card .image_container:hover img {
        transform : scale(1.20);
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .image_lists .image_card .image_container img {
        object-fit          : cover;
        width               : 100%;
        height              : 100%;
        transition-duration : 75ms;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .image_lists .image_card .image_description input {
        width            : 100%;
        border           : 1px solid #CBD5E1;
        background-color : #F1F5F9;
        border-radius    : 8px;
        padding          : 12px;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .image_lists .image_card .image_description input:focus {
        border-color : #3B82F6;
        outline      : 1px solid #3B82F6;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .image_lists .image_card .remove_image {
        position         : absolute;
        right            : 10px;
        top              : 10px;
        padding          : 8px;
        border-radius    : 50%;
        border           : none;
        background-color : #EF4444;
        cursor           : pointer;
    }

    .realtor_subpage_container section.picture .subpage_content .upload_picture .image_lists .image_card .remove_image:hover {
        background-color : #F85454;
    }

    /* +============| SUBPAGE PREVIEW |===========+ */

    .realtor_subpage_container section.preview .subpage_content {
        display        : flex;
        flex-direction : column;
        flex-grow      : 1;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content {
        display        : flex;
        flex-direction : column;
        gap            : 20px;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .img_container {
        height        : 300px;
        width         : 100%;
        overflow      : hidden;
        border        : 1px solid #CBD5E1;
        border-radius : 8px;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .img_container img {
        object-fit : cover;
        width      : 100%;
        height     : 100%;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content h4 {
        margin     : 0;
        font-size  : 20px;
        text-align : center;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details {
        display        : flex;
        flex-direction : column;
        gap            : 20px;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .main_details {
        display        : flex;
        flex-direction : column;
        gap            : 10px;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .main_details .purpose {
        width   : fit-content;
        padding : 4px 10px;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .main_details .price_house {
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .main_details .price_house .left {
        display        : flex;
        flex-direction : column;
        gap            : 5px;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .main_details .price_house .left .price {
        display        : flex;
        flex-direction : row;
        gap            : 5px;
        align-items    : center;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .main_details .price_house .left .price h5 {
        margin    : 0;
        font-size : 35px;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .main_details .price_house .left .price span {
        font-size : 16px;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .main_details .price_house .right {
        height : fit-content;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .main_details .price_house .right .house_type {
        background-color : #F97316;
        display          : flex;
        flex-direction   : row;
        gap              : 8px;
        padding          : 8px 22px;
        color            : #FFF;
        font-size        : 16px;
        font-weight      : 500;
        text-transform   : capitalize;
        border-radius    : 999px;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .feature_details {
        display         : flex;
        flex-direction  : row;
        justify-content : space-evenly;
        align-items     : center;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .feature_details .feature {
        display         : flex;
        flex-direction  : column;
        justify-content : center;
        align-items     : center;
        gap             : 5px;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .feature_details .feature div {
        display        : flex;
        flex-direction : row;
        gap            : 5px;
        align-items    : center;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .feature_details .feature b {
        font-size : 35px;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .other_details {
        display        : flex;
        flex-direction : column;
        gap            : 15px;
        width          : 60%;
        margin         : 20 auto 0;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .other_details .details {
        display        : flex;
        flex-direction : column;
        gap            : 10px;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .other_details .details h5 {
        margin    : 0;
        font-size : 22px;
    }

    .realtor_subpage_container section.preview .subpage_content .preview_content .preview_details .other_details .details .item {
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
        font-weight     : 500;
        padding-bottom  : 7px;
        border-bottom   : 1px solid #CBD5E1;
    }

    .preview .property_submitted_container {
        display        : flex;
        flex-direction : column;
        gap            : 30px;
        height         : 100%;
    }

    .preview .property_submitted_container .property_submitted {
        height          : 100%;
        width           : 100%;
        display         : flex;
        flex-direction  : column;
        justify-content : center;
        align-items     : center;
        text-align      : center;
        gap             : 20px;
        flex-grow       : 1;
    }

    .preview .property_submitted .message_submmitted b {
        font-weight : 700;
        font-size   : 20px;
    }

    .preview .property_submitted_container .actions {
        display        : flex;
        flex-direction : column;
        gap            : 15px;
        height         : fit-content;
    }

    .preview .property_submitted_container .actions button {
        border-radius : 8px;
        border        : none;
        padding       : 10px;
        cursor        : pointer;
        width         : 100%;
        font-weight   : 600;
    }

    .preview .property_submitted_container .actions button.new {
        background-color : #F97316;
        color            : white;
        display          : flex;
        justify-content  : center;
    }

    .preview .property_submitted_container .actions button.new:hover {
        background-color : #F58433;
    }

    .preview .property_submitted_container .actions button.see {
        margin-right     : 10px;
        border           : 1px solid #CBD5E1;
        background-color : #F1F5F9;
    }

    .preview .property_submitted_container .actions button.see:hover {
        border : 1px solid #F97316;
        color  : #F97316;
    }
</style>