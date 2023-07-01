<script>
   import { onMount } from 'svelte';
   import { UserLogout } from '../jsApi.GEN.js';
  
   export let access = {
      'admin': false,
      'buyer': false,
      'realtor': false,
      'user': false,
   };
  
   let segment1;
   onMount( () => {
      segment1 = window.location.pathname.split( '/' )[ 1 ];
      console.log(segment1);
   } );
  
   async function userLogout() {
      await UserLogout( {}, function( o ) {
         console.log( o );
         if( o.error ) return alert( o.error );
         window.location = '/';
      } );
   }
</script>

<aside class="side_menu_admin">
   <div class="side_menu_admin_container">
      <h3>STREET</h3>
      <div class="menu_container">
         <!-- PAGES -->
         <hr />
         <h6>LAYOUT PAGES</h6>
         <nav class='menu'>
            <a href='/' class:active={segment1 === ''}><i class="fas fa-home" aria-hidden="true"></i> HOME</a>
            {#if access.buyer }
               <a href='/buyer' class:active={segment1 === 'buyer'}><i class="fa-solid fa-money-bill-wave"></i> BUYER</a>
            {/if}
            {#if access.realtor}
               <a href='/realtor' class:active={segment1 === 'realtor'}><i class="fa-solid fa-user-tie"></i> REALTOR</a>
            {/if}
            {#if access.admin }
               <a href='/admin' class:active={segment1 === 'admin'}><i class="fa-sharp fa-solid fa-user-gear"></i> ADMIN</a>
            {/if}
         </nav>
         <!-- NO LAYOUT PAGES -->
         <hr />
         <h6>NO LAYOUT PAGES</h6>
         <nav class='menu'>
            {#if access.user}
               <a href='/user' class:active={segment1 === 'user'}><i class="fa-solid fa-user-circle"></i> PROFILE</a>
            {/if}
         </nav>
         <!-- AUTH -->
         <hr />
         <h6>AUTH</h6>
         <nav class='menu'>
            {#if access.user}
               <button on:click={userLogout} class="logout"><i class="fa-solid fa-right-from-bracket"></i> LOGOUT</button>
            {/if}
         </nav>
      </div>
   </div>
</aside>

<style>
  /* Responsive */
  @media (min-width: 768px) {
    .side_menu_admin {
      left: 0;
      display: block;
      position: fixed;
      top: 0;
      bottom: 0;
      overflow-y: auto;
      flex-direction: row;
      flex-wrap: nowrap;
      overflow: hidden;
      background-color: white;
      color: #475569;
      padding: 16px 24px;
      width: 256px;
    }
    .side_menu_admin_container {
      flex-direction: column;
      align-items: stretch;
      min-height: 100%;
      flex-wrap: nowrap;
      padding: 0;
      display: flex;
      width: 100%;
      margin: 0 auto;
    }
    .side_menu_admin_container h3 {
      font-size: 16px;
      line-height: 1.5rem;
      padding: 15px 0 10px 0;
      margin: 0;
    }
    
    .menu_container {
      margin-top: 1rem;
      align-items: stretch;
      flex-direction: column;
      display: flex;
    }
    .menu_container hr {
      margin: 1rem 0;
    }
    .menu_container h6 {
      font-size: 15px;
      margin: 8px 0;
    }

    .menu_container .menu { 
      display: flex;
      flex-direction: column;
      margin-bottom: 10px;
    }
    .menu_container .menu a, .menu .logout{ /*MENU LISTS*/
      color: #475569;
      text-decoration: none;
      margin: 0;
      padding: 0.75rem 0;
      font-size: 0.875rem;
      line-height: 1.25rem;
      font-weight: 700;
      text-transform: uppercase;
      text-align: left;
      background-color: transparent;
      border: none;
    }
    .menu_container .menu a i, .menu_container .menu button i { /*ICON*/
      margin-right: 0.8rem;
      font-size: 14px;
      color: #CBD5E1;
    }
    .menu_container .menu a:hover, .menu_container .menu button:hover { /*HOVER*/
      color: #64748B;
    }
    .active, .active i { /*ACTIVE Navigation*/
      color: #EF4444 !important;
    }
  }
</style>