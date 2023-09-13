<script>
  // TODO: Use this component for using Translator
  // @ts-nocheck
  import {currentLang} from "./uiState";
  import translation from "../translation.json";
  import {onMount} from "svelte";
  
  export function translate( key ) {
    if( $currentLang==='EN' ) {
      return translation[ key ]
    }
    const keyTranslate = key + $currentLang;
    console.log( keyTranslate )
    return translation[ keyTranslate ];
  }
  
  onMount(() => {
    console.log('loaded')
  })
</script>