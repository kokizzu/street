<script>
	/** @typedef {import('../_types/user').User} User */
	/** @typedef {import('../_types/master').Access} Access */
	/** @typedef {import('../_types/master').PagerOut} PagerOut */
	/** @typedef {import('../_types/master').PagerIn} PagerIn */
	/** @typedef {import('../_types/master').Field} Field */

	import Main from '../_layouts/Main.svelte';
	import AdminSubMenu from '../_components/AdminSubMenu.svelte';
	import TableView from '../_components/TableView.svelte';
	import { AdminAccessLogs } from '../jsApi.GEN';
	import { notifier } from '../_components/notifier.js';

	let user 		= /** @type {User} */ ({/* user */});
	let fields	= /** @type {Field[]}	*/ ([/* fields */]);
	let logs 		= /** @type {any[]} */ ([/* logs */]);
	let pager		= /** @type {PagerOut} */ ({/* pager */});

	function handleResponse(res) {
		if (res.error) {
			notifier.showError(res.error);
			return true;
		}
		if (res.logs && res.logs.length) logs = res.logs;
		if (res.pager && res.pager.page) pager = res.pager;
	}

	async function refreshTableView(/** @type {PagerIn}*/ pagerIn) {
		await AdminAccessLogs({ // @ts-ignore
			pager: pagerIn,
		}, function(/** @type {any}*/ res) {
			handleResponse(res);
		});
	}
</script>


<Main {user}>
  <div class="admin-container">
    <AdminSubMenu />
    <div class="admin-content">
			<TableView arrayOfArray={false}
				bind:pager={pager}
				{fields}
				onRefreshTableView={refreshTableView}
				rows={logs}
			/>
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