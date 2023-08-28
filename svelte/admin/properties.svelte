<script>
  // @ts-nocheck
  import Menu from '../_components/Menu.svelte';
  import AdminSubMenu from '../_components/AdminSubMenu.svelte';
  import ProfileHeader from '../_components/ProfileHeader.svelte';
  import Footer from '../_components/Footer.svelte';
  import TableView from '../_components/TableView.svelte';
  import ModalForm from '../_components/ModalForm.svelte';
  import HiSolidEye from 'svelte-icons-pack/hi/HiSolidEye';
  import { AdminProperties, UserPropHistory } from '../jsApi.GEN';
  import { datetime2 } from '_components/formatter';
  
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidPlusCircle from 'svelte-icons-pack/fa/FaSolidPlusCircle';
  import ModalDialog from '../_components/ModalDialog.svelte';
  
  let segments = {/* segments */};
  let fields = [/* fields */];
  let properties = [/* properties */] || [];
  let pager = {/* pager */};
  let currentPropHistory = [];
  let extraActions = [
    {
      icon: HiSolidEye,
      tooltip: 'Show Property History',
      onClick: function( item ) {
        let propertyKey = item[ 1 ]; // property key for taiwan data
        UserPropHistory(
          {
            propertyKey: propertyKey,
          },
          function( res ) {
            if( res.error ) alert( res.error );
            propHistoryModal.showModal();
            currentPropHistory = res.history || [];
            console.log( res.history );
          },
        );
      },
    },
  ];
  
  $: console.log( 'properties=', properties );
  
  let propHistoryModal = ModalDialog;
  
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
    if( res.pager && res.pager.filters && !res.properties ) properties = []; // if nothing found but filter exists, clear table
  }
  
  async function refreshTableView( pagerIn ) {
    // console.log( 'pagerIn=',pagerIn );
    await AdminProperties(
      {
        pager: pagerIn,
        action: 'list',
      },
      function( res ) {
        handleResponse( res );
      },
    );
  }
  
  let form = ModalForm; // for lookup
  
  async function editRow( id, row ) {
    await AdminProperties(
      {
        property: {id},
        action: 'form',
      },
      function( res ) {
        if( !handleResponse( res ) ) form.showModal( res.property );
      },
    );
  }
  
  function addRow() {
    form.showModal( {id: ''} );
  }
  
  async function saveRow( action, row ) {
    let property = {...row};
    if( !property.id ) property.id = '0';
    console.log( property );
    try {
      property.coord = JSON.parse( '[' + property.coord + ']' );
    } catch( e ) {
      property.coord = [0, 0];
    }
    await AdminProperties(
      {
        property: property,
        action: action,
        pager: pager, // force refresh page, will be slow
      },
      function( res ) {
        if( handleResponse( res ) ) {
          return form.setLoading( false ); // has error
        }
        form.hideModal(); // success
      },
    );
  }
</script>

<section class='dashboard'>
  <Menu access={segments} />
  <div class='dashboard_main_content'>
    <ProfileHeader />
    <AdminSubMenu />
    <div class='content'>
      <ModalForm {fields} rowType='Property' bind:this={form} onConfirm={saveRow} />
      <section class='tableview_container'>
        <TableView
          {fields}
          {extraActions}
          bind:pager
          rows={properties}
          widths={{
            mainUse: '320px',
            address: '240px',
          }}
          onRefreshTableView={refreshTableView}
          onEditRow={editRow}
        >
          <button on:click={addRow} class='add_button'>
            <Icon size={18} color='#FFFF' src={FaSolidPlusCircle} />
            <span>Add</span>
          </button>
        </TableView>
      </section>
    </div>
    <Footer />
  </div>
  <ModalDialog bind:this={propHistoryModal} title='Property History'>
    <div slot='content'>
      {#if currentPropHistory && currentPropHistory.length}
        {#each currentPropHistory as row}
          <div class='prop_item'>
            <div class='row'>
              <p class='name'>Property ID</p><span>:</span>
              <p class='value'>{row.id}</p>
            </div>
            <div class='row'>
              <p class='name'>Created At</p><span>:</span>
              <p class='value'>{datetime2( row.createdAt ) || '-'}</p>
            </div>
            <div class='row'>
              <p class='name'>Created By</p><span>:</span>
              <p class='value'>{row.createdBy}</p>
            </div>
            <div class='row'>
              <p class='name'>Deleted At</p><span>:</span>
              <p class='value'>{datetime2( row.deletedAt ) || '-'}</p>
            </div>
            <div class='row'>
              <p class='name'>Updated At</p><span>:</span>
              <p class='value'>{datetime2( row.updatedAt ) || '-'}</p>
            </div>
            <div class='row'>
              <p class='name'>Updated By</p><span>:</span>
              <p class='value'>{datetime2( row.updatedBy )}</p>
            </div>
            <div class='row'>
              <p class='name'>Address</p><span>:</span>
              <p class='value'>{row.address}</p>
            </div>
            <div class='row'>
              <p class='name'>District</p><span>:</span>
              <p class='value'>{row.district}</p>
            </div>
            <div class='row'>
              <p class='name'>Price</p><span>:</span>
              <p class='value'>{row.price}</p>
            </div>
            <div class='row'>
              <p class='name'>Price NTD</p><span>:</span>
              <p class='value'>{row.priceNtd}</p>
            </div>
            <div class='row'>
              <p class='name'>Price per Unit</p><span>:</span>
              <p class='value'>{row.pricePerUnit}</p>
            </div>
            <div class='row'>
              <p class='name'>Property Key</p><span>:</span>
              <p class='value'>{row.propertyKey}</p>
            </div>
            <div class='row'>
              <p class='name'>Serial Number</p><span>:</span>
              <p class='value'>{row.serialNumber}</p>
            </div>
            <div class='row'>
              <p class='name'>Transaction date normal</p><span>:</span>
              <p class='value'>{row.transactionDateNormal || '-'}</p>
            </div>
            <div class='row'>
              <p class='name'>Transaction Key</p><span>:</span>
              <p class='value'>{row.transactionKey}</p>
            </div>
            <div class='row'>
              <p class='name'>Transaction Number</p><span>:</span>
              <p class='value'>{row.transactionNumber}</p>
            </div>
            <div class='row'>
              <p class='name'>Transaction Sign</p><span>:</span>
              <p class='value'>{row.transactionSign}</p>
            </div>
            <div class='row'>
              <p class='name'>Transaction Time</p><span>:</span>
              <p class='value'>{datetime2( row.transactionTime )}</p>
            </div>
            <div class='row'>
              <p class='name'>Transaction Type</p><span>:</span>
              <p class='value'>{row.transactionType}</p>
            </div>
          </div>
        {/each}
      {:else}
        no history for this property
      {/if}
    </div>
  </ModalDialog>
</section>

<style>
  .prop_item {
    display          : flex;
    flex-direction   : column;
    gap              : 5px;
    padding          : 8px;
    border           : 1px solid #CBD5E1;
    background-color : #F1F5F9;
    border-radius    : 8px;
  }

  .prop_item .row {
    display        : flex;
    flex-direction : row;
    align-items    : center;
    gap            : 4px;
  }

  .prop_item .row p {
    margin : 0;
  }

  .prop_item .row .name {
    width : 35%;
  }

  .prop_item .row .name,
  .prop_item .row span {
    font-weight : 700;
  }

  .prop_item .row .value {
    flex-grow : 1;
  }
</style>
