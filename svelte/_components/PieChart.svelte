<script>
  export let size = 200;
  export let bgColor = 'cornflowerblue'
  export let shares = [
    {
      percent: 44,
      color: 'orange'
    }
  ];
  
  let viewBox, radius, halfCircumference;
  $: viewBox = `0 0 ${size} ${size}`;
  $: radius = size / 2;
  $: halfCircumference = Math.PI * radius;
  
  $: {
    let last = 0;
    for( const z in shares ) {
      const percent = shares[ z ].percent || 0;
      const pieSize = halfCircumference * (percent / 100)
      shares[ z ].dashArray = `0 ${halfCircumference - pieSize} ${pieSize}`
      shares[ z ].rotate = last * 360 / 100;
      last -= percent;
    }
  }
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