<script>
  //@ts-nocheck
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidAngleLeft from "svelte-icons-pack/fa/FaSolidAngleLeft";
  import {onMount} from "svelte";
  import {RealtorUpsertProperty} from "../jsApi.GEN";
  import FaSolidImage from "svelte-icons-pack/fa/FaSolidImage";
  import FaSolidPen from "svelte-icons-pack/fa/FaSolidPen";
  import FaSolidHome from "svelte-icons-pack/fa/FaSolidHome";
  import {formatPrice} from "./formatter";
  import FaSolidBed from "svelte-icons-pack/fa/FaSolidBed";
  import FaSolidBath from "svelte-icons-pack/fa/FaSolidBath";
  import FaSolidChair from "svelte-icons-pack/fa/FaSolidChair";
  import FaSolidBorderStyle from "svelte-icons-pack/fa/FaSolidBorderStyle";
  import FaSolidExchangeAlt from "svelte-icons-pack/fa/FaSolidExchangeAlt";
  import FaTrashAlt from "svelte-icons-pack/fa/FaTrashAlt";
  import FaSolidCamera from "svelte-icons-pack/fa/FaSolidCamera";
  import FaSolidTimes from "svelte-icons-pack/fa/FaSolidTimes";
  import FaSolidCircleNotch from "svelte-icons-pack/fa/FaSolidCircleNotch";
  
  export let user;
  export let property;
  export let countries;
  
  let approvalStatus = 'approved';
  let submitLoading = false;
  onMount( () => {
    console.log( 'Property = ', property )
    if( property.approvalState!=='pending' && property.approvalState!=='' ) {
      approvalStatus = 'rejected';
    }
    if( property.approvalState==='pending' ) {
      approvalStatus = 'pending'
    }
  } )
  
  function GetPayload() {
    if( property.countryCode==='US' ) property.city = property.countyName;
    return {
      countryCode: property.countryCode,
      city: property.city,
      countyName: property.countyName,
      district: property.district,
      street: property.street,
      street2: property.street2 || '',
      livingroom: property.livingroom,
      parking: parseFloat( property.parking ),
      depositFee: property.depositFee,
      minimumDurationYear: property.minimumDurationYear,
      otherFee: property.otherFee || [],
      imageLabels: property.imageLabels || [],
      altitude: property.altitude,
      id: '' + property.id,
      formattedAddress: property.formattedAddress,
      coord: [property.lat, property.lng],
      houseType: property.houseType,
      purpose: property.purpose,
      images: property.images || [],
      bedroom: property.bedroom,
      bathroom: property.bathroom,
      sizeM2: '' + property.sizeM2, // have to be string because of taiwan data
      mainUse: property.mainUse,
      note: property.note,
      lastPrice: '' + property.lastPrice,
      agencyFeePercent: property.agencyFeePercent,
      numberOfFloors: '' + property.floors,
    };
  }
  
  let approvalStates = {
    'pending': {
      description: 'Waiting for review 🔍',
      reason: 'We are reviewing your property. It will takes 1-3 days.'
    },
    'rejected': {
      description: 'Sorry, your property information has been rejected 😢',
      reason: property.approvalState
    },
    'approved': {
      description: 'Congratulations, your property has been successfully listed on the App. 😄',
      reason: ''
    },
  }
  
  let EDIT_PICTURE = 'picture', EDIT_FEATURE = 'feature', EDIT_FACILITY = 'facility', EDIT_ABOUT = 'about', EDIT_FLOORS = 'floors';
  let PART_TO_EDIT = '';
  
  // +================| Edit Picture |================+ //
  let imageHouseInput;
  let houseImgUploading = false;
  let uploadHouseStatus = '';
  let uploadHousePercent = 0;
  let images = property.images;
  let imageLabels = property.imageLabels || [];
  
  function handlerHouseImage() {
    if( !imageHouseInput ) return;
    const file = imageHouseInput.files[ 0 ];
    if( file ) {
      let formData = new FormData();
      formData.append( 'rawFile', file );
      formData.append( 'purpose', 'property' ); // property or floorPlan
      let ajax = new XMLHttpRequest();
      ajax.addEventListener( 'progress', function( event ) {
        houseImgUploading = true;
        let percent = (event.loaded / event.total) * 100;
        uploadHousePercent = Math.round( percent );
        uploadHouseStatus = `${uploadHousePercent}% uploaded... please wait`;
      } );
      ajax.addEventListener( 'load', function( event ) {
        houseImgUploading = false;
        if( ajax.status===200 ) {
          const out = JSON.parse( event.target.responseText );
          if( !out.error ) {
            images = [...images, out.urlPattern]; // push house image url to array
            imageLabels = [...imageLabels, ''];
          }
          alert( 'Image uploaded' );
        } else if( ajax.status===413 ) {
          alert( 'Image too large' );
        } else {
          alert( `Error: ${ajax.status}  ${ajax.statusText}` );
        }
      } );
      ajax.addEventListener( 'error', function( event ) {
        alert( 'Network error' );
      } );
      ajax.addEventListener( 'abort', function( event ) {
        alert( 'Upload aborted' );
      }, false );
      ajax.open( 'POST', '/user/uploadFile' );
      ajax.send( formData );
    }
    imageHouseInput.value = null;
  }
  
  function removeImage( index ) {
    images = images.filter( ( _, i ) => i!==index );
    imageLabels = imageLabels.filter( ( _, i ) => i!==index );
  }
  
  async function SaveEditPicture() {
    submitLoading = true;
    property.images = images;
    property.imageLabels = imageLabels;
    let payload = GetPayload();
    const prop = {property: payload};
    await RealtorUpsertProperty( prop, function( res ) {
      if( res.error ) {
        submitLoading = false;
        alert( res.error );
        return;
      }
      console.log( res );
      submitLoading = false;
    } );
  }
  
  // +================================================+ //
</script>

<div class="edit_property_root">
	<div class="upper_action">
		<a class='back_button' href='/realtor'>
			<Icon className='iconBack' color='#475569' size={18} src={FaSolidAngleLeft}/>
		</a>
		<span>Property ID: {property.id}</span>
	</div>
	<div class="edit_property_container">
		
		{#if PART_TO_EDIT===''}
			<div class="property_status">
				<div class={`status ${approvalStatus}`}>
					<p>{approvalStates[ approvalStatus ].description}</p>
				</div>
				{#if approvalStates[ approvalStatus ].reason!==''}
					<div class="reason">
						<p class:text_danger={approvalStatus === 'rejected'}>{approvalStates[ approvalStatus ].reason}</p>
						<button class="edit_btn">Review again</button>
					</div>
				{/if}
			</div>
			<div class="property_images_container">
				<div class='property_images'>
					{#if property.images && property.images.length}
						<img src={property.images[0]} alt=''/>
					{:else}
						<div class='image_empty'>
							<Icon size={40} color='#848d96' src={FaSolidImage}/>
							<span>No Image !</span>
						</div>
					{/if}
				</div>
				<button class="edit_btn" on:click={() => PART_TO_EDIT = EDIT_PICTURE}>
					<Icon color='#FFF' size={10} src={FaSolidPen}/>
					<span>Edit</span>
				</button>
			</div>
			<div class="main_details">
				<div class="col1">
					<div class="left">
						<div class="labels">
							<div class={property.purpose === 'rent' ? `purpose label_rent` : `purpose label_sale`}>
								{property.purpose==='rent' ? 'For Rent' : 'On Sale'}
							</div>
							<div class='house_type'>
								<Icon color='#FFF' size={13} src={FaSolidHome}/>
								<span>{property.houseType==='' ? 'House' : property.houseType}</span>
							</div>
						</div>
						<div class='price_details'>
							<h1>{formatPrice( property.lastPrice, 'TWD' ) || '0.00'}</h1>
							<p>Agency Fee : {property.agencyFeePercent}%</p>
						</div>
					</div>
					<div class="right">
						<button class="edit_btn">
							<Icon color='#FFF' size={10} src={FaSolidPen}/>
							<span>Edit</span>
						</button>
					</div>
				</div>
				<div class="col2">
					<div class='feature_item'>
						<b>{property.bedroom}</b>
						<div class="labels">
							<Icon className="labels_icon" color='#475569' size={13} src={FaSolidBed}/>
							<span>Beds</span>
						</div>
					</div>
					<div class='feature_item'>
						<b>{property.bathroom}</b>
						<div class="labels">
							<Icon className="labels_icon" color='#475569' size={13} src={FaSolidBath}/>
							<span>Baths</span>
						</div>
					</div>
					<div class='feature_item'>
						<b>{property.livingroom}</b>
						<div class="labels">
							<Icon className="labels_icon" color='#475569' size={12} src={FaSolidChair}/>
							<span>Livings</span>
						</div>
					</div>
					<div class='feature_item'>
						<b>{property.sizeM2}</b>
						<div class="labels">
							<Icon className="labels_icon" color='#475569' size={13} src={FaSolidBorderStyle}/>
							<span>M2</span>
							<button class='unit_toggle'>
								<span class="bg"></span>
								<Icon className="labels_icon" color='#F97316' size={12} src={FaSolidExchangeAlt}/>
							</button>
						</div>
					</div>
				</div>
			</div>
			<div class="second_details">
				<div class="facility">
					<div class="upper">
						<h3>Facility</h3>
						<button class="edit_btn">
							<Icon color='#FFF' size={10} src={FaSolidPen}/>
							<span>Edit</span>
						</button>
					</div>
					<p>
						Air conditioning/Heating, Balcony/Patio, Air conditioning/Heating, Balcony/Patio, Air conditioning/Heating, Balcony/Patio, Air
						conditioning/Heating, Balcony/Patio, Air conditioning/Heating, Balcony/Patio,
					</p>
				</div>
				<div class="about">
					<div class="upper">
						<h3>About</h3>
						<button class="edit_btn">
							<Icon color='#FFF' size={10} src={FaSolidPen}/>
							<span>Edit</span>
						</button>
					</div>
					<p>
						Looking for convenient city living? This spacious apartment is perfectly situated in the heart of the city, with nearby amenities including
						grocery stores, restaurants, and shopping. Enjoy plenty of natural light and modern conveniences like updated appliances and ample storage.
						This amazing apartment has everything you need for comfortable living in a fantastic location.
					</p>
				</div>
				<div class="parking">
					<div class="upper">
						<h3>Parking</h3>
						<button class="edit_btn">
							<Icon color='#FFF' size={10} src={FaSolidPen}/>
							<span>Edit</span>
						</button>
					</div>
					<div class="details">
						<span>Parking</span>
						<span>Yes</span>
					</div>
				</div>
			</div>
			<div class="floors">
				<div class="upper">
					<h1>Floors</h1>
					<button class="edit_btn">
						<Icon color='#FFF' size={10} src={FaSolidPen}/>
						<span>Edit</span>
					</button>
				</div>
				<div class="floor_lists">
					<div class="floor_item">
						<h3>Floor #1</h3>
						<div class="floor_details">
							<div class="room_lists">
								<div class="room">
									<span>Bedroom #1</span>
									<span>122.1 sq ft</span>
								</div>
								<div class="room">
									<span>Bedroom #1</span>
									<span>122.1 sq ft</span>
								</div>
								<div class="room">
									<span>Bedroom #1</span>
									<span>122.1 sq ft</span>
								</div>
							</div>
							<div class="floor_plan">
								<div class="img_container">
									<img alt="floor_plan" src="/assets/img/realtor/floor-plan-pen-ruler.webp"/>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		{/if}
		
		<!-- Edit partials -->
		{#if PART_TO_EDIT===EDIT_PICTURE}
			<div class="edit_part">
				<div class="upper">
					<button class='back_button' on:click={() => PART_TO_EDIT = ''}>
						<Icon className='iconBack' color='#475569' size={18} src={FaSolidAngleLeft}/>
					</button>
					<h3>Facility</h3>
				</div>
				<div class="edit_content">
					<div class="upload_picture">
						<div class='upload_container'>
							<label class='image_upload_button' for='upload_image'>
								{#if !houseImgUploading}
									<input
										bind:this={imageHouseInput}
										on:change={handlerHouseImage}
										type='file'
										accept='image/*'
										id='upload_image'
									/>
									<Icon size={35} className='upload_icon' color='#848D96' src={FaSolidCamera}/>
									<p>Select file to Upload</p>
								{:else}
									<progress value={uploadHousePercent} max='100'></progress>
									<p>{uploadHouseStatus}</p>
								{/if}
							</label>
						</div>
						<div class='image_lists'>
							{#if images && images.length}
								{#each images as img, idx}
									<div class='image_card'>
										<div class='image_container'>
											<img alt='' src={img}/>
										</div>
										<div class='image_description'>
											{#if imageLabels && imageLabels.length}
												<input placeholder='Description' type='text' bind:value={imageLabels[idx]}/>
											{:else}
												<input placeholder='Description' type='text'/>
											{/if}
										</div>
										<button
											class='remove_image'
											title='remove this image'
											on:click|preventDefault={() => removeImage(idx)}
										>
											<Icon color='#FFF' size={12} src={FaSolidTimes}/>
										</button>
									</div>
								{/each}
							{/if}
						</div>
					</div>
					<button class='next_button' on:click={SaveEditPicture}>
						{#if submitLoading===false}
							<span>Save</span>
						{/if}
						{#if submitLoading===true}
							<Icon className='spin' color='#FFF' size={15} src={FaSolidCircleNotch}/>
						{/if}
					</button>
				</div>
			</div>
		{/if}
	</div>
	<div class="delete_property_container">
		<button class="delete_property">
			<Icon color='#FFF' size={10} src={FaTrashAlt}/>
			<span>Delete Property</span>
		</button>
	</div>
</div>

<style>
    /* General purpose selector */
    :global(.back_button:hover .iconBack) {
        fill : #EF4444;
    }

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

    .edit_property_root {
        margin   : -40px 0 0 0;
        position : relative;
        z-index  : 50;
        padding  : 0;
    }

    .edit_property_root h1 {
        margin    : 0;
        font-size : 32px;
    }

    .edit_property_root h3 {
        margin    : 0;
        font-size : 20px;
    }

    .edit_property_root img {
        object-fit : cover;
        width      : 100%;
        height     : 100%;
    }

    .edit_property_root .text_danger {
        color : rgba(255, 126, 118, 1);
    }

    .edit_property_root .delete_property {
        background-color : #EF4444;
        display          : flex;
        flex-direction   : row;
        gap              : 6px;
        align-items      : center;
        color            : #FFF;
        padding          : 7px 15px;
        font-weight      : 600;
        border           : none;
        border-radius    : 8px;
        cursor           : pointer;
    }

    .edit_property_root .delete_property:hover {
        background-color : #F85454;
    }

    .edit_property_root .upper_action {
        display          : flex;
        flex-direction   : row;
        align-items      : center;
        gap              : 10px;
        font-size        : 20px;
        padding          : 0 55px 70px 55px;
        background-color : #EF4444;
        font-weight      : 600;
        color            : #FFF;
    }

    .edit_property_root .back_button {
        padding          : 8px;
        background-color : #FFF;
        border-radius    : 50%;
        font-size        : 14px;
        cursor           : pointer;
        border           : none;
        text-decoration  : none;
        height           : fit-content;
        color            : #334155;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    }

    .edit_property_root .edit_btn {
        display          : flex;
        flex-direction   : row;
        gap              : 6px;
        align-items      : center;
        color            : #FFF;
        background-color : #6366F1;
        padding          : 5px 15px;
        font-weight      : 600;
        border           : none;
        border-radius    : 8px;
        cursor           : pointer;
        width            : fit-content;
    }

    .edit_property_root .edit_btn:hover {
        background-color : #7E80F1;
    }

    .edit_property_root .delete_property_container {
        margin : 20px auto 0 auto;
        width  : 60%;
    }

    /* ========================== */

    .edit_property_container {
        position         : relative;
        width            : 60%;
        z-index          : 60;
        margin           : -50px auto 0 auto;
        color            : #475569;
        background-color : #FFF;
        display          : flex;
        flex-direction   : column;
        gap              : 20px;
        padding          : 20px 0;
        border-radius    : 10px;
        min-height       : 500px;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    }

    .edit_property_container .property_status {
        display        : flex;
        flex-direction : column;
        gap            : 8px;
        width          : 100%;
        height         : fit-content;
        padding        : 0 20px;
    }

    .edit_property_container .property_status .status {
        width         : 100%;
        height        : fit-content;
        padding       : 15px;
        border-radius : 8px;
        font-weight   : bold;
    }

    .edit_property_container .property_status p {
        margin : 0;
    }

    .edit_property_container .property_status .status.pending {
        background-color : rgba(255, 208, 118, 1);
    }

    .edit_property_container .property_status .status.rejected {
        background-color : rgba(255, 126, 118, 1);
        color            : #FFF;
    }

    .edit_property_container .property_status .status.approved {
        background-color : rgba(140, 216, 107, 1);
        color            : #FFF;
    }

    .edit_property_container .property_status .reason {
        display        : flex;
        flex-direction : column;
        gap            : 8px;
    }

    .edit_property_container .property_images_container {
        display         : flex;
        flex-direction  : row;
        justify-content : center;
        align-items     : center;
        height          : fit-content;
        width           : 100%;
        position        : relative;
    }

    .edit_property_container .property_images_container .property_images {
        width    : 340px;
        height   : 190px;
        flex     : none;
        overflow : hidden;
    }

    .edit_property_container .property_images_container .property_images .image_empty {
        object-fit       : cover;
        width            : 100%;
        height           : 100%;
        background-color : rgb(0 0 0 / 0.06);
        display          : flex;
        flex-direction   : column;
        gap              : 10px;
        font-size        : 18px;
        color            : #848D96;
        justify-content  : center;
        align-items      : center;
    }

    .edit_property_container .property_images_container .edit_btn {
        position : absolute !important;
        right    : 20px;
        top      : 15px;
    }

    .edit_property_container .main_details {
        display        : flex;
        flex-direction : column;
        gap            : 15px;
        padding        : 0 20px;
        position       : relative;
    }

    .edit_property_container .main_details .col1 {
        display        : flex;
        gap            : 15px;
        flex-direction : row;
    }

    .edit_property_container .main_details .col1 .left {
        flex-grow      : 1;
        display        : flex;
        flex-direction : column;
        gap            : 10px;
    }

    .edit_property_container .main_details .col1 .left .labels {
        width          : fit-content;
        display        : flex;
        flex-direction : row;
        gap            : 10px;
    }

    .edit_property_container .main_details .col1 .left .labels .purpose,
    .edit_property_container .main_details .col1 .left .labels .house_type {
        padding        : 5px 10px;
        width          : fit-content;
        display        : flex;
        flex-direction : row;
        gap            : 7px;
        align-items    : center;
        color          : #FFF;
        text-transform : capitalize;
        font-weight    : 600;
    }

    .edit_property_container .main_details .col1 .left .labels .house_type {
        background-color : #F97316;
    }

    .edit_property_container .main_details .col1 .left .price_details p {
        margin : 10px 0 0 0;
    }

    .edit_property_container .main_details .col1 .right {
        width : fit-content;
    }

    .edit_property_container .main_details .col2 {
        display         : flex;
        flex-direction  : row;
        justify-content : space-around;
        align-items     : center;
        text-align      : center;
    }

    .edit_property_container .main_details .col2 .feature_item {
        display        : flex;
        flex-direction : column;
        gap            : 8px;
    }

    .edit_property_container .main_details .col2 .feature_item b {
        font-size : 25px;
    }

    .edit_property_container .main_details .col2 .feature_item .labels {
        display         : flex;
        flex-direction  : row;
        justify-content : center;
        gap             : 6px;
    }

    .edit_property_container .main_details .col2 .feature_item .labels .unit_toggle {
        border     : none;
        background : transparent;
        position   : relative;
        cursor     : pointer;
    }

    .edit_property_container .main_details .col2 .feature_item .labels .unit_toggle .bg {
        width            : 0;
        height           : 0;
        border-radius    : 50%;
        background-color : rgb(0 0 0 / 0.06);
        z-index          : 1;
        position         : absolute;
        top              : -4px;
        left             : 0;
    }

    .edit_property_container .main_details .col2 .feature_item .labels .unit_toggle:hover .bg {
        width  : 24px;
        height : 24px;
    }

    .edit_property_container .second_details {
        display        : flex;
        flex-direction : column;
        gap            : 20px;
        padding        : 0 20px;
    }

    .edit_property_container .second_details div {
        display        : flex;
        flex-direction : column;
        gap            : 15px;
    }

    .edit_property_container .second_details div .upper,
    .edit_property_container .floors .upper {
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
        align-items     : center;
    }

    .edit_property_container .second_details div p {
        margin : 0;
    }

    .edit_property_container .second_details .parking .details {
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
        align-items     : center;
    }

    .edit_property_container .second_details .parking .details span:nth-child(2) {
        font-weight : 600;
    }

    .edit_property_container .floors {
        display        : flex;
        flex-direction : column;
        gap            : 20px;
        padding        : 0 20px;
    }

    .edit_property_container .floors .floor_lists {
        display        : flex;
        flex-direction : column;
        gap            : 15px;
    }

    .edit_property_container .floors .floor_lists .floor_item {
        display        : flex;
        flex-direction : column;
        gap            : 10px;
    }

    .edit_property_container .floors .floor_lists .floor_item .floor_details {
        display        : flex;
        flex-direction : row;
        gap            : 15px;
    }

    .edit_property_container .floors .floor_lists .floor_item .floor_details .room_lists {
        flex-grow      : 1;
        display        : flex;
        flex-direction : column;
        gap            : 8px;
    }

    .edit_property_container .floors .floor_lists .floor_item .floor_details .room_lists .room span:nth-child(1) {
        font-weight : 600;
    }

    .edit_property_container .floors .floor_lists .floor_item .floor_details .room_lists .room {
        display        : flex;
        flex-direction : row;
        align-items    : center;
        gap            : 20px;
    }

    .edit_property_container .floors .floor_lists .floor_item .floor_details .floor_plan {
        width : fit-content;
    }

    .edit_property_container .floors .floor_lists .floor_item .floor_details .floor_plan .img_container {
        width         : 340px;
        height        : 190px;
        flex          : none;
        overflow      : hidden;
        border-radius : 8px;
        border        : 1px solid #CBD5E1;
    }

    .edit_property_container .edit_part {
        width            : 100%;
        height           : 100%;
        background-color : #FFF;
        border-radius    : 8px;
        padding          : 0 20px;
        display          : flex;
        flex-direction   : column;
        gap              : 20px;
        flex-grow        : 1;
    }

    .edit_property_container .edit_part .upper {
        display        : flex;
        flex-direction : row;
        align-items    : center;
        gap            : 20px;
    }

    .edit_property_container .edit_part .upper .back_button {
        background-color : #F1F5F9;
    }

    .edit_property_container .edit_part .edit_content {
        display         : flex;
        flex-direction  : column;
        justify-content : space-between;
        flex-grow       : 1;
        gap             : 20px;
    }

    /* +=== UPDATE PICTURE ===+ */

    .edit_property_container .edit_part .upload_picture {
        display        : flex;
        flex-direction : column;
        gap            : 20px;
    }

    .edit_property_container .edit_part .upload_picture .upload_container {
        width  : 100%;
        height : fit-content;
    }

    .edit_property_container .edit_part .upload_picture .upload_container .image_upload_button {
        display          : flex;
        flex-direction   : column;
        justify-content  : center;
        align-items      : center;
        border           : 1px solid #CBD5E1;
        background-color : #F1F5F9;
        border-radius    : 8px;
        gap              : 8px;
        width            : 100%;
        color            : #848D96;
        height           : 110px;
        cursor           : pointer;
    }

    .edit_property_container .edit_part .upload_picture .upload_container .image_upload_button:hover {
        border : 1px solid #F97316;
        color  : #F97316;
    }

    :global(.edit_property_container .edit_part  .upload_picture .upload_container .image_upload_button:hover .upload_icon) {
        fill : #F97316;
    }

    .edit_property_container .edit_part .upload_picture .upload_container .image_upload_button input {
        position       : absolute;
        opacity        : 0;
        pointer-events : none;
    }

    .edit_property_container .edit_part .upload_picture .upload_container .image_upload_button p {
        font-size : 16px;
        margin    : 0;
    }

    .edit_property_container .edit_part .upload_picture .upload_container .image_upload_button progress {
        appearance    : none;
        border-radius : 8px;
        height        : 13px;
        overflow      : hidden;
        margin-bottom : 8px;
    }

    .edit_property_container .edit_part .upload_picture .upload_container .image_upload_button progress::-webkit-progress-bar {
        background-color   : aliceblue;
        box-shadow         : -1px 1px 10px 0 rgba(0, 0, 0, 0.3) inset;
        -webkit-box-shadow : -1px 1px 10px 0 rgba(0, 0, 0, 0.3) inset;
        -moz-box-shadow    : -1px 1px 10px 0 rgba(0, 0, 0, 0.3) inset;
    }

    .edit_property_container .edit_part .upload_picture .upload_container .image_upload_button progress::-webkit-progress-value {
        background-color : #F97316;
    }

    .edit_property_container .edit_part .upload_picture .upload_container .image_upload_button progress::-moz-progress-bar {
        background-color : #F97316;
    }

    .edit_property_container .edit_part .upload_picture .image_lists {
        display               : grid;
        gap                   : 20px;
        grid-auto-columns     : 1fr 1fr;
        grid-auto-rows        : 1fr 1fr;
        grid-template-columns : 1fr 1fr;
    }

    .edit_property_container .edit_part .upload_picture .image_lists .image_card {
        display          : flex;
        position         : relative;
        flex-direction   : column;
        padding          : 15px;
        background-color : #F1F5F9;
        border-radius    : 10px;
        gap              : 15px;
        border           : 1px solid #CBD5E1;
        height           : 270px;
    }

    .edit_property_container .edit_part .upload_picture .image_lists .image_card .image_container {
        width         : 100%;
        height        : 250px;
        overflow      : hidden;
        border-radius : 8px;
        border        : 1px solid #CBD5E1;
        cursor        : pointer;
    }

    .edit_property_container .edit_part .upload_picture .image_lists .image_card .image_container:hover img {
        transform : scale(1.20);
    }

    .edit_property_container .edit_part .upload_picture .image_lists .image_card .image_container img {
        object-fit          : cover;
        width               : 100%;
        height              : 100%;
        transition-duration : 75ms;
    }

    .edit_property_container .edit_part .upload_picture .image_lists .image_card .image_description input {
        width            : 100%;
        border           : 1px solid #CBD5E1;
        background-color : #F1F5F9;
        border-radius    : 8px;
        padding          : 12px;
    }

    .edit_property_container .edit_part .upload_picture .image_lists .image_card .image_description input:focus {
        border-color : #3B82F6;
        outline      : 1px solid #3B82F6;
    }

    .edit_property_container .edit_part .upload_picture .image_lists .image_card .remove_image {
        position         : absolute;
        right            : 10px;
        top              : 10px;
        padding          : 8px;
        border-radius    : 50%;
        border           : none;
        background-color : #EF4444;
        cursor           : pointer;
    }

    .edit_property_container .edit_part .upload_picture .image_lists .image_card .remove_image:hover {
        background-color : #F85454;
    }

    /* Responsive to mobile device */
    @media (max-width : 768px) {
        .edit_property_root {
            margin : -65px 20px 0 20px;
        }

        .edit_property_root .upper_action {
            padding   : 0 0 70px 0;
            font-size : 16px;
        }

        .edit_property_root .delete_property_container {
            margin : 20px 0;
            width  : 100%;
        }

        .edit_property_container {
            width   : 100%;
            padding : 15px 0;
        }

        .edit_property_container .property_status,
        .edit_property_container .main_details,
        .edit_property_container .second_details,
        .edit_property_container .floors {
            padding : 0 15px;
        }

        .edit_property_container .property_images_container {
            width    : 100%;
            position : relative;
        }

        .edit_property_container .property_images_container .property_images {
            width    : 100%;
            height   : 190px;
            flex     : none;
            overflow : hidden;
        }

        .edit_property_container .property_images_container .edit_btn {
            right : 15px;
        }

        .edit_property_container .main_details .col1 .right {
            position : absolute;
            top      : 0;
            right    : 15px;
        }

        .edit_property_container .main_details .col2 {
            justify-content : space-between;
        }

        .edit_property_container .floors .floor_lists .floor_item .floor_details .floor_plan .img_container {
            width  : 100px;
            height : 50px;
            flex   : none;
        }

        .edit_property_container .main_details .col2 .feature_item .labels {
            font-size : 10px;
            gap       : 3px;
        }

        :global(.labels_icon) {
            width  : 10px;
            height : 10px;
        }

        .edit_property_container .edit_part {
            padding : 0 15px;
        }

        .edit_property_container .edit_part .upload_picture {
            gap : 15px;
        }

        .edit_property_container .edit_part .upload_picture .image_lists {
            gap                   : 15px;
            grid-auto-columns     : 1fr;
            grid-auto-rows        : 1fr 1fr;
            grid-template-columns : 1fr;
        }

        .edit_property_container .edit_part .upload_picture .image_lists .image_card {
            padding : 10px;
        }

        .edit_property_container .edit_part .upload_picture .image_lists .image_card .remove_image {
            right : 5px;
            top   : 5px;
        }
    }
</style>