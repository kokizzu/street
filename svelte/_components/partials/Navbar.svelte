<script>

</script>

<header>
  <div class="label">
    <a href="/" class="logo">
      <h2>HapSTR</h2>
    </a>
  </div>
</header>

<style>
  header {
    top: 0;
    left: 0;
    right: 0;
    position: fixed;
    width: 100%;
    height: var(--navbar-height);
    display: flex;
    align-items: center;
    justify-content: start;
    padding: 0 20px;
    border-bottom: 1px solid var(--gray-002);
    background-color: #FFF;
  }

  header .label {
    width: fit-content;
    display: flex;
    flex-direction: row;
    gap: 20px;
    align-items: center;
  }

  header .label .logo {
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 17px;
    color: var(--gray-008);
    text-decoration: none;
  }
</style>