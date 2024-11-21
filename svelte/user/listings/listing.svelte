<script>
  /** @typedef {import('../../_types/user').User} User */
  /** @typedef {import('../../_types/master').Access} Access */
  /** @typedef {import('../../_types/property').PropertyWithNote} PropertyWithNote */
  /** @typedef {import('../../_types/master').CountryData} CountryData */

  import Main from '../../_layouts/Main.svelte';
  import { Icon } from '../../node_modules/svelte-icons-pack/dist';
  import { LuSearch } from '../../node_modules/svelte-icons-pack/dist/lu';
  import GoogleMapJs from '../../_components/GoogleMap/GoogleMapJS.svelte';
  
  const title     = /** @type {string} */ ('#{title}')
  const user      = /** @type {User} */ ({/* user */});
  const access    = /** @type {Access} */ ({/* segments */});
  const property  = /** @type {PropertyWithNote} */ ({/* property */});
  const countries = /** @type {CountryData[]} */ ([/* countries */]);

  let gmapComponent = /** @type {import('svelte').SvelteComponent} */ (null);

  function gmapReady() {
    gmapComponent.SetMarker(
      property.coord[0],
      property.coord[1],
      title
    )
  }
</script>

<Main {user} {access}>
  <div class="listing-container">
    <div class="content">
      <div class="property">
        <div class="go-back">

        </div>
      </div>
      <div class="map-full">
        <div class="map-container">
          <GoogleMapJs
            bind:this={gmapComponent}
            lat={property.coord[0]}
            lng={property.coord[1]}
            zoom={10}
            on:ready={gmapReady}
          />
        </div>
      </div>
    </div>
  </div>
</Main>

<style>
  .listing-container {
    display: flex;
    flex-direction: column;
    width: 100%;
  }

  .listing-container div.content {
    display: grid;
    grid-template-columns: auto 55%;
    position: relative;
  }

  .listing-container .content .map-full {
    display: flex;
    flex-direction: column;
    position: sticky;
    top: 0;
    height: calc(100vh - var(--navbar-height));
  }

  .listing-container .content .map-full .map-container {
    display: block;
    width: 100%;
    height: 100%;
    overflow: hidden;
  }
</style>