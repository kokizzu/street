<script>
  // @ts-nocheck
  import Menu from '../_components/Menu.svelte';
  import ProfileHeader from '../_components/ProfileHeader.svelte';
  import Footer from '../_components/Footer.svelte';
  import Property from '../_components/Property.svelte'
  import {onMount} from "svelte";
  
  let propItem = {/* property */};
  let meta = {/* propertyMeta */}
  let user = {/* user */};
  let segments = {/* segments */};
  
  let isAdmin = false;
  
  let approvalStatus = 'approved';
  onMount(() => {
    console.log('onMount.realtor/ownedProperty')
    console.log('Property = ', propItem)
	  if (propItem.approvalState !== 'pending' && propItem.approvalState !== '') {
      approvalStatus = 'rejected';
	  }
    if (propItem.approvalState === 'pending') {
      approvalStatus = 'pending'
    }
  })
  let approvalStates = {
    'pending': {
      description: 'Waiting for review 🔍',
      reason: 'We are reviewing your property. It will takes 1-3 days.'
    },
    'rejected': {
      description: 'Sorry, your property information has been rejected 😢',
      reason: propItem.approvalState
    },
    'approved': {
      description: 'Congratulations, your property has been successfully listed on the App. 😄',
      reason: ''
    },
  }
</script>

<section class='dashboard'>
	<Menu
		access={segments}
	/>
	<div class='dashboard_main_content'>
		<ProfileHeader {user} access={segments}/>
		<div class='content'>
			<div class="property_container">
				<div class="property_status">
					<div class={`status ${approvalStatus}`}>
						<p>{approvalStates[approvalStatus].description}</p>
					</div>
					<div class="reason">
						<p>{approvalStates[approvalStatus].reason}</p>
					</div>
				</div>
				<div class='property'>
					<Property {meta} {propItem} {isAdmin}/>
				</div>
			</div>
		</div>
		<Footer></Footer>
	</div>
</section>

<style>
    .property_container {
        position        : relative;
        margin          : -40px auto 50px auto;
        justify-content : center;
        color           : #475569;
        width           : 70%;
        display         : flex;
        flex-direction  : column;
        gap             : 20px;
    }

    .property_status {
        display          : flex;
        flex-direction   : column;
        gap              : 8px;
        margin-top       : 15px;
        padding          : 20px;
        border-radius    : 8px;
        background-color : #F0F0F0;
        border           : 1px solid #D9DDE3;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    }

    .property_status .status {
        width         : 100%;
        height        : fit-content;
        text-align    : center;
        padding       : 15px 0;
        border-radius : 8px;
        font-weight   : bold;
    }

    .property_status p {
        margin : 0;
    }

    .property_status .status.pending {
        background-color : rgba(255, 208, 118, 1);
    }

    .property_status .status.rejected {
        background-color : rgba(255, 126, 118, 1);
		    color: #FFF;
    }

    .property_status .status.approved {
        background-color : rgba(140, 216, 107, 1);
        color: #FFF;
    }

    .property {
        filter        : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
        border        : 1px solid #D9DDE3;
        border-radius : 8px;
    }
</style>