<script>
  import Menu from './_components/menu.svelte'
  
  let user = {/* user */}
  let segments = {/* segments */}
</script>

<Menu access={segments}/>
<ul>
  <li><a href="/admin/users">Users</a></li>
  <li><a href="/admin/sessions">Sessions</a></li>
  <li><a href="/admin/actionLogs">Sessions</a></li>
</ul>