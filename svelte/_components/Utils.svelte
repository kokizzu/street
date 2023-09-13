<script>
  // @ts-nocheck
  import {currentLang} from "./uiState";
  import translation from "../translation.json";
  
  export function translate( key ) {
    if( $currentLang==='EN' ) {
      return translation[ key ]
    }
    const keyTranslate = key + $currentLang;
    console.log( keyTranslate )
    return translation[ keyTranslate ];
  }
</script>