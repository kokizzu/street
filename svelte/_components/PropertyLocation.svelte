<script>
  // @ts-nocheck
  import { onMount } from "svelte";
  import { UserSearchProp, UserNearbyFacilities } from "jsApi.GEN";
  import { formatPrice } from "./formatter";
  import { GoogleMap, GoogleSdk } from "./GoogleMap/components";
  
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
  import FaSolidInfoCircle from "svelte-icons-pack/fa/FaSolidInfoCircle";
  import FaSolidUndoAlt from "svelte-icons-pack/fa/FaSolidUndoAlt";
  import FaSolidBan from "svelte-icons-pack/fa/FaSolidBan";

  let random_props = [];
  let facilities = [];
  let markers = [];
  onMount( async () => {
    await UserSearchProp({}, async res => {
      random_props = res.properties;
      console.log(random_props)
    });
    await UserNearbyFacilities({
      centerLat: myLatLng.lat,
      centerLong: myLatLng.lng,
    }, async res => {
      facilities = await res.facilities;
      facilities.forEach(fac => {
        markers.push(gmapsComponent.createMarker({
          lat: fac.lat,
          lng: fac.lng,
        }));
      });
    })
  } );
  // Maps
  let gmapsComponent;
  let myLatLng = {lat: 23.6978, lng: 120.9605};
  let mapOptions = {
    center: myLatLng,
    zoom: 8,
    mapTypeId: 'roadmap',
    mapId: 'street_project',
  }
  let places_service;
  let geocoder;
  // Search Autocomplete
  let input_search_value, autocomplete_service;
  let autocomplete_lists = [];
  
  async function initGoogleService() {
    const { AutocompleteService } = await google.maps.importLibrary('places');
    autocomplete_service = new AutocompleteService();
    geocoder = new google.maps.Geocoder();
  }
  
  function searchLocationHandler() {
    autocomplete_service.getPlacePredictions(
      {
        input: input_search_value,
        types: ["establishment", "geocode"]
      },
      function (predictions, status) {
        if (status === google.maps.places.PlacesServiceStatus.OK) {
          autocomplete_lists = predictions;
          console.log(predictions);
        }
      }
    )
  }
  
  function searchByLocationEvent( event ) {
    myLatLng.lat = event.detail.center.lat();
    myLatLng.lng = event.detail.center.lng();
  }
  async function searchByLocationHandler() {
    await UserSearchProp({
      centerLat: myLatLng.lat,
      centerLong: myLatLng.lng,
      offset: 0,
      limit: 0,
      maxDistanceKM: 500,
    }, async res => {
      random_props = res.properties || [];
    });
    
    await UserNearbyFacilities({
      centerLat: myLatLng.lat,
      centerLong: myLatLng.lng,
    }, async res => {
      facilities = await res.facilities;
      markers = gmapsComponent.clearMarkers(markers);
      facilities.forEach(fac => {
        markers.push(gmapsComponent.createMarker({
          lat: fac.lat,
          lng: fac.lng,
        }));
      });
    });
  }
  
  async function searchByAddressHandler( place_id ) {
    let lat, lng;
    await geocoder
      .geocode({ placeId: place_id})
      .then(({ results }) => {
        if (results[0]) {
          lat = results[0].geometry.location.lat();
          lng = results[0].geometry.location.lng()
        } else {
          alert("No result found")
        }
      }).catch((e) => {
        alert("Geocoder failed due to: " + e)
      });
    await UserSearchProp({
      centerLat: lat,
      centerLong: lng,
      offset: 0,
      limit: 0,
      maxDistanceKM: 500,
    }, async res => {
      random_props = res.properties;
      autocomplete_lists = [];
      input_search_value = "";
    });
  }

  // Toggle search mode, currently focus on `Search by location`
  const srch_loc = "location", srch_addrs = "address";
  let search_mode = srch_loc;
  function toggleSearchMode( srchMode ) {
    search_mode = srchMode;
  }
</script>

<GoogleSdk on:ready={initGoogleService} />
<div class="property_location_container">
  <div class="header">
    <div class="tabs">
      <button class:active={search_mode === srch_loc} on:click={() => toggleSearchMode(srch_loc)}>
        <Icon size={15} color={search_mode === srch_loc ? "#3b82f6" : "#475569"} src={FaSolidMapMarkedAlt} />
        <span>Search by Location</span>
      </button>
      <button class:active={search_mode === srch_addrs} on:click={() => toggleSearchMode(srch_addrs)}>
        <Icon size={14} color={search_mode === srch_addrs ? "#3b82f6" : "#475569"} src={FaSolidHotel} />
        <span>Search by Address</span>
      </button>
    </div>
    <div class="info">
      <Icon size={12} color="#475569" src={FaSolidInfoCircle} />
      <span>
        {search_mode === srch_loc
          ? "Drag map to find nearby property"
          : "Type on search bar to find properties near an address"
        }
      </span>
    </div>
  </div>
  <div class="content">
    <div class="left">
      <div class="props_container">
        {#if random_props.length}
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
                  <div class={prop.purpose === 'sell'
                    ? `purpose label_sale`
                    : `purpose label_rent`
                  }>
                    {prop.purpose === 'sell' ? `On Sale` : `For Rent`}
                  </div>
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
        {:else }
          <div class="no_properties">
            <div class="warn">
              <Icon size={17} color="#475569" src={FaSolidBan} />
              <span>No properties in this area</span>
            </div>
          </div>
        {/if}
      </div>
    </div>
    <div class="right">
      {#if search_mode === srch_loc}
        <div class='map_container'>
          <button class="btn_sync_map" on:click={searchByLocationHandler}>
            <Icon size={12} color="#475569" src={FaSolidUndoAlt} />
            <span>Search this area</span>
          </button>
          <GoogleMap options={mapOptions} bind:this={gmapsComponent} on:mapDragged={searchByLocationEvent}/>
        </div>
      {/if}
      {#if search_mode === srch_addrs}
        <div class="search_by_address">
          <div class="search_box">
            <label for="search_location">
              <Icon
                size={18}
                className="icon_search_location"
                color="#9fa9b5"
                src={FaSolidSearch}
              />
            </label>
            <input
              type="text"
              id="search_location"
              placeholder="Search property..."
              on:input={() => {
                searchLocationHandler();
              }}
              bind:value={input_search_value}
            />
          </div>
          <div class="autocomplete_container">
            {#if autocomplete_lists.length}
              {#each autocomplete_lists as place}
                <button class="autocomplete_item" on:click|preventDefault={() => searchByAddressHandler(place.place_id)}>
                  <Icon size={17} color="#9fa9b5" src={FaSolidMapMarkerAlt} />
                  <span>{place.description}</span>
                </button>
              {/each}
            {:else}
              <span class="empty">Address lists will appear here...</span>
            {/if}
          </div>
        </div>
      {/if}
    </div>
  </div>
</div>

<style>
  .property_location_container {
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
    align-items: center;
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
  .property_location_container .header .info {
    display: flex;
    flex-direction: row;
    align-items: center;
    font-size: 12px;
    gap: 8px;
    margin-right: 20px;
  }
  
  .property_location_container .content {
    display: grid;
    grid-template-columns: auto 500px;
    gap: 20px;
    flex-grow: 1;
    max-width: 100%;
    overflow-y: scroll;
  }
  .property_location_container .content::-webkit-scrollbar-thumb,
  .property_location_container .content .left::-webkit-scrollbar-thumb,
  .property_location_container .content .right::-webkit-scrollbar-thumb {
    background-color : transparent;
  }
  .property_location_container .content::-webkit-scrollbar,
  .property_location_container .content .left::-webkit-scrollbar,
  .property_location_container .content .right::-webkit-scrollbar {
    width: 0;
  }
  .property_location_container .content::-webkit-scrollbar-track,
  .property_location_container .content .left::-webkit-scrollbar-track,
  .property_location_container .content .right::-webkit-scrollbar-track {
    background-color : transparent;
  }
  .property_location_container .content .left {
    height: 100%;
    border-top: 1px solid #cbd5e1;
    border-bottom: 1px solid #cbd5e1;
    overflow-y: scroll;
  }
  .property_location_container .content .left .props_container {
    height: 100%;
    display: flex;
    flex-direction: column;
    gap: 18px;
    padding: 18px 0;
    overflow: auto;
  }
  .property_location_container .content .left .props_container .no_properties {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    width: 100%;
  }
  .property_location_container .content .left .props_container .no_properties .warn {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 8px;
    width: fit-content;
    height: fit-content;
    background-color: #f1f5f9;
    padding: 15px;
    border-radius: 8px;
    font-size: 15px;
  }
  .property_location_container .content .left .props_container::-webkit-scrollbar-thumb {
    background-color : #3b82f6;
  }
  .property_location_container .content .left .props_container::-webkit-scrollbar {
    width: 8px;
  }
  .property_location_container .content .left .props_container::-webkit-scrollbar-track {
    background-color : transparent;
  }
  .property_location_container .content .left .props_container .prop_item {
    display: flex;
    flex-direction: row;
    gap: 15px;
    padding-right: 15px;
    border-radius: 8px;
    cursor: pointer;
    height: 190px;
    min-height: 190px;
  }
  .property_location_container .content .left .props_container .prop_item:hover .prop_info .main_info .address {
    text-decoration: underline;
  }
  .property_location_container .content .left .props_container .prop_item:hover .img_container .image_empty,
  .property_location_container .content .left .props_container .prop_item:hover .img_container img {
    transform: scale(1.20);
  }
  .property_location_container .content .left .props_container .prop_item .img_container {
    min-width : 240px;
    width: 240px;
    height: 100%;
    overflow: hidden;
    border: 1px solid #cbd5e1;
    border-radius: 8px;
  }
  .property_location_container .content .left .props_container .prop_item .img_container img {
    object-fit: cover;
    width: 100%;
    height: 100%;
    transition-duration: 75ms;
  }
  .property_location_container .content .left .props_container .prop_item .img_container .image_empty {
    object-fit: cover;
    width: 100%;
    height: 100%;
    background-color: #f1f5f9;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 5px;
    transition-duration: 75ms;
  }
  .property_location_container .content .left .props_container .prop_item .prop_info {
    flex-grow: 1;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }
  .property_location_container .content .left .props_container .prop_item .prop_info .main_info {
    display: flex;
    flex-direction: column;
    gap: 16px;
  }
  .property_location_container .content .left .props_container .prop_item .prop_info .main_info .label_info {
    display: flex;
    flex-direction: row;
    gap: 8px;
  }
  .property_location_container .content .left .props_container .prop_item .prop_info .main_info .label_info .purpose,
  .property_location_container .content .left .props_container .prop_item .prop_info .main_info .label_info .house_type {
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
  .property_location_container .content .left .props_container .prop_item .prop_info .main_info .address {
    display: flex;
    flex-direction: row;
    align-items: flex-start;
    font-size: 15px;
    gap: 8px;
  }
  .property_location_container .content .left .props_container .prop_item .prop_info .main_info .feature {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    font-size: 13px;
  }
  .property_location_container .content .left .props_container .prop_item .prop_info .main_info .feature .item {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 8px;
  }
  .property_location_container .content .left .props_container .prop_item .prop_info .secondary_info {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: flex-end;
  }
  .property_location_container .content .left .props_container .prop_item .prop_info .secondary_info .size {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 6px;
  }
  .property_location_container .content .left .props_container .prop_item .prop_info .secondary_info .price {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    gap: 4px;
  }
  .property_location_container .content .left .props_container .prop_item .prop_info .secondary_info .price .agency_fee {
    font-size: 12px;
  }
  .property_location_container .content .left .props_container .prop_item .prop_info .secondary_info .price .last_price {
    font-size: 20px;
    font-weight: 700;
  }
  .property_location_container .content .right {
    border: 1px solid #cbd5e1;
    border-radius: 8px;
    height: 100%;
    position: relative;
  }
  .property_location_container .content .right .search_by_address {
    display: flex;
    flex-direction: column;
    height: 100%;
    width: 100%;
    gap: 20px;
  }
  .property_location_container .content .right .search_by_address .search_box {
    position: relative;
    width: 100%;
    height: fit-content;
    padding: 20px 20px 0 20px;
  }
  .property_location_container .content .right .search_by_address .search_box input {
      width: 100%;
      border: 1px solid #cbd5e1;
      background-color: #f1f5f9;
      border-radius: 8px;
      padding: 12px 12px 12px 40px;
    }
  .property_location_container .content .right .search_by_address .search_box input:focus {
      border-color: #3b82f6;
      outline: 1px solid #3b82f6;
    }
  :global(.icon_search_location) {
    position: absolute;
    left: 0;
    bottom: 0;
    top: 20px;
    z-index: 40;
    margin: auto 0 auto 32px;
  }
  .property_location_container .content .right .autocomplete_container {
    height: 100%;
    display: flex;
    flex-direction: column;
    padding: 0;
    overflow: auto;
    border-top: 1px solid #cbd5e1;
  }
  .property_location_container .content .right .autocomplete_container {
    height: 100%;
    display: flex;
    flex-direction: column;
    overflow: auto;
  }
  .property_location_container .content .right .autocomplete_container::-webkit-scrollbar-thumb {
    background-color : #3b82f6;
  }
  .property_location_container .content .right .autocomplete_container::-webkit-scrollbar {
    width: 8px;
  }
  .property_location_container .content .right .autocomplete_container::-webkit-scrollbar-track {
    background-color : transparent;
  }
  .property_location_container .content .right .autocomplete_container .empty {
    padding: 10px;
    text-align: center;
  }
  .property_location_container .content .right .autocomplete_container .autocomplete_item {
    display: flex;
    flex-direction: row;
    gap: 8px;
    align-items: center;
    padding: 10px;
    border: none;
    background: none;
    border-bottom: 1px solid #cbd5e1;
    cursor: pointer;
  }
  .property_location_container .content .right .autocomplete_container .autocomplete_item:hover {
    background-color: #f1f5f9;
  }
  .property_location_container .content .right .map_container {
    display: block;
    position: relative;
    width: 100%;
    height: 100%;
    border-radius: 8px;
    overflow: hidden;
  }
  .property_location_container .content .right .map_container .btn_sync_map {
    position: absolute;
    display: flex;
    flex-direction: row;
    gap: 8px;
    align-items: center;
    width: fit-content;
    padding: 0 20px;
    height: 40px;
    top: 10px;
    left: 188px;
    border-radius: 3px;
    font-size: 15px;
    background-color: #FFFFFF;
    border: none;
    z-index: 20;
    cursor: pointer;
    color: #475569;
    box-shadow: 0 4px 24px 0 rgba(0, 0, 0, 0.25);
  }
  .property_location_container .content .right .map_container .btn_sync_map:hover {
    background-color: #f1f5f9;
  }
</style>