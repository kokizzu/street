<script>
  // @ts-nocheck
  import {UserNearbyFacilities, UserSearchProp, UserLikeProp} from '../jsApi.GEN.js';
  import {formatPrice} from './formatter.js';
  import {T} from './uiState.js';
  import {GoogleMap, GoogleSdk} from './GoogleMap/components';
  import Growl from './Growl.svelte';
  import {mapComponent} from './GoogleMap/stores';
  
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidSearch from 'svelte-icons-pack/fa/FaSolidSearch';
  import FaSolidMapMarkerAlt from 'svelte-icons-pack/fa/FaSolidMapMarkerAlt';
  import FaSolidImage from 'svelte-icons-pack/fa/FaSolidImage';
  import FaSolidHome from 'svelte-icons-pack/fa/FaSolidHome';
  import FaSolidRulerCombined from 'svelte-icons-pack/fa/FaSolidRulerCombined';
  import FaSolidBuilding from 'svelte-icons-pack/fa/FaSolidBuilding';
  import FaSolidBath from 'svelte-icons-pack/fa/FaSolidBath';
  import FaSolidBed from 'svelte-icons-pack/fa/FaSolidBed';
  import FaSolidUndoAlt from 'svelte-icons-pack/fa/FaSolidUndoAlt';
  import FaSolidBan from 'svelte-icons-pack/fa/FaSolidBan';
  import FaSolidReceipt from 'svelte-icons-pack/fa/FaSolidReceipt';
  import FaSolidShareAlt from "svelte-icons-pack/fa/FaSolidShareAlt";
  import FaBrandsLinkedin from "svelte-icons-pack/fa/FaBrandsLinkedin";
  import FaBrandsTwitter from "svelte-icons-pack/fa/FaBrandsTwitter";
  import FaCopy from "svelte-icons-pack/fa/FaCopy";
  import FaBrandsFacebook from "svelte-icons-pack/fa/FaBrandsFacebook";
  import FaBrandsTelegram from "svelte-icons-pack/fa/FaBrandsTelegram";
  import FaBrandsWhatsapp from "svelte-icons-pack/fa/FaBrandsWhatsapp";
  import FaSolidCircleNotch from "svelte-icons-pack/fa/FaSolidCircleNotch";
  import FaSolidAngleLeft from 'svelte-icons-pack/fa/FaSolidAngleLeft';
  import FaHeart from "svelte-icons-pack/fa/FaHeart";
  import {distanceKM} from './GoogleMap/distance';
  import {notifier} from './notifier.js';
  
  export let randomProps = [];
  export let defaultDistanceKm = 20;
  export let initialLatLong = [0, 0];
  let facilities = [], markersFacility = [], markersProperty = [], propItemBinds = [], infoWindows, propItemHighlight = null;
  let gmapsComponent = GoogleMap;
  let gmapBounds = {}; // top-left bottom-right of map
  let myLatLng = {lat: initialLatLong[ 0 ], lng: initialLatLong[ 1 ]};
  let mapOptions = {
    center: myLatLng,
    zoom: 11,
    mapTypeId: 'roadmap',
    mapId: 'street_project',
  };
  const markers_icon = {
    school: {path: '/assets/icons/marker-school.svg'},
    restaurant: {path: '/assets/icons/marker-restaurant.svg'},
    convenience_store: {path: '/assets/icons/marker-mall.svg'},
    hospital: {path: '/assets/icons/marker-hospital.svg'},
    subway_station: {path: '/assets/icons/marker-subway.svg'},
  };
  let geocoder, input_search_value, autocomplete_service;
  let autocomplete_lists = [];
  let shareItemIndex = null;
  let isSearchingMap = false;
  
  const highLightMapMarker = {
    enter: ( index ) => {
      propItemHighlight = index;
      if( !markersProperty || !markersProperty[ index ] ) return;
      const marker = markersProperty[ index ];
      marker.setZIndex( google.maps.Marker.MAX_ZINDEX + 1 )
      marker.setIcon( {
        url: '/assets/icons/marker-2.svg', // URL to your custom icon image
        scaledSize: new google.maps.Size( 40, 40 ),
      } );
    },
    leave: ( index ) => {
      if( !markersProperty || !markersProperty[ index ] ) return;
      const marker = markersProperty[ index ];
      marker.setZIndex( google.maps.Marker.MAX_ZINDEX - 1 )
      marker.setIcon( {
        url: '/assets/icons/marker-2.svg', // URL to your custom icon image
        scaledSize: new google.maps.Size( 32, 32 ),
      } );
      propItemHighlight = null;
    },
  };
  
  async function searchProperty( search ) {
    if( search ) {
      let bestDistance = defaultDistanceKm;
      if( gmapBounds && gmapBounds.getNorthEast ) {
        let ne = gmapBounds.getNorthEast();
        bestDistance = distanceKM( ne.lat(), ne.lng(), myLatLng.lat, myLatLng.lng );
      }
      await UserSearchProp( {
        centerLat: myLatLng.lat,
        centerLong: myLatLng.lng,
        offset: 0,
        limit: 40, // this is apparently the culprit XD if we show too many it would slow, but if it's too little it won't spread
        maxDistanceKM: bestDistance,
      }, async res => {
        if( res.error ) {
          notifier.showError(res.error );
          return;
        }
        randomProps = res.properties || [];
      } );
      markersProperty = gmapsComponent.clearMarkers( markersProperty );
    }
    randomProps.forEach( prop => {
      markersProperty.push( gmapsComponent.createMarker( prop.lat, prop.lng, '/assets/icons/marker-2.svg', 32, prop.uniqPropKey ) );
    } );
    markersProperty.forEach( ( marker, idx ) => {
      marker.addListener( 'mouseover', () => {
        marker.setZIndex( google.maps.Marker.MAX_ZINDEX + 1 )
        marker.setIcon( {
          url: '/assets/icons/marker-2.svg', // URL to your custom icon image
          scaledSize: new google.maps.Size( 40, 40 ),
        } );
      } );
      marker.addListener( 'mouseout', () => {
        marker.setZIndex( google.maps.Marker.MAX_ZINDEX - 1 )
        marker.setIcon( {
          url: '/assets/icons/marker-2.svg', // URL to your custom icon image
          scaledSize: new google.maps.Size( 32, 32 ),
        } );
      } );
      marker.addListener( 'click', () => {
        let propItem = propItemBinds[ idx ];
        propItemHighlight = idx;
        propItem.scrollIntoView( {behavior: 'smooth'} );
        setTimeout( () => {
          propItemHighlight = null;
        }, 2200 );
      } );
    } );
  }
  
  async function searchNearbyFacility() {
    await UserNearbyFacilities( {
      centerLat: myLatLng.lat,
      centerLong: myLatLng.lng,
    }, async res => {
      if( res.error )  {
        notifier.showError(res.error );
        return;
      }
      markersFacility = gmapsComponent.clearMarkers( markersFacility );
      facilities = await res.facilities;
      facilities.forEach( fac => {
        let iconmarkerpath = '/assets/icons/marker.svg';
        if( markers_icon[ fac.type ] ) {
          iconmarkerpath = markers_icon[ fac.type ].path;
        }
        markersFacility.push( gmapsComponent.createMarker( fac.lat, fac.lng, iconmarkerpath, 32, fac.name ) );
      } );
    } );
    markersFacility.forEach( ( marker, idx ) => {
      marker.addListener( 'click', () => {
        if( infoWindows ) {
          infoWindows.close();
        }
        infoWindows = gmapsComponent.infoWindow( facilities[ idx ].name, facilities[ idx ].address, facilities[ idx ].type );
        infoWindows.open( {anchor: marker} );
      } );
    } );
  }
  
  async function initGoogleService() {
    const {AutocompleteService} = await google.maps.importLibrary( 'places' );
    autocomplete_service = new AutocompleteService();
    geocoder = new google.maps.Geocoder();
    if( initialLatLong[ 0 ]!==0 && initialLatLong[ 1 ]!==0 ) {
      await searchNearbyFacility();
    }
    await searchProperty( false );
    $mapComponent.addListener( 'click', () => {
      if( infoWindows ) {
        infoWindows.close();
      }
    } );
  }
  
  function searchLocationHandler() {
    autocomplete_service.getPlacePredictions( {
        input: input_search_value,
        types: ['establishment', 'geocode'],
      },
      function( predictions, status ) {
        if( status===google.maps.places.PlacesServiceStatus.OK ) {
          autocomplete_lists = predictions;
        }
      },
    );
  }
  
  function searchByLocationEvent( event ) {
    myLatLng.lat = event.detail.center.lat();
    myLatLng.lng = event.detail.center.lng();
  }
  
  function zoomEvent( event ) {
    console.log( event.detail.zoom );
  }
  
  async function searchByLocationHandler() {
    isSearchingMap = true;
    await searchProperty( true );
    await searchNearbyFacility();
    isSearchingMap = false;
  }
  
  async function searchByAddressHandler( place_id ) {
    await geocoder
      .geocode( {placeId: place_id} )
      .then( ( {results} ) => {
        if( results[ 0 ] ) {
          myLatLng.lat = results[ 0 ].geometry.location.lat();
          myLatLng.lng = results[ 0 ].geometry.location.lng();
        } else {
			    notifier.showError( 'No result found' );
        }
      } ).catch( ( e ) => {
		    notifier.showError(`Geocoder failed due to: ${e}` );
      } );
    autocomplete_lists = [];
    input_search_value = '';
    await searchProperty( true );
    await searchNearbyFacility();
    await gmapsComponent.setCentre( {
      lat: myLatLng.lat,
      lng: myLatLng.lng,
    } );
  }
  
  function showShareItems( idx ) {
    if( idx===shareItemIndex ) {
      return shareItemIndex = null;
    }
    return shareItemIndex = idx;
  }
  
  function copyToClipboard( text ) {
    shareItemIndex = null;
    navigator.clipboard.writeText( text );
	  notifier.showSuccess('Link copied to clipboard' );
  }
  
  function propertyUrl( id ) {
    let url = window.location.href.split( '#' )[ 0 ]
    return url + 'guest/property/' + id;
  }
  
  async function likeProperty( propId ) {
    await UserLikeProp( {
      propId: propId, // uint64
      like: true, // bool
    }, async res => {
      if( res.error ) {
        notifier.showError( res.error );
        return
      }
		  notifier.showSuccess('Property liked' );
    } )
  }

  let mobileClickSearchLocation = false;
</script>


<GoogleSdk on:ready={initGoogleService}/>

{#if mobileClickSearchLocation}
  <div class='mobile_autocomplete_container'>
    <header>
      <button class="back_button" on:click={() => mobileClickSearchLocation = false}>
        <Icon color='#475569' size={27} src={FaSolidAngleLeft} />
      </button>
      <div class='search_box'>
				<label for='search_location'>
					<Icon
						className='icon_search_location'
						color='#9fa9b5'
						size={18}
						src={FaSolidSearch}
					/>
				</label>
				<input
					bind:value={input_search_value}
					id='search_location'
					on:input={() => {
            searchLocationHandler();
          }}
					placeholder='Search for address...'
					type='text'
				/>
			</div>
    </header>
    <div class='autocomplete_container'>
      {#if autocomplete_lists.length}
        {#each autocomplete_lists as place}
          <button
            class='autocomplete_item'
            on:click|preventDefault={() => {
              mobileClickSearchLocation = false;
              searchByAddressHandler(place.place_id)
            }}
          >
            <Icon size={17} color='#9fa9b5' src={FaSolidMapMarkerAlt}/>
            <span>{place.description}</span>
          </button>
        {/each}
      {:else}
        <div class='no_autocomplete'>
          <div class='warn'>
            <Icon size={17} color='#475569' src={FaSolidReceipt}/>
            <span class='empty'>Address lists will appear here...</span>
          </div>
        </div>
      {/if}
    </div>
  </div>
{/if}

<div class='property_location_container'>
  <div class="search_mobile">
    <button class='search_location_btn' on:click={() => mobileClickSearchLocation = true}>
      <Icon
        color='#475569'
        size={18}
        src={FaSolidSearch}
      />
      <span>Search for address...</span>
    </button>
  </div>
	<div class='left'>
		<div class='props_container'>
			{#if randomProps.length}
				{#each randomProps as prop, index}
					<button
						class={propItemHighlight === index ? `prop_item highlight` : 'prop_item' }
						bind:this={propItemBinds[index]}
						on:mouseenter={() => highLightMapMarker.enter(index)}
						on:mouseleave={() => highLightMapMarker.leave(index)}
					>
						<picture class='img_container'>
							{#if prop.images && prop.images.length}
								<img src={prop.images[0]} alt=''/>
							{:else}
								<div class='image_empty'>
									<Icon size={40} className='no_image_icon' color='#848d96' src={FaSolidImage}/>
									<span>No Image !</span>
								</div>
							{/if}
						</picture>
						<div class='prop_info'>
							<div class='main_info'>
								<div class='top_xd'>
									<div class='label_info'>
										<div class={prop.purpose === 'rent' ? 'purpose label_rent' : 'purpose label_sale' }>
											{prop.purpose==='rent' ? $T.forRent : $T.onSale}
										</div>
										<div class='house_type'>
											<Icon size={12} className='house_type_icon' color='#475569' src={FaSolidHome}/>
											<span>{prop.houseType==="" ? 'House' : prop.houseType}</span>
										</div>
									</div>
									<div class='right_buttons'>
										<button class="like_btn" on:click={() => likeProperty(prop.id)}>
											<Icon color='#9fa9b5' className='like_icon' size={18} src={FaHeart}/>
										</button>
										<button class='share_btn' on:click={() => showShareItems(index)}>
											<Icon size={17} color='#9fa9b5' className='share_icon' src={FaSolidShareAlt}/>
										</button>
									</div>
									{#if shareItemIndex===index}
										<div class='share_container'>
											<button class='share_item copy' title='Copy link address'
											        on:click={() => copyToClipboard(propertyUrl(prop.id))}>
												<Icon size={14} color='#475569' src={FaCopy}/>
											</button>
											<a class='share_item'
											   aria-label="Share to Facebook"
											   href='https://www.facebook.com/sharer/sharer.php?u={propertyUrl(prop.id)}?utm_source=facebook&utm_medium=social&utm_campaign=user-share'
											   target="_blank"
											   rel="noopener"
											>
												<Icon size={14} color='#475569' src={FaBrandsFacebook}/>
											</a>
											<a class='share_item'
											   aria-label="Share to LinkedIn"
											   href='https://www.linkedin.com/shareArticle?mini=true&url={propertyUrl(prop.id)}&title=I%20Found%20Awesome%House%20{propertyUrl(prop.id)}'
											   target="_blank"
											   rel="noopener"
											>
												<Icon size={14} color='#475569' src={FaBrandsLinkedin}/>
											</a>
											<a class='share_item'
											   aria-label="Share to Twitter"
											   href='https://twitter.com/intent/tweet?url={propertyUrl(prop.id)}'
											   target="_blank"
											   rel="noopener"
											>
												<Icon size={14} color='#475569' src={FaBrandsTwitter}/>
											</a>
											<a class='share_item'
											   aria-label="Share to Telegram"
											   href='https://t.me/share/url?url={propertyUrl(prop.id)}'
											   target="_blank"
											   rel="noopener"
											>
												<Icon size={14} color='#475569' src={FaBrandsTelegram}/>
											</a>
											<a class='share_item'
											   aria-label="Share to WhatsApp"
											   href='https://api.whatsapp.com/send?text=I%20Found%20Awesome%20House%20{propertyUrl(prop.id)}'
											   target="_blank"
											   rel="noopener"
											>
												<Icon size={16} color='#475569' src={FaBrandsWhatsapp}/>
											</a>
										</div>
									{/if}
								</div>
								<div class='address'>
									<Icon size={17} className='icon_address' color='#f97316' src={FaSolidMapMarkerAlt}/>
									<span>{prop.formattedAddress==="" ? prop.address : prop.formattedAddress}</span>
								</div>
								<div class='feature'>
									<div class='item'>
										<div>
											<Icon size={13} className='icon_feature' color='#FFF' src={FaSolidBuilding}/>
											<span>{$T.floors}</span>
										</div>
										<span class='value'>{prop.numberOfFloors===0 ? 'no-data' : prop.numberOfFloors}</span>
									</div>
									<div class='item'>
										<div>
											<Icon size={14} className='icon_feature' color='#FFF' src={FaSolidBed}/>
											<span>{$T.bed}</span>
										</div>
										<span class='value'>{prop.bedroom===0 ? 'no-data' : prop.bedroom}</span>
									</div>
									<div class='item'>
										<div>
											<Icon size={13} className='icon_feature' color='#FFF' src={FaSolidBath}/>
											<span>{$T.bath}</span>
										</div>
										<span class='value'>{prop.bathroom===0 ? 'no-data' : prop.bathroom}</span>
									</div>
                  <div class='item sizeM2'>
										<div>
											<Icon size={13} className='icon_feature' color='#FFF' src={FaSolidRulerCombined}/>
											<span>Size</span>
										</div>
										<span class='value'>{prop.sizeM2} {$T.m}2</span>
									</div>
								</div>
							</div>
							<div class='secondary_info'>
								<div class='size'>
									<Icon size={12} color='#f97316' src={FaSolidRulerCombined}/>
									<span>{prop.sizeM2} {$T.m}2</span>
								</div>
								<div class='price'>
									<span class='agency_fee'>{$T.agencyFee} {prop.agencyFeePercent || '0'}%</span>
									<span class='last_price'>{formatPrice( prop.lastPrice || 0, 'TWD' )}</span>
								</div>
							</div>
						</div>
					</button>
				{/each}
			{:else }
				<div class='no_properties'>
					<div class='warn'>
						<Icon size={17} color='#475569' src={FaSolidBan}/>
						<span>No properties in this area</span>
					</div>
				</div>
			{/if}
		</div>
	</div>
	<div class='right'>
		<div class='map_container'>
			<button class='btn_sync_map' on:click={searchByLocationHandler}>
				{#if !isSearchingMap}
					<Icon color='#1080e8' size={12} src={FaSolidUndoAlt}/>
				{/if}
				{#if isSearchingMap}
					<Icon className="spin" color='#1080e8' size={12} src={FaSolidCircleNotch}/>
				{/if}
				<span>Search this area</span>
			</button>
			<GoogleMap
				bind:bounds={gmapBounds}
				bind:this={gmapsComponent}
				on:mapDragged={searchByLocationEvent}
				on:zoomChanged={zoomEvent}
				options={mapOptions}
			/>
		</div>
		<div class='search_by_address'>
			<div class='search_box'>
				<label for='search_location'>
					<Icon
						className='icon_search_location'
						color='#9fa9b5'
						size={18}
						src={FaSolidSearch}
					/>
				</label>
				<input
					bind:value={input_search_value}
					id='search_location'
					on:input={() => {
            searchLocationHandler();
          }}
					placeholder='Search for address...'
					type='text'
				/>
			</div>
			<div class='autocomplete_container'>
				{#if autocomplete_lists.length}
					{#each autocomplete_lists as place}
						<button
							class='autocomplete_item'
							on:click|preventDefault={() => searchByAddressHandler(place.place_id)}
						>
							<Icon size={17} color='#9fa9b5' src={FaSolidMapMarkerAlt}/>
							<span>{place.description}</span>
						</button>
					{/each}
				{:else}
					<div class='no_autocomplete'>
						<div class='warn'>
							<Icon size={17} color='#475569' src={FaSolidReceipt}/>
							<span class='empty'>Address lists will appear here...</span>
						</div>
					</div>
				{/if}
			</div>
		</div>
	</div>
</div>

<style>
    @keyframes spin {
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

    :global(.gm-style-iw) {
        overflow : visible !important;
    }

    :global(.gm-style-iw-c) {
        padding : 0 0 0 10px !important;
    }

    :global(.gm-ui-hover-effect) {
        background      : #EF4444 !important;
        border-radius   : 50% !important;
        display         : block !important;
        border          : 0 !important;
        margin          : 0 !important;
        padding         : 0 !important;
        text-transform  : none !important;
        appearance      : none !important;
        position        : absolute !important;
        cursor          : pointer !important;
        user-select     : none !important;
        top             : -10px !important;
        right           : -10px !important;
        width           : 27px !important;
        height          : 27px !important;
        z-index         : 999 !important;
        backdrop-filter : none !important;
        opacity         : inherit !important;
        color           : #FFF !important;
    }

    :global(.gm-ui-hover-effect:hover) {
        background : #E55D5D !important;
    }

    :global(.gm-ui-hover-effect span) {
        margin           : auto !important;
        background-color : #FFF !important;
    }

    .search_box {
        position : relative;
        width    : 100%;
        height   : fit-content;
        padding  : 20px 20px 0 20px;
    }

    .search_box input {
        width            : 100%;
        border           : 1px solid #CBD5E1;
        background-color : #F1F5F9;
        border-radius    : 8px;
        padding          : 12px 12px 12px 40px;
    }

    .search_box input:focus {
        border-color : #3B82F6;
        outline      : 1px solid #3B82F6;
    }

    .search_mobile, .mobile_autocomplete_container {
      display: none;
    }

    /* ================================ */

    .property_location_container {
        margin                : -40px auto 0 auto;
        border-radius         : 8px;
        filter                : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
        padding               : 20px;
        background-color      : white;
        color                 : #475569;
        width                 : 90%;
        min-height            : 1200px;
        height                : 1200px;
        display               : grid;
        grid-template-columns : 1.2fr 0.8fr;
        gap                   : 10px;
        max-width             : 100%;
        overflow-y            : scroll;
    }

    .property_location_container::-webkit-scrollbar-thumb,
    .property_location_container .left::-webkit-scrollbar-thumb,
    .property_location_container .right::-webkit-scrollbar-thumb {
        background-color : transparent;
    }

    .property_location_container::-webkit-scrollbar,
    .property_location_container .left::-webkit-scrollbar,
    .property_location_container .right::-webkit-scrollbar {
        width : 0;
    }

    .property_location_container::-webkit-scrollbar-track,
    .property_location_container .left::-webkit-scrollbar-track,
    .property_location_container .right::-webkit-scrollbar-track {
        background-color : transparent;
    }

    .property_location_container .left {
        height     : 100%;
        overflow-y : scroll;
    }

    .property_location_container .left .props_container {
        height         : 100%;
        display        : flex;
        flex-direction : column;
        gap            : 18px;
        overflow-y     : scroll;
    }

    .property_location_container .left .props_container .no_properties,
    .autocomplete_container .no_autocomplete {
        display         : flex;
        justify-content : center;
        align-items     : center;
        height          : 100%;
        width           : 100%;
    }

    .property_location_container .left .props_container .no_properties .warn,
    .autocomplete_container .no_autocomplete .warn {
        display          : flex;
        flex-direction   : row;
        align-items      : center;
        gap              : 8px;
        width            : fit-content;
        height           : fit-content;
        background-color : #F1F5F9;
        padding          : 15px;
        border-radius    : 8px;
        font-size        : 15px;
    }

    .property_location_container .left .props_container::-webkit-scrollbar-thumb {
        background-color : #3B82F6;
        border-radius    : 8px;
    }

    .property_location_container .left .props_container::-webkit-scrollbar {
        width : 8px;
    }

    .property_location_container .left .props_container::-webkit-scrollbar-track {
        background-color : transparent;
        margin           : 4px;
    }

    .property_location_container .left .props_container .prop_item {
        display        : flex;
        flex-direction : row;
        gap            : 15px;
        padding        : 10px 10px;
        margin-right   : 5px;
        border-radius  : 8px;
        cursor         : pointer;
        height         : 190px;
        min-height     : 190px;
        border         : 2px solid transparent;
        text-transform : capitalize;
        color          : #475569;
    }

    .property_location_container .left .props_container .prop_item.highlight {
        border : 2px solid #1080E8;
    }

    .property_location_container .left .props_container .prop_item:nth-child(odd) {
        background-color : #F1F5F9;
    }

    .property_location_container .left .props_container .prop_item:nth-child(even) {
        background-color : transparent;
    }

    .property_location_container .left .props_container .prop_item:hover .prop_info .main_info .address {
        text-decoration : underline;
    }

    .property_location_container .left .props_container .prop_item:hover .img_container .image_empty,
    .property_location_container .left .props_container .prop_item:hover .img_container img {
        transform : scale(1.20);
    }

    .property_location_container .left .props_container .prop_item .img_container {
        min-width     : 240px;
        width         : 240px;
        height        : 100%;
        overflow      : hidden;
        border        : 1px solid #CBD5E1;
        border-radius : 8px;
    }

    .property_location_container .left .props_container .prop_item .img_container img {
        object-fit          : cover;
        width               : 100%;
        height              : 100%;
        transition-duration : 75ms;
    }

    .property_location_container .left .props_container .prop_item .img_container .image_empty {
        object-fit          : cover;
        width               : 100%;
        height              : 100%;
        background-color    : #F1F5F9;
        display             : flex;
        flex-direction      : column;
        justify-content     : center;
        align-items         : center;
        gap                 : 5px;
        color: #848d96;
        transition-duration : 75ms;
    }

    .property_location_container .left .props_container .prop_item .prop_info {
        flex-grow       : 1;
        display         : flex;
        flex-direction  : column;
        justify-content : space-between;
        height          : 100%;
    }

    .property_location_container .left .props_container .prop_item .prop_info .main_info {
        display        : flex;
        flex-direction : column;
        gap            : 12px;
    }

    .property_location_container .left .props_container .prop_item .prop_info .main_info .top_xd {
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
        align-items     : center;
        align-content   : center;
        position        : relative;
    }

    .property_location_container .left .props_container .prop_item .prop_info .main_info .top_xd .right_buttons {
        display        : flex;
        flex-direction : row;
        align-items    : center;
        margin-right   : 4px;
        gap            : 18px;
    }

    .share_container {
        display          : flex;
        flex-direction   : row;
        gap              : 8px;
        position         : absolute;
        right            : 0;
        top              : 40px;
        padding          : 10px;
        border-radius    : 8px;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
        background-color : #FFF;
        height           : fit-content;
        width            : fit-content;
        border           : 1px solid #CBD5E1;
    }

    .share_item {
        padding       : 5px;
        border-radius : 5px;
        border        : none;
        background    : none;
        cursor        : pointer;
    }

    .share_item:hover {
        background-color : #F1F5F9;
    }

    .property_location_container .left .props_container .prop_item .prop_info .main_info .top_xd .share_btn {
        padding    : 0;
        border     : none;
        background : none;
        cursor     : pointer;
    }

    :global(.property_location_container .left .props_container .prop_item .prop_info .main_info .top_xd .share_btn:hover .share_icon),
    :global(.property_location_container .left .props_container .prop_item .prop_info .main_info .top_xd .like_btn:hover .like_icon) {
        fill : #F97316;
    }

    .property_location_container .left .props_container .prop_item .prop_info .main_info .top_xd .like_btn {
        padding    : 0;
        border     : none;
        background : none;
        cursor     : pointer;
    }

    .property_location_container .left .props_container .prop_item .prop_info .main_info .label_info {
        display        : flex;
        flex-direction : row;
        gap            : 8px;
    }

    .property_location_container .left .props_container .prop_item .prop_info .main_info .label_info .purpose,
    .property_location_container .left .props_container .prop_item .prop_info .main_info .label_info .house_type {
        padding        : 4px 10px;
        font-size      : 13px;
        border         : 1px solid #CBD5E1;
        border-radius  : 5px;
        display        : flex;
        flex-direction : row;
        align-items    : center;
        gap            : 7px;
        width          : fit-content;
    }

    .property_location_container .left .props_container .prop_item .prop_info .main_info .address {
        display        : flex;
        flex-direction : row;
        align-items    : flex-start;
        font-size      : 15px;
        gap            : 8px;
        text-align     : left;
    }

    .property_location_container .left .props_container .prop_item .prop_info .main_info .feature {
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
        align-items     : center;
        font-size       : 11px;
    }

    .property_location_container .left .props_container .prop_item .prop_info .main_info .feature .item {
        display        : flex;
        flex-direction : row;
        align-items: stretch;
        width          : fit-content;
        height         : fit-content;
        border-radius  : 5px;
    }

    .property_location_container .left .props_container .prop_item .prop_info .main_info .feature .item.sizeM2 {
        display: none;
    }

    .property_location_container .left .props_container .prop_item .prop_info .main_info .feature .item div {
        display                   : flex;
        flex-direction            : row;
        align-items               : center;
        padding                   : 3px 7px;
        background-color          : #F97316;
        border                    : 1px solid #F97316;
        gap                       : 4px;
        color                     : #FFFF;
        border-top-left-radius    : 5px;
        border-bottom-left-radius : 5px;
    }

    .property_location_container .left .props_container .prop_item .prop_info .main_info .feature .item .value {
        padding                    : 3px 7px;
        border                     : 1px solid #CBD5E1;
        border-top-right-radius    : 5px;
        border-bottom-right-radius : 5px;
    }

    .property_location_container .left .props_container .prop_item .prop_info .secondary_info {
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
        align-items     : flex-end;
    }

    .property_location_container .left .props_container .prop_item .prop_info .secondary_info .size {
        display        : flex;
        flex-direction : row;
        align-items    : center;
        gap            : 6px;
    }

    .property_location_container .left .props_container .prop_item .prop_info .secondary_info .price {
        display        : flex;
        flex-direction : column;
        align-items    : flex-end;
        gap            : 4px;
    }

    .property_location_container .left .props_container .prop_item .prop_info .secondary_info .price .agency_fee {
        font-size : 12px;
    }

    .property_location_container .left .props_container .prop_item .prop_info .secondary_info .price .last_price {
        font-size   : 20px;
        font-weight : 700;
    }

    .property_location_container .right {
        border-radius      : 8px;
        height             : 100%;
        position           : relative;
        display            : grid;
        grid-template-rows : 500px auto;
        gap                : 20px;
    }

    .property_location_container .right .search_by_address {
        border         : 1px solid #CBD5E1;
        border-radius  : 8px;
        display        : flex;
        flex-direction : column;
        height         : 100%;
        width          : 100%;
        gap            : 20px;
    }

    :global(.icon_search_location) {
        position : absolute;
        left     : 0;
        bottom   : 0;
        top      : 20px;
        z-index  : 40;
        margin   : auto 0 auto 32px;
    }

    .autocomplete_container {
        height         : 100%;
        display        : flex;
        flex-direction : column;
        padding        : 0;
        overflow       : auto;
        border-top     : 1px solid #CBD5E1;
    }

    .autocomplete_container {
        height         : 100%;
        display        : flex;
        flex-direction : column;
        overflow       : auto;
    }

    .autocomplete_container::-webkit-scrollbar-thumb {
        background-color : #3B82F6;
    }

    .autocomplete_container::-webkit-scrollbar {
        width : 8px;
    }

    .autocomplete_container::-webkit-scrollbar-track {
        background-color : transparent;
    }

    .autocomplete_container .autocomplete_item {
        display        : flex;
        flex-direction : row;
        gap            : 8px;
        align-items    : center;
        padding        : 10px;
        border         : none;
        background     : none;
        border-bottom  : 1px solid #CBD5E1;
        cursor         : pointer;
    }

    .autocomplete_container .autocomplete_item:hover {
        background-color : #F1F5F9;
    }

    .property_location_container .right .map_container {
        border        : 1px solid #CBD5E1;
        display       : block;
        position      : relative;
        width         : 100%;
        height        : 100%;
        border-radius : 8px;
        overflow      : hidden;
    }

    .property_location_container .right .map_container .btn_sync_map {
        position         : absolute;
        display          : flex;
        flex-direction   : row;
        gap              : 8px;
        align-items      : center;
        width            : fit-content;
        padding          : 0 20px;
        height           : 40px;
        top              : 10px;
        left             : auto;
        right            : 70px;
        border-radius    : 3px;
        font-size        : 15px;
        background-color : #FFF;
        border           : none;
        z-index          : 20;
        cursor           : pointer;
        color            : #475569;
        box-shadow       : 0 4px 24px 0 rgba(0, 0, 0, 0.25);
    }

    .property_location_container .right .map_container .btn_sync_map:hover {
        background-color : #F1F5F9;
    }

    /* Responsive to mobile device */
    @media (max-width: 768px) {
      .property_location_container {
        margin                : -40px auto 0 auto;
        padding               : 15px;
        min-height            : 700px;
        height                : 1200px;
        display               : flex;
        flex-direction: column;
        gap: 15px;
        max-width             : 100%;
        overflow-y            : scroll;
      }

      .mobile_autocomplete_container {
        display: flex;
        flex-direction: column;
        position: fixed;
        left: 0;
        bottom: 0;
        right: 0;
        top: 0;
        background-color: #FFF;
        z-index: 9999;
        overflow-y: scroll;
        padding: 20px;
        gap: 15px;
      }

      .mobile_autocomplete_container header {
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 10px;
      }

      .mobile_autocomplete_container header .back_button {
        width: fit-content;
        padding: 0;
        background: transparent;
        border: none;
      }

      .mobile_autocomplete_container header .search_box {
        padding: 0;
        flex-grow: 1;
      }

      :global(.icon_search_location) {
        top: 0;
        margin: auto 0 auto 15px;
      }

      .search_mobile {
        display: block;
      }

      .search_location_btn {
        border: none;
        background: transparent;
        padding  : 12px;
        border           : 1px solid #CBD5E1;
        background-color : #F1F5F9;
        border-radius    : 8px;
        width    : 100%;
        height   : fit-content;
        display: flex;
        flex-direction: row;
        gap: 12px;
        color: #475569;
        align-items: center;
        cursor: pointer;
      }
      .property_location_container .left .props_container::-webkit-scrollbar-thumb,
      .mobile_autocomplete_container::-webkit-scrollbar-thumb {
        background: transparent;
      }

      .property_location_container .left .props_container::-webkit-scrollbar,
      .mobile_autocomplete_container::-webkit-scrollbar {
        width : 0;
      }

      .property_location_container .left .props_container::-webkit-scrollbar-track,
      .mobile_autocomplete_container::-webkit-scrollbar-track {
        background-color : transparent;
      }
      .property_location_container .right {
        display: none;
      }

      .property_location_container .left .props_container .prop_item:hover .prop_info .main_info .address {
        text-decoration : none;
      }

      .property_location_container .left .props_container .prop_item .img_container {
        min-width     : 70px;
        width         : 70px;
        height        : 100%;
      }

      .property_location_container .left .props_container .prop_item {
        gap            : 10px;
        padding        : 20px 0 0 0;
        border-top        : 1px solid #CBD5E1;
        border-bottom: none;
        border-left: none;
        border-right: none;
        margin-right   : 0;
        border-radius  : 0;
        cursor         : pointer;
      }

      .property_location_container .left .props_container .prop_item:nth-child(1) {
        padding-top: 0;
        border-top: none;
      }

      .property_location_container .left .props_container .prop_item:nth-child(odd) {
        background-color : transparent;
      }

      .property_location_container .left .props_container .prop_item.highlight {
        border : none;
      }

      :global(.no_image_icon) {
        width: 35px;
        height: auto;
      }

      .property_location_container .left .props_container .prop_item .img_container .image_empty span {
        display: none;
      }

      .property_location_container .left .props_container .prop_item .prop_info .main_info .label_info {
        gap            : 4px;
      }

      .property_location_container .left .props_container .prop_item .prop_info .main_info .label_info .purpose,
      .property_location_container .left .props_container .prop_item .prop_info .main_info .label_info .house_type {
        padding        : 3px 7px;
        font-size      : 9px;
        border         : 1px solid #CBD5E1;
        border-radius  : 5px;
        display        : flex;
        flex-direction : row;
        align-items    : center;
        gap            : 4px;
        width          : fit-content;
      }

      :global(.house_type_icon) {
        width: 9px;
        height: auto;
      }

      .property_location_container .left .props_container .prop_item .prop_info .main_info .top_xd .right_buttons {
        gap            : 10px;
      }

      :global(.share_icon),
      :global(.like_icon) {
        width: 13px;
        height: auto;
      }

      .property_location_container .left .props_container .prop_item .prop_info .main_info .address {
        font-size      : 10px;
        align-items: center;
      }

      :global(.icon_address) {
        flex-shrink : 0;
        width: 14px;
        height: auto;
      }

      .property_location_container .left .props_container .prop_item .prop_info .main_info .address span {
        flex-shrink: 1;
      }

      .property_location_container .left .props_container .prop_item .prop_info .main_info .feature {
        display         : grid;
        grid-template-columns: 1fr 1fr;
        grid-template-rows: 1fr 1fr;
        font-size: 8px;
        line-height: 1em;
        gap: 8px;
      }

      .property_location_container .left .props_container .prop_item .prop_info .main_info .feature .item.sizeM2 {
        display: flex;
      }

      :global(.icon_feature) {
        width: 10px;
        height: auto;
      }

      .property_location_container .left .props_container .prop_item .prop_info .secondary_info {
        justify-content : flex-end;
      }

      .property_location_container .left .props_container .prop_item .prop_info .secondary_info .size {
        display        : none;
      }
    }
</style>