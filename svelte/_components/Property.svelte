<script>
  /** @typedef {import('../_types/property').Property} TypeProperty */
  /** @typedef {import('../_types/property').PropertyUS} TypePropertyUS */
  /** @typedef {import('../_types/property').PropertyExtraUS} TypePropertyExtraUS */

  import FaSolidImage from 'svelte-icons-pack/fa/FaSolidImage';
  import FaSolidHome from 'svelte-icons-pack/fa/FaSolidHome';
  import { formatPrice, localeDatetime, M2ToPing } from './formatter';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidMapMarkerAlt from 'svelte-icons-pack/fa/FaSolidMapMarkerAlt';
  import PillBox from './PillBox.svelte';
  import { onMount } from 'svelte';

  export let propHistories;
  export let propItem = /** @type {TypeProperty | TypePropertyUS} */ ({});
  export let meta;
  export let isAdmin = false;
  export let isGuest = false;
  export let propExtraUS = /** @type {TypePropertyExtraUS} */ ({});

  let approvalStatus = 'approved';
  let signs = [
    {
      status: 'approved',
      sign: '***',
    }, {
      status: 'rejected',
      sign: '**',
    }, {
      status: 'pending',
      sign: '*',
    },
  ];

  let facilityInfoJSON;
  let mlsDisclaimerInfo;

  try {
    facilityInfoJSON = JSON.parse(String(propExtraUS.facilityInfo))
    mlsDisclaimerInfo = JSON.parse(String(propExtraUS.mlsDisclaimerInfo))
  } catch (error) {
    facilityInfoJSON = {};
    mlsDisclaimerInfo = {};
  }

  onMount(() => {
    console.log('onMount.Property');
    console.log('Property = ', propItem);
    console.log('Meta =', meta);
    console.log('IsAdmin =', isAdmin);
    console.log('isGuest =', isGuest);
    console.log('propItem.note =', propItem.note);
    if (propItem.approvalState !== 'pending' && propItem.approvalState !== '') {
      approvalStatus = 'rejected';
    }
    if (propItem.approvalState === 'pending') {
      approvalStatus = 'pending';
    }
  });
</script>

<div class='property'>
  <div class='property_main'>
      <div class='property_images'>
          {#if propItem.images && (propItem.images).length}
              <img src={String(propItem.images[0]).split(' ')[0]} alt='' />
          {:else}
              <div class='image_empty'>
                  <Icon size={40} color='#475569' src={FaSolidImage} />
                  <span>No Image !</span>
              </div>
          {/if}
      </div>
      <div class='property_info'>
          <div class='col1'>
              <div class='left'>
                  <div class={propItem.purpose === 'rent' ? `purpose label_rent` : `purpose label_sale`}>
                      {propItem.purpose === 'rent' ? 'For Rent' : 'On Sale'}
                  </div>
                  <div class='house_type'>
                      <Icon color='#FFF' size={16} src={FaSolidHome} />
                      <span>{propItem.houseType === '' ? 'House' : propItem.houseType}</span>
                  </div>
              </div>
              {#each signs as s}
                  {#if s.status === approvalStatus && s.status !== 'approved'}
                      <span class={`prop_status ${s.status}`}>{s.sign + ' Property is ' + s.status}</span>
                  {/if}
              {/each}
          </div>
          <div class='col2'>
              <h1>{formatPrice(propItem.lastPrice, 'TWD') || '0.00'}</h1>
              <p>Agency Fee : {propItem.agencyFeePercent}%</p>
              <div class='address'>
                  <Icon className='icon_address' color='#f97316' size={18} src={FaSolidMapMarkerAlt} />
                  <span>{propItem.formattedAddress === '' ? propItem.address : propItem.formattedAddress}</span>
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
              {#if propItem.countryCode === 'TW' || propItem.countryCode === ''}
                  <b>{M2ToPing(propItem.sizeM2 || 0)} Ping</b>
              {:else}
                  <b>{propItem.sizeM2 || '0'} M2</b>
              {/if}
              <p>Size</p>
          </div>
      </div>
  </div>
  <div class='property_attributes'>
    {#each meta as m}
      {#if !isAdmin && !(m.name === 'createdAt' || m.name === 'updatedAt' || m.name === 'countryCode')}
        {#if m.inputType === 'datetime'}
          <PillBox label={m.label} content={localeDatetime(propItem[m.name])} />
        {:else if m.name !== 'note'}
          <PillBox label={m.label} content={propItem[m.name]} />
          {#if facilityInfoJSON[m.name] != ``}
            <PillBox label={m.label} content={facilityInfoJSON[m.name]} />
          {/if}
          {#if mlsDisclaimerInfo[m.name] != ``}
            <PillBox label={m.label} content={mlsDisclaimerInfo[m.name]} />
          {/if}
        {/if}
      {:else if isAdmin}
        {#if m.inputType === 'datetime'}
          <PillBox label={m.label} content={localeDatetime(propItem[m.name])} />
        {:else if m.name !== 'note'}
          <PillBox label={m.label} content={propItem[m.name]} />
        {/if}
      {/if}
    {/each}
    {#if (!isGuest && propItem.note)}
      {@const note = JSON.parse(propItem.note)}
      <PillBox label='About' content={note.about} />
      <PillBox label='Seller/Broker Email' content={note.contactEmail} />
      <PillBox label='Seller/Broker Phone' content={note.contactPhone} />
    {/if}
  </div>
  {#if propItem.floorList && propItem.floorList.length}
      <div class='property_floors'>
          <h3>Floors</h3>
          <div class='floor_lists'>
              {#each propItem.floorList as floors}
                  <div class='floor_item'>
                      <div class='left'>
                          <h5>
                              {floors.type === 'basement' ? 'Basement' : `Floor #${floors.floor}`}
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
                          {#if floors.planImageUrl === ''}
                              <div class='image_empty'>
                                  <Icon size={30} color='#475569' src={FaSolidImage} />
                                  <span>No Image !</span>
                              </div>
                          {:else}
                              <img src={floors.planImageUrl} alt='' />
                          {/if}
                      </div>
                  </div>
              {/each}
          </div>
      </div>
  {/if}
  {#if propHistories && propHistories.length}
      <div class='property_history'>
          <h3>Price History</h3>
          <div class='history_list'>
              {#each propHistories as ph}
                  <div class='item'>
                      <span>{formatPrice(ph.price || ph.priceNtd, 'TWD')}</span>
                      <span>{localeDatetime(ph.updatedAt)}</span>
                  </div>
              {/each}
          </div>
      </div>
  {/if}
</div>

<style>
  .property {
      height           : fit-content;
      background-color : #F0F0F0;
      border-radius    : 8px;
      padding          : 15px 15px 70px 15px;
      display          : flex;
      flex-direction   : column;
      gap              : 20px;
      position         : relative;
  }

  .property h3 {
      font-size  : 22px;
      margin     : 0 0 15px 0;
      text-align : center;
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
      flex-direction   : column;
      gap              : 10px;
      font-size        : 18px;
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

  .property_main .property_info .col1 .prop_status {
      padding       : 5px 10px;
      border-radius : 5px;
      font-size     : 15px;
  }

  .property_main .property_info .col1 .approved {
      background-color : rgba(140, 216, 107, 1);
      color            : #FFF;
  }

  .property_main .property_info .col1 .pending {
      background-color : rgba(255, 208, 118, 1);
      color            : #475569;
  }

  .property_main .property_info .col1 .rejected {
      background-color : rgba(255, 126, 118, 1);
      color            : #FFF;
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

  .property_floors .floor_lists .floor_item .floor_plan_image .image_empty {
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

  .property_history .history_list {
      padding          : 15px;
      background-color : rgb(0 0 0 / 0.06);
      border-radius    : 8px;
      display          : flex;
      flex-direction   : column;
      min-height       : 200px;
  }

  .property_history .history_list .item {
      display         : flex;
      flex-direction  : row;
      justify-content : space-between;
      padding         : 10px 0;
      border-bottom   : 1px solid #CBD5E1;
      font-size       : 15px;
      font-weight     : 600;
  }
</style>