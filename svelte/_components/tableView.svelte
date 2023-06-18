<script>
  export let fields = []; // array of field object
  export let rows = []; // 2 dimension array
  
  function date( unixSec ) {
    if( !unixSec ) return '';
    const date = new Date( unixSec * 1000 );
    return date.toISOString().substring( 0, 10 );
  }
</script>
<table>
  <thead>
  <tr>
    {#each fields as field}
      <th>{field.label}</th>
    {/each}
  </tr>
  </thead>
  <tbody>
  {#each rows as row}
    <tr>
      {#each fields as field, i}
        {#if field.inputType==='checkbox'}
          <td>{!!row[ i ]}</td>
        {:else if field.inputType==='datetime'}
          <td>{date( row[ i ] )}</td>
        {:else}
          <td>{row[ i ]}</td>
        {/if}
      {/each}
    </tr>
  {/each}
  </tbody>
  <tfoot>
  </tfoot>
</table>

<style>
    table {
        border-collapse : collapse;
    }

    table, th, td {
        border : 1px solid black;
    }
</style>