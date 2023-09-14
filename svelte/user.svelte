<script>
  // @ts-nocheck
  import Menu from './_components/Menu.svelte';
  import ProfileHeader from './_components/ProfileHeader.svelte';
  import Footer from './_components/Footer.svelte';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import {datetime} from './_components/formatter';
  import {onMount} from 'svelte';
  import {UserChangePassword, UserUpdateProfile} from './jsApi.GEN';
  import FaSolidAngleLeft from 'svelte-icons-pack/fa/FaSolidAngleLeft';
  import FaSolidAngleRight from 'svelte-icons-pack/fa/FaSolidAngleRight';
  
  let sideMenuOpen = false;
  
  function openSideMenu() {
    sideMenuOpen = true;
  }
  
  function closeSideMenu() {
    sideMenuOpen = false;
  }
  
  let user = {/* user */};
  let segments = {/* segments */};
  
  let oldPassword = '';
  let newPassword = '';
  let repeatNewPassword = '';
  
  let oldProfileJson = '';
  onMount( () => {
    oldProfileJson = JSON.stringify( user );
  } );
  
  async function updateProfile() {
    if( JSON.stringify( user )===oldProfileJson ) return alert( 'No changes' );
    UserUpdateProfile( user, function( res ) {
      if( res.error ) return alert( res.error );
      oldProfileJson = JSON.stringify( res.user );
      user = res.user;
      alert( 'profile updated' );
    } );
  }
  
  async function changePassword() {
    if( newPassword!==repeatNewPassword ) return alert( 'New password and repeat new password must be same' );
    let input = {
      oldPass: oldPassword,
      newPass: newPassword,
    };
    UserChangePassword( input, function( res ) {
      if( res.error ) return alert( res.error );
      oldPassword = '';
      newPassword = '';
      repeatNewPassword = '';
      alert( 'password changed' );
    } );
  }
</script>

<section class='dashboard'>
	<Menu
		access={segments}
		isSideMenuOpen={sideMenuOpen}
		on:closesidemenu={closeSideMenu}
	/>
	<div class='dashboard_main_content'>
		<ProfileHeader on:opensidemenu={openSideMenu}></ProfileHeader>
		<div class='content'>
			<div class='profile_container'>
				<div class='profile_set'>
					<h2>Update Profile</h2>
					<div class='profile_input'>
						<label for='email'>Email</label>
						<input bind:value={user.email} id='email' type='email'/>
					</div>
					<div class='profile_input'>
						<label for='userName'>Username</label>
						<input bind:value={user.userName} id='userName' type='text'/>
					</div>
					<div class='profile_input'>
						<label for='fullName'>Full Name</label>
						<input bind:value={user.fullName} id='fullName' type='text'/>
					</div>
					<div class='profile_info'>
						<label for='registered'>Registered</label>
						<span id='registered'>{datetime( user.createdAt )}</span>
					</div>
					<div class='profile_info'>
						<label for='lastLogin'>Last login</label>
						<span id='lastLogin'>{datetime( user.lastLoginAt )}</span>
					</div>
					<div class='profile_info'>
						<label for='verified'>Verified</label>
						<span id='verified'>{datetime( user.verifiedAt ) || '0'}</span>
					</div>
					<label for='updateProfile'></label>
					<button id='updateProfile' on:click={updateProfile}>
						<span>SUBMIT</span>
						<Icon color='#FFF' size={18} src={FaSolidAngleLeft}/>
					</button>
				</div>
				
				<div class='password_set'>
					<h2>Change password</h2>
					<div class='profile_input'>
						<label for='oldPassword'>Old Password</label>
						<input bind:value={oldPassword} id='oldPassword' type='password'/>
					</div>
					<div class='profile_input'>
						<label for='newPassword'>New Password</label>
						<input bind:value={newPassword} id='newPassword' type='password'/>
					</div>
					<div class='profile_input'>
						<label for='repeatNewPassword'>Repeat New Password</label>
						<input bind:value={repeatNewPassword} id='repeatNewPassword' type='password'/>
					</div>
					<label for='changePassword'></label>
					<button id='changePassword' on:click={changePassword}>
						<span>SUBMIT</span>
						<Icon color='#FFF' size={18} src={FaSolidAngleRight}/>
					</button>
				</div>
			</div>
		</div>
		<Footer></Footer>
	</div>
</section>

<style>
    .profile_container {
        position       : relative;
        margin-top     : -40px;
        margin-left    : auto;
        margin-right   : auto;
        display        : flex;
        flex-direction : row;
        width          : 88%;
        color          : #475569;
        font-size      : 14px;
    }

    .profile_set, .password_set {
        flex-basis       : 50%;
        display          : flex;
        flex-direction   : column;
        border-radius    : 8px;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
        padding          : 20px;
        background-color : white;
        height           : fit-content;
    }

    .profile_set h2, .password_set h2 {
        font-size   : 18px;
        margin      : 0 0 20px 0;
        font-weight : 700;
        text-align  : center;
        color       : #6366F1;
    }

    .profile_set button, .password_set button {
        margin-left      : auto;
        margin-top       : 20px;
        width            : fit-content;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
        padding          : 6px 20px;
        font-size        : 13px;
        display          : inline-flex;
        font-weight      : 600;
        flex-direction   : row;
        align-items      : center;
        align-content    : center;
        justify-content  : center;
        border           : none;
        background-color : #6366F1;
        border-radius    : 8px;
        color            : white;
        cursor           : pointer;
    }

    .profile_set button:hover, .password_set button:hover {
        background-color : #7E80F1;
    }

    .profile_set {
        margin-right : 15px;
    }

    .password_set {
        margin-left : 15px;
    }

    /* Input box */
    .profile_input {
        display        : flex;
        flex-direction : column;
        width          : 100%;
        margin-bottom  : 10px;
    }

    .profile_input label {
        font-size     : 13px;
        font-weight   : 700;
        margin-left   : 10px;
        margin-bottom : 8px;
    }

    .profile_input input {
        width            : 100%;
        border           : 1px solid #CBD5E1;
        background-color : #F1F5F9;
        border-radius    : 8px;
        padding          : 12px;
    }

    .profile_input input:focus {
        border-color : #3B82F6;
        outline      : 1px solid #3B82F6;
    }

    /* Info */
    .profile_info {
        display       : inline-flex;
        align-items   : center;
        font-size     : 13px;
        margin-left   : 10px;
        margin-bottom : 5px;
    }

    .profile_info label {
        font-weight  : 700;
        margin-right : 10px;
    }
</style>