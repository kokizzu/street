<script>
  import Menu from '../_components/Menu.svelte';
  import AdminSubMenu from './_adminSubMenu.svelte';
  import TableView from '../_components/TableView.svelte';
  import { AdminUsers } from '../jsApi.GEN';
  import MasterForm from '../_components/MasterForm.svelte';
  
  let segments = {/* segments */ };
  let fields = [/* fields */];
  let users = [/* users */];
  let pager = {/* pager */ };
  
  $: console.log( users, fields, pager );
  
  async function refreshTableView( e ) {
    const pagerIn = e.detail;
    // console.log( pager );
    await AdminUsers( {
      pager: pagerIn,
      action: 'list',
    }, function( res ) {
      // console.log( res );
      users = res.users;
      pager = res.pager;
    } );
  }
  
  let masterForm = MasterForm; // for lookup
  
  async function editRow( e ) {
    const id = e.detail;
    await AdminUsers( {
      user: { id },
      action: 'form',
    }, function( res ) {
      console.log( res );
      masterForm.showDialog(res)
    } );
  }


</script>
<Menu access={segments} />
<AdminSubMenu></AdminSubMenu>
<MasterForm {fields}
            bind:this={masterForm}
></MasterForm>
<TableView {fields}
           bind:pager={pager}
           rows={users}
           on:refreshTableView={refreshTableView}
           on:editRow={editRow}
></TableView>