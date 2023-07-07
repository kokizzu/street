<script>
   import Menu from './_components/Menu.svelte';
   import ProfileHeader from './_components/ProfileHeader.svelte';
   import Footer from './_components/Footer.svelte';

   let sideMenuOpen = false;
   function openSideMenu() {
      sideMenuOpen = true;
   }
   function closeSideMenu() {
      sideMenuOpen = false;
   }
  
   let user = {/* user */};
   let segments = {/* segments */};
</script>

<section class="dashboard">
   <Menu
      access={segments}
      isSideMenuOpen={sideMenuOpen}
      on:closesidemenu={closeSideMenu}
   />
   <div class="dashboard_main_content">
      <ProfileHeader on:opensidemenu={openSideMenu}></ProfileHeader>
      <div class="content">
         buyer content goes here
      </div>
      <Footer></Footer>
   </div>
</section>