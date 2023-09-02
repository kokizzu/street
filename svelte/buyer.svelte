<script>
  // @ts-nocheck
  import Menu from './_components/Menu.svelte';
  import ProfileHeader from './_components/ProfileHeader.svelte';
  import Footer from './_components/Footer.svelte';
  import { GoogleMap } from './_components/GoogleMap/components';
  
  let user = {/* user */};
  let segments = {/* segments */};
  
  let myLatLng = {lat: 23.6978, lng: 120.9605};
  let mapOptions = {
     center: myLatLng,
     zoom: 8,
     mapTypeId: 'roadmap',
     mapId: 'street_project',
  }
</script>

<section class="dashboard">
   <Menu
      access={segments}
   />
   <div class="dashboard_main_content">
      <ProfileHeader></ProfileHeader>
      <div class="content">
        <div class='map_container'>
          <GoogleMap options={mapOptions} />
        </div>
      </div>
      <Footer></Footer>
   </div>
</section>

<style>
  .map_container {
    width: 70%;
    height        : 430px;
    position: relative;
    margin-top: -40px;
    margin-left: auto;
    margin-right: auto;
    border-radius: 8px;
  }
</style>