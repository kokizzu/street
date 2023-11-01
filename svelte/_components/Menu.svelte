<script>
  // @ts-nocheck
  import {UserLogout} from '../jsApi.GEN.js';
  import {onMount} from 'svelte';
  import {isSideMenuOpen} from './uiState.js';
  
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidHome from 'svelte-icons-pack/fa/FaSolidHome';
  import FaSolidShoppingBag from 'svelte-icons-pack/fa/FaSolidShoppingBag';
  import FaSolidBuilding from 'svelte-icons-pack/fa/FaSolidBuilding';
  import FaSolidSlidersH from 'svelte-icons-pack/fa/FaSolidSlidersH';
  import FaSolidUserCircle from 'svelte-icons-pack/fa/FaSolidUserCircle';
  import FaSolidSignInAlt from 'svelte-icons-pack/fa/FaSolidSignInAlt';
  import FaSolidTimes from 'svelte-icons-pack/fa/FaSolidTimes';
  
  export let doToggle = function() {
    isSideMenuOpen.set( !$isSideMenuOpen );
  };
  export let access = {
    'admin': false,
    'buyer': false,
    'realtor': false,
    'user': false,
  };
  
  let segment1;
  onMount( () => {
    console.log( 'onMount.Menu' );
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

{#if $isSideMenuOpen}
	<button class="backdrop" on:click={() => $isSideMenuOpen = !$isSideMenuOpen}></button>
{/if}
<aside class={$isSideMenuOpen ? `side_menu_admin open` : `side_menu_admin`}>
	<div class='side_menu_admin_container'>
		<header>
			<h3>STREET</h3>
			<button on:click|preventDefault={doToggle}>
				<Icon color='#475569' size={20} src={FaSolidTimes}/>
			</button>
		</header>
		<div class='menu_container'>
			<!-- PAGES -->
			<hr/>
			<h6>MENU</h6>
			<nav class='menu'>
				<a class:active={segment1 === ''} href='/'>
					<Icon className={segment1 === '' ? 'icon_active' : 'icon_dark'} size={22} src={FaSolidHome}/>
					<span>HOME</span>
				</a>
				{#if access.buyer }
					<a href='/buyer' class:active={segment1 === 'buyer'}>
						<Icon size={22} className={segment1 === 'buyer' ? 'icon_active' : 'icon_dark'} src={FaSolidShoppingBag}/>
						<span>BUYER</span>
					</a>
				{/if}
				{#if access.realtor}
					<a href='/realtor' class:active={segment1 === 'realtor'}>
						<Icon size={20} className={segment1 === 'realtor' ? 'icon_active' : 'icon_dark'} src={FaSolidBuilding}/>
						<span>REALTOR</span>
					</a>
				{/if}
				{#if access.admin }
					<a href='/admin' class:active={segment1 === 'admin'}>
						<Icon size={20} className={segment1 === 'admin' ? 'icon_active' : 'icon_dark'} src={FaSolidSlidersH}/>
						<span>ADMIN</span>
					</a>
				{/if}
			</nav>
			<!-- SETTING -->
			<hr/>
			<h6>SETTING</h6>
			<nav class='menu'>
				{#if access.user}
					<a href='/user' class:active={segment1 === 'user'}>
						<Icon size={22} className={segment1 === 'user' ? 'icon_active' : 'icon_dark'} src={FaSolidUserCircle}/>
						<span>PROFILE</span>
					</a>
				{/if}
				{#if access.user}
					<button on:click={userLogout} class='logout'>
						<Icon size={22} className='icon_dark' src={FaSolidSignInAlt}/>
						<span>LOGOUT</span>
					</button>
				{/if}
			</nav>
		</div>
	</div>
</aside>

<style>
    :global(.icon_dark) {
        fill : #475569;
    }

    :global(.icon_active) {
        fill : #EF4444;
    }

    .backdrop {
        z-index          : 8888;
        position         : fixed;
        top              : 0;
        bottom           : 0;
        left             : 0;
        right            : 0;
        border           : none;
        padding          : 0;
        background-color : rgba(0 0 0 / 20%);
    }

    .side_menu_admin {
        left             : -300px;
        display          : block;
        position         : fixed;
        z-index          : 9999;
        top              : 0;
        bottom           : 0;
        overflow-y       : auto;
        flex-direction   : row;
        flex-wrap        : nowrap;
        background-color : white;
        color            : #475569;
        padding          : 16px 24px;
        width            : 300px;
        transition       : 0.3s;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    }

    .side_menu_admin.open {
        left : 0;
    }

    .side_menu_admin_container {
        flex-direction : column;
        align-items    : stretch;
        min-height     : 100%;
        flex-wrap      : nowrap;
        padding        : 0;
        display        : flex;
        width          : 100%;
        margin         : 0 auto;
    }

    .side_menu_admin_container header {
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
        align-items     : center;
    }

    .side_menu_admin_container header h3 {
        font-size   : 16px;
        line-height : 1.5rem;
        padding     : 0;
        margin      : 0;
    }

    .side_menu_admin_container header button {
        padding       : 5px;
        border        : none;
        background    : none;
        border-radius : 5px;
        font-size     : 14px;
        cursor        : pointer;
    }

    .side_menu_admin_container header button:hover {
        background-color : rgb(0 0 0 / 0.07);
        color            : #EF4444;
    }

    .menu_container {
        margin-top     : 1rem;
        align-items    : stretch;
        flex-direction : column;
        display        : flex;
    }

    .menu_container hr {
        margin : 1rem 0;
    }

    .menu_container h6 {
        font-size : 15px;
        margin    : 12px 0;
    }

    .menu_container .menu {
        display        : flex;
        flex-direction : column;
        margin-bottom  : 10px;
    }

    .menu_container .menu a, .menu .logout { /*MENU LISTS*/
        color            : #475569;
        text-decoration  : none;
        margin           : 0;
        padding          : 0.75rem 0;
        font-size        : 0.875rem !important;
        line-height      : 1.25rem;
        font-weight      : 700;
        text-transform   : uppercase;
        text-align       : left;
        background-color : transparent;
        border           : none;
        display          : flex;
        flex-direction   : row;
        align-items      : center;
        gap              : 15px;
    }

    .menu_container .menu .logout {
        cursor        : pointer;
        margin-top    : 0;
        margin-bottom : 0;
        margin-right  : 0;
    }


    .menu_container .menu a:hover, .menu_container .menu .logout:hover { /*HOVER*/
        color : #64748B;
    }

    .active { /*ACTIVE Navigation*/
        color : #EF4444 !important;
    }

    /* Responsive to mobile device */
    @media (max-width : 768px) {
        .side_menu_admin {
            width : 240px;
        }
    }
</style>