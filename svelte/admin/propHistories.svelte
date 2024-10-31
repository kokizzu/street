<script>
  /** @typedef {import('../_types/user').User} User */
  /** @typedef {import('../_types/master').Access} Access */
  /** @typedef {import('../_types/master').PagerIn} PagerIn */
  /** @typedef {import('../_types/master').PagerOut} PagerOut */
  /** @typedef {import('../_types/master').Field} Field */
  /** @typedef {import('../_types/property').Property} Property*/
  /** @typedef {import('../_types/master').ExtendedAction} ExtendedAction*/

  import Main from '../_layouts/Main.svelte';
  import AdminSubMenu from '../_components/AdminSubMenu.svelte';
  import TableView from '../_components/TableView.svelte';
  import ModalForm from '../_components/ModalForm.svelte';
  import { AdminPropHistories } from '../jsApi.GEN';
  import { notifier } from '../_components/notifier.js';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from '../node_modules/svelte-icons-pack/dist/ri';
  
  let user          = /** @type {User} */ ({/* user */});
  let access        = /** @type {Access} */ ({/* segments */});
  let fields        = /** @type {Field[]} */ ([/* fields */]);
  let pager         = /** @type {PagerOut} */ ({/* pager */});
  let propHistories = /** @type {any[][]} */ ([/* propHistories */]);

  let modalForm = /** @type {import('svelte').SvelteComponent}*/ (null);
  
  /**
	 * @description Handle AJAX response
	 * @param res {any}
	 * @returns {boolean}
	 */
	function handleResponse(res) {
    console.log(res);
    if (res.error) {
      notifier.showError( res.error );
      return false
    }

    if (res.propHistories && res.propHistories.length) {
      propHistories = res.propHistories;
    }

    if (res.pager && res.pager.page) {
      pager = res.pager;
    }
  }
  
  async function refreshTableView(/** @type {PagerIn} */ pagerIn ) {
    await AdminPropHistories( { // @ts-ignore
      pager: pagerIn,
      cmd: 'list',
    }, function( res ) {
      handleResponse( res );
    } );
  }
  
  async function editRow(/** @type {number|string} */ id, /** @type {any[] | any} */ row) {
    await AdminPropHistories( { // @ts-ignore
      propHistory: {id},
      cmd: 'form',
    }, async function( /** @type {any} */ res ) {
      if( !handleResponse( res ) )
        modalForm.showModal( res.propHistory );
    } );
  }
  
  function addRow() {
    modalForm.showModal( {id: ''} );
  }
  
  async function saveRow(/** @type {string} */ action, /** @type {any} */ row ) {
    let ph = {...row};
    if( !ph.id ) ph.id = '0';
    console.log( ph );
    await AdminPropHistories( { // @ts-ignore
      propHistory: ph,
      cmd: action,
      pager: pager, // force refresh page, will be slow
    }, function( res ) {
      if( handleResponse( res ) ) {
        return modalForm.setLoading( false ); // has error
      }
      modalForm.hideModal(); // success
    } );
  }

</script>

<Main {user} {access}>
  <div class="admin-container">
    <AdminSubMenu />
    <div class="admin-content">
      <ModalForm
        fields={fields}
        rowType="Prop History"
        bind:this={modalForm}
        onConfirm={saveRow}
      />
      <TableView
        fields={fields}
        bind:pager={pager}
        rows={propHistories}
        onRefreshTableView={refreshTableView}
        onEditRow={editRow}
      >
        <button
          class="btn"
          on:click={addRow}
          title="Add property history"
        >
          <Icon
            size={17}
            src={RiSystemAddBoxLine}
            color="var(--gray-008)"
          />
        </button>
      </TableView>
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