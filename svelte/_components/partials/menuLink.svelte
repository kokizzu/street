<script>
   import { Icon } from 'svelte-icons-pack';

  const pathname = /** @type {string}*/ (window.location.pathname);

  export let href = /** @type {string} */ ('/');
  export let icon = /** @type {import('svelte-icons-pack').IconType} */ ({});
  export let iconSize = /** @type {number} */ (17);
  export let label = /** @type {string} */ ('Home');
</script>

<a
  href={href}
  class:active={href === pathname}
  title={label}
  >
  <Icon
    src={icon}
    size={iconSize}
  />
  <span>{label}</span>
</a>

<style>
  a {
    text-decoration: none;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-start;
    gap: 10px;
    padding: 10px 15px;
    border-radius: 8px;
    cursor: pointer;
  }

  a:hover {
    background-color: var(--orange-transparent);
    color: var(--orange-007);
  }

  a.active {
    background-color: var(--orange-transparent);
    color: var(--orange-007);
  }
</style>