// notifier by Javier Sanahuja Liebana
// this require css on body, see for "/* notifier */" on _layout.html
let Notifier = function() {
  let a = {
    container: null, default_time: 4500, vanish_time: 300, fps: 30, success: {
      classes: 'alert-success',
      icon: '<svg xmlns="http://www.w3.org/2000/svg" width="8" height="8" viewBox="0 0 8 8"><path d="M6.41 0l-.69.72-2.78 2.78-.81-.78-.72-.72-1.41 1.41.72.72 1.5 1.5.69.72.72-.72 3.5-3.5.72-.72-1.44-1.41z" transform="translate(0 1)" /></svg>',
    }, error: {
      classes: 'alert-danger',
      icon: '<svg xmlns="http://www.w3.org/2000/svg" width="8" height="8" viewBox="0 0 8 8"><path d="M1.41 0l-1.41 1.41.72.72 1.78 1.81-1.78 1.78-.72.69 1.41 1.44.72-.72 1.81-1.81 1.78 1.81.69.72 1.44-1.44-.72-.69-1.81-1.78 1.81-1.81.72-.72-1.44-1.41-.69.72-1.78 1.78-1.81-1.78-.72-.72z" /></svg>',
    }, warning: {
      classes: 'alert-warning',
      icon: '<svg xmlns="http://www.w3.org/2000/svg" width="8" height="8" viewBox="0 0 8 8"><path d="M3.09 0c-.06 0-.1.04-.13.09l-2.94 6.81c-.02.05-.03.13-.03.19v.81c0 .05.04.09.09.09h6.81c.05 0 .09-.04.09-.09v-.81c0-.05-.01-.14-.03-.19l-2.94-6.81c-.02-.05-.07-.09-.13-.09h-.81zm-.09 3h1v2h-1v-2zm0 3h1v1h-1v-1z" /></svg>',
    }, info: {
      classes: 'alert-info',
      icon: '<svg xmlns="http://www.w3.org/2000/svg" width="8" height="8" viewBox="0 0 8 8"><path d="M3 0c-.55 0-1 .45-1 1s.45 1 1 1 1-.45 1-1-.45-1-1-1zm-1.5 2.5c-.83 0-1.5.67-1.5 1.5h1c0-.28.22-.5.5-.5s.5.22.5.5-1 1.64-1 2.5c0 .86.67 1.5 1.5 1.5s1.5-.67 1.5-1.5h-1c0 .28-.22.5-.5.5s-.5-.22-.5-.5c0-.36 1-1.84 1-2.5 0-.81-.67-1.5-1.5-1.5z" transform="translate(2)"/></svg>',
    },
  }, b = function( a, b, c, d, e, f, g ) {
    this.pushed = !1, this.element = document.createElement( 'div' ), this.element.className = 'notifyjs-notification ' + c.classes, this.element.innerHTML = '<div class=\'notifyjs-icon\'>' + c.icon + '</div>' + b, a.appendChild( this.element );// Notification progress
    let h = document.createElement( 'p' );
    h.className = 'progress', this.element.appendChild( h ), this.callback = g;
    let j = this;
    this.push = function() {
      if( !j.pushed ) {
        j.pushed = !0;
        let a = 0, b = 1e3 / f;
        j.element.style.display = 'block', j.interval = setInterval( function() {
          a++;
          let c = 100 * (1 - b * a / d);
          h.style.right = c + '%', 0>=c && ('function'== typeof g && g(), j.clear());
        }, b );
      }
    }, this.clear = function() {
      if( j.pushed ) {
        let a = 1e3 / f, b = 1;
        'undefined'!= typeof j.interval && clearInterval( j.interval ), j.interval = setInterval( function() {b -= 1 / (e / a), j.element.style.opacity = b, 0>=b && (clearInterval( j.interval ), j.destroy());}, a );
      }
    }, this.destroy = function() {j.pushed && (j.pushed = !1, 'undefined'!= typeof j.interval && clearInterval( j.interval ), a.removeChild( j.element ));};
  };
  return function Notifier( c ) {
    this.options = Object.assign( {}, a ),
      this.options = Object.assign( this.options, c ), null===this.options.container && (this.options.container = document.createElement( 'div' ), document.getElementsByTagName( 'body' )[ 0 ].appendChild( this.options.container )), this.options.container.className += ' notifyjs-container',
      this.notify = function( a, c, d, e ) {
        return 'undefined'== typeof this.options[ a ] ? void console.error( 'Notify.js: Error, undefined \'' + a + '\' notification type' ) : ('undefined'== typeof d && (d = this.options.default_time), new b( this.options.container, c, this.options[ a ], d, this.options.vanish_time, this.options.fps, e ));
      };
    this.showInfo = function( infoMsg ) {
      let notification = this.notify( 'info', infoMsg );
      notification.push();
    };
    this.showSuccess = function( okMsg ) {
      let notification = this.notify( 'success', okMsg );
      notification.push();
    };
    this.showError = function( errMsg ) {
      let notification = this.notify( 'error', errMsg );
      notification.push();
    };
    this.showWarning = function( warnMsg ) {
      let notification = this.notify( 'warning', warnMsg );
      notification.push();
    };
  };
}();

let notifier = new Notifier();

export {
  notifier,
};