<script>
  import Menu from '../_components/Menu.svelte';
  import AdminSubMenu from './_adminSubMenu.svelte';
  import TableView from '../_components/TableView.svelte';
  import ModalForm from '../_components/ModalForm.svelte';
  import { AdminProperties } from '../jsApi.GEN';
  
  let segments = {/* segments */};
  let fields = [/* fields */];
  let properties = [/* properties */] || [];
  let pager = {/* pager */};
  
  $: console.log( properties );
  
  // return true if got error
  function handleResponse( res ) {
    console.log( res );
    if( res.error ) {
      // TODO: growl error
      alert( res.error );
      return true;
    }
    if( res.properties && res.properties.length ) properties = res.properties;
    if( res.pager && res.pager.page ) pager = res.pager;
  }
  
  async function refreshTableView( e ) {
    const pagerIn = e.detail;
    // console.log( pager );
    await AdminProperties( {
      pager: pagerIn,
      action: 'list',
    }, function( res ) {
      handleResponse( res );
    } );
  }
  
  let form = ModalForm; // for lookup
  
  async function editRow( e ) {
    const id = e.detail;
    await AdminProperties( {
      property: {id},
      action: 'form',
    }, function( res ) {
      if( !handleResponse( res ) )
        form.showModal( res.property );
    } );
  }
  
  function addRow() {
    form.showModal( {id: ''} );
  }
  
  async function saveRow( action, row ) {
    let property = {...row};
    if( !property.id ) property.id = '0';
    console.log(property)
    property.coord = JSON.parse( '[' + property.coord + ']' );
    await AdminProperties( {
      property: property,
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
           rowType='User'
           bind:this={form}
           onConfirm={saveRow}
></ModalForm>
<TableView {fields}
           bind:pager={pager}
           rows={properties}
           on:refreshTableView={refreshTableView}
           on:editRow={editRow}
></TableView>