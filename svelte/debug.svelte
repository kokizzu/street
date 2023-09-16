<script>

    import {T} from './_components/uiState.js';

</script>

{$T.currentLang}
{$T.forRent}
{$T.forSale}
<button on:click={() => T.changeLanguage('EN')}>EN</button>
<button on:click={() => T.changeLanguage('TW')}>TW</button>
