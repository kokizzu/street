<script>
   // TODO, Make conditional in mode
   import {stackPageCount} from '../uiState';

   let count = 0;
   stackPageCount.subscribe((value) => {
      count = value;
   })
   function nextPage() {
      if(count < 4)
      stackPageCount.update((n) => n + 1);
   }
   function backPage() {
      if(count > 1)
      stackPageCount.update((n) => n - 1);
   }

   let modeCount = 0;
   // +=== House info properties ===+
   const ABOUT_THE_HOUSE = 'ABOUT-THE-HOUSE';
   const UPLOAD_HOUSE_PHOTO = 'UPLOAD-HOUSE-PHOTO';
   const FEATURE_OR_FACILITY = 'FEATURE-OR-FACILITY';
   const DESCRIPTION_PROPERTY = 'DESCRIPTION-PROPERTY';
   const modeLists = [
      ABOUT_THE_HOUSE,
      UPLOAD_HOUSE_PHOTO,
      FEATURE_OR_FACILITY,
      DESCRIPTION_PROPERTY
   ]
   let mode = modeLists[0];
   function nextBtnHandler() {
      if (modeCount < modeLists.length) {
         modeCount++;
      } else {
         nextPage();
      }
   }

   // **** House Type
   let house_type = '';
   let showHouseOption = false;
   let isApartment = false;
   function handleHouseTypeOption() {
      showHouseOption = !showHouseOption;
   }
   function selectHouseType() {
      isApartment = house_type === 'apartment';
      showHouseOption = !showHouseOption;
      console.log(house_type);
   }

   // **** Rent or Sell
   let rent_or_sell = ''
   let showRentSellOption = false;
   function handleRentSellOption() {
      showRentSellOption = !showRentSellOption;
   }
   function selectRentOrSell() {
      showRentSellOption = !showRentSellOption;
      console.log(rent_or_sell);
   }
</script>

<div class='info_container'>
   <div>
      <button class='back_button' on:click|preventDefault={backPage}><i class='gg-chevron-left'></i></button>
      <h2>About the House</h2>

      <div class='info_input'>
         <div class='house_type_container'>
            <div class='house_type'>
               <label for='house_type'>House Type</label>
               <button on:click={handleHouseTypeOption}>{house_type ? house_type : 'select'}</button>
               {#if showHouseOption}
               <div class='option_container'>
                  <label class='option' for='house'>
                     <input type='radio' bind:group={house_type} on:change={() => selectHouseType()} id='house' value='house'/>
                     House
                  </label>
                  
                  <label class='option' for='apartment'>
                     <input type='radio' bind:group={house_type} on:change={() => selectHouseType()} id='apartment' value='apartment'>
                     Apartment
                  </label>
               </div>
               {/if}
            </div>
            {#if isApartment}
               <div class='apartment_floor'>
                  <label for='apartment_floor'>Floor</label>
                  <input type='text' name='apartment_floor' id='apartment_floor' placeholder='Your property based'/>
               </div>
            {/if}
         </div>
         <div class='rent_or_sell'>
            <label for='rent_or_sell'>Rent or Sell</label>
            <button disabled={!house_type} on:click={handleRentSellOption}>{rent_or_sell ? rent_or_sell : 'select'}</button>
            {#if showRentSellOption}
               <div class='option_container'>
                  <label class='option' for='rent'>
                     <input type='radio' bind:group={rent_or_sell} on:change={() => selectRentOrSell()} id='rent' value='rent'/>
                     Rent
                  </label>
                  
                  <label class='option' for='sell'>
                     <input type='radio' bind:group={rent_or_sell} on:change={() => selectRentOrSell()} id='sell' value='sell'>
                     Sell
                  </label>
               </div>
               {/if}
         </div>
      </div>
   </div>

   <button class='next_button' on:click|preventDefault={nextBtnHandler}>NEXT</button>
</div>

<style>
   .info_container {
      color: #334155;
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      height: 100%;
   }
   .info_container .next_button {
      background-color: #f97316;
      border-radius: 8px;
      border: none;
      padding: 10px;
      color: white;
      font-weight: 500;
      margin-top: 20px;
      cursor: pointer;
   }
   .info_container .next_button:hover {
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
      color: #EF4444;
   }

   .info_input {
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
      border: 1px solid #CBD5E1;
      background-color: #F1F5F9;
   	border-radius: 8px;
   	padding: 10px 12px;
      cursor: pointer;
      font-weight: 600;
      text-align: left;
      color: #f97316;
   }
   .house_type button:hover, .apartment_floor input:hover, .rent_or_sell button:hover {
      border: 1px solid #f97316;
   }
   .apartment_floor input:focus {
      outline: 1px solid #f97316;
   }
   .info_input .option_container {
      width: 100%;
      position: absolute;
      top: 70px;
      z-index: 3;
      display: flex;
      flex-direction: column;
      border-radius: 8px;
      border: 1px solid #CBD5E1;
      background-color: #F1F5F9;
   }
   .info_input .option_container .option {
      margin: 0;
      padding: 10px 12px;
      font-weight: 500;
      cursor: pointer;
   }
   .info_input .option_container .option:hover {
      color: #f97316;
   }
   .option input[type="radio"] {
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
</style>