<script>
  /** @typedef {import('../_types/user').User} User */
  /** @typedef {import('../_types/master').Access} Access */

  import Main from '../_layouts/Main.svelte';
  import AdminSubMenu from '../_components/AdminSubMenu.svelte';

  let user = /** @type {User} */ ({/* user */});
  let access = /** @type {Access} */ ({/* segments */});
</script>

<Main {user} {access}>
  <div class="admin-container">
    <AdminSubMenu />
    <div class="admin-content" style="flex: 1 !important; display: flex !important; flex-direction: column !important; gap: 0 !important; overflow: hidden;">
      <section class="posthog-section" style="flex: 1; display: flex; min-height: 0; overflow: hidden;">
        <div class="posthog-container" style="width: 100%; height: 100vh; display: flex; flex-direction: column; padding: 20px; box-sizing: border-box;">
          <h2 style="margin: 0 0 20px 0; flex-shrink: 0;">PostHog Analytics Dashboard</h2>
          <iframe 
            width="100%" 
            frameborder="0" 
            allowfullscreen 
            src="https://us.posthog.com/embedded/h0BhXQ5AT_fQgxTvUR66po4T62VLSA" 
            key="0" 
            sandbox="allow-scripts allow-same-origin allow-popups"
            style="border: 1px solid #e9ecef; border-radius: 8px; flex: 1; min-height: 600px; height: 100%;">
          </iframe>
        </div>
      </section>
    </div>
  </div>
</Main>

<style>
  .admin-container {
    display: flex;
    flex-direction: column;
    width: 100%;
    gap: 20px;
    padding: 10px 20px 20px;
  }

  .admin-container .admin-content {
    display: flex;
    flex-direction: column;
    width: 100%;
    gap: 20px;
  }
</style>
