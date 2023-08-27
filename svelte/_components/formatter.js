function datetime(unixSec) {
    if (!unixSec) return '';
    const dt = new Date(unixSec * 1000);
    const options = {day: '2-digit', month: 'long', year: 'numeric'};
    const formattedDate = dt.toLocaleDateString(undefined, options);
    // return dt.substring( 0, 10 ) + ' ' + dt.substring( 11, 16 );
    return formattedDate;
}

function priceNtd(val) {
    return (new Number(val * 1000)).toLocaleString('zh-TW');
}

module.exports = {
    datetime: datetime,
    priceNtd: priceNtd,
};