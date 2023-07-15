<script>
   import Menu from './_components/Menu.svelte';
   import ProfileHeader from './_components/ProfileHeader.svelte';
   import Footer from './_components/Footer.svelte';

   let subPage = 1;
   let subpageRef;
   async function nextPage() {
      // const nextCard = document.getElementById(`subpage_${pageCount - 1}`);
      // nextCard.scrollIntoView({ behavior: 'smooth', block: 'end', inline: 'center' });
      if (subPage <= 4) {
         subPage++;
         const scrollOption = {
            left: 860,
            behavior: 'smooth'
         }
         subpageRef.scrollTo(scrollOption)
      }
   }
   
   let user = {/* user */};
   let segments = {/* segments */};
</script>

<section class='dashboard'>
   <Menu
      access={segments}
   />
   <div class='dashboard_main_content'>
      <ProfileHeader></ProfileHeader>
      <div class='realtor_step_progress_bar'>
         <div class='step_wrapper'>
            <div class='step_item completed'>
               <span></span>
               <p>Location</p>
            </div>
            <div class='step_item active'>
               <span></span>
               <p>Info</p>
            </div>
            <div class='step_item'>
               <span></span>
               <p>Floor</p>
            </div>
            <div class='step_item'>
               <span></span>
               <p>Preview</p>
            </div>
         </div>
      </div>
      <div class='content'>
         <div class='realtor_subpage_container' bind:this={subpageRef}>
            {#if subPage >= 1}
               <section class='location'>
                  <p>location</p>
                  <button>BACK</button>
                  <button on:click|preventDefault={nextPage}>NEXT</button>
               </section>
            {/if}
            {#if subPage >= 2}
               <section class='info'>
                  <p>info</p>
                  <button>BACK</button>
                  <button on:click|preventDefault={nextPage}>NEXT</button>
               </section>
            {/if}
            {#if subPage >= 3}
               <section class='floor'>
                  <p>floor</p>
                  <button>BACK</button>
                  <button on:click|preventDefault={nextPage}>NEXT</button>
               </section>
            {/if}
            {#if subPage >= 4}
               <section class='preview'>
                  <p>preview</p>
                  <button>BACK</button>
                  <button on:click|preventDefault={nextPage}>NEXT</button>
               </section>
            {/if}
         </div>
      </div>
      <Footer></Footer>
   </div>
</section>

<style>
   .realtor_step_progress_bar {
      position: relative;
      margin-top: -40px;
      margin-left: auto;
      margin-right: auto;
      display: flex;
      justify-content: center;
      background-color: #EF4444;
      padding-bottom: 70px;
   }
   .realtor_step_progress_bar .step_wrapper {
      width: 70%;
      margin-left: auto;
      margin-right: auto;
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      background-color: white;
      border-radius: 6px;
      filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
      padding: 19px 10% 11px 10%;
   }
   .realtor_step_progress_bar .step_wrapper .step_item {
      position: relative;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      flex: 1;
      font-size: 12px;
      color: #334155;
   }
   .realtor_step_progress_bar .step_wrapper .step_item::before {
      position: absolute;
      content: '';
      border-bottom: 2px solid #CBD5E1;
      width: 100%;
      top: 5px;
      left: -50%;
      z-index: 2;
   }
   .realtor_step_progress_bar .step_wrapper .step_item::after {
      position: absolute;
      content: '';
      border-bottom: 2px solid #CBD5E1;
      width: 100%;
      top: 5px;
      left: 50%;
      z-index: 2;
   }
   .realtor_step_progress_bar .step_wrapper .step_item:first-child::before {
      content: none;
   }
   .realtor_step_progress_bar .step_wrapper .step_item:last-child::after {
      content: none;
   }
   .realtor_step_progress_bar .step_wrapper .step_item span {
      width: 13px;
      height: 13px;
      background-color: #CBD5E1;
      border-radius: 100%;
      z-index: 4;
   }
   .realtor_step_progress_bar .step_wrapper .step_item p {
      margin: 8px 0 0 0;
   }
   .realtor_step_progress_bar .step_wrapper .step_item.active span {
      width: 11px;
      height: 11px;
      background-color: white;
      outline: 3px solid #f97316;
   }
   .step_item.completed::after {
      position: absolute !important;
      content: ''  !important;
      border-bottom: 2px solid #f97316  !important;
      width: 100%  !important;
      top: 5px  !important;
      left: 50%  !important;
      z-index: 3  !important;
   }
   .step_item.completed span {
      background-color: #f97316 !important;
   }

   .realtor_subpage_container {
      position: relative;
      margin-top: -40px;
      margin-left: auto;
      margin-right: auto;
      display: flex;
      border-radius: 8px;
      width: 860px;
      min-height: 500px;
      overflow-x: scroll;
      scroll-snap-type: x mandatory;
      filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
   }
   .realtor_subpage_container::-webkit-scrollbar-thumb {
      background-color: #EF4444;
      border-radius: 4px;
   }
   .realtor_subpage_container::-webkit-scrollbar {
      height: 0;
   }
   .realtor_subpage_container::-webkit-scrollbar-track {
      background-color: transparent;
      
   }
   .realtor_subpage_container section {
      padding: 20px;
      margin: 0 10px;
      background-color: white;
      border-radius: 8px;
      min-height: 700px;
      flex: 0 0 860px;
      scroll-snap-align: start;
   }
</style>