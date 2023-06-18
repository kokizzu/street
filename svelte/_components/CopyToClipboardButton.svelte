<script>
  import FloatingNotification from './FloatingNotification.svelte'
  import Icon from './Icon.svelte'
  
  export let value = '';
  export let title = 'copy to clipboard';
  
  // local state
  let copied = 0;
  
  function showNotification() {
    copied = setTimeout( function() {
      copied = 0;
    }, 1500 )
  }
  
  function copyToClipboard() {
    showNotification()
    if( navigator.clipboard ) {
      return navigator.clipboard.writeText( value )
    }
    // Fall back to document.execCommand (deprecated)
    const el = document.createElement( 'textarea' )
    el.value = value
    el.style.opacity = 0
    document.body.append( el )
    el.select()
    document.execCommand( 'copy' )
    el.remove()
    return Promise.resolve()
  }

</script>
<button title="{title}" type="button" class="iconButton"
        on:click={copyToClipboard}>
	<Icon name="copy"/>
</button>
{#if copied}
	<FloatingNotification icon='copy' text={value} subtext='copied to clipboard'/>
{/if}
