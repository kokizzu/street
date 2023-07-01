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
      <hr />
      <!-- PAGES -->
      <nav class='menu'>
        <h6>LAYOUT PAGES</h6>
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
      <hr />
      <!-- AUTH -->
      <nav class='menu'>
        <h6>NO LAYOUT PAGES</h6>
        {#if access.user}
          <a href='/user' class:active={segment1 === 'user'}><i class="fa-solid fa-user-circle"></i> PROFILE</a>
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
      color: #475569;
      padding: 20px 23px;
      width: 11rem;
      left: 0;
      bottom: 0;
      top: 0;
      position: fixed;
      overflow-y: scroll;
      overflow: hidden;
      display: block;
    }
    .side_menu_admin_container {
      width: 100%;
    }
    .side_menu_admin_container h6 {
      font-size: 15px;
      margin: 8px 0;
    }
    .menu_container {
      margin-top: 1rem;
      align-items: stretch;
      flex-direction: column;
      display: flex;
    }
    .menu_container .menu a{
      color: #475569;
      text-decoration: none;
      margin: 8px 0;
      font-size: 13px;
      line-height: 1rem;
      font-weight: 700;
    }
    .menu_container .menu a i, .menu_container .menu button i {
      margin-right: 0.3rem;
      font-size: 14px;
      color: #CBD5E1;
    }
    .menu_container .menu a:hover, .menu_container .menu button:hover {
      color: #64748B;
    }
    .active, .active i {
      color: #EF4444 !important;
    }
    .menu_container .menu {
      display: flex;
      flex-direction: column;
      margin-bottom: 10px;
    }
    .menu_container .menu .logout {
      text-align: left;
      padding: 0;
      margin: 8px 0;
      font-weight: 700;
      background-color: transparent;
      border: none;
      font-size: 13px;
      color: #475569;
    }
  }
</style>