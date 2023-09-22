<script>
  import Growl from "../_components/Growl.svelte";
  import {formatPrice, localeDatetime} from '../_components/formatter.js';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidHome from "svelte-icons-pack/fa/FaSolidHome";
  import FaSolidMapMarkerAlt from "svelte-icons-pack/fa/FaSolidMapMarkerAlt";
  import FaSolidArrowRight from "svelte-icons-pack/fa/FaSolidArrowRight";
  import FaSolidShareAlt from "svelte-icons-pack/fa/FaSolidShareAlt";
  import FaCopy from "svelte-icons-pack/fa/FaCopy";
  import FaBrandsFacebook from "svelte-icons-pack/fa/FaBrandsFacebook";
  import FaBrandsLinkedin from "svelte-icons-pack/fa/FaBrandsLinkedin";
  import FaBrandsTwitter from "svelte-icons-pack/fa/FaBrandsTwitter";
  import PillBox from "../_components/PillBox.svelte";
  
  let propItem = {/* propItem */};
  let meta = {/* propertyMeta */}
  let showGrowl = false, gMsg = '', gType = '';
  
  function useGrowl( type, msg ) {
    showGrowl = true;
    gMsg = msg;
    gType = type;
    setTimeout( () => {
      showGrowl = false;
    }, 2000 );
  }
  
  function copyToClipboard( text ) {
    navigator.clipboard.writeText( text );
    useGrowl( 'success', 'Link copied to clipboard' );
  }
  
  function getBaseURL() {
    const url = `${window.location}`
    const regex = /^.+?[^\/:](?=[?\/]|$)/gm;
    const match = url.match( regex );
    const result = match;
    
    return result;
  }
</script>

<svelte:head>
	<meta content="HapSTR" property="og:site_name"/>
	<meta content="Property" property="og:title"/>
	<meta content={propItem.note} property="og:description"/>
	<meta content='tw_TW' property='og:locale'/>
	<meta content='en_EN' property='og:locale:alternate'/>
	<meta content="website" property="og:type"/>
	<meta content={`${window.location}`} property="og:url"/>
	<meta content={`${getBaseURL()}${propItem.images[0]}`} property="og:image"/>
	<meta content={`${getBaseURL()}${propItem.images[0]}`} property="og:image:secure_url"/>
	<meta content="1200" property="og:image:width"/>
	<meta content="630" property="og:image:height"/>
	<meta content="Property Image" property="og:image:alt"/>
	<meta content={propItem.createdAt} property="article:published_time"/>
	<meta content={propItem.updatedAt} property="article:modified_time"/>
	<meta content={propItem.updatedAt} property="article:updated_time"/>
	
	<meta content="summary_large_image" name="twitter:card"/>
	<meta content="I found an awesome house" name="twitter:title"/>
	<meta content={propItem.note} name="twitter:description"/>
	<meta content={`${getBaseURL()}${propItem.images[0]}`} name="twitter:image"/>
	<!--	<meta name="twitter:site" content="@xd"/>-->
	<!--	<meta name="twitter:creator" content="@xd"/>-->
</svelte:head>

{#if showGrowl}
	<Growl message={gMsg} growlType={gType}/>
{/if}
<section class="property_container">
	<div class='property'>
		<div class='property_main'>
			<div class='property_images'>
				{#if propItem.images && propItem.images.length}
					<img src={propItem.images[0]} alt=''/>
				{:else}
					<div class='image_empty'>
						<i class='gg-image'/>
					</div>
				{/if}
			</div>
			<div class='property_info'>
				<div class='col1'>
					<div class='left'>
						<div class={propItem.purpose === 'rent' ? `purpose label_rent` : `purpose label_sale`}>
							{propItem.purpose==='rent' ? 'For Rent' : 'On Sale'}
						</div>
						<div class='house_type'>
							<Icon color='#FFFF' size={16} src={FaSolidHome}/>
							<span>{propItem.houseType}</span>
						</div>
					</div>
				</div>
				<div class='col2'>
					<h1>{formatPrice( propItem.lastPrice, 'TWD' ) || '0.00'}</h1>
					<p>Agency Fee : {propItem.agencyFeePercent}%</p>
					<div class='address'>
						<Icon className='icon_address' color='#f97316' size={18} src={FaSolidMapMarkerAlt}/>
						<span>{propItem.formattedAddress}</span>
					</div>
				</div>
			</div>
		</div>
		<div class='property_secondary'>
			<div class='feature_number'>
				<div class='feature_item'>
					<b>{propItem.numberOfFloors || '0'}</b>
					<p>Floors</p>
				</div>
				<div class='feature_item'>
					<b>{propItem.bathroom || '0'}</b>
					<p>Baths</p>
				</div>
				<div class='feature_item'>
					<b>{propItem.bedroom || '0'}</b>
					<p>Beds</p>
				</div>
				<div class='feature_item'>
					<b>{propItem.sizeM2 || '0'} M2</b>
					<p>Size</p>
				</div>
			</div>
		</div>
		<div class='property_attributes'>
			{#each meta as m}
				{#if propItem[ m.name ]}
					{#if m.inputType==='datetime'}
						<PillBox label={m.label} content={localeDatetime(propItem[m.name])}/>
					{:else}
						<PillBox label={m.label} content={propItem[m.name]}/>
					{/if}
				{/if}
			{/each}
		</div>
		<div class='property_floors'>
			<h3>Floors</h3>
			{#if propItem.floorList && propItem.floorList.length}
				<div class='floor_lists'>
					{#each propItem.floorList as floors}
						<div class='floor_item'>
							<div class='left'>
								<h5>
									{floors.type==='basement' ? 'Basement' : `Floor #${floors.floor}`}
								</h5>
								<!-- TODO: currently room list only 1 object, fix Tarantool to accept array -->
								{#if floors.rooms}
									<div class='room_lists'>
										{#each floors.rooms as room}
											<div class='room_item'>
												<span>{room.name || '-'}</span>
												<span>{room.sizeM2 || '-'} M2</span>
											</div>
										{/each}
									</div>
								{/if}
							</div>
							<div class='floor_plan_image'>
								{#if floors.planImageUrl===''}
                        <span>
                           <i class='gg-image'/>
                           <p>No Image</p>
                        </span>
								{:else}
									<img src={floors.planImageUrl} alt=''/>
								{/if}
							</div>
						</div>
					{/each}
				</div>
			{:else}
				<div class='no_floors'>
					<p>No Floors</p>
				</div>
			{/if}
		</div>
	</div>
	<div class='side_attribute'>
		<div class='login_container'>
			<div>
				<h3>Login to</h3>
				<h3>HapSTR</h3>
			</div>
			<a class='login_btn' href='/'>
				<span>Login</span>
				<Icon color='#FFFF' size={13} src={FaSolidArrowRight}/>
			</a>
		</div>
		<div class='share_container'>
			<header>
				<span>Share this</span>
				<Icon className='share_icon' color='#9fa9b5' size={14} src={FaSolidShareAlt}/>
			</header>
			<div class='share_options'>
				<button class='share_item' on:click={() => copyToClipboard(`${window.location}`)} title='Copy link address'>
					<Icon className='share_icon' color='#475569' size={25} src={FaCopy}/>
				</button>
				<a aria-label="Share to Facebook"
				   class='share_item'
				   href={`https://www.facebook.com/sharer/sharer.php?u=${window.location}?utm_source=facebook&utm_medium=social&utm_campaign=user-share`}
				   rel="noopener"
				   target="_blank"
				>
					<Icon className='share_icon' color='#475569' size={25} src={FaBrandsFacebook}/>
				</a>
				<a aria-label="Share to LinkedIn"
				   class='share_item'
				   href={`http://www.linkedin.com/shareArticle?mini=true&url=${window.location}&title=I%20Found%20Awesome%House%20${window.location}property/${propItem.id}`}
				   rel="noopener"
				   target="_blank"
				>
					<Icon className='share_icon' color='#475569' size={25} src={FaBrandsLinkedin}/>
				</a>
				<a aria-label="Share to Twitter"
				   class='share_item'
				   href={`https://twitter.com/intent/tweet?text=I%20Found%20Awesome%House ${window.location}property/${propItem.id}`}
				   rel="noopener"
				   target="_blank"
				>
					<Icon className='share_icon' color='#475569' size={25} src={FaBrandsTwitter}/>
				</a>
			</div>
		</div>
	</div>
</section>

<style>
    .property_container {
        width           : 100%;
        margin          : 30px auto 50px auto;
        color           : #475569;
        display         : flex;
        flex-direction  : row;
        justify-content : center;
        gap             : 30px;
    }

    .property {
        width            : 60%;
        height           : fit-content;
        background-color : #F0F0F0;
        border-radius    : 8px;
        padding          : 15px 15px 70px 15px;
        display          : flex;
        flex-direction   : column;
        gap              : 20px;
        position         : relative;
    }

    .property_main {
        display         : flex;
        flex-direction  : row;
        flex-wrap       : nowrap;
        align-content   : flex-start;
        justify-content : flex-start;
        align-items     : flex-start;
        gap             : 20px;
    }

    .property_main .property_images {
        width         : 340px;
        height        : 190px;
        flex          : none;
        overflow      : hidden;
        border-radius : 8px;
    }

    .property_main .property_images img {
        object-fit : cover;
        width      : 100%;
        height     : 100%;
    }

    .property_main .property_images .image_empty {
        border-radius    : 8px;
        object-fit       : cover;
        width            : 100%;
        height           : 100%;
        background-color : rgb(0 0 0 / 0.06);
        display          : flex;
        justify-content  : center;
        align-items      : center;
    }

    .property_main .property_info {
        flex-grow      : 1;
        display        : flex;
        flex-direction : column;
        gap            : 10px;
        margin-top     : 20px;
    }

    .property_main .property_info .col1 {
        display         : flex;
        justify-content : space-between;
        align-items     : center;
        flex-direction  : row;
    }

    .property_main .property_info .col1 .left {
        display     : flex;
        gap         : 15px;
        align-items : center;
    }

    .property_main .property_info .col1 .left .purpose,
    .property_main .property_info .col1 .left .house_type,
    .property_main .property_info .col1 .edit_property {
        background-color : #F97316;
        padding          : 7px 18px;
        border-radius    : 8px;
        color            : white;
        font-size        : 14px;
        text-transform   : capitalize;
        text-decoration  : none;
    }

    .property_main .property_info .col1 .left .house_type,
    .property_main .property_info .col1 .edit_property {
        display     : flex;
        gap         : 10px;
        align-items : center;
    }

    .property_main .property_info .col1 .edit_property:hover {
        background-color : #F58433;
    }

    .property_main .property_info .col2 {
        display        : flex;
        flex-direction : column;
        gap            : 10px;
    }

    .property_main .property_info .col2 h1 {
        margin    : 0;
        font-size : 32px;
    }

    .property_main .property_info .col2 p {
        margin : 0 0 10px 0;
    }

    .property_main .property_info .col2 .address {
        display        : flex;
        flex-direction : row;
        gap            : 10px;
        align-items    : center;
    }

    :global(.icon_address) {
        flex-shrink : 0;
    }

    .property_main .property_info .col2 .address span {
        flex-shrink : 1;
    }

    .property_secondary {
        display        : flex;
        flex-direction : column;
        gap            : 15px;
    }

    .property_secondary .feature_number {
        display         : flex;
        flex-direction  : row;
        align-items     : center;
        justify-content : space-evenly;
        width           : 100%;
    }

    .property_secondary .feature_number .feature_item {
        display        : flex;
        flex-direction : column;
        align-items    : center;
        gap            : 8px;
        text-transform : capitalize;
    }

    .property_secondary p {
        margin : 0;
    }

    .property_secondary b {
        font-size : 22px;
    }

    .property_floors {
        width  : 70%;
        margin : 0 auto;
    }

    .property_floors h3 {
        font-size  : 22px;
        margin     : 0 0 15px 0;
        text-align : center;
    }

    .property_floors .floor_lists {
        display        : flex;
        flex-direction : column;
        gap            : 20px;
    }

    .property_floors .floor_lists .floor_item {
        display     : flex;
        align-items : flex-start;
        gap         : 35px;
    }

    .property_floors .floor_lists .floor_item .left {
        flex-grow : 1;
    }

    .property_floors .floor_lists .floor_item .left h5 {
        font-size      : 22px;
        text-transform : capitalize;
        font-weight    : normal;
        margin         : 0 0 15px 0;
        width          : 100%;
    }

    .property_floors .floor_lists .floor_lists .floor_item .left .room_lists {
        font-size : 14px;
        width     : 100%;
        color     : #475569;
        border    : 1px solid #CBD5E1;
    }

    .property_floors .floor_lists .floor_item .left .room_lists .room_item {
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
        text-transform  : capitalize;
        border-bottom   : 1px solid #CBD5E1;
        padding         : 5px 0;
    }

    .property_floors .floor_lists .floor_item .floor_plan_image {
        width         : 240px;
        height        : 140px;
        position      : relative;
        border-radius : 8px;
        overflow      : hidden;
        border        : 1px solid #CBD5E1;
        margin-top    : 45px;
    }

    .property_floors .floor_lists .floor_item .floor_plan_image img {
        width      : 100%;
        height     : 100%;
        object-fit : cover;
    }

    .property_floors .floor_lists .floor_item .floor_plan_image span {
        border-radius   : 8px;
        object-fit      : cover;
        width           : 100%;
        height          : 100%;
        display         : flex;
        flex-direction  : column;
        gap             : 8px;
        justify-content : center;
        align-items     : center;
    }

    .property_floors .floor_lists .floor_item .floor_plan_image span p {
        margin : 0;
    }

    .side_attribute {
        width          : 250px;
        height         : fit-content;
        display        : flex;
        flex-direction : column;
        gap            : 20px;
    }

    .side_attribute .login_container,
    .side_attribute .share_container {
        text-align       : center;
        height           : fit-content;
        background-color : #F0F0F0;
        border-radius    : 8px;
        padding          : 15px;
        display          : flex;
        flex-direction   : column;
        gap              : 10px;
    }

    .side_attribute .login_container h3 {
        margin    : 0 0 10px 0;
        font-size : 24px;
    }

    .side_attribute .login_container .login_btn {
        background-color : #6366F1;
        width            : 100%;
        height           : fit-content;
        padding          : 12px;
        text-align       : center;
        border-radius    : 8px;
        color            : #FFF;
        font-size        : 13px;
        display          : flex;
        flex-direction   : row;
        gap              : 6px;
        justify-content  : center;
        align-items      : center;
        text-decoration  : none;
    }

    .side_attribute .login_container .login_btn:hover {
        background-color : #7E80F1;
    }

    .side_attribute .share_container header {
        display         : flex;
        flex-direction  : row;
        justify-content : center;
        gap             : 7px;
        align-items     : center;
        font-size       : 15px;
    }

    .side_attribute .share_container .share_options {
        display               : grid;
        grid-template-columns : repeat(4, minmax(0, 1fr));
        align-items           : center;
        justify-items         : center;
        justify-content       : center;
        align-content         : center;
    }

    .side_attribute .share_container .share_options .share_item {
        border     : none;
        background : none;
        cursor     : pointer;
    }

    :global(.side_attribute .share_container .share_options .share_item:hover .share_icon) {
        fill : #57667A !important;
    }
</style>
