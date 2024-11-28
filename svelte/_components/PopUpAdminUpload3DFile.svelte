<script>
  import { notifier } from './notifier';
  import InputBox from './InputBox.svelte';

  let isShow      = /** @type {boolean} */ (false);
  export let isSubmitted = /** @type {boolean} */ (false);
  export let uploadingProgressStr = /** @type {string} */ ('');

  const allowedExtensions = /** @type {string[]} */ ([
    '.obj', '.fbx', '.stl', '.amf', '.iges', '.glb'
  ]);

  let fileName    = /** @type {string} */ ('');
  let file3d      = /** @type {File} */ (null);
  let inputElm    = /** @type {HTMLInputElement} */ (null);
  let isDragging  = /** @type {boolean} */ (false);
  let propertyCountry = /** @type {string} */ ('Default');
  let propertyKey = /** @type {string} */ ('');

  export const Reset = () => {
    isSubmitted = false;
    fileName = ''
    file3d = null;
    inputElm.files = null;
    inputElm.value = '';
    isDragging = false;
    uploadingProgressStr = '';
  }
  export const Show = () => isShow = true;
  export const Close = () => {
    isShow = false;
    Reset();
  }

  /**
   * @description Submit Upload 3D File
   * @type {Function}
   * @param file {File}
   * @param country {string}
   * @param propKey {string}
   * @returns {Promise<void>}
   */
   export let OnSubmit = async function(file, country, propKey) {
    console.log('File Name =', file.name);
    console.log('Country =', country);
    console.log('Property Key =', propKey);
  };

  function handleDrop(/** @type {DragEvent}*/ event) {
    event.preventDefault();
    isDragging = false;

    const droppedFile = event.dataTransfer.files[0];

    if (!droppedFile) {
      notifier.showError('No file dropped.');
      return;
    }

    // Validate file type
    const fileExtension = droppedFile.name.slice(
      droppedFile.name.lastIndexOf('.')
    ).toLowerCase();

    if (!allowedExtensions.includes(fileExtension)) {
      notifier.showError(`Invalid file type. Please upload a 3D file (${allowedExtensions.join(', ')}).`);
      return;
    }

    file3d = droppedFile;
    fileName = file3d.name;
  }

  function handleDragOver(/** @type {DragEvent} */event) {
    event.preventDefault();
    isDragging = true;
  }

  function handleDragLeave() {
    isDragging = false;
  }

  function handleFileInputChange(/** @type {
    Event & { currentTarget: EventTarget & HTMLInputElement
  }} */ event) {
    const file = event.currentTarget.files[0];
    file3d = file;
    fileName = file.name;
    console.log('FileName:',fileName)
  }
</script>

<div class="backdrop {isShow ? 'show' : ''}">
  <div class="form-upload">
    <h3>Upload 3D File</h3>
    <InputBox
      label="Property Key"
      type="text"
      id="propertyKey"
      autocomplete="on"
      placeholder="#US78193"
      bind:value={propertyKey}
    />
    <InputBox
      label="Property Country"
      type="combobox-arr"
      id="propertyCountry"
      valuesArr={[
        'Default', 'US', 'TW'
      ]}
      bind:value={propertyCountry}
    />
    <div class="upload-zone">
      <label
        for="3dFile"
        on:dragover={handleDragOver}
        on:dragleave={handleDragLeave}
        on:drop={handleDrop}
      >
        {#if uploadingProgressStr === ''}
          {#if fileName !== ''}
            <h3>{fileName}</h3>
            <span>Click or Drop file here to change</span>
          {/if}
          {#if fileName === ''}
            <b>Click or Drag and Drop your 3D file here (obj, fbx, stl, amf, iges, glb)</b>
          {/if}
        {/if}

        {#if uploadingProgressStr !== ''}
          <b>{uploadingProgressStr}</b>
        {/if}
      </label>
      <input
        bind:this={inputElm}
        on:change={handleFileInputChange}
        type="file"
        name="3dFile"
        id="3dFile"
        accept=".obj,.fbx,.stl,.amf,.iges,.glb" 
        required
      />
    </div>
    <div class="buttons">
      <button class="cancel-btn" on:click={Close}>
        Cancel
      </button>
      <button class="submit-btn" on:click={() => OnSubmit(file3d, propertyCountry, propertyKey)} disabled={!file3d}>
        {#if isSubmitted}
          <span>Uploading....</span>
        {/if}
        {#if !isSubmitted}
          <span>Upload</span>
        {/if}
      </button> 
    </div>
  </div>
</div>

<style>
  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }

  :global(.spin) {
    animation: spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
  }

  .backdrop {
    position: fixed;
    z-index: 2001;
    background: rgba(41, 41, 41, 0.5);
    width: 100%;
    left: 0;
    top: 0;
    right: 0;
    bottom: 0;
    height: auto;
    display: none;
    justify-content: center;
    align-items: center;
  }

  .backdrop.show {
    display: flex;
  }

  .form-upload {
    padding: 20px;
    border-radius: 10px;
    background-color: #FFF;
    display: flex;
    flex-direction: column;
    gap: 20px;
    color: #475569;
    width: 600px;
  }

  .form-upload h3 {
    margin: 0 0 0 20px;
    font-size: 18px;
  }

  .form-upload .upload-zone {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 250px;
    width: 100%;
  }

  .form-upload .upload-zone label {
    cursor: pointer;
    height: 100%;
    width: 100%;
    border: 3px dashed var(--gray-003);
    display: flex;
    flex-direction: column;
    gap: 8px;
    justify-content: center;
    align-items: center;
    font-size: 14px;
    border-radius: 8px;
  }

  .form-upload .upload-zone label h3 {
    margin: 0;
    font-size: 20px;
  }

  .form-upload .upload-zone label:hover {
    border-color: var(--orange-005);
    color: var(--orange-005);
  }

  .form-upload .upload-zone input {
    display: none;
    visibility: hidden;
  }

  .form-upload .buttons {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    align-items: stretch;
    font-weight: 500;
    width: 100%;
    gap: 10px;
  }

  .form-upload .buttons button {
    border-radius: 999px;
    border: none;
    padding: 10px;
    cursor: pointer;
    width: fit-content;
    padding: 12px 20px;
    text-transform: capitalize;
  }

  .form-upload .buttons .cancel-btn {
    border: none;
    background-color: transparent;
  }

  .form-upload .buttons .cancel-btn:hover {
    background-color: var(--orange-transparent);
    color: var(--orange-005);
  }

  .form-upload .buttons .submit-btn {
    background-color: var(--orange-006);
    color: #FFF;
  }

  .form-upload .buttons .submit-btn:disabled,
  .form-upload .buttons .submit-btn:disabled:hover {
    background-color: var(--gray-003);
    color: var(--gray-008);
  }

  .form-upload .buttons .submit-btn:hover {
    background-color: var(--orange-005);
    color: #FFF;
  }
</style>