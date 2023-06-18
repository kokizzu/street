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
  } );
  
  async function userLogout() {
    await UserLogout( {}, function( o ) {
      console.log( o );
      if( o.error ) return alert( o.error );
      window.location = '/';
    } );
  }
</script>

<ul class='menu'>
  <li class:active={segment1 === ''}><a href='/'>Home</a></li>
  {#if access.buyer }
    <li class:active={segment1 === 'buyer'}><a href='/buyer'>Buyer</a></li>
  {/if}
  {#if access.realtor}
    <li class:active={segment1 === 'realtor'}><a href='/realtor'>Realtor</a></li>
  {/if}
  {#if access.admin }
    <li class:active={segment1 === 'admin'}><a href='/admin'>Admin</a></li>
  {/if}
  {#if access.user}
    <li style='float:right'>
      <button on:click={userLogout}>Logout</button>
    </li>
    <li class:active={segment1 === 'user'} style='float:right'>
      <a href='/user'>Profile</a>
    </li>
  {/if}
</ul>

<style>
    .menu {
        list-style-type  : none;
        margin           : 0;
        padding          : 0;
        overflow         : hidden;
        background-color : #333;
    }

    .menu li {
        float : left;
    }

    .menu li a {
        display         : block;
        color           : white;
        text-align      : center;
        padding         : 14px 16px;
        text-decoration : none;
    }

    .menu li a:hover:not(.active) {
        background-color : #111;
    }

    .menu li button {
        padding : 0.5em;
        margin  : 0.3em;
    }

    .active {
        background-color : #4CAF50;
    }
</style>