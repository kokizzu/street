function datetime( unixSec, humanize ) {
  if( !unixSec ) return '';
  let dt = new Date( unixSec * 1000 );
  if( typeof unixSec==='string' ) {
    dt = new Date( unixSec );
  }
  if( !humanize ) {
    dt = dt.toISOString();
    return dt.substring( 0, 10 ) + ' ' + dt.substring( 11, 16 );
  }
  const options = {day: '2-digit', month: 'long', year: 'numeric'};
  const formattedDate = dt.toLocaleDateString( undefined, options );
  return formattedDate;
}

function priceNtd( val ) {
  return (new Number( val )).toLocaleString( 'zh-TW' );
}

function localeDatetime( unixSec ) {
  if( !unixSec ) return '';
  const dt = new Date( unixSec * 1000 );
  const day = dt.toLocaleDateString( 'default', {weekday: 'long'} );
  const date = dt.getDate();
  const month = dt.toLocaleDateString( 'default', {month: 'long'} );
  const year = dt.getFullYear();
  let hh = dt.getHours();
  if( hh<10 ) hh = '0' + hh;
  let mm = dt.getMinutes();
  if( mm<10 ) mm = '0' + mm;
  const formattedDate = `${day}, ${date} ${month} ${year} ${hh}:${mm}`;
  return formattedDate;
}

function datetime2( unixSec ) {
  if( !unixSec ) return '-';
  
  const dt = new Date( unixSec * 1000 );
  const year = dt.getFullYear();
  const month = dt.toLocaleDateString( 'default', {month: 'long'} ); // String(dt.getMonth() + 1).padStart(2, '0');
  const day = dt.toLocaleDateString( 'default', {weekday: 'long'} );
  const date = String( dt.getDate() ).padStart( 2, '0' );
  const hours = String( dt.getHours() ).padStart( 2, '0' );
  const minutes = String( dt.getMinutes() ).padStart( 2, '0' );
  const seconds = String( dt.getSeconds() ).padStart( 2, '0' );

  return `${day}, ${date} ${month} ${year} - ${hours}:${minutes}:${seconds}`;
}

function formatPrice( price, currency ) {
  try {
    return new Intl.NumberFormat( 'en-US', {
      style: 'currency',
      currency: currency,
      maximumSignificantDigits: 3,
    } ).format( price );
  } catch( e ) {
    console.log( 'formatPrice failed', e, price, currency );
    return price + ' ' + currency;
  }
}

function getApprovalState(s) {
  if (s == '') {
    return 'approved'
  } else if ( s.startsWith('pending')) {
    return 'pending'
  } else {
    return 'rejected'
  }
}

function M2ToPing( sizeM2 ) {
  const value = sizeM2 / 3.30579;
  const minifiedValue = value.toFixed( 2 );
  return parseFloat( minifiedValue );
}

module.exports = {
  datetime: datetime,
  priceNtd: priceNtd,
  localeDatetime: localeDatetime,
  datetime2: datetime2,
  formatPrice: formatPrice,
  getApprovalState: getApprovalState,
  M2ToPing: M2ToPing
};