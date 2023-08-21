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
                  <img src={property.images[0]} alt="" />
                </div>

                <div class="property_info">

                  <div class="col1">
                    <div class="left">
                      <div>{property.purpose}</div>
                      <div class="house_type">
                        <i class='gg-home-alt'></i>
                        <span>{property.houseType}</span>
                      </div>
                    </div>
                    <a href='/realtor/property/{property.id}' class="edit_property">
                      <i class="gg-pen"></i>
                      <span>Edit Property</span>
                    </a>
                  </div>

                  <div class="col2">
                    <h1>$ {property.lastPrice}</h1>
                    <p>{property.agencyFeePercent}</p>
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
  .property .property_images {
    width: 260px;
    height: 160px;
    flex: none;
  }
  .property .property_images img {
    object-fit: cover;
    width: 100%;
    height: 100%;
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
  .left .house_type, .edit_property {
    display: flex;
    gap: 8px;
    align-items: center;
   }
</style>
