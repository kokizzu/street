<script>
	/** @typedef {import('./_types/user').User} User */
  /** @typedef {import('./_types/property').Property} Property */
  /**
   * @typedef {Object} MarkerIcon
   * @property {string} path
   * @property {string} alt
   */

  import Main from './_layouts/Main.svelte';
  import { GoogleMap, GoogleSdk } from './_components/GoogleMap/components';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { LuSearch } from './node_modules/svelte-icons-pack/dist/lu';
  import { FaSolidBan } from './node_modules/svelte-icons-pack/dist/fa';
  import PropertyImage from './_components/propertyImage.svelte';
  import { onMount } from 'svelte';

	const user              = /** @type {User} */ ({/* user */});
  const coord             = /** @type {number[]} */ ([/* initialLatLong */]);
  const properties        = /** @type {Property[]} */ ([/* randomProps */]);
  const defaultDistanceKm = /** @type {number} */ (Number('#{defaultDistanceKm}') || 20);

  let gmapsComponent = /** @type {import('svelte').SvelteComponent} */ (null);

  const markerIcons = /** @type {Record<string, MarkerIcon>} */ ({
    'school': {
      path: '/assets/icons/marker-school.svg',
      alt: 'School',
    },
    'restaurant': {
      path: '/assets/icons/marker-restaurant.svg',
      alt: 'Restaurant',
    },
    'convenience_store': {
      path: '/assets/icons/marker-mall.svg',
      alt: 'Convenience Store',
    },
    'hospital': {
      path: '/assets/icons/marker-hospital.svg',
      alt: 'Hospital',
    },
    'subway_station': {
      path: '/assets/icons/marker-subway.svg',
      alt: 'Subway Station',
    }
  });

  onMount(() => {
    console.log('user = ', user);
    console.log('coord = ', coord);
    console.log('properties = ', properties);
    console.log('markerIcons = ', markerIcons);
    console.log('defaultDistanceKm = ', defaultDistanceKm);
  })
</script>

<GoogleSdk />

<Main {user} >
	<div class="listings-root">
    <div class="content">
      <div class="properties">
        {#if properties && properties.length}
          {#each properties as prop, _ (prop.id)}
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
                  <button>Listing</button>
                  <button>Upload 3D file</button>
                </div>
              </div>
            </div>
          {/each}
        {:else}
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
              <input type="text" />
            </div>
            <button class="search-btn">
              <span>Search</span>
            </button>
          </div>
        </div>
        <div class="map-container">
          <GoogleMap
            bind:this={gmapsComponent}
            options={{
              center: {
                lat: coord[0],
                lng: coord[1],
              },
              zoom: 8,
              mapTypeId: 'roadmap',
              mapId: 'street_project',
            }}
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