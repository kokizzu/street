<script>
  /** @typedef {import('../_types/user').User} User */
  /** @typedef {import('../_types/master').Access} Access */
  /** @typedef {import('../_types/business').AdminRevenue} Revenue */

  import Main from '../_layouts/Main.svelte';
  import { Icon } from '../node_modules/svelte-icons-pack/dist';
  import { RiSystemAddLargeFill } from '../node_modules/svelte-icons-pack/dist/ri';
  import { formatPrice, datetime } from '../_components/formatter';
  
  let user      = /** @type {User} */ ({/* user */});
  let access    = /** @type {Access} */ ({/* segments */});
  let revenues  = /** @type {Revenue[]} */ ([/* revenues */]);

  console.log('revenues =', revenues);
</script>

<Main {user} {access}>
  <haha class="revenue-container">
    <hehe class="header">
      <miaw></miaw>
      <button class="add-btn">
        <Icon
          size="20"
          color="#FFF"
          src={RiSystemAddLargeFill}
        />
        <span>Add</span>
      </button>
    </hehe>
    <huhu class="table-root">
      <div class="table-container">
        <table>
          <thead>
            <tr>
              <th>Property ID</th>
              <th>Realtor ID</th>
              <th>Property Country</th>
              <th>Revenue</th>
              <th>Property Bought</th>
              <th>Sales Date</th>
              <th>Created At</th>
              <th>Updated At</th>
            </tr>
          </thead>
          <tbody>
            {#if revenues && revenues.length > 0}
              {#each (revenues || []) as rv}
                <tr>
                  <td>{rv.propertyId}</td>
                  <td>{rv.realtorId}</td>
                  <td>{rv.propertyCountry || 'Default'}</td>
                  <td>{formatPrice(Number(rv.revenue), 'USD')}</td>
                  <td>{rv.propertyBought}</td>
                  <td>{rv.salesDate}</td>
                  <td>{datetime(rv.createdAt)}</td>
                  <td>{datetime(rv.updatedAt)}</td>
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
    </huhu>
  </haha>
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

  .revenue-container .header .add-btn:hover {
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