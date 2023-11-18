<script>
  import {notifier} from './notifier.js';
</script>
<hr/>
checking whether component conflict or not
<button on:click={() => notifier.showInfo('test')}>Notify</button>
<button on:click={() => notifier.showSuccess('test')}>Success</button>
<button on:click={() => notifier.showError('test')}>Error</button>
<button on:click={() => notifier.showWarning('test')}>Warn</button>
