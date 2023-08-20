<script>
  // @ts-nocheck
  // import { onMount } from 'svelte';
  import Menu from './_components/Menu.svelte';
  import ProfileHeader from './_components/ProfileHeader.svelte';
  import Footer from './_components/Footer.svelte';
  // import { RealtorOwnedProperties } from 'jsApi.GEN';
  
  let ownedProperties = [/* ownedProperties */];
  
  // let ownedProperties = {};
  // onMount(async () => {
  //   await RealtorOwnedProperties(payload, function(res){
  //     console.log( res );
  //     if( res.error ) return alert( res.error );
  //     ownedProperties = res;
  //   })
  // });
  
  // TODO: Find out what payload to use 💀
  let user = {/* user */};
  let segments = {/* segments */};
</script>

<section class='dashboard'>
  <Menu
    access={segments}
  />
  <div class='dashboard_main_content'>
    <ProfileHeader></ProfileHeader>
    <div class='content'>
      <div class='property_lists_container'>
        <div class='property_lists_header'>
          <a href='/realtor/property' class='add_button' style=''>
            <i class='gg-add'></i>
            <span>Add</span>
          </a>
          <h1>Property owned</h1>
        </div>
        <section class='property_lists'>
          <!-- TODO, loop owned properties here -->
          <div class='property'>
            <ul>
              {#each ownedProperties as property}
                <li>
                  {JSON.stringify( property )}
                  <a href='/realtor/property/{property.id}'>
                    <span>{property.formattedAddress}</span>
                </li>
              {/each}
            </ul>
          </div>
          
          <div class='property'>
          
          </div>
        </section>
      </div>
    </div>
    <Footer></Footer>
  </div>
</section>

<style>
  .property_lists_container {
    position         : relative;
    margin-top       : -40px;
    margin-left      : auto;
    margin-right     : auto;
    padding          : 20px;
    display          : flex;
    flex-direction   : column;
    gap              : 20px;
    background-color : white;
    border-radius    : 8px;
    min-height       : 500px;
    width            : 70%;
    color            : #475569;
  }

  .property_lists_container .property_lists_header {
    display         : flex;
    flex-direction  : row;
    justify-content : space-between;
    align-items     : center;
    height          : fit-content;
    width           : 100%;
  }

  .property_lists_container .property_lists_header .add_button {
    padding          : 8px 20px;
    font-size        : 16px;
    display          : inline-flex;
    flex-direction   : row;
    align-items      : center;
    align-content    : center;
    justify-content  : center;
    border           : none;
    background-color : #6366F1;
    border-radius    : 8px;
    color            : white;
    cursor           : pointer;
    text-decoration  : none;
  }

  .property_lists_container .property_lists_header h1 {
    font-size : 18px;
    margin    : 0;
  }

  .property_lists_container .property_lists_header .add_button:hover {
    background-color : #7E80F1;
  }

  .property_lists_container .property_lists_header .add_button i {
    margin-right : 8px;
  }

  .property_lists_container .property_lists {
    display        : flex;
    flex-direction : column;
    gap            : 20px;
    width          : 100%;
  }

  .property_lists_container .property_lists .property {
    background-color : #F0F0F0;
    border-radius    : 8px;
    height           : 250px;
  }
</style>