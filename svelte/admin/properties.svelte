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
            onClick: function(item) {
                let propertyKey = item[1]; // property key for taiwan data
                UserPropHistory({
                    propertyKey: propertyKey,
                }, function(res) {
                    if (res.error) alert(res.error);
                    propHistoryModal.showModal();
                    currentPropHistory = res.History || [];
                });
            },
        },
    ];

    $: console.log('properties=', properties);

    let propHistoryModal = ModalDialog;

    // return true if got error
    function handleResponse(res) {
        console.log(res);
        if (res.error) {
            // TODO: growl error
            alert(res.error);
            return true;
        }
        if (res.properties && res.properties.length) properties = res.properties;
        if (res.pager && res.pager.page) pager = res.pager;
    }

    async function refreshTableView(e) {
        const pagerIn = e.detail;
        // console.log( pager );
        await AdminProperties({
            pager: pagerIn,
            action: 'list',
        }, function(res) {
            handleResponse(res);
        });
    }

    let form = ModalForm; // for lookup

    async function editRow(e) {
        const id = e.detail;
        await AdminProperties({
            property: {id},
            action: 'form',
        }, function(res) {
            if (!handleResponse(res))
                form.showModal(res.property);
        });
    }

    function addRow() {
        form.showModal({id: ''});
    }

    async function saveRow(action, row) {
        let property = {...row};
        if (!property.id) property.id = '0';
        console.log(property);
        try {
            property.coord = JSON.parse('[' + property.coord + ']');
        } catch (e) {
            property.coord = [0, 0];
        }
        await AdminProperties({
            property: property,
            action: action,
            pager: pager, // force refresh page, will be slow
        }, function(res) {
            if (handleResponse(res)) {
                return form.setLoading(false); // has error
            }
            form.hideModal(); // success
        });
    }
</script>

<section class='dashboard'>
    <Menu access={segments} />
    <div class='dashboard_main_content'>
        <ProfileHeader />
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
                           on:refreshTableView={refreshTableView}
                           on:editRow={editRow}
                >
                    <button on:click={addRow} class='add_button'>
                        <Icon size={18} color='#FFFF' src={FaSolidPlusCircle} />
                        <span>Add</span>
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
                    {JSON.stringify(row)}<br />
                {/each}
            {:else}
                no history for this property
            {/if}
        </div>
    </ModalDialog>
</section>

<style>

</style>
