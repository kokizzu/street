<script>
  import Menu from '../_components/Menu.svelte';
  import AdminSubMenu from './_adminSubMenu.svelte';
  import TableView from '../_components/TableView.svelte';
  import { AdminUsers } from '../jsApi.GEN';
  import ModalForm from '../_components/ModalForm.svelte';
  
  let segments = {/* segments */ };
  let fields = [/* fields */];
  let users = [/* users */];
  let pager = {/* pager */ };
  
  $: console.log( users, fields, pager );
  
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
  
  async function refreshTableView( e ) {
    const pagerIn = e.detail;
    // console.log( pager );
    await AdminUsers( {
      pager: pagerIn,
      action: 'list',
    }, function( res ) {
      handleResponse( res );
    } );
  }
  
  let form = ModalForm; // for lookup
  
  async function editRow( e ) {
    const id = e.detail;
    await AdminUsers( {
      user: { id },
      action: 'form',
    }, function( res ) {
      if( !handleResponse( res ) )
        form.showModal( res.user );
    } );
  }
  
  async function saveRow( action, row ) {
    await AdminUsers( {
      user: row,
      action: action,
      pager: pager, // force refresh page, will be slow
    }, function( res ) {
      if( handleResponse( res ) ) {
        return form.setLoading(false); // has error
      }
      form.hideModal(); // success
    } );
  }


</script>
<Menu access={segments} />
<AdminSubMenu></AdminSubMenu>
<ModalForm {fields}
           bind:this={form}
           onConfirm={saveRow}
></ModalForm>
<TableView {fields}
           bind:pager={pager}
           rows={users}
           on:refreshTableView={refreshTableView}
           on:editRow={editRow}
></TableView>