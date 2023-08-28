function datetime( unixSec ) {
  if( !unixSec ) return '';
  const dt = new Date( unixSec * 1000 );
  const options = {day: '2-digit', month: 'long', year: 'numeric'};
  const formattedDate = dt.toLocaleDateString( undefined, options );
  // return dt.substring( 0, 10 ) + ' ' + dt.substring( 11, 16 );
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

module.exports = {
  datetime: datetime,
  priceNtd: priceNtd,
  localeDatetime: localeDatetime,
};