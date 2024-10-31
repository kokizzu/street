<script>
  /** @typedef {import('../_types/user').User} User */
  /** @typedef {import('../_types/master').Access} Access */
  /** @typedef {import('../_types/master').PagerIn} PagerIn */
  /** @typedef {import('../_types/master').PagerOut} PagerOut */
  /** @typedef {import('../_types/master').Field} Field */

  import Main from '../_layouts/Main.svelte';
  import AdminSubMenu from '../_components/AdminSubMenu.svelte';
  import { AdminUsers } from '../jsApi.GEN';
  import ModalForm from '../_components/ModalForm.svelte';
  import { notifier } from '../_components/notifier.js';
  import TableView from '../_components/TableView.svelte';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from '../node_modules/svelte-icons-pack/dist/ri';
  
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
    await AdminUsers( { // @ts-ignore
      pager: pagerIn,
      cmd: 'list',
    }, function( res ) {
      handleResponse( res );
    } );
  }
  
  let form = /** @type {import('svelte').SvelteComponent} */ (null);
  
  async function editRow(/** @type {string | number}*/ id, /** @type {any[] | any}*/ row ) {
    await AdminUsers( { // @ts-ignore
      user: { id },
      cmd: 'form',
    }, function( res ) {
      if( !handleResponse(/** @type {any} */ res ) ) // @ts-ignore
        form.showModal( res.user );
    } );
  }
  
  function addRow() {
    form.showModal( {id: ''} );
  }
  
  async function saveRow( action, row ) {
    let user = {...row};
    if( !user.id ) user.id = '0';
    await AdminUsers( { // @ts-ignore
      user: user,
      cmd: action,
      pager: pager, // force refresh page, will be slow
    }, function( res ) {
      if( handleResponse( res ) ) {
        return form.setLoading( false ); // has error
      }
      form.hideModal(); // success
    } );
  }
</script>


<Main {user} {access}>
  <div class="admin-container">
    <AdminSubMenu />
    <div class="admin-content">
      <ModalForm {fields}
        rowType="User"
        bind:this={form}
        onConfirm={saveRow}
      />
      <section class="users-master-container">
        <TableView
          fields={fields}
          bind:pager={pager}
          rows={users}
          onRefreshTableView={refreshTableView}
          onEditRow={editRow}
        >
          <button on:click={addRow} class="btn">
            <Icon
              size="17"
              src={RiSystemAddBoxLine}
              color="var(--gray-008)"
            />
          </button>
        </TableView>
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

  .users-master-container {
    display: flex;
    flex-direction: column;
    width: 100%;
    gap: 20px;
  }
</style>