<script>
	/** @typedef {import('./_types/user').User} User */
  /** @typedef {import('./_types/property').Property} Property */

  import Main from './_layouts/Main.svelte';
  import { GoogleMap, GoogleSdk } from './_components/GoogleMap/components';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { LuSearch } from './node_modules/svelte-icons-pack/dist/lu';
  import { FaSolidBan, FaSolidImage } from './node_modules/svelte-icons-pack/dist/fa';
  import { T } from './_components/uiState';

	const user              = /** @type {User} */ ({/* user */});
  const coord             = /** @type {number[]} */ ([/* initialLatLong */]);
  const properties        = /** @type {Property[]} */ ([/* randomProps */]);

  let gmapsComponent = /** @type {import('svelte').SvelteComponent} */ (null);

  console.log('Properties = ', properties);
</script>

<GoogleSdk />

<Main {user} >
	<div class="listings-root">
    <div class="content">
      <div class="properties">
        {#if properties && properties.length}
          {#each properties as prop, _ (prop.id)}
            <div class="property">
              <picture class="img-container">
                {#if prop.images && prop.images.length}
                  <img
                    src={prop.images[0]}
                    alt={prop.formattedAddress}
                  />
                {:else}
                  <div class='image-empty'>
                    <Icon
                      size="40"
                      className="no-image-icon"
                      color="#848d96"
                      src={FaSolidImage}
                    />
                    <span>No Image !</span>
                  </div>
                {/if}
              </picture>
              <div class="prop_info">
                <p>Ingformation</p>
              </div>
            </div>
          {/each}
        {:else}
          <div class="no_properties">
            <div class="warn">
              <Icon size="17" color='#475569' src={FaSolidBan}/>
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
    grid-template-columns: auto 60%;
    position: relative;
  }

  .listings-root .content .properties {
    height: 4000px;
    padding: 20px;
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