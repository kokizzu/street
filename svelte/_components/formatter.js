function datetime( unixSec ) {
  if( !unixSec ) return '';
  const dt = new Date( unixSec * 1000 ).toISOString();
  return dt.substring( 0, 10 ) + ' ' + dt.substring( 11, 16 );
}

module.exports = {
  datetime: datetime,
};