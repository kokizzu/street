<script>
    /** @typedef {import('../_types/user').User} User */
    /** @typedef {import('../_types/master').Access} Access */
    /** @typedef {import('../_types/master').PagerIn} PagerIn */
    /** @typedef {import('../_types/master').PagerOut} PagerOut */
    /** @typedef {import('../_types/master').Field} Field */
  
    import Main from '../_layouts/Main.svelte';
    import AdminSubMenu from '../_components/AdminSubMenu.svelte';
    import TableView from '../_components/TableView.svelte';
    import { Admin3DFiles } from '../jsApi.GEN';
    import { notifier } from '../_components/notifier';
    import PopUpUpload3DFile from '../_components/PopUpUpload3DFile.svelte';
    import { Icon } from '../node_modules/svelte-icons-pack/dist';
    import { RiSystemAddBoxLine } from '../node_modules/svelte-icons-pack/dist/ri';
  
    let user 		= /** @type {User} */ ({/* user */});
    let access	= /** @type {Access} */ ({/* segments */});
    let fields	= /** @type {Field[]} */ ([/* fields */]);
    let files   = /** @type {any}*/ ([/* files */]);
    let pager   = /** @type {PagerOut} */ ({/* pager */});

    let popUpUpload3DFile = /** @type {import('svelte').SvelteComponent} */ (null);
  
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
      await Admin3DFiles({ // @ts-ignore
        pager: pagerIn,
        cmd: 'list',
      }, async function(res) {
        handleResponse(res);
      });
    }
  </script>

<PopUpUpload3DFile
  bind:this={popUpUpload3DFile}
/>
  
  <Main {user} {access}>
    <div class="admin-container">
      <AdminSubMenu />
      <div class="admin-content">
        <section class="tableview-container">
          <TableView {fields}
            bind:pager={pager}
            rows={files}
            onRefreshTableView={refreshTableView}
          >
          <button
            class="btn"
            on:click={() => popUpUpload3DFile.Show()}
            title="Upload 3D file"
          >
            <Icon
              size="17"
              src={RiSystemAddBoxLine}
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