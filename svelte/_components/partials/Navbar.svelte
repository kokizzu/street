<script>
  import { Icon } from '../../node_modules/svelte-icons-pack/dist';
  import { RiUserFacesUserLine } from '../../node_modules/svelte-icons-pack/dist/ri';

  export let username = /** @type {string} */ ('unknown');
</script>

<header>
  <div class="label">
    <a href="/" class="logo">
      <h2>HapSTR</h2>
    </a>
  </div>
  <div class="user-info">
    <a href="/user" class="user">
      <span>{username}</span>
      <Icon src={RiUserFacesUserLine} size="15" />
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
    justify-content: space-between;
    padding: 0 20px;
    border-bottom: 1px solid var(--gray-002);
    background-color: #FFF;
    z-index: 100;
  }

  header .label {
    width: fit-content;
    display: flex;
    flex-direction: row;
    gap: 0;
    align-items: center;
  }

  header .label .logo {
    display: flex;
    align-items: center;
    justify-content: flex-start;
    font-size: 17px;
    color: var(--gray-008);
    text-decoration: none;
    width: calc(var(--sidemenu-width) - 20px);
  }

  header .user-info {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 10px;
  }

  header .user-info a.user {
    display: flex;
    align-items: center;
    flex-direction: row;
    justify-content: center;
    gap: 6px;
    text-decoration: none;
    color: var(--gray-008);
    padding: 5px 10px;
    border-radius: 5px;
    border: 1px solid var(--gray-002);
  }

  header .user-info a.user:hover {
    background-color: var(--gray-001);
  }
</style>