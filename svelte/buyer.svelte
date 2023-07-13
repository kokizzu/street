<script>
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
         buyer content goes here
      </div>
      <Footer></Footer>
   </div>
</section>