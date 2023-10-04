<script>
  // @ts-nocheck
  import {onMount} from 'svelte';
  import Menu from '../_components/Menu.svelte';
  import ProfileHeader from '../_components/ProfileHeader.svelte';
  import Footer from '../_components/Footer.svelte';
  import AddFloorDialog from '../_components/AddFloorDialog.svelte';
  import AddOrEditRoomDialog from '../_components/AddOrEditRoomDialog.svelte';
  import {RealtorUpsertProperty} from '../jsApi.GEN';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidAngleLeft from 'svelte-icons-pack/fa/FaSolidAngleLeft';
  import FaSolidMapMarkerAlt from 'svelte-icons-pack/fa/FaSolidMapMarkerAlt';
  import FaSolidFlagUsa from 'svelte-icons-pack/fa/FaSolidFlagUsa';
  import FaSolidCloudUploadAlt from 'svelte-icons-pack/fa/FaSolidCloudUploadAlt';
  import FaSolidImage from "svelte-icons-pack/fa/FaSolidImage";
  import FaSolidPlusCircle from 'svelte-icons-pack/fa/FaSolidPlusCircle';
  import FaSolidPen from 'svelte-icons-pack/fa/FaSolidPen';
  import FaSolidTrashAlt from 'svelte-icons-pack/fa/FaSolidTrashAlt';
  import FaSolidHome from 'svelte-icons-pack/fa/FaSolidHome';
  import FaHeart from 'svelte-icons-pack/fa/FaHeart';
  import FaSolidBan from 'svelte-icons-pack/fa/FaSolidBan';
  import FaCheckCircle from 'svelte-icons-pack/fa/FaCheckCircle';
  import siElsevier from "svelte-icons-pack/si/SiElsevier";
  
  // Use property from backend if backend is ready
  let property = {/* property */};
  let user = {/* user */};
  let segments = {/* segments */};
  let countryData = {/* countryData */}
  let currentPage = 0;
  let cards = [{}, {}, {}, {}];
  
  function progressDotHandler( toPage ) {
    let card = cards[ toPage ];
    card.scrollIntoView( {behavior: 'smooth'} );
  }
  
  function nextPage() {
    if( currentPage<3 ) {
      currentPage++;
      let card = cards[ currentPage ];
      card.scrollIntoView( {behavior: 'smooth'} );
    }
  }
  
  function backPage() {
    if( currentPage>0 ) {
      currentPage--;
      let card = cards[ currentPage ];
      card.scrollIntoView( {behavior: 'smooth'} );
    }
  }
  
  // +=============| Location and Signage |=============+ //
  const LOC_ADDR = 'Address';
  const LOC_MAP = 'Put the pin on your house location by clicking the map';
  const LOC_STREETVIEW = 'Put signage on your house location';
  let modeLocationCount = 0, nextLocationLocked = true;
  const modeLocationLists = [
    {mode: LOC_ADDR},
    {mode: LOC_MAP},
    {mode: LOC_STREETVIEW}
  ];
  let modeLocation = modeLocationLists[ modeLocationCount ].mode;
  let locationObj = {
    country: 'Afghanistan',
    city: '',
    street1: '',
    street2: '',
    floors: 0,
    coord: {
      lat: 0,
      lng: 0
    }
  }
  $: if( locationObj.country==='' || locationObj.city==='' || locationObj.street1==='' ) {
    nextLocationLocked = true;
  } else {
    nextLocationLocked = false;
  }
  
  const handleNextLocation = {
    'LOC_ADDR': () => {
      if( locationObj.country==='' || locationObj.city==='' || locationObj.street1==='' ) {
        alert( 'Please fill required field' );
        return
      }
      modeLocationCount++;
      modeLocation = modeLocationLists[ modeLocationCount ].mode;
    },
    'LOC_MAP': () => {
      nextPage();
    }
  }
</script>

<section class='dashboard'>
	<Menu
		access={segments}
	/>
	<div class='dashboard_main_content'>
		<ProfileHeader></ProfileHeader>
		<div class='realtor_step_progress_bar'>
			<div class='container'>
				<a class='back_button' href='/realtor'>
					<Icon className="iconBack" color='#475569' size={18} src={FaSolidAngleLeft}/>
				</a>
				<div class='step_wrapper'>
					<div class={currentPage >= 0 ? 'step_item completed' : 'step_item active'}>
						<button on:click={() => progressDotHandler(0)}></button>
						<p>Location</p>
					</div>
					<div class={currentPage === 1 ? 'step_item active' : 'step_item' && currentPage > 1 ? 'step_item completed' : 'step_item'}>
						<button on:click={() => progressDotHandler(1)}></button>
						<p>Info</p>
					</div>
					<div class={currentPage === 2 ? 'step_item active' : 'step_item' && currentPage > 2 ? 'step_item completed' : 'step_item'}>
						<button on:click={() => progressDotHandler(2)}></button>
						<p>Picture</p>
					</div>
					<div class={currentPage === 3 ? 'step_item active' : 'step_item'}>
						<button on:click={() => progressDotHandler(3)}></button>
						<p>Preview</p>
					</div>
				</div>
			</div>
		</div>
		<div class='content'>
			<div class='realtor_subpage_container'>
				<section bind:this={cards[0]} class='location' id='subpage_1'>
					{#if currentPage!==0}
						<button class='back_button' on:click={backPage}>
							<Icon className="iconBack" color='#475569' size={18} src={FaSolidAngleLeft}/>
						</button>
					{/if}
					<div class='subpage_content'>
						<h3>{ modeLocation }</h3>
						{#if modeLocation===LOC_ADDR}
							<div class='address'>
								<div class='row'>
									<div class='input_box'>
										<label for='country'>Country or Region <span class='asterisk'>*</span></label>
										<select id='country' name='country' bind:value={locationObj.country}>
											{#each countryData as country}
												<option value={country.country}>{country.country}</option>
											{/each}
										</select>
									</div>
									<div class='input_box'>
										<label for='city'>City or County <span class='asterisk'>*</span></label>
										<input id='city' type='text' placeholder='Required' bind:value={locationObj.city}/>
									</div>
								</div>
								<div class='row'>
									<div class='input_box'>
										<label for='street1'>Street 1 <span class='asterisk'>*</span></label>
										<input id='street1' type='text' placeholder='Required' bind:value={locationObj.street1}/>
									</div>
									<div class='input_box'>
										<label for='street2'>Street 2</label>
										<input id='street2' type='text' placeholder='Optional' bind:value={locationObj.street2}/>
									</div>
								</div>
								<div class='row'>
									<div class='input_box'>
										<label for='floors'>Floors</label>
										<input id='floors' type='number' placeholder='10' min='0' bind:value={locationObj.floors}/>
									</div>
								</div>
							</div>
						{/if}
						{#if modeLocation===LOC_MAP}
							<div class='address'>
								<div class='row'>
									<div class='input_box'>
										<label for='country'>Country or Region <span class='asterisk'>*</span></label>
										<select id='country' name='country'>
											{#each countryData as country}
												<option value={country.country}>{country.country}</option>
											{/each}
										</select>
									</div>
									<div class='input_box'>
										<label for='city'>City or County <span class='asterisk'>*</span></label>
										<input id='city' type='text' placeholder='Required'/>
									</div>
								</div>
								<div class='row'>
									<div class='input_box'>
										<label for='street1'>Street 1 <span class='asterisk'>*</span></label>
										<input id='street1' type='text' placeholder='Required'/>
									</div>
									<div class='input_box'>
										<label for='street2'>Street 2</label>
										<input id='street2' type='text' placeholder='Optional'/>
									</div>
								</div>
								<div class='row'>
									<div class='input_box'>
										<label for='floors'>Floors</label>
										<input id='floors' type='number' placeholder='10' min='0'/>
									</div>
								</div>
							</div>
						{/if}
					</div>
					<button
						class='next_button'
						disabled={nextLocationLocked === true}
						on:click={() => {
                     if (modeLocationCount === 0) {
                       handleNextLocation.LOC_ADDR()
                     }
                  }}>
						{#if nextLocationLocked===true}
							<Icon size={16} color='#FFF' src={FaSolidBan}/>
						{/if}
						{#if nextLocationLocked===false}
							<span>NEXT</span>
						{/if}
					</button>
				</section>
				<section bind:this={cards[1]} class='info' id='subpage_2'>
					<button class='back_button'>
						<Icon className="iconBack" color='#475569' size={18} src={FaSolidAngleLeft}/>
					</button>
				</section>
				<section bind:this={cards[2]} class='picture' id='subpage_3'>
					<button class='back_button'>
						<Icon className="iconBack" color='#475569' size={18} src={FaSolidAngleLeft}/>
					</button>
				</section>
				<section bind:this={cards[3]} class='preview' id='subpage_4'>
					<button class='back_button'>
						<Icon className="iconBack" color='#475569' size={18} src={FaSolidAngleLeft}/>
					</button>
				</section>
			</div>
		</div>
		<Footer></Footer>
	</div>
</section>

<style>
    /* General purpose selector*/
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

    .realtor_subpage_container section .back_button {
        padding          : 8px;
        border           : none;
        background-color : rgb(0 0 0 / 0.06);
        border-radius    : 5px;
        font-size        : 14px;
        cursor           : pointer;
        display          : flex;
        flex-direction   : row;
        align-items      : center;
        gap              : 8px;
        position         : absolute;
        left             : 20px;
        top              : 20px;
    }

    .realtor_subpage_container section .back_button:hover {
        background-color : rgb(0 0 0 / 0.05);
        color            : #EF4444;
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
    }

    .input_box input:focus,
    .input_box select:focus {
        border-color : #3B82F6;
        outline      : 1px solid #3B82F6;
    }

    .next_button {
        background-color : #F97316;
        color            : white;
        display          : flex;
        justify-content  : center;
        border-radius    : 8px;
        border           : none;
        padding          : 10px;
        cursor           : pointer;
        width            : 100%;
        font-weight      : 600;
    }

    .next_button:disabled,
    .next_button:hover {
        background-color : #F58433;
    }

    .asterisk {
        color      : #EF4444;
        font-size  : 14px;
        margin-top : -3px;
    }

    /* =========================*/

    .realtor_step_progress_bar {
        position         : relative;
        margin-top       : -40px;
        margin-left      : auto;
        margin-right     : auto;
        display          : flex;
        justify-content  : center;
        background-color : #EF4444;
        padding-bottom   : 70px;
    }

    .realtor_step_progress_bar .container {
        width            : 940px;
        margin-left      : auto;
        margin-right     : auto;
        display          : flex;
        flex-direction   : row;
        justify-content  : center;
        align-items      : center;
        background-color : white;
        border-radius    : 10px;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
        padding          : 13px 5%;
    }

    .realtor_step_progress_bar .container .back_button {
        padding          : 8px;
        background-color : rgb(0 0 0 / 0.06);
        border-radius    : 5px;
        font-size        : 14px;
        cursor           : pointer;
        border           : none;
        text-decoration  : none;
        height           : fit-content;
        color            : #334155;
    }

    .realtor_step_progress_bar .container .back_button:hover {
        color : #EF4444;
    }

    .realtor_step_progress_bar .step_wrapper {
        width           : 500px;
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
        padding-top     : 5px;
    }

    .realtor_step_progress_bar .step_wrapper .step_item {
        position        : relative;
        display         : flex;
        flex-direction  : column;
        justify-content : center;
        align-items     : center;
        flex            : 1;
        font-size       : 12px;
        color           : #334155;
    }

    .realtor_step_progress_bar .step_wrapper .step_item::before {
        position      : absolute;
        content       : '';
        border-bottom : 2px solid #CBD5E1;
        width         : 100%;
        top           : 5px;
        left          : -50%;
        z-index       : 2;
    }

    .realtor_step_progress_bar .step_wrapper .step_item::after {
        position      : absolute;
        content       : '';
        border-bottom : 2px solid #CBD5E1;
        width         : 100%;
        top           : 5px;
        left          : 50%;
        z-index       : 2;
    }

    .realtor_step_progress_bar .step_wrapper .step_item:first-child::before {
        content : none;
    }

    .realtor_step_progress_bar .step_wrapper .step_item:last-child::after {
        content : none;
    }

    .realtor_step_progress_bar .step_wrapper .step_item button {
        width            : 13px;
        height           : 13px;
        background-color : #CBD5E1;
        border-radius    : 100%;
        border           : none;
        z-index          : 4;
        cursor           : pointer;
    }

    .realtor_step_progress_bar .step_wrapper .step_item button:hover {
        outline : 5px solid rgb(0 0 0 / 0.09);
    }

    .realtor_step_progress_bar .step_wrapper .step_item p {
        margin : 8px 0 0 0;
    }

    .realtor_step_progress_bar .step_wrapper .step_item.active button {
        width            : 11px;
        height           : 11px;
        background-color : white;
        outline          : 3px solid #F97316;
        cursor           : pointer;
    }

    .step_item.completed::after {
        position      : absolute !important;
        content       : '' !important;
        border-bottom : 2px solid #F97316 !important;
        width         : 100% !important;
        top           : 5px !important;
        left          : 50% !important;
        z-index       : 3 !important;
    }

    .step_item.completed button {
        background-color : #F97316 !important;
    }

    /* Main Container Subpage */
    .realtor_subpage_container {
        color            : #475569;
        position         : relative;
        margin-top       : -40px;
        margin-left      : auto;
        margin-right     : auto;
        display          : flex;
        flex-direction   : row;
        gap              : 20px;
        border-radius    : 10px;
        width            : 940px;
        min-height       : 700px;
        overflow-x       : scroll;
        scroll-snap-type : x mandatory;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    }

    .realtor_subpage_container::-webkit-scrollbar-thumb {
        background-color : #EF4444;
        border-radius    : 4px;
    }

    .realtor_subpage_container::-webkit-scrollbar {
        height : 0;
        width  : 10px;
    }

    .realtor_subpage_container::-webkit-scrollbar-track {
        background-color : transparent;
    }

    .realtor_subpage_container section {
        padding           : 20px;
        background-color  : white;
        border-radius     : 10px;
        min-height        : 700px;
        flex              : 0 0 940px;
        scroll-snap-align : start;
        position          : relative;
        display           : flex;
        flex-direction    : column;
        justify-content   : space-between;
    }

    /* +============| SUBPAGE LOCATION |===========+ */
    .realtor_subpage_container section.location .subpage_content {
        display        : flex;
        flex-direction : column;
        height         : fit-content !important;
    }

    .realtor_subpage_container section.location .subpage_content h3 {
        margin     : 8px 0;

        text-align : center;
        font-size  : 20px;
    }

    .realtor_subpage_container section.location .subpage_content .address {
        margin-top     : 20px;
        display        : flex;
        flex-direction : column;
        gap            : 15px;
    }

    .realtor_subpage_container section.location .subpage_content .address .row {
        display               : grid;
        grid-template-columns : 1fr 1fr;
        gap                   : 20px;
    }
</style>