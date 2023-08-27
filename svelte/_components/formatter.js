function datetime( unixSec ) {
  if( !unixSec ) return '-';
  const dt = new Date( unixSec * 1000 )
  const options = { day: '2-digit', month: 'long', year: 'numeric' };
  const formattedDate = dt.toLocaleDateString(undefined, options);
  // return dt.substring( 0, 10 ) + ' ' + dt.substring( 11, 16 );
  return formattedDate
}

function datetime2( unixSec ) {
  if( !unixSec ) return '-';

  const dt = new Date(unixSec * 1000);
  const year = dt.getFullYear();
  const month = dt.toLocaleDateString('default', { month: 'long'}); // String(dt.getMonth() + 1).padStart(2, '0');
  const day = dt.toLocaleDateString('default', {weekday: 'long'});
  const date = String(dt.getDate()).padStart(2, '0');
  const hours = String(dt.getHours()).padStart(2, '0');
  const minutes = String(dt.getMinutes()).padStart(2, '0');
  const seconds = String(dt.getSeconds()).padStart(2, '0');

  const formattedDate = `${day}, ${date} ${month} ${year} - ${hours}:${minutes}:${seconds}`;
  return formattedDate
}

module.exports = {
  datetime: datetime,
  datetime2: datetime2
};