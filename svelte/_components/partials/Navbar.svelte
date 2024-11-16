<script>
  import { Icon } from '../../node_modules/svelte-icons-pack/dist';
  import { RiUserFacesUserLine, RiCommunicationChatQuoteLine } from '../../node_modules/svelte-icons-pack/dist/ri';
  import { UserSendFeedback } from '../../jsApi.GEN';
  import { notifier } from '../notifier';

  export let username = /** @type {string} */ ('unknown');

  let isShowFeedBackPopUp = /** @type {boolean} */ (false);
  let feedbackMessage = /** @type {string} */ ('');

  async function sendFeedback() {
    if (!feedbackMessage) {
      notifier.showWarning('Please enter your feedback');
      return;
    }
    const inUserSendFeedback = /** @type {import('../../jsApi.GEN').UserSendFeedbackIn} */ ({
      userMessage: feedbackMessage
    })
    await UserSendFeedback(inUserSendFeedback,
      async function (/** @type {any} */ res) {
        if (res.error) {
          notifier.showError(res.error || 'Failed to send feedback');
          return;
        }

        notifier.showSuccess('Feedback sent');
        isShowFeedBackPopUp = false;
      }
    )
  }
</script>

{#if isShowFeedBackPopUp}
  <div class="backdrop">
    <div class="feedback-dialog">
      <h3>Send Feedback</h3>
      <textarea
        bind:value={feedbackMessage}
        name="feedback"
        id="feedback"
        cols="30"
        rows="10"
        placeholder="Your feedback here"
      />
      <div class="buttons">
        <button class="cancel-btn" on:click={() => {
          isShowFeedBackPopUp = false
        }}>Cancel</button>
        <button
          class="submit-btn"
          on:click={sendFeedback}
        >Submit</button>
      </div>
    </div>
  </div>
{/if}

<header>
  <div class="label">
    <a href="/" class="logo">
      <h2>HapSTR</h2>
    </a>
  </div>
  <div class="info">
    <button on:click={() => isShowFeedBackPopUp = true} class="feedback-btn">
      <span>Feedback</span>
      <Icon
        src={RiCommunicationChatQuoteLine}
        size="15"
        color="#FFF"
      />
    </button>
    <a href="/user" class="user">
      <span>{username}</span>
      <Icon src={RiUserFacesUserLine} size="15" />
    </a>
  </div>
</header>

<style>
  .backdrop {
    position: fixed;
    z-index: 2001;
    background: rgba(41, 41, 41, 0.5);
    width: 100%;
    left: 0;
    top: 0;
    right: 0;
    bottom: 0;
    height: auto;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .feedback-dialog {
    padding: 20px;
    border-radius: 10px;
    background-color: #FFF;
    display: flex;
    flex-direction: column;
    gap: 20px;
    color: #475569;
    width: 600px;
  }

  .feedback-dialog h3 {
    margin: 0 0 0 20px;
    font-size: 18px;
  }

  .feedback-dialog textarea {
    border: 1px solid var(--gray-002);
    background-color: var(--gray-001);
    border-radius: 8px;
    padding: 12px;
    resize: none;
  }

  .feedback-dialog textarea:focus {
    border-color: var(--orange-005);
    outline: 1px solid var(--orange-005);
  }

  .feedback-dialog .buttons {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    align-items: stretch;
    font-weight: 500;
    width: 100%;
    gap: 10px;
  }

  .feedback-dialog .buttons button {
    border-radius: 999px;
    border: none;
    padding: 10px;
    cursor: pointer;
    width: fit-content;
    padding: 12px 20px;
    text-transform: capitalize;
  }

  .feedback-dialog .buttons .cancel-btn {
    border: none;
    background-color: transparent;
  }

  .feedback-dialog .buttons .cancel-btn:hover {
    background-color: var(--orange-transparent);
    color: var(--orange-005);
  }

  .feedback-dialog .buttons .submit-btn {
    background-color: var(--orange-006);
    color: #FFF;
  }

  .feedback-dialog .buttons .submit-btn:hover {
    background-color: var(--orange-005);
  }

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

  header .info {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 10px;
  }

  header .info .feedback-btn {
    display: flex;
    flex-direction: row;
    gap: 8px;
    align-items: center;
    justify-content: center;
    padding: 10px 20px;
    border: none;
    border-radius: 9999px;
    background-color: var(--blue-006);
    color: #FFF;
    cursor: pointer;
  }

  header .info .feedback-btn:hover {
    background-color: var(--blue-005);
  }

  header .info a.user {
    display: flex;
    align-items: center;
    flex-direction: row;
    justify-content: center;
    gap: 6px;
    text-decoration: none;
    color: var(--gray-008);
    padding: 10px 20px;
    border-radius: 9999px;
    border: 1px solid var(--gray-002);
  }

  header .info a.user:hover {
    background-color: var(--gray-001);
  }
</style>