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
        <div class="slot">
          <slot />
        </div>
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
		overflow-y: auto;
		min-height: calc(100vh - var(--navbar-height));
		transition: 0.3s;
		width: 100%;
  }

  .root_layout .root_container .root_content .content {
    display: flex;
    flex-direction: column;
    min-height: 100%;
    overflow: inherit;
    margin-left: var(--sidemenu-width);
    margin-top: calc(var(--navbar-height) + 1px);
  }

  .root_layout .root_container .root_content .content .slot {
    height: 100%;
    width: 100%;
    padding: 20px;
  }
</style>