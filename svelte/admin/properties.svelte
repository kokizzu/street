<script>
  /** @typedef {import('../_types/user').User} User */
  /** @typedef {import('../_types/master').Access} Access */
  /** @typedef {import('../_types/master').PagerIn} PagerIn */
  /** @typedef {import('../_types/master').PagerOut} PagerOut */
  /** @typedef {import('../_types/master').Field} Field */
  /** @typedef {import('../_types/property').Property} Property*/
  /** @typedef {import('../_types/master').ExtendedAction} ExtendedAction*/

  import Main from '../_layouts/Main.svelte';
  import AdminSubMenu from '../_components/AdminSubMenu.svelte';
  import TableView from '../_components/TableView.svelte';
  import ModalForm from '../_components/ModalForm.svelte';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import {
    RiSystemEyeLine, RiSystemCheckboxCircleLine, RiSystemAddBoxLine,
    RiSystemCloseCircleLine, RiBusinessLinksLine, RiSystemFilter3Fill,
    RiSystemProhibitedLine
  } from '../node_modules/svelte-icons-pack/dist/ri';
  import { AdminProperties, UserPropHistory } from '../jsApi.GEN';
  import { notifier } from '../_components/notifier.js';
  import ModalDialog from '../_components/ModalDialog.svelte';
  import PillBox from '../_components/PillBox.svelte';
  import { priceNtd } from '../_components/formatter.js';
  import { fieldsArrToMap } from '../_components/mapper.js';
  
  let user        = /** @type {User} */ ({/* user */});
  let access      = /** @type {Access} */ ({/* segments */});
  let fields      = /** @type {Field[]} */ ([/* fields */]);
  let properties  = /** @type {Property[]} */ ([/* properties */]);
  let pager       = /** @type {PagerOut} */ ({/* pager */});

  let propHistoryModal  = /** @type {import('svelte').SvelteComponent}*/ (null);
  let modalForm         = /** @type {import('svelte').SvelteComponent}*/ (null);

  let fieldByKey = fieldsArrToMap( fields );

  let currentPropHistory = [];

  /** @type {ExtendedAction[]} */
  let extraActions = [
    {
      icon: RiSystemEyeLine,
      label: 'Show Property History',
      onClick: function (/** @type {any} */ item) {
        let propertyKey = item[fieldByKey['uniqPropKey'].idx]; // property key for taiwan data
        UserPropHistory({
          propertyKey: propertyKey,
        }, async function (/** @type {any} */ res) {
          if (res.error) alert(res.error);
          propHistoryModal.showModal();
          currentPropHistory = res.history || [];
        });
      },
      link: null,
      showIf: null
    },
    {
      icon: RiSystemCheckboxCircleLine,
      showIf: function (/** @type {any} */ row) {
        return !!row[fieldByKey['approvalState'].idx];
      },
      label: 'Approve property',
      onClick: function (item) {
        const id = item[0];
        AdminProperties({
          cmd: 'upsert', // @ts-ignore
          property: {
            id: id,
            approvalState: ' ' // empty is approved, will be trimmed on server side
          },
          pager: pager,
        },
          function (/** @type {any} */ res) {
            if (res.error) notifier.showError(res.error || 'Failed to update property');
          });
      },
      link: null
    },
    {
      icon: RiSystemCloseCircleLine,
      showIf: function (/** @type {any} */ row) {
        return !row[fieldByKey['approvalState'].idx];
      },
      label: 'Reject property',
      onClick: function (/** @type {any} */ item) {
        const id = item[0];
        const reason = prompt('input refusal reason for #' + id + ', use "pending" to set state to pending');
        if (!reason) return;
        AdminProperties({
          cmd: 'upsert', // @ts-ignore
          property: {
            id: id,
            approvalState: reason
          },
        }, async function (/** @type {any} */ res) {
          if (res.error) {
            notifier.showError(res.error || 'Failed to reject property');
          }
        });
      },
      link: null
    },
    {
      icon: RiBusinessLinksLine,
      label: 'Go to property page',
      link: function (/** @type {any} */ item) {
        return '/guest/property/' + item[fieldByKey['id'].idx];
      },
      onClick: null,
      showIf: null
    },
  ];
  
  /**
	 * @description Handle AJAX response
	 * @param res {any}
	 * @returns {boolean}
	 */
	function handleResponse(res) {
    if( res.error ) {
      notifier.showError( res.error );
      return true;
    }
    if( res.properties && res.properties.length ) properties = res.properties;
    if( res.pager && res.pager.page ) pager = res.pager;
    if( res.pager && res.pager.filters && !res.properties ) properties = []; // if nothing found but filter exists, clear table
  }
  
  async function refreshTableView(/** @type {PagerIn} */ pagerIn ) {
    await AdminProperties( { // @ts-ignore
      pager: pagerIn,
      cmd: 'list',
    }, function( res ) {
      handleResponse( res );
    } );
  }
  
  async function editRow(/** @type {number|string} */ id, /** @type {any[]} */ row ) {
    await AdminProperties( { // @ts-ignore
      property: {id},
      cmd: 'form',
    }, function(/** @type {any} */ res ) {
      if( !handleResponse( res ) )
        modalForm.showModal( res.property );
    });
  }
  
  let filterdByPending = false;
  let filterdByRejected = false;

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
      cmd: 'list', // @ts-ignore
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
      cmd: 'list', // @ts-ignore
      pager: pager,
    }, handleResponse );
  }
  
  function addRow() {
    modalForm.showModal( {id: ''} );
  }
  
  async function saveRow(/** @type {string}*/ action, /** @type {any} */ row ) {
    let property = {...row};
    if( !property.id ) property.id = '0';
    try {
      property.coord = JSON.parse( '[' + property.coord + ']' );
    } catch( e ) {
      property.coord = [0, 0];
    }
    await AdminProperties( { // @ts-ignore
      property: property,
      cmd: action,
      pager: pager, // force refresh page, will be slow
    }, function( res ) {
      if( handleResponse( res ) ) {
        return modalForm.setLoading( false ); // has error
      }
      modalForm.hideModal(); // success
    } );
  }
  
  let renderMap = {
    note: function(/** @type {any} */ note ) {
      return JSON.stringify( note );
    },
  };
</script>


<Main {user} {access}>
  <div class="admin-container">
    <AdminSubMenu />
    <div class="admin-content">
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
      <ModalForm
        bind:this={modalForm}
        fields={fields}
        onConfirm={saveRow}
        rowType="Property"
      />
      <section class="tableview-container">
        <TableView
          bind:pager={pager}
          extraActions={extraActions}
          fields={fields}
          onEditRow={editRow}
          onRefreshTableView={refreshTableView}
          rows={properties||[]}
          renderFuncs={renderMap}
          widths={{
            mainUse: '320px',
            address: '240px'
          }}
        >
          <button
            class="btn"
            on:click={addRow}
            title="Add property"
          >
            <Icon
              size="17"
              src={RiSystemAddBoxLine}
              color="var(--gray-008)"
            />
          </button>
          <button
            class="btn"
            disabled={!filterdByPending}
            on:click={filterPendingApproval}
            title="Filter by pending approval"
          >
            <Icon
              size="17"
              src={RiSystemFilter3Fill}
              color="var(--gray-008)"
            />
          </button>
          <button
            class="btn"
            disabled={!filterdByRejected}
            on:click={filterRejectedApproval}
            title="Filter by rejected approval"
          >
            <Icon
              size="17"
              src={RiSystemProhibitedLine}
              color="var(--gray-008)"
            />
          </button>
        </TableView>
      </section>
    </div>
  </div>
</Main>

<style>
  .admin-container {
    display: flex;
    flex-direction: column;
    width: 100%;
    gap: 20px;
    padding: 10px 20px 20px;
  }

  .admin-container .admin-content {
    display: flex;
    flex-direction: column;
    width: 100%;
    gap: 20px;
  }
</style>