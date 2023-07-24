<script>
  
  import AdminSubMenu from '../_components/AdminSubMenu.svelte';
  import Menu from '../_components/Menu.svelte';
  import ProfileHeader from '../_components/ProfileHeader.svelte';
  import Footer from '../_components/Footer.svelte';
  import Dropzone from '../_components/UploadDropzone.svelte';
  
  let files = {
    accepted: [],
    rejected: [],
  };
  
  // TODO: create component out of this
  
  function handleFilesSelect( e ) {
    const {acceptedFiles, fileRejections} = e.detail;
    files.accepted = [...files.accepted, ...acceptedFiles];
    files.rejected = [...files.rejected, ...fileRejections];
    if( acceptedFiles.length<1 ) return;
    
    let file = acceptedFiles[ 0 ];
    //console.log(file.name,file.size,file.type);
    var formData = new FormData();
    formData.append( 'rawFile', file );
    formData.append( 'purpose', 'property' ); // property or floorPlan
    var ajax = new XMLHttpRequest();
    ajax.upload.addEventListener( 'progress', progressHandler, false );
    ajax.addEventListener( 'load', completeHandler, false );
    ajax.addEventListener( 'error', errorHandler, false );
    ajax.addEventListener( 'abort', abortHandler, false );
    ajax.open( 'POST', '/user/uploadFile' );
    ajax.send( formData );
  }
  
  let uploadStatus = '';
  let uploadPercent = 0;
  let lastImage = '';
  
  function progressHandler( event ) {
    var percent = (event.loaded / event.total) * 100;
    uploadPercent = Math.round( percent );
    uploadStatus = uploadPercent + '% uploaded... please wait';
  }
  
  function completeHandler( event ) {
    uploadStatus = event.target.responseText;
    uploadPercent = 0; //wil clear progress bar after successful upload
    const out = JSON.parse(event.target.responseText)
    if(!out.error) {
      lastImage = out.resizedUrl; // .originalUrl also available
    }
  }
  
  function errorHandler( event ) {
    uploadStatus = 'Upload Failed';
  }
  
  function abortHandler( event ) {
    uploadStatus = 'Upload Aborted';
  }
  
  let segments = {/* segments */};
</script>

<section class='dashboard'>
  <Menu access={segments} />
  <div class='dashboard_main_content'>
    <ProfileHeader />
    <AdminSubMenu></AdminSubMenu>
    <div class='content'>
      
      <Dropzone on:drop={handleFilesSelect} />
      <ol>
        {#each files.accepted as item}
          <li>{item.name}</li>
        {/each}
      </ol>
      <h3>{uploadStatus}</h3>
      <progress value={uploadPercent} max='100' style='width:300px;'></progress>
      
      {#if files.rejected && files.rejected.length}
        <h3>Errors</h3>
        <ol>
          {#each files.rejected as item}
            <li>{item.name}</li>
          {/each}
        </ol>
      {/if}
      
      {#if lastImage}
        <hr/>
        <img src={lastImage} alt='last uploaded'/>
      {/if}
    </div>
    <Footer></Footer>
  </div>
</section>

<style>

</style>