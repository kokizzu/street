<script>
  /** @typedef {import('../../_types/master.js').Access} Access */

  import { UserLogout } from '../../jsApi.GEN';
  import { notifier } from '../notifier';

  export let access = /** @type {Access} */ ({});

  const pathAll = /** @type {string}*/ (window.location.pathname);
  const pathLv1 = /** @type {string}*/ (window.location.pathname.split( '/' )[ 1 ]);
  const pathLv2 = /** @type {string}*/ (window.location.pathname.split( '/' )[ 2 ]);

  async function logout() {
    await UserLogout( {}, function( /** @type any */ o ) {
      if( o.error ) {
        notifier.showError( o.error );
        console.log( o );
        return;
      }

      notifier.showSuccess( 'Logged out' );
      window.location = /** @type {Location | (string & Location)} */ ('/');

      return new Promise( resolve => {
        setTimeout( () => {
          resolve();
        }, 1000 );
      });
    });
  }
</script>

<aside>
  <div class="container">
    <nav class="nav-menu">
      <a href="/" class:active={pathLv1 === ''}>Home</a>
      <a href="/user/buyers" class:active={pathAll === '/user/buyers'}>Buyers</a>
      <a href="/realtor" class:active={pathAll === '/realtor'}>Realtors</a>
      <a href="/user/listings" class:active={pathLv2 === 'listings'}>Listings</a>
      {#if access.admin}
        <a href="/admin/revenue" class:active={pathAll === '/admin/revenue'}>Revenue</a>
      {:else}
        <a href="/realtor/revenue" class:active={pathAll === '/realtor/revenue'}>Revenue</a>
      {/if}
    </nav>
    <span class="separator" />
    {#if access.user}
      <nav class="nav-menu">
        {#if access.admin}
          <a href="/admin" class:active={pathLv1 === 'admin' && pathAll !== '/admin/revenue'}>Admin</a>
        {/if}
        <a href="/user" class:active={pathAll === '/user'}>Profile</a>
        <button class="red" on:click={logout}>Logout</button>
      </nav>
    {/if}
  </div>
</aside>

<style>
  aside {
    position: fixed;
    top: var(--navbar-height);
    left: 0;
    bottom: 0;
    height: 100%;
    width: var(--sidemenu-width);
    border-right: 1px solid var(--gray-002);
    background-color: #FFF;
  }

  aside .separator {
    width: 80%;
    margin: auto;
    height: 1px;
    background-color: var(--gray-002);
  }

  aside .container {
    padding: 10px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    height: fit-content;
    width: 100%;
    margin: 0;
  }

  aside .container .nav-menu {
    display: flex;
    flex-direction: column;
    gap: 5px;
    height: fit-content;
    flex-wrap: nowrap;
    width: 100%;
    margin: 0;
  }

  aside .container .nav-menu a,
  aside .container .nav-menu button {
    text-decoration: none;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    padding: 10px 15px;
    border-radius: 8px;
    cursor: pointer;
    font-size: 17px;
    color: var(--gray-008);
    border: none;
    background-color: transparent;
  }

  aside .container .nav-menu a:hover,
  aside .container .nav-menu button:hover {
    background-color: var(--gray-001);
  }

  aside .container .nav-menu a.active {
    background-color: var(--orange-transparent);
    color: var(--orange-007);
  }

  aside .container .nav-menu button.red:hover {
    background-color: var(--red-transparent);
    color: var(--red-007);
  }
</style>