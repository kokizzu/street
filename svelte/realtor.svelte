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

  // <!-- {JSON.stringify( property )} -->
  // <a href='/realtor/property/{property.id}'>
  // <span>{property.formattedAddress}</span>

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
  }
  .property .property_main {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
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
</style>
