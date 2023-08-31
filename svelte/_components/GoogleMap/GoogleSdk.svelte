<script>
  import loader from '../asyncScriptLoader.js'
  import { onMount, createEventDispatcher } from 'svelte'
  import { mapsLoaded, mapsLoading } from './stores.js'

  const dispatch = createEventDispatcher()

  export let apiKey
  
  $: $mapsLoaded && dispatch('ready')

  onMount(() => {
    window.byGmapsReady = () => {
      mapsLoaded.set(true)
      delete window.byGmapsReady
    }

    if ($mapsLoaded) {
      dispatch('ready')
    }

    if (!$mapsLoading) {
      const url = [
        '//maps.googleapis.com/maps/api/js?',
        apiKey ? `key=${apiKey}&` : '',
        'libraries=places&callback=byGmapsReady'
      ].join('')

      mapsLoading.set(true)

      loader(
        [
          { type: 'script', url }
        ],
        () => { return $mapsLoaded },
        () => {}
      )
    }
  })
</script>
