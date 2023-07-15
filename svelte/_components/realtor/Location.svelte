<script>
   import { createEventDispatcher } from 'svelte';
   import {stackPageCount, subpages} from '../uiState';

   const dispatch = createEventDispatcher();
   let count;
   stackPageCount.subscribe((value) => {
      count = value;
   })
   function nextPage() {
      dispatch('nextPage');
   }
</script>

<div class='location_container'>
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
</div>

<style>
   .location_container {
      color: #334155;
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      height: 100%;
   }
   .location_container .next_button {
      background-color: #f97316;
      border-radius: 8px;
      border: none;
      padding: 10px;
      color: white;
      font-weight: 600;
      margin-top: 20px;
      cursor: pointer;
   }
   .location_container .next_button:hover {
      background-color: #f58433
   }
   .location_container h2 {
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
      border-color: #3b82f6;
      outline: 1px solid #3b82f6;
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