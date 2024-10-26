<script>
  /** @typedef {import('../_types/master').Access} Access */
  /** @typedef {import('../_types/property').Property} TypeProperty */
  /** @typedef {import('../_types/property').PropertyUS} TypePropertyUS */
  /** @typedef {import('../_types/property').PropertyExtraUS} TypePropertyExtraUS */
  /** @typedef {import('../_types/master').PagerOut} PagerOut */
  /** @typedef {import('../_types/master').Field} Field */

  import Menu from '../_components/Menu.svelte';
  import AdminSubMenu from '../_components/AdminSubMenu.svelte';
  import ProfileHeader from '../_components/ProfileHeader.svelte';
  import Footer from '../_components/partials/Footer.svelte';
  import TableView from '../_components/TableView.svelte';
  import ModalForm from '../_components/ModalForm.svelte';
  import HiSolidEye from 'svelte-icons-pack/hi/HiSolidEye';
  import HiSolidXCircle from 'svelte-icons-pack/hi/HiSolidXCircle';
  import HiSolidCheckCircle from 'svelte-icons-pack/hi/HiSolidCheckCircle';
  import { AdminPropertiesUS, UserPropHistory } from '../jsApi.GEN';
  import {notifier} from '../_components/notifier.js';
  
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidPlusCircle from 'svelte-icons-pack/fa/FaSolidPlusCircle';
  import FaSolidCheckDouble from 'svelte-icons-pack/fa/FaSolidCheckDouble';
  import FaSolidRecycle from 'svelte-icons-pack/fa/FaSolidRecycle';
  import HiOutlineLink from 'svelte-icons-pack/hi/HiOutlineLink';
  
  import ModalDialog from '../_components/ModalDialog.svelte';
  import PillBox from '../_components/PillBox.svelte';
  import { priceNtd } from '../_components/formatter.js';
  import { fieldsArrToMap } from '../_components/mapper.js';
  import AccessLog from './accessLog.svelte';
  import AddFloorDialog from '../_components/AddFloorDialog.svelte';
  import AddOrEditRoomDialog from '../_components/AddOrEditRoomDialog.svelte';
  import AddOtherFeesDialog from '../_components/AddOtherFeesDialog.svelte';
  import Admin from '../admin.svelte';
  import Buyer from '../buyer.svelte';
  
  let segments    = /** @type {Access} */ ({/* segments */});
  let fields      = /** @type {Field[]} */ ([/* fields */]);
  let fieldByKey  = /** @type {Object.<string, any>}*/ fieldsArrToMap( fields );
  let properties  = /** @type {TypeProperty[] | TypePropertyUS[]} */ ([/* properties */]);
  let pager       = /** @type {PagerOut} */ ({/* pager */});

  /** @type {import('svelte').SvelteComponent}*/
  let propHistoryModal;

  let currentPropHistory = [];
  let extraActions = [
    {
      icon: HiSolidEye,
      label: 'Show Property History',
      onClick: function( item ) {
        let propertyKey = item[ fieldByKey[ 'uniqPropKey' ].idx ]; // property key for taiwan data
        UserPropHistory( {
          propertyKey: propertyKey,
        }, // @ts-ignore
        function( /** @type {any}*/ res ) {
          if( res.error ) {
            notifier.showError( res.error );
            return;
          }
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
        const adminPropUSIn = /** @type {import('../jsApi.GEN').AdminPropertiesUSIn | any} */ ({
          cmd: 'upsert',
          property: { id, approvalState: ' ' }, // empty is approved, will be trimmed on server side
          pager: pager,
        })
        AdminPropertiesUS( adminPropUSIn, // @ts-ignore
        function( /** @type {any}*/ res ) {
          if( res.error ) {
            notifier.showError( res.error );
            return;
          }
          refreshTableView( pager );
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

        const adminPropUSIn = /** @type {import('../jsApi.GEN').AdminPropertiesUSIn | any} */ ({
          cmd: 'upsert',
          property: { id, approvalState: reason },
          pager: pager,
        })
        AdminPropertiesUS( adminPropUSIn, // @ts-ignore
        function( /** @type {any}*/ res ) {
          if( res.error ) {
            notifier.showError( res.error );
            return;
          }
          refreshTableView( pager );
        } );
      },
    },
    {
      icon: HiOutlineLink,
      label: 'Go to property page',
      link: function( item ) {
        return '/guest/property/US' + item[ fieldByKey[ 'id' ].idx ];
      },
    },
  ];
  
  $: console.log( 'properties=', properties );
  
  // return true if got error
  function handleResponse( res ) {
    console.log( res );
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

    const adminPropUSIn = /** @type {import('../jsApi.GEN').AdminPropertiesUSIn | any} */ ({
      cmd: 'list',
      pager: pagerIn,
    })
    await AdminPropertiesUS( adminPropUSIn, // @ts-ignore
    function( res ) {
      handleResponse( res );
    } );
  }
  
  /** @type {import('svelte').SvelteComponent} */
  let form; // for lookup
  
  async function editRow( id, row ) {
    const adminPropUSIn = /** @type {import('../jsApi.GEN').AdminPropertiesUSIn | any} */ ({
      cmd: 'form',
      property: { id },
    })
    await AdminPropertiesUS( adminPropUSIn, // @ts-ignore
    function( /** @type {any}*/ res ) {
      if( !handleResponse( res ) ) {
        form.showModal( res.property )
      };
    } );
  }
  
  let filterdByPending = false, filterdByRejected = false, filter = [];
  $: approvalStateFilter = (pager.filters || {}).approvalState || []
  $: filterdByPending = approvalStateFilter[0] === 'pending';
  $: filterdByRejected = approvalStateFilter[0] === '<>' && approvalStateFilter[1] === '<>pending'
  
  async function filterPendingApproval() {
    if( !pager.filters ) pager.filters = {};
    if( filterdByPending ) {
      delete pager.filters[ 'approvalState' ];
    } else {
      pager.filters.approvalState = ['pending'];
    }

    const adminPropUSIn = /** @type {import('../jsApi.GEN').AdminPropertiesUSIn | any} */ ({
      cmd: 'list',
      pager: pager,
    }) // @ts-ignore
    await AdminPropertiesUS( adminPropUSIn, handleResponse );
  }
  
  async function filterRejectedApproval() {
    if(!pager.filters) pager.filters = {};
    if( filterdByRejected ) {
      delete pager.filters[ 'approvalState' ];
    } else {
      pager.filters.approvalState = ['<>', '<>pending'];
    }

    const adminPropUSIn = /** @type {import('../jsApi.GEN').AdminPropertiesUSIn | any} */ ({
      cmd: 'list',
      pager: pager,
    }) // @ts-ignore
    await AdminPropertiesUS( adminPropUSIn, handleResponse );
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

    const adminPropUSIn = /** @type {import('../jsApi.GEN').AdminPropertiesUSIn | any} */ ({
      cmd: action,
      property: property,
      pager: pager, // force refresh page, will be slow
    })
    await AdminPropertiesUS( adminPropUSIn, function( res ) {
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
    <ProfileHeader access={segments}/>
    <AdminSubMenu></AdminSubMenu>
    <div class='content'>
      <ModalForm fields={fields}
                 rowType='Property'
                 bind:this={form}
                 onConfirm={saveRow}
      ></ModalForm>
      <section class='tableview_container'>
        <TableView fields={fields}
                   extraActions={extraActions}
                   bind:pager={pager}
                   rows={properties}
                   widths={{
                               mainUse: '320px',
                               address: '240px',
                               street: '240px',
                               city: '120px',
                               countyName: '120px',
                           }}
                   onRefreshTableView={refreshTableView}
                   onEditRow={editRow}>
          <button on:click={addRow} class='add_button'>
            <Icon size="17" color='#FFF' src={FaSolidPlusCircle} />
            <span>Add</span>
          </button>
          <button on:click={filterPendingApproval} class='filter_pending_button' class:not_filtered={!filterdByPending}>
            <Icon size="17" color='{!filterdByPending ? "#FFF" : "#000"}' src={FaSolidCheckDouble} />
            <span>Filter Pending</span>
          </button>
          <button on:click={filterRejectedApproval} class='filter_pending_button' class:not_filtered={!filterdByRejected}>
            <Icon size="17" color='{!filterdByRejected ? "#FFF" : "#000"}' src={FaSolidRecycle} />
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
              <PillBox label={key} content={new Date(val*1000).toString()} />
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
  .filter_pending_button {
    padding         : 8px 20px;
    font-size       : 14pt;
    font-weight     : bold;
    border-radius   : 8px;
    filter          : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    margin-left     : 4px;
    cursor          : pointer;
    align-items     : center;
    align-content   : center;
    justify-content : center;
    border          : 1px solid grey;
    gap             : 8px;
    display         : inline-flex;
  }

  .filter_pending_button.not_filtered {
    background-color : #6366F1;
    color            : white;
    flex-direction   : row;
  }

</style>
