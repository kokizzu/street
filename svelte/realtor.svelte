<script>
  // @ts-nocheck
  import { onMount } from 'svelte';
  import Menu from './_components/Menu.svelte';
  import ProfileHeader from './_components/ProfileHeader.svelte';
  import Footer from './_components/Footer.svelte';
  import OwnedProperty from '_components/OwnedProperty.svelte';

  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidPlusCircle from "svelte-icons-pack/fa/FaSolidPlusCircle";

  let ownedProperties = [/* ownedProperties */];
  let pager = [/* pager */];

  onMount(() => {
    console.log('ownedProperties=',ownedProperties);
    console.log('pager=',pager);
  });

  

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
          <h1>Owned Properties: {pager.total}</h1>
          <a href="/realtor/property" class="add_button" style="">
            <Icon size={20} color="#FFFF" src={FaSolidPlusCircle} />
            <span>Add</span>
          </a>
        </div>
        {#if ownedProperties.length}
          <section class="property_lists">
            {#each ownedProperties as property}
              <OwnedProperty property={property} />
            {/each}
          </section>
        {:else}
          <div class="no_property">
            <p>No Property</p>
          </div>
        {/if}
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
    gap: 8px;
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

  .property_lists_container .no_property {
    background-color: #f0f0f0;
    border-radius: 8px;
    height: fit-content;
    flex-grow: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 15px;
    font-weight: 600;
  }

  .property_lists_container .property_lists {
    display: flex;
    flex-direction: column;
    gap: 20px;
    width: 100%;
  }

  
</style>
