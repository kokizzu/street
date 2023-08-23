<script>
  // @ts-nocheck
  import { onMount } from 'svelte';
  import Menu from './_components/Menu.svelte';
  import ProfileHeader from './_components/ProfileHeader.svelte';
  import Footer from './_components/Footer.svelte';

  let ownedProperties = [/* ownedProperties */];

  onMount(() => {
    console.log(ownedProperties);
  });

  function datetime( unixSec ) {
    if( !unixSec ) return '';
    const dt = new Date( unixSec * 1000 )
    const day = dt.toLocaleDateString('default', {weekday: 'long'});
    const date = dt.getDay();
    const month = dt.toLocaleDateString('default', { month: 'long'});
    const year = dt.getFullYear()
    const formattedDate = `${day}, ${date} ${month} ${year}`
    return formattedDate
  }

  let user = {/* user */};
  let segments = {/* segments */};
</script>

<section class="dashboard">
  <Menu access={segments} />
  <div class="dashboard_main_content">
    <ProfileHeader />
    <div class="content">
      <div class="property_lists_container">
        <div class="property_lists_header">
          <a href="/realtor/property" class="add_button" style="">
            <i class="gg-add" />
            <span>Add</span>
          </a>
          <h1>Property owned</h1>
        </div>
        <section class="property_lists">
          {#each ownedProperties as property}
            <div class="property">
              <div class="property_main">
                <div class="property_images">
                  <!-- TODO: make it as slider, there are more images to display -->
                  {#if property.images.length}
                    <img src={property.images[0]} alt="" />
                  {:else}
                    <div class="image_empty">
                      <i class="gg-image" />
                    </div>
                  {/if}
                </div>
                <div class="property_info">
                  <div class="col1">
                    <div class="left">
                      <div class="purpose">On {property.purpose}</div>
                      <div class="house_type">
                        <i class="gg-home-alt"></i>
                        <span>{property.houseType}</span>
                      </div>
                    </div>
                    <a href='/realtor/property/{property.id}' class="edit_property">
                      <i class="gg-pen"></i>
                      <span>Edit Property</span>
                    </a>
                  </div>

                  <div class="col2">
                    <h1>$ {property.lastPrice || '0.00'}</h1>
                    <p>Agency Fee : {property.agencyFeePercent}%</p>
                    <div class='address'>
                      <i class='gg-pin'/>
                      <span>{property.formattedAddress}</span>
                    </div>
                  </div>
                </div>
              </div>
              <div class="property_secondary">
                <div class="feature_number">
                  <div class="feature_item">
                    <b>{property.numberOfFloors || '0'}</b>
                    <p>FLoors</p>
                  </div>
                  <div class="feature_item">
                    <b>{property.bathroom || '0'}</b>
                    <p>Bathroom</p>
                  </div>
                  <div class="feature_item">
                    <b>{property.bedroom || '0'}</b>
                    <p>Bedroom</p>
                  </div>
                  <div class="feature_item">
                    <b>{property.sizeM2 || '0'} M2</b>
                    <p>Size</p>
                  </div>
                </div>
              </div>
              <table class="property_table">
                <tbody>
                  <tr>
                    <td class="name">Property ID</td>
                    <td class="value">{property.id || '-'}</td>
                  </tr>
                  <tr>
                    <td class="name">Created At</td>
                    <td class="value">{datetime(property.createdAt) || '0'}</td>
                  </tr>
                  <tr>
                    <td class="name">Updated At</td>
                    <td class="value">{datetime(property.updatedAt) || '0'}</td>
                  </tr>
                  <tr>
                    <td class="name">Building Lamination</td>
                    <td class="value">{property.buildingLamination || '-'}</td>
                  </tr>
                  <tr>
                    <td class="name">Construct Completed Date</td>
                    <td class="value">{datetime(property.constructCompletedDate) || '-'}</td>
                  </tr>
                  <tr>
                    <td class="name">Updated At</td>
                    <td class="value">{datetime(property.updatedAt) || '0'}</td>
                  </tr>
                  <tr>
                    <td class="name">Main Use</td>
                    <td class="value">{property.mainUse || '-'}</td>
                  </tr>
                  <tr>
                    <td class="name">Note</td>
                    <td class="value">{property.note || '-'}</td>
                  </tr>
                </tbody>
              </table>
              <div class="property_floors">
                <h3>FLOORS</h3>
                {#if property.floorList.length}
                  <div class="floor_lists">
                  {#each property.floorList as floors}
                    <div class="floor_item">
                      <div class="left">
                        <h5>
                          {floors.type==='basement' ? floors.type : `${floors.type} #${floors.floor}`}
                        </h5>
                        <!-- TODO: currently room list only 1 object, fix Tarantool to accept array -->
                        {#if floors.rooms}
                          <div class="room_lists">
                            <div class="room_item">
                              <span>{@html floors.rooms.name || '-'}</span>
                              <span>{@html floors.rooms.sizeM2 || '-'} M2</span>
                            </div>
                            <div class="room_item">
                              <span>{@html floors.rooms.name || '-'}</span>
                              <span>{@html floors.rooms.sizeM2 || '-'} M2</span>
                            </div>
                            <div class="room_item">
                              <span>{@html floors.rooms.name || '-'}</span>
                              <span>{@html floors.rooms.sizeM2 || '-'} M2</span>
                            </div>
                            <div class="room_item">
                              <span>{@html floors.rooms.name || '-'}</span>
                              <span>{@html floors.rooms.sizeM2 || '-'} M2</span>
                            </div>
                            <div class="room_item">
                              <span>{@html floors.rooms.name || '-'}</span>
                              <span>{@html floors.rooms.sizeM2 || '-'} M2</span>
                            </div>
                            <div class="room_item">
                              <span>{@html floors.rooms.name || '-'}</span>
                              <span>{@html floors.rooms.sizeM2 || '-'} M2</span>
                            </div>
                            <div class="room_item">
                              <span>{@html floors.rooms.name || '-'}</span>
                              <span>{@html floors.rooms.sizeM2 || '-'} M2</span>
                            </div>
                          </div>
                        {/if}
                      </div>

                      <div class="floor_plan_image">
                        {#if floors.planImageUrl === ''}
                          <span>
                            <i class='gg-image'></i>
                            <p>No Image</p>
                          </span>
                        {:else}
                          <img src={floors.planImageUrl} alt="" />
                        {/if}
                      </div>
                    </div>
                  {/each}
                  </div>
                {:else}
                  <div class="no_floors">
                    <p>No Floors</p>
                  </div>
                {/if}
              </div>
            </div>
          {/each}
        </section>
      </div>
    </div>
    <Footer />
  </div>
</section>

<style>
  .property_lists_container {
    position: relative;
    margin-top: -40px;
    margin-left: auto;
    margin-right: auto;
    padding: 20px;
    display: flex;
    flex-direction: column;
    gap: 20px;
    background-color: white;
    border-radius: 8px;
    min-height: 500px;
    width: 70%;
    color: #475569;
  }

  .property_lists_container .property_lists_header {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    height: fit-content;
    width: 100%;
  }

  .property_lists_container .property_lists_header .add_button {
    padding: 8px 20px;
    font-size: 16px;
    display: inline-flex;
    flex-direction: row;
    align-items: center;
    align-content: center;
    justify-content: center;
    border: none;
    background-color: #6366f1;
    border-radius: 8px;
    color: white;
    cursor: pointer;
    text-decoration: none;
  }

  .property_lists_container .property_lists_header h1 {
    font-size: 18px;
    margin: 0;
  }

  .property_lists_container .property_lists_header .add_button:hover {
    background-color: #7e80f1;
  }

  .property_lists_container .property_lists_header .add_button i {
    margin-right: 8px;
  }

  .property_lists_container .property_lists {
    display: flex;
    flex-direction: column;
    gap: 20px;
    width: 100%;
  }

  .property_lists_container .property_lists .property {
    background-color: #f0f0f0;
    border-radius: 8px;
    height: fit-content;
    padding: 15px;
    display: flex;
    flex-direction: column;
    gap: 20px;
  }
  .property .property_main {
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    align-content: flex-start;
    justify-content: flex-start;
    align-items: flex-start;
    gap: 15px;
  }
  .property .property_main .property_images {
    width: 260px;
    height: 160px;
    flex: none;
    overflow: hidden;
    border-radius: 8px;
  }
  .property .property_main .property_images img {
    object-fit: cover;
    width: 100%;
    height: 100%;
    
  }
  .property .property_main .property_images .image_empty {
    border-radius    : 8px;
    object-fit       : cover;
    width            : 100%;
    height           : 100%;
    background-color : rgb(0 0 0 / 0.06);
    display          : flex;
    justify-content  : center;
    align-items      : center;
  }
  
  .property .property_main .property_info {
    flex-grow: 1;
    display: flex;
    flex-direction: column;
    gap: 10px;
  }
  .property .property_main .property_info .col1 {
    display: flex;
    justify-content: space-between;
    align-items: center;
    flex-direction: row;
  }
  .property .property_main .property_info .col1 .left {
    display: flex;
    gap: 15px;
    align-items: center;
  }
  .property .property_main .property_info .col1 .left .purpose,
  .property .property_main .property_info .col1 .left .house_type,
  .property .property_main .property_info .col1 .edit_property {
    background-color: #f97316;
    padding: 10px 25px;
    border-radius: 8px;
    color: white;
    font-size: 14px;
    text-transform: capitalize;
    text-decoration: none;
  }
  .property .property_main .property_info .col1 .left .house_type,
  .property .property_main .property_info .col1 .edit_property {
    display: flex;
    gap: 10px;
    align-items: center;
  }
  .property .property_main .property_info .col1 .edit_property:hover {
    background-color: #F58433;
  }

  .property .property_main .property_info .col2 {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }
  .property .property_main .property_info .col2 h1 {
    margin: 0;
    font-size: 32px;
  }
  .property .property_main .property_info .col2 p {
    margin: 0 0 10px 0;
  }
  .property .property_main .property_info .col2 .address {
    display: flex;
    flex-direction: row;
    gap: 10px;
    align-items: center;
  }
  .property .property_main .property_info .col2 .address i {
    color: #f97316;
  }

  .property .property_secondary {
    display: flex;
    flex-direction: column;
    gap: 15px;
  }
  .property .property_secondary .feature_number {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-evenly;
    width: 100%;
  }
  .property .property_secondary .feature_number .feature_item {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 8px;
  }
  .property .property_secondary p {
    margin: 0;
  }
  .property .property_secondary b {
    font-size: 22px;
  }
  
  .property .property_table {
    border-collapse: collapse;
    font-size: 14px;
    width: 100%;
    color: #475569;
    border: 1px solid #cbd5e1;
  }
  .property .property_table .name {
    width: 210px;
    max-width: 210px;
    font-weight: 700;
  }
  .property .property_table .name,
  .property .property_table .value {
    border-top: 1px solid #cbd5e1;
    border-bottom: 1px solid #cbd5e1;
    padding: 12px;
  }

  .property .property_floors {
    width: 70%;
    margin: 0 auto;
  }

  .property .property_floors h3 {
    font-size: 22px;
    margin: 0 0 15px 0;
    text-align: center;
  }
  .property .property_floors .floor_lists {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }
  .property .property_floors .floor_lists .floor_item {
    display: flex;
    align-items: flex-start;
    gap: 35px;
  }
  .property .property_floors .floor_lists .floor_item .left {
    flex-grow: 1;
  }
  .property .property_floors .floor_lists .floor_item .left h5 {
    font-size: 22px;
    text-transform: capitalize;
    margin: 0 0 15px 0;
    width: 100%;
  }
  .property .property_floors .floor_lists .floor_lists .floor_item .left .room_lists {
    font-size: 14px;
    width: 100%;
    color: #475569;
    border: 1px solid #cbd5e1;
  }
  .property .property_floors .floor_lists .floor_item .left .room_lists .room_item {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    text-transform: capitalize;
    border-bottom: 1px solid #cbd5e1;
    padding: 5px 0;
  }
  .property .property_floors .floor_lists .floor_item .floor_plan_image {
    width: 240px;
    height: 140px;
    position      : relative;
    border-radius : 8px;
    overflow      : hidden;
    border: 1px solid #cbd5e1;
    margin-top: 45px;
  }
  .property .property_floors .floor_lists .floor_item .floor_plan_image img {
    width      : 100%;
    height     : 100%;
    object-fit : cover;
  }
  .property .property_floors .floor_lists .floor_item .floor_plan_image span {
    border-radius    : 8px;
    object-fit       : cover;
    width            : 100%;
    height           : 100%;
    display          : flex;
    flex-direction: column;
    gap: 8px;
    justify-content  : center;
    align-items      : center;
  }
  .property .property_floors .floor_lists .floor_item .floor_plan_image span p {
    margin: 0;
  }
</style>
