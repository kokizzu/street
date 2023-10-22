<script>
  // @ts-nocheck
  import Menu from '../_components/Menu.svelte';
  import AdminSubMenu from '../_components/AdminSubMenu.svelte';
  import TableView from '../_components/TableView.svelte';
  import ModalForm from '../_components/ModalForm.svelte';
  import { AdminPropHistories } from '../jsApi.GEN';
  import Footer from '../_components/Footer.svelte';
  import ProfileHeader from '../_components/ProfileHeader.svelte';
  import Growl from '../_components/Growl.svelte';

  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidPlusCircle from "svelte-icons-pack/fa/FaSolidPlusCircle";
  
  let segments = {/* segments */};
  let fields = [/* fields */];
  let propHistories = [/* propHistories */] || [];
  let pager = {/* pager */};
  
  $: console.log( propHistories );

  let growl = Growl;
  // return true if got error
  function handleResponse( res ) {
    console.log( res );
    if( res.error ) {
      growl.showError( res.error );
      return true;
    }
    if( res.propHistories && res.propHistories.length ) propHistories = res.propHistories;
    if( res.pager && res.pager.page ) pager = res.pager;
  }
  
  async function refreshTableView( pagerIn ) {
    // console.log( 'pagerIn=',pagerIn );
    await AdminPropHistories( {
      pager: pagerIn,
      cmd: 'list',
    }, function( res ) {
      handleResponse( res );
    } );
  }
  
  let form = ModalForm; // for lookup
  
  async function editRow(id, row) {
    await AdminPropHistories( {
      propHistory: {id},
      cmd: 'form',
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
    console.log( ph );
    await AdminPropHistories( {
      propHistory: ph,
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

<Growl bind:this={growl} />
<section class='dashboard'>
  <Menu access={segments} />
  <div class='dashboard_main_content'>
    <ProfileHeader />
    <AdminSubMenu></AdminSubMenu>
    <div class='content'>
      <ModalForm {fields}
                 rowType='Prop History'
                 bind:this={form}
                 onConfirm={saveRow}
      ></ModalForm>
      <section class='tableview_container'>
        <TableView {fields}
                   bind:pager={pager}
                   rows={propHistories}
                   onRefreshTableView={refreshTableView}
                   onEditRow={editRow}
        >
          <button on:click={addRow} class='action_btn'>
            <Icon size={17} color='#FFF' src={FaSolidPlusCircle} />
            <span>Add</span>
          </button>
        </TableView>
      </section>
    </div>
    <Footer></Footer>
  </div>
</section>
