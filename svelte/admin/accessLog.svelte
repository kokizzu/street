<script>
    // @ts-nocheck
    import Menu from '../_components/Menu.svelte';
    import AdminSubMenu from '../_components/AdminSubMenu.svelte';
    import ProfileHeader from '../_components/ProfileHeader.svelte';
    import Footer from '../_components/Footer.svelte';
    import TableView from '../_components/TableView.svelte';
    import {AdminAccessLogs} from '../jsApi.GEN';
    import Growl from '../_components/Growl.svelte';

    let segments = {/* segments */};
    let fields = [/* fields */];
    let logs = [/* logs */];
    let pager = {/* pager */};

    $: console.log(logs, fields, pager);
    let growl4 = Growl;

    // return true if got error
    function handleResponse(res) {
        if (res.error) {
            alert(res.error);
            return true;
        }
        if (res.logs && res.logs.length) logs = res.logs;
        if (res.pager && res.pager.page) pager = res.pager;
    }

    async function refreshTableView(pagerIn) {
        // console.log( 'pagerIn=',pagerIn );
        await AdminAccessLogs({
            pager: pagerIn,
        }, function(res) {
            handleResponse(res);
        });
    }
</script>


<section class='dashboard'>
    <Menu access={segments}/>
    <div class='dashboard_main_content'>
        <ProfileHeader access={segments}/>
        <AdminSubMenu></AdminSubMenu>
        <div class='content'>
            <section class='tableview_container'>
                <TableView arrayOfArray={false}
                           bind:pager={pager}
                           {fields}
                           onRefreshTableView={refreshTableView}
                           rows={logs}
                >
                </TableView>
            </section>
        </div>
        <Footer></Footer>
    </div>
</section>

<style>
</style>