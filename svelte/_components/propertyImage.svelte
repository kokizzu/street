<script>
  /**
   * @typedef {Object} Props
   * @property {string} [src]
   * @property {string} [alt]
   */

  /** @type {Props} */
  let { src = $bindable('/assets/img/placeholder.webp'), alt = '' } = $props();

  function onerror() {
    src = '/assets/img/placeholder.webp';
  }
</script>

<picture>
  <img {src} {alt} {onerror} />
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