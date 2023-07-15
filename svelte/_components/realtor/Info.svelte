<script>
  import { createEventDispatcher } from 'svelte';
  import {stackPageCount, subpages} from '../uiState';

  const dispatch = createEventDispatcher();

  let count;
  stackPageCount.subscribe(value => {
    count = value;
  });
  function nextPage() {
    dispatch('nextPage')
  }
  function backPage() {
    dispatch('backPage')
  }

  let modeCount = 0;
  let modeSkippable = false;
  // +=== House info properties ===+
  const ABOUT_THE_HOUSE = 'About The House';
  const UPLOAD_HOUSE_PHOTO = 'Upload House Photo';
  const FEATURE_OR_FACILITY = 'Feature or Facility';
  const DESCRIPTION_PROPERTY = 'Description of this property';
  const PRICE = 'Price';
  const modeLists = [
    { mode: ABOUT_THE_HOUSE, skip: false },
    { mode: UPLOAD_HOUSE_PHOTO, skip: true },
    { mode: FEATURE_OR_FACILITY, skip: true },
    { mode: DESCRIPTION_PROPERTY, skip: true },
    { mode: PRICE, skip: false }
  ];
  let mode = modeLists[modeCount].mode;
  function nextBtnHandler() {
    if (modeCount < modeLists.length - 1) {
      modeCount++;
      mode = modeLists[modeCount].mode;
      modeSkippable = modeLists[modeCount].skip;
    } else {
      nextPage();
    }
  }
  function backBtnHandler() {
    if (modeCount > 0) {
      modeCount--;
      mode = modeLists[modeCount].mode;
      modeSkippable = modeLists[modeCount].skip;
    } else {
      backPage();
    }
  }
  function skipBtnHandler() {
    if (modeCount < modeLists.length - 1) {
      modeCount++;
      mode = modeLists[modeCount].mode;
      modeSkippable = modeLists[modeCount].skip;
    } else {
      nextPage();
    }
  }

  // +==========| ABOUT THE HOUSE |============+
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

  // +==========| UPLOAD HOUSE PHOTO |============+
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
</script>

<div class='info_container'>
  <div class='main_info'>
    <button class='back_button' on:click|preventDefault={backBtnHandler}>
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
                type='text'
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
      <button class='skip_button' on:click|preventDefault={skipBtnHandler}>
        SKIP
      </button>
    {/if}
    <button class='next_button' on:click|preventDefault={nextBtnHandler}>
      NEXT
    </button>
  </div>
</div>

<style>
  .info_container {
    color: #334155;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    height: 100%;
  }
  .info_container .next_skip_button {
    display: flex;
    flex-direction: row;
    align-items: stretch;
    font-weight: 500;
    margin-top: 20px;
    width: 100%;
  }
  .info_container .next_skip_button button {
    border-radius: 8px;
    border: none;
    padding: 10px;
    cursor: pointer;
    width: 100%;
  }
  .info_container .next_skip_button .skip_button {
    margin-right: 10px;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    color: #f97316;
  }
  .info_container .next_skip_button .skip_button:hover {
    border: 1px solid #f97316;
  }
  .info_container .next_skip_button .next_button {
    background-color: #f97316;
    color: white;
  }
  .info_container .next_skip_button .next_button:hover {
    background-color: #f58433;
  }
  .info_container h2 {
    font-size: 18px;
    font-weight: 600;
    margin-bottom: 20px;
  }
  .back_button {
    padding: 5px;
    border: none;
    background-color: rgb(0 0 0 / 0.06);
    border-radius: 5px;
    font-size: 14px;
    cursor: pointer;
  }
  .back_button:hover {
    background-color: rgb(0 0 0 / 0.05);
    color: #ef4444;
  }

  /* +==========| ABOUT THE HOUSE |===========+ */
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
  
  /* +=============| UPLOAD HOUSE PHOTO |=============+ */
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

  /* +===============| Feature and Facility |================+ */
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

  /* +====================| Description of Property |==================+ */
  .description_of_property textarea {
    resize: vertical;
    width: 100%;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    border-radius: 8px;
    padding: 10px 12px;
    text-align: left;
  }

  /* +==============| Price |===================+ */
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
</style>
