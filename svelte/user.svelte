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
  import FaSolidTrashAlt from "svelte-icons-pack/fa/FaSolidTrashAlt";
  
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

<section class="dashboard">
	<Menu access={segments}/>
	<div class="dashboard_main_content">
		<ProfileHeader />
		<div class="content">
			<div class="profile_details_container">
				<div class="left">
					<div class="profile_details">
						<h2>Profile Details</h2>
						<div class="profile_pictures">
							<div class="img_container">
								<img alt="profile" src="/assets/img/team-1-200x200.jpg"/>
							</div>
							<div class="actions">
								<button class='btn_upload_photo'>Upload Profile Photo</button>
								<button class='btn_delete'>
									<Icon color="#FFF" size={15} src={FaSolidTrashAlt}/>
								</button>
							</div>
						</div>
						<div class="input_container">
							<div class="name">
								<div class="profile_input">
									<label for="userName">Username</label>
									<input bind:value={user.userName} id="userName" type="text"/>
								</div>
								<div class="profile_input">
									<label for="fullName">Full Name</label>
									<input bind:value={user.fullName} id="fullName" type="text"/>
								</div>
							</div>
							<div class="profile_input email">
								<label for="email">Email</label>
								<input bind:value={user.email} id="email" type="email"/>
							</div>
						</div>
						<div class="info_container">
							<div class="profile_info">
								<label for="registered">Registered:</label>
								<span id="registered">{datetime( user.createdAt )}</span>
							</div>
							<div class="profile_info">
								<label for="lastLogin">Last login:</label>
								<span id="lastLogin">{datetime( user.lastLoginAt )}</span>
							</div>
							<div class="profile_info">
								<label for="verified">Verified:</label>
								<span id="verified">{datetime( user.verifiedAt ) || '0'}</span>
							</div>
						</div>
						<label for="updateProfile"/>
						<button id="updateProfile" on:click={updateProfile}>
							<span>SUBMIT</span>
							<Icon color="#FFF" size={18} src={FaSolidAngleLeft}/>
						</button>
					</div>
				</div>
				
				<div class="right">
					<div class="password_set">
						<h2>Change password</h2>
						<div class="input_container">
							<div class="profile_input">
								<label for="oldPassword">Old Password</label>
								<input bind:value={oldPassword} id="oldPassword" type="password"/>
							</div>
							<div class="profile_input">
								<label for="newPassword">New Password</label>
								<input bind:value={newPassword} id="newPassword" type="password"/>
							</div>
							<div class="profile_input">
								<label for="repeatNewPassword">Repeat New Password</label>
								<input bind:value={repeatNewPassword} id="repeatNewPassword" type="password"/>
							</div>
							<label for="changePassword"/>
							<button id="changePassword" on:click={changePassword}>
								<span>SUBMIT</span>
								<Icon color="#FFF" size={18} src={FaSolidAngleRight}/>
							</button>
						</div>
					</div>
					<div class="country_details">
						<h2>Country Details</h2>
						<div class="country_list">
							<select id="country" name="country">
								<option value="Afghanistan">Afghanistan</option>
								<option value="Albania">Albania</option>
								<option value="Algeria">Algeria</option>
								<option value="Andorra">Andorra</option>
								<option value="Angola">Angola</option>
								<option value="Antigua and Barbuda">Antigua and Barbuda</option>
								<option value="Argentina">Argentina</option>
								<option value="Armenia">Armenia</option>
								<option value="Australia">Australia</option>
								<option value="Austria">Austria</option>
								<option value="Azerbaijan">Azerbaijan</option>
								<option value="Bahamas">Bahamas</option>
								<option value="Bahrain">Bahrain</option>
								<option value="Bangladesh">Bangladesh</option>
								<option value="Barbados">Barbados</option>
								<option value="Belarus">Belarus</option>
								<option value="Belgium">Belgium</option>
								<option value="Belize">Belize</option>
								<option value="Benin">Benin</option>
								<option value="Bhutan">Bhutan</option>
								<option value="Bolivia">Bolivia</option>
								<option value="Bosnia and Herzegovina">Bosnia and Herzegovina</option>
								<option value="Botswana">Botswana</option>
								<option value="Brazil">Brazil</option>
								<option value="Brunei">Brunei</option>
								<option value="Bulgaria">Bulgaria</option>
								<option value="Burkina Faso">Burkina Faso</option>
								<option value="Burundi">Burundi</option>
								<option value="Côte d'Ivoire">Côte d'Ivoire</option>
								<option value="Cabo Verde">Cabo Verde</option>
								<option value="Cambodia">Cambodia</option>
								<option value="Cameroon">Cameroon</option>
								<option value="Canada">Canada</option>
								<option value="Central African Republic">Central African Republic</option>
								<option value="Chad">Chad</option>
								<option value="Chile">Chile</option>
								<option value="China">China</option>
								<option value="Colombia">Colombia</option>
								<option value="Comoros">Comoros</option>
								<option value="Congo (Congo-Brazzaville)">Congo (Congo-Brazzaville)</option>
								<option value="Costa Rica">Costa Rica</option>
								<option value="Croatia">Croatia</option>
								<option value="Cuba">Cuba</option>
								<option value="Cyprus">Cyprus</option>
								<option value="Czechia (Czech Republic)">Czechia (Czech Republic)</option>
								<option value="Democratic Republic of the Congo (Congo-Kinshasa)">Democratic Republic of the Congo (Congo-Kinshasa)</option>
								<option value="Denmark">Denmark</option>
								<option value="Djibouti">Djibouti</option>
								<option value="Dominica">Dominica</option>
								<option value="Dominican Republic">Dominican Republic</option>
								<option value="Ecuador">Ecuador</option>
								<option value="Egypt">Egypt</option>
								<option value="El Salvador">El Salvador</option>
								<option value="Equatorial Guinea">Equatorial Guinea</option>
								<option value="Eritrea">Eritrea</option>
								<option value="Estonia">Estonia</option>
							</select>
						</div>
					</div>
				</div>
			</div>
		</div>
		<Footer/>
	</div>
</section>

<style>
    .profile_details_container {
        position              : relative;
        margin-top            : -40px;
        margin-left           : auto;
        margin-right          : auto;
        display               : grid;
        grid-template-columns : 70% auto;
        gap                   : 30px;
        width                 : 88%;
        color                 : #475569;
        font-size             : 14px;
    }

    .profile_details_container h2 {
        font-size   : 18px;
        margin      : 0 0 20px 0;
        font-weight : 700;
        text-align  : center;
        color       : #6366F1;
    }

    .profile_details_container .left .profile_details,
    .profile_details_container .right .password_set,
    .profile_details_container .right .country_details {
        display          : flex;
        flex-direction   : column;
        border-radius    : 8px;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
        padding          : 20px;
        background-color : white;
        height           : fit-content;
    }

    .profile_details_container .left,
    .profile_details_container .right {
        display        : flex;
        flex-direction : column;
        gap            : 30px;
    }

    .profile_details_container .left .profile_details .profile_pictures {
        display        : flex;
        flex-direction : row;
        align-items    : center;
        gap            : 25px;
        margin         : 0 0 20px 0;
    }

    .profile_details_container .left .profile_details .profile_pictures .img_container {
        width         : 60px;
        height        : 60px;
        overflow      : hidden;
        border        : 2px solid #86909F;
        border-radius : 50%;
    }

    .profile_details_container .left .profile_details .profile_pictures .img_container img {
        width      : 100%;
        height     : 100%;
        object-fit : cover;
    }

    .profile_details_container .left .profile_details .profile_pictures .actions {
        flex-grow      : 1;
        display        : flex;
        flex-direction : row;
        align-items    : center;
        gap            : 10px;
    }

    .profile_details_container .left .profile_details .profile_pictures .actions .btn_upload_photo {
        width            : fit-content;
        height           : fit-content;
        padding          : 7px 15px;
        border-radius    : 8px;
        border           : 2px solid #F1F5F9;
        background-color : #F1F5F9;
        color            : #3B82F6;
        font-weight      : 700;
        cursor           : pointer;
    }

    .profile_details_container .left .profile_details .profile_pictures .actions .btn_upload_photo:hover {
        text-decoration : underline;
    }

    .profile_details_container .left .profile_details .profile_pictures .actions .btn_delete {
        width            : fit-content;
        height           : fit-content;
        padding          : 6px 10px;
        border-radius    : 8px;
        background-color : #EF4444;
        border           : 1px solid #EF4444;
        cursor           : pointer;
    }

    .profile_details_container .left .profile_details .profile_pictures .actions .btn_delete:hover {
        background-color : #F85454;
        border           : 1px solid #F85454;
    }

    .profile_details_container .left .profile_details .input_container,
    .profile_details_container .right .password_set .input_container {
        display        : flex;
        flex-direction : column;
        gap            : 10px;
    }

    .profile_details_container .left .profile_details .input_container .name {
        display               : grid;
        grid-template-columns : 1fr 1fr;
        gap                   : 20px;
    }

    .profile_details_container .left .profile_details .name .profile_input {
        width : 100%;
    }

    .profile_details_container .left .profile_details .profile_input.email {
        width        : 49%;
        margin-right : 10px;
    }

    .profile_details_container .left .profile_details .info_container {
        display        : flex;
        flex-direction : column;
        gap            : 5px;
        margin-top     : 15px;
    }

    .profile_details_container .left .profile_details .info_container .profile_info {
        display     : inline-flex;
        align-items : center;
        font-size   : 13px;
        margin-left : 10px;
    }

    .profile_details_container .left .profile_details .info_container .profile_info label {
        font-weight  : 700;
        margin-right : 10px;
    }

    .profile_details_container .right .country_details .country_list #country {
        width            : 100%;
        border           : 1px solid #CBD5E1;
        background-color : #F1F5F9;
        border-radius    : 8px;
        padding          : 12px;
        cursor           : pointer;
    }

    .profile_details_container .right .country_details .country_list #country:focus {
        outline : 2px solid #3B82F6;
    }

    .profile_details #updateProfile,
    .password_set #changePassword {
        margin-left      : auto;
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

    .profile_details #updateProfile:hover,
    .password_set #changePassword:hover {
        background-color : #7E80F1;
    }

    /* Input box */
    .profile_input {
        display        : flex;
        flex-direction : column;
        gap            : 8px;
        width          : 100%;
    }

    .profile_input label {
        font-size   : 13px;
        font-weight : 700;
        margin-left : 10px;
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


</style>
