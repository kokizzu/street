<script>
  import { fromEvent } from 'file-selector';
  import {
    fileAccepted, fileMatchSize, isEvtWithFiles, isIeOrEdge, isPropagationStopped,
    TOO_MANY_FILES_REJECTION
  } from './UploadDropzone.js';
  import { createEventDispatcher, onDestroy } from 'svelte';
  
  //props
  
  /**
   * @typedef {Object} Props
   * @property {any} accept - Set accepted file types.
See https://github.com/okonet/attr-accept for more information. - string or string[]
   * @property {boolean} [disabled]
   * @property {any} [getFilesFromEvent]
   * @property {any} [maxSize]
   * @property {number} [minSize]
   * @property {boolean} [multiple]
   * @property {boolean} [preventDropOnDocument]
   * @property {boolean} [noClick]
   * @property {boolean} [noKeyboard]
   * @property {boolean} [noDrag]
   * @property {boolean} [noDragEventsBubbling]
   * @property {string} [containerClasses]
   * @property {string} [containerStyles]
   * @property {boolean} [disableDefaultStyles]
   * @property {string} [name]
   * @property {any} inputElement
   * @property {boolean} [required]
   * @property {import('svelte').Snippet} [children]
   */

  /** @type {Props} */
  let {
    accept,
    disabled = false,
    getFilesFromEvent = fromEvent,
    maxSize = Infinity,
    minSize = 0,
    multiple = true,
    preventDropOnDocument = true,
    noClick = false,
    noKeyboard = false,
    noDrag = false,
    noDragEventsBubbling = false,
    containerClasses = '',
    containerStyles = '',
    disableDefaultStyles = false,
    name = '',
    inputElement = $bindable(),
    required = false,
    children
  } = $props();
  const dispatch = createEventDispatcher();
  
  //state
  
  let dropzoneState = {
    isFocused: false,
    isFileDialogActive: false,
    isDragActive: false,
    isDragAccept: false,
    isDragReject: false,
    draggedFiles: [],
    acceptedFiles: [],
    fileRejections: [],
  };
  
  let rootRef = $state();
  
  function resetState() {
    dropzoneState.isFileDialogActive = false;
    dropzoneState.isDragActive = false;
    dropzoneState.draggedFiles = [];
    dropzoneState.acceptedFiles = [];
    dropzoneState.fileRejections = [];
  }
  
  // Fn for opening the file dialog programmatically
  function openFileDialog() {
    if( inputElement ) {
      inputElement.value = null; // TODO check if null needs to be set
      dropzoneState.isFileDialogActive = true;
      inputElement.click();
    }
  }
  
  // Cb to open the file dialog when SPACE/ENTER occurs on the dropzone
  function onKeyDownCb( event ) {
    // Ignore keyboard events bubbling up the DOM tree
    if( !rootRef || !rootRef.isEqualNode( event.target ) ) {
      return;
    }
    
    if( event.keyCode===32 || event.keyCode===13 ) {
      event.preventDefault();
      openFileDialog();
    }
  }
  
  // Update focus state for the dropzone
  function onFocusCb() {
    dropzoneState.isFocused = true;
  }
  
  function onBlurCb() {
    dropzoneState.isFocused = false;
  }
  
  // Cb to open the file dialog when click occurs on the dropzone
  function onClickCb() {
    if( noClick ) {
      return;
    }
    
    // In IE11/Edge the file-browser dialog is blocking, therefore, use setTimeout()
    // to ensure React can handle state changes
    // See: https://github.com/react-dropzone/react-dropzone/issues/450
    if( isIeOrEdge() ) {
      setTimeout( openFileDialog, 0 );
    } else {
      openFileDialog();
    }
  }
  
  function onDragEnterCb( event ) {
    event.preventDefault();
    stopPropagation( event );
    
    dragTargetsRef = [...dragTargetsRef, event.target];
    
    if( isEvtWithFiles( event ) ) {
      Promise.resolve( getFilesFromEvent( event ) ).then( draggedFiles => {
        if( isPropagationStopped( event ) && !noDragEventsBubbling ) {
          return;
        }
        
        dropzoneState.draggedFiles = draggedFiles;
        dropzoneState.isDragActive = true;
        
        dispatch( 'dragenter', {
          dragEvent: event,
        } );
      } );
    }
  }
  
  function onDragOverCb( event ) {
    event.preventDefault();
    stopPropagation( event );
    
    if( event.dataTransfer ) {
      try {
        event.dataTransfer.dropEffect = 'copy';
      } catch {
      } /* eslint-disable-line no-empty */
    }
    
    if( isEvtWithFiles( event ) ) {
      dispatch( 'dragover', {
        dragEvent: event,
      } );
    }
    
    return false;
  }
  
  function onDragLeaveCb( event ) {
    event.preventDefault();
    stopPropagation( event );
    
    // Only deactivate once the dropzone and all children have been left
    const targets = dragTargetsRef.filter(
      target => rootRef && rootRef.contains( target ),
    );
    // Make sure to remove a target present multiple times only once
    // (Firefox may fire dragenter/dragleave multiple times on the same element)
    const targetIdx = targets.indexOf( event.target );
    if( targetIdx!== -1 ) {
      targets.splice( targetIdx, 1 );
    }
    dragTargetsRef = targets;
    if( targets.length>0 ) {
      return;
    }
    
    dropzoneState.isDragActive = false;
    dropzoneState.draggedFiles = [];
    
    if( isEvtWithFiles( event ) ) {
      dispatch( 'dragleave', {
        dragEvent: event,
      } );
    }
  }
  
  function onDropCb( event ) {
    event.preventDefault();
    stopPropagation( event );
    
    dragTargetsRef = [];
    
    if( isEvtWithFiles( event ) ) {
      dispatch( 'filedropped', {
        event,
      } );
      Promise.resolve( getFilesFromEvent( event ) ).then( files => {
        if( isPropagationStopped( event ) && !noDragEventsBubbling ) {
          return;
        }
        
        const acceptedFiles = [];
        const fileRejections = [];
        
        files.forEach( file => {
          const [accepted, acceptError] = fileAccepted( file, accept );
          const [sizeMatch, sizeError] = fileMatchSize( file, minSize, maxSize );
          if( accepted && sizeMatch ) {
            acceptedFiles.push( file );
          } else {
            const errors = [acceptError, sizeError].filter( e => e );
            fileRejections.push( {file, errors} );
          }
        } );
        
        if( !multiple && acceptedFiles.length>1 ) {
          // Reject everything and empty accepted files
          acceptedFiles.forEach( file => {
            fileRejections.push( {file, errors: [TOO_MANY_FILES_REJECTION]} );
          } );
          acceptedFiles.splice( 0 );
        }
        
        // Files dropped keep input in sync
        if( event.dataTransfer ) {
          inputElement.files = event.dataTransfer.files;
        }
        
        dropzoneState.acceptedFiles = acceptedFiles;
        dropzoneState.fileRejections = fileRejections;
        
        dispatch( 'drop', {
          acceptedFiles,
          fileRejections,
          event,
        } );
        
        if( fileRejections.length>0 ) {
          dispatch( 'droprejected', {
            fileRejections,
            event,
          } );
        }
        
        if( acceptedFiles.length>0 ) {
          dispatch( 'dropaccepted', {
            acceptedFiles,
            event,
          } );
        }
      } );
    }
    resetState();
  }
  
  function composeHandler( fn ) {
    return disabled ? null : fn;
  }
  
  function composeKeyboardHandler( fn ) {
    return noKeyboard ? null : composeHandler( fn );
  }
  
  function composeDragHandler( fn ) {
    return noDrag ? null : composeHandler( fn );
  }
  
  function stopPropagation( event ) {
    if( noDragEventsBubbling ) {
      event.stopPropagation();
    }
  }
  
  // allow the entire document to be a drag target
  function onDocumentDragOver( event ) {
    if( preventDropOnDocument ) {
      event.preventDefault();
    }
  }
  
  let dragTargetsRef = [];
  
  function onDocumentDrop( event ) {
    if( !preventDropOnDocument ) {
      return;
    }
    if( rootRef && rootRef.contains( event.target ) ) {
      // If we intercepted an event for our instance, let it propagate down to the instance's onDrop handler
      return;
    }
    event.preventDefault();
    dragTargetsRef = [];
  }
  
  // Update file dialog active state when the window is focused on
  function onWindowFocus() {
    // Execute the timeout only if the file dialog is opened in the browser
    if( dropzoneState.isFileDialogActive ) {
      setTimeout( () => {
        if( inputElement ) {
          const {files} = inputElement;
          
          if( !files.length ) {
            dropzoneState.isFileDialogActive = false;
            dispatch( 'filedialogcancel' );
          }
        }
      }, 300 );
    }
  }
  
  onDestroy( () => {
    // This is critical for canceling the timeout behaviour on `onWindowFocus()`
    inputElement = null;
  } );
  
  function onInputElementClick( event ) {
    event.stopPropagation();
  }
</script>

<style>
    .dropzone {
        flex             : 1;
        display          : flex;
        flex-direction   : column;
        align-items      : center;
        padding          : 20px;
        border-width     : 2px;
        border-radius    : 2px;
        border-color     : #EEE;
        border-style     : dashed;
        background-color : #FAFAFA;
        color            : #BDBDBD;
        outline          : none;
        transition       : border 0.24s ease-in-out;
    }

    .dropzone:focus {
        border-color : #2196F3;
    }
</style>

<svelte:window onfocus={onWindowFocus} ondragover={onDocumentDragOver} ondrop={onDocumentDrop} />

<div
  bind:this={rootRef}
  tabindex='0'
  role='button'
  class="{disableDefaultStyles ? '' : 'dropzone'}
  {containerClasses}"
  style={containerStyles}
  onkeydown={composeKeyboardHandler(onKeyDownCb)}
  onfocus={composeKeyboardHandler(onFocusCb)}
  onblur={composeKeyboardHandler(onBlurCb)}
  onclick={composeHandler(onClickCb)}
  ondragenter={composeDragHandler(onDragEnterCb)}
  ondragover={composeDragHandler(onDragOverCb)}
  ondragleave={composeDragHandler(onDragLeaveCb)}
  ondrop={composeDragHandler(onDropCb)}
>
  <input
    {accept}
    {multiple}
    {required}
    type='file'
    {name}
    autocomplete='off'
    tabindex='-1'
    onchange={onDropCb}
    onclick={onInputElementClick}
    bind:this={inputElement}
    style='display: none;'
  />
  {#if children}{@render children()}{:else}
    <p>Drag 'n' drop some files here, or click to select files</p>
  {/if}
</div>
