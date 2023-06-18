<script>
  import Menu from '../_components/menu.svelte';
  import TableView from '../_components/tableView.svelte';
  
  let segments = {/* segments */};
  let fields = [/* fields */ ];
  let users = [/* users */];

  $: console.log(users, fields)
  
</script>
<Menu access={segments} />
<TableView {fields} rows={users}></TableView>