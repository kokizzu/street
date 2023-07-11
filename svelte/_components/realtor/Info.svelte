<script>
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
</script>

<div class='info_container'>
   <div>
      <button class='back_button' on:click|preventDefault={backPage}><i class='gg-chevron-left'></i></button>
      <h2>About the House</h2>

      <div class='location_input'>
         
      </div>
   </div>

   <button class='next_button' on:click|preventDefault={nextPage}>NEXT</button>
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
      font-weight: 600;
      margin-top: 20px;
      cursor: pointer;
   }
   .info_container .next_button:hover {
      background-color: #f58433
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
</style>