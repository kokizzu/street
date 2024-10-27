<script>
  const pathLv1 = /** @type {string}*/ (window.location.pathname.split( '/' )[ 1 ]);
</script>

<aside>
  <div class="container">
    <nav class="nav-menu">
      <a href="/" class:active={pathLv1 === ''}>Home</a>
      <a href="/buyer" class:active={pathLv1 === 'buyer'}>Buyers</a>
      <a href="/realtor" class:active={pathLv1 === 'realtor'}>Realtors</a>
      <a href="/listings" class:active={pathLv1 === 'listings'}>Listings</a>
      <a href="/revenue" class:active={pathLv1 === 'revenue'}>Revenue</a>
    </nav>
    <span class="separator" />
    <nav class="nav-menu">
      <a href="/admin" class:active={pathLv1 === 'admin'}>Admin</a>
      <a href="/user" class:active={pathLv1 === 'user'}>Profile</a>
      <button class="red">Logout</button>
    </nav>
  </div>
</aside>

<style>
  aside {
    position: fixed;
    top: var(--navbar-height);
    left: 0;
    bottom: 0;
    height: 100%;
    width: var(--sidemenu-width);
    border-right: 1px solid var(--gray-002);
    background-color: #FFF;
  }

  aside .separator {
    width: 80%;
    margin: auto;
    height: 1px;
    background-color: var(--gray-002);
  }

  aside .container {
    padding: 10px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    height: fit-content;
    width: 100%;
    margin: 0;
  }

  aside .container .nav-menu {
    display: flex;
    flex-direction: column;
    gap: 5px;
    height: fit-content;
    flex-wrap: nowrap;
    width: 100%;
    margin: 0;
  }

  aside .container .nav-menu a,
  aside .container .nav-menu button {
    text-decoration: none;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    padding: 10px 15px;
    border-radius: 8px;
    cursor: pointer;
    font-size: 17px;
    color: var(--gray-008);
    border: none;
    background-color: transparent;
  }

  aside .container .nav-menu a:hover,
  aside .container .nav-menu button:hover {
    background-color: var(--gray-001);
  }

  aside .container .nav-menu a.active {
    background-color: var(--orange-transparent);
    color: var(--orange-007);
  }

  aside .container .nav-menu button.red:hover {
    background-color: var(--red-transparent);
    color: var(--red-007);
  }
</style>