<script>
  export let text = '';
  export let subtext = '';
  export let modifier = '';
  
  export function hidingAfter( sec ) {
    return function() {
      hideAfter( sec )
    }
  }
  
  export function hideAfter( sec ) {
    if( !sec ) sec = 5;
    setTimeout( hideNotif, sec * 1000 )
  }
  
  export function hideNotif() {
    text = ''
    subtext = ''
    modifier = ''
  }
  
  export function updateSubtext( newSubtext, newModifier ) {
    subtext = newSubtext
    modifier = newModifier || '';
  }
  
  export function updateOk(status) {
    subtext = status || 'OK'
    modifier = 'primary'
  }
  
  export function updatingErr( err ) {
    updateErr( err );
  }
  
  export function updateErr( err ) {
    console.log( err )
    modifier = 'danger'
    const resp = err.response;
    let lastErr = '';
    if( resp ) {
      if( resp.data ) {
        if( resp.data.error ) {
          lastErr = resp.data.error
        } else {
          lastErr = JSON.stringify( resp.data )
        }
      }
    } else if( err.request ) {
      lastErr = 'error sending request: ' + err;
    }
    if( !lastErr ) {
      lastErr = err
    }
    subtext = lastErr
  }
  
  export function showNotif( newText, newSubtext ) {
    text = newText;
    subtext = newSubtext || '';
  }
</script>
{#if text}
	<div class="wrapper">
		<div class="no-data float">
			<p><b>{@html text}</b> <i class={modifier}>{@html subtext}</i></p>
		</div>
	</div>
{/if}
<style>
    .wrapper {
        position : absolute;
    }

    .no-data {
        background-color       : #FFF;
        color                  : #707070;
        padding                : 18px;
        box-shadow             : 0 10px 20px rgba(0, 0, 0, 0.19), 0 6px 6px rgba(0, 0, 0, 0.23);
        text-rendering         : optimizeLegibility;
        -webkit-font-smoothing : antialiased;
        font-size              : 12px;
        line-height            : 16px;
        display                : block;
        font-family            : roboto;
        border-radius          : 1em;
    }


    .no-data p {
        margin-top : 0;
    }

    .float {
        animation-name                    : float;
        -webkit-animation-name            : float;
        animation-duration                : 1.5s;
        -webkit-animation-duration        : 1.5s;
        animation-iteration-count         : infinite;
        -webkit-animation-iteration-count : infinite;
    }

    @-webkit-keyframes float {
        0% {
            -webkit-transform : translateY(0%);
        }
        50% {
            -webkit-transform : translateY(8%);
        }
        100% {
            -webkit-transform : translateY(0%);
        }
    }

    @-moz-keyframes float {
        0% {
            transform : translateY(0%);
        }
        50% {
            transform : translateY(8%);
        }
        100% {
            transform : translateY(0%);
        }
    }

    @-webkit-keyframes float {
        0% {
            transform : translateY(0%);
        }
        50% {
            transform : translateY(8%);
        }
        100% {
            transform : translateY(0%);
        }
    }

    @-o-keyframes float {
        0% {
            transform : translateY(0%);
        }
        50% {
            transform : translateY(8%);
        }
        100% {
            transform : translateY(0%);
        }
    }

    @keyframes float {
        0% {
            transform : translateY(0%);
        }
        50% {
            transform : translateY(8%);
        }
        100% {
            transform : translateY(0%);
        }
    }
</style>