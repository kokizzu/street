<script>
	const year = /** @type {number} */ (new Date().getFullYear());
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
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		margin: 0;
		padding: 10px 20px;
		border-top: 1px solid var(--gray-002);
	}

	footer nav {
		display: flex;
		flex-direction: row;
		justify-content: center;
		align-items: center;
		gap: 15px;
	}

	footer nav a {
		color: var(--gray-008);
		text-decoration: none;
	}

	footer nav a:hover {
		color: var(--orange-005);
		text-decoration: underline;
	}
</style>