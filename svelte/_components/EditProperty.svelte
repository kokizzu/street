<script>
  //@ts-nocheck
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidAngleLeft from 'svelte-icons-pack/fa/FaSolidAngleLeft';
  import { onMount } from 'svelte';
  import { AdminProperties, RealtorUpsertProperty } from '../jsApi.GEN';
  import FaSolidImage from 'svelte-icons-pack/fa/FaSolidImage';
  import FaSolidPen from 'svelte-icons-pack/fa/FaSolidPen';
  import FaSolidHome from 'svelte-icons-pack/fa/FaSolidHome';
  import { formatPrice, getApprovalState } from './formatter';
  import FaSolidBed from 'svelte-icons-pack/fa/FaSolidBed';
  import FaSolidBath from 'svelte-icons-pack/fa/FaSolidBath';
  import FaSolidChair from 'svelte-icons-pack/fa/FaSolidChair';
  import FaSolidBorderStyle from 'svelte-icons-pack/fa/FaSolidBorderStyle';
  import FaSolidExchangeAlt from 'svelte-icons-pack/fa/FaSolidExchangeAlt';
  import FaTrashAlt from 'svelte-icons-pack/fa/FaTrashAlt';
  import FaSolidCamera from 'svelte-icons-pack/fa/FaSolidCamera';
  import FaSolidTimes from 'svelte-icons-pack/fa/FaSolidTimes';
  import FaSolidCircleNotch from 'svelte-icons-pack/fa/FaSolidCircleNotch';
  import FaCheckCircle from 'svelte-icons-pack/fa/FaCheckCircle';
  import AddOtherFeesDialog from './AddOtherFeesDialog.svelte';
  
  export let isOwner = false;
  export let property;
  export let countries;
  export let isAdmin = false;
  
  let approvalStatus = 'approved';
  let submitLoading = false;
  let countryCurrency = 'TWD';

  onMount( () => {
    console.log( 'Property = ', property );
    approvalStatus = getApprovalState( property.approvalState );
    for( let i = 0; i<countries.length; i++ ) {
      if( countries[ i ].iso_2===property.countryCode ) {
        countryCurrency = countries[ i ].currency.code;
      }
    }
    console.log( 'Approval status = ', approvalStatus );
  } );
  
  function GetPayload() {
    if( property.countryCode==='US' ) property.city = property.countyName;
    if( property.numberOfFloors==='undefined' ) property.numberOfFloors = 1;
    return {
      countryCode: property.countryCode,
      city: property.city,
      countyName: property.countyName,
      district: property.district,
      street: property.street + "\n" + (property.street2 || ''),
      livingroom: property.livingroom,
      parking: parseFloat( property.parking ),
      depositFee: property.depositFee,
      minimumDurationYear: property.minimumDurationYear,
      otherFees: property.otherFees || [],
      imageLabels: property.imageLabels || [],
      altitude: property.altitude,
      id: property.id,
      formattedAddress: property.formattedAddress,
      coord: property.coord,
      houseType: property.houseType,
      purpose: property.purpose,
      images: property.images || [],
      bedroom: property.bedroom,
      bathroom: property.bathroom,
      sizeM2: '' + property.sizeM2,
      mainUse: property.mainUse,
      note: property.note,
      lastPrice: '' + property.lastPrice,
      agencyFeePercent: property.agencyFeePercent || 0,
    };
  }
  
  let approvalStates = {
    'pending': {
      description: 'Waiting for review 🔍',
      reason: 'We are reviewing your property. It will takes 1-3 days.',
    },
    'rejected': {
      description: 'Sorry, your property information has been rejected 😢',
      reason: property.approvalState,
    },
    'approved': {
      description: 'Congratulations, your property has been successfully listed on the App. 😄',
      reason: '',
    },
  };
  
  const EDIT_PICTURE = 'picture', EDIT_FEATURE = 'feature', EDIT_FACILITY = 'facility', EDIT_ABOUT = 'about', EDIT_FLOORS = 'floors';
  let PART_TO_EDIT = '';
  
  // +================| Edit Picture |================+ //
  let imageHouseInput;
  let houseImgUploading = false, uploadHouseStatus = '', uploadHousePercent = 0;
  let images = property.images, imageLabels = property.imageLabels || [];
  
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
    const payload = GetPayload();
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
  
  // +================| Edit Feature |===================+ //
  let houseTypeLists = [
    'house', 'land', 'apartment', 'townhouse', 'condo', 'villa', 'factory', 'parking', 'other',
  ], editFeatureCount = 0;
  let houseType = property.houseType, bedroom = property.bedroom, bathroom = property.bathroom, livingroom = property.livingroom;
  let sizeM2 = property.sizeM2, parking = property.parking.toString();
  let purpose = property.purpose, lastPrice = property.lastPrice, agencyFeePercent = property.agencyFeePercent;
  let depositFee = property.depositFee, minimumDurationYear = property.minimumDurationYear, otherFees = property.otherFees;
  let otherFeeObj = {name: '', fee: 0};
  let addOtherFeeDialog = AddOtherFeesDialog, agencyFee = 'true';
  let submitApprove = false, submitReject = false;
  
  function addOtherFee() {
    otherFees = [...otherFees, otherFeeObj];
    otherFeeObj = {
      name: '',
      fee: 0,
    };
    addOtherFeeDialog.hideModal();
  }
  
  async function SaveEditFeature() {
    if( editFeatureCount===0 ) {
      editFeatureCount = 1;
      console.log( parking );
      return;
    }
    if( editFeatureCount===1 ) {
      submitLoading = true;
      property.houseType = houseType;
      property.bedroom = bedroom;
      property.bathroom = bathroom;
      property.livingroom = livingroom;
      property.sizeM2 = sizeM2;
      property.parking = parking;
      property.purpose = purpose;
      property.lastPrice = lastPrice;
      property.agencyFeePercent = agencyFeePercent;
      property.depositFee = depositFee;
      property.minimumDurationYear = minimumDurationYear;
      property.otherFees = otherFees;
      const payload = GetPayload();
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
  }
  
  // +================| Edit Facility |===================+ //
  let mainUse = property.mainUse;
  
  async function SaveEditFacility() {
    submitLoading = true;
    property.mainUse = mainUse;
    const payload = GetPayload();
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
  
  // +================| Edit About |===================+ //
  let note = property.note;
  
  async function SaveEditAbout() {
    submitLoading = true;
    property.note = note;
    const payload = GetPayload();
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
  
  function ApproveProperty() {
    submitApprove = true;
    AdminProperties( {
        cmd: 'upsert',
        property: {id: property.id, approvalState: ' '}, // empty is approved, will be trimmed on server side
      },
      function( res ) {
        if( res.error ) {
          submitApprove = false;
          alert( res.error );
          return;
        }
        submitApprove = false;
      } );
  }
  
  let rejectReason = '', showRejectDialog = false;
  
  function RejectProperty() {
    if( rejectReason==='' ) return alert( 'Reason cannot be empty' );
    AdminProperties( {
        cmd: 'upsert',
        property: {id: property.id, approvalState: rejectReason},
      },
      function( res ) {
        if( res.error ) {
          showRejectDialog = false;
          alert( res.error );
          return;
        }
        showRejectDialog = false;
        console.log( res );
      } );
  }

  function ReviewProperty() {
    RealtorUpsertProperty( {
      askReview: true,
      property: property
    },
    function(res) {
      if(res.error) {
        console.log(res);
        return
      }
      console.log(res)
    }
    )
  }
</script>

<div class='edit_property_root'>
  <div class='upper_action'>
    <a class='back_button' href='/realtor'>
      <Icon className='iconBack' color='#475569' size={18} src={FaSolidAngleLeft} />
    </a>
    <span>Property ID: {property.id}</span>
  </div>
  <div class='edit_property_container'>
    {#if showRejectDialog===true}
      <div class='backdrop_reject_reason'>
        <div class='reject_reason'>
          <div class='reject_reason_content'>
            <h3>Reject Reason</h3>
            <textarea
              bind:value={rejectReason}
              name='reject-reason'
              id='reject-reason'
              placeholder={'input refusal reason for #' + property.id + ', use "pending" to set state to pending'}
              class='textarea_input'
            ></textarea>
          </div>
          <div class='buttons'>
            <button class='cancel_btn' on:click={()=>showRejectDialog=false}>Cancel</button>
            <button class='reject_btn' on:click={RejectProperty}>Reject</button>
          </div>
        </div>
      </div>
    {/if}
    
    {#if PART_TO_EDIT===''}
      <div class='property_status'>
        <div class={`status ${approvalStatus}`}>
          <p>{approvalStates[ approvalStatus ].description}</p>
        </div>
        {#if approvalStates[ approvalStatus ].reason!==''}
          <div class='reason'>
            <p class:text_danger={approvalStatus === 'rejected'}>{approvalStates[ approvalStatus ].reason}</p>
          </div>
        {/if}
        {#if isOwner && approvalStatus!==''}
            <button class='edit_btn' on:click={ReviewProperty}>
              Review again
            </button>
        {/if}
        {#if isAdmin}
          <div class='action_btns'>
            {#if approvalStatus==='pending' || approvalStatus==='rejected'}
              <button class='approve_btn' on:click={ApproveProperty}>
                {#if !submitApprove}
                  <Icon size={10} color='#FFF' src={FaCheckCircle} />
                {/if}
                {#if submitApprove}
                  <Icon className='spin' color='#FFF' size={10} src={FaSolidCircleNotch} />
                {/if}
                <span>Approve</span>
              </button>
            {/if}
            {#if approvalStatus==='pending' || approvalStatus==='approved'}
              <button class='reject_btn' on:click={()=>showRejectDialog=true}>
                <Icon size={10} color='#FFF' src={FaSolidTimes} />
                <span>Reject</span>
              </button>
            {/if}
          </div>
        {/if}
      </div>
      <div class='property_images_container'>
        <div class='property_images'>
          {#if property.images && property.images.length}
            <img src={property.images[0]} alt='' />
          {:else}
            <div class='image_empty'>
              <Icon size={40} color='#848d96' src={FaSolidImage} />
              <span>No Image !</span>
            </div>
          {/if}
        </div>
        {#if isOwner}
        <button class='edit_btn' on:click={() => PART_TO_EDIT = EDIT_PICTURE}>
          <Icon color='#FFF' size={10} src={FaSolidPen} />
          <span>Edit</span>
        </button>
        {/if}
      </div>
      <div class='main_details'>
        <div class='col1'>
          <div class='left'>
            <div class='labels'>
              <div class={property.purpose === 'rent' ? `purpose label_rent` : `purpose label_sale`}>
                {property.purpose==='rent' ? 'For Rent' : 'On Sale'}
              </div>
              <div class='house_type'>
                <Icon color='#FFF' size={13} src={FaSolidHome} />
                <span>{property.houseType==='' ? 'House' : property.houseType}</span>
              </div>
            </div>
            <div class='price_details'>
              <h1>{formatPrice( property.lastPrice, 'TWD' ) || '0.00'}</h1>
              <p>Agency Fee : {property.agencyFeePercent}%</p>
            </div>
          </div>
          <div class='right'>
            {#if isOwner}
            <button class='edit_btn' on:click={() => PART_TO_EDIT = EDIT_FEATURE}>
              <Icon color='#FFF' size={10} src={FaSolidPen} />
              <span>Edit</span>
            </button>
            {/if}
          </div>
        </div>
        <div class='col2'>
          <div class='feature_item'>
            <b>{property.bedroom}</b>
            <div class='labels'>
              <Icon className='labels_icon' color='#848D96' size={13} src={FaSolidBed} />
              <span>Beds</span>
            </div>
          </div>
          <div class='feature_item'>
            <b>{property.bathroom}</b>
            <div class='labels'>
              <Icon className='labels_icon' color='#848D96' size={13} src={FaSolidBath} />
              <span>Baths</span>
            </div>
          </div>
          <div class='feature_item'>
            <b>{property.livingroom}</b>
            <div class='labels'>
              <Icon className='labels_icon' color='#848D96' size={12} src={FaSolidChair} />
              <span>Livings</span>
            </div>
          </div>
          <div class='feature_item'>
            <b>{property.sizeM2}</b>
            <div class='labels'>
              <Icon className='labels_icon' color='#848D96' size={13} src={FaSolidBorderStyle} />
              <span>M2</span>
              <button class='unit_toggle'>
                <span class='bg'></span>
                <Icon className='labels_icon' color='#F97316' size={12} src={FaSolidExchangeAlt} />
              </button>
            </div>
          </div>
        </div>
      </div>
      <div class='second_details'>
        <div class='facility'>
          <div class='upper'>
            <h3>Facility</h3>
            {#if isOwner}
            <button class='edit_btn' on:click={() => PART_TO_EDIT = EDIT_FACILITY}>
              <Icon color='#FFF' size={10} src={FaSolidPen} />
              <span>Edit</span>
            </button>
            {/if}
          </div>
          <p>
            {property.mainUse || '--'}
          </p>
        </div>
        <div class='about'>
          <div class='upper'>
            <h3>About</h3>
            {#if isOwner}
            <button class='edit_btn' on:click={() => PART_TO_EDIT = EDIT_ABOUT}>
              <Icon color='#FFF' size={10} src={FaSolidPen} />
              <span>Edit</span>
            </button>
            {/if}
          </div>
          <p>
            {property.note || '--'}
          </p>
        </div>
        <div class='parking'>
          <div class='upper'>
            <h3>Parking</h3>
            {#if isOwner}
            <button class='edit_btn'>
              <Icon color='#FFF' size={10} src={FaSolidPen} />
              <span>Edit</span>
            </button>
            {/if}
          </div>
          <div class='details'>
            <span>Parking</span>
            <span>{property.parking>=1 ? 'Yes' : 'No'}</span>
          </div>
        </div>
      </div>
<!--      <div class='floors'>-->
<!--        <div class='upper'>-->
<!--          <h1>Floors</h1>-->
<!--          <button class='edit_btn'>-->
<!--            <Icon color='#FFF' size={10} src={FaSolidPen} />-->
<!--            <span>Edit</span>-->
<!--          </button>-->
<!--        </div>-->
<!--        <div class='floor_lists'>-->
<!--          <div class='floor_item'>-->
<!--            <h3>Floor #1</h3>-->
<!--            <div class='floor_details'>-->
<!--              <div class='room_lists'>-->
<!--                <div class='room'>-->
<!--                  <span>Bedroom #1</span>-->
<!--                  <span>122.1 sq ft</span>-->
<!--                </div>-->
<!--                <div class='room'>-->
<!--                  <span>Bedroom #1</span>-->
<!--                  <span>122.1 sq ft</span>-->
<!--                </div>-->
<!--                <div class='room'>-->
<!--                  <span>Bedroom #1</span>-->
<!--                  <span>122.1 sq ft</span>-->
<!--                </div>-->
<!--              </div>-->
<!--              <div class='floor_plan'>-->
<!--                <div class='img_container'>-->
<!--                  <img alt='floor_plan' src='/assets/img/realtor/floor-plan-pen-ruler.webp' />-->
<!--                </div>-->
<!--              </div>-->
<!--            </div>-->
<!--          </div>-->
<!--        </div>-->
<!--      </div>-->
    {/if}
    <!-- Edit partials -->
    {#if PART_TO_EDIT===EDIT_PICTURE}
      <div class='edit_part'>
        <div class='upper'>
          <button class='back_button' on:click={() => PART_TO_EDIT = ''}>
            <Icon className='iconBack' color='#475569' size={18} src={FaSolidAngleLeft} />
          </button>
          <h3>Upload property photo</h3>
        </div>
        <div class='edit_content'>
          <div class='upload_picture'>
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
                  <Icon size={35} className='upload_icon' color='#848D96' src={FaSolidCamera} />
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
                      <img alt='' src={img} />
                    </div>
                    <div class='image_description'>
                      {#if imageLabels && imageLabels.length}
                        <input placeholder='Description' type='text' bind:value={imageLabels[idx]} />
                      {:else}
                        <input placeholder='Description' type='text' />
                      {/if}
                    </div>
                    <button
                      class='remove_image'
                      title='remove this image'
                      on:click|preventDefault={() => removeImage(idx)}
                    >
                      <Icon color='#FFF' size={12} src={FaSolidTimes} />
                    </button>
                  </div>
                {/each}
              {/if}
            </div>
          </div>
          <button disabled={submitLoading===true} class='next_button' on:click={SaveEditPicture}>
            {#if submitLoading===false}
              <span>Save</span>
            {/if}
            {#if submitLoading===true}
              <Icon className='spin' color='#FFF' size={15} src={FaSolidCircleNotch} />
            {/if}
          </button>
        </div>
      </div>
    {/if}
    {#if PART_TO_EDIT===EDIT_FEATURE}
      <AddOtherFeesDialog
        bind:fee={otherFeeObj.fee}
        bind:name={otherFeeObj.name}
        bind:this={addOtherFeeDialog}
      >
        <button class='add_fee_btn' on:click={addOtherFee}>
          Add
        </button>
      </AddOtherFeesDialog>
      <div class='edit_part'>
        <div class='upper'>
          <button class='back_button' on:click={() => {
            if( editFeatureCount>0 ) {
              editFeatureCount = 0;
            } else if( editFeatureCount===0 ) {
              PART_TO_EDIT = '';
            }
					}}>
            <Icon className='iconBack' color='#475569' size={18} src={FaSolidAngleLeft} />
          </button>
          <h3>Feature</h3>
        </div>
        <div class='edit_content'>
          {#if editFeatureCount===0}
            <div class='feature'>
              <div class='row'>
                <div class='input_box'>
                  <label for='houseType'>House type <span class='asterisk'>*</span></label>
                  <select id='houseType' name='houseType' bind:value={houseType}>
                    {#each houseTypeLists as t}
                      <option value={t}>{t}</option>
                    {/each}
                  </select>
                </div>
              </div>
              <div class='room_area'>
                {#if property.houseType!=='land'}
                  <div class='input_box beds'>
                    <label for='beds' class='labels'>
                      <Icon className='labels_icon' color='#848D96' s size={13} src={FaSolidBed} />
                      <span>Beds</span>
                    </label>
                    <input id='beds' type='number' min='0' bind:value={bedroom} />
                  </div>
                  <div class='input_box baths'>
                    <label for='baths' class='labels'>
                      <Icon className='labels_icon' color='#848D96' size={13} src={FaSolidBath} />
                      <span>Baths</span>
                    </label>
                    <input id='baths' type='number' min='0' bind:value={bathroom} />
                  </div>
                  <div class='input_box livings'>
                    <label for='livings' class='labels'>
                      <Icon className='labels_icon' color='#848D96' size={13} src={FaSolidChair} />
                      <span>Livings</span>
                    </label>
                    <input id='livings' type='number' min='0' bind:value={livingroom} />
                  </div>
                {/if}
                <div class='input_box area'>
                  <label for='area' class='labels'>
                    <Icon className='labels_icon' color='#848D96' size={13} src={FaSolidBorderStyle} />
                    <span>M2</span>
                    <button class='unit_toggle'>
                      <span class='bg'></span>
                      <Icon className='labels_icon' color='#F97316' size={13} src={FaSolidExchangeAlt} />
                    </button>
                  </label>
                  <input id='area' type='number' min='0' bind:value={sizeM2} />
                </div>
              </div>
              <div class='row'>
                <div class='input_box'>
                  <label for='parking'>Parking <span class='asterisk'>*</span></label>
                  <select id='parking' name='parking' bind:value={parking}>
                    <option value='1'>Yes</option>
                    <option value='0'>No</option>
                  </select>
                </div>
              </div>
            </div>
          {/if}
          {#if editFeatureCount===1}
            <div class='feature'>
              <div class='rent_or_sell'>
                <label class={purpose === 'sell' ? 'option clicked': 'option'} for='sell'>
                  <input
                    type='radio'
                    name='rent_or_sell'
                    on:click={() => (purpose = 'sell')}
                    id='sell'
                    value='sell'
                  />
                  Sell
                </label>
                <label class={purpose === 'rent' ? 'option clicked': 'option'} for='rent'>
                  <input
                    type='radio'
                    name='rent_or_sell'
                    on:click={() => (purpose = 'rent')}
                    id='rent'
                    value='rent'
                  />
                  Rent
                </label>
              </div>
              <div class='row'>
                <div class='input_box prop_price'>
                  <label for='price'>{purpose==='sell' ? 'Property Price' : 'Rent'}</label>
                  <input id='price' type='number' min='0' bind:value={lastPrice} />
                  <span>$</span>
                </div>
                {#if purpose==='rent'}
                  <p class='permonth'>/month</p>
                {/if}
              </div>
              <div class='row'>
                <div class='input_box'>
                  <label for='agency_fee'>Agency Fee</label>
                  <select id='agency_fee' bind:value={agencyFee}>
                    <option value='true'>Yes</option>
                    <option value='false'>No</option>
                  </select>
                </div>
                {#if agencyFee==='true'}
                  <div class='input_box agency_fee'>
                    <label for='agency_fee_percent'>Charge to Buyer</label>
                    <input id='agency_fee_percent' type='number' min='0' max='99' bind:value={agencyFeePercent} />
                    <span>%</span>
                  </div>
                {/if}
              </div>
              {#if purpose==='rent'}
                <div class='row'>
                  <div class='input_box deposit_fee'>
                    <label for='deposit_fee'>Deposit Fee</label>
                    <input id='deposit_fee' type='number' min='0' bind:value={depositFee} />
                    <span>$</span> <!-- TODO: Use current country currency sign -->
                  </div>
                  <div class='input_box min_duration'>
                    <label for='minimum_duration'>Minimum Duration</label>
                    <input id='minimum_duration' type='number' min='1' bind:value={minimumDurationYear} />
                    <span>Year</span>
                  </div>
                </div>
                <div class='other_fee'>
                  <header>
                    <h4>Other Fee</h4>
                    <button class='add_fee' on:click={addOtherFeeDialog.showModal()}>
                      Add
                    </button>
                  </header>
                  <div class='other_fee_lists'>
                    {#if otherFees && otherFees.length}
                      {#each otherFees as otherFee}
                        <div class='fee'>
                          <span>{otherFee.name}</span>
                          <b>{formatPrice( otherFee.fee, countryCurrency )}/mo</b>
                        </div>
                      {/each}
                    {:else}
                      <div class='fee'>
                        <span>Example Fee #1</span>
                        <span>$$$</span>
                      </div>
                      <div class='fee'>
                        <span>Example Fee #1</span>
                        <span>$$$</span>
                      </div>
                    {/if}
                  </div>
                </div>
              {/if}
            </div>
          {/if}
          <button disabled={submitLoading===true} class='next_button' on:click={SaveEditFeature}>
            {#if submitLoading===false}
							<span>
								{editFeatureCount===0 ? 'Next' : 'Save'}
							</span>
            {/if}
            {#if submitLoading===true}
              <Icon className='spin' color='#FFF' size={15} src={FaSolidCircleNotch} />
            {/if}
          </button>
        </div>
      </div>
    {/if}
    {#if PART_TO_EDIT===EDIT_FACILITY}
      <div class='edit_part'>
        <div class='upper'>
          <button class='back_button' on:click={() => PART_TO_EDIT = ''}>
            <Icon className='iconBack' color='#475569' size={18} src={FaSolidAngleLeft} />
          </button>
          <h3>Facility</h3>
        </div>
        <div class='edit_content'>
          <div class='facility'>
						<textarea
              bind:value={mainUse}
              name='facility'
              id='facility'
              placeholder='Facility'
              class='textarea_input'
            ></textarea>
          </div>
          <button disabled={submitLoading===true} class='next_button' on:click={SaveEditFacility}>
            {#if submitLoading===false}
              <span>Save</span>
            {/if}
            {#if submitLoading===true}
              <Icon className='spin' color='#FFF' size={15} src={FaSolidCircleNotch} />
            {/if}
          </button>
        </div>
      </div>
    {/if}
    {#if PART_TO_EDIT===EDIT_ABOUT}
      <div class='edit_part'>
        <div class='upper'>
          <button class='back_button' on:click={() => PART_TO_EDIT = ''}>
            <Icon className='iconBack' color='#475569' size={18} src={FaSolidAngleLeft} />
          </button>
          <h3>Facility</h3>
        </div>
        <div class='edit_content'>
          <div class='about'>
						<textarea
              bind:value={note}
              name='about'
              id='about'
              placeholder='About your property...'
              class='textarea_input'
            ></textarea>
          </div>
          <button disabled={submitLoading===true} class='next_button' on:click={SaveEditAbout}>
            {#if submitLoading===false}
              <span>Save</span>
            {/if}
            {#if submitLoading===true}
              <Icon className='spin' color='#FFF' size={15} src={FaSolidCircleNotch} />
            {/if}
          </button>
        </div>
      </div>
    {/if}
  </div>
  <div class='delete_property_container'>
    <button class='delete_property'>
      <Icon color='#FFF' size={10} src={FaTrashAlt} />
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

  .textarea_input {
    white-space      : normal;
    height           : 300px;
    resize           : none;
    width            : 100%;
    border           : 1px solid #CBD5E1;
    background-color : #F1F5F9;
    border-radius    : 8px;
    padding          : 12px;
  }

  .textarea_input:focus {
    border-color : #3B82F6;
    outline      : 1px solid #3B82F6;
  }

  .textarea_input::-webkit-scrollbar {
    width : 7px;
  }

  .textarea_input::-webkit-scrollbar-thumb {
    background-color : #3B82F6;
    border-radius    : 7px;
  }

  .textarea_input::-webkit-scrollbar-track {
    background : transparent;
  }


  .add_fee_btn {
    background-color : #F97316;
    color            : white;
    border-radius    : 8px;
    border           : none;
    padding          : 10px;
    cursor           : pointer;
    width            : 100%;
  }

  .room_area {
    display        : flex;
    flex-direction : row;
    gap            : 20px;
  }

  .input_box {
    display        : flex;
    flex-direction : column;
    gap            : 8px;
    width          : 100%;
  }

  .input_box label {
    font-size   : 13px;
    font-weight : 700;
    margin-left : 10px;
  }

  .input_box input,
  .input_box select {
    width            : 100%;
    border           : 1px solid #CBD5E1;
    background-color : #F1F5F9;
    border-radius    : 8px;
    padding          : 12px;
    text-transform   : capitalize;
  }

  .input_box input:focus,
  .input_box select:focus {
    border-color : #3B82F6;
    outline      : 1px solid #3B82F6;
  }

  .row {
    display               : grid;
    grid-template-columns : 1fr 1fr;
    gap                   : 20px;
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

  .unit_toggle {
    border     : none;
    background : transparent;
    position   : relative;
    cursor     : pointer;
  }

  .unit_toggle .bg {
    width            : 0;
    height           : 0;
    border-radius    : 50%;
    background-color : rgb(0 0 0 / 0.06);
    z-index          : 1;
    position         : absolute;
    top              : -4px;
    left             : 0;
  }

  .unit_toggle:hover .bg {
    width  : 24px;
    height : 24px;
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

  .edit_property_container .backdrop_reject_reason {
    position        : absolute;
    z-index         : 40;
    background      : rgba(41, 41, 41, 0.5);
    width           : 100%;
    left            : 0;
    top             : 0;
    bottom          : 0;
    height          : auto;
    border-radius   : 8px;
    display         : flex;
    justify-content : center;
  }

  .edit_property_container .backdrop_reject_reason .reject_reason {
    color            : #334155;
    width            : 400px;
    height           : fit-content;
    gap              : 20px;
    background-color : white;
    padding          : 20px;
    margin           : 40px 15px 0;
    border-radius    : 8px;
    display          : flex;
    flex-direction   : column;
    justify-content  : space-between;
  }

  .edit_property_container .backdrop_reject_reason .reject_reason .reject_reason_content {
    display        : flex;
    flex-direction : column;
    gap            : 15px;
  }

  .edit_property_container .backdrop_reject_reason .reject_reason .reject_reason_content h3 {
    text-align : center;
  }

  .edit_property_container .backdrop_reject_reason .reject_reason .buttons {
    display        : flex;
    flex-direction : row;
    gap            : 10px;
    align-items    : stretch;
    font-weight    : 500;
    width          : 100%;
  }

  .edit_property_container .backdrop_reject_reason .reject_reason .buttons button {
    border-radius  : 8px;
    border         : none;
    padding        : 10px;
    cursor         : pointer;
    width          : 100%;
    text-transform : capitalize;
  }

  .edit_property_container .backdrop_reject_reason .reject_reason .buttons .cancel_btn {
    border           : 1px solid #CBD5E1;
    background-color : #F1F5F9;
    color            : #F97316;
  }

  .edit_property_container .backdrop_reject_reason .reject_reason .buttons .cancel_btn:hover {
    border : 1px solid #F97316;
  }

  .edit_property_container .backdrop_reject_reason .reject_reason .buttons .reject_btn {
    border           : 1px solid rgba(255, 126, 118, 1);
    background-color : rgba(255, 126, 118, 1);
    color            : #FFF;
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

  .edit_property_container .property_status .action_btns {
    display        : flex;
    flex-direction : row;
    align-items    : center;
    gap            : 10px;
  }

  .edit_property_container .property_status .approve_btn {
    display          : flex;
    flex-direction   : row;
    gap              : 6px;
    align-items      : center;
    color            : #FFF;
    background-color : rgba(140, 216, 107, 1);
    padding          : 5px 15px;
    font-weight      : 600;
    border           : none;
    border-radius    : 8px;
    cursor           : pointer;
    width            : fit-content;
  }

  .edit_property_container .property_status .approve_btn:hover {
    background-color : rgb(118, 216, 75);
  }

  .edit_property_container .property_status .reject_btn {
    display          : flex;
    flex-direction   : row;
    gap              : 6px;
    align-items      : center;
    color            : #FFF;
    background-color : rgba(255, 126, 118, 1);
    padding          : 5px 15px;
    font-weight      : 600;
    border           : none;
    border-radius    : 8px;
    cursor           : pointer;
    width            : fit-content;
  }

  .edit_property_container .property_status .reject_btn:hover,
  .edit_property_container .backdrop_reject_reason .reject_reason .buttons .reject_btn:hover {
    background-color : rgb(248, 106, 96);
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
    color           : #848D96;
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

  /* +=== UPDATE FEATURE ===+ */

  .edit_property_container .edit_part .edit_content .feature {
    display        : flex;
    flex-direction : column;
    gap            : 20px;
  }

  .edit_property_container .edit_part .edit_content .feature .room_area {
    display        : flex;
    flex-direction : row;
    align-items    : flex-start;
    gap            : 20px;
  }

  .edit_property_container .edit_part .edit_content .feature .room_area .input_box {
    flex-direction : column-reverse !important;
  }

  .edit_property_container .edit_part .edit_content .feature .room_area .beds label,
  .edit_property_container .edit_part .edit_content .feature .room_area .baths label,
  .edit_property_container .edit_part .edit_content .feature .room_area .livings label,
  .edit_property_container .edit_part .edit_content .feature .room_area .area label {
    display         : flex !important;
    flex-direction  : row !important;
    justify-content : left !important;
    align-items     : center !important;
    gap             : 8px !important;
    color           : #848D96;
  }

  .edit_property_container .edit_part .edit_content .feature .permonth {
    line-height : 1em;
    margin      : 20px 0 0;
    font-size   : 17px;
    font-weight : 500
  }

  .edit_property_container .edit_part .edit_content .feature .rent_or_sell {
    width                 : 100%;
    display               : grid;
    grid-template-columns : 1fr 1fr;
    border-collapse       : collapse;
  }

  .edit_property_container .edit_part .edit_content .feature .rent_or_sell .option {
    margin           : 0;
    padding          : 10px 12px;
    border           : 1px solid #CBD5E1;
    background-color : #F1F5F9;
    font-weight      : 500;
    text-align       : center;
    cursor           : pointer;
  }

  .edit_property_container .edit_part .edit_content .feature .rent_or_sell .option:nth-child(1) {
    border-top-left-radius    : 8px;
    border-bottom-left-radius : 8px;
  }

  .edit_property_container .edit_part .edit_content .feature .rent_or_sell .option:nth-child(2) {
    border-top-right-radius    : 8px;
    border-bottom-right-radius : 8px;
  }

  .edit_property_container .edit_part .edit_content .feature .rent_or_sell .option:hover {
    border : 1px solid #F97316;
    color  : #F97316;
  }

  .edit_property_container .edit_part .edit_content .feature .rent_or_sell .option.clicked {
    background-color : #F97316;
    color            : white;
    border           : none;
  }

  .edit_property_container .edit_part .edit_content .feature .rent_or_sell .option input[type='radio'] {
    position       : absolute;
    opacity        : 0;
    pointer-events : none;
  }

  .edit_property_container .edit_part .edit_content .feature .row {
    display               : grid;
    grid-template-columns : 1fr 1fr;
    gap                   : 20px;
    align-items           : center;
  }

  .edit_property_container .edit_part .edit_content .feature .row .prop_price,
  .edit_property_container .edit_part .edit_content .feature .row .agency_fee,
  .edit_property_container .edit_part .edit_content .feature .row .deposit_fee,
  .edit_property_container .edit_part .edit_content .feature .row .min_duration {
    position : relative;
  }

  .edit_property_container .edit_part .edit_content .feature .row .agency_fee input {
    padding : 12px 25px 12px 12px;
  }

  .edit_property_container .edit_part .edit_content .feature .row .agency_fee span {
    position  : absolute;
    right     : 10px;
    top       : 35px;
    font-size : 15px;
    color     : black;
  }

  .edit_property_container .edit_part .edit_content .feature .row .prop_price input,
  .edit_property_container .edit_part .edit_content .feature .row .deposit_fee input {
    padding : 12px 12px 12px 25px;
  }

  .edit_property_container .edit_part .edit_content .feature .row .prop_price span,
  .edit_property_container .edit_part .edit_content .feature .row .deposit_fee span {
    position  : absolute;
    left      : 10px;
    top       : 35px;
    font-size : 14px;
    color     : black;
  }

  .edit_property_container .edit_part .edit_content .feature .row .min_duration input {
    padding : 12px 42px 12px 12px;
  }

  .edit_property_container .edit_part .edit_content .feature .row .min_duration span {
    position  : absolute;
    right     : 10px;
    top       : 35px;
    font-size : 15px;
    color     : #F97316;
  }

  .edit_property_container .edit_part .edit_content .feature .other_fee {
    display        : flex;
    flex-direction : column;
    gap            : 15px;
  }

  .edit_property_container .edit_part .edit_content .feature .other_fee header {
    display         : flex;
    flex-direction  : row;
    justify-content : space-between;
    align-items     : center;
    margin-top      : 20px;
  }

  .edit_property_container .edit_part .edit_content .feature .other_fee header h4 {
    margin    : 0;
    font-size : 18px;
  }

  .edit_property_container .edit_part .edit_content .feature .other_fee header .add_fee {
    color            : #F97316;
    border           : none;
    padding          : 10px 15px;
    font-weight      : 600;
    font-size        : 15px;
    border-radius    : 8px;
    cursor           : pointer;
    background-color : transparent;
  }

  .edit_property_container .edit_part .edit_content .feature .other_fee header .add_fee:hover {
    background-color : #F1F5F9;
  }

  .edit_property_container .edit_part .edit_content .feature .other_fee .other_fee_lists {
    padding          : 15px;
    background-color : #F1F5F9;
    border-radius    : 8px;
    display          : flex;
    flex-direction   : column;
    min-height       : 200px;
  }

  .edit_property_container .edit_part .edit_content .feature .other_fee .other_fee_lists .fee {
    display         : flex;
    flex-direction  : row;
    justify-content : space-between;
    padding         : 10px 0;
    border-bottom   : 1px solid #CBD5E1;
    font-size       : 15px;
    font-weight     : 500;
  }

  /* Responsive to mobile device */
  @media (max-width : 768px) {
    .unit_toggle .bg {
      top  : -4px;
      left : 1px;
    }

    .unit_toggle:hover .bg {
      width  : 20px;
      height : 20px;
    }

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

    .edit_property_container .main_details .col2 .feature_item .labels,
    .edit_property_container .edit_part .edit_content .feature .room_area .input_box .labels {
      font-size : 10px;
      gap       : 1px;
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

    .edit_property_container .edit_part .edit_content .feature {
      gap : 15px;
    }

    .edit_property_container .edit_part .edit_content .feature .room_area {
      gap : 10px !important;
    }
  }
</style>