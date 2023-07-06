<script>
  import Menu from '../_components/Menu.svelte';
  import AdminSubMenu from './_adminSubMenu.svelte';
  import TableView from '../_components/TableView.svelte';
  import ModalForm from '../_components/ModalForm.svelte';
  import { AdminPropHistories } from '../jsApi.GEN';
  
  let segments = {/* segments */};
  let fields = [/* fields */];
  let propHistories = [/* propHistories */] || [];
  let pager = {/* pager */};
  
  $: console.log( propHistories );
  
  // return true if got error
  function handleResponse( res ) {
    console.log( res );
    if( res.error ) {
      // TODO: growl error
      alert( res.error );
      return true;
    }
    if( res.propHistories && res.propHistories.length ) propHistories = res.propHistories;
    if( res.pager && res.pager.page ) pager = res.pager;
  }
  
  async function refreshTableView( e ) {
    const pagerIn = e.detail;
    // console.log( pager );
    await AdminPropHistories( {
      pager: pagerIn,
      action: 'list',
    }, function( res ) {
      handleResponse( res );
    } );
  }
  
  let form = ModalForm; // for lookup
  
  async function editRow( e ) {
    const id = e.detail;
    await AdminPropHistories( {
      propHistory: {id},
      action: 'form',
    }, function( res ) {
      if( !handleResponse( res ) )
        form.showModal( res.propHistory );
    } );
  }
  
  function addRow() {
    form.showModal( {id: ''} );
  }
  
  async function saveRow( action, row ) {
    let ph = {...row};
    if( !ph.id ) ph.id = '0';
    console.log(ph)
    await AdminPropHistories( {
      propHistory: ph,
      action: action,
      pager: pager, // force refresh page, will be slow
    }, function( res ) {
      if( handleResponse( res ) ) {
        return form.setLoading( false ); // has error
      }
      form.hideModal(); // success
    } );
  }

</script>


<Menu access={segments} />
<AdminSubMenu></AdminSubMenu>
<button on:click={addRow}>Add</button>
<ModalForm {fields}
           rowType='Prop History'
           bind:this={form}
           onConfirm={saveRow}
></ModalForm>
<TableView {fields}
           bind:pager={pager}
           rows={propHistories}
           on:refreshTableView={refreshTableView}
           on:editRow={editRow}
></TableView>