<script>
  export let fields = []; // list of editable properties
  let visible = false;
  
  function showDialog( row ) {
  
  }
  
  // TODO: continue this

</script>
<div class:hidden={!visible}>
  {#each fields as field}
    <div>{JSON.stringify( field )}</div>
  {/each}
</div>
<style>
  div.hidden {
      display: none;
  }
</style>