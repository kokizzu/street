<script>
   import Menu from '../_components/Menu.svelte';
   import AdminSubMenu from './_adminSubMenu.svelte';
   import ProfileHeader from '../_components/ProfileHeader.svelte';
   import Footer from '../_components/Footer.svelte';
   import TableView from '../_components/TableView.svelte';
   import { AdminUsers } from '../jsApi.GEN';
   import ModalForm from '../_components/ModalForm.svelte';

   let sideMenuOpen = false;
   function openSideMenu() {
      sideMenuOpen = true;
   }
   function closeSideMenu() {
      sideMenuOpen = false;
   }
  
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
         user: {id},
         action: 'form',
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

<section class="dashboard">
   <Menu
      access={segments}
      isSideMenuOpen={sideMenuOpen}
      on:closesidemenu={closeSideMenu}
   />
   <div class="dashboard_main_content">
      <ProfileHeader on:opensidemenu={openSideMenu}></ProfileHeader>
      <AdminSubMenu></AdminSubMenu>
      <div class="content">
         <ModalForm {fields}
            rowType='User'
            bind:this={form}
            onConfirm={saveRow}
         ></ModalForm>
         <section class="users_dashboard">
            <button on:click={addRow} class="add_button">
               <i class="gg-add"></i>
               <span>Add</span>
            </button>
            <TableView {fields}
               bind:pager={pager}
               rows={users}
               on:refreshTableView={refreshTableView}
               on:editRow={editRow}
            ></TableView>
         </section>
      </div>
      <Footer></Footer>
   </div>
</section>

<style>
   .users_dashboard {
      position: relative;
      z-index: 30;
      margin-top: -40px;
      margin-left: auto;
      margin-right: auto;
      border-radius: 8px;
      filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
      padding: 20px;
      background-color: white;
      width: 88%;
      height: fit-content;
      color: #475569;
      font-size: 16px;
   }
   .users_dashboard .add_button {
      padding: 8px 20px;
      font-size: 16px;
      display: inline-flex;
      flex-direction: row;
      align-items: center;
      align-content: center;
      justify-content: center;
      border: none;
      background-color: #6366F1;
      border-radius: 8px;
      color: white;
      cursor: pointer;
   }
   .users_dashboard .add_button:hover {
      background-color: #7E80F1;
   }
   .users_dashboard .add_button i {
      margin-right: 8px;
   }
</style>