<script>
  // @ts-nocheck
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidInfoCircle from 'svelte-icons-pack/fa/FaSolidInfoCircle'; // info
  import FaSolidCheckCircle from 'svelte-icons-pack/fa/FaSolidCheckCircle'; // success
  import FaSolidExclamationTriangle from 'svelte-icons-pack/fa/FaSolidExclamationTriangle'; // warning
  import FaSolidTimesCircle from 'svelte-icons-pack/fa/FaSolidTimesCircle'; // error
  import FaSolidQuestionCircle from 'svelte-icons-pack/fa/FaSolidQuestionCircle'; // default
  
  // BUG: get error when calling method
  
  let message = 'This is growl, XD';
  let glowlType = 'default'; // default, info, warning, error, success
  let duration = 3000;
  let isVisible = true;
  setTimeout( () => {
    isVisible = false;
  }, duration );
  // export function showGlowl(msg, type, duration) {
  //   isVisible = true;
  //   message = msg;
  //   glowlType = type;
  //   setTimeout( () => {
  //     isVisible = false;
  //   }, duration );
  // }
</script>

{#if isVisible}
  <div class={`growl ${glowlType}`}>
	 {#if glowlType === 'default'}
		<Icon className='icon_growl' size={20} color='#475569' src={FaSolidQuestionCircle} />
	 {/if}
	 {#if glowlType === 'info'}
		<Icon className='icon_growl' size={20} color='#ffff' src={FaSolidInfoCircle} />
	 {/if}
	 {#if glowlType === 'success'}
		<Icon className='icon_growl' size={20} color='#ffff' src={FaSolidCheckCircle} />
	 {/if}
	 {#if glowlType === 'error'}
		<Icon className='icon_growl' size={20} color='#ffff' src={FaSolidTimesCircle} />
	 {/if}
	 {#if glowlType === 'warning'}
		<Icon className='icon_growl' size={19} color='#ffff' src={FaSolidExclamationTriangle} />
	 {/if}
	 <span>{message}</span>
  </div>
{/if}

<style>
  .growl {
	 position: fixed;
	 display: flex;
	 flex-direction: row;
	 align-items: center;
	 gap: 8px;
	 font-size: 14px;
	 bottom: 20px;
	 right: 20px;
	 padding: 10px 20px;
	 border-radius: 3px;
	 box-shadow: 0px 4px 24px 0px rgba(0, 0, 0, 0.25);
	 z-index: 9999;
	 min-width: 120px;
	 max-width: 350px;
	 height: fit-content;
  }
  :global(.icon_growl) {
    flex-shrink: 0;
  }
  .growl span {
	 flex-shrink: 1;
  }
  .growl.default {
	 background-color: #FFFFFF;
	 color: #475569;
  }
  .growl.info {
	 background-color: #1080e8;
	 color: #FFFFFF;
  }
  .growl.success {
	 background-color: #059669;
    color: #FFFFFF;
  }
  .growl.error {
	 background-color: #ef4444;
	 color: #FFFFFF;
  }
  .growl.warning {
	 background-color: #d97706;
	 color: #FFFFFF;
  }
</style>
