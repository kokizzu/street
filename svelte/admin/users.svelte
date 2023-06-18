<script>
  import Menu from '../_components/menu.svelte';
  import AdminSubMenu from './_adminSubMenu.svelte';
  import TableView from '../_components/tableView.svelte';
  import { AdminUsers } from '../jsApi.GEN';
  
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

</script>
<Menu access={segments} />
<AdminSubMenu></AdminSubMenu>
<TableView {fields} {pager} rows={users}
           on:refreshTableView={refreshTableView}
></TableView>