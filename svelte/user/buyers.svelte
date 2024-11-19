<script>
  /** @typedef {import('../_types/user').User} User */
  /** @typedef {import('../_types/master').Access} Access */
  /** @typedef {import('../_types/business').Buyer} Buyer */

  import Main from '../_layouts/Main.svelte';
  import { formatPrice, getYearMonth } from '../_components/formatter';
  import InputBox from '../_components/InputBox.svelte';
  import { UserBuyers } from '../jsApi.GEN';
  import { notifier } from '../_components/notifier';
  
  const user      = /** @type {User} */ ({/* user */});
  const access    = /** @type {Access} */ ({/* segments */});

  let buyers = /** @type {Buyer[]} */ ([/* buyers */]);

  let monthFilter = /** @type {Date|string|any} */ (getYearMonth());
  async function applyFilterYearMonth() {
    const inUserBuyers = /** @type {import('../jsApi.GEN').UserBuyersIn} */ ({
      cmd: 'list',
      yearMonth: monthFilter
    });
    await UserBuyers(inUserBuyers, async function(/** @type {any} */ res) {
      if (res.error) {
        console.log('error =', res.error);
        notifier.showError(res.error || 'failed to get buyers data');
        return;
      }

      buyers = res.buyers;
    })
  }
</script>

<Main {user} {access}>
  <div class="buyers-container">
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
    </div>
    <div class="table-root">
      <div class="table-container">
        <table>
          <thead>
            <tr>
              <th class="no">No</th>
              <th>Property ID</th>
              <th>Buyer Email</th>
							<th>Property Price</th>
              <th>Sales Date</th>
            </tr>
          </thead>
          <tbody>
						{#if buyers && buyers.length > 0}
              {#each (buyers || []) as by, idx}
                <tr>
                  <td>{idx + 1}</td>
                  <td>{by.propertyId}</td>
                  <td>{by.buyerEmail}</td>
                  <td>{formatPrice(Number(by.propertyPrice), 'USD')}</td>
                  <td>{by.salesDate}</td>
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
  .buyers-container {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }

  .buyers-container .header {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
  }

  .buyers-container .header .filter {
    display: flex;
    flex-direction: row;
    align-items: stretch;
    gap: 10px;
  }

  .buyers-container .header .filter .btn-filter {
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

  .buyers-container .header .filter .btn-filter:hover {
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