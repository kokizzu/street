<script>
  import Menu from './_components/menu.svelte'
  import AdminSubMenu from './admin/_adminSubMenu.svelte'
  
  let user = {/* user */}
  let segments = {/* segments */}
</script>

<Menu access={segments}/>
<AdminSubMenu></AdminSubMenu>
TODO: show summary statistics here