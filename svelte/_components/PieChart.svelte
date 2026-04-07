<script>
  import { run } from 'svelte/legacy';

  /**
   * @typedef {Object} Props
   * @property {number} [size]
   * @property {string} [bgColor]
   * @property {any} [shares]
   */

  /** @type {Props} */
  let { size = 200, bgColor = 'cornflowerblue', shares = $bindable([
    {
      percent: 44,
      color: 'orange'
    }
  ]) } = $props();
  
  let viewBox = $derived(`0 0 ${size} ${size}`), radius = $derived(size / 2), halfCircumference = $state();
  
  
  run(() => {
    halfCircumference = Math.PI * radius;
  });
  
  run(() => {
    let last = 0;
    for( const z in shares ) {
      const percent = shares[ z ].percent || 0;
      const pieSize = halfCircumference * (percent / 100)
      shares[ z ].dashArray = `0 ${halfCircumference - pieSize} ${pieSize}`
      shares[ z ].rotate = last * 360 / 100;
      last -= percent;
    }
  });
</script>
<svg width={size} height={size} {viewBox}>
	<circle r={radius} cx={radius} cy={radius} fill={bgColor}/>
	{#each shares as share}
		<circle
			r={radius / 2}
			cx={radius}
			cy={radius}
			fill="transparent"
			stroke={share.color}
			stroke-width={radius}
			stroke-dasharray={share.dashArray}
			transform="rotate({share.rotate}, {radius}, {radius})"/>
	{/each}
</svg>