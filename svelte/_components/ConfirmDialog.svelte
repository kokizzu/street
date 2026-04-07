<script>
  
  /**
   * @typedef {Object} Props
   * @property {boolean} [show]
   * @property {string} [title]
   * @property {string} [content]
   * @property {string} [modifier]
   * @property {string} [confirmText]
   * @property {any} [extraButtons] - {label:'', onClick: function(){}, modifier:'danger'}
   * @property {any} [onConfirm]
   * @property {any} [onCancel]
   * @property {import('svelte').Snippet} [children]
   */

  /** @type {Props} */
  let {
    show = $bindable(false),
    title = $bindable(''),
    content = $bindable(''),
    modifier = $bindable('primary'),
    confirmText = $bindable('Ok'),
    extraButtons = $bindable([]),
    onConfirm = $bindable(function() {}),
    onCancel = function() {},
    children
  } = $props();
  
  function okPressed() {
    show = false;
    onConfirm();
  }
  
  function cancelPressed() {
    show = false;
    onCancel();
  }
  
  export function closeDialog() {
    show = false;
    title = '';
    content = '';
  }
  
  export function showDialog( opts ) {
    show = true;
    title = opts.title || '';
    content = opts.content || '';
    confirmText = opts.confirmText || 'Ok';
    modifier = opts.modifier || 'primary';
    extraButtons = opts.extraButtons || [];
    onConfirm = opts.onConfirm || function() {};
  }
  
  function closeCallback( onClick ) {
    onClick = onClick || function() {};
    return function() {
      closeDialog();
      onClick();
    };
  }

</script>
{#if show}
  <div class='backdrop'></div>
  <div class='modal'>
    <div class='modal_background'></div>
    <div class='modal_content box -small' style='outline: none;'>
      <button tabindex='0' aria-label='close' title='Close' type='button' class='smallBtn close' onclick={cancelPressed}>
        <i class='gg-cross'></i>
      </button>
      <div class='title'>{title}</div>
      <p>
        {@html content}
        {@render children?.()}
      </p>
      <div class='layout-h'>
        <button tabindex='0' type='button' class="button {modifier || 'primary'}" onclick={okPressed}>{confirmText}</button>
        <button tabindex='0' type='button' class='button secondary' onclick={cancelPressed}>Cancel</button>
        {#each extraButtons as prop}
          <button tabindex='0' type='button' class='button {prop.modifier}' onclick={closeCallback(prop.onClick)}>{prop.label}</button>
        {/each}
      </div>
    </div>
  </div>
{/if}
<style>
    .backdrop {
        position   : fixed;
        top        : 0;
        left       : 0;
        background : grey;
        width      : 100%;
        height     : 100%;
        opacity    : .8;
    }

    .modal {
        background-color : white;
        position         : fixed;
        top              : 50%;
        left             : 50%;
        width            : 600px;
        margin-left      : -300px;
        margin-top       : -120px;
        overflow-y       : auto;
        box-shadow       : .25rem .25rem .5rem #DADADA, -.25rem -.25rem .5rem #FFF;
        border-radius    : 1em;
        padding          : 1em;
        opacity          : 1;
    }

    .title {
        font-weight : bold;
        font-size   : 1.2em;
    }

    .close {
        float : right;
    }

</style>