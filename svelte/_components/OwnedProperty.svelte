<script>
  // @ts-nocheck
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidHome from 'svelte-icons-pack/fa/FaSolidHome';
  import FaSolidPen from 'svelte-icons-pack/fa/FaSolidPen';
  import FaSolidMapMarkerAlt from 'svelte-icons-pack/fa/FaSolidMapMarkerAlt';
  import PillBox from './PillBox.svelte';
  import { localeDatetime } from './formatter';
  
  export let property;
  export let meta;
  let showMore = false;
  
  function handleShowMore() {
    showMore = !showMore;
  }

</script>

<div class="{showMore ? 'property_more' : 'property'}">
  <div class='property_main'>
    <div class='property_images'>
      <!-- TODO: make it as slider, there are more images to display -->
      {#if property.images && property.images.length}
        <img src={property.images[0]} alt='' />
      {:else}
        <div class='image_empty'>
          <i class='gg-image' />
        </div>
      {/if}
    </div>
    <div class='property_info'>
      <div class='col1'>
        <div class='left'>
          <div class={property.purpose === 'sell' ? `purpose label_sale` : `purpose label_rent`}>{property.purpose === 'sell' ? `On Sale` : `For ${property.purpose}`}</div>
          <div class='house_type'>
            <Icon size={16} color='#FFFF' src={FaSolidHome} />
            <span>{property.houseType}</span>
          </div>
        </div>
        <a href='/realtor/property/{property.id}' class='edit_property'>
          <Icon size={13} color='#FFFF' src={FaSolidPen} />
          <span>Edit Property</span>
        </a>
      </div>
      
      <div class='col2'>
        <h1>$ {property.lastPrice || '0.00'}</h1>
        <p>Agency Fee : {property.agencyFeePercent}%</p>
        <div class='address'>
          <Icon size={18} color='#f97316' src={FaSolidMapMarkerAlt} />
          <span>{property.formattedAddress}</span>
        </div>
      </div>
    </div>
  </div>
  <div class='property_secondary'>
    <div class='feature_number'>
      <div class='feature_item'>
        <b>{property.numberOfFloors || '0'}</b>
        <p>Floors</p>
      </div>
      <div class='feature_item'>
        <b>{property.bathroom || '0'}</b>
        <p>Bathroom</p>
      </div>
      <div class='feature_item'>
        <b>{property.bedroom || '0'}</b>
        <p>Bedroom</p>
      </div>
      <div class='feature_item'>
        <b>{property.sizeM2 || '0'} M2</b>
        <p>Size</p>
      </div>
    </div>
  </div>
  <div class='property_attributes'>
    {#each meta as m}
      {#if property[ m.name ]}
        {#if m.inputType==='datetime'}
          <PillBox label={m.label} content={localeDatetime(property[m.name])} />
        {:else}
          <PillBox label={m.label} content={property[m.name]} />
        {/if}
      {/if}
    {/each}
  </div>
  <div class='property_floors'>
    <h3>FLOORS</h3>
    {#if property.floorList && property.floorList.length}
      <div class='floor_lists'>
        {#each property.floorList as floors}
          <div class='floor_item'>
            <div class='left'>
              <h5>
                {floors.type==='basement' ? floors.type : `${floors.type} #${floors.floor}`}
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
                                <i class='gg-image'></i>
                                <p>No Image</p>
                              </span>
              {:else}
                <img src={floors.planImageUrl} alt='' />
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
  <div class='property_less_more'>
    <button on:click={handleShowMore} class='toggle_show_more'>
      Show {showMore===true ? 'Less' : 'More'}
    </button>
  </div>
</div>

<style>
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
</style>