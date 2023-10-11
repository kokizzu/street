<script>
  let scriptLoaded = false, cssLoaded = false;
  
  function init3dTiles() {
    // Enable simultaneous requests.
    Cesium.RequestScheduler.requestsByServer["tile.googleapis.com:443"] = 18;
    
    // Create the viewer.
    const viewer = new Cesium.Viewer('cesiumContainer', {
      imageryProvider: false,
      baseLayerPicker: false,
      geocoder: false,
      globe: false,
      // https://cesium.com/blog/2018/01/24/cesium-scene-rendering-performance/#enabling-request-render-mode
      requestRenderMode: true,
    });
    
    // Add 3D Tiles tileset.
    const tileset = viewer.scene.primitives.add(new Cesium.Cesium3DTileset({
      url: "https://tile.googleapis.com/v1/3dtiles/root.json?key=AIzaSyBKF5w6NExgYbmNMvlbMqF6sH2X4dFvMBg",
      // This property is needed to appropriately display attributions
      // as required.
      showCreditsOnScreen: true,
    }));
  }
  
  function checkFilesLoaded() {
    if (scriptLoaded && cssLoaded) {
      init3dTiles();
    }
  }
  
  const scriptElement = document.createElement("script");
  scriptElement.src = "https://ajax.googleapis.com/ajax/libs/cesiumjs/1.105/Build/Cesium/Cesium.js";
  scriptElement.onload = () => {
    scriptLoaded = true;
    checkFilesLoaded();
  };
  document.body.appendChild(scriptElement);
  
  const linkElement = document.createElement("link");
  linkElement.rel = "stylesheet";
  linkElement.href = "https://ajax.googleapis.com/ajax/libs/cesiumjs/1.105/Build/Cesium/Widgets/widgets.css";
  linkElement.onload = () => {
    cssLoaded = true;
    checkFilesLoaded();
  };
  document.head.appendChild(linkElement);
</script>

<div id="cesiumContainer"></div>

<style>
    #cesiumContainer {
        height : 100%;
        width  : 100%;
    }
</style>