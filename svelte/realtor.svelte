<script>
  /** @typedef {import('./_types/master.js').Access} Access */
  /** @typedef {import('./_types/user').User} User */
  /** @typedef {import('./_types/property').PropertyWithNote} PropertyWithNote */
  /** @typedef {import('./_types/master').PagerIn} PagerIn */
  /** @typedef {import('./_types/master').PagerOut} PagerOut */
  /** @typedef {import('./_types/master').Field} Field */

  import Main from './_layouts/Main.svelte';
  import OwnedProperty from './_components/OwnedProperty.svelte';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddLargeFill } from './node_modules/svelte-icons-pack/dist/ri';
  
  const user              = /** @type {User} */ ({/* user */});
  const access            = /** @type {Access} */ ({/* segments */});
  let ownedProperties     = /** @type {PropertyWithNote[]} */ [/* ownedProperties */];
  let pager               = /** @type {PagerOut} */ ({/* pager */});
  let propertyMeta        = /** @type {Field[]} */ ([/* propertyMeta */]);
</script>

<Main {user} {access}>
  <div class="properties-container">
    <div class="info-actions">
      <h2>Owned/Managed Properties: {pager.countResult}</h2>
      <a href="/realtor/property" class="add-btn">
        <Icon
          size="20"
          color="#FFF"
          src={RiSystemAddLargeFill}
        />
        <span>Add</span>
      </a>
    </div>
    {#if ownedProperties && ownedProperties.length}
      <section class="properties">
        {#each ownedProperties as property}
          {#if property.deletedAt<=0}
            <OwnedProperty {property} meta={propertyMeta} />
          {/if}
        {/each}
      </section>
    {:else}
      <div class="no-property">
        <p>No Property</p>
      </div>
    {/if}
  </div>
</Main>

<style>
  .properties-container {
    display: flex;
    flex-direction: column;
    width: 100%;
    gap: 20px;
    padding: 20px;
  }

  .properties-container .info-actions {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    gap: 20px;
    width: 100%;
  }

  .properties-container .info-actions .add-btn {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
    background-color: var(--orange-006);
    border-radius: 9999px;
    padding: 15px 30px;
    color: #FFF;
    font-size: 16px;
    font-weight: 600;
    text-decoration: none;
  }

  .properties-container .info-actions .add-btn:hover {
    background-color: var(--orange-005);
  }

  .properties-container .info-actions h2 {
    font-size: 18px;
    margin: 0;
  }

  .properties-container .no-property {
    background-color: var(--gray-002);
    border-radius: 8px;
    height: fit-content;
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight      : 600;
  }

  .properties-container .properties {
    display: flex;
    flex-direction: column;
    gap: 20px;
    width: 100%;
  }

  /* Responsive to mobile device */
  @media (max-width : 768px) {
    :global(.icon) {
      width  : 15px;
      height : 15px;
    }

    .properties-container {
      margin  : -40px 20px 0 20px;
      padding : 15px;
      width   : auto;
    }

    .properties-container .info-actions h2 {
      font-size : 16px !important;
    }

    .properties-container .info-actions .add-btn {
      padding   : 5px 15px !important;
      font-size : 14px !important;
      gap       : 8px !important;
    }
  }
</style>