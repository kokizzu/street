<script>
  // @ts-nocheck
  import { onMount } from "svelte";
  import { UserSearchProp } from "jsApi.GEN";
  import { formatPrice } from "./formatter";

  import Icon from "svelte-icons-pack/Icon.svelte";
  import FaSolidSearch from "svelte-icons-pack/fa/FaSolidSearch";
  import FaSolidHotel from "svelte-icons-pack/fa/FaSolidHotel";
  import FaSolidMapMarkedAlt from "svelte-icons-pack/fa/FaSolidMapMarkedAlt";
  import FaSolidMapMarkerAlt from "svelte-icons-pack/fa/FaSolidMapMarkerAlt";
  import FaSolidImage from "svelte-icons-pack/fa/FaSolidImage";
  import FaSolidHome from "svelte-icons-pack/fa/FaSolidHome";
  import FaSolidRulerCombined from "svelte-icons-pack/fa/FaSolidRulerCombined";
  import FaSolidBuilding from "svelte-icons-pack/fa/FaSolidBuilding";
  import FaSolidBath from "svelte-icons-pack/fa/FaSolidBath";
  import FaSolidBed from "svelte-icons-pack/fa/FaSolidBed";

  let random_props = [];
  onMount( async () => {
    await UserSearchProp({}, async res => {
      console.log("Property : ", res.properties);
      random_props = res.properties;
    })
    await initMap();
  } );

  // Maps
  // const google_map_api_key = "AIzaSyBKF5w6NExgYbmNMvlbMqF6sH2X4dFvMBg";
  let map;
  let autocomplete;
  let map_container;
  let map_autocomplete;
  let myLatLng = {lat: 23.6978, lng: 120.9605};
  
  async function initMap() {
    let markers = [
      {lat: -34.03360401120961, lng: 149.86401361846924},
      {lat: -34.40606480160426, lng: 149.94091791534424},
      {lat: -34.53286884912313, lng: 149.08398432159424},
      {lat: -34.242748228904865, lng: 149.32568354034424},
      {lat: -34.87157170169384, lng: 150.00683588409424},
      {lat: -34.18369473718617, lng: 150.24853510284424},
    ];
    const { Autocomplete } = await google.maps.importLibrary( 'places');
    autocomplete = new Autocomplete(
       map_autocomplete,
       {
         types: ['establishment'],
         componentRestrictions: {'country': ['AU']},
         fields: ['place_id', 'geometry', 'name']
       }
    )
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

  const srch_loc = "location", srch_prop = "property";
  let search_mode = srch_loc;
  async function toggleSearchMode( srchMode ) {
    search_mode = srchMode;
    if (search_mode === srch_loc) {await initMap()}
  }
  
  let searchbox_value; // bind value
  
  let props_by_loc = []; // Fill this array when search by location
  let show_area_lists = false;
  let area_list_predictions = {};
  
  // async function AreaListHandler() {
  //   show_area_lists = true;
  //   const resp = await fetch(
  //     "https://maps.googleapis.com/maps/api/place/autocomplete/json" +
  //     `?input=${searchbox_value}` +
  //     `&location=${myLatLng.lat}%2C${myLatLng.lng}` +
  //     "&types=geocode" +
  //     `&key=${google_map_api_key}`,
  //     {
  //        method: "GET"
  //     }
  //   );
  //   area_list_predictions = await resp.json();
  //   console.log(area_list_predictions)
  // }
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
      v: 'weekly'
    } );
  </script>
<!--  <script-->
<!--	  async defer-->
<!--     src="https://maps.googleapis.com/maps/api/js?key=AIzaSyBKF5w6NExgYbmNMvlbMqF6sH2X4dFvMBg&libraries=places&callback=initSearchMapAutoComplete">-->
<!--  </script>-->
</svelte:head>

<div class="property_location_container">
  <div class="header">
    <div class="tabs">
      <button class:active={search_mode === srch_loc} on:click={() => toggleSearchMode(srch_loc)}>
        <Icon size={15} color={search_mode === srch_loc ? "#3b82f6" : "#475569"} src={FaSolidMapMarkedAlt} />
        <span>Search by Location</span>
      </button>
      <button class:active={search_mode === srch_prop} on:click={() => toggleSearchMode(srch_prop)}>
        <Icon size={14} color={search_mode === srch_prop ? "#3b82f6" : "#475569"} src={FaSolidHotel} />
        <span>Search by Properties</span>
      </button>
    </div>
    <div class="search_box">
      <label for="search_location">
        <Icon size={18} className="icon_search_location" color="#9fa9b5" src={FaSolidSearch} />
      </label>
      <input type="text" id="search_location" placeholder="Search property..." bind:this={map_autocomplete} />
    </div>
  </div>
  {#if search_mode === srch_loc}
    <div class="search_by_location">
      <div class="left">
        <div class="props_container">
          {#each random_props as prop}
            <div class="prop_item">
              <div class="img_container">
                {#if prop.images && prop.images.length}
                  <img src={prop.images[0]} alt="" />
                {:else}
                  <div class="image_empty">
                    <Icon size={40} color="#475569" src={FaSolidImage} />
                    <span>No Image !</span>
                  </div>
                {/if}
              </div>
              <div class="prop_info">
                <div class="main_info">
                  <div class="label_info">
                    <div class="purpose">On {prop.purpose === "" ? 'Sale' : prop.purpose}</div>
                    <div class="house_type">
                      <Icon size={12} color="#475569" src={FaSolidHome} />
                      <span>{prop.houseType === "" ? 'House' : prop.houseType}</span>
                    </div>
                  </div>
                  <div class="address">
                    <Icon size={17} color="#f97316" src={FaSolidMapMarkerAlt} />
                    <span>{prop.formattedAddress === "" ? prop.address : prop.formattedAddress}</span>
                  </div>
                  <div class="feature">
                    <div class="item">
                      <Icon size={13} color="#f97316" src={FaSolidBuilding} />
                      <span><b>Floor</b>: {prop.numberOfFloors || 0}</span>
                    </div>
                    <div class="item">
                      <Icon size={14} color="#f97316" src={FaSolidBed} />
                      <span><b>Beds</b>: {prop.bedroom || 0}</span>
                    </div>
                    <div class="item">
                      <Icon size={14} color="#f97316" src={FaSolidBath} />
                      <span><b>Baths</b>: {prop.bathroom || 0}</span>
                    </div>
                  </div>
                </div>
                <div class="secondary_info">
                  <div class="size">
                    <Icon size={12} color="#f97316" src={FaSolidRulerCombined} />
                    <span>{prop.sizeM2} M2</span>
                  </div>
                  <div class="price">
                    <span class="agency_fee">Agency Fee: {prop.agencyFeePercent || '0'}%</span>
                    <span class="last_price">{formatPrice(prop.lastPrice || 0, 'USD')}</span>
                  </div>
                </div>
              </div>
            </div>
          {/each}
        </div>
      </div>
      <div class="right">
        <div class='map_container' bind:this={map_container}>
          <!-- Map goes here, rendered automatically -->
        </div>
      </div>
    </div>
  {:else if search_mode === srch_prop}
    <div class="search_by_property"></div>
  {/if}
</div>

<style>
  .property_location_container {
    position: relative;
    margin-top: -40px;
    margin-left: auto;
    margin-right: auto;
    border-radius: 8px;
    filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    padding: 25px;
    background-color: white;
    color: #475569;
    width: 88%;
    height: 800px;
    display: flex;
    flex-direction: column;
    gap: 20px;
  }
  .property_location_container .header {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }
  .property_location_container .header .tabs {
    display: flex;
    flex-direction: row;
    gap: 15px;
  }
  .property_location_container .header .tabs button {
    padding: 12px 20px;
    border-radius: 8px;
    border: 1px solid #cbd5e1;
    background-color: transparent;
    font-weight: 600;
    cursor: pointer;
    color: #475569;
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 8px;
  }
  .property_location_container .header .tabs button.active {
    border: 1px solid #3b82f6;
    color: #3b82f6;
  }
  .property_location_container .header .tabs button:hover {
    background-color: #f1f5f9;
  }
  .property_location_container .header .search_box {
    position: relative;
    width: 450px;
  }
  :global(.icon_search_location) {
    position: absolute;
    left: 0;
    bottom: 0;
    top: 0;
    z-index: 40;
    margin: auto 12px;
  }
  .property_location_container .header .search_box input {
    width: 100%;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    border-radius: 8px;
    padding: 12px 12px 12px 40px;
  }
  .property_location_container .header .search_box input:focus {
    border-color: #3b82f6;
    outline: 1px solid #3b82f6;
  }

  .search_by_location {
    display: grid;
    grid-template-columns: auto 500px;
    gap: 20px;
    flex-grow: 1;
    max-width: 100%;
    overflow-y: scroll;
  }
  .search_by_location::-webkit-scrollbar-thumb {
    background-color : transparent;
  }
  .search_by_location::-webkit-scrollbar {
    width: 0;
  }
  .search_by_location::-webkit-scrollbar-track {
    background-color : transparent;
  }
  .search_by_location .left {
    height: 100%;
    border-top: 1px solid #cbd5e1;
    border-bottom: 1px solid #cbd5e1;
    overflow-y: scroll;
  }
  .search_by_location .left::-webkit-scrollbar-thumb {
    background-color : transparent;
  }
  .search_by_location .left::-webkit-scrollbar {
    width: 0;
  }
  .search_by_location .left::-webkit-scrollbar-track {
    background-color : transparent;
  }
  .search_by_location .left .props_container {
    height: 100%;
    display: flex;
    flex-direction: column;
    gap: 18px;
    padding: 18px 0;
    overflow: auto;
  }
  .search_by_location .left .props_container::-webkit-scrollbar-thumb {
    background-color : #3b82f6;
  }
  .search_by_location .left .props_container::-webkit-scrollbar {
    width: 8px;
  }
  .search_by_location .left .props_container::-webkit-scrollbar-track {
    background-color : transparent;
  }
  .search_by_location .left .props_container .prop_item {
    display: flex;
    flex-direction: row;
    gap: 15px;
    padding-right: 15px;
    border-radius: 8px;
    cursor: pointer;
  }
  .search_by_location .left .props_container .prop_item:hover .prop_info .main_info .address {
    text-decoration: underline;
  }
  .search_by_location .left .props_container .prop_item .img_container {
    width: 240px;
    height: 170px;
  }
  .search_by_location .left .props_container .prop_item .img_container img {
    object-fit: cover;
    width: 100%;
    height: 100%;
  }
  .search_by_location .left .props_container .prop_item .img_container .image_empty {
    border-radius: 8px;
    object-fit: cover;
    width: 100%;
    height: 100%;
    background-color: #f1f5f9;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 5px;
    border: 1px solid #cbd5e1;
  }
  .search_by_location .left .props_container .prop_item .prop_info {
    flex-grow: 1;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }
  .search_by_location .left .props_container .prop_item .prop_info .main_info {
    display: flex;
    flex-direction: column;
    gap: 16px;
  }
  .search_by_location .left .props_container .prop_item .prop_info .main_info .label_info {
    display: flex;
    flex-direction: row;
    gap: 8px;
  }
  .search_by_location .left .props_container .prop_item .prop_info .main_info .label_info .purpose,
  .search_by_location .left .props_container .prop_item .prop_info .main_info .label_info .house_type {
    padding: 4px 10px;
    font-size: 13px;
    border: 1px solid #cbd5e1;
    border-radius: 5px;
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 7px;
    width: fit-content;
  }
  .search_by_location .left .props_container .prop_item .prop_info .main_info .address {
    display: flex;
    flex-direction: row;
    align-items: flex-start;
    font-size: 15px;
    gap: 8px;
  }
  .search_by_location .left .props_container .prop_item .prop_info .main_info .feature {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    font-size: 13px;
  }
  .search_by_location .left .props_container .prop_item .prop_info .main_info .feature .item {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 8px;
  }
  .search_by_location .left .props_container .prop_item .prop_info .secondary_info {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: flex-end;
  }
  .search_by_location .left .props_container .prop_item .prop_info .secondary_info .size {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 6px;
  }
  .search_by_location .left .props_container .prop_item .prop_info .secondary_info .price {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    gap: 4px;
  }
  .search_by_location .left .props_container .prop_item .prop_info .secondary_info .price .agency_fee {
    font-size: 12px;
  }
  .search_by_location .left .props_container .prop_item .prop_info .secondary_info .price .last_price {
    font-size: 20px;
    font-weight: 700;
  }
  .search_by_location .left .props_container .prop_item .prop_info .secondary_info
  .search_by_location .right {
    border: 1px solid #cbd5e1;
    border-radius: 8px;
  }
  .search_by_location .right .map_container {
    display: block;
    width: 100%;
    height: 100%;
    border-radius: 8px;
  }
</style>