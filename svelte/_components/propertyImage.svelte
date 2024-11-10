<script>
  export let src = '/assets/img/placeholder.webp';
  export let alt = '';

  function onerror() {
    src = '/assets/img/placeholder.webp';
  }
</script>

<picture>
  <img {src} {alt} on:error={onerror} />
</picture>

<style>
  picture {
    width: 100%;
    height: 130px;
    overflow: hidden;
  }

  picture img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition-duration : 75ms;
  }
</style>