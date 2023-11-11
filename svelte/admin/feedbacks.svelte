<script>
	// @ts-nocheck
	import Menu from '../_components/Menu.svelte';
	import AdminSubMenu from '../_components/AdminSubMenu.svelte';
	import ProfileHeader from '../_components/ProfileHeader.svelte';
	import Footer from '../_components/Footer.svelte';
	import TableView from '../_components/TableView.svelte';
	import ModalForm from '../_components/ModalForm.svelte';
	import HiSolidEye from 'svelte-icons-pack/hi/HiSolidEye';
	import HiSolidXCircle from 'svelte-icons-pack/hi/HiSolidXCircle';
	import HiSolidCheckCircle from 'svelte-icons-pack/hi/HiSolidCheckCircle';
	import { AdminFeedbacks, AdminProperties, UserPropHistory } from '../jsApi.GEN';
	
	import Icon from 'svelte-icons-pack/Icon.svelte';
	import FaSolidPlusCircle from 'svelte-icons-pack/fa/FaSolidPlusCircle';
	import HiOutlineLink from 'svelte-icons-pack/hi/HiOutlineLink';
	
	import ModalDialog from '../_components/ModalDialog.svelte';
	import { fieldsArrToMap } from '../_components/mapper.js';
	
	let segments = {/* segments */};
  let fields = [/* fields */];
  let fieldByKey = fieldsArrToMap( fields );
  let feedbacks = [/* feedbacks */] || [];
  let pager = {/* pager */};
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
  }
  
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
    console.log('editRow', id, row)
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
          widths={{mainUse: '320px', address: '240px'}}
        >
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
