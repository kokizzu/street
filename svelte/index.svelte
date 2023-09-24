<script>
  //@ts-nocheck
  import {GuestForgotPassword, GuestLogin, GuestRegister, GuestResendVerificationEmail} from './jsApi.GEN.js';
  import {onMount, tick} from 'svelte';
  import Menu from './_components/Menu.svelte';
  import PropertyLocation from '_components/PropertyLocation.svelte';
  import ProfileHeader from './_components/ProfileHeader.svelte';
  import Footer from './_components/Footer.svelte';
  
  let sideMenuOpen = false;
  let randomProps = [/* randomProps */] || [];
  let initialLatLong = [/* initialLatLong */];
  let defaultDistanceKm = +'#{defaultDistanceKm}';
  
  function openSideMenu() {
    sideMenuOpen = true;
  }
  
  function closeSideMenu() {
    sideMenuOpen = false;
  }
  
  let user = {/* user */};
  let segments = {/* segments */};
  let google = '#{google}';
  
  function getCookie( name ) {
    var match = document.cookie.match( new RegExp( '(^| )' + name + '=([^;]+)' ) );
    if( match ) return match[ 2 ];
  }
  
  // server state
  const title = '#{title}'; // /*! title */ {/* title */} [/* title */]
  // TODO: print session or fetch from cookie
  
  // local state
  let email = '';
  let password = '';
  let confirmPass = '';
  
  // binding to element
  let emailInput = {};
  let passInput = {};
  
  const LOGIN = 'LOGIN';
  const REGISTER = 'REGISTER';
  const RESEND_VERIFICATION_EMAIL = 'RESEND_VERIFICATION_EMAIL';
  const FORGOT_PASSWORD = 'FORGOT_PASSWORD';
  const USER = '';
  let mode = LOGIN;
  
  async function onHashChange() {
    const auth = getCookie( 'auth' );
    console.log( auth, user );
    if( auth && user && !auth.startsWith( 'TEMP__' ) ) {
      location.hash = '';
      mode = USER;
      return;
    }
    
    let hash = location.hash || '';
    if( hash[ 0 ]==='#' ) hash = hash.substring( 1 );
    
    if( hash===LOGIN ) mode = LOGIN;
    else if( hash===REGISTER ) mode = REGISTER;
    else if( hash===RESEND_VERIFICATION_EMAIL ) mode = RESEND_VERIFICATION_EMAIL;
    else if( hash===FORGOT_PASSWORD ) mode = FORGOT_PASSWORD;
    else location.hash = LOGIN;
    await tick();
    emailInput.focus();
  }
  
  onMount(() => {
    onHashChange();
    console.log("User = ", user)
  } )
  
  async function guestRegister() {
    // TODO: replace all alert with growl
    if( !email ) return alert( 'email is required' );
    if( password.length<12 ) return alert( 'password must be at least 12 characters' );
    if( password!==confirmPass ) return alert( 'passwords do not match' );
    // TODO: send to backend
    const i = {email, password};
    await GuestRegister( i, async function( o ) {
      // TODO: codegen commonResponse (o.error, etc)
      // TODO: codegen list of possible errors
      console.log( o );
      if( o.error ) return alert( o.error );
      alert( 'registered successfully, a registration verification has been sent to your email' );
      mode = LOGIN;
      password = '';
      await tick();
      passInput.focus();
    } );
  }
  
  async function guestLogin() {
    if( !email ) return alert( 'email is required' );
    if( password.length<12 ) return alert( 'password must be at least 12 characters' );
    const i = {email, password};
    await GuestLogin( i, function( o ) {
      console.log( o );
      if( o.error ) return alert( o.error );
      user = o.user;
      segments = o.segments;
      onHashChange();
      window.document.location = '/'
    } );
  }
  
  async function guestResendVerificationEmail() {
    if( !email ) return alert( 'email is required' );
    const i = {email};
    await GuestResendVerificationEmail( i, function( o ) {
      console.log( o );
      if( o.error ) return alert( o.error );
      onHashChange();
      alert( 'a email verification link has been sent to your email' );
    } );
  }
  
  async function guestForgotPassword() {
    if( !email ) return alert( 'email is required' );
    const i = {email};
    await GuestForgotPassword( i, function( o ) {
      console.log( o );
      if( o.error ) return alert( o.error );
      onHashChange();
      alert( 'a reset password link has been sent to your email' );
    } );
  }
</script>

<svelte:window on:hashchange={onHashChange}/>
{#if mode===USER}
	<section class="dashboard">
		<Menu access={segments} isSideMenuOpen={sideMenuOpen} on:closesidemenu={closeSideMenu}/>
		<div class="dashboard_main_content">
			<ProfileHeader on:opensidemenu={openSideMenu}/>
			<div class="content">
				<PropertyLocation initialLatLong={initialLatLong} randomProps={randomProps}/>
			</div>
			<Footer/>
		</div>
	</section>
{:else}
	<section class="auth_section">
		<div class="main_container">
			<div class="title_container">
				<p>{title}</p>
				<h1>{mode.split( '_' ).join( ' ' )}</h1>
			</div>
			<div class="sign_in_container">
				<div class="input_container">
					{#if mode===LOGIN || mode===REGISTER || mode===RESEND_VERIFICATION_EMAIL || mode===FORGOT_PASSWORD}
						<div class="input_box">
							<label for="email">Email</label>
							<input type="text" id="email" bind:value={email} bind:this={emailInput}/>
						</div>
					{/if}
					{#if mode===LOGIN || mode===REGISTER}
						<div class="input_box">
							<label for="password">Password</label>
							<input type="password" id="password" bind:value={password} bind:this={passInput}/>
						</div>
					{/if}
					{#if mode===REGISTER}
						<div class="input_box">
							<label for="confirmPass">Confirm Password</label>
							<input type="password" id="confirmPass" bind:value={confirmPass}/>
						</div>
					{/if}
				</div>
				<!-- Forgot Password -->
				{#if mode===LOGIN}
					<p class="forgot_password">
						Forgot Password?
						<a href="#FORGOT_PASSWORD" on:click|preventDefault={() => (mode = FORGOT_PASSWORD)}>Reset here</a>
					</p>
				{/if}
				<div class="button_container">
					{#if mode===REGISTER}
						<button on:click={guestRegister}>Register</button>
					{/if}
					{#if mode===LOGIN}
						<button on:click={guestLogin}>Login</button>
					{/if}
					{#if mode===RESEND_VERIFICATION_EMAIL}
						<button on:click={guestResendVerificationEmail}>Resend Verification Email</button>
					{/if}
					{#if mode===FORGOT_PASSWORD}
						<button on:click={guestForgotPassword}>Request Reset Password Link</button>
					{/if}
				</div>
				<!-- Oauth Buttons -->
				{#if mode===REGISTER || mode===LOGIN}
					<div class="oauth_container">
						<div class="or_separator">
							<span/>
							<p>or</p>
							<span/>
						</div>
						<!-- Google OAuth -->
						{#if google}
							<a class="button" href={google}>
								<img src="/assets/icons/google.svg" alt="Google"/>
								<span>Continue with Google</span>
							</a>
						{/if}
					</div>
				{/if}
				<div class="foot_auth">
					{#if mode!==REGISTER}
						<p>Have no account? <a href="#REGISTER" on:click={() => (mode = REGISTER)}>register</a></p>
					{/if}
					{#if mode!==LOGIN}
						<p>Already have account? <a href="#LOGIN" on:click={() => (mode = LOGIN)}>login</a></p>
					{/if}
					{#if mode!==RESEND_VERIFICATION_EMAIL}
						<p>
							Email not yet verified? <a
							href="#RESEND_VERIFICATION_EMAIL"
							on:click={() => (mode = RESEND_VERIFICATION_EMAIL)}>request verification email</a
						>
						</p>
					{/if}
				</div>
			</div>
		</div>
	</section>
{/if}

<style>
    .auth_section {
        height           : 100%;
        width            : 100%;
        background-color : #F1F5F9;
        display          : flex;
        color            : #475569;
    }

    .main_container {
        width            : 480px;
        height           : fit-content;
        padding          : 20px;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
        border-radius    : 15px;
        display          : flex;
        flex-direction   : column;
        background-color : white;
        margin           : 50px auto;
        border           : 1px solid #CBD5E1;
    }

    .title_container {
        display        : flex;
        flex-direction : column;
        width          : 100%;
        text-align     : center;
    }

    .title_container p {
        font-size   : 16px;
        font-weight : 600;
        color       : #EF4444;
        margin      : 0;
    }

    .title_container h1 {
        margin      : 5px 0 0 0;
        font-size   : 22px;
        font-weight : 700;
    }

    .input_container {
        display        : flex;
        flex-direction : column;
        margin-bottom  : 15px;
    }

    .input_container .input_box {
        display        : flex;
        flex-direction : column;
        width          : 100%;
        margin-top     : 10px;
    }

    .input_container .input_box label {
        font-size     : 13px;
        font-weight   : 700;
        margin-left   : 10px;
        margin-bottom : 8px;
    }

    .input_container .input_box input {
        width            : 100%;
        border           : 1px solid #CBD5E1;
        background-color : #F1F5F9;
        border-radius    : 8px;
        padding          : 12px;
    }

    .input_container .input_box input:focus {
        border-color : #3B82F6;
        outline      : 1px solid #3B82F6;
    }

    .forgot_password {
        margin-top    : 7px;
        margin-bottom : 15px;
        width         : 100%;
        text-align    : center;
        font-size     : 14px;
        font-weight   : 600;
    }

    .forgot_password a {
        color           : #3B82F6;
        text-decoration : none;
    }

    .forgot_password a:hover {
        color           : #5892F5;
        text-decoration : underline;
    }

    .button_container button {
        margin           : 0;
        width            : 100%;
        padding          : 10px;
        font-size        : 16px;
        font-weight      : 700;
        background-color : #3B82F6;
        border-radius    : 8px;
        color            : white;
        border           : none;
        cursor           : pointer;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    }

    .button_container button:hover {
        background-color : #5892F5;
    }

    .oauth_container .or_separator {
        display        : flex;
        flex-direction : row;
        align-items    : center;
        width          : 100%;
    }

    .oauth_container .or_separator span {
        flex-grow  : 1;
        height     : 0;
        border-top : 1px solid #CBD5E1;
        padding    : 0;
    }

    .oauth_container .or_separator p {
        width       : fit-content;
        font-weight : 600;
        padding     : 0 10px;
    }

    .oauth_container .button {
        padding          : 10px;
        background-color : white;
        border           : 1px solid #CBD5E1;
        display          : flex;
        flex-direction   : row;
        align-items      : center;
        justify-content  : center;
        font-weight      : 600;
        border-radius    : 8px;
        text-decoration  : none;
        color            : #334155;
    }

    .oauth_container .button:hover {
        background-color : #F1F5F9;
        /* #94a3b8 */
    }

    .oauth_container .button img {
        width  : 20px;
        height : auto;
    }

    .oauth_container .button span {
        margin-left : 8px;
    }

    .foot_auth {
        margin-top     : 10px;
        display        : flex;
        flex-direction : column;
    }

    .foot_auth p {
        margin-top    : 10px;
        margin-bottom : 0;
        text-align    : center;
        font-weight   : 600;
    }

    .foot_auth a {
        color           : #3B82F6;
        text-decoration : none;
    }

    .foot_auth a:hover {
        color           : #5892F5;
        text-decoration : underline;
    }
</style>
