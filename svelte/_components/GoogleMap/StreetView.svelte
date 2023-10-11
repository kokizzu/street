<script>
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidSearch from "svelte-icons-pack/fa/FaSolidSearch";
  import GoogleSdk from './GoogleSdk.svelte';
  import FaSolidMapMarkerAlt from "svelte-icons-pack/fa/FaSolidMapMarkerAlt";
  
  // TODO: use it to payload property
  export let elevation;
  export let resolution;
  
  let scriptLoaded = false, cssLoaded = false;
  let viewer, tileset;
  let streetViewInput, streetViewInputValue, autocompleteService, geocoder;
  let showAutoCompleteList = false, autocompleteLists = [];
  
  function init3dTiles() {
    // Enable simultaneous requests.
    Cesium.RequestScheduler.requestsByServer[ "tile.googleapis.com:443" ] = 18;
    
    // Create the viewer.
    viewer = new Cesium.Viewer( 'cesiumContainer', {
      imageryProvider: false,
      baseLayerPicker: false,
      infobox: false,
      geocoder: false,
      globe: false,
      // https://cesium.com/blog/2018/01/24/cesium-scene-rendering-performance/#enabling-request-render-mode
      requestRenderMode: true,
    } );
    
    // Add 3D Tiles tileset.
    tileset = viewer.scene.primitives.add( new Cesium.Cesium3DTileset( {
      url: "https://tile.googleapis.com/v1/3dtiles/root.json?key=AIzaSyBKF5w6NExgYbmNMvlbMqF6sH2X4dFvMBg",
      // This property is needed to appropriately display attributions
      // as required.
      showCreditsOnScreen: true,
    } ) );
  }
  
  async function initGoogleSdk() {
    const {AutocompleteService} = await google.maps.importLibrary( 'places' );
    autocompleteService = new AutocompleteService();
    geocoder = new google.maps.Geocoder();
  }
  
  // Point the camera at a location and elevation, at a viewport-appropriate distance.
  function pointCameraAt( location, viewport, elevation ) {
    const distance = Cesium.Cartesian3.distance(
      Cesium.Cartesian3.fromDegrees(
        viewport.getSouthWest().lng(), viewport.getSouthWest().lat(), elevation ),
      Cesium.Cartesian3.fromDegrees(
        viewport.getNorthEast().lng(), viewport.getNorthEast().lat(), elevation )
    ) / 2;
    const target = new Cesium.Cartesian3.fromDegrees( location.lng(), location.lat(), elevation );
    const pitch = -Math.PI / 4;
    const heading = 0;
    viewer.camera.lookAt( target, new Cesium.HeadingPitchRange( heading, pitch, distance ) );
  }
  
  // Rotate the camera around a location and elevation, at a viewport-appropriate distance.
  let unsubscribe = null;
  
  function rotateCameraAround( location, viewport, elevation ) {
    if( unsubscribe ) unsubscribe();
    pointCameraAt( location, viewport, elevation );
    unsubscribe = viewer.clock.onTick.addEventListener( () => {
      viewer.camera.rotate( Cesium.Cartesian3.UNIT_Z );
    } );
  }
  
  function searchLocationHandler() {
    autocompleteService.getPlacePredictions( {
        input: streetViewInputValue,
        types: ['establishment', 'geocode'],
      },
      function( predictions, status ) {
        if( status===google.maps.places.PlacesServiceStatus.OK ) {
          autocompleteLists = predictions;
        }
      },
    );
  }
  
  async function searchByAddressHandler( place_id ) {
    await geocoder
      .geocode( {placeId: place_id} )
      .then( async ( {results} ) => {
        if( results[ 0 ] ) {
          // Get place elevation using the ElevationService.
          const elevatorService = new google.maps.ElevationService();
          const elevationResponse = await elevatorService.getElevationForLocations( {
            locations: [results[ 0 ].geometry.location],
          } );
          if( !(elevationResponse.results && elevationResponse.results.length) ) {
            alert( `Insufficient elevation data for place: ${results[ 0 ].formatted_address}` );
          }
          const elevation = elevationResponse.results[ 0 ].elevation || 10;
          rotateCameraAround(
            results[ 0 ].geometry.location,
            results[ 0 ].geometry.viewport,
            elevation
          );
        } else {
          alert( 'No result found' );
        }
      } ).catch( ( e ) => {
        alert( `Geocoder failed due to: ${e}` );
      } );
    autocompleteLists = [];
    showAutoCompleteList = false;
  }
  
  // Cesium file
  function checkFilesLoaded() {
    if( scriptLoaded && cssLoaded ) {
      init3dTiles();
    }
  }
  
  const scriptElement = document.createElement( "script" );
  scriptElement.src = "https://ajax.googleapis.com/ajax/libs/cesiumjs/1.105/Build/Cesium/Cesium.js";
  scriptElement.onload = () => {
    scriptLoaded = true;
    checkFilesLoaded();
  };
  document.body.appendChild( scriptElement );
  const linkElement = document.createElement( "link" );
  linkElement.rel = "stylesheet";
  linkElement.href = "https://ajax.googleapis.com/ajax/libs/cesiumjs/1.105/Build/Cesium/Widgets/widgets.css";
  linkElement.onload = () => {
    cssLoaded = true;
    checkFilesLoaded();
  };
  document.head.appendChild( linkElement );
</script>

<GoogleSdk on:ready={initGoogleSdk}/>
<div class="streetview_input_box">
	<Icon
		className='icon_search_location'
		color='#9fa9b5'
		size={18}
		src={FaSolidSearch}
	/>
	<input
		bind:this={streetViewInput}
		bind:value={streetViewInputValue}
		id="pacViewPlace"
		name="pacViewPlace"
		on:click={() => showAutoCompleteList = !showAutoCompleteList}
		on:input={searchLocationHandler}
		placeholder="Enter a location..."
		type="text"
	/>
</div>
{#if showAutoCompleteList===true}
	<div class='autocomplete_container'>
		{#if autocompleteLists.length}
			{#each autocompleteLists as place}
				<button
					class='autocomplete_item'
					on:click|preventDefault={() => searchByAddressHandler(place.place_id)}
				>
					<Icon size={17} color='#9fa9b5' src={FaSolidMapMarkerAlt}/>
					<span>{place.description}</span>
				</button>
			{/each}
		{:else}
			<button
				class='autocomplete_item'
			>
				<Icon size={17} color='#9fa9b5' src={FaSolidMapMarkerAlt}/>
				<span>Place name...</span>
			</button>
		{/if}
	</div>
{/if}
<div id="cesiumContainer"></div>

<style>
    .streetview_input_box {
        position : absolute;
        z-index  : 20;
        left     : 20px;
        top      : 10px;
    }

    :global(.icon_search_location) {
        position : absolute;
        left     : 0;
        bottom   : 0;
        top      : 10px;
        z-index  : 40;
        margin   : auto 0 auto 13px;
        cursor   : pointer;
    }

    :global(.icon_search_location:hover) {
        fill : #F97316;
    }

    .streetview_input_box #pacViewPlace {
        margin-top    : 10px;
        padding       : 12px 12px 12px 40px;
        border-radius : 8px;
        border        : none;
        width         : 450px;
        filter        : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    }

    .streetview_input_box #pacViewPlace:focus {
        border-color : #F97316;
        outline      : 2px solid #F97316;
    }

    .autocomplete_container {
        position         : absolute;
        z-index          : 20;
        left             : 20px;
        top              : 70px;
        width            : 450px;
        min-height       : 200px;
        max-height       : 400px;
        background-color : #FFF;
        border-radius    : 8px;
        overflow         : auto;
        display          : flex;
        flex-direction   : column;
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

    #cesiumContainer {
        height        : 100%;
        width         : 100%;
        overflow      : hidden;
        border-radius : 8px;
    }
</style>