<script>
  /** @typedef {import('../_types/user').User} User */
  /** @typedef {import('../_types/master').Access} Access */
  /** @typedef {import('../_types/business').Sales} Sales */
  /** @typedef {import('../_types/business').Revenue} Revenue */

  import Main from '../_layouts/Main.svelte';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { RiSystemAddLargeFill } from '../node_modules/svelte-icons-pack/dist/ri';
  import { formatPrice, getYearMonth } from '../_components/formatter';
  import PopUpAddSales from '../_components/PopUpAdminAddSales.svelte';
  import { AdminRevenue } from '../jsApi.GEN';
  import { notifier } from '../_components/notifier';
  import InputBox from '../_components/InputBox.svelte';
  
  let user      = /** @type {User} */ ({/* user */});
  let access    = /** @type {Access} */ ({/* segments */});
  let revenues  = /** @type {Revenue[]} */ ([/* revenues */]);

  console.log('revenues =', revenues);

  let monthFilter = /** @type {Date|string|any} */ (getYearMonth());

  async function applyFilterYearMonth() {
    const inAdminRevenue = /** @type {import('../jsApi.GEN').AdminRevenueIn} */ ({
      cmd: 'list',
      yearMonth: monthFilter
    });
    await AdminRevenue(inAdminRevenue, async function(/** @type {any} */ res) {
      if (res.error) {
        console.log('error =', res.error);
        notifier.showError(res.error || 'failed to get revenues');
        isSubmitAddSales = false;
        return;
      }

      revenues = res.revenues;
    })
  }

  let popUpAddSales = /** @type {import('svelte').SvelteComponent} */ (null);
  let isSubmitAddSales = /** @type {boolean} */ (false);

  /**
   * @description Submit add sales
   * @type {Function}
   * @param salesObj {Sales}
   * @param propKey {string}
   * @param realtorEmail {string}
   * @returns {Promise<void>}
   */
  const SubmitAddSales = async function(salesObj, propKey, realtorEmail) {
    console.log('Sales =', salesObj);
    console.log('Property Key =', propKey);
    console.log('Realtor Email =', realtorEmail);
    await AdminRevenue({
      cmd: 'upsert', // @ts-ignore
      sales: salesObj,
      propKey: propKey,
      realtorEmail: realtorEmail
    }, async function(/** @type {any} */ res) {
      
    })
  };
</script>

<PopUpAddSales
  bind:this={popUpAddSales}
  bind:isSubmitted={isSubmitAddSales}
  heading="Add Sales"
  OnSubmit={SubmitAddSales}
/>

<Main {user} {access}>
  <div class="revenue-container">
    <div class="header">
      <div class="filter">
        <InputBox
          type="month"
          id="month-filter"
          bind:value={monthFilter}
        />
        <button class="btn-filter" on:click={applyFilterYearMonth}>
          <span>Apply Filter</span>
        </button>
      </div>
      <button class="add-btn">
        <Icon
          size="20"
          color="#FFF"
          src={RiSystemAddLargeFill}
        />
        <span>Add</span>
      </button>
    </div>
    <div class="table-root">
      <div class="table-container">
        <table>
          <thead>
            <tr>
              <th class="no">No</th>
              <th>Revenue</th>
              <th>Property Bought</th>
              <th>Sales Date</th>
            </tr>
          </thead>
          <tbody>
            {#if revenues && revenues.length > 0}
              {#each (revenues || []) as rv, idx}
                <tr>
                  <td>{idx + 1}</td>
                  <td>{formatPrice(Number(rv.revenue), 'USD')}</td>
                  <td>{rv.propertyBought}</td>
                  <td>{rv.salesDate}</td>
                </tr>
              {/each}
            {:else}
              <tr>
                <td colspan={5}>No data</td>
              </tr>
            {/if}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</Main>

<style>
  .revenue-container {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }

  .revenue-container .header {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
  }

  .revenue-container .header .filter {
    display: flex;
    flex-direction: row;
    align-items: stretch;
    gap: 10px;
  }

  .revenue-container .header .filter .btn-filter {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
    background-color: var(--orange-006);
    border-radius: 999px;
    border: none;
    cursor: pointer;
    padding: 10px 30px;
    color: #FFF;
    font-size: 16px;
    font-weight: 500;
    text-decoration: none;
  }

  .revenue-container .header .add-btn {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
    background-color: var(--orange-006);
    border-radius: 9999px;
    border: none;
    cursor: pointer;
    padding: 15px 30px;
    color: #FFF;
    font-size: 16px;
    font-weight: 600;
    text-decoration: none;
  }

  .revenue-container .header .add-btn:hover,
  .revenue-container .header .filter .btn-filter:hover {
    background-color: var(--orange-005);
  }

  .table-root {
    display: flex;
    flex-direction: column;
    background-color: #fff;
    border-radius: 8px;
    border: 1px solid var(--gray-003);
    padding: 0;
    overflow: hidden;
		width: 100%;
  }

  .table-root .table-container {
    overflow-x: auto;
    scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: thin;
  }

  .table-root .table-container table {
    width: 100%;
    background: #fff;
    box-shadow: none;
    text-align: left;
    border-collapse: separate;
    border-spacing: 0;
    overflow: hidden;
		font-size: 13px;
  }

  .table-root .table-container table thead {
    box-shadow: none;
    border-bottom: 1px solid var(--gray-003);
  }

  .table-root .table-container table thead tr th {
    padding: 12px;
		background-color: var(--gray-001);
		text-transform: capitalize;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-003);
		min-width: fit-content;
		width: auto;
    text-wrap: nowrap;
  }

  .table_root .table_container table thead tr th.no {
    width: 30px;
  }

  .table-root .table-container table thead tr th:last-child {
    border-right: none;
  }

  .table-root .table-container table tbody tr td {
    padding: 8px 12px;
  }

	.table-root .table-container table tbody tr td {
    padding: 8px 12px;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-004);
  }

	.table-root .table-container table tbody tr:last-child {
		border-bottom: none !important;
  }

	.table-root .table-container table tbody tr:last-child td,
	.table-root .table-container table tbody tr:last-child th {
		border-bottom: none !important;
	}

  .table-root .table-container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .table-root .table-container table tbody tr:last-child td,
  .table-root .table-container table tbody tr:last-child th {
    border-bottom: none !important;
  }

  .table-root .table-container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .table-root .table-container table tbody tr td:last-child {
    border-right: none !important;
  }

  .table-root .table-container table tbody tr th {
    text-align: center;
    border-right: 1px solid var(--gray-004);
    border-bottom: 1px solid var(--gray-004);
  }
</style>