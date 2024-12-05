<script>
  /** @typedef {import('../../_types/user').User} User */
  /** @typedef {import('../../_types/master').Access} Access */
  /** @typedef {import('../../_types/property').PropertyWithNote} PropertyWithNote */

  import Main from '../../_layouts/Main.svelte';
  import { Icon } from '../../node_modules/svelte-icons-pack/dist';
  import { RiArrowsArrowLeftSLine, RiDocumentFileDownloadLine } from '../../node_modules/svelte-icons-pack/dist/ri'
  import GoogleMapJs from '../../_components/GoogleMap/GoogleMapJS.svelte';
  import { formatPrice } from '../../_components/formatter';
  import PopUpUpload3DFile from '../../_components/PopUpUpload3DFile.svelte';
  import axios, { HttpStatusCode } from 'axios';
  import { notifier } from '../../_components/notifier';
  
  const title     = /** @type {string} */ ('#{title}')
  const user      = /** @type {User} */ ({/* user */});
  const access    = /** @type {Access} */ ({/* segments */});
  const property  = /** @type {PropertyWithNote} */ ({/* property */});

  console.log('Property=',property)

  let gmapComponent     = /** @type {import('svelte').SvelteComponent} */ (null);
  let popUpUpload3DFile = /** @type {import('svelte').SvelteComponent} */ (null);

  const defImgUrl = /** @type {string} */ ('/assets/img/placeholder.webp');
  let imgUrl = /** @type {string} */ (property.images[0] || defImgUrl);
  let img3dUrl = /** @type {string} */ (property.image3dUrl);

  function gmapReady() {
    gmapComponent.SetMarker(
      property.coord[0],
      property.coord[1],
      title
    );
  }

  let isSubmitUpload3dFile = false;
  let uploadingProgressStr = '';

  async function SubmitUpload3DFile(/** @type {File} */ file) {
    isSubmitUpload3dFile = true;

    const formData = new FormData();
    formData.append('propertyId', String(property.id));
    formData.append('country', String(property.countryCode));
    formData.append('rawFile', file, file.name);

    await axios.postForm('/user/upload3DFile', formData, {
      headers: {
        "Content-Type": "multipart/form-data"
      },
      onUploadProgress: (/** @type {import('axios').AxiosProgressEvent} */ event) => {
        const propgressNum = Math.round((event.loaded * 100) / Number(event.total));
        uploadingProgressStr = `Uploading ${propgressNum}%`;
      }
    }).then((/** @type {import('axios').AxiosResponse}*/ res) => {
      notifier.showSuccess('3D file uploaded !!');
      popUpUpload3DFile.Close();

      img3dUrl = res.data.imageUrl;
    }).catch((/** @type {import('axios').AxiosError}*/ err) => {
      const res = /** @type {import('axios').AxiosResponse} */ (err.response);
      switch (res.status) {
        case HttpStatusCode.PayloadTooLarge:
          notifier.showError('File size too large');
          break;
        default:
          notifier.showError(res.data.error || 'Failed to upload 3D file');
      }
      popUpUpload3DFile.Reset();
    });
  }
</script>

<PopUpUpload3DFile
  bind:this={popUpUpload3DFile}
  bind:isSubmitted={isSubmitUpload3dFile}
  bind:uploadingProgressStr
  OnSubmit={SubmitUpload3DFile}
/>

<Main {user} {access}>
  <div class="listing-container">
    <div class="content">
      <div class="property">
        <a class="go-back" href="/user/listings">
          <Icon
            src={RiArrowsArrowLeftSLine}
            size="20"
            color="var(--orange-006)"
          />
          <span>Back</span>
        </a>
        <picture class="img-container">
          <img
            src={imgUrl}
            alt={title}
            on:error={() => imgUrl = defImgUrl}
          />
        </picture>
        <div class="info-1">
          <h2 class="address">
            {property.address || property.formattedAddress}
          </h2>
          <p class="about">{property.about || '--'}</p>
          {#if img3dUrl === ''}
            <button class="upload-btn" on:click={() => popUpUpload3DFile.Show()}>
              Upload 3D File
            </button>
          {/if}
          {#if img3dUrl !== ''}
            <a class="download-btn"
              href="/user/download3dFile?country={property.countryCode}&propertyId={property.id}"
              target="_blank"
            >
              <Icon
                src={RiDocumentFileDownloadLine}
                size="20"
                color="#FFF"
              />
              <span>Download 3D File</span>
            </a>
          {/if}
          <table>
            <tbody>
              <tr>
                <th><span>Property Key</span></th>
                <td>{property.uniqPropKey}</td>
              </tr>
              <tr>
                <th><span>Price</span></th>
                <td>{formatPrice((Number(property.lastPrice || 0)), 'USD')}</td>
              </tr>
              <tr>
                <th><span>Email</span></th>
                <td>{property.contactEmail || '--@--'}</td>
              </tr>
              <tr>
                <th><span>Phone</span></th>
                <td>{property.contactPhone || '0 000 000 0000'}</td>
              </tr>
              <tr>
                <th><span>Purpose</span></th>
                <td>{property.purpose || 'For Sale'}</td>
              </tr>
              <tr>
                <th><span>Bathroom</span></th>
                <td>{property.bathroom || 0}</td>
              </tr>
              <tr>
                <th><span>Bedroom</span></th>
                <td>{property.bedroom || 0}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <div class="map-full">
        <div class="map-container">
          <GoogleMapJs
            bind:this={gmapComponent}
            lat={property.coord[0]}
            lng={property.coord[1]}
            zoom={10}
            on:ready={gmapReady}
          />
        </div>
      </div>
    </div>
  </div>
</Main>

<style>
  .listing-container {
    display: flex;
    flex-direction: column;
    width: 100%;
  }

  .listing-container div.content {
    display: grid;
    grid-template-columns: auto 55%;
    position: relative;
  }

  .listing-container .content .property {
    display: flex;
    flex-direction: column;
    gap: 10px;
    padding: 20px;
  }

  .listing-container .content .property a.go-back {
    display: flex;
    flex-direction: row;
    gap: 5px;
    align-items: center;
    padding: 5px 10px 5px 7px;
    color: var(--orange-006);
    text-decoration: none;
    width: fit-content;
    border-radius: 5px;
    font-size: 15px;
    font-weight: 600;
  }

  .listing-container .content .property a.go-back:hover {
    background-color: var(--orange-transparent);
  }

  .listing-container .content .property picture.img-container {
    width: 100%;
    height: 200px;
    overflow: hidden;
    border-radius: 8px;
    cursor: pointer;
  }

  .listing-container .content .property picture.img-container img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition-duration : 200ms;
  }

  .listing-container .content .property picture.img-container img:hover {
    transform: scale(1.2);
  }

  .listing-container .content .property .info-1 {
    display: flex;
    flex-direction: column;
    gap: 15px;
  }

  .listing-container .content .property .info-1 .address {
    margin: 10px 0 0 0;
  }

  .listing-container .content .property .info-1 .about {
    margin: 0;
    padding: 0;
  }

  .listing-container .content .property .info-1 button.upload-btn {
    background-color: transparent;
    border: none;
    font-weight: 600;
    color: var(--orange-006);
    cursor: pointer;
    width: fit-content;
    padding: 0;
  }

  .listing-container .content .property .info-1 button.upload-btn:hover {
    color: var(--orange-005);
  }

  .listing-container .content .property .info-1 a.download-btn {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    gap: 8px;
    padding: 8px 13px;
    border-radius: 5px;
    background-color: var(--blue-006);
    color: #FFF;
    width: fit-content;
    text-decoration: none;
  }

  .listing-container .content .property .info-1 a.download-btn:hover {
    background-color: var(--blue-005);
  }

  .listing-container .content .property .info-1 table {
    width: 100%;
    background: #fff;
    box-shadow: none;
    text-align: left;
    border-collapse: separate;
    border-spacing: 0;
    overflow: hidden;
		font-size: 13px;
  }

  .listing-container .content .property .info-1 table tbody th {
    padding: 3px 5px 0 0;
    min-width: 100px;
    width: 130px;
  }

  .listing-container .content .property .info-1 table tbody td {
    padding: 3px 0;
  }

  .listing-container .content .map-full {
    display: flex;
    flex-direction: column;
    position: sticky;
    top: 0;
    height: calc(100vh - var(--navbar-height));
  }

  .listing-container .content .map-full .map-container {
    display: block;
    width: 100%;
    height: 100%;
    overflow: hidden;
  }
</style>