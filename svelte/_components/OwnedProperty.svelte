<script>
    import { Icon } from '../node_modules/svelte-icons-pack/dist';
    import {
        FaSolidHouse, FaSolidPen, FaSolidBed, FaSolidBath,
				FaSolidChair, FaSolidBorderTopLeft, FaSolidMapLocationDot
    } from '../node_modules/svelte-icons-pack/dist/fa';
    import PillBox from './PillBox.svelte';
    import {localeDatetime, M2ToPing, getApprovalState} from './formatter';
    import {T} from './uiState.js';
    import { onMount } from 'svelte';
    import { notifier } from './notifier';
    
    export let property;
    export let meta;
    let showMore = false;
  
    let approvalStatus = 'approved';
    let noteObj;
    onMount(() => {
      try {
        noteObj = JSON.parse( property.note );
      } catch (/** @type {any | Error} */ e) {
        noteObj = {
          contactPhone:"",
          contactEmail: "",
          about: property.note
        }
        notifier.showError('Error convert string to object = ' + e)
      }
      approvalStatus = getApprovalState( property.approvalState );
      if (property.countryCode==='TW' || property.countryCode==='') {
          property.sizeM2 = M2ToPing(property.sizeM2);
      }
    })
    
    function handleShowMore() {
      showMore = !showMore;
    }
  </script>
  
  <div class={showMore ? 'property_more' : 'property'}>
      <div class={`approve_status ${approvalStatus}`}>
          {approvalStatus}
      </div>
      <div class='property_main'>
          <div class='property_images'>
              {#if property.images && property.images.length}
                  <img src={property.images[0]} alt=''/>
              {:else}
                  <div class='image_empty'>
                      <i class='gg-image'/>
                  </div>
              {/if}
          </div>
          <div class='property_info'>
              <div class='col1'>
                  <div class='left'>
                      <div class={property.purpose === 'rent' ? `purpose label_rent` : `purpose label_sale`}>
                          {property.purpose==='rent' ? $T.forRent : $T.onSale}
                      </div>
                      <div class='house_type'>
                          <Icon color='#FFF' size="14" src={FaSolidHouse}/>
                          <span>{property.houseType}</span>
                      </div>
                  </div>
                  <a class='edit_property' href='/realtor/property/{property.id}'>
                      <Icon color='#FFF' size="13" src={FaSolidPen}/>
                      <span>Edit Property</span>
                  </a>
              </div>
              <div class='col2'>
                  <h1>$ {property.lastPrice || '0.00'}</h1>
                  <p>{$T.agencyFee} : {property.agencyFeePercent}%</p>
                  <div class='address'>
                      <Icon color='#f97316' className="icon_address" size="18" src={FaSolidMapLocationDot}/>
                      <span>{property.formattedAddress}</span>
                  </div>
              </div>
          </div>
      </div>
      <div class='property_secondary'>
          <div class='col2'>
              <div class='feature_item'>
                <b>{property.bedroom}</b>
                <div class='labels'>
                  <Icon className='labels_icon' color='#848D96' size="13" src={FaSolidBed} />
                  <span>Beds</span>
                </div>
              </div>
              <div class='feature_item'>
                <b>{property.bathroom}</b>
                <div class='labels'>
                  <Icon className='labels_icon' color='#848D96' size="13" src={FaSolidBath} />
                  <span>Baths</span>
                </div>
              </div>
              <div class='feature_item'>
                <b>{property.livingroom}</b>
                <div class='labels'>
                  <Icon className='labels_icon' color='#848D96' size="12" src={FaSolidChair} />
                  <span>Livings</span>
                </div>
              </div>
              <div class='feature_item'>
                <b>{property.sizeM2}</b>
                <div class='labels'>
                  <Icon className='labels_icon' color='#848D96' size="13" src={FaSolidBorderTopLeft} />
                  <span>{property.countryCode==='TW' || property.countryCode==='' ? 'Ping' : 'M2'}</span>
                </div>
              </div>
            </div>
      </div>
      <div class='property_attributes'>
          {#each meta as m}
              {#if property[ m.name ]}
                  {#if m.inputType==='datetime'}
                      <PillBox label={m.label} content={localeDatetime(property[m.name])}/>
                  {:else}
                      {#if m.name==='note' && ( noteObj && noteObj.about)}
                          <PillBox label={m.label} content={noteObj.about}/>
                      {:else}
                          <PillBox label={m.label} content={property[m.name]}/>
                      {/if}
                  {/if}
              {/if}
          {/each}
      </div>
      {#if property.floorList && property.floorList.length}
      <div class='property_floors'>
          <h3>{$T.floors}</h3>
          <div class='floor_lists'>
              {#each property.floorList as floors}
                  <div class='floor_item'>
                      <div class='left'>
                          <h5>
                              {floors.type==='basement' ? $T.basement : `${$T.floorN}${floors.floor}`}
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
      </div>
      {/if}
      <div class='property_less_more'>
          <button class='toggle_show_more' on:click={handleShowMore}>
              Show {showMore===true ? 'Less' : 'More'}
          </button>
      </div>
  </div>
  
  <style>
      .unit_toggle {
      border     : none;
      background : transparent;
      position   : relative;
      cursor     : pointer;
    }
    
      .property {
          background-color : #F0F0F0;
          border-radius    : 8px;
          height           : 250px;
          overflow-y       : hidden;
          padding          : 15px 15px 70px 15px;
          display          : flex;
          flex-direction   : column;
          gap              : 20px;
          position         : relative;
      }
  
      .property_more {
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
          gap             : 15px;
      }
  
      .approve_status {
          padding          : 7px 18px;
          border-radius    : 8px;
          font-size        : 14px;
          text-transform   : capitalize;
          text-decoration  : none;
      }
  
      .approve_status.approved {
          background-color : rgba(140, 216, 107, 1);
          color            : #FFF;
      }
  
      .approve_status.pending {
          background-color : rgba(255, 208, 118, 1);
          color            : #475569;
      }
  
      .approve_status.rejected {
          background-color : rgba(255, 126, 118, 1);
          color            : #FFF;
      }
  
      .property_main .property_images {
          width         : 260px;
          height        : 160px;
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
  
      :global(.icon_address) {
          flex-shrink : 0;
      }
  
      .property_main .property_info .col2 .address span {
          flex-shrink: 1;
      }
  
      .property_main .property_info .col2 .address {
          display        : flex;
          flex-direction : row;
          gap            : 10px;
          align-items    : center;
      }
  
      .property_secondary {
          display        : flex;
          flex-direction : column;
          gap            : 15px;
      }
  
      .property_secondary .col2 {
      display         : flex;
      flex-direction  : row;
      justify-content : space-around;
      align-items     : center;
      text-align      : center;
    }
  
    .property_secondary .col2 .feature_item {
      display        : flex;
      flex-direction : column;
      gap            : 8px;
    }
  
    .property_secondary .col2 .feature_item b {
      font-size : 25px;
    }
  
    .property_secondary .col2 .feature_item .labels {
      display         : flex;
      flex-direction  : row;
      justify-content : center;
      gap             : 6px;
      color           : #848D96;
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
  
      .property_less_more {
          height          : 50px;
          background      : linear-gradient(to top, #F0F0F0, #F0F0F0, #F0F0F0, transparent);
          bottom          : 0;
          position        : absolute;
          z-index         : 90;
          width           : -webkit-fill-available;
          display         : flex;
          justify-content : center;
      }
  
      .toggle_show_more {
          color            : #F97316;
          border           : none;
          cursor           : pointer;
          font-weight      : 700;
          background-color : transparent;
          text-align       : center;
          width            : fit-content;
          bottom           : 10px;
          padding-left     : 1em;
      }
  
      .toggle_show_more:hover {
          color           : #F58433;
          text-decoration : underline;
      }
  
      /* Responsive to mobile device */
      @media (max-width: 768px) {
          .property_main {
              flex-direction  : column;
              flex-wrap       : wrap;
              gap             : 15px;
          }
  
          .property_main .property_images {
              width         : 100%;
              height        : 180px;
          }
  
          .property_main .property_info {
              width: 100%;
          }
  
          .property_main .property_info .col1 {
              display         : flex;
              justify-content : space-between;
              align-items     : center;
              flex-direction  : row;
              font-size: 10px;
              width: 100%;
          }
  
          .property_main .property_info .col1 .left {
              gap         : 10px;
          }
  
          .property_main .property_info .col1 .left .purpose,
          .property_main .property_info .col1 .left .house_type,
          .property_main .property_info .col1 .edit_property {
              padding          : 8px 10px;
              border-radius    : 8px;
              font-size        : 12px;
          }
  
          .property_main .property_info .col1 .edit_property span {
              display : none;
          }
      }
  </style>