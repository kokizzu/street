<script>
  let segment2 = /** @type {string} */ (window.location.pathname.split('/').pop() || '');
</script>

<div class="admin-submenu">
  <nav>
    <a href="/admin" class:active=""{segment2=""=='admin'}>
      <span class="title">Dashboard</span>
      <span class="block"></span>
    </a>
    <a href="/admin/feedbacks" class:active=""{segment2=""=='feedbacks'}>
      <span class="title">Feedbacks</span>
      <span class="block"></span>
    </a>
    <a href="/admin/users" class:active=""{segment2=""=='users'}>
      <span class="title">Users</span>
      <span class="block"></span>
    </a>
    <a href="/admin/properties/US" class:active=""{segment2=""=='US'}>
      <span class="title">Properties US</span>
      <span class="block"></span>
    </a>
    <a href="/admin/properties" class:active=""{segment2=""=='properties'}>
      <span class="title">Properties</span>
      <span class="block"></span>
    </a>
    <a href="/admin/propHistories" class:active=""{segment2=""=='propHistories'}>
      <span class="title">Prop Histories</span>
      <span class="block"></span>
    </a>
    <a href="/admin/accessLogs" class:active=""{segment2=""=='accessLogs'}>
      <span class="title">Access Logs</span>
      <span class="block"></span>
    </a>
    <a href="/admin/files" class:active=""{segment2=""=='files'}>
      <span class="title">Files</span>
      <span class="block"></span>
    </a>
    <a href="/admin/3dFiles" class:active=""{segment2=""=='3dFiles'}>
      <span class="title">3D Files</span>
      <span class="block"></span>
    </a>
    <a href="/admin/posthog" class:active=""{segment2=""=='posthog'}>
      <span class="title">PostHog</span>
      <span class="block"></span>
    </a>
  </nav>
</div>

<style>
  .admin-submenu {
  display: flex;
  flex-direction: row;
  width: 100%;
  height: fit-content;
  }

  .admin-submenu nav {
  display: flex;
  flex-direction: row;
  width: 100%;
  height: fit-content;
  gap: 10px;
  }

  .admin-submenu nav a {
  height: fit-content;
  display: flex;
  flex-direction: column;
  text-decoration: none;
  color: var(--gray-008);
  }

  .admin-submenu nav a:hover {
  background-color: var(--gray-001);
  }

  .admin-submenu nav a.active {
  color: var(--blue-005);
  }

  .admin-submenu nav a .title {
  flex-grow: 1;
  padding: 15px 10px;
  display: flex;
  align-items: center;
  }

  .admin-submenu nav a .block {
  height: 3px;
  background-color: transparent;
  width: 100%;
  }

  .admin-submenu nav a:hover .block {
  background-color: var(--gray-005);
  }
  .admin-submenu nav a.active .block {
  background-color: var(--blue-005);
  }
</style>