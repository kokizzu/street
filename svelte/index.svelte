<script>
  /** @typedef {import('./_types/master.js').Access} Access */
  /** @typedef {import('./_types/user.js').User} User */
  /** @typedef {import('./_types/user.js').MostLoggedInUser} MostLoggedInUser */
  /** @typedef {import('./_types/business.js').Revenue} Revenue */
  /** @typedef {import('./_types/business.js').Order} Order */
  /** @typedef {import('./_types/property.js').ScannedAreasToRender} ScannedAreasToRender */
  /** @typedef {import('./_types/property.js').ScannedPropertiesToRender} ScannedPropertiesToRender */
  /**
   * @typedef {Object} UserRegistered
   * @property {string} date
   * @property {number} count
   */
  /**
   * @typedef {Object} RealtorStat
   * @property {string} date
   * @property {number} totalActivity
  */
  /**
   * @typedef {Object} BuyerStat
   * @property {string} date
   * @property {number} totalActivity
  */

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

  const revenues          = /** @type {Revenue[]} */ ([/* revenues */]);
  const orders            = /** @type {Order[]} */ ([/* orders */]);
  const usersRegistered   = /** @type {UserRegistered[]} */ ([/* user_registered */ ]);
  const realtorStats      = /** @type {RealtorStat[]} */ ([/* realtor_stats */ ]);
  const buyerStats        = /** @type {BuyerStat[]} */ ([/* buyer_stats */ ]);
  const usersMostLoggedIn = /** @type {MostLoggedInUser[]} */ ([/* users_most_logged_in */ ]);
  const mostScannedAreas  = /** @type {ScannedAreasToRender[]} */ ([/* most_scanned_areas */ ]);
  const mostScannedProps  = /** @type {ScannedPropertiesToRender[]} */ ([/* most_scanned_properties */ ]);

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

  let chart = /** @type {import('chart.js').Chart} */ (null);

  const STAT_REVENUE    = `revenue`;
  const STAT_REGISTERED = `registered`;
  const STAT_REALTORS   = `realtors`;
  const STAT_ORDERS     = `orders`;
  const STAT_BUYERS     = `buyers`;

  let MODE_STATS = STAT_REVENUE;

  function remove2ndOrdersData() {
    if (MODE_STATS === STAT_ORDERS) {
      if (chart) chart.data.datasets.pop();
    }
  }

  function renderRevenueChart() {
    remove2ndOrdersData()
    MODE_STATS = STAT_REVENUE;
    if (chart) {
      chart.data.labels = (revenues || []).map((/** @type {Revenue} */ i) => {
        const dt = /** @type {Date} */ (new Date(i.salesDate));
        return dt.toLocaleDateString('en-US', {
          month: 'short',
          day: '2-digit'
        });
      });
      chart.data.datasets[0].data = (revenues || []).map((/** @type {Revenue} */ i) => Number(i.revenue));
      chart.data.datasets[0].label = 'Revenue';
      chart.options.scales.y = {
        ticks: {
          stepSize: 10000000,
          callback: function(value) {
            if (Number(value) >= 1000000000) return Number(value) / 1000000000 + 'B';
            if (Number(value) >= 1000000) return Number(value) / 1000000 + 'M';
            if (Number(value) >= 1000) return Number(value) / 1000 + 'K';
            return Number(value);
          }
        }
      };
      chart.update();
    }
  }

  function renderRegisteredChart() {
    remove2ndOrdersData()
    MODE_STATS = STAT_REGISTERED;
    if (chart) {
      chart.data.labels = (usersRegistered || []).map((/** @type {UserRegistered} */ i) => {
        const dt = /** @type {Date} */ (new Date(i.date));
        return dt.toLocaleDateString('en-US', {
          month: 'short',
          day: '2-digit'
        });
      });
      chart.data.datasets[0].data = (usersRegistered || []).map((i) => i.count);
      chart.data.datasets[0].label = 'Registered';
      chart.options.scales.y = {
        ticks: {
          stepSize: 10
        }
      };
      chart.update();
    }
  }

  function renderRealtorsChart() {
    remove2ndOrdersData()
    MODE_STATS = STAT_REALTORS;
    if (chart) {
      chart.data.labels = (realtorStats || []).map((/** @type {RealtorStat} */ i) => {
        const dt = /** @type {Date} */ (new Date(i.date));
        return dt.toLocaleDateString('en-US', {
          month: 'short',
          day: '2-digit'
        });
      });
      chart.data.datasets[0].data = (realtorStats || []).map((i) => i.totalActivity);
      chart.data.datasets[0].label = 'Realtors';
      chart.options.scales.y = {
        ticks: {
          stepSize: 10
        }
      };
      chart.update();
    }
  }

  function renderBuyersChart() {
    remove2ndOrdersData()
    MODE_STATS = STAT_BUYERS;
    if (chart) {
      chart.data.labels = (buyerStats || []).map((/** @type {BuyerStat} */ i) => {
        const dt = /** @type {Date} */ (new Date(i.date));
        return dt.toLocaleDateString('en-US', {
          month: 'short',
          day: '2-digit'
        });
      });
      chart.data.datasets[0].data = (buyerStats || []).map((i) => i.totalActivity);
      chart.data.datasets[0].label = 'Realtors';
      chart.options.scales.y = {
        ticks: {
          stepSize: 10
        }
      };
      chart.update();
    }
  }
  
  function renderOrdersChart() {
    MODE_STATS = STAT_ORDERS;
    if (chart) {
      chart.data.labels = (orders || []).map((/** @type {Order} */ i) => {
        const dt = /** @type {Date} */ (new Date(i.salesDate));
        return dt.toLocaleDateString('en-US', {
          month: 'short',
          day: '2-digit'
        });
      });
      chart.data.datasets[0].data = (orders || []).map((i) => Number(i.revenue));
      chart.data.datasets[0].label = 'Revenue';
      chart.options.scales.y = {
        ticks: {
          stepSize: 10000000,
          callback: function(value) {
            if (Number(value) >= 1000000000) return Number(value) / 1000000000 + 'B';
            if (Number(value) >= 1000000) return Number(value) / 1000000 + 'M';
            if (Number(value) >= 1000) return Number(value) / 1000 + 'K';
            return Number(value);
          }
        }
      };
      chart.data.datasets.push({
        label: 'Total Transactions',
        data: (orders || []).map((i) => i.totalTransaction),
        borderColor: '#3b82f6',
        backgroundColor: '#3b82f630',
        pointRadius: 0,
        tension: 0.1
      })
      chart.update();
    }
  }
  
  onMount(() => {
    onHashChange();
    if (MODE === MODE_USER) {
      setTimeout(() => {
        const ElmChart = /** @type {HTMLCanvasElement} */ (document.getElementById('chart'));
        chart = new Chart(ElmChart, {
          type: 'line',
          data: {
            labels: (revenues || []).map((/** @type {Revenue} */ i) => {
              const dt = /** @type {Date} */ (new Date(i.salesDate));
              return dt.toLocaleDateString('en-US', {
                month: 'short',
                day: '2-digit'
              });
            }),
            datasets: [{
              label: 'Revenue',
              data: (revenues || []).map((/** @type {Revenue} */ i) => Number(i.revenue)),
              borderColor: '#f97316',
              backgroundColor: '#f9731630',
              pointRadius: 0,
              tension: 0.1
            }]
          },
          options: {
            plugins: {
              legend: {
                display: false
              }
            },
            maintainAspectRatio: false,
            responsive: true,
            scales: {
              y: {
                beginAtZero: true,
                ticks: {
                  stepSize: 10000000,
                  callback: function(value) {
                    if (Number(value) >= 1000000000) return Number(value) / 1000000000 + 'B';
                    if (Number(value) >= 1000000) return Number(value) / 1000000 + 'M';
                    if (Number(value) >= 1000) return Number(value) / 1000 + 'K';
                    return Number(value);
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
          <button on:click={renderRevenueChart} class:active={MODE_STATS===STAT_REVENUE} disabled={MODE_STATS===STAT_REVENUE}>
            <span class="block"></span>
            <span class="title">Revenue</span>
          </button>
          <button on:click={renderRegisteredChart} class:active={MODE_STATS===STAT_REGISTERED} disabled={MODE_STATS===STAT_REGISTERED}>
            <span class="block"></span>
            <span class="title">Registered</span>
          </button>
          <button on:click={renderRealtorsChart} class:active={MODE_STATS===STAT_REALTORS} disabled={MODE_STATS===STAT_REALTORS}>
            <span class="block"></span>
            <span class="title">Realtors</span>
          </button>
          <button on:click={renderOrdersChart} class:active={MODE_STATS===STAT_ORDERS} disabled={MODE_STATS===STAT_ORDERS}>
            <span class="block"></span>
            <span class="title">Orders</span>
          </button>
          <button on:click={renderBuyersChart} class:active={MODE_STATS===STAT_BUYERS} disabled={MODE_STATS===STAT_BUYERS}>
            <span class="block"></span>
            <span class="title">Buyers</span>
          </button>
        </nav>
        <div class="chart">
          <canvas id="chart"></canvas>
        </div>
      </div>
      <div class="rows-table">
        <div class="table-root">
          <header>
            <span>Most Logged in Buyers</span>
          </header>
          <div class="table-container">
            <table>
              <thead>
                <tr>
                  <th>Time period</th>
                  <th>Email </th>
                  <th>Full Name</th>
                  <th>Total</th>
                </tr>
              </thead>
              <tbody>
                {#if usersMostLoggedIn && usersMostLoggedIn.length > 0}
                  {#each (usersMostLoggedIn || []) as msUser}
                    <tr>
                      <td>{msUser.time_period}</td>
                      <td>{msUser.email}</td>
                      <td>{msUser.full_name}</td>
                      <td>{msUser.total}</td>
                    </tr>
                  {/each}
                {:else}
                  <tr>
                    <td colspan="4">No data</td>
                  </tr>
                {/if}
              </tbody>
            </table>
          </div>
        </div>
        <div class="table-root">
          <header>
            <span>Most Scanned Areas</span>
          </header>
          <div class="table-container">
            <table>
              <thead>
                <tr>
                  <th>Time period</th>
                  <th>Views</th>
                  <th>City</th>
                  <th>State</th>
                  <th>Coord</th>
                </tr>
              </thead>
              <tbody>
                {#if mostScannedAreas && mostScannedAreas.length > 0}
                  {#each (mostScannedAreas || []) as msArea}
                    <tr>
                      <td>{msArea.time_period}</td>
                      <td>{msArea.views}</td>
                      <td>{msArea.city}</td>
                      <td>{msArea.state}</td>
                      <td>{msArea.latitude}, {msArea.longitude}</td>
                    </tr>
                  {/each}
                {:else}
                  <tr>
                    <td colspan="4">No data</td>
                  </tr>
                {/if}
              </tbody>
            </table>
          </div>
        </div>
      </div>
      <div class="table-root">
        <header>
          <span>Most Scanned Properties</span>
        </header>
        <div class="table-container">
          <table>
            <thead>
              <tr>
                <th>Time period</th>
                <th>Views</th>
                <th>Price</th>
                <th>Size</th>
                <th>City</th>
                <th>State</th>
                <th>Address</th>
              </tr>
            </thead>
            <tbody>
              {#if mostScannedProps && mostScannedProps.length > 0}
                {#each (mostScannedProps || []) as msProperty}
                  <tr>
                    <td>{msProperty.time_period}</td>
                    <td>{msProperty.views}</td>
                    <td>{msProperty.price}</td>
                    <td>{msProperty.total_sqft}</td>
                    <td>{msProperty.city}</td>
                    <td>{msProperty.state}</td>
                    <td>{msProperty.address}</td>
                  </tr>
                {/each}
              {:else}
                <tr>
                  <td colspan="4">No data</td>
                </tr>
              {/if}
            </tbody>
          </table>
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

  .home-container .rows-table {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 20px;
    align-items: start;

  }

  .table-root {
    display: flex;
    flex-direction: column;
    background-color: #fff;
    border-radius: 8px;
    border: 1px solid var(--gray-003);
    padding: 0;
    overflow: hidden;
		width: 100%;
  }

  .table-root header {
    padding: 16px 12px;
  }

  .table-root header span {
    font-weight: 600;
    font-size: 14px;
    color: var(--gray-008);
  }

  .table-root .table-container {
    overflow-x: auto;
    scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: thin;
  }

  .table-root .table-container table {
    width: 100%;
    background: #fff;
    box-shadow: none;
    text-align: left;
    border-collapse: separate;
    border-spacing: 0;
    overflow: hidden;
		font-size: 13px;
  }

  .table-root .table-container table thead {
    box-shadow: none;
    border-bottom: 1px solid var(--gray-003);
    border-top: 1px solid var(--gray-003);
  }

  .table-root .table-container table thead tr th {
    padding: 12px;
		background-color: var(--gray-001);
		text-transform: capitalize;
    border-top: 1px solid var(--gray-003);
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-003);
		min-width: fit-content;
		width: auto;
    text-wrap: nowrap;
  }

  .table-root .table-container table thead tr th:last-child {
    border-right: none;
  }

  .table-root .table-container table tbody tr td {
    padding: 8px 12px;
  }

	.table-root .table-container table tbody tr td {
    padding: 8px 12px;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-004);
  }

	.table-root .table-container table tbody tr:last-child {
		border-bottom: none !important;
  }

	.table-root .table-container table tbody tr:last-child td,
	.table-root .table-container table tbody tr:last-child th {
		border-bottom: none !important;
	}

  .table-root .table-container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .table-root .table-container table tbody tr:last-child td,
  .table-root .table-container table tbody tr:last-child th {
    border-bottom: none !important;
  }

  .table-root .table-container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .table-root .table-container table tbody tr td:last-child {
    border-right: none !important;
  }

  .table-root .table-container table tbody tr th {
    text-align: center;
    border-right: 1px solid var(--gray-004);
    border-bottom: 1px solid var(--gray-004);
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