<script>
// @ts-nocheck

	/** @typedef {import('../_types/user').User} User */
  /** @typedef {import('../_types/master').Access} Access */
  /** @typedef {import('../_types/master').PagerIn} PagerIn */
  /** @typedef {import('../_types/master').PagerOut} PagerOut */
  /** @typedef {import('../_types/master').Field} Field */

	import Main from '../_layouts/Main.svelte';
	import AdminSubMenu from '../_components/AdminSubMenu.svelte';
	import TableView from '../_components/TableView.svelte';
	import { AdminFiles } from '../jsApi.GEN';
	import Dropzone from '../_components/UploadDropzone.svelte';
	import {notifier} from '../_components/notifier.js';

	let user 		= /** @type {User} */ ({/* user */});
  let access	= /** @type {Access} */ ({/* segments */});
  let fields	= /** @type {Field[]} */ ([/* fields */]);
	let files   = /** @type {any}*/ ([/* files */]);
	let pager   = /** @type {PagerOut} */ ({/* pager */});

	/**
	 * @description Handle AJAX response
	 * @param res {any}
	 * @returns {boolean}
	 */
	function handleResponse(res) {
		console.log(res);
		if (res.error) {
			notifier.showError(res.error);
			return true;
		}
		if (res.files && res.files.length) files = res.files;
		if (res.pager && res.pager.page) pager = res.pager;
	}

	async function refreshTableView(/** @type {PagerIn} */ pagerIn ) {
		await AdminFiles({ // @ts-ignore
			pager: pagerIn,
			cmd: 'list',
		}, async function(res) {
			handleResponse(res);
		});
	}


	let uploads = {
		accepted: [],
		rejected: [],
	};

	// TODO: create component out of this

	function handleFilesSelect(/**  @type {DragEvent | any} */e) {
		const {acceptedFiles, fileRejections} = e.detail;
		files.accepted = [...uploads.accepted, ...acceptedFiles];
		files.rejected = [...uploads.rejected, ...fileRejections];
		if (acceptedFiles.length < 1) return;

		let file = acceptedFiles[0];
		//console.log(file.name,file.size,file.type);
		var formData = new FormData();
		formData.append('rawFile', file);
		formData.append('purpose', 'property'); // property or floorPlan
		var ajax = new XMLHttpRequest();
		ajax.upload.addEventListener('progress', progressHandler, false);
		ajax.addEventListener('load', completeHandler, false);
		ajax.addEventListener('error', errorHandler, false);
		ajax.addEventListener('abort', abortHandler, false);
		ajax.open('POST', '/user/uploadFile');
		ajax.send(formData);
	}

	let uploadStatus = '';
	let uploadPercent = 0;
	let lastImage = '';

	function progressHandler(event) {
			var percent = (event.loaded / event.total) * 100;
			uploadPercent = Math.round(percent);
			uploadStatus = uploadPercent + '% uploaded... please wait';
	}

	function completeHandler(event) {
			uploadStatus = event.target.responseText;
			uploadPercent = 0; //wil clear progress bar after successful upload
			const out = JSON.parse(event.target.responseText);
			if (!out.error) {
					lastImage = out.resizedUrl; // .originalUrl also available
			}
	}

	function errorHandler(event) {
			uploadStatus = 'Upload Failed';
	}

	function abortHandler(event) {
			uploadStatus = 'Upload Aborted';
	}
</script>

<Main {user} {access}>
  <div class="admin-container">
    <AdminSubMenu />
    <div class="admin-content">
			<section class="tableview-container">
				<TableView {fields}
					bind:pager={pager}
					rows={files}
					onRefreshTableView={refreshTableView}
				/>
			</section>
			<div class="dropzone-container">
				<Dropzone
					accept="image/*"
					on:drop={handleFilesSelect}
				/>
				<ol>
					{#each uploads.accepted as item}
						<li>{item.name}</li>
					{/each}
				</ol>
				<h3>{uploadStatus}</h3>
				<progress value={uploadPercent} max="100" style="width:300px;"></progress>
				{#if uploads.rejected && uploads.rejected.length}
					<h3>Errors</h3>
					<ol>
						{#each uploads.rejected as item}
							<li>{item.name}</li>
						{/each}
					</ol>
				{/if}
				{#if lastImage}
					<hr />
					<img src={lastImage} alt="last uploaded" />
				{/if}
			</div>
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