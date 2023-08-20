<script>
  // @ts-nocheck
  import { onMount } from 'svelte';

  import Menu from './_components/Menu.svelte';
  import ProfileHeader from './_components/ProfileHeader.svelte';
  import Footer from './_components/Footer.svelte';
  import { RealtorUpsertProperty } from './jsApi.GEN';


  let user = {
    /* user */
  };
  let segments = {
    /* segments */
  };
</script>
  
<section class="dashboard">
  <Menu access={segments} />
  <div class="dashboard_main_content">
    <ProfileHeader />
    <div class="content">
      <p>Realtor</p>
    </div>
    <Footer />
  </div>
</section>