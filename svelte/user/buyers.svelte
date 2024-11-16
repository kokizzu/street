<script>
  /** @typedef {import('../_types/user').User} User */
  /** @typedef {import('../_types/master').Access} Access */

  import Main from '../_layouts/Main.svelte';
  
  const user      = /** @type {User} */ ({/* user */});
  const access    = /** @type {Access} */ ({/* segments */});  
</script>

<Main {user} {access}>
  <div class="buyers-container">
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
						{#each Array(20) as _, idx}
							<tr>
								<td>{idx + 1}</td>
								<td>#2H8JQPXMV</td>
								<td>johndoe@example.com</td>
								<td>$ 1,000,000</td>
                <td>2025-01-01</td>
							</tr>
						{/each}
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