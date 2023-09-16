function toRadians(value) {
    return value * Math.PI / 180;
}

function distanceKM(x1, y1, x2, y2) {

    var R = 6371.0710; // kilometer
    var rlat1 = toRadians(x1); // Convert degrees to radians
    var rlat2 = toRadians(x2); // Convert degrees to radians
    var difflat = rlat2 - rlat1; // Radian difference (latitudes)
    var difflon = toRadians(y2 - y1); // Radian difference (longitudes)
    return 2 * R * Math.asin(Math.sqrt(Math.sin(difflat / 2) * Math.sin(difflat / 2) + Math.cos(rlat1) * Math.cos(rlat2) * Math.sin(difflon / 2) * Math.sin(difflon / 2)));
}

module.exports = {
  distanceKM: distanceKM,
  toRadians: toRadians,
};