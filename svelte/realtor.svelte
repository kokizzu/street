<script>
   import Menu from './_components/Menu.svelte';
   import ProfileHeader from './_components/ProfileHeader.svelte';
   import Footer from './_components/Footer.svelte';
   import { stackPageRealtor } from '_components/uiState';

   let newStack = [];
   let arrLength = 0;
   let newElement = {
      subroute: 'addNewPropStep',
      attrs: {
         prop1: 'hehe',
         prop2: 'haha'
      }
   }
   stackPageRealtor.subscribe((value) => {
      newStack = value;
      arrLength = value.length;
   });
   function nextPage() {
      if (arrLength < 4)
      stackPageRealtor.update(arr => [...arr, newElement] );
   }
   function backPage() {
      if (arrLength > 1)
      stackPageRealtor.update(arr => {
         arr.pop();
         return arr;
      });
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
         <div class='realtor_subpage_container'>
               {#each newStack as subpage}
                  <div class='subpage'>
                     <h4>{subpage.subroute}</h4>
                     <button on:click|preventDefault={nextPage}>NEXt</button>
                     <button on:click|preventDefault={backPage}>BACk</button>
                  </div>
               {/each}
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
      content: "";
      border-bottom: 2px solid #CBD5E1;
      width: 100%;
      top: 5px;
      left: -50%;
      z-index: 2;
   }
   .realtor_step_progress_bar .step_wrapper .step_item::after {
      position: absolute;
      content: "";
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
      content: ""  !important;
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
      flex-direction: row;
      justify-content: center;
      flex-wrap: nowrap;
      background-color: white;
      border-radius: 8px;
      filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
      width: 88%;
      padding: 20px;
   }
   .realtor_subpage_container .subpage {
      margin: 0 30px;
      padding: 20px;
      filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
      background-color: white;
      border-radius: 8px;
   }
</style>