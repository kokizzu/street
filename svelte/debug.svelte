<script>
    import {T} from './_components/uiState.js';
    import { UserSendFeedback } from './jsApi.GEN';
    import {notifier} from './_components/notifier.js';
    import DebugComponent from './_components/DebugComponent.svelte';

    let userMessage = $state('');
    function sendFeedback() {
        UserSendFeedback({userMessage}, function( res ){
            if(res.error)  return alert(res.error);
            console.log(res);
            alert('success')
        })
    }
</script>

{$T.currentLang}
{$T.forRent}
{$T.forSale}
<button onclick={() => T.changeLanguage('EN')}>EN</button>
<button onclick={() => T.changeLanguage('TW')}>TW</button>

<hr/>
<label for='userMessage'>User Mesage</label> <br/>
<textarea id='userMessage' cols='40' rows='3' placeholder='your feedback here' bind:value={userMessage}></textarea><br/>


<button onclick={sendFeedback}>Send Feedback</button>
<button onclick={() => notifier.showInfo('test')}>Notify</button>
<button onclick={() => notifier.showSuccess('test')}>Success</button>
<button onclick={() => notifier.showError('test')}>Error</button>
<button onclick={() => notifier.showWarning('test')}>Warn</button>

<DebugComponent/>