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
         <h6>MENU</h6>
         <nav class='menu'>
            <a href='/' class:active={segment1 === ''}>
               <i class="gg-home"></i>
               <span>HOME</span>
            </a>
            {#if access.buyer }
               <a href='/buyer' class:active={segment1 === 'buyer'}>
                  <i class="gg-shopping-bag"></i>
                  <span>BUYER</span>
               </a>
            {/if}
            {#if access.realtor}
               <a href='/realtor' class:active={segment1 === 'realtor'}>
                  <i class="gg-hello"></i>
                  <span>REALTOR</span>
               </a>
            {/if}
            {#if access.admin }
               <a href='/admin' class:active={segment1 === 'admin'}>
                  <i class="gg-options" style="margin-left: 5px !important; margin-right: 5px;"></i>
                  <span>ADMIN</span>
               </a>
            {/if}
         </nav>
         <!-- SETTING -->
         <hr />
         <h6>SETTING</h6>
         <nav class='menu'>
            {#if access.user}
               <a href='/user' class:active={segment1 === 'user'}>
                  <i class="gg-profile"></i>
                  <span>PROFILE</span>
               </a>
            {/if}
            {#if access.user}
               <button on:click={userLogout} class="logout">
                  <i class="gg-log-out"></i>
                  <span>LOGOUT</span>
               </button>
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
      overflow: auto;
      background-color: white;
      color: #475569;
      padding: 16px 24px;
      width: 256px;
      filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
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
      margin: 12px 0;
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
      font-size: 0.875rem !important;
      line-height: 1.25rem;
      font-weight: 700;
      text-transform: uppercase;
      text-align: left;
      background-color: transparent;
      border: none;
      display: flex;
      flex-direction: row;
      align-items: center;
      align-content: center;
    }
    .menu_container .menu .logout {
      margin-left: 4px !important;
      margin-top: 0;
      margin-bottom: 0;
      margin-right: 0;
    }
    .menu_container .menu .logout span {
      margin-left: 25px !important;
    }
    .menu_container .menu a i, .menu_container .menu .logout i { /*ICON*/
      margin: 0;
      color: #CBD5E1;
    }
    .menu_container .menu a:hover, .menu_container .menu .logout:hover { /*HOVER*/
      color: #64748B;
    }
    .menu_container .menu a span, .menu_container .menu .logout span {
      margin-left: 15px;
    }
    .active, .active i { /*ACTIVE Navigation*/
      color: #EF4444 !important;
    }
  }
</style>