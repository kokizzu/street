<script>
    import {T} from './_components/uiState.js';
    import Growl from './_components/Growl.svelte';

    let myGrowl// = Growl;
</script>

{$T.currentLang}
{$T.forRent}
{$T.forSale}
<button on:click={() => T.changeLanguage('EN')}>EN</button>
<button on:click={() => T.changeLanguage('TW')}>TW</button>
<button on:click={() => alert('This is Growl')}>
    Show Growl
</button>

