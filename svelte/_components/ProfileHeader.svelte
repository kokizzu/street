<script>
	import { createEventDispatcher } from 'svelte';
	import PieChart from './PieChart.svelte'
	import PieLegend from './PieLegend.svelte'

	const dispatch = createEventDispatcher();

	function openSideMenu() {
      dispatch("opensidemenu");
   }
</script>

<header class="profile_header">
	<nav class="navbar">
		<div class='label_menu'>
			<button on:click|preventDefault={openSideMenu}>
				<i class='gg-menu'></i>
			</button>
			<p>DASHBOARD</p>
		</div>
		<div class="right_nav">
			<form class="search_input">
				<span>
					<i class="gg-search"></i>
				</span>
				<input type="text" placeholder="Search here..." name="search" />
			</form>
			<button class="profile_button">
				<img src="/assets/img/team-1-200x200.jpg" alt="Profile Picture" />
			</button>
		</div>
	</nav>
	<div class="stats">
		<!-- Chart Goes here -->
	</div>
</header>

<style>
	.profile_header {
      background-color: #EF4444;
      height: fit-content;
      position: relative;
      padding-top: 18px;
      padding-bottom: 80px;
      padding-right: 50px;
      padding-left: 50px;
      display: flex;
      flex-direction: column;
   }
   .profile_header .navbar {
   	display: flex;
   	flex-direction: row;
   	justify-content: space-between;
   	align-items: center;
   }
	.profile_header .navbar .label_menu {
		display: flex;
		flex-direction: row;
		align-items: center;
		color: white;
	}
	.profile_header .navbar .label_menu button {
		padding: 15px 8px;
		border: none;
		background: none;
		border-radius: 5px;
		font-size: 14px;
		color: white;
		cursor: pointer;
	}
	.profile_header .navbar .label_menu button:hover {
		background-color: rgba(255, 255, 255, 0.3);
	}
   .profile_header .navbar .label_menu p {
   	font-size: 17px;
   	font-weight: 600;
      line-height: 1.5rem;
      padding: 0;
      margin: 0 0 0 15px;
   }
   .profile_header .navbar .right_nav {
   	display: flex;
   	flex-direction: row;
   }
   .profile_header .navbar .right_nav .search_input {
   	margin-right: 15px;
   	margin-top: auto;
   	margin-left: 0;
   	margin-bottom: auto;
   }
   .profile_header .navbar .right_nav .search_input > span {
   	padding: 13px;
   	position: absolute;
   	z-index: 10;
   	font-size: 16px;
   	color: #CBD5E1;
   }
   .profile_header .navbar .right_nav .search_input > input {
   	width: 250px;
   	border: none;
   	border-radius: 8px;
   	padding-top: 12px;
   	padding-bottom: 12px;
   	padding-left: 40px;
   	padding-right: 12px;
   }
   .profile_header .navbar .right_nav .search_input > input:focus {
   	outline: 3px solid #3b82f6;
   }
   .profile_header .navbar .right_nav .profile_button {
      padding: 0;
    	border: none;
    	margin: 0;
    	width: fit-content;
    	height: fit-content;
    	border-radius: 9999px;
 	}
 	.profile_header .navbar .right_nav .profile_button:hover {
	 	box-shadow: 0 0 0 3px rgba(255, 255, 255, 0.5);
	 	cursor: pointer;
 	}
   .profile_header .navbar .right_nav .profile_button > img {
   	width: 50px;
   	height: 50px;
   	border-radius: 9999px;
   }
</style>