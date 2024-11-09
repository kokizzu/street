/**
 * @description Format datetime
 * @param {number|string|any} unixSec 
 * @param {any} humanize 
 * @returns {string}
 */
function datetime( unixSec, humanize ) {
  if (!unixSec) return '';

  let dt = /** @type {Date} */ (new Date( unixSec * 1000 ));

  if(typeof unixSec === 'string') {
    dt = new Date(unixSec);
  }

  if( !humanize ) {
    const newDt = /** @type {string} */ (dt.toISOString());
    return newDt.substring( 0, 10 ) + ' ' + newDt.substring( 11, 16 );
  }
  
  return dt.toLocaleDateString('default', {
    day: '2-digit',
    month: 'long',
    year: 'numeric'
  });
}

function priceNtd( val ) {
  return (new Number( val )).toLocaleString( 'zh-TW' );
}

/**
 * @description Format to local datetime
 * @param {number|string} unixSec 
 * @returns {string}
 */
function localeDatetime( unixSec ) {
  if( !unixSec ) return '';

  const dt = /** @type {Date} */ (new Date(Number(unixSec) * 1000));
  const day = /** @type {string} */ (dt.toLocaleDateString('default', {
    weekday: 'long'
  }));
  const date = /** @type {number} */ (dt.getDate());
  const month = /** @type {string} */ (dt.toLocaleDateString('default', {
    month: 'long'
  }));
  const year = /** @type {number} */ (dt.getFullYear());

  let hh = /** @type {string} */ (String(dt.getHours()));
  if (Number(hh) < 10) hh = '0' + hh;

  let mm = /** @type {string} */ (String(dt.getMinutes()));
  if (Number(mm) < 10) mm = '0' + mm;

  return `${day}, ${date} ${month} ${year} ${hh}:${mm}`;
}

/**
 * @description Format to local datetime (with seconds)
 * @param {number|string} unixSec 
 * @returns {string}
 */
function datetime2( unixSec ) {
  if (!unixSec) return '-';
  
  const dt = /** @type {Date} */ (new Date(Number(unixSec) * 1000 ));
  const year = /** @type {number} */ (dt.getFullYear());
  const month = /** @type {string} */ (dt.toLocaleDateString('default', {
    month: 'long'
  }));
  const day = /** @type {string} */ (dt.toLocaleDateString('default', {
    weekday: 'long'
  }));
  const date = /** @type {string} */ (String(dt.getDate()).padStart(2, '0'));
  const hours = /** @type {string} */ (String(dt.getHours()).padStart(2, '0'));
  const minutes = /** @type {string} */ (String(dt.getMinutes()).padStart(2, '0'));
  const seconds = /** @type {string} */ (String(dt.getSeconds()).padStart(2, '0'));

  return `${day}, ${date} ${month} ${year} - ${hours}:${minutes}:${seconds}`;
}

/**
 * @description Format price
 * @param {number | bigint | Intl.StringNumericLiteral} price 
 * @param {string} currency 
 * @returns {string}
 */
function formatPrice(price, currency) {
  try {
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: currency,
      maximumSignificantDigits: 3,
    }).format(price);
  } catch(err) {
    console.log('formatPrice failed', err, price, currency);
    return price + ' ' + currency;
  }
}

function getApprovalState(/** @type {string} */ s) {
  if (s == '') {
    return 'approved'
  } else if (s.startsWith('pending')) {
    return 'pending'
  } else {
    return 'rejected'
  }
}

function M2ToPing(/** @type {number} */ sizeM2 ) {
  const value = sizeM2 / 3.30579;
  const minifiedValue = value.toFixed( 2 );
  return parseFloat( minifiedValue );
}

function dateISOFormat(/** @type number */ dayTo = 0) {
  const dt = new Date();
  dt.setDate(dt.getDate() + dayTo);

  const date = String(dt.getDate()).padStart(2, '0');
  const month = String(dt.getMonth() + 1).padStart(2, '0');
  const year = dt.getFullYear();

  return `${year}-${month}-${date}`;
}

module.exports = {
  datetime: datetime,
  priceNtd: priceNtd,
  localeDatetime: localeDatetime,
  datetime2: datetime2,
  formatPrice: formatPrice,
  getApprovalState: getApprovalState,
  M2ToPing: M2ToPing,
  dateISOFormat: dateISOFormat
};