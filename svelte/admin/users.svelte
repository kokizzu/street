<script>
  // @ts-nocheck
  import Menu from '../_components/Menu.svelte';
  import AdminSubMenu from '../_components/AdminSubMenu.svelte';
  import ProfileHeader from '../_components/ProfileHeader.svelte';
  import Footer from '../_components/Footer.svelte';
  import TableView from '../_components/TableView.svelte';
  import { AdminUsers } from '../jsApi.GEN';
  import ModalForm from '../_components/ModalForm.svelte';

  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidPlusCircle from "svelte-icons-pack/fa/FaSolidPlusCircle";
  
  let segments = {/* segments */};
  let fields = [/* fields */];
  let users = [/* users */];
  let pager = {/* pager */};
  
  // $: console.log( users, fields, pager );
  
  // return true if got error
  function handleResponse( res ) {
    console.log( res );
    if( res.error ) {
      // TODO: growl error
      alert( res.error );
      return true;
    }
    if( res.users && res.users.length ) users = res.users;
    if( res.pager && res.pager.page ) pager = res.pager;
  }
  
  async function refreshTableView( pagerIn ) {
    // console.log( 'pagerIn=',pagerIn );
    await AdminUsers( {
      pager: pagerIn,
      cmd: 'list',
    }, function( res ) {
      handleResponse( res );
    } );
  }
  
  let form = ModalForm; // for lookup
  
  async function editRow( id, row ) {
    await AdminUsers( {
      user: {id},
      cmd: 'form',
    }, function( res ) {
      if( !handleResponse( res ) )
        form.showModal( res.user );
    } );
  }
  
  function addRow() {
    form.showModal( {id: ''} );
  }
  
  async function saveRow( action, row ) {
    let user = {...row};
    if( !user.id ) user.id = '0';
    await AdminUsers( {
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

<section class='dashboard'>
  <Menu access={segments} />
  <div class='dashboard_main_content'>
    <ProfileHeader></ProfileHeader>
    <AdminSubMenu></AdminSubMenu>
    <div class='content'>
      <ModalForm {fields}
                 rowType='User'
                 bind:this={form}
                 onConfirm={saveRow}
      ></ModalForm>
      <section class='tableview_container'>
        <TableView {fields}
                   bind:pager={pager}
                   rows={users}
                   onRefreshTableView={refreshTableView}
                   onEditRow={editRow}
        >
          <button on:click={addRow} class='add_button'>
            <Icon size={18} color="#FFFF" src={FaSolidPlusCircle} />
            <span>Add</span>
          </button>
        </TableView>
      </section>
    </div>
    <Footer></Footer>
  </div>
</section>

<style>
</style>