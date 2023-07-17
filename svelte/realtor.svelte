<script>
  import Menu from './_components/Menu.svelte';
  import ProfileHeader from './_components/ProfileHeader.svelte';
  import Footer from './_components/Footer.svelte';
  import AddFloorDialog from '_components/AddFloorDialog.svelte';

  let user = {/* user */};
  let segments = {/* segments */};

  let subPage = 1;
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
  async function skipPage() {
    if (subPage < 4) {
      await subPage++;
      const nextCard = document.getElementById(`subpage_${subPage}`);
      nextCard.scrollIntoView({ behavior: 'smooth', block: 'end', inline: 'center' });
    }
  }

  // +=============| Location |=============+ //

  // +=============| House Info |=============+ //
  let modeHouseInfoCount = 0;
  let modeSkippable = false;
  const ABOUT_THE_HOUSE = 'About The House';
  const UPLOAD_HOUSE_PHOTO = 'Upload House Photo';
  const FEATURE_OR_FACILITY = 'Feature or Facility';
  const DESCRIPTION_PROPERTY = 'Description of this property';
  const PRICE = 'Price';
  const modeHouseLists = [
    { mode: ABOUT_THE_HOUSE, skip: false },
    { mode: UPLOAD_HOUSE_PHOTO, skip: true },
    { mode: FEATURE_OR_FACILITY, skip: true },
    { mode: DESCRIPTION_PROPERTY, skip: true },
    { mode: PRICE, skip: false }
  ];
  let mode = modeHouseLists[modeHouseInfoCount].mode;
  function houseInfoNext() {
    if (modeHouseInfoCount < modeHouseLists.length - 1) {
      modeHouseInfoCount++;
      mode = modeHouseLists[modeHouseInfoCount].mode;
      modeSkippable = modeHouseLists[modeHouseInfoCount].skip;
    } else {
      nextPage();
    }
  }
  function houseInfoBack() {
    if (modeHouseInfoCount > 0) {
      modeHouseInfoCount--;
      mode = modeHouseLists[modeHouseInfoCount].mode;
      modeSkippable = modeHouseLists[modeHouseInfoCount].skip;
    } else {
      backPage();
    }
  }
  function houseInfoSkip() {
    if (modeHouseInfoCount < modeHouseLists.length - 1) {
      modeHouseInfoCount++;
      mode = modeHouseLists[modeHouseInfoCount].mode;
      modeSkippable = modeHouseLists[modeHouseInfoCount].skip;
    } else {
      nextPage();
    }
  }
  // ______About The House
  // **** House Type
  let house_type = '';
  let showHouseOption = false;
  let isApartment = false;
  function handleHouseTypeOption() {
    isApartment = house_type === 'apartment';
    showHouseOption = !showHouseOption;
  }
  // **** Rent or Sell
  let rent_or_sell = '';
  let showRentSellOption = false;
  function handleRentSellOption() {
    showRentSellOption = !showRentSellOption;
  }
  // ______Upload House Photo
  let imageCount = 0;
  let imageInput;
  let images = [/*{ image: data, preview }*/];
  let showImage = false;
  function inputImageHandler() {
    const file = imageInput.files[0];
    if (file) {
      showImage = true;
      const reader = new FileReader();
      reader.addEventListener('load', function() {
        images = [...images, {
          image: file,
          preview: reader.result
        }]
      });
      reader.readAsDataURL(file);
      imageCount++
      return;
    }
    showImage = false;
  }
  function removeImage(index) {
    images = images.filter((_, i) => i !== index);
  }

  // +=============| Floors |=============+ //
  let add_floor_dialog = AddFloorDialog;
  let floor_type = '';
  let floor_lists = [];
  let floor_attribute = {
    type: '',
    floor: 0,
    beds: 0,
    baths: 0
  }
  let floorCount = 1;
  let basement_added = false;
  let floor_edit_mode = false;
  let floor_index_to_edit = 0;

  function showAddFloorDialog() {
    add_floor_dialog.showModal();
  }
  function handlerAddFloor() {
    if (floor_type === 'basement' && basement_added === true) {
      alert('basement already added')
      add_floor_dialog.hideModal()
      return
    }
    if (floor_type === 'basement') {
      floor_attribute = {
        type: floor_type,
        floor: 0,
        beds: 0,
        baths: 0
      }
      floor_lists = [...floor_lists, floor_attribute]
      basement_added = true
    } else {
      floor_attribute = {
        type: floor_type,
        floor: floorCount,
        beds: 0,
        baths: 0
      }
      floor_lists = [...floor_lists, floor_attribute]
      floorCount++
    }
    add_floor_dialog.hideModal();
    return
  }
  function handlerEditFloor(index) {
    floor_edit_mode = true;
    floor_index_to_edit = index;
  }
  // _______Upload Floor Plan photo
  // let imageCount = 0;
  // let imageInput;
  // let images = [/*{ image: data, preview }*/];
  // let showImage = false;
  // function inputImageHandler() {
  //   const file = imageInput.files[0];
  //   if (file) {
  //     showImage = true;
  //     const reader = new FileReader();
  //     reader.addEventListener('load', function() {
  //       images = [...images, {
  //         image: file,
  //         preview: reader.result
  //       }]
  //     });
  //     reader.readAsDataURL(file);
  //     imageCount++
  //     return;
  //   }
  //   showImage = false;
  // }
</script>

<section class='dashboard'>
  <Menu
    access ={segments}
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
      <div class='realtor_subpage_container'>
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
                  <iframe
                    title='location'
                    src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d7890.4948540156465!2d116.0808121910197!3d-8.572191565888081!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2dcdc0619c44f171%3A0x3a8c129c834e7cc2!2sPejeruk%2C%20Ampenan%2C%20Mataram%20City%2C%20West%20Nusa%20Tenggara!5e0!3m2!1sen!2sid!4v1689595526347!5m2!1sen!2sid"
                    loading="lazy"
                    referrerpolicy="no-referrer-when-downgrade">
                  </iframe>
                </div>
              </div>
            </div>               
            <button class='next_button' on:click|preventDefault={nextPage}>NEXT</button>
          </section>
        {/if}
        {#if subPage >= 2}
          <section class='info' id='subpage_2'>
            <div class='main_info'>
              <button class='back_button' on:click|preventDefault={houseInfoBack}>
                <i class='gg-chevron-left' />
              </button>
              {#if mode === ABOUT_THE_HOUSE}
                <h2>{mode}</h2>
                <div class='about_the_house'>
                  <div class='house_type_container'>
                    <div class='house_type'>
                      <label for='house_type'>House Type</label>
                      <button on:click={handleHouseTypeOption}>{house_type ? house_type : 'select'}</button>
                      {#if showHouseOption}
                        <div class='option_container'>
                          <label class='option' for='house'>
                            <input
                              type='radio'
                              bind:group={house_type}
                              on:change={() => handleHouseTypeOption()}
                              id='house'
                              value='house'
                            />
                            House
                          </label>
                          <label class='option' for='apartment'>
                            <input
                              type='radio'
                              bind:group={house_type}
                              on:change={() => handleHouseTypeOption()}
                              id='apartment'
                              value='apartment'
                            />
                            Apartment
                          </label>
                        </div>
                      {/if}
                    </div>
                    {#if isApartment}
                      <div class='apartment_floor'>
                        <label for='apartment_floor'>Floor</label>
                        <input
                          type='number'
                          name='apartment_floor'
                          id='apartment_floor'
                          placeholder='Your property based'
                        />
                      </div>
                    {/if}
                  </div>
                  <div class='rent_or_sell'>
                    <label for='rent_or_sell'>Rent or Sell</label>
                    <button disabled={!house_type} on:click={handleRentSellOption}
                      >{rent_or_sell ? rent_or_sell : 'select'}</button
                    >
                    {#if showRentSellOption}
                      <div class='option_container'>
                        <label class='option' for='rent'>
                          <input
                            type='radio'
                            bind:group={rent_or_sell}
                            on:change={() => handleRentSellOption()}
                            id='rent'
                            value='rent'
                          />
                          Rent
                        </label>
          
                        <label class='option' for='sell'>
                          <input
                            type='radio'
                            bind:group={rent_or_sell}
                            on:change={() => handleRentSellOption()}
                            id='sell'
                            value='sell'
                          />
                          Sell
                        </label>
                      </div>
                    {/if}
                  </div>
                </div>
              {/if}
              {#if mode === UPLOAD_HOUSE_PHOTO}
                <h2>{mode}</h2>
                <div class='upload_house_photo'>
                  
                  <div class='image_preview_container'>
                    <label class='image_upload_button' for='upload_image'>
                      <input
                        bind:this={imageInput}
                        on:change={inputImageHandler}
                        type='file'
                        accept='image/*'
                        id='upload_image'
                      />
                      <i class='gg-software-upload'></i>
                      <p>Select file to Upload</p>
                    </label>
                    {#if showImage}
                      {#each images as imgFile, index}
                        <div class='image_card'>
                          <img src={imgFile.preview} alt=''>
                          <button on:click={() => removeImage(index)} title='remove this image'>
                            <i class='gg-close'></i>
                          </button>
                        </div>
                      {/each}
                    {/if}
                  </div>
                </div>
              {/if}
              {#if mode === FEATURE_OR_FACILITY}
                <h2>Feature</h2>
                <div class='feature_section'>
                  <div class='beds'>
                    <input type='number' min='0' name='beds' id='beds'>
                    <label for='beds'>Beds</label>
                  </div>
                  <div class='baths'>
                    <input type='number' min='0' name='baths' id='baths'>
                    <label for='baths'>Baths</label>
                  </div>
                  <div class='square_foot'>
                    <input type='number' min='0' name='square_foot' id='square_foot' step='0.01'>
                    <label for='square_foot'>Sq Ft</label>
                  </div>
                </div>
                <h2>Facility</h2>
                <div class='facility_section'>
                  <textarea rows='10' placeholder='Type the facility in the property.' name='facility' id='facility'></textarea>
                </div>
              {/if}
              {#if mode === DESCRIPTION_PROPERTY}
                <h2>{mode}</h2>
                <div class='description_of_property'>
                  <textarea rows='20' placeholder='Writing a description can help potential buyers become more interested in your property.' name='description' id='description'></textarea>
                </div>
              {/if}
              {#if mode === PRICE}
                <h2>{mode}</h2>
                <div class='price'>
                  <div class='property_price'>
                    <label for='property_price'>Property Price</label>
                    <input type='number' min='0' name='property_price' id='property_price' />
                  </div>
                  <div class='agency_fee'>
                    <label for='agency_fee'>Agency Fee</label>
                    <div>
                      <input type='number' min='0' name='agency_fee' id='agency_fee'>
                      <span>%</span>
                    </div>
                  </div>
                </div>
              {/if}
            </div>
          
            <div class='next_skip_button'>
              {#if modeSkippable === true}
                <button class='skip_button' on:click|preventDefault={houseInfoSkip}>
                  SKIP
                </button>
              {/if}
              <button class='next_button' on:click|preventDefault={houseInfoNext}>
                NEXT
              </button>
            </div>
          </section>
        {/if}
        {#if subPage >= 3}
          <section class='floor' id='subpage_3'>
            <AddFloorDialog bind:this={add_floor_dialog} bind:floor_type={floor_type}>
              <button disabled={floor_type === ''} class='add_floor_button' on:click={handlerAddFloor}>Add</button>
            </AddFloorDialog>
            <div class='floor_main_content'>
              <button
                class='back_button'
                on:click|preventDefault={() => {
                  if (floor_edit_mode === true) {
                    floor_edit_mode = false;
                  } else {
                    backPage();
                  }
                }}>
                <i class='gg-chevron-left' />
              </button>
              <div class='floor_header'>
                {#if floor_edit_mode === true}
                  <h2>Edit {floor_lists[floor_index_to_edit].type} {floor_lists[floor_index_to_edit].type === 'basement' ? '' : `#${floor_lists[floor_index_to_edit].floor}` }</h2>
                {:else}
                  <h2>Floors</h2>
                  <button on:click|preventDefault={showAddFloorDialog}>Add</button>
                {/if}
              </div>
              {#if floor_edit_mode === false}
              <div class='floor_items_container'>
                {#each floor_lists as floor, index}
                  <div class='floor_item'>
                    <div class='left_item'>
                      <h4>{floor.type} {floor.type === 'basement' ? '' : `#${floor.floor}` }</h4>
                      <div class='rooms_total'>
                        <div class='beds'>
                          <b>3</b>
                          <p>Beds</p>
                        </div>
                        <div class='baths'>
                          <b>2</b>
                          <p>Baths</p>
                        </div>
                      </div>
                    </div>
                    <div class='right_item'>
                      <button class='edit_floor' on:click={() => handlerEditFloor(index)}>Edit</button>
                      <div class='floor_plan'>
                        <span>
                          <i class='gg-image'></i>
                        </span>
                      </div>
                    </div>
                  </div>
                {/each}
              </div>
              {/if}
              {#if floor_edit_mode === true}
                <div class='edit_floor_container'>
                  <label class='image_upload_button' for='upload_image'>
                    <input
                      bind:this={imageInput}
                      on:change={inputImageHandler}
                      type='file'
                      accept='image/*'
                      id='upload_image'
                    />
                    <i class='gg-software-upload'></i>
                    <p>Select file to Upload</p>
                  </label>
                </div>
              {/if}
            </div>
            <button class='next_button' on:click|preventDefault={nextPage}>NEXT</button>
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

  /* Main Container Subpage */
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
  .realtor_subpage_container .back_button {
    padding: 5px;
    border: none;
    background-color: rgb(0 0 0 / 0.06);
    border-radius: 5px;
    font-size: 14px;
    cursor: pointer;
  }
  .realtor_subpage_container .back_button:hover {
    background-color: rgb(0 0 0 / 0.05);
    color: #ef4444;
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
    background-color: #f58433;
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
    height: 430px;
  }
  .location_input .map_container iframe {
    width: 100%;
    height: 100%;
    border: none;
    border-radius: 8px;
  }

  /* +============| SUBPAGE INFO |===========+ */
  .info {
    color: #334155;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    height: 100%;
  }
  .info .next_skip_button {
    display: flex;
    flex-direction: row;
    align-items: stretch;
    font-weight: 500;
    margin-top: 20px;
    width: 100%;
  }
  .info .next_skip_button button {
    border-radius: 8px;
    border: none;
    padding: 10px;
    cursor: pointer;
    width: 100%;
  }
  .info .next_skip_button .skip_button {
    margin-right: 10px;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    color: #f97316;
  }
  .info .next_skip_button .skip_button:hover {
    border: 1px solid #f97316;
  }
  .info .next_skip_button .next_button {
    background-color: #f97316;
    color: white;
  }
  .info .next_skip_button .next_button:hover {
    background-color: #f58433;
  }
  .info h2 {
    font-size: 18px;
    font-weight: 600;
    margin-bottom: 20px;
  }
  /* __SUBPAGE INFO - About The House */
  .about_the_house {
    display: grid;
    grid-template-rows: 1fr 1fr;
    gap: 20px;
  }
  .house_type_container {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 20px;
  }
  .house_type, .apartment_floor, .rent_or_sell {
    position: relative;
    display: flex;
    flex-direction: column;
    width: 100%;
    height: fit-content;
  }
  .house_type button, .apartment_floor input, .rent_or_sell button {
    width: auto;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    border-radius: 8px;
    padding: 10px 12px;
    font-weight: 600;
    text-align: left;
    color: #f97316;
    text-transform: capitalize;
  }
  .rent_or_sell button {
    width: 48.3% !important;
  }
  .house_type button:hover, .apartment_floor input:hover, .rent_or_sell button:hover {
    border: 1px solid #f97316;
  }
  .apartment_floor input:focus {
    outline: 1px solid #f97316;
  }
  .rent_or_sell .option_container {
    width: 50% !important;
  }
  .about_the_house .option_container {
    width: 100%;
    position: absolute;
    top: 70px;
    z-index: 3;
    display: flex;
    flex-direction: column;
    border-radius: 8px;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
  }
  .about_the_house .option_container .option {
    margin: 0;
    padding: 10px 12px;
    font-weight: 500;
    cursor: pointer;
  }
  .about_the_house .option_container .option:hover {
    color: #f97316;
  }
  .option input[type='radio'] {
    position: absolute;
    opacity: 0;
    pointer-events: none;
  }
  .house_type label, .apartment_floor label, .rent_or_sell label {
    font-size: 13px;
    font-weight: 700;
    margin-left: 10px;
    margin-bottom: 8px;
  }
  /* __SUBPAGE INFO - Upload House Photo*/
  .upload_house_photo {
    display: grid;
    gap: 20px;
  }
  .image_preview_container {
    display: grid;
    justify-items: center;
    grid-template-columns: repeat(3, 1fr);
    gap: 20px;
  }
  .image_card {
    position: relative;
    border-radius: 8px;
    width: 100%;
    height: 110px;
  }
  .image_card img {
    border-radius: 8px;
    object-fit: cover;
    width: 100%;
    height: 100%;
    border: 1px solid #cbd5e1;
  }
  .image_card button {
    position: absolute;
    z-index: 10;
    top: -10px;
    right: -10px;
    background-color: #ef4444;
    border: none;
    color: white;
    padding: 5px;
    border-radius: 50%;
    cursor: pointer;
    filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
  }
  .image_upload_button {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    border-radius: 8px;
    width: 100%;
    height: 110px;
    cursor: pointer;
  }
  .image_upload_button:hover {
    border: 1px solid #f97316;
    color: #f97316;
  }
  .image_upload_button input {
    position: absolute;
    opacity: 0;
    pointer-events: none;
  }
  .image_upload_button i {
    font-size: 60px;
    margin-bottom: 10px;
  }
  .image_upload_button p {
    font-size: 11px;
    margin: 0;
  }
  /* __SUBPAGE INFO - Feature and Facility */
  .feature_section {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    gap: 20px;
  }
  .feature_section .beds, .baths, .square_foot {
    display: flex;
    flex-direction: column;
  }
  .feature_section .beds input, .baths input, .square_foot input {
    width: 100%;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    border-radius: 8px;
    padding: 10px 12px;
    text-align: center;
    margin-bottom: 10px;
  }
  .facility_section textarea {
    resize: vertical;
    width: 100%;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    border-radius: 8px;
    padding: 10px 12px;
    text-align: left;
  }
  /* __SUBPAGE INFO - Description of Property */
  .description_of_property textarea {
    resize: vertical;
    width: 100%;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    border-radius: 8px;
    padding: 10px 12px;
    text-align: left;
  }
  /* __SUBPAGE INFO - Price */
  .price {
    display: grid;
    grid-template-rows: 1fr 1fr;
    gap: 20px;
  }
  .price .property_price, .agency_fee {
    display: flex;
    flex-direction: column;
  }
  .price label {
    font-size: 13px;
    font-weight: 700;
    margin-left: 10px;
    margin-bottom: 8px;
  }
  .price .property_price input, .agency_fee input {
    width: 50%;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    border-radius: 8px;
    padding: 10px 12px;
    text-align: left;
  }
  .price .agency_fee div {
    display: flex;
    flex-direction: row;
    align-items: center;
  }
  .price .agency_fee div span {
    margin-left: 8px;
    font-size: 16px;
    font-weight: 600;
  }

  /* +============| SUBPAGE FLOORS |===========+ */
  .floor {
    color: #334155;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    height: 100%;
    position: relative;
  }
  .floor .next_button {
    background-color: #f97316;
    border-radius: 8px;
    border: none;
    padding: 10px;
    color: white;
    font-weight: 600;
    margin-top: 20px;
    cursor: pointer;
  }
  .floor .next_button:hover {
    background-color: #f58433;
  }
  .floor .floor_header {
    margin: 20px 0;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    width: 100%;
  }
  .floor .floor_header h2 {
    font-weight: 600;
    margin: 0;
    font-size: 18px;
  }
  .floor .floor_header button {
    border-radius: 8px;
    border: none;
    padding: 10px 12px;
    color: #f97316;
    font-weight: 600;
    cursor: pointer;
    background: none;
    font-size: 14px;
  }
  .floor .floor_header button:hover {
    background-color : rgb(0 0 0 / 0.07);
  }
  .floor .add_floor_button {
    background-color: #f97316;
    color: white;
    border-radius: 8px;
    border: none;
    padding: 10px;
    cursor: pointer;
    width: 100%;
  }
  .floor .add_floor_button:hover {
    background-color: #f58433;
  }
  .floor .floor_items_container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    width: 100%;
  }
  .floor .floor_items_container .floor_item {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    padding: 20px;
    background-color: rgba(0 0 0 / 0.06);
    border-radius: 8px;
    height: fit-content;
    width: 60%;
    margin: 0 auto 20px auto;
  }
  .floor .floor_items_container .floor_item .left_item {
    flex-basis: 40%;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }
  .floor .floor_items_container .floor_item .left_item h4 {
    font-size: 16px;
    font-weight: 600;
    text-transform: capitalize;
    margin: 0;
  }
  .floor .floor_items_container .floor_item .left_item .rooms_total {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 20px;
    margin: 20px 0;
  }
  .floor .floor_items_container .floor_item .beds, .floor_item .baths {
    display: flex;
    flex-direction: column;
    text-align: center;
  }
  .floor .floor_items_container .floor_item .beds b, .floor_item .baths b {
    font-size: 25px;
  }
  .floor .floor_items_container .floor_item .beds p, .floor_item .baths p {
    margin: 10px 0 0 0;
  }
  .floor .floor_items_container .floor_item .right_item {
    flex-basis: 30%;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }
  .floor .floor_items_container .floor_item .right_item .edit_floor {
    background-color: #f97316;
    border-radius: 8px;
    border: none;
    padding: 7px;
    color: white;
    font-weight: 600;
    cursor: pointer;
    width: 100%;
    margin: 0 0 0 auto;
  }
  .floor .floor_items_container .floor_item .right_item .edit_floor:hover {
    background-color: #f58433;
  }
  .floor .floor_items_container .floor_item .right_item .floor_plan {
    position: relative;
    border-radius: 8px;
    width: 100%;
    height: 90px;
    border: 1px solid #cbd5e1;
    margin-top: 20px;
  }
  .floor .floor_items_container .floor_item .right_item .floor_plan span {
    border-radius: 8px;
    object-fit: cover;
    width: 100%;
    height: 100%;
    background-color: rgb(0 0 0 / 0.06);
    display: flex;
    justify-content: center;
    align-items: center;
  }
</style>