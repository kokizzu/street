<script>
  import Menu from './_components/Menu.svelte';
  import ProfileHeader from './_components/ProfileHeader.svelte';
  import Footer from './_components/Footer.svelte';

  let subPage = 1;
  let subpageRef;
   
  async function nextPage() {
    if (subPage < 4) {
      await subPage++;
      const nextCard = document.getElementById(`subpage_${subPage}`);
      nextCard.scrollIntoView({ behavior: 'smooth', block: 'end', inline: 'center' });
    }
  }
  function backPage() {
    if (subPage > 1) {
      subPage--;
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
        <div class={subPage > 1 ? 'step_item completed' : 'step_item active'}>
          <span></span>
          <p>Location</p>
        </div>
        <div class={subPage == 2 ? 'step_item active' : 'step_item' && subPage > 2 ? 'step_item completed' : 'step_item'}>
          <span></span>
          <p>Info</p>
        </div>
        <div class={subPage == 3 ? 'step_item active' : 'step_item' && subPage > 2 ? 'step_item completed' : 'step_item'}>
          <span></span>
          <p>Floor</p>
        </div>
        <div class={subPage == 4 ? 'step_item active' : 'step_item'}>
          <span></span>
          <p>Preview</p>
        </div>
      </div>
    </div>
    <div class='content'>
      <div class='realtor_subpage_container' bind:this={subpageRef}>
        {#if subPage >= 1}
          <section class='location' id='subpage_1'>
            <div>
              <h2>Type property's address or move the map to autofill it</h2>               
              <div class='location_input'>
                <div class='input_box'>
                  <label for="input_address"></label>
                  <input type="text" id='input_address' />
                </div>
                <div class='map_container'>
                  <img src='/assets/img/realtor/GoogleMapTA.webp' alt='Map'>
                </div>
              </div>
            </div>               
            <button class='next_button' on:click|preventDefault={nextPage}>NEXT</button>
          </section>
        {/if}
        {#if subPage >= 2}
          <section class='info' id='subpage_2'>
            <p>info</p>
            <button on:click|preventDefault={backPage}>BACK</button>
            <button on:click|preventDefault={nextPage}>NEXT</button>
          </section>
        {/if}
        {#if subPage >= 3}
          <section class='floor' id='subpage_3'>
            <p>floor</p>
            <button on:click|preventDefault={backPage}>BACK</button>
            <button on:click|preventDefault={nextPage}>NEXT</button>
          </section>
        {/if}
        {#if subPage >= 4}
          <section class='preview' id='subpage_4'>
            <p>preview</p>
            <button on:click|preventDefault={backPage}>BACK</button>
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

  /* +============| SUBPAGE LOCATION |===========+ */
  .location {
    color: #334155;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    height: 100%;
  }
  .location .next_button {
    background-color: #f97316;
    border-radius: 8px;
    border: none;
    padding: 10px;
    color: white;
    font-weight: 600;
    margin-top: 20px;
    cursor: pointer;
  }
  .location .next_button:hover {
    background-color: #f58433
  }
  .location h2 {
    font-size: 18px;
    font-weight: 600;
    margin-bottom: 20px;
  }
  .location_input {
    display: flex;
    flex-direction: column;
  }
  .location_input .input_box input {
    width: 100%;
    border: 1px solid #CBD5E1;
    background-color: #F1F5F9;
   	border-radius: 8px;
   	padding: 10px 12px;
  }
  .location_input .input_box input:focus {
    border-color: #f97316;
    outline: 1px solid #f97316;
  }

  .location_input .map_container {
    margin-top: 20px;
    border-radius: 8px;
    width: 100%;
    height: fit-content;
  }
  .location_input .map_container img {
    width: 100%;
    height: auto;
    border-radius: 8px;
  }
</style>