<script>
  /** @typedef {import('../_types/user').User} User */
  /** @typedef {import('../_types/master').Access} Access */
  /** @typedef {import('../_types/master').PagerIn} PagerIn */
  /** @typedef {import('../_types/master').PagerOut} PagerOut */
  /** @typedef {import('../_types/master').Field} Field */

  import Main from '../_layouts/Main.svelte';
  import AdminSubMenu from '../_components/AdminSubMenu.svelte';
  import TableView from '../_components/TableView.svelte';
  import ModalForm from '../_components/ModalForm.svelte';
  import { AdminFeedbacks } from '../jsApi.GEN';
  import { notifier } from '../_components/notifier.js';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { RiSystemInformation2Line } from '../node_modules/svelte-icons-pack/dist/ri';
  
  let user        = /** @type {User} */ ({/* user */});
  let access      = /** @type {Access} */ ({/* segments */});
  let fields      = /** @type {Field[]} */ ([/* fields */]);
  let feedbacks   = /** @type {any[][]}*/ [/* feedbacks */];
  let pager       = /** @type {PagerOut} */ ({/* pager */});
  let users       = /** @type {User[]} */ ([/* users */]);
  let modalForm   = /** @type {import('svelte').SvelteComponent}*/ (null);
  
  /**
	 * @description Handle AJAX response
	 * @param res {any}
	 * @returns {boolean}
	 */
	function handleResponse(res) {
    if( res.error ) {
      notifier.showError( res.error );
      return true;
    }
    if( res.feedbacks && res.feedbacks.length ) feedbacks = res.feedbacks;
    if( res.pager && res.pager.page ) pager = res.pager;
    if( res.pager && res.pager.filters && !res.feedbacks ) feedbacks = []; // if nothing found but filter exists, clear table
    if( res.users && res.users.length ) users = res.users;
  }
  
  const userFunc = (/** @type {number|string} */ userId, /** @type {any} */ row ) => {
    if(!users) return userId;
    return (users[ userId ] + ' (' + userId + ')') || 'unknown';
  }
  
  const renderMap = {
    createdBy: userFunc,
    updatedBy: userFunc,
  };
  
  async function refreshTableView(/** @type {PagerIn} */ pagerIn ) {
    await AdminFeedbacks( { // @ts-ignore
      pager: pagerIn,
      cmd: 'list',
    }, function( res ) {
      handleResponse( res );
    } );
  }
  
  async function editRow(/** @type {number|string}*/ id, /** @type {any[] | any}*/ row ) {
    console.log( 'editRow', id, row );
    await AdminFeedbacks( { // @ts-ignore
      feedback: {id},
      cmd: 'form',
    }, function(/** @type {any} */ res ) {
      if( !handleResponse( res ) ) {
        modalForm.showModal( res.feedback );
      }
    } );
  }
  
  async function saveRow( action, row ) {
    let feedback = {...row};
    if( !feedback.id ) feedback.id = '0';
    console.log( feedback );
    try {
      feedback.coord = JSON.parse( '[' + feedback.coord + ']' );
    } catch( e ) {
      feedback.coord = [0, 0];
    }
    await AdminFeedbacks( { // @ts-ignore
      feedback: feedback,
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
        bind:this={modalForm}
        fields={fields}
        onConfirm={saveRow}
        rowType='Feedback'
      >
      </ModalForm>
      <TableView
        bind:pager={pager}
        fields={fields}
        onEditRow={editRow}
        onRefreshTableView={refreshTableView}
        rows={feedbacks || []}
        renderFuncs={renderMap}
      >
        <div class="info">
          <Icon
            src={RiSystemInformation2Line}
            size="16"
            color="var(--blue-005)"
          />
          <span>Editing a feedback's admin's reply will send email to the user.</span>
        </div>
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
    gap: 10px;
  }

  .admin-container .admin-content .info {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
    color: var(--blue-005);
    background-color: var(--blue-transparent);
    border-radius: 8px;
    padding: 9px 15px;
    width: fit-content;
  }
</style>