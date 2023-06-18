<script>
  import {GuestForgotPassword, GuestLogin, GuestRegister, GuestResendVerificationEmail} from "./jsApi.GEN.js"
  import {onMount, tick} from "svelte";
  import Menu from "./_components/Menu.svelte"
  
  let user = {/* user */};
  let segments = {/* segments */};
  let google = '#{google}';
  
  function getCookie( name ) {
    var match = document.cookie.match( new RegExp( '(^| )' + name + '=([^;]+)' ) );
    if( match ) return match[ 2 ];
  }
  
  // server state
  const title = '#{title}' // /*! title */ {/* title */} [/* title */]
  // TODO: print session or fetch from cookie
  
  // local state
  let email = ''
  let password = ''
  let confirmPass = ''
  
  // binding to element
  let emailInput = {}
  let passInput = {}
  
  const LOGIN = 'LOGIN'
  const REGISTER = 'REGISTER'
  const RESEND_VERIFICATION_EMAIL = 'RESEND-VERIFICATION-EMAIL'
  const FORGOT_PASSWORD = 'FORGOT-PASSWORD'
  const USER = ''
  let mode = LOGIN
  
  async function onHashChange() {
    const auth = getCookie( 'auth' )
    console.log( auth, user )
    if( auth && user && !auth.startsWith( 'TEMP__' ) ) {
      location.hash = ''
      mode = USER
      return
    }
    
    let hash = (location.hash || '')
    if( hash[ 0 ]==='#' ) hash = hash.substring( 1 )
    
    if( hash===LOGIN ) mode = LOGIN
    else if( hash===REGISTER ) mode = REGISTER
    else if( hash===RESEND_VERIFICATION_EMAIL ) mode = RESEND_VERIFICATION_EMAIL
    else if( hash===FORGOT_PASSWORD ) mode = FORGOT_PASSWORD
    else location.hash = LOGIN
    await tick()
    emailInput.focus()
  }
  
  onMount( onHashChange )
  
  async function guestRegister() {
    // TODO: replace all alert with growl
    if( !email ) return alert( 'email is required' )
    if( password.length<12 ) return alert( 'password must be at least 12 characters' )
    if( password!==confirmPass ) return alert( 'passwords do not match' )
    // TODO: send to backend
    const i = {email, password}
    await GuestRegister( i, async function( o ) {
      // TODO: codegen commonResponse (o.error, etc)
      // TODO: codegen list of possible errors
      console.log( o )
      if( o.error ) return alert( o.error )
      alert( 'registered successfully, a registration verification has been sent to your email' )
      mode = LOGIN
      password = ''
      await tick()
      passInput.focus()
    } )
  }
  
  async function guestLogin() {
    if( !email ) return alert( 'email is required' )
    if( password.length<12 ) return alert( 'password must be at least 12 characters' )
    const i = {email, password}
    await GuestLogin( i, function( o ) {
      console.log( o )
      if( o.error ) return alert( o.error );
      user = o.user
      segments = o.segments
      onHashChange()
    } )
  }
  
  async function guestResendVerificationEmail() {
    if( !email ) return alert( 'email is required' )
    const i = {email}
    await GuestResendVerificationEmail( i, function( o ) {
      console.log( o )
      if( o.error ) return alert( o.error );
      onHashChange()
      alert( 'a email verification link has been sent to your email' )
    } )
  }
  
  async function guestForgotPassword() {
    if( !email ) return alert( 'email is required' )
    const i = {email}
    await GuestForgotPassword( i, function( o ) {
      console.log( o )
      if( o.error ) return alert( o.error );
      onHashChange()
      alert( 'a reset password link has been sent to your email' )
    } )
  }
</script>

<svelte:window on:hashchange={onHashChange}/>
{#if mode===USER}
	<Menu access={segments} />
	already logged in
	<hr/>
	TODO: import other svelte component here (menu, content, etc)
{:else}
	<div class="guest">
		<h1>{title} - {mode}</h1>
		<div class="mainContainer">
			<label for="email">Email</label>
			<input type="text" id="email" bind:value={email} bind:this={emailInput}/><br/>
			
			{#if mode===LOGIN || mode===REGISTER}
				<label for="password">Password</label>
				<input type="password" id="password" bind:value={password} bind:this={passInput}><br/>
			{/if}
			
			{#if mode===REGISTER}
				<label for="confirmPass">Confirm Password</label>
				<input type="password" id="confirmPass" bind:value={confirmPass}><br/>
				<label>
					<button on:click={guestRegister}>Register</button>
				</label>
			{/if}
			
			{#if mode===LOGIN}
				<label>
					<button on:click={guestLogin}>Login</button>
				</label>
			{/if}
			
			{#if mode===REGISTER || mode===LOGIN}
				{#if google}
					or
					<a class="button" href={google}>sign in/sign up using Google</a>
				{/if}
			{/if}
			
			{#if mode===RESEND_VERIFICATION_EMAIL}
				<span class="label"></span>
				<button on:click={guestResendVerificationEmail}>Resend Verification Email</button>
			{/if}
			
			{#if mode===FORGOT_PASSWORD}
				<span class="label"></span>
				<button on:click={guestForgotPassword}>Request Reset Password Link</button>
			{/if}
			
			<br/>
			<hr/>
			
			{#if mode!==REGISTER}
				Have no account?
				<a href="#REGISTER" on:click={()=> mode=REGISTER}>register</a><br/>
			{/if}
			{#if mode!==LOGIN}
				Already have account?
				<a href="#LOGIN" on:click={()=> mode=LOGIN}>login</a><br/>
			{/if}
			{#if mode!==RESEND_VERIFICATION_EMAIL}
				Email not yet verified?
				<a href="#{RESEND_VERIFICATION_EMAIL}" on:click={()=> mode=RESEND_VERIFICATION_EMAIL}>request verification email</a><br/>
			{/if}
			{#if mode!==FORGOT_PASSWORD}
				Forgot your password?
				<a href="#{FORGOT_PASSWORD}" on:click={()=> mode=FORGOT_PASSWORD}>request reset password link</a><br/>
			{/if}
		</div>
	</div>
{/if}
<style>
    h1 {
        text-align : center;
    }

    /* pad label and input so they are equal in size */
    label {
        display    : inline-block;
        padding    : 0.2em;
        width      : 11em;
        text-align : right;
    }

    span.label {
        width   : 11em;
        display : inline-block;
    }

    input {
        width   : 16em;
        padding : 0.1em;
        margin  : 0.2em
    }

    button {
        margin-top    : 0.5em;
        padding       : 0.3em;
        border-radius : 0.3em;
    }

    * {
        font-family : sans-serif;
        font-size   : 16pt;
    }

    .mainContainer {
        margin : 0 auto;
    }

    div.guest {
        margin           : 0 auto;
        max-width        : 640px;
        background-color : lightgrey;
        border           : 1px solid grey;
        border-radius    : 1em;
        padding          : 1em
    }
</style>