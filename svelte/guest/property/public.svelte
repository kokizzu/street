<script>
  /** @typedef {import('../../_types/property').Property} TypeProperty */
  /** @typedef {import('../../_types/property').PropertyUS} TypePropertyUS */
  /** @typedef {import('../../_types/property').PropertyExtraUS} TypePropertyExtraUS */
  /** @typedef {import('../../_types/master').Field} Field */

  import { Icon } from '../../node_modules/svelte-icons-pack/dist';
  import {
    FaSolidArrowRight, FaSolidShareNodes, FaCopy, FaBrandsFacebook,
    FaBrandsLinkedin, FaBrandsTwitter, FaBrandsWhatsapp,
    FaBrandsTelegram,
  } from '../../node_modules/svelte-icons-pack/dist/fa';
  import Property from '../../_components/Property.svelte';
  import {notifier} from '../../_components/notifier.js';
  
  let propItem    = /** @type {TypeProperty | TypePropertyUS} */ ({/* propItem */});
  let meta        = /** @type {Field[]}*/ ([/* propertyMeta */]);
  let propExtraUS = /** @type {TypePropertyExtraUS} */ ({/* propExtraUS */});
  let isAdmin     = /** @type {boolean} */ (false);

  function copyToClipboard(/** @type {string} */ text ) {
    navigator.clipboard.writeText( text );
    notifier.showInfo( 'Link copied to clipboard' );
  }
</script>

<section class='property_container'>
  {#if propItem.deletedAt > 0}
    <h2>This property has been deleted !!</h2>
  {:else}
    <div class='property'>
      <Property
        propItem={propItem}
        meta={meta}
        isAdmin={isAdmin}
        isGuest={true}
        propExtraUS={propExtraUS}
        propHistories={[]}
      />
    </div>
    <div class='side_attribute'>
      <div class='login_container'>
        <div>
          <h3>Login to</h3>
          <h3>HapSTR</h3>
        </div>
        <a class='login_btn' href='/'>
          <span>Login</span>
          <Icon color='#FFF' size="13" src={FaSolidArrowRight} />
        </a>
      </div>
      <div class='share_container'>
        <header>
          <span>Share this</span>
          <Icon className='share_icon' color='#9fa9b5' size="14" src={FaSolidShareNodes} />
        </header>
        <div class='share_options'>
          <button class='share_item' on:click={() => copyToClipboard(String(window.location))} title='Copy link address'>
            <Icon className='share_icon' color='#475569' size="28" src={FaCopy} />
          </button>
          <a aria-label='Share to Facebook'
             class='share_item'
             href='https://www.facebook.com/sharer/sharer.php?u={window.location}?utm_source=facebook&utm_medium=social&utm_campaign=user-share'
             rel='noopener'
             target='_blank'
          >
            <Icon className='share_icon' color='#475569' size="30" src={FaBrandsFacebook} />
          </a>
          <a aria-label='Share to LinkedIn'
             class='share_item'
             href='https://www.linkedin.com/shareArticle?mini=true&url={window.location}&title=I%20Found%20Awesome%House%20{window.location}'
             rel='noopener'
             target='_blank'
          >
            <Icon className='share_icon' color='#475569' size="30" src={FaBrandsLinkedin} />
          </a>
          <a aria-label='Share to Twitter'
             class='share_item'
             href='https://twitter.com/intent/tweet?url={window.location}'
             rel='noopener'
             target='_blank'
          >
            <Icon className='share_icon' color='#475569' size="30" src={FaBrandsTwitter} />
          </a>
          <a aria-label='Share to Telegram'
             class='share_item'
             href='https://t.me/share/url?url={window.location}'
             rel='noopener'
             target='_blank'
          >
            <Icon className='share_icon' color='#475569' size="30" src={FaBrandsTelegram} />
          </a>
          <a aria-label='Share to WhatsApp'
             class='share_item'
             href='https://api.whatsapp.com/send?text=I%20Found%20Awesome%20House%20{window.location}'
             rel='noopener'
             target='_blank'
          >
            <Icon className='share_icon' color='#475569' size="33" src={FaBrandsWhatsapp} />
          </a>
        </div>
      </div>
    </div>
  {/if}
</section>

<style>
  .property_container {
    width           : 100%;
    margin          : 30px auto 50px auto;
    color           : #475569;
    display         : flex;
    flex-direction  : row;
    justify-content : center;
    gap             : 30px;
  }

  .property {
    width : 60%;
  }

  .side_attribute {
    width          : 250px;
    height         : fit-content;
    display        : flex;
    flex-direction : column;
    gap            : 20px;
  }

  .side_attribute .login_container,
  .side_attribute .share_container {
    text-align       : center;
    height           : fit-content;
    background-color : #F0F0F0;
    border-radius    : 8px;
    padding          : 15px;
    display          : flex;
    flex-direction   : column;
    gap              : 10px;
  }

  .side_attribute .login_container h3 {
    margin    : 0 0 10px 0;
    font-size : 24px;
  }

  .side_attribute .login_container .login_btn {
    background-color : #6366F1;
    width            : 100%;
    height           : fit-content;
    padding          : 12px;
    text-align       : center;
    border-radius    : 8px;
    color            : #FFF;
    font-size        : 13px;
    display          : flex;
    flex-direction   : row;
    gap              : 6px;
    justify-content  : center;
    align-items      : center;
    text-decoration  : none;
  }

  .side_attribute .login_container .login_btn:hover {
    background-color : #7E80F1;
  }

  .side_attribute .share_container header {
    display         : flex;
    flex-direction  : row;
    justify-content : center;
    gap             : 7px;
    align-items     : center;
    font-size       : 15px;
  }

  .side_attribute .share_container .share_options {
    display               : grid;
    grid-template-columns : repeat(4, minmax(0, 1fr));
    grid-template-rows    : repeat(2, minmax(0, 1fr));
    align-items           : center;
    justify-items         : center;
    justify-content       : center;
    align-content         : center;
    row-gap               : 12px;
  }

  .side_attribute .share_container .share_options .share_item {
    border     : none;
    background : none;
    cursor     : pointer;
  }

  :global(.side_attribute .share_container .share_options .share_item:hover .share_icon) {
    fill : #57667A !important;
  }
</style>