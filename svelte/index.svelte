<script>
  /** @typedef {import('./_types/master.js').Access} Access */
  /** @typedef {import('./_types/user.js').User} User */

  import {
    GuestForgotPassword, GuestLogin, GuestRegister,
    GuestResendVerificationEmail
  } from './jsApi.GEN.js';
  import {onMount, tick} from 'svelte';
  import Main from './_layouts/Main.svelte';
  import { Icon } from './node_modules/svelte-icons-pack/dist'
  import { FaSolidCircleNotch } from './node_modules/svelte-icons-pack/dist/fa';
  import { notifier } from './_components/notifier.js';
  import InputBox from './_components/InputBox.svelte';
  import Chart from 'chart.js/auto';
  
  let title     = /** @type {string} */ ('#{title}');
  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let google    = /** @type {string} */ ('#{google}');
  let apple     = /** @type {string} */ ('#{apple}');

  // Generate Apple OAuth URL
  const clientId      = 'com.hapstr.app'; //
  const redirectUri   = 'https://admin.hapstr.xyz/guest/oauthCallback'; // Your frontend callback URL
  const state         = 'random_state_value'; // Use a random string to prevent CSRF attacks
  const scope         = 'email';
  const response_mode = 'form_post'

  apple = `https://appleid.apple.com/auth/authorize?response_type=code&client_id=${clientId}&redirect_uri=${encodeURIComponent(redirectUri)}&scope=${scope}&state=${state}&response_mode=${response_mode}`;
  
  function getCookie( name ) {
    var match = document.cookie.match( new RegExp( '(^| )' + name + '=([^;]+)' ) );
    if( match ) return match[ 2 ];
  }
  
  // local state
  let email       = /** @type {string} */ ('');
  let password    = /** @type {string} */ ('');
  let confirmPass = /** @type {string} */ ('');
  
  // binding to element
  let passInput   = /** @type {HTMLInputElement} */ ({});
  
  const MODE_LOGIN        = /** @type {string} */ ('LOGIN');
  const MODE_REGISTER     = /** @type {string} */ ('REGISTER');
  const MODE_VERIF_EMAIL  = /** @type {string} */ ('RESEND_VERIFICATION_EMAIL');
  const MODE_FG_PASSWD    = /** @type {string} */ ('FORGOT_PASSWORD');
  const MODE_USER         = /** @type {string} */ ('');

  let MODE = /** @type {string} */ (MODE_LOGIN);

  let isSubmitted = /** @type {boolean} */ (false);
  
  async function onHashChange() {
    const auth = getCookie( 'auth' );
    if( auth && user && !auth.startsWith( 'TEMP__' ) ) {
      location.hash = MODE_USER;
      MODE = MODE_USER;
      return;
    }
    
    let hash = /** @type {string} */ (location.hash || '');

    if ( hash[ 0 ] === '#' ) hash = hash.substring( 1 );
    
    switch ( hash ) {
      case MODE_LOGIN:
        MODE = MODE_LOGIN;
        break;
      case MODE_REGISTER:
        MODE = MODE_REGISTER;
        break;
      case MODE_VERIF_EMAIL:
        MODE = MODE_VERIF_EMAIL;
        break;
      case MODE_FG_PASSWD:
        MODE = MODE_FG_PASSWD;
        break;
      default:
        location.hash = MODE_LOGIN;
    }

    await tick();
  }
  
  onMount(() => {
    onHashChange();
    if (MODE === MODE_USER) {
      setTimeout(() => {
        const ElmRevenueChart = /** @type {HTMLCanvasElement} */ (document.getElementById('revenue-chart'));
        new Chart(ElmRevenueChart, {
          type: 'line',
          data: {
            labels: ['Jan 31', 'Feb 28', 'Mar 30', 'Apr 30', 'May 30', 'Jun 30'],
            datasets: [{
              label: 'Revenue',
              data: [40000, 130000, 70000, 250000, 145000, 220000],
              borderColor: '#f97316',
              backgroundColor: '#f9731630',
              pointRadius: 0,
              tension: 0.1
            }]
          },
          options: {
            plugins: {
              legend: {
                rtl: true,
                labels: {
                  usePointStyle: true,
                  color: 'var(--gray-008)',
                  textAlign: 'right',
                }
              }
            },
            maintainAspectRatio: false,
            responsive: true,
            scales: {
              y: {
                beginAtZero: true,
                ticks: {
                  stepSize: 100000,
                  callback: function(value) {
                    return value >= 1000 ? value / 1000 + 'K' : value;
                  }
                }
              }
            }
          }
        });
      }, 400);
    }
  });
  
  async function guestRegister() {
    isSubmitted = true;
    if( !email ) {
      isSubmitted = false;
      notifier.showError('Email is required' );
      return
    }
    if( password.length<12 ) {
      isSubmitted = false;
      notifier.showError('Password must be at least 12 characters' );
      return
    }
    if( password!==confirmPass ) {
      isSubmitted = false;
      notifier.showError('Passwords do not match' );
      return
    }

    const i = {email, password};
    await GuestRegister( i, async function( /** @type any */ o ) {
      if( o.error ) {
        isSubmitted = false;
        notifier.showError(o.error );
		return
      }
      isSubmitted = false;
      notifier.showSuccess('Registered successfully, a registration verification has been sent to your email' );
      MODE = MODE_LOGIN;
      password = '';
      await tick();
      passInput.focus();
    } );
  }
  
  async function guestLogin() {
    isSubmitted = true;
    if( !email ) {
      isSubmitted = false;
      notifier.showError('Email is required' );
      return
    }
    if( password.length<12 ) {
      isSubmitted = false;
      notifier.showError('Password must be at least 12 characters' );
      return
    }
    const i = {email, password};
    await GuestLogin( i, async function(/** @type any */ o ) {
      if( o.error ) {
        isSubmitted = false;
        notifier.showError( o.error );
        return
      }
      isSubmitted = false;
      notifier.showSuccess('Login successfully' );
      setTimeout( () => {
        user = o.user;
        segments = o.segments;
        onHashChange();

        window.document.location = '/realtor';
      }, 1500 );
    } );
  }
  
  async function guestResendVerificationEmail() {
    isSubmitted = true;
    if( !email ) {
      isSubmitted = false;
      notifier.showError( 'Email is required' );
      return
    }
    const i = {email};
    await GuestResendVerificationEmail( i, async function( /** @type any */ o ) {
      if( o.error ) {
        isSubmitted = false;
        notifier.showError(o.error );
        return
      }
      isSubmitted = false;
      onHashChange();
      notifier.showInfo('An email verification link has been sent to your email' );
    } );
  }
  
  async function guestForgotPassword() {
    isSubmitted = true;
    if( !email ) {
      isSubmitted = false;
      notifier.showError('Email is required' );
      return
    }
    const i = {email};
    await GuestForgotPassword( i, async function(/** @type any */ o  ) {
      if( o.error ) {
        isSubmitted = false;
        notifier.showError( o.error );
        return
      }
      onHashChange();
      notifier.showInfo('A reset password link has been sent to your email' );
    } );
  }
</script>

<svelte:window on:hashchange={onHashChange}/>
{#if MODE === MODE_USER}
  <Main {user} access={segments}>
    <div class="home-container">
      <div class="stats-chart">
        <nav>
          <button class="active">
            <span class="block"></span>
            <span class="title">Revenue</span>
          </button>
          <button>
            <span class="block"></span>
            <span class="title">Registered</span>
          </button>
          <button>
            <span class="block"></span>
            <span class="title">Realtors</span>
          </button>
          <button>
            <span class="block"></span>
            <span class="title">Orders</span>
          </button>
        </nav>
        <div class="chart">
          <canvas id="revenue-chart"></canvas>
        </div>
      </div>
    </div>
  </Main>
{:else}
	<section class="auth-section">
		<div class="main-container">
			<div class="title-container">
				<p>{title}</p>
				<h1>{MODE.split( '_' ).join( ' ' )}</h1>
			</div>
			<div class="form-container">
				<div class="input-container">
					{#if MODE === MODE_LOGIN
            || MODE === MODE_REGISTER
            || MODE === MODE_VERIF_EMAIL
            || MODE === MODE_FG_PASSWD
          }
            <InputBox
              id="email"
              label="Email"
              type="text"
              bind:value={email}
            />
					{/if}

					{#if MODE === MODE_LOGIN || MODE === MODE_REGISTER}
            <InputBox
              id="password"
              label="Password"
              type="password"
              bind:value={password}
            />
					{/if}

					{#if MODE === MODE_REGISTER}
            <InputBox
              id="confirmPass"
              label="Confirm Password"
              type="password"
              bind:value={confirmPass}
            />
					{/if}
				</div>
				<!-- Forgot Password -->
				{#if MODE===MODE_LOGIN}
					<p class="forgot-password">
						Forgot Password?
						<a href="#FORGOT_PASSWORD" on:click|preventDefault={() => (MODE = MODE_FG_PASSWD)}>Reset here</a>
					</p>
				{/if}
				<div class="button-container">
					{#if MODE===MODE_REGISTER}
						<button on:click={guestRegister}>
							{#if isSubmitted===true}
								<Icon className="spin" color='#FFF' size="15" src={FaSolidCircleNotch}/>
							{/if}
							{#if isSubmitted===false}
								<span>Register</span>
							{/if}
						</button>
					{/if}
					{#if MODE===MODE_LOGIN}
						<button on:click={guestLogin}>
							{#if isSubmitted===true}
								<Icon className="spin" color='#FFF' size="15" src={FaSolidCircleNotch}/>
							{/if}
							{#if isSubmitted===false}
								<span>Login</span>
							{/if}
						</button>
					{/if}
					{#if MODE=== MODE_VERIF_EMAIL}
						<button on:click={guestResendVerificationEmail}>
							{#if isSubmitted===true}
								<Icon className="spin" color='#FFF' size="15" src={FaSolidCircleNotch}/>
							{/if}
							{#if isSubmitted===false}
								<span>Resend Verification Email</span>
							{/if}
						</button>
					{/if}
					{#if MODE===MODE_FG_PASSWD}
						<button on:click={guestForgotPassword}>
							{#if isSubmitted===true}
								<Icon className="spin" color='#FFF' size="15" src={FaSolidCircleNotch}/>
							{/if}
							{#if isSubmitted===false}
								<span>Request Reset Password Link</span>
							{/if}
						</button>
					{/if}
				</div>
				<!-- Oauth Buttons -->
				{#if MODE===MODE_REGISTER || MODE===MODE_LOGIN}
					<div class="oauth-container">
						<div class="or-separator">
							<span/>
							<p>or</p>
							<span/>
						</div>
            <div class="oauth-buttons">
              <!-- Google OAuth -->
              {#if google}
                <a class="button" href={google}>
                  <img src="/assets/icons/google.svg" alt="Google"/>
                  <span>Continue with Google</span>
                </a>
              {/if}
              <!-- Apple OAuth -->
              {#if apple}
                <a class="button" href={apple}>
                  <img src="/assets/icons/apple.png" alt="Apple"/>
                  <span>Continue with Apple</span>
                </a>
              {/if}
            </div>
					</div>
				{/if}
				<div class="foot-auth">
					{#if MODE!==MODE_REGISTER}
						<p>Have no account? <a href="#REGISTER" on:click={() => (MODE = MODE_REGISTER)}>register</a></p>
					{/if}
					{#if MODE!==MODE_LOGIN}
						<p>Already have account? <a href="#LOGIN" on:click={() => (MODE = MODE_LOGIN)}>login</a></p>
					{/if}
					{#if MODE!==MODE_VERIF_EMAIL}
						<p>
							Email not yet verified? <a
							href="#RESEND_VERIFICATION_EMAIL"
							on:click={() => (MODE = MODE_VERIF_EMAIL)}>request verification email</a
						>
						</p>
					{/if}
				</div>
			</div>
		</div>
	</section>
{/if}

<style>
  @keyframes spin {
    from {
      transform : rotate(0deg);
    }
    to {
      transform : rotate(360deg);
    }
  }

  :global(.spin) {
    animation : spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
  }

  .home-container {
    display: flex;
    flex-direction: column;
    gap: 20px;
    width: 100%;
    padding: 20px;
  }

  .home-container .stats-chart {
    display: flex;
    flex-direction: column;
    gap: 0;
    padding: 0;
    border: 1px solid var(--gray-003);
    border-radius: 8px;
    height: 350px;
    width: 100%;
  }

  .home-container .stats-chart nav {
    display: flex;
    flex-direction: row;
    gap: 10px;
    height: fit-content;
    padding: 0 16px;
  }

  .home-container .stats-chart nav button {
    border: none;
    background-color: transparent;
    color: var(--gray-008);
    font-size: 15px;
    cursor: pointer;
    display: flex;
    flex-direction: column;
    align-items: center;
    height: fit-content;
    padding: 0;
  }

  .home-container .stats-chart nav button:hover {
    background-color: var(--gray-001);
  }

  .home-container .stats-chart nav button.active {
    color: var(--orange-005);
  }

  .home-container .stats-chart nav button .title {
    flex-grow: 1;
    padding: 15px 10px;
    display: flex;
    align-items: center;
  }

  .home-container .stats-chart nav button .block {
    height: 3px;
		background-color: transparent;
		width: 100%;
  }

  .home-container .stats-chart nav button:hover .block {
		background-color: var(--gray-005);
	}
	.home-container .stats-chart nav button.active .block {
		background-color: var(--orange-005);
	}

  .home-container .stats-chart .chart {
    height: 92%;
    width: 100%;
    display: flex;
    width: 100%;
    padding: 0 16px 45px 16px;
  }

  .home-container .stats-chart .chart canvas {
    width: 100% !important;
    height: 100% !important;
  }

  .auth-section {
    height: 100%;
    width: 100%;
    background-color: var(--gray-001);
    filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    display: flex;
    color: var(--gray-008);
  }

  .auth-section .main-container {
    width: 480px;
    height: fit-content;
    padding: 20px;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    background-color: white;
    margin: 50px auto;
    border: 1px solid var(--gray-003);
    gap: 10px;
  }

  .auth-section .main-container .title-container {
    display: flex;
    flex-direction: column;
    gap: 10px;
    width: 100%;
    text-align: center;
  }

  .auth-section .main-container .title-container p {
    font-size: 16px;
    font-weight: 600;
    color: var(--orange-005);
    margin: 0;
  }

  .auth-section .main-container .title-container h1 {
    margin: 0 0 5px 0;
    font-size: 22px;
    font-weight: 700;
  }

  .auth-section .main-container .form-container {
    display: flex;
    flex-direction: column;
    gap: 20px;
    width: 100%;
  }

  .auth-section .main-container .form-container .input-container {
    display: flex;
    flex-direction: column;
    gap: 15px;
  }

  .auth-section .main-container .form-container .forgot-password {
    width: 100%;
    text-align: center;
    font-weight: 500;
    margin: 0;
  }

  .auth-section .main-container .form-container .forgot-password a {
    color: var(--blue-006);
    text-decoration: none;
  }

  .auth-section .main-container .form-container .forgot-password a:hover {
    color: var(--blue-005);
    text-decoration: underline;
  }

  .auth-section .main-container .form-container .button-container {
    height: fit-content;
    width: 100%;
    display: flex;
  }

  .auth-section .main-container .form-container .button-container button {
    margin: 0;
    width: 100%;
    padding: 12px;
    font-size: 15px;
    font-weight: 600;
    background-color: var(--orange-006);
    border-radius: 8px;
    color: white;
    border: none;
    cursor: pointer;
  }

  .auth-section .main-container .form-container .button-container button:hover {
    background-color : var(--orange-005);
  }

  .auth-section .main-container .form-container .oauth-container {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .auth-section .main-container .form-container .oauth-container .or-separator {
    display: flex;
    flex-direction: row;
    align-items: center;
    width: 100%;
  }

  .auth-section .main-container .form-container .oauth-container .or-separator span {
    flex-grow: 1;
    height: 0;
    border-top: 1px solid var(--gray-002);
    padding: 0;
  }

  .auth-section .main-container .form-container .oauth-container .or-separator p {
    width       : fit-content;
    font-weight : 600;
    padding     : 0 10px;
    margin: 0;
  }

  .auth-section .main-container .form-container .oauth-container .oauth-buttons {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .auth-section .main-container .form-container .oauth-container .button {
    padding: 10px;
    background-color: white;
    border: 1px solid var(--gray-002);
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    font-weight: 600;
    border-radius: 8px;
    text-decoration: none;
    color: var(--gray-008);
  }

  .auth-section .main-container .form-container .oauth-container .button:hover {
    background-color: var(--gray-001);
    border: 1px solid var(--gray-003);
  }

  .auth-section .main-container .form-container .oauth-container .button img {
    width: 20px;
    height: auto;
  }

  .auth-section .main-container .form-container .oauth-container .button span {
    margin-left: 8px;
  }

  .auth-section .main-container .form-container .foot-auth {
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .auth-section .main-container .form-container .foot-auth p {
    font-weight: 500;
    margin: 0;
  }

  .auth-section .main-container .form-container .foot-auth a {
    color: var(--blue-006);
    text-decoration: none;
  }

  .auth-section .main-container .form-container .foot-auth a:hover {
    color: var(--blue-005);
    text-decoration: underline;
  }
</style>