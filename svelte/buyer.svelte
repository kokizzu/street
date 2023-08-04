<script>
  // @ts-nocheck
  import Menu from './_components/Menu.svelte';
  import ProfileHeader from './_components/ProfileHeader.svelte';
  import Footer from './_components/Footer.svelte';
  
  let user = {/* user */};
  let segments = {/* segments */};
</script>

<section class="dashboard">
   <Menu
      access={segments}
   />
   <div class="dashboard_main_content">
      <ProfileHeader></ProfileHeader>
      <div class="content">
        <div id="street_view" class="street_view">
          <iframe
            title="Street view"
            src="https://www.google.com/maps/embed?pb=!4v1691159921635!6m8!1m7!1sCAoSLEFGMVFpcE9UZ0pyRHBhSjZ5bGhha2QtRHBGNzZ2cDRBZk82alZwMFhjUndM!2m2!1d-33.8562164!2d151.2153473!3f60.305466!4f15.305465999999996!5f0.7820865974627469"
            allowfullscreen="true" loading="lazy" referrerpolicy="no-referrer-when-downgrade"
          ></iframe>
        </div>
      </div>
      <Footer></Footer>
   </div>
</section>

<style>
  .street_view {
    width: 70%;
    position: relative;
    margin-top: -40px;
    margin-left: auto;
    margin-right: auto;
    border-radius: 8px;
  }
  .street_view iframe {
    width: 100%;
    height: 450px;
    border: none;
    border-radius: 8px;
  }
</style>