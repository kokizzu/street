<script>
  /** @typedef {import('../_types/user').User} User */

  import SideMenu from '../_components/partials/SideMenu.svelte';
  import Navbar from '../_components/partials/Navbar.svelte';
  import Footer from '../_components/partials/Footer.svelte';
  import { isShrinkMenu } from '../_states/page';

  export let user = /** @type {User} */ ({});
</script>

<div class="root_layout">
  <div class="root_container">
    <Navbar username={user.userName}/>
    <div class="root_content { $isShrinkMenu ? 'shrink' : 'expand' }">
      <SideMenu />
      <main class="content">
        <slot />
        <Footer />
      </main>
    </div>
  </div>
</div>

<style>
  .root_layout {
    display: block;
		top: 0;
		bottom: 0;
		left: 0;
		right: 0;
		height: 100vh;
		width: 100vw;
  }

  .root_layout .root_container {
    height: 100%;
		width: 100%;
		display: flex;
  }

  .root_layout .root_container .root_content {
		display: flex;
		flex-direction: column;
		-webkit-box-orient: vertical;
		-webkit-box-direction: normal;
		min-height: calc(100vh - var(--navbar-height));
		transition: 0.3s;
		width: 100%;
    max-width: 100%;
    margin-top: var(--navbar-height);
    margin-left: var(--sidemenu-width);
    margin-right: 0;
    margin-bottom: 0;
    overflow: hidden;
  }

  .root_layout .root_container .root_content .content {
    display: flex;
		flex-direction: column;
		justify-content: space-between;
		min-height: 100%;
    overflow-y: auto;
    padding: 0;
    margin: 0;
  }
</style>