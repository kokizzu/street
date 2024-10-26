<script>
  /** @typedef {import('../_types/master').Access} Access */
  /** @typedef {import('../_types/property').Property} TypeProperty */
  /** @typedef {import('../_types/master').CountryData} CountryData */
  /** @typedef {import('../_types/user').User} User */

  import Menu from '../_components/Menu.svelte';
  import ProfileHeader from '../_components/ProfileHeader.svelte';
  import Footer from '../_components/partials/Footer.svelte';
  import CreateProperty from "../_components/CreateProperty.svelte";
  import EditProperty from "../_components/EditProperty.svelte";
  import { onMount } from 'svelte';
  import { GoogleSdk } from '../_components/GoogleMap/components';
  
  let property  = /** @type {TypeProperty} */ ({/* property */});
  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access}*/ ({/* segments */});
  let countries = /** @type {CountryData[]}*/ ([/* countries */]);

  let isAdmin = false;
  let isOwner = false;

  onMount(() => {
    console.log('User = ', user);
    console.log('Segments = ', segments);
    if (segments.admin === true) isAdmin = true;
    if (user.id == property.createdBy) isOwner = true;
  });

  let isGoogleReady = false;
</script>

<GoogleSdk on:ready={()=>isGoogleReady = true}/>

<section class='dashboard'>
  <Menu access={segments}/>
  <div class='dashboard_main_content'>
    <ProfileHeader {user} access={segments}/>
    {#if isGoogleReady}
      {#if property.id}
        <EditProperty {property} {countries} {isAdmin} {isOwner} />
      {:else}
        <CreateProperty {property} {user} {countries}/>
      {/if}
    {/if}
    <Footer></Footer>
  </div>
</section>