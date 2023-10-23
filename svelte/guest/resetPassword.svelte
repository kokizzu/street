<script>
  //@ts-nocheck
  import { GuestResetPassword } from '../jsApi.GEN';
  import Growl from '../_components/Growl.svelte'

  let password = '';
  let pass2 = '';
  let growl2 = Growl;

  async function resetPassword() {
    if (password.length < 12) {
      alert('password must be at least 12 characters');
      return
    }
    if (password !== pass2) {
      alert('password confirmation does not match');
      return
    }
    const queryParam = window.location.href.split('?')[1];
    const qps = queryParam.split('&');
    let secretCode = '';
    let hash = '';
    for (let i = 0; i < qps.length; i++) {
      const [key, value] = qps[i].split('=');
      if (key === 'secretCode') secretCode = value;
      if (key === 'hash') hash = value;
    }
    let i = { secretCode, hash, password };
    await GuestResetPassword(i, function (o) {
      console.log(o);
      if (o.error) {
        alert(o.error);
        return
      }
      alert('password reset successful');
      window.location.href = '/';
    });
  }
</script>


<section class="reset_password_container">
  <div class="main_content">
    <h1>
      <i class="gg-lock" />
      <span>Reset Password</span>
    </h1>

    <div class="input_box">
      <label for="new_password">New Password</label>
      <input type="password" id="new_password" bind:value={password} />
    </div>

    <div class="input_box">
      <label for="confirm_password">Confirm New Password</label>
      <input type="password" bind:value={pass2} /><br />
    </div>

    <button on:click={resetPassword}>Reset Password</button>
  </div>
</section>

<style>
  .reset_password_container {
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
    display: flex;
    flex-direction: row;
    align-content: center;
    justify-content: center;
    align-items: center;
  }
  .main_content h1 i {
    margin-right: 15px;
  }
  .main_content .input_box {
    display: flex;
    flex-direction: column;
    width: 100%;
    margin-top: 10px;
  }

  .main_content .input_box label {
    font-size: 13px;
    font-weight: 700;
    margin-left: 10px;
    margin-bottom: 8px;
  }

  .main_content .input_box input {
    width: 100%;
    border: 1px solid #cbd5e1;
    background-color: #f1f5f9;
    border-radius: 8px;
    padding: 12px;
  }

  .main_content .input_box input:focus {
    border-color: #3b82f6;
    outline: 1px solid #3b82f6;
  }

  .main_content button {
    padding: 12px 0;
    font-size: 16px;
    border: none;
    background-color: #6366f1;
    border-radius: 8px;
    color: white;
    cursor: pointer;
    text-decoration: none;
    width: 100%;
    margin: 10px auto 0 auto;
  }
  .main_content button:hover {
    background-color: #7e80f1;
  }
</style>
