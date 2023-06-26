<script>
  
  import { onMount } from 'svelte';
  let segment2;
  onMount(() => {
    segment2 = window.location.pathname.split('/')[2] || '';
  })
</script>

<ul class='submenu'>
  <li class:active={segment2===''}><a href="/admin">Dashboard</a></li>
  <li class:active={segment2==='users'}><a href="/admin/users">Users</a></li>
  <li class:active={segment2==='properties'}><a href="/admin/properties">Properties</a></li>
<!--  <li class:active={segment2==='sessions'}><a href="/admin/sessions">Sessions</a></li>-->
<!--  <li class:active={segment2==='actionLogs'}><a href="/admin/actionLogs">Action Logs</a></li>-->
</ul>

<style>
    .submenu {
        list-style-type  : none;
        margin           : 0;
        padding          : 0;
        overflow         : hidden;
        background-color : #666;
    }

    .submenu li {
        float : left;
    }

    .submenu li a {
        display         : block;
        color           : white;
        text-align      : center;
        padding         : 14px 16px;
        text-decoration : none;
    }

    .submenu li a:hover:not(.active) {
        background-color : #111;
    }

    .active {
        background-color : #4CAF50;
    }
</style>