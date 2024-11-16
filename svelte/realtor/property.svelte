<script>
  /** @typedef {import('../_types/user').User} User */
  /** @typedef {import('../_types/master').Access} Access */

  import Main from '../_layouts/Main.svelte';
  
  let user      = /** @type {User} */ ({/* user */});
  let access    = /** @type {Access} */ ({/* segments */});
</script>

<Main {user} {access}>
  <div class="property-container">
    <a href="/realtor/propertyOld">Go to Old Page</a>
  </div>
</Main>

<style>
  
</style>