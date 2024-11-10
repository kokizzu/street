<script>
	import { Icon } from '../../node_modules/svelte-icons-pack/dist';
	import { RiSystemSearchLine, RiMapMapPin2Line} from '../../node_modules/svelte-icons-pack/dist/ri';
	import GoogleSdk from './GoogleSdk.svelte';
	import Growl from '../Growl.svelte'

	export let elevation;
	export let resolution;
	export let lat;
	export let lng;

	let scriptLoaded = false, cssLoaded = false;
	let viewer, tileset, handler;
	let streetViewInput, streetViewInputValue, autocompleteService, geocoder, elevationService;
	let showAutoCompleteList = false, autocompleteLists = [];
	let growl1 = Growl;
	let cameraOrientation;

	function init3dTiles() {
		// Enable simultaneous requests.
		Cesium.RequestScheduler.requestsByServer["tile.googleapis.com:443"] = 18;

		cameraOrientation = new Cesium.HeadingPitchRoll(
			4.6550106925119925,
			-0.2863894863138836,
			1.3561760425773173e-7
		);

		// Create the viewer.
		viewer = new Cesium.Viewer("cesiumContainer", {
			imageryProvider: false,
			baseLayerPicker: false,
			requestRenderMode: true,
			geocoder: false,
			timeline: false,
			animation: false,
			sceneModePicker: false,
			homeButton: false,
			infoBox: false,
			globe: false,
		});

		// Enable rendering the sky
		viewer.scene.skyAtmosphere.show = true;

		// Add 3D Tiles tileset.
		tileset = viewer.scene.primitives.add(
			new Cesium.Cesium3DTileset({
				url: "https://tile.googleapis.com/v1/3dtiles/root.json?key=AIzaSyBKF5w6NExgYbmNMvlbMqF6sH2X4dFvMBg",
				// This property is required to display attributions as required.
				showCreditsOnScreen: true,
			})
		);

		viewer.scene.camera.setView({
			destination: Cesium.Cartesian3.fromDegrees(
				lat,
				lng,
				elevation || 1500.0,
			),
			orientation: cameraOrientation,
		});

		handler = new Cesium.ScreenSpaceEventHandler(viewer.scene.canvas);
		handler.setInputAction(async (click) => {
			var cartesian = viewer.scene.pickPosition(new Cesium.Cartesian2(click.position.x, click.position.y));
			var ellipsoid = viewer.scene.primitives.get(0).ellipsoid;
			if (cartesian) {
				var cartographic = ellipsoid.cartesianToCartographic(cartesian);
				lat = Cesium.Math.toDegrees(cartographic.latitude);
				lng = Cesium.Math.toDegrees(cartographic.longitude);
				elevationService
					.getElevationForLocations({
						locations: [{lat, lng}],
					})
					.then(({ results }) => {
						if (results[0]) {
							elevation = results[0].elevation;
							resolution = results[0].resolution;
						}
				})
				.catch((e) =>
					alert('Elevation service failed due to: ' + e),
				)
			}
		}, Cesium.ScreenSpaceEventType.LEFT_CLICK);
	}

	async function initGoogleSdk() {
		const {AutocompleteService} = await google.maps.importLibrary('places');
		autocompleteService = new AutocompleteService();
		geocoder = new google.maps.Geocoder();
		elevationService = new google.maps.ElevationService();
	}

	function searchLocationHandler() {
		showAutoCompleteList = true;
		autocompleteService.getPlacePredictions({
				input: streetViewInputValue,
				types: ['establishment', 'geocode'],
			},
			function(predictions, status) {
				if (status === google.maps.places.PlacesServiceStatus.OK) {
					autocompleteLists = predictions;
				} else {
					alert('Cannot get address')
				}
			},
		);
	}

	window.addEventListener('click', () => {
		showAutoCompleteList = false;
		autocompleteLists = [];
	});

	async function searchByAddressHandler(place_id) {
		await geocoder
			.geocode({placeId: place_id})
			.then(async ({results}) => {
				if (results[0]) {
					const elevationResponse = await elevationService.getElevationForLocations({
						locations: [results[0].geometry.location],
					});
					if (!(elevationResponse.results && elevationResponse.results.length)) {
						alert(`Insufficient elevation data for place: ${results[0].formatted_address}`)
					}
					const elv = elevationResponse.results[0].elevation;
					elevation = elv;
					resolution = elevationResponse.results[0].resolution;
					lat = results[0].geometry.location.lat();
					lng = results[0].geometry.location.lng();
					viewer.scene.camera.setView({
						destination: Cesium.Cartesian3.fromDegrees(
							results[0].geometry.location.lng(),
							results[0].geometry.location.lat(),
							elv || 15000.0
						),
						orientation: cameraOrientation,
					})
					// rotateCameraAround(
					//   results[ 0 ].geometry.location,
					//   results[ 0 ].geometry.viewport,
					//   elv
					// );
				} else {
					alert('No result found');
				}
			}).catch((e) => {
				alert(`Geocoder failed due to: ${e}`);
			});
		autocompleteLists = [];
		showAutoCompleteList = false;
	}

	// Cesium file
	function checkFilesLoaded() {
		if (scriptLoaded && cssLoaded) {
			init3dTiles();
		}
	}

	const scriptElement = document.createElement("script");
	scriptElement.src = "https://ajax.googleapis.com/ajax/libs/cesiumjs/1.105/Build/Cesium/Cesium.js";
	scriptElement.onload = () => {
		scriptLoaded = true;
		checkFilesLoaded();
	};
	document.body.appendChild(scriptElement);
	const linkElement = document.createElement("link");
	linkElement.rel = "stylesheet";
	linkElement.href = "https://ajax.googleapis.com/ajax/libs/cesiumjs/1.105/Build/Cesium/Widgets/widgets.css";
	linkElement.onload = () => {
		cssLoaded = true;
		checkFilesLoaded();
	};
	document.head.appendChild(linkElement);
</script>


<GoogleSdk on:ready={initGoogleSdk}/>
<div class="streetview_input_box">
	<Icon
		className='icon_search_location'
		color='#9fa9b5'
		size={18}
		src={RiSystemSearchLine}
	/>
	<input
		bind:this={streetViewInput}
		bind:value={streetViewInputValue}
		id="pacViewPlace"
		name="pacViewPlace"
		on:input={searchLocationHandler}
		placeholder="Enter a location..."
		type="text"
	/>
</div>
{#if showAutoCompleteList === true}
	<div class='autocomplete_container'>
	{#if autocompleteLists.length}
		{#each autocompleteLists as place}
			<button
				class='autocomplete_item'
				on:click|preventDefault={() => searchByAddressHandler(place.place_id)}
			>
				<Icon size={17} color='#9fa9b5' src={RiMapMapPin2Line}/>
				<span>{place.description}</span>
			</button>
		{/each}
	{:else}
		<button class='autocomplete_item' >
			<Icon size={17} color='#9fa9b5' src={RiMapMapPin2Line}/>
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
        cursor        : pointer;
    }
</style>