<script>
  
  
  import {GuestResetPassword} from "../jsApi.GEN";
  
  let password = ''
  let pass2 = ''
  
  async function resetPassword() {
    if( password.length<12 ) return alert( 'password must be at least 12 characters')
    if( password!==pass2 ) return alert( 'password confirmation does not match' )
    const queryParam = window.location.href.split( '?' )[ 1 ]
    const qps = queryParam.split( '&' )
    let secretCode = ''
    let hash = ''
    for( let i = 0; i<qps.length; i++ ) {
      const [key, value] = qps[ i ].split( '=' )
      if( key==='secretCode' ) secretCode = value
      if( key==='hash' ) hash = value
    }
    let i = {secretCode, hash, password}
    await GuestResetPassword( i, function( o ) {
      console.log( o )
      if( o.error ) return alert( o.error )
      alert( 'password reset successful' )
      window.location.href = '/'
    } )
  }


</script>

New Password<br/>
<input type="password" bind:value={password}/><br/>
Confirm New Password<br/>
<input type="password" bind:value={pass2}/><br/>

<button on:click={resetPassword}>Reset Password</button>