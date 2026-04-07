<script>
  import {notifier} from './notifier.js';
</script>
<hr/>
checking whether component conflict or not
<button onclick={() => notifier.showInfo('test')}>Notify</button>
<button onclick={() => notifier.showSuccess('test')}>Success</button>
<button onclick={() => notifier.showError('test')}>Error</button>
<button onclick={() => notifier.showWarning('test')}>Warn</button>
