function datetime( unixSec, humanize ) {
  if( !unixSec ) return '';
  if( typeof unixSec==='string' ) return unixSec; // might not be unix time
  let dt = new Date( unixSec * 1000 );
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
  
  const formattedDate = `${day}, ${date} ${month} ${year} - ${hours}:${minutes}:${seconds}`;
  return formattedDate;
}

function formatPrice( price, currency ) {
  return new Intl.NumberFormat( 'en-US', {
    style: 'currency',
    currency: currency,
    maximumSignificantDigits: 3,
  } ).format( price );
}

module.exports = {
  datetime: datetime,
  priceNtd: priceNtd,
  localeDatetime: localeDatetime,
  datetime2: datetime2,
  formatPrice: formatPrice,
};