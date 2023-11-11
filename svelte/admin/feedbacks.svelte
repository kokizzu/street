<script>
  // @ts-nocheck
  import Menu from '../_components/Menu.svelte';
  import AdminSubMenu from '../_components/AdminSubMenu.svelte';
  import ProfileHeader from '../_components/ProfileHeader.svelte';
  import Footer from '../_components/Footer.svelte';
  import TableView from '../_components/TableView.svelte';
  import ModalForm from '../_components/ModalForm.svelte';
  import { AdminFeedbacks } from '../jsApi.GEN';
  import { fieldsArrToMap } from '../_components/mapper.js';
  
  let segments = {/* segments */};
  let fields = [/* fields */];
  let fieldByKey = fieldsArrToMap( fields );
  let feedbacks = [/* feedbacks */] || [];
  let pager = {/* pager */};
  let users = {/* users */} || {};
  $: console.log( 'feedbacks=', feedbacks );
  
  // return true if got error
  function handleResponse( res ) {
    if( res.error ) {
      alert( res.error );
      return true;
    }
    if( res.feedbacks && res.feedbacks.length ) feedbacks = res.feedbacks;
    if( res.pager && res.pager.page ) pager = res.pager;
    if( res.pager && res.pager.filters && !res.feedbacks ) feedbacks = []; // if nothing found but filter exists, clear table
    if( res.users && res.users.length ) users = res.users;
  }
  
  const userFunc = ( userId, row ) => {
    if(!users) return userId;
    return (users[ userId ] + ' (' + userId + ')') || 'unknown';
  }
  
  const renderMap = {
    createdBy: userFunc,
    updatedBy: userFunc,
  };
  
  async function refreshTableView( pagerIn ) {
    // console.log( 'pagerIn=',pagerIn );
    await AdminFeedbacks( {
      pager: pagerIn,
      cmd: 'list',
    }, function( res ) {
      handleResponse( res );
    } );
  }
  
  let form = ModalForm; // for lookup
  
  async function editRow( id, row ) {
    console.log( 'editRow', id, row );
    await AdminFeedbacks( {
      feedback: {id},
      cmd: 'form',
    }, function( res ) {
      if( !handleResponse( res ) )
        form.showModal( res.feedback );
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
    await AdminFeedbacks( {
      feedback: feedback,
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


<section class='dashboard'>
  <Menu access={segments} />
  <div class='dashboard_main_content'>
    <ProfileHeader access={segments} />
    <AdminSubMenu></AdminSubMenu>
    <div class='content'>
      <ModalForm
        bind:this={form}
        fields={fields}
        onConfirm={saveRow}
        rowType='Feedback'
      >
      </ModalForm>
      <section class='tableview_container'>
        <TableView
          bind:pager={pager}
          fields={fields}
          onEditRow={editRow}
          onRefreshTableView={refreshTableView}
          rows={feedbacks || []}
          renderFuncs={renderMap}
          widths={{mainUse: '320px', address: '240px'}}
        >
          Editing a feedback's admin's reply will send email to the user.
        </TableView>
      </section>
    </div>
    <Footer></Footer>
  </div>
</section>

<style>
  .action_btn.not_filtered {
    background-color : #6366F1;
    color            : white;
    flex-direction   : row;
  }

  .action_btn.not_filtered:hover {
    background-color : #7E80F1;
  }
</style>
