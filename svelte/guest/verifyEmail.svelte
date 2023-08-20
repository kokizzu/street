<script>
  //@ts-nocheck
  let verified = '#{verified}';
  let error = '#{error}';
  let email = '#{email}';
</script>

<section class="verify_email_container">
  <div class="main_content">
    <h1>/*! title */</h1>
    {#if verified === 'true'}
      <p>{email} is verified</p>
    {/if}
    {#if error}
      <p class="error">{error}</p>
    {/if}
    <a href="/">
      <i class="gg-arrow-left-r" />
      <p>Back to main page</p>
    </a>
  </div>
</section>

<style>
  .verify_email_container {
    height: 100%;
    width: 100%;
    background-color: #f1f5f9;
    display: flex;
    color: #475569;
  }
  .main_content {
    width: 420px;
    height: fit-content;
    padding: 20px;
    filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    border-radius: 15px;
    display: flex;
    flex-direction: column;
    background-color: white;
    margin: 50px auto;
    border: 1px solid #cbd5e1;
  }
  .main_content h1 {
    font-weight: 700;
    font-size: 22px;
    margin: 0 auto 15px;
  }
  .main_content p {
    margin: 0;
    font-size: 15px;
    text-align: center;
  }
  .main_content a {
    padding: 8px 30px;
    font-size: 16px;
    display: inline-flex;
    flex-direction: row;
    align-items: center;
    align-content: center;
    justify-content: center;
    border: none;
    background-color: #6366f1;
    border-radius: 8px;
    color: white;
    cursor: pointer;
    text-decoration: none;
    width: fit-content;
    margin: 30px auto 0 auto;
  }
  .main_content a:hover {
    background-color: #7e80f1;
  }
  .main_content a i {
    margin-right: 8px;
  }
</style>
