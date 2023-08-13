<script>
  // @ts-nocheck
  import { onMount } from 'svelte';
  
  import Menu from './_components/Menu.svelte';
  import ProfileHeader from './_components/ProfileHeader.svelte';
  import Footer from './_components/Footer.svelte';
  import AddFloorDialog from '_components/AddFloorDialog.svelte';
  import AddRoomDialog from '_components/AddRoomDialog.svelte';
  import { RealtorUpsertProperty } from './jsApi.GEN';
  
  onMount( async () => {
    await initMap();
  } );
  
  let user = {/* user */};
  let segments = {/* segments */};
  
  let realtorStack = [
    {
      subroute: 'location',
      attrs: {
        address: '',
        long: 0,
        lat: 0,
      },
    }, {
      subroute: 'information',
      attrs: {
        house_type: '', /*House or Apartment*/
        purpose: '', /*rent or sell*/
        images: [
          // "/url/to/images",
          // "/url/to/images"
        ],
        feature: {
          beds: 0,
          baths: 0,
          area: 0,
        },
        facility: '', /*Facility Description*/
        description: '', /*Description of this property*/
        price: 0, /*Price*/
        agency_fee: 0, /*Fee Percentage*/
      },
    }, {
      subroute: 'floors',
      attrs: {
        floor_lists: [
          {
            type: '', /*Floor or Basement*/
            floor: 0,
            beds: 0,
            baths: 0,
            rooms: [],
          }, {
            type: '', /*Floor or Basement*/
            floor: 0,
            beds: 0,
            baths: 0,
            rooms: [
              {
                name: '',
                sizeM2: 0,
                unit: 'm2',
              },
            ],
          },
        ],
      },
    },
  ];
  let mapStack, infoStack, floorStack, previewStack;
  $: mapStack = realtorStack[ 0 ];
  $: infoStack = realtorStack[ 1 ];
  $: floorStack = realtorStack[ 2 ];
  $: previewStack = realtorStack[ 3 ];
  let currentPage = 0;
  let cards = [{}, {}, {}, {}];
  
  let payload = {};
  $: payload = {
    formattedAddress: mapStack.attrs.address,
    coord: [mapStack.attrs.lat, mapStack.attrs.long],
    houseType: infoStack.attrs.house_type,
    purpose: infoStack.attrs.purpose,
    images: infoStack.attrs.images,
    bedroom: infoStack.attrs.feature.beds,
    bathroom: infoStack.attrs.feature.baths,
    sizeM2: '' + infoStack.attrs.feature.area, // have to be string because of taiwan data
    mainUse: infoStack.attrs.facility,
    note: infoStack.attrs.description,
    price: infoStack.attrs.property_price,
    agencyFeePercent: infoStack.attrs.agency_fee_percent,
    numberOfFloors: '' + floorStack.attrs.floor_lists.length, // have to be string because of taiwan data
    floorList: floorStack.attrs.floor_lists,
  };
  
  async function nextPage() {
    if( currentPage<3 ) {
      currentPage++;
      let card = cards[ currentPage ];
      console.log( card );
      card.scrollIntoView( {behavior: 'smooth'} );
    }
  }
  
  function backPage() {
    if( currentPage>0 ) {
      currentPage--;
    }
  }
  
  // +=============| Location |=============+ //
  let map;
  let map_container;
  let input_address;
  let lng = 0;
  let lat = 0;
  let formatted_address = '';
  
  async function initMap() {
    const myLatLng = {lat: 23.6978, lng: 120.9605}; // taiwan
    let markers = [];
    const {Map} = await google.maps.importLibrary( 'maps' );
    const geocoder = new google.maps.Geocoder();
    map = new Map( map_container, {
      center: myLatLng,
      zoom: 8,
      mapTypeId: 'roadmap',
      mapId: 'street_project',
    } );
    const {SearchBox} = await google.maps.importLibrary( 'places' );
    const searchBox = new SearchBox( input_address );
    map.controls[ google.maps.ControlPosition.TOP_LEFT ].push( input_address );
    // Convert coordinate to formatted_address
    const getAddress = ( latLng ) => {
      geocoder.geocode( {location: latLng}, ( results, status ) => {
        if( status===google.maps.GeocoderStatus.OK && results.length>0 ) {
          formatted_address = results[ 0 ].formatted_address;
        } else {
          console.log( 'Address not found' );
          formatted_address = 'unknown';
        }
      } );
    };
    // Clickable Map
    map.addListener( 'click', ( event ) => {
      markers.forEach( ( marker ) => {
        marker.setMap( null );
      } );
      markers = [];
      markers.push(
        new google.maps.Marker( {
          map,
          position: event.latLng,
          draggable: true,
        } ),
      );
      // Update data structure
      lng = event.latLng.lng();
      lat = event.latLng.lat();
      getAddress( event.latLng );
      // Callback for dragend event
      google.maps.event.addListener( markers[ 0 ], 'dragend', ( event ) => {
        lng = event.latLng.lng();
        lat = event.latLng.lat();
        getAddress( event.latLng );
      } );
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
      lng = places[ 0 ].geometry.location.lng();
      lat = places[ 0 ].geometry.location.lat();
      formatted_address = places[ 0 ].formatted_address;
      // Clear out the old markers
      markers.forEach( ( marker ) => {
        marker.setMap( null );
      } );
      markers = [];
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
  
  function handlerLocationNext() {
    if( lng===0 || lat===0 || formatted_address==='' ) {
      alert( 'Location must be added' );
    } else {
      realtorStack[ currentPage ] = {
        subroute: 'location',
        attrs: {
          address: formatted_address,
          long: lng,
          lat: lat,
        },
      };
      nextPage();
    }
  }
  
  // +=============| House Info |=============+ //
  let house_info_obj = {
    house_type: '',
    floor: 0,
    purpose: '',
    images: [],
    feature: {
      beds: 0,
      baths: 0,
      area: 0,
    },
    facility: '',
    description: '',
    price: {
      property_price: 0,
      agency_fee: 0,
    },
  };
  
  let modeHouseInfoCount = 0;
  const ABOUT_THE_HOUSE = 'About The House';
  const UPLOAD_HOUSE_PHOTO = 'Upload House Photo';
  const FEATURE_OR_FACILITY = 'Feature or Facility';
  const DESCRIPTION_PROPERTY = 'Description of this property';
  const PRICE = 'Price';
  const modeHouseLists = [
    {mode: ABOUT_THE_HOUSE, skip: false},
    {mode: UPLOAD_HOUSE_PHOTO, skip: true},
    {mode: FEATURE_OR_FACILITY, skip: true},
    {mode: DESCRIPTION_PROPERTY, skip: true},
    {mode: PRICE, skip: false},
  ];
  let mode = modeHouseLists[ modeHouseInfoCount ].mode;
  let modeSkippable = modeHouseLists[ modeHouseInfoCount ].skip;
  
  function houseInfoNext() {
    if( modeHouseInfoCount<modeHouseLists.length - 1 ) {
      modeHouseInfoCount++;
      mode = modeHouseLists[ modeHouseInfoCount ].mode;
      modeSkippable = modeHouseLists[ modeHouseInfoCount ].skip;
    } else {
      realtorStack[ currentPage ] = {
        subroute: 'information',
        attrs: house_info_obj,
      };
      nextPage();
    }
  }
  
  function houseInfoBack() {
    if( modeHouseInfoCount>0 ) {
      modeHouseInfoCount--;
      mode = modeHouseLists[ modeHouseInfoCount ].mode;
      modeSkippable = modeHouseLists[ modeHouseInfoCount ].skip;
    } else {
      backPage();
    }
  }
  
  function houseInfoSkip() {
    if( modeHouseInfoCount<modeHouseLists.length - 1 ) {
      modeHouseInfoCount++;
      mode = modeHouseLists[ modeHouseInfoCount ].mode;
      modeSkippable = modeHouseLists[ modeHouseInfoCount ].skip;
    } else {
      nextPage();
    }
  }
  
  // ______About The House
  // **** House Type
  let house_type = '';
  let showHouseOption = false;
  let isApartment = false;
  let floor = 1;
  
  function handleHouseTypeOption() {
    isApartment = house_type==='apartment';
    showHouseOption = !showHouseOption;
  }
  
  // **** Rent or Sell
  let rent_or_sell = '';
  let showRentSellOption = false;
  
  function handleRentSellOption() {
    showRentSellOption = !showRentSellOption;
  }
  
  function handleNextAboutHouse() {
    if( house_type==='' || rent_or_sell==='' ) {
      alert( 'Must fill the form' );
    } else {
      house_info_obj.house_type = house_type;
      house_info_obj.floor = floor;
      house_info_obj.purpose = rent_or_sell;
      houseInfoNext();
    }
  }
  
  // ______Upload House Photo
  let imageHouseInput;
  let house_images = [/*"/url/of/house/image"*/];
  let houseImgUploading = false;
  let uploadHouseStatus = '';
  let uploadHousePercent = 0;
  
  function handlerHouseImage() {
    if( !imageHouseInput ) return;
    const file = imageHouseInput.files[ 0 ];
    if( file ) {
      var formData = new FormData();
      formData.append( 'rawFile', file );
      formData.append( 'purpose', 'property' ); // property or floorPlan
      var ajax = new XMLHttpRequest();
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
            house_images = [...house_images, out.urlPattern]; // push house image url to array
          }
          console.log( 'Upload successful', out );
        } else if( ajax.status===413 ) {
          alert( 'Image too large' );
        } else {
          alert( 'Error: ' + ajax.status + ' ' + ajax.statusText );
        }
        imageHouseInput.value = '';
      } );
      ajax.addEventListener( 'error', function( event ) {
        alert( 'Network error' );
      } );
      ajax.addEventListener( 'abort', function( event ) {
        alert( 'Upload aborted' );
      }, false );
      ajax.open( 'POST', '/user/uploadFile' );
      ajax.send( formData );
    }
  }
  
  function removeImage( index ) {
    house_images = house_images.filter( ( _, i ) => i!==index );
  }
  
  function handleNextUploadHouseImage() {
    house_info_obj.images = house_images;
    houseInfoNext();
  }
  
  // ______Feature and Facility
  let feature = {
    beds: 0,
    baths: 0,
    area: 0,
  };
  let facility = '';
  
  function handleNextFeatureFacility() {
    house_info_obj.feature.beds = feature.beds;
    house_info_obj.feature.baths = feature.baths;
    house_info_obj.feature.area = feature.area;
    house_info_obj.facility = facility;
    houseInfoNext();
  }
  
  // _______Description of Property
  let description_of_property = '';
  
  function handleNextDescriptionProperty() {
    house_info_obj.description = description_of_property;
    houseInfoNext();
  }
  
  // _______Price of Property
  let property_price = 0;
  let agency_fee = 0;
  
  function handleNextPriceProperty() {
    if( property_price===0 ) {
      alert( `Price cannot be "${property_price}"` );
    } else {
      house_info_obj.price = property_price;
      house_info_obj.agency_fee_percent = agency_fee;
      houseInfoNext();
    }
  }
  
  // +=============| Floors |=============+ //
  let add_floor_dialog = AddFloorDialog;
  let floor_type = '';
  let floor_lists = [];
  let floor_attribute = {
    type: '',
    floor: 0,
    beds: 0,
    baths: 0,
    rooms: [/*{ name, sizeM2, unit }*/],
    planImageUrl: '',
  };
  let floorCount = 1;
  let basement_added = false;
  let floor_edit_mode = false;
  let floor_index_to_edit = 0;
  
  function showAddFloorDialog() {
    add_floor_dialog.showModal();
  }
  
  function handlerAddFloor() {
    if( floor_type==='basement' && basement_added===true ) {
      alert( 'basement already added' );
      add_floor_dialog.hideModal();
      return;
    }
    if( floor_type==='basement' ) {
      floor_attribute = {
        type: floor_type,
        floor: 0,
        beds: 0,
        baths: 0,
        rooms: [],
        planImageUrl: '',
      };
      floor_lists = [...floor_lists, floor_attribute];
      basement_added = true;
      floor_type = '';
    } else {
      floor_attribute = {
        type: floor_type,
        floor: floorCount,
        beds: 0,
        baths: 0,
        rooms: [/*{name, sizeM2, unit}*/],
        planImageUrl: '',
      };
      floor_lists = [...floor_lists, floor_attribute];
      floor_type = '';
      floorCount++;
    }
    add_floor_dialog.hideModal();
    return;
  }
  
  function handlerEditFloor( index ) {
    floor_edit_mode = true;
    floor_index_to_edit = index;
  }
  
  // _______Upload Floor Plan photo
  let imgFlrPlanUploading = false;
  let uploadFlrPlanStatus = '';
  let uploadFlrPlanPercent = 0;
  let imageFloorPlanInput;
  
  function handlerImageFloorPlan() {
    const file = imageFloorPlanInput.files[ 0 ];
    if( file ) {
      var formData = new FormData();
      formData.append( 'rawFile', file );
      formData.append( 'purpose', 'floorPlan' ); // property or floorPlan
      var ajax = new XMLHttpRequest();
      ajax.upload.addEventListener( 'progress', function( event ) {
        imgFlrPlanUploading = true;
        let percent = (event.loaded / event.total) * 100;
        uploadFlrPlanPercent = Math.round( percent );
        uploadFlrPlanStatus = `${uploadFlrPlanPercent}% uploaded... please wait`;
      } );
      ajax.addEventListener( 'load', function( event ) {
        imgFlrPlanUploading = false;
        if( ajax.status===200 ) {
          const out = JSON.parse( event.target.responseText );
          if( !out.error ) {
            floor_lists[ floor_index_to_edit ].planImageUrl = out.urlPattern; // .originalUrl also available
          }
          console.log( 'Upload successful', out );
        } else if( ajax.status===413 ) {
          alert( 'Image too large' );
        } else {
          alert( 'Error: ' + ajax.status + ' ' + ajax.statusText );
        }
        imageFloorPlanInput.value = '';
      } );
      ajax.addEventListener( 'error', function( event ) {
        alert( 'Network error' );
      } );
      ajax.addEventListener( 'abort', function( event ) {
        alert( 'Upload aborted' );
      }, false );
      ajax.open( 'POST', '/user/uploadFile' );
      ajax.send( formData );
    }
  }
  
  // ________Rooms Edit
  let add_room_dialog = AddRoomDialog;
  let room_type = '';
  let room_obj = {
    name: '',
    sizeM2: 0,
    unit: 'm2',
  };
  let size_m2;
  let room_size;
  let unit_mode;
  let bedroom_total = 0;
  let bathroom_total = 0;
  
  function showAddRoomDialog() {
    add_room_dialog.showModal();
  }
  
  function handlerAddRoom( index ) {
    if( unit_mode==='SqFt' && size_m2===0 ) {
      size_m2 = add_room_dialog.sqftToM2( room_size );
    } else if( unit_mode==='SqFt' && size_m2!==0 ) {
      size_m2 = add_room_dialog.sqftToM2( room_size );
    } else {
      size_m2 = room_size;
    }
    
    if( room_type==='living room' ) {
      for( let i = 0; i<floor_lists[ index ][ 'rooms' ].length; i++ ) {
        if( floor_lists[ index ][ 'rooms' ][ i ].name==='living room' ) {
          alert( 'Living Room already added' );
          add_room_dialog.hideModal();
          room_type = '';
          room_size = 0;
          size_m2 = 0;
          return;
        }
      }
      
      room_obj = {
        name: 'living room',
        sizeM2: size_m2,
        unit: 'm2',
      };
      floor_lists[ index ].rooms = [...floor_lists[ index ].rooms, room_obj];
      room_type = '';
      room_size = 0;
      size_m2 = 0;
    }
    if( room_type==='bedroom' ) {
      room_obj = {
        name: 'bedroom',
        sizeM2: size_m2,
        unit: 'm2',
      };
      bedroom_total++;
      floor_lists[ index ].rooms = [...floor_lists[ index ].rooms, room_obj];
      floor_lists[ index ].beds = bedroom_total;
      room_type = '';
      room_size = 0;
      size_m2 = 0;
    }
    if( room_type==='bathroom' ) {
      room_obj = {
        name: 'bathroom',
        sizeM2: size_m2,
        unit: 'm2',
      };
      bathroom_total++;
      floor_lists[ index ].rooms = [...floor_lists[ index ].rooms, room_obj];
      floor_lists[ index ].baths = bathroom_total;
      room_type = '';
      room_size = 0;
      size_m2 = 0;
    }
    add_room_dialog.hideModal();
    return;
  }
  
  function handleNextFloor() {
    realtorStack[ currentPage ] = {
      subroute: 'floors',
      attrs: {
        floor_lists: floor_lists,
      },
    };
    nextPage();
  }
  
  // +=============| PREVIEW |=============+ //
  let isPropertySubmitted = false;
  
  function formatPrice( price ) {
    return price.toLocaleString( 'en-US', {
      style: 'currency',
      currency: 'USD',
    } );
  }
  
  async function handleSubmit() {
    console.log( realtorStack, payload );
    const prop = {property: payload};
    await RealtorUpsertProperty( prop, function( res ) {
      console.log( res );
      if( res.error ) return alert( res.error );
      isPropertySubmitted = true; // uncomment when debugging so can back to previous screen
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

<section class='dashboard'>
  <Menu
    access={segments}
  />
  <div class='dashboard_main_content'>
    <ProfileHeader></ProfileHeader>
    <div class='realtor_step_progress_bar'>
      <div class='step_wrapper'>
        <div class={currentPage >= 0 ? 'step_item completed' : 'step_item active'}>
          <span></span>
          <p>Location</p>
        </div>
        <div class={currentPage === 1 ? 'step_item active' : 'step_item' && currentPage > 1 ? 'step_item completed' : 'step_item'}>
          <span></span>
          <p>Info</p>
        </div>
        <div class={currentPage === 2 ? 'step_item active' : 'step_item' && currentPage > 2 ? 'step_item completed' : 'step_item'}>
          <span></span>
          <p>Floor</p>
        </div>
        <div class={currentPage === 3 ? 'step_item active' : 'step_item'}>
          <span></span>
          <p>Preview</p>
        </div>
      </div>
    </div>
    <div class='content'>
      <div class='realtor_subpage_container'>
        <section class='location' id='subpage_1' bind:this={cards[0]}>
          <div>
            <h2>Type property's address or move the map to autofill it</h2>
            <div class='location_input'>
              <div class='input_box'>
                <label for='input_address'></label>
                <input bind:this={input_address} type='text' id='input_address' />
              </div>
              <div class='map_container' bind:this={map_container}>
                <!-- Map goes here, rendered automatically -->
              </div>
            </div>
          </div>
          <button class='next_button' on:click|preventDefault={handlerLocationNext}>NEXT</button>
        </section>
        <section class='info' id='subpage_2' bind:this={cards[1]}>
          <div class='main_info'>
            <button class='back_button' on:click|preventDefault={houseInfoBack}>
              <i class='gg-chevron-left' />
            </button>
            {#if mode===ABOUT_THE_HOUSE}
              <h2>{mode}</h2>
              <div class='about_the_house'>
                <div class='house_type_container'>
                  <div class='house_type'>
                    <label for='house_type'>House Type</label>
                    <button on:click={handleHouseTypeOption}>{house_type ? house_type : 'select'}</button>
                    {#if showHouseOption}
                      <div class='option_container'>
                        <label class='option' for='house'>
                          <input
                            type='radio'
                            bind:group={house_type}
                            on:change={() => handleHouseTypeOption()}
                            id='house'
                            value='house'
                          />
                          House
                        </label>
                        <label class='option' for='apartment'>
                          <input
                            type='radio'
                            bind:group={house_type}
                            on:change={() => handleHouseTypeOption()}
                            id='apartment'
                            value='apartment'
                          />
                          Apartment
                        </label>
                      </div>
                    {/if}
                  </div>
                  {#if isApartment}
                    <div class='apartment_floor'>
                      <label for='apartment_floor'>Floor</label>
                      <input
                        type='number'
                        name='apartment_floor'
                        id='apartment_floor'
                        placeholder='Your property based'
                        min='1'
                        bind:value={floor}
                      />
                    </div>
                  {/if}
                </div>
                <div class='rent_or_sell'>
                  <label for='rent_or_sell'>Rent or Sell</label>
                  <button disabled={!house_type} on:click={handleRentSellOption}
                  >{rent_or_sell ? rent_or_sell : 'select'}</button
                  >
                  {#if showRentSellOption}
                    <div class='option_container'>
                      <label class='option' for='rent'>
                        <input
                          type='radio'
                          bind:group={rent_or_sell}
                          on:change={() => handleRentSellOption()}
                          id='rent'
                          value='rent'
                        />
                        Rent
                      </label>
                      
                      <label class='option' for='sell'>
                        <input
                          type='radio'
                          bind:group={rent_or_sell}
                          on:change={() => handleRentSellOption()}
                          id='sell'
                          value='sell'
                        />
                        Sell
                      </label>
                    </div>
                  {/if}
                </div>
              </div>
            {/if}
            {#if mode===UPLOAD_HOUSE_PHOTO}
              <h2>{mode}</h2>
              <div class='upload_house_photo'>
                <div class='image_preview_container'>
                  <label class='image_upload_button' for='upload_image'>
                    {#if !houseImgUploading}
                      <input
                        bind:this={imageHouseInput}
                        on:change={handlerHouseImage}
                        type='file'
                        accept='image/*'
                        id='upload_image'
                      />
                      <i class='gg-software-upload'></i>
                      <p>Select file to Upload</p>
                    {:else}
                      <progress value={uploadHousePercent} max='100'></progress>
                      <p>{uploadHouseStatus}</p>
                    {/if}
                  </label>
                  {#if house_images.length}
                    {#each house_images as imgFile, index}
                      <div class='image_card'>
                        <img src={imgFile} alt=''>
                        <button on:click={() => removeImage(index)} title='remove this image'>
                          <i class='gg-close'></i>
                        </button>
                      </div>
                    {/each}
                  {/if}
                </div>
              </div>
            {/if}
            {#if mode===FEATURE_OR_FACILITY}
              <h2>Feature</h2>
              <div class='feature_section'>
                <div class='beds'>
                  <input bind:value={feature.beds} type='number' min='0' name='beds' id='beds'>
                  <label for='beds'>Beds</label>
                </div>
                <div class='baths'>
                  <input bind:value={feature.baths} type='number' min='0' name='baths' id='baths'>
                  <label for='baths'>Baths</label>
                </div>
                <div class='square_foot'>
                  <input bind:value={feature.area} type='number' min='0' name='square_foot' id='square_foot' step='0.01'>
                  <label for='square_foot'>Sq Ft</label>
                </div>
              </div>
              <h2>Facility</h2>
              <div class='facility_section'>
                <textarea bind:value={facility} rows='10' placeholder='Type the facility in the property.' name='facility' id='facility'></textarea>
              </div>
            {/if}
            {#if mode===DESCRIPTION_PROPERTY}
              <h2>{mode}</h2>
              <div class='description_of_property'>
                  <textarea bind:value={description_of_property} rows='20'
                            placeholder='Writing a description can help potential buyers become more interested in your property.' name='description'
                            id='description'></textarea>
              </div>
            {/if}
            {#if mode===PRICE}
              <h2>{mode}</h2>
              <div class='price'>
                <div class='property_price'>
                  <label for='property_price'>Property Price</label>
                  <input bind:value={property_price} type='number' min='0' name='property_price' id='property_price' />
                </div>
                <div class='agency_fee'>
                  <label for='agency_fee'>Agency Fee</label>
                  <div>
                    <input bind:value={agency_fee} type='number' min='0' max='100' name='agency_fee' id='agency_fee'>
                    <span>%</span>
                  </div>
                </div>
              </div>
            {/if}
          </div>
          <div class='next_skip_button'>
            {#if modeSkippable===true}
              <button class='skip_button' on:click|preventDefault={houseInfoSkip}>
                SKIP
              </button>
            {/if}
            <button
              disabled={houseImgUploading === true}
              class='next_button'
              on:click|preventDefault={() => {
                  if (modeHouseInfoCount === 0) {
                    handleNextAboutHouse();
                  } else if (modeHouseInfoCount === 1) {
                    handleNextUploadHouseImage();
                  } else if (modeHouseInfoCount === 2) {
                    handleNextFeatureFacility();
                  } else if (modeHouseInfoCount === 3) {
                    handleNextDescriptionProperty();
                  } else if (modeHouseInfoCount === 4) {
                    handleNextPriceProperty();
                  }
                }}
            >
              {#if houseImgUploading===true}
                <i class='gg-disc'></i>
              {:else}
                <p>NEXT</p>
              {/if}
            </button>
          </div>
        </section>
        <section class='floor' id='subpage_3' bind:this={cards[2]}>
          <!-- Add Floor Dialog -->
          <AddFloorDialog bind:this={add_floor_dialog} bind:floor_type={floor_type}>
            <button disabled={floor_type === ''} class='add_floor_button' on:click={handlerAddFloor}>Add</button>
          </AddFloorDialog>
          <!-- Add Room Dialog -->
          <AddRoomDialog bind:this={add_room_dialog} bind:room_type={room_type} bind:room_size={room_size} bind:m2_size={size_m2} bind:unit_mode={unit_mode}>
            <button disabled={room_type === ''} class='add_room_button' on:click={() => handlerAddRoom(floor_index_to_edit)}>Add</button>
          </AddRoomDialog>
          <div class='floor_main_content'>
            <button
              class='back_button'
              on:click|preventDefault={() => {
                  if (floor_edit_mode === true) {
                    floor_edit_mode = false;
                  } else {
                    backPage();
                  }
                }}>
              <i class='gg-chevron-left' />
            </button>
            <div class='floor_header'>
              {#if floor_edit_mode===true}
                <h3>
                  Edit {floor_lists[ floor_index_to_edit ].type} {floor_lists[ floor_index_to_edit ].type==='basement' ? '' : `#${floor_lists[ floor_index_to_edit ].floor}` }</h3>
              {:else}
                <h2>Floors</h2>
                <button on:click|preventDefault={showAddFloorDialog}>Add</button>
              {/if}
            </div>
            {#if floor_edit_mode===false}
              <div class='floor_items_container'>
                {#if floor_lists.length}
                  {#each floor_lists as floor, index}
                    <div class='floor_item'>
                      <div class='left_item'>
                        <h4>{floor.type==='basement' ? floor.type : `${floor.type} #${floor.floor}`}</h4>
                        <div class='rooms_total'>
                          <div class='beds'>
                            <b>{floor.beds}</b>
                            <p>Beds</p>
                          </div>
                          <div class='baths'>
                            <b>{floor.baths}</b>
                            <p>Baths</p>
                          </div>
                        </div>
                      </div>
                      <div class='right_item'>
                        <button class='edit_floor' on:click={() => handlerEditFloor(index)}>Edit</button>
                        <div class='floor_plan'>
                          {#if floor.planImageUrl===''}
                              <span>
                                <i class='gg-image'></i>
                              </span>
                          {:else}
                            <img src={floor.planImageUrl} alt='' />
                          {/if}
                        </div>
                      </div>
                    </div>
                  {/each}
                {:else}
                  <div class='no_content'>
                    <p>No Floor</p>
                  </div>
                {/if}
              </div>
            {/if}
            {#if floor_edit_mode===true}
              <div class='edit_floor_container'>
                {#if imgFlrPlanUploading===true}
                  <div class='uploading'>
                    <progress value={uploadFlrPlanPercent} max='100'></progress>
                    <p>{uploadFlrPlanStatus}</p>
                  </div>
                {:else}
                  {#if floor_lists[ floor_index_to_edit ].planImageUrl===''}
                    <label class='floor_plan_upload' for='floor_plan_upload'>
                      <input
                        bind:this={imageFloorPlanInput}
                        on:change={handlerImageFloorPlan}
                        type='file'
                        accept='image/*'
                        id='floor_plan_upload'
                      />
                      <img src='/assets/img/realtor/floor-plan-pen-ruler.jpg.webp' alt=''>
                      <div>
                        <i class='gg-add'></i>
                        <p>Floor Plan Picture</p>
                      </div>
                    </label>
                  {:else}
                    <div class='floor_plan_preview'>
                      <img src={floor_lists[floor_index_to_edit].planImageUrl} alt=''>
                    </div>
                  {/if}
                {/if}
                <div class='room_list_container'>
                  <div class='room_list_header'>
                    <h3>Rooms</h3>
                    <button on:click={showAddRoomDialog}>Add</button>
                  </div>
                  {#if floor_lists[ floor_index_to_edit ][ 'rooms' ].length}
                    {#each floor_lists[ floor_index_to_edit ].rooms as room}
                      <div class='room_list_item'>
                        <span>{room.name}</span>
                        <div class='right_item'>
                          <span>{room.sizeM2} {room.unit}</span>
                          <button class='remove_room'>
                            <i class='gg-trash'></i>
                          </button>
                        </div>
                      </div>
                    {/each}
                  {:else}
                    <div class='no_room'>
                      <p>No room</p>
                    </div>
                  {/if}
                </div>
              </div>
            {/if}
          </div>
          <button disabled={imgFlrPlanUploading === true} class='next_button' on:click|preventDefault={handleNextFloor}>
            {#if imgFlrPlanUploading===true}
              <i class='gg-disc'></i>
            {:else}
              <p>NEXT</p>
            {/if}
          </button>
        </section>
        <section class='preview' id='subpage_4' bind:this={cards[3]}>
          {#if isPropertySubmitted===false}
            <div class='preview_main'>
              <h2>Preview Your Property</h2>
              <div class='image_preview_wrapper'>
                {#if infoStack.attrs.images.length}
                  <img src={infoStack.attrs.images[0]} alt=''>
                {:else}
                  <div class='image_preview_empty'>
                    <i class='gg-image'></i>
                    <p>No Image to Preview</p>
                  </div>
                {/if}
              </div>
              <div class='preview_price_house_type'>
                <div class='left_item'>
                  <span>On Sale</span>
                  <div class='price'>
                    <h3>{formatPrice( infoStack.attrs.price )}</h3>
                    <p>Agency Fee: {infoStack.attrs.agency_fee_percent}%</p>
                  </div>
                </div>
                <div class='right_item'>
                  <button class='like_button'>
                    <i class='gg-heart'></i>
                  </button>
                  <div class='house_type'>
                    <i class='gg-home-alt'></i>
                    <span>{infoStack.attrs.house_type}</span>
                  </div>
                </div>
              </div>
              <div class='preview_feature'>
                <div>
                  <b>{infoStack.attrs.feature.beds}</b>
                  <p>Beds</p>
                </div>
                <div>
                  <b>{infoStack.attrs.feature.baths}</b>
                  <p>Baths</p>
                </div>
                <div>
                  <b>{infoStack.attrs.feature.area}</b>
                  <p>Sq Ft</p>
                </div>
              </div>
              <article class='preview_description'>
                <div class='preview_facility'>
                  <h3>Facility</h3>
                  <p>{infoStack.attrs.facility!=='' ? infoStack.attrs.facility : '--'}</p>
                </div>
                <div class='preview_about'>
                  <h3>About</h3>
                  <p>{infoStack.attrs.description!=='' ? infoStack.attrs.description : '--'}</p>
                </div>
              </article>
              <div class='preview_floors'>
                <h2>Floors</h2>
                <div class='floor_container'>
                  {#each realtorStack[ 2 ].attrs.floor_lists as floors}
                    <div class='floor_item'>
                      <h3>{floors.type==='basement' ? floors.type : `${floors.type} #${floors.floor}`}</h3>
                      <div class='floor_attr'>
                        <div class='floor_rooms'>
                          {#if floors[ 'rooms' ].length}
                            {#each floors.rooms as rooms}
                              <div class='room_item'>
                                <span>{rooms.name}</span>
                                <span>{rooms.sizeM2} {rooms.unit}</span>
                              </div>
                            {/each}
                          {/if}
                        </div>
                        <div class='floor_plan'>
                          {#if floors.planImageUrl===''}
                              <span>
                                <i class='gg-image'></i>
                              </span>
                          {:else}
                            <img src={floors.planImageUrl} alt='' />
                          {/if}
                        </div>
                      </div>
                    </div>
                  {/each}
                </div>
              </div>
            </div>
            <button class='submit_button' on:click|preventDefault={handleSubmit}>SUBMIT</button>
          {/if}
          {#if isPropertySubmitted===true}
            <div class='property_submitted'>
              <div class='icon_submitted'>
                <i class='gg-check-r'></i>
                <span>SUBMITTED</span>
              </div>
              <div class='message_submmitted'>
                <b>We will review it soon</b>
                <p>Thanks you for submitting your property.</p>
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
  @keyframes spin {
    from {
      transform : rotate(0deg);
    }
    to {
      transform : rotate(360deg);
    }
  }

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

  .realtor_step_progress_bar .step_wrapper {
    width            : 70%;
    margin-left      : auto;
    margin-right     : auto;
    display          : flex;
    flex-direction   : row;
    justify-content  : space-between;
    background-color : white;
    border-radius    : 6px;
    filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    padding          : 19px 10% 11px 10%;
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

  .realtor_step_progress_bar .step_wrapper .step_item span {
    width            : 13px;
    height           : 13px;
    background-color : #CBD5E1;
    border-radius    : 100%;
    z-index          : 4;
  }

  .realtor_step_progress_bar .step_wrapper .step_item p {
    margin : 8px 0 0 0;
  }

  .realtor_step_progress_bar .step_wrapper .step_item.active span {
    width            : 11px;
    height           : 11px;
    background-color : white;
    outline          : 3px solid #F97316;
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

  .step_item.completed span {
    background-color : #F97316 !important;
  }

  /* Main Container Subpage */
  .realtor_subpage_container {
    position         : relative;
    margin-top       : -40px;
    margin-left      : auto;
    margin-right     : auto;
    display          : flex;
    border-radius    : 8px;
    width            : 860px;
    min-height       : 500px;
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
  }

  .realtor_subpage_container::-webkit-scrollbar-track {
    background-color : transparent;
  }

  .realtor_subpage_container section {
    padding           : 20px;
    margin            : 0 10px;
    background-color  : white;
    border-radius     : 8px;
    min-height        : 700px;
    flex              : 0 0 860px;
    scroll-snap-align : start;
  }

  .realtor_subpage_container .back_button {
    padding          : 5px;
    border           : none;
    background-color : rgb(0 0 0 / 0.06);
    border-radius    : 5px;
    font-size        : 14px;
    cursor           : pointer;
  }

  .realtor_subpage_container .back_button:hover {
    background-color : rgb(0 0 0 / 0.05);
    color            : #EF4444;
  }

  /* +============| SUBPAGE LOCATION |===========+ */
  .location {
    color           : #334155;
    display         : flex;
    flex-direction  : column;
    justify-content : space-between;
    height          : 100%;
  }

  .location .next_button {
    background-color : #F97316;
    border-radius    : 8px;
    border           : none;
    padding          : 10px;
    color            : white;
    font-weight      : 600;
    margin-top       : 20px;
    cursor           : pointer;
  }

  .location .next_button:hover {
    background-color : #F58433;
  }

  .location h2 {
    font-size     : 18px;
    font-weight   : 600;
    margin-bottom : 20px;
  }

  .location_input {
    display        : flex;
    flex-direction : column;
  }

  .location_input #input_address {
    margin-top    : 10px;
    padding       : 12px 15px;
    border-radius : 2px;
    border        : none;
    width         : 400px;
    filter        : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
  }

  .location_input #input_address:focus {
    border-color : #F97316;
    outline      : 2px solid #F97316;
  }

  .location_input .map_container {
    margin-top    : 20px;
    border-radius : 8px;
    width         : 100%;
    height        : 430px;
  }

  /* +============| SUBPAGE INFO |===========+ */
  .info {
    color           : #334155;
    display         : flex;
    flex-direction  : column;
    justify-content : space-between;
    height          : 100%;
  }

  .info .next_skip_button {
    display        : flex;
    flex-direction : row;
    align-items    : stretch;
    font-weight    : 500;
    margin-top     : 20px;
    width          : 100%;
  }

  .info .next_skip_button button {
    border-radius : 8px;
    border        : none;
    padding       : 10px;
    cursor        : pointer;
    width         : 100%;
  }

  .info .next_skip_button button p {
    margin : 0;
  }

  .info .next_skip_button .skip_button {
    margin-right     : 10px;
    border           : 1px solid #CBD5E1;
    background-color : #F1F5F9;
    color            : #F97316;
  }

  .info .next_skip_button .skip_button:hover {
    border : 1px solid #F97316;
  }

  .info .next_skip_button .next_button {
    background-color : #F97316;
    color            : white;
    display          : flex;
    justify-content  : center;
  }

  .info .next_skip_button .next_button:hover {
    background-color : #F58433;
  }

  .info .next_skip_button .next_button i {
    animation : spin 1.5s infinite linear;
  }

  .info h2 {
    font-size     : 18px;
    font-weight   : 600;
    margin-bottom : 20px;
  }

  /* __SUBPAGE INFO - About The House */
  .about_the_house {
    display            : grid;
    grid-template-rows : 1fr 1fr;
    gap                : 20px;
  }

  .house_type_container {
    display               : grid;
    grid-template-columns : 1fr 1fr;
    gap                   : 20px;
  }

  .house_type, .apartment_floor, .rent_or_sell {
    position       : relative;
    display        : flex;
    flex-direction : column;
    width          : 100%;
    height         : fit-content;
  }

  .house_type button, .apartment_floor input, .rent_or_sell button {
    width            : auto;
    border           : 1px solid #CBD5E1;
    background-color : #F1F5F9;
    border-radius    : 8px;
    padding          : 10px 12px;
    font-weight      : 600;
    text-align       : left;
    color            : #F97316;
    text-transform   : capitalize;
  }

  .rent_or_sell button {
    width : 48.3% !important;
  }

  .house_type button:hover, .apartment_floor input:hover, .rent_or_sell button:hover {
    border : 1px solid #F97316;
  }

  .apartment_floor input:focus {
    outline : 1px solid #F97316;
  }

  .rent_or_sell .option_container {
    width : 50% !important;
  }

  .about_the_house .option_container {
    width            : 100%;
    position         : absolute;
    top              : 70px;
    z-index          : 3;
    display          : flex;
    flex-direction   : column;
    border-radius    : 8px;
    border           : 1px solid #CBD5E1;
    background-color : #F1F5F9;
  }

  .about_the_house .option_container .option {
    margin      : 0;
    padding     : 10px 12px;
    font-weight : 500;
    cursor      : pointer;
  }

  .about_the_house .option_container .option:hover {
    color : #F97316;
  }

  .option input[type='radio'] {
    position       : absolute;
    opacity        : 0;
    pointer-events : none;
  }

  .house_type label, .apartment_floor label, .rent_or_sell label {
    font-size     : 13px;
    font-weight   : 700;
    margin-left   : 10px;
    margin-bottom : 8px;
  }

  /* __SUBPAGE INFO - Upload House Photo*/
  .upload_house_photo {
    display : grid;
    gap     : 20px;
  }

  .upload_house_photo .image_preview_container {
    display               : grid;
    justify-items         : center;
    grid-template-columns : repeat(3, 1fr);
    gap                   : 20px;
  }

  .upload_house_photo .image_card {
    position      : relative;
    border-radius : 8px;
    width         : 100%;
    height        : 110px;
  }

  .upload_house_photo .image_card img {
    border-radius : 8px;
    object-fit    : cover;
    width         : 100%;
    height        : 100%;
    border        : 1px solid #CBD5E1;
  }

  .upload_house_photo .image_card button {
    position         : absolute;
    z-index          : 10;
    top              : -10px;
    right            : -10px;
    background-color : #EF4444;
    border           : none;
    color            : white;
    padding          : 5px;
    border-radius    : 50%;
    cursor           : pointer;
    filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
  }

  .upload_house_photo .image_upload_button {
    display          : flex;
    flex-direction   : column;
    justify-content  : center;
    align-items      : center;
    border           : 1px solid #CBD5E1;
    background-color : #F1F5F9;
    border-radius    : 8px;
    width            : 100%;
    height           : 110px;
    cursor           : pointer;
  }

  .upload_house_photo .image_upload_button:hover {
    border : 1px solid #F97316;
    color  : #F97316;
  }

  .upload_house_photo .image_upload_button input {
    position       : absolute;
    opacity        : 0;
    pointer-events : none;
  }

  .upload_house_photo .image_upload_button i {
    font-size     : 60px;
    margin-bottom : 10px;
  }

  .upload_house_photo .image_upload_button p {
    font-size : 11px;
    margin    : 0;
  }

  .upload_house_photo .image_upload_button progress {
    appearance    : none;
    border-radius : 8px;
    height        : 13px;
    overflow      : hidden;
    margin-bottom : 8px;
  }

  .upload_house_photo .image_upload_button progress::-webkit-progress-bar {
    background-color   : aliceblue;
    box-shadow         : -1px 1px 10px 0px rgba(0, 0, 0, 0.3) inset;
    -webkit-box-shadow : -1px 1px 10px 0px rgba(0, 0, 0, 0.3) inset;
    -moz-box-shadow    : -1px 1px 10px 0px rgba(0, 0, 0, 0.3) inset;
  }

  .upload_house_photo .image_upload_button progress::-webkit-progress-value {
    background-color : #F97316;
  }

  .upload_house_photo .image_upload_button progress::-moz-progress-bar {
    background-color : #F97316;
  }

  /* __SUBPAGE INFO - Feature and Facility */
  .feature_section {
    display               : grid;
    grid-template-columns : 1fr 1fr 1fr;
    gap                   : 20px;
  }

  .feature_section .beds, .baths, .square_foot {
    display        : flex;
    flex-direction : column;
  }

  .feature_section .beds input, .baths input, .square_foot input {
    width            : 100%;
    border           : 1px solid #CBD5E1;
    background-color : #F1F5F9;
    border-radius    : 8px;
    padding          : 10px 12px;
    text-align       : center;
    margin-bottom    : 10px;
  }

  .facility_section textarea {
    resize           : vertical;
    width            : 100%;
    border           : 1px solid #CBD5E1;
    background-color : #F1F5F9;
    border-radius    : 8px;
    padding          : 10px 12px;
    text-align       : left;
  }

  /* __SUBPAGE INFO - Description of Property */
  .description_of_property textarea {
    resize           : vertical;
    width            : 100%;
    border           : 1px solid #CBD5E1;
    background-color : #F1F5F9;
    border-radius    : 8px;
    padding          : 10px 12px;
    text-align       : left;
  }

  /* __SUBPAGE INFO - Price */
  .price {
    display            : grid;
    grid-template-rows : 1fr 1fr;
    gap                : 20px;
  }

  .price .property_price, .agency_fee {
    display        : flex;
    flex-direction : column;
  }

  .price label {
    font-size     : 13px;
    font-weight   : 700;
    margin-left   : 10px;
    margin-bottom : 8px;
  }

  .price .property_price input, .agency_fee input {
    width            : 50%;
    border           : 1px solid #CBD5E1;
    background-color : #F1F5F9;
    border-radius    : 8px;
    padding          : 10px 12px;
    text-align       : left;
  }

  .price .agency_fee_percent div {
    display        : flex;
    flex-direction : row;
    align-items    : center;
  }

  .price .agency_fee_percent div span {
    margin-left : 8px;
    font-size   : 16px;
    font-weight : 600;
  }

  /* +============| SUBPAGE FLOORS |===========+ */
  .floor {
    color           : #334155;
    display         : flex;
    flex-direction  : column;
    justify-content : space-between;
    height          : 100%;
    position        : relative;
  }

  .floor .next_button {
    background-color : #F97316;
    border-radius    : 8px;
    border           : none;
    padding          : 10px;
    color            : white;
    font-weight      : 600;
    margin-top       : 20px;
    cursor           : pointer;
    display          : flex;
    justify-content  : center;
  }

  .floor .next_button:hover {
    background-color : #F58433;
  }

  .floor .next_button i {
    animation : spin 1.5s infinite linear;
  }

  .floor .next_button p {
    margin : 0;
  }

  .floor .floor_header {
    margin          : 20px 0;
    display         : flex;
    flex-direction  : row;
    justify-content : space-between;
    align-items     : center;
    width           : 100%;
  }

  .floor .floor_header h2 {
    font-weight : 600;
    margin      : 0;
    font-size   : 18px;
  }

  .floor .floor_header h3 {
    font-weight : 600;
    margin      : 0;
    font-size   : 18px;
    width       : 100%;
    text-align  : center;
  }

  .floor .floor_header button {
    border-radius : 8px;
    border        : none;
    padding       : 10px 12px;
    color         : #F97316;
    font-weight   : 600;
    cursor        : pointer;
    background    : none;
    font-size     : 14px;
  }

  .floor .floor_header button:hover {
    background-color : rgb(0 0 0 / 0.07);
  }

  .floor .add_floor_button, .add_room_button {
    background-color : #F97316;
    color            : white;
    border-radius    : 8px;
    border           : none;
    padding          : 10px;
    cursor           : pointer;
    width            : 100%;
  }

  .floor .add_floor_button:hover, .add_room_button:hover {
    background-color : #F58433;
  }

  .floor .floor_items_container {
    display         : flex;
    flex-direction  : column;
    justify-content : center;
    width           : 100%;
  }

  .floor .floor_items_container .no_content {
    display          : flex;
    justify-content  : center;
    border-radius    : 8px;
    background-color : rgb(0 0 0 / 0.06);
    padding          : 80px 20px;
    height           : 100%;
    font-weight      : 600;
    font-size        : 16px;
    width            : 60%;
    margin           : 0 auto;
  }

  .floor .floor_items_container .floor_item {
    display          : flex;
    flex-direction   : row;
    justify-content  : space-between;
    padding          : 20px;
    background-color : rgba(0 0 0 / 0.06);
    border-radius    : 8px;
    height           : fit-content;
    width            : 60%;
    margin           : 0 auto 20px auto;
  }

  .floor .floor_items_container .floor_item .left_item {
    flex-basis      : 40%;
    display         : flex;
    flex-direction  : column;
    justify-content : space-between;
  }

  .floor .floor_items_container .floor_item .left_item h4 {
    font-size      : 16px;
    font-weight    : 600;
    text-transform : capitalize;
    margin         : 0;
  }

  .floor .floor_items_container .floor_item .left_item .rooms_total {
    display               : grid;
    grid-template-columns : 1fr 1fr;
    gap                   : 20px;
    margin                : 20px 0;
  }

  .floor .floor_items_container .floor_item .beds, .floor_item .baths {
    display        : flex;
    flex-direction : column;
    text-align     : center;
  }

  .floor .floor_items_container .floor_item .beds b, .floor_item .baths b {
    font-size : 25px;
  }

  .floor .floor_items_container .floor_item .beds p, .floor_item .baths p {
    margin : 10px 0 0 0;
  }

  .floor .floor_items_container .floor_item .right_item {
    flex-basis      : 30%;
    display         : flex;
    flex-direction  : column;
    justify-content : space-between;
  }

  .floor .floor_items_container .floor_item .right_item .edit_floor {
    background-color : #F97316;
    border-radius    : 8px;
    border           : none;
    padding          : 7px;
    color            : white;
    font-weight      : 600;
    cursor           : pointer;
    width            : 100%;
    margin           : 0 0 0 auto;
  }

  .floor .floor_items_container .floor_item .right_item .edit_floor:hover {
    background-color : #F58433;
  }

  .floor .floor_items_container .floor_item .right_item .floor_plan {
    position      : relative;
    border-radius : 8px;
    width         : 100%;
    height        : 90px;
    border        : 1px solid #CBD5E1;
    margin-top    : 20px;
    overflow      : hidden;
  }

  .floor .floor_items_container .floor_item .right_item .floor_plan img {
    width      : 100%;
    height     : 100%;
    object-fit : cover;
  }

  .floor .floor_items_container .floor_item .right_item .floor_plan span {
    border-radius    : 8px;
    object-fit       : cover;
    width            : 100%;
    height           : 100%;
    background-color : rgb(0 0 0 / 0.06);
    display          : flex;
    justify-content  : center;
    align-items      : center;
  }

  .floor .edit_floor_container {
    height         : fit-content;
    width          : 60%;
    margin         : 0 auto 20px auto;
    display        : flex;
    flex-direction : column;
  }

  .floor .edit_floor_container .uploading {
    border          : 1px solid #CBD5E1;
    border-radius   : 8px;
    width           : 100%;
    height          : 130px;
    display         : flex;
    flex-direction  : column;
    justify-content : center;
    align-items     : center;
  }

  .floor .edit_floor_container .uploading p {
    font-size : 11px;
    margin    : 0;
  }

  .floor .edit_floor_container .uploading progress {
    appearance    : none;
    border-radius : 8px;
    height        : 13px;
    overflow      : hidden;
    margin-bottom : 8px;
    width         : 65%;
  }

  .floor .edit_floor_container .uploading progress::-webkit-progress-bar {
    background-color   : aliceblue;
    box-shadow         : -1px 1px 10px 0px rgba(0, 0, 0, 0.3) inset;
    -webkit-box-shadow : -1px 1px 10px 0px rgba(0, 0, 0, 0.3) inset;
    -moz-box-shadow    : -1px 1px 10px 0px rgba(0, 0, 0, 0.3) inset;
  }

  .floor .edit_floor_container .uploading progress::-webkit-progress-value {
    background-color : #F97316;
  }

  .floor .edit_floor_container .uploading progress::-moz-progress-bar {
    background-color : #F97316;
  }

  .floor .edit_floor_container .floor_plan_upload {
    border        : 1px solid #CBD5E1;
    border-radius : 8px;
    width         : 100%;
    height        : 130px;
    cursor        : pointer;
    overflow      : hidden;
    position      : relative;
  }

  .floor .edit_floor_container .floor_plan_upload:hover {
    border : 1px solid #F97316;
  }

  .floor .edit_floor_container .floor_plan_upload input {
    position       : absolute;
    opacity        : 0;
    pointer-events : none;
  }

  .floor .edit_floor_container .floor_plan_upload img {
    width  : 100%;
    height : auto;
  }

  .floor .edit_floor_container .floor_plan_upload div {
    background-color : rgb(255 255 255 / 0.7);
    position         : absolute;
    top              : 0;
    right            : 0;
    bottom           : 0;
    left             : 0;
    display          : flex;
    flex-direction   : row;
    align-items      : center;
    padding          : 0 20px;
    font-size        : 20px;
  }

  .floor .edit_floor_container .floor_plan_upload div p {
    margin-left : 25px;
  }

  .floor .edit_floor_container .floor_plan_preview {
    border        : 1px solid #CBD5E1;
    border-radius : 8px;
    width         : 100%;
    height        : 130px;
    overflow      : hidden;
  }

  .floor .edit_floor_container .floor_plan_preview img {
    width  : 100%;
    height : auto;
    cursor : pointer;
  }

  .floor .edit_floor_container .room_list_container .room_list_header {
    margin          : 20px 0;
    width           : 100%;
    display         : flex;
    flex-direction  : row;
    justify-content : space-between;
    align-items     : center;
  }

  .floor .edit_floor_container .room_list_container .room_list_header h3 {
    font-weight : 600;
    margin      : 0;
    font-size   : 18px;
  }

  .floor .edit_floor_container .room_list_container .room_list_header button {
    border-radius : 8px;
    border        : none;
    padding       : 10px 12px;
    color         : #F97316;
    font-weight   : 600;
    cursor        : pointer;
    background    : none;
    font-size     : 14px;
  }

  .floor .edit_floor_container .room_list_container .room_list_header button:hover {
    background-color : rgb(0 0 0 / 0.07);
  }

  .floor .edit_floor_container .room_list_container .room_list_item {
    width           : 100%;
    display         : flex;
    flex-direction  : row;
    align-items     : center;
    justify-content : space-between;
    padding         : 8px 0;
    border-bottom   : 1px solid #334155;
    font-weight     : 500;
    text-transform  : capitalize;
  }

  .floor .edit_floor_container .room_list_container .room_list_item .right_item {
    display        : flex;
    flex-direction : row;
    align-items    : center;
  }

  .floor .edit_floor_container .room_list_container .room_list_item .right_item .remove_room {
    border        : none;
    background    : none;
    padding       : 10px 13px;
    border-radius : 50%;
    margin-left   : 8px;
    cursor        : pointer;
    color         : #F97316;
  }

  .floor .edit_floor_container .room_list_container .room_list_item .right_item .remove_room:hover {
    background-color : rgb(0 0 0 / 0.06);
  }

  .floor .room_list_container .no_room {
    display          : flex;
    justify-content  : center;
    border-radius    : 8px;
    background-color : rgb(0 0 0 / 0.06);
    padding          : 80px 20px;
    height           : fit-content;
    font-weight      : 600;
    font-size        : 16px;
    width            : 100%;
  }

  /* +============| SUBPAGE PREVIEW |===========+ */
  .preview {
    color           : #334155;
    display         : flex;
    flex-direction  : column;
    justify-content : space-between;
    min-height      : 100%;
    width           : 100%;
  }

  .preview .preview_main {
    display        : flex;
    flex-direction : column;
  }

  .preview h2 {
    font-weight : 600;
    margin      : 0;
    font-size   : 18px;
  }

  .preview .image_preview_wrapper {
    border-radius : 8px;
    width         : 60%;
    height        : 200px;
    border        : 1px solid #CBD5E1;
    margin        : 20px auto;
    overflow      : hidden;
  }

  .preview .image_preview_wrapper img {
    width      : 100%;
    height     : 100%;
    object-fit : cover;
  }

  .preview .image_preview_wrapper .image_preview_empty {
    border-radius    : 8px;
    object-fit       : cover;
    width            : 100%;
    height           : 100%;
    background-color : rgb(0 0 0 / 0.06);
    display          : flex;
    justify-content  : center;
    align-items      : center;
  }

  .preview .image_preview_wrapper .image_preview_empty p {
    margin-left : 10px;
  }

  .preview .preview_price_house_type {
    display         : flex;
    flex-direction  : row;
    justify-content : space-between;
    width           : 80%;
    margin          : 0 auto;
  }

  .preview .preview_price_house_type .left_item {
    display        : flex;
    flex-direction : column;
  }

  .preview .preview_price_house_type .left_item span {
    padding          : 4px 10px;
    font-size        : 12;
    background-color : #F97316;
    width            : fit-content;
    color            : white;
    margin-bottom    : 10px;
    text-transform   : capitalize;
  }

  .preview .preview_price_house_type .left_item .price {
    display : block;
  }

  .preview .preview_price_house_type .left_item .price h3 {
    font-size   : 30px;
    font-weight : 700;
    margin      : 0;
  }

  .preview .preview_price_house_type .left_item p {
    color     : #64748B;
    font-size : 13px;
    margin    : 10px 0 0 0;
  }

  .preview .preview_price_house_type .right_item {
    display         : flex;
    flex-direction  : column;
    justify-content : flex-start;
    align-items     : flex-end;
  }

  .preview .preview_price_house_type .right_item .like_button {
    border        : none;
    background    : none;
    padding       : 10px;
    border-radius : 50%;
    margin-bottom : 10px;
    cursor        : pointer;
  }

  .preview .preview_price_house_type .right_item .like_button:hover {
    background-color : rgb(0 0 0 / 0.06);
    color            : #F97316;
  }

  .preview .preview_price_house_type .right_item .house_type {
    border-radius    : 999px;
    background-color : #F97316;
    color            : white;
    padding          : 10px 25px;
    width            : fit-content;
    display          : flex;
    flex-direction   : row;
    align-items      : center;
  }

  .preview .preview_price_house_type .right_item .house_type span {
    margin-left    : 5px;
    font-size      : 18px;
    text-transform : capitalize;
  }

  .preview .preview_feature {
    display               : grid;
    grid-template-columns : 1fr 1fr 1fr;
    align-items           : center;
    margin                : 30px auto 0 auto;
    width                 : 60%;
  }

  .preview .preview_feature div {
    text-align : center;
  }

  .preview .preview_feature div b {
    font-size : 22px;
  }

  .preview .preview_feature div p {
    margin : 10px 0 0 0;
  }

  .preview .preview_description {
    display        : flex;
    flex-direction : column;
    margin         : 20px auto 0 auto;
    width          : 80%;
  }

  .preview .preview_description div:nth-child(1) {
    margin-bottom : 20px;
  }

  .preview .preview_description div h3 {
    font-weight : 600;
    margin      : 0 0 10px 0;
    font-size   : 18px;
  }

  .preview .preview_description div p {
    margin : 0;
  }

  .preview .preview_floors {
    display        : flex;
    flex-direction : column;
    margin         : 30px auto 0 auto;
    width          : 80%;
  }

  .preview .preview_floors h2 {
    font-size     : 22px;
    margin-bottom : 20px;
  }

  .preview .preview_floors .floor_container {
    display        : flex;
    flex-direction : column;
    width          : 100%;
  }

  .preview .preview_floors .floor_container .floor_item {
    display        : flex;
    flex-direction : column;
    margin-bottom  : 15px;
  }

  .preview .preview_floors .floor_container .floor_item h3 {
    font-weight    : 600;
    margin         : 0 0 8px 0;
    font-size      : 18px;
    text-transform : capitalize;
  }

  .preview .preview_floors .floor_container .floor_item .floor_attr {
    display         : flex;
    flex-direction  : row;
    justify-content : space-between;
  }

  .preview .preview_floors .floor_container .floor_item .floor_attr .floor_rooms {
    flex-basis     : 60%;
    display        : flex;
    flex-direction : column;
  }

  .preview .preview_floors .floor_container .floor_item .floor_attr .floor_rooms .room_item {
    display         : flex;
    flex-direction  : row;
    justify-content : space-between;
    align-items     : center;
    font-weight     : 500;
    text-transform  : capitalize;
    padding         : 5px 0;
  }

  .preview .preview_floors .floor_container .floor_item .floor_attr .floor_plan {
    position      : relative;
    border-radius : 8px;
    width         : 200px;
    height        : 110px;
    border        : 1px solid #CBD5E1;
    overflow      : hidden;
  }

  .preview .preview_floors .floor_container .floor_item .floor_attr .floor_plan img {
    width      : 100%;
    height     : 100%;
    object-fit : cover;
  }

  .preview .preview_floors .floor_container .floor_item .floor_attr .floor_plan span {
    border-radius    : 8px;
    object-fit       : cover;
    width            : 100%;
    height           : 100%;
    background-color : rgb(0 0 0 / 0.06);
    display          : flex;
    justify-content  : center;
    align-items      : center;
  }

  .preview .submit_button {
    background-color : #F97316;
    border-radius    : 8px;
    border           : none;
    padding          : 10px;
    color            : white;
    font-weight      : 600;
    margin-top       : 20px;
    cursor           : pointer;
  }

  .preview .submit_button:hover {
    background-color : #F58433;
  }

  .preview .property_submitted {
    height          : 100%;
    width           : 100%;
    display         : flex;
    flex-direction  : column;
    justify-content : center;
    align-items     : center;
    text-align      : center;
  }

  .preview .property_submitted .icon_submitted {
    background-color : rgb(0 0 0 / 0.06);
    padding          : 30px 40px;
    border-radius    : 8px;
    display          : flex;
    flex-direction   : row;
    margin-bottom    : 20px;
  }

  .preview .property_submitted .icon_submitted span {
    margin-left : 10px;
    font-size   : 20px;
    font-weight : 700;
  }

  .preview .property_submitted .message_submmitted b {
    font-weight : 700;
    font-size   : 20px;
  }
</style>