<script>
  // @ts-nocheck
  import { onMount } from 'svelte';
  import { isSideMenuOpen, isLangTWN } from './uiState.js';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import FaSolidBars from 'svelte-icons-pack/fa/FaSolidBars';

  let storedLang = localStorage.getItem('isLangTWN');
  onMount(() => {
    if (!storedLang) {
      localStorage.setItem('isLangTWN', 'false');
    }
    isLangTWN.set(JSON.parse(storedLang));
  });

  function openSideMenu() {
    isSideMenuOpen.set(!$isSideMenuOpen);
  }
  function toggleLang() {
    isLangTWN.set(!$isLangTWN);
    localStorage.setItem('isLangTWN', $isLangTWN.toString());
  }
</script>

<header class="profile_header">
  <nav class="navbar">
    <div class="label_menu">
      <button on:click|preventDefault={openSideMenu}>
        <Icon size={20} color="#FFF" src={FaSolidBars} />
      </button>
      <p>DASHBOARD</p>
    </div>
    <div class="right_nav">
      <button on:click={toggleLang} class="toggle_lang">
        <span class:lang={$isLangTWN}>TWN</span>
        <span class:lang={!$isLangTWN}>ENG</span>
      </button>
      <button class="profile_button">
        <img src="/assets/img/team-1-200x200.jpg" alt="profile" />
      </button>
    </div>
  </nav>
</header>

<style>
  .profile_header {
    background-color: #ef4444;
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
    padding: 10px;
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
    gap: 30px;
    align-items: center;
  }
  .profile_header .navbar .right_nav .toggle_lang {
    border: none;
    padding: 6px 8px;
    background-color: #ffffff;
    border-radius: 8px;
    height: fit-content;
    font-size: 10px;
    font-weight: 600;
    display: flex;
    gap: 4px;
    align-items: center;
    color: #475569;
    cursor: pointer;
  }
  .profile_header .navbar .right_nav .toggle_lang:hover {
    background-color: #f1f5f9;
    filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
  }
  .profile_header .navbar .right_nav .toggle_lang span {
    padding: 5px;
    background-color: transparent;
  }
  .profile_header .navbar .right_nav .toggle_lang span.lang {
    background-color: #3b82f6;
    border-radius: 5px;
    color: #ffffff;
  }
  .profile_header .navbar .right_nav .profile_button {
    padding: 0;
    border: none;
    margin: 0;
    width: fit-content;
    height: fit-content;
    border-radius: 50%;
    filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
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
