<script>
  /** @typedef {import('../_types/user').User} User */
  /** @typedef {import('../_types/master').Access} Access */
  /** @typedef {import('../_types/master').PagerIn} PagerIn */
  /** @typedef {import('../_types/master').PagerOut} PagerOut */
  /** @typedef {import('../_types/master').Field} Field */

  import Main from '../_layouts/Main.svelte';
  import { formatPrice } from '../_components/formatter';
  import InputBox from '../_components/InputBox.svelte';
  import { UserBuyers } from '../jsApi.GEN';
  import { notifier } from '../_components/notifier';
  import TableView from '../_components/TableView.svelte';
  
  let user    = /** @type {User} */ ({/* user */});
  let access = /** @type {Access} */ ({/* segments */});
  let fields  = /** @type {Field[]} */ ([/* fields */]);
  let users   = /** @type {User[]} */ ([/* users */]);
  let pager   = /** @type {PagerOut} */ ({/* pager */});

  /**
	 * @description Handle AJAX response
	 * @param res {any}
	 * @returns {boolean}
	 */
	function handleResponse(res) {
    console.log( res );
    if( res.error ) {
      notifier.showError(res.error );
      return true;
    }
    if( res.users && res.users.length ) users = res.users;
    if( res.pager && res.pager.page ) pager = res.pager;
  }
  
  async function refreshTableView(/** @type {PagerIn}*/ pagerIn ) {
    await UserBuyers( { // @ts-ignore
      pager: pagerIn,
      cmd: 'list',
    }, function( res ) {
      handleResponse( res );
    } );
  }
</script>

<Main {user} {access}>
  <div class="buyers-container">
    <section class="users-master-container">
      <TableView
        fields={fields}
        bind:pager={pager}
        rows={users}
        onRefreshTableView={refreshTableView}
        isNoActions={true}
      />
    </section>
  </div>
</Main>

<style>
  .buyers-container {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }
</style>