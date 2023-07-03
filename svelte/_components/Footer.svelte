<script>
	const year = new Date().getFullYear();
</script>

<footer>
	<p>Copyright © {year}</p>
	<nav>
		<a href="/privacy">Privacy</a>
		<a href="/tos">Terms and Conditions </a>
		<a href="/about">About Us</a>
	</nav>
</footer>

<style>
	footer {
		font-size: 16px;
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		margin: 0 50px;
		padding: 16px 0;
		border-top: 1px solid #CBD5E1;
		color: #475569;
	}
	footer nav {
		display: flex;
		flex-direction: row;
	}
	footer nav a {
		text-decoration: none;
		margin-left: 15px;
		color: #475569;
	}
	footer nav a:hover {
		color: #EF4444;
	}
</style>