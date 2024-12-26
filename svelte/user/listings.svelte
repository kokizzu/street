<script>
	/** @typedef {import('../_types/user').User} User */
  /** @typedef {import('../_types/property').PropertyWithNote} PropertyWithNote */
  /** @typedef {import('../_types/master').Access} Access }*/

  import Main from '../_layouts/Main.svelte';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { LuSearch } from '../node_modules/svelte-icons-pack/dist/lu';
  import { FaSolidBan } from '../node_modules/svelte-icons-pack/dist/fa';
  import PropertyImage from '../_components/propertyImage.svelte';
  import GoogleMapJs from '../_components/GoogleMap/GoogleMapJS.svelte';
  import axios, { HttpStatusCode } from 'axios';
  import { notifier } from '../_components/notifier';
  import PopUpUpload3DFile from '../_components/PopUpUpload3DFile.svelte';
  import { gmap, mapLoader } from '../_components/GoogleMap/stores';

	const user    = /** @type {User} */ ({/* user */});
  const access  = /** @type {Access} */ ({/* segments */});

  let coord         = /** @type {number[]} */ ([/* initialLatLong */]);
  let properties    = /** @type {PropertyWithNote[]} */ ([/* randomProps */]);
  let distanceKm    = /** @type {number} */ (Number('#{defaultDistanceKm}') || 20);
  let gmapComponent = /** @type {import('svelte').SvelteComponent} */ (null);
  let mapMarkers    = /** @type {google.maps.marker.AdvancedMarkerElement[]} */ ([]);

  let popUpUpload3DFile = /** @type {import('svelte').SvelteComponent} */ (null);
  let isSubmitUpload3d  = /** @type {boolean} */ (false);
  let upload3dProgress  = /** @type {string} */ ('');
  let upload3dPropId    = /** @type {number|string} */ (0);
  let upload3dCountry   = /** @type {string} */ ('');

  async function SubmitUpload3DFile(/** @type {File} */ file) {
    isSubmitUpload3d = true;

    const formData = new FormData();
    formData.append('propertyId', String(upload3dPropId));
    formData.append('country', String(upload3dCountry));
    formData.append('rawFile', file, file.name);

    await axios.postForm('/user/upload3DFile', formData, {
      headers: {
        "Content-Type": "multipart/form-data"
      },
      onUploadProgress: (/** @type {import('axios').AxiosProgressEvent} */ event) => {
        const propgressNum = Math.round((event.loaded * 100) / Number(event.total));
        upload3dProgress = `Uploading ${propgressNum}%`;
      }
    }).then((/** @type {import('axios').AxiosResponse}*/ res) => {
      notifier.showSuccess('3D file uploaded !!');
      popUpUpload3DFile.Close();

      for (let i = 0; i < (properties.length -1); i++) {
        if (properties[i].id === upload3dPropId) {
          properties[i].image3dUrl = res.data.imageUrl;
          console.log('Img 3D URL:', properties[i].image3dUrl);
          break;
        } 
      }
    }).catch((/** @type {import('axios').AxiosError}*/ err) => {
      const res = /** @type {import('axios').AxiosResponse} */ (err.response);
      switch (res.status) {
        case HttpStatusCode.PayloadTooLarge:
          notifier.showError('File size too large');
          break;
        default:
          notifier.showError(res.data.error || 'Failed to upload 3D file');
      }
      popUpUpload3DFile.Reset();
    });
  }

  let inputSearchAddressValue   = /** @type {string} */ ('');
  let isFocusInputSearchAddress = /** @type {boolean} */ (false);
  let isSearchingAddress        = /** @type {boolean} */ (false);
  let placeSuggestions          = /** @type {google.maps.places.AutocompleteSuggestion[]} */ ([]);
  let autoCompleteSuggestion    = /** @type {typeof google.maps.places.AutocompleteSuggestion} */ (null);
  let autoCompleteSessToken     = /** @type {typeof google.maps.places.AutocompleteSessionToken} */ (null);
  let geocoder                  = /** @type {google.maps.Geocoder} */ (null);

  async function HandlerSearchLocation() {
    const latLngFunc = /** @type {google.maps.LatLng} */ ($gmap.getCenter());
    const request = /** @type {google.maps.places.AutocompleteRequest} */ ({
      sessionToken: new autoCompleteSessToken(),
      input: inputSearchAddressValue,
      origin: {
        lat: latLngFunc.lat(),
        lng: latLngFunc.lng()
      }
    });

    const { suggestions } = await autoCompleteSuggestion.fetchAutocompleteSuggestions(request);

    placeSuggestions = suggestions;
    console.log('placeSuggestions=', placeSuggestions);
  }

  async function HandlerSearchByPlaceID(/** @type {string} */ placeId) {
    await geocoder.geocode({ placeId }).then(({ results }) => {
      if (results.length > 0) {
        const res = results[0];
        $gmap.setCenter(res.geometry.location);
      }
    });

    isFocusInputSearchAddress = false;
    inputSearchAddressValue = '';

    
  }

  async function gmapReady() {
    (properties || []).forEach((p) => {
      mapMarkers = [...mapMarkers,
        gmapComponent.SetMarker(p.coord[0], p.coord[1], p.address)
      ];
    });

    geocoder = new google.maps.Geocoder();

    $mapLoader.importLibrary('places').then(({ AutocompleteSessionToken, AutocompleteSuggestion }) => {
      autoCompleteSuggestion = AutocompleteSuggestion;
      autoCompleteSessToken = AutocompleteSessionToken;
    }).catch((err) => {
      console.error('failed to load AutocompleteSuggestion :', err);
    });
  }
</script>

<PopUpUpload3DFile
  bind:this={popUpUpload3DFile}
  bind:isSubmitted={isSubmitUpload3d}
  bind:uploadingProgressStr={upload3dProgress}
  OnSubmit={SubmitUpload3DFile}
/>

<Main {user} {access}>
	<div class="listings-root">
    <div class="content">
      <div class="properties">
        {#each (properties || []) as prop, _ (prop.id)}
          <div class="property">
            <PropertyImage
              src={prop.images[0]}
              alt={prop.address || prop.formattedAddress}
            />
            <div class="info">
              <div class="top">
                <b>{prop.address || prop.formattedAddress}</b>
                <div class="features">
                  <div class="pill-box">Floors: {prop.numberOfFloors}</div>
                  <div class="pill-box">Bed: {prop.bedroom}</div>
                  <div class="pill-box">Bath: {prop.bathroom}</div>
                  <div class="pill-box">Size (m2): {prop.sizeM2}</div>
                  <div class="pill-box">Siz (sqft): {prop.totalSqft}</div>
                </div>
              </div>
              <div class="actions">
                <button on:click={() => window.location.href = `/user/listings/${prop.id}`}>Listing</button>
                {#if prop.image3dUrl === ''}
                  <button on:click={() => {
                    upload3dCountry = prop.countryCode;
                    upload3dPropId = prop.id;
                    popUpUpload3DFile.Show();
                  }}>Upload 3D file</button>
                {/if}
                {#if prop.image3dUrl !== ''}
                  <button class="download-btn" on:click={() => {
                    window.open(
                      `/guest/download3dFile?country=${prop.countryCode}&propertyId=${prop.id}`,
                      '_blank'
                    );
                  }}>
                    Download 3D file
                  </button>
                {/if}
              </div>
            </div>
          </div>
        {/each}
        {#if !properties || properties.length === 0}
          <div class="no-properties">
            <div class="warn">
              <Icon
                size="17"
                color="#475569"
                src={FaSolidBan}
              />
              <span>No properties in this area</span>
            </div>
          </div>
        {/if}
      </div>
      <div class="searcher">
        <div class="search-container">
          <div class="search-form">
            <div class="search-box">
              <Icon
                src={LuSearch}
                size="15"
                className="search-icon"
              />
              <input
                type="text"
                id="search-address"
                bind:value={inputSearchAddressValue}
                on:input={HandlerSearchLocation}
                on:focus={() => isFocusInputSearchAddress = true}
                on:blur={() => isFocusInputSearchAddress = false}
                placeholder="Search for address..."
              />
            </div>
            <button class="search-btn">
              <span>Search</span>
            </button>
          </div>
        </div>
        <div class="map-container">
          <GoogleMapJs
            bind:this={gmapComponent}
            lat={coord[0]}
            lng={coord[1]}
            zoom={11}
            on:ready={gmapReady}
          />
        </div>
      </div>
    </div>
	</div>
</Main>

<style>
  .listings-root {
    display: flex;
    flex-direction: column;
    width: 100%;
  }

  .listings-root div.content {
    display: grid;
    grid-template-columns: auto 55%;
    position: relative;
  }

  .listings-root .content .properties {
    height: fit-content;
    width: 100%;
    padding: 15px 20px;
    display: flex;
    flex-direction: column;
    gap: 15px;
  }

  .listings-root .content .properties .property {
    display: grid;
    grid-template-columns: 35% auto;
    border-radius: 10px;
    border: 1px solid var(--gray-002);
    overflow: hidden;
  }

  

  .listings-root .content .properties .property .info {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    gap: 10px;
    padding: 10px;
  }

  .listings-root .content .properties .property .info .top {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .listings-root .content .properties .property .info .top .features {
    display: flex;
    flex-direction: row;
    gap: 5px;
    flex-wrap: wrap;
  }

  .listings-root .content .properties .property .info .top .features .pill-box {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 5px 10px;
    border-radius: 8px;
    border: 1px solid var(--gray-002);
    font-size: var(--font-sm);
  }

  .listings-root .content .properties .property .info .actions {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
  }

  .listings-root .content .properties .property .info .actions button {
    background-color: transparent;
    border: none;
    font-weight: 600;
    color: var(--orange-006);
    cursor: pointer;
  }

  .listings-root .content .properties .property .info .actions button:hover {
    color: var(--orange-005);
  }

  .listings-root .content .properties .property .info .actions button.download-btn {
    color: var(--blue-006);
  }

  .listings-root .content .properties .property .info .actions button.download-btn:hover {
    color: var(--blue-005);
  }

  .listings-root .content .searcher {
    display: flex;
    flex-direction: column;
    position: sticky;
    top: 0;
    height: calc(100vh - var(--navbar-height));
  }

  .listings-root .content .searcher .search-container {
    display: flex;
    flex-direction: row;
    padding: 15px 0;
  }

  .listings-root .content .searcher .search-container .search-form {
    display: flex;
    flex-direction: row;
    width: 100%;
    gap: 10px;
  }

  .listings-root .content .searcher .search-container .search-form .search-box {
    display: flex;
    flex-direction: row;
    align-items: center;
    position: relative;
    width: 60%;
  }

  .listings-root .content .searcher .search-container .search-form .search-box input {
    width: 100%;
    border: 1px solid var(--gray-003);
    border-radius: 8px;
    background-color: transparent;
    padding: 12px 12px 12px 40px;
    width: 100%;
  }

  .listings-root .content .searcher .search-container .search-form .search-box input:focus {
    border-color: var(--orange-005);
    outline: 1px solid var(--orange-005);
  }

  :global(.listings-root .content .searcher .search-container .search-form .search-box .search-icon) {
    position: absolute;
    left: 15px;
    top: 50%;
    transform: translateY(-50%);
    color: var(--gray-003);
  }

  .listings-root .content .searcher .search-container .search-form button.search-btn {
    display: flex;
    flex-direction: row;
    align-items: center;
    padding: 10px 30px;
    background-color: var(--orange-006);
    color: #FFF;
    border-radius: 8px;
    font-weight: 600;
    border: none;
    cursor: pointer;
  }

  .listings-root .content .searcher .search-container .search-form button.search-btn:hover {
    background-color: var(--orange-005);
  }

  .listings-root .content .searcher .map-container {
    display: block;
    width: 100%;
    height: 100%;
    overflow: hidden;
  }
</style>