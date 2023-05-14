<script>
  import {GuestRegister} from "./jsApi.GEN.js"
  import {onMount} from "svelte";
  
  // TODO: codegen axios part
  
  // server state
  const title = '#{title}' // /*! title */ {/* title */} [/* title */]
  // TODO: print session or fetch from cookie
  
  // local state
  let email = ''
  let password = ''
  let confirmPass = ''
  
  const LOGIN = 'LOGIN'
  const REGISTER = 'REGISTER'
  let mode = LOGIN
  
  function onHashChange() {
    const hash = location.hash
    if( hash===`#${LOGIN}` ) mode = LOGIN
    else if( hash===`#${REGISTER}` ) mode = REGISTER
    else location.hash = LOGIN
  }
  
  onMount( onHashChange )
  
  async function guestRegister() {
    // TODO: replace all alert with growl
    if( !email ) return alert( 'email is required' )
    if( password.length<12 ) return alert( 'password must be at least 12 characters' )
    if( password!==confirmPass ) return alert( 'passwords do not match' )
    // TODO: send to backend
    const i = { email, password }
    await GuestRegister( i, function( o ) {
      // TODO: codegen commonResponse
      // TODO: generate jsdoc proper callback type is something like this:
      /*
         if succeed (same as raw.response)
	      config:{transitional: {…}, adapter: Array(2), transformRequest: Array(1), transformResponse: Array(1), timeout: 0, …}
			data:{sessionToken: '', error: 'email already used', status: 400, user: {…}}
			headers: AxiosHeaders {content-length: '277', content-type: 'application/json', date: 'Sat, 13 May 2023 22:53:11 GMT'}
			request:XMLHttpRequest {onreadystatechange: null, readyState: 4, timeout: 0, withCredentials: false, upload: XMLHttpRequestUpload, …}
			status:200
			statusText:"Status OK"
			
			if error
			code:"ERR_BAD_REQUEST"
			config:{transitional: {…}, adapter: Array(2), transformRequest: Array(1), transformResponse: Array(1), timeout: 0, …}
			message:"Request failed with status code 400"
			name:"AxiosError"
			request:XMLHttpRequest {onreadystatechange: null, readyState: 4, timeout: 0, withCredentials: false, upload: XMLHttpRequestUpload, …}
			data: same as response.data
			response: same as above
       */
      console.log( o )
      if(o.error) alert(o.error)
    } )
  }
  
  function guestLogin() {
  
  }

</script>

<svelte:window on:hashchange={onHashChange}/>
<h1>{title} - {mode}</h1>
<div class="mainContainer">
	<label for="email">Email</label>
	<input type="text" id="email" bind:value={email}/><br/>
	<label for="password">Password</label>
	<input type="password" id="password" bind:value={password}><br/>
	{#if mode===REGISTER}
		<label for="confirmPass">Confirm Password</label>
		<input type="password" id="confirmPass" bind:value={confirmPass}><br/>
		<button on:click={guestRegister}>Register</button>
		<br/>
		Already have account?
		<a href="#LOGIN" on:click={()=> mode=LOGIN}>Login</a>
	{:else}
		<button on:click={guestLogin}>Login</button>
		<br/>
		Have no account?
		<a href="#REGISTER" on:click={()=> mode=REGISTER}>Register</a>
	{/if}
</div>

<style>
    /* pad label and input so they are equal in size */
    label {
        display : inline-block;
        padding : 0.2em;
        width   : 5em;
    }

    input {
        width   : 10em;
        padding : 0.1em;
        margin  : 0.2em
    }

    button {
        padding : 0.3em;
    }

    * {
        font-family : sans-serif;
        font-size   : 16pt;
    }

    .mainContainer {
        margin : 0 auto;
    }
</style>