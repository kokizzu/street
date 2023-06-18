<script>
  import Menu from './_components/Menu.svelte'
  
  let user = {/* user */}
  let segments = {/* segments */}
</script>

<Menu access={segments}/>

logged user profile content goes here

<br/>
Debug:<br/>
<textarea cols="20" rows="10">{JSON.stringify( user, null, 2 )}</textarea>