<script>
  import Menu from '../_components/Menu.svelte';
  import AdminSubMenu from '../_components/AdminSubMenu.svelte';
  import ProfileHeader from '../_components/ProfileHeader.svelte';
  import Footer from '../_components/partials/Footer.svelte';
  import TableView from '../_components/TableView.svelte';
  import ModalForm from '../_components/ModalForm.svelte';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import {
    HiSolidEye, HiSolidXCircle, HiSolidCheckCircle, HiOutlineLink
  } from '../node_modules/svelte-icons-pack/dist/hi';
  import {
    FaSolidCirclePlus, FaSolidCheckDouble, FaSolidRecycle
  } from '../node_modules/svelte-icons-pack/dist/fa';
  import { AdminProperties, UserPropHistory } from '../jsApi.GEN';
  import {notifier} from '../_components/notifier.js';

  import ModalDialog from '../_components/ModalDialog.svelte';
  import PillBox from '../_components/PillBox.svelte';
  import { priceNtd } from '../_components/formatter.js';
  import { fieldsArrToMap } from '../_components/mapper.js';
  
  let segments = {/* segments */};
  let fields = [/* fields */];
  let fieldByKey = fieldsArrToMap( fields );
  let properties = [/* properties */];
  console.log('Properties', properties);

  let pager = {/* pager */};
  let currentPropHistory = [];
  let extraActions = [
    {
      icon: HiSolidEye,
      label: 'Show Property History',
      onClick: function( item ) {
        let propertyKey = item[ fieldByKey[ 'uniqPropKey' ].idx ]; // property key for taiwan data
        UserPropHistory( {
          propertyKey: propertyKey,
        }, function( res ) {
          if( res.error ) alert( res.error );
          propHistoryModal.showModal();
          currentPropHistory = res.history || [];
        } );
      },
    },
    {
      icon: HiSolidCheckCircle,
      showIf: function( row ) {
        return !!row[ fieldByKey[ 'approvalState' ].idx ];
      },
      label: 'approve property',
      onClick: function( item ) {
        const id = item[ 0 ];
        AdminProperties( {
            cmd: 'upsert',
            property: {id: id, approvalState: ' '}, // empty is approved, will be trimmed on server side
            pager: pager,
          },
          function( res ) {
            if( res.error ) alert( res.error );
          } );
      },
    },
    {
      icon: HiSolidXCircle,
      showIf: function( row ) {
        return !row[ fieldByKey[ 'approvalState' ].idx ];
      },
      label: 'reject property',
      onClick: function( item ) {
        const id = item[ 0 ];
        const reason = prompt( 'input refusal reason for #' + id + ', use "pending" to set state to pending' );
        if( !reason ) return;
        AdminProperties( {
          cmd: 'upsert',
          property: {id: id, approvalState: reason},
        }, function( res ) {
          if( res.error ) alert( res.error );
        } );
      },
    },
    {
      icon: HiOutlineLink,
      label: 'Go to property page',
      link: function( item ) {
        return '/guest/property/' + item[ fieldByKey[ 'id' ].idx ];
      },
    },
  ];
  
  let propHistoryModal = ModalDialog;
  
  // return true if got error
  function handleResponse( res ) {
    if( res.error ) {
      notifier.showError( res.error );
      return true;
    }
    if( res.properties && res.properties.length ) properties = res.properties;
    if( res.pager && res.pager.page ) pager = res.pager;
    if( res.pager && res.pager.filters && !res.properties ) properties = []; // if nothing found but filter exists, clear table
  }
  
  async function refreshTableView( pagerIn ) {
    // console.log( 'pagerIn=',pagerIn );
    await AdminProperties( {
      pager: pagerIn,
      cmd: 'list',
    }, function( res ) {
      handleResponse( res );
    } );
  }
  
  let form = ModalForm; // for lookup
  
  async function editRow( id, row ) {
    await AdminProperties( {
      property: {id},
      cmd: 'form',
    }, function( res ) {
      if( !handleResponse( res ) )
        form.showModal( res.property );
    } );
  }
  
  let filterdByPending = false, filterdByRejected = false, filter = [];
  $: approvalStateFilter = (pager.filters || {}).approvalState || [];
  $: filterdByPending = approvalStateFilter[ 0 ]==='pending';
  $: filterdByRejected = approvalStateFilter[ 0 ]==='<>' && approvalStateFilter[ 1 ]==='<>pending';
  
  async function filterPendingApproval() {
    if( !pager.filters ) pager.filters = {};
    if( filterdByPending ) {
      delete pager.filters[ 'approvalState' ];
    } else {
      pager.filters.approvalState = ['pending'];
    }
    await AdminProperties( {
      cmd: 'list',
      pager: pager,
    }, handleResponse );
  }
  
  async function filterRejectedApproval() {
    if( !pager.filters ) pager.filters = {};
    if( filterdByRejected ) {
      delete pager.filters[ 'approvalState' ];
    } else {
      pager.filters.approvalState = ['<>', '<>pending'];
    }
    await AdminProperties( {
      cmd: 'list',
      pager: pager,
    }, handleResponse );
  }
  
  function addRow() {
    form.showModal( {id: ''} );
  }
  
  async function saveRow( action, row ) {
    let property = {...row};
    if( !property.id ) property.id = '0';
    try {
      property.coord = JSON.parse( '[' + property.coord + ']' );
    } catch( e ) {
      property.coord = [0, 0];
    }
    await AdminProperties( {
      property: property,
      cmd: action,
      pager: pager, // force refresh page, will be slow
    }, function( res ) {
      if( handleResponse( res ) ) {
        return form.setLoading( false ); // has error
      }
      form.hideModal(); // success
    } );
  }
  
  let renderMap = {
    note: function( note ) {
      return JSON.stringify( note );
    },
  };
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
        rowType='Property'
      >
      </ModalForm>
      <section class='tableview_container'>
        <TableView
          bind:pager={pager}
          extraActions={extraActions}
          fields={fields}
          onEditRow={editRow}
          onRefreshTableView={refreshTableView}
          rows={properties||[]}
          renderFuncs={renderMap}
          widths={{mainUse: '320px', address: '240px'}}
        >
          <button class='action_btn' on:click={addRow}>
            <Icon color='#FFF' size={17} src={FaSolidCirclePlus} />
            <span>Add</span>
          </button>
          <button
            class='action_btn'
            class:not_filtered={!filterdByPending}
            on:click={filterPendingApproval}
          >
            <Icon color='{!filterdByPending ? "#FFF" : "#000"}' size={17} src={FaSolidCheckDouble} />
            <span>Filter Pending</span>
          </button>
          <button
            class='action_btn'
            class:not_filtered={!filterdByRejected}
            on:click={filterRejectedApproval}
          >
            <Icon color='{!filterdByRejected ? "#FFF" : "#000"}' size={17} src={FaSolidRecycle} />
            <span>Filter Rejected</span>
          </button>
        </TableView>
      </section>
    </div>
    <Footer></Footer>
  </div>
  <ModalDialog bind:this={propHistoryModal}>
    <div slot='content'>
      <h3>Property History</h3>
      {#if currentPropHistory && currentPropHistory.length}
        {#each currentPropHistory as row}
          {#each Object.entries( row ) as [key, val]}
            {#if val==='0' || !val}
              &nbsp;
            {:else if key==='createdAt' || key==='updatedAt'}
              <PillBox label={key} content={new Date(val*1000)} />
            {:else if key==='priceNtd' || key==='pricePerUnit'}
              <PillBox label={key} content={priceNtd(val)} />
            {:else}
              <PillBox label={key} content={val} />
            {/if}
          {/each}
        {/each}
      {:else}
        no history for this property
      {/if}
    </div>
  </ModalDialog>
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