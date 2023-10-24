<script>
  //@ts-nocheck
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidAngleLeft from "svelte-icons-pack/fa/FaSolidAngleLeft";
  
  export let user;
  export let property;
  
</script>

<div class="upper_action">
	<a class='back_button' href='/realtor'>
		<Icon className='iconBack' color='#475569' size={18} src={FaSolidAngleLeft}/>
	</a>
	<span>Property ID: XXX123XXX</span>
</div>

<div>

</div>

<style>
    /* General purpose selector */
    :global(.back_button:hover .iconBack) {
        fill : #EF4444;
    }

    .upper_action {
        display          : flex;
        flex-direction   : row;
        align-items      : center;
        gap              : 10px;
        font-size        : 20px;
        margin           : -40px 0 0 0;
        padding          : 0 55px 70px 55px;
        position         : relative;
        background-color : #EF4444;
        font-weight      : 600;
        color            : #FFF;
        z-index          : 40;
    }

    .upper_action .back_button {
        padding          : 8px;
        background-color : #FFF;
        border-radius    : 5px;
        font-size        : 14px;
        cursor           : pointer;
        border           : none;
        text-decoration  : none;
        height           : fit-content;
        color            : #334155;
    }
</style>