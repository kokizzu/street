<script>
  /** @typedef {import('../_types/user').User} User */
  /** @typedef {import('../_types/master').Access} Access */

  import Main from '../_layouts/Main.svelte';
  
  let user   = /** @type {User} */ ({/* user */});
  let access = /** @type {Access} */ ({/* segments */});
</script>

<Main {user} {access}>
  <div class="revenue-container">
    <div class="table-root">
      <div class="table-container">
        <table>
          <thead>
            <tr>
              <th>Invoice #</th>
              <th>Revenue</th>
              <th>Realtor</th>
              <th>No. of listings</th>
              <th>Register date</th>
              <th>Purchase date</th>
            </tr>
          </thead>
          <tbody>
            {#each Array(10) as _}
              <tr>
                <td>US9340-A343</td>
                <td>$50,454,899</td>
                <td>Dennis Lowe</td>
                <td>45</td>
                <td>March 34, 2024</td>
                <td>March 34, 2024</td>
              </tr>
            {/each}
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