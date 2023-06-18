<script>
  import Menu from './_components/Menu.svelte';
  import { datetime } from './_components/formatter';
  import { onMount } from 'svelte';
  import { UserChangePassword, UserUpdateProfile } from './jsApi.GEN.js';
  
  let user = {/* user */ };
  let segments = {/* segments */ };
  
  let oldPassword = '';
  let newPassword = '';
  let repeatNewPassword = '';
  
  let oldProfileJson = '';
  onMount( () => {
    oldProfileJson = JSON.stringify( user );
  } );
  
  async function updateProfile() {
    if( JSON.stringify( user )==oldProfileJson ) return alert( 'No changes' );
    UserUpdateProfile( user, function( res ) {
      if( res.error ) return alert( res.error );
      oldProfileJson = JSON.stringify( res.user );
      user = res.user;
      alert( 'profile updated' );
    } );
  }
  
  async function changePassword() {
    if( newPassword!=repeatNewPassword ) return alert( 'New password and repeat new password must be same' );
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

<Menu access={segments} />

<hr />

<h2>Update Profile</h2><br />
<label for='email'>Email</label>
<input id='email' type='email' bind:value={user.email} /><br />
<label for='userName'>Username</label>
<input id='userName' type='text' bind:value={user.userName} /><br />
<label for='fullName'>Full Name</label>
<input id='fullName' type='text' bind:value={user.fullName} /><br />
<label for='registered'>Registered</label>
<span id='registered'>{datetime( user.createdAt )}</span><br />
<label for='lastLogin'>Last login</label>
<span id='lastLogin'>{datetime( user.lastLoginAt )}</span><br />
<label for='verified'>Verified</label>
<span id='verified'>{datetime( user.verifiedAt )}</span><br />

<label for='updateProfile'></label>
<button id='updateProfile' on:click={updateProfile}>Update Profile</button>

<hr />

<h2>Change password</h2>


<label for='oldPassword'>Old Password</label>
<input id='oldPassword' type='password' bind:value={oldPassword} /><br />

<label for='newPassword'>New Password</label>
<input id='newPassword' type='password' bind:value={newPassword} /><br />
<label for='repeatNewPassword'>Repeat New Password</label>
<input id='repeatNewPassword' type='password' bind:value={repeatNewPassword} /><br />

<label for='changePassword'></label>
<button id='changePassword' on:click={changePassword}>Change Password</button>