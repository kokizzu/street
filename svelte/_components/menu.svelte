<script>
  import {onMount} from "svelte";
  
  export let access = {
    "admin": false,
    "buyer": false,
    "realtor": false,
    "user": false
  }
  
  let segment;
  onMount(() => {
    segment = ((window || {}).location + '').split('/')[4] // [http, '', domain, first-segment]
  })
</script>

<ul class="menu">
		<li class:active={segment === ''}><a href="/">Home</a></li>
	{#if access.buyer }
		<li class:active={segment === 'buyer'}><a href="/buyer">Buyer</a></li>
	{/if}
	{#if access.realtor}
		<li class:active={segment === 'realtor'}><a href="/realtor">Realtor</a></li>
	{/if}
	{#if access.admin }
		<li class:active={segment === 'admin'}><a href="/admin">Admin</a></li>
	{/if}
	{#if access.user}
		<li class:active={segment === 'user'} style="float: right"><a href="/user">Profile</a></li>
	{/if}
</ul>

<style>
    .menu {
        list-style-type  : none;
        margin           : 0;
        padding          : 0;
        overflow         : hidden;
        background-color : #333;
    }

    .menu li {
        float : left;
    }

    .menu li a {
        display         : block;
        color           : white;
        text-align      : center;
        padding         : 14px 16px;
        text-decoration : none;
    }

    .menu li a:hover:not(.active) {
        background-color : #111;
    }

    .menu li a.active {
        float : right;
    }

    .active {
        background-color : #4CAF50;
    }
</style>