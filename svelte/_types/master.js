/**
 * @typedef {Object} Access
 * @property {boolean} admin
 * @property {boolean} buyer
 * @property {boolean} realtor
 * @property {boolean} user
 */
module.exports = {};

/**
 * @typedef {Object} PagerOut
 * @property {number} page
 * @property {number} perPage
 * @property {number} pages
 * @property {number} total
 * @property {Object.<string, string[]>} filters
 * @property {string[]} order
 */
module.exports = {};

/**
 * @typedef {Object} PagerIn
 * @property {number} page
 * @property {number} perPage
 * @property {Object.<string, string[]>} filters
 * @property {string[]} order
 */
module.exports = {}

/**
 * @typedef {Object} Field
 * @property {string} name
 * @property {string} label
 * @property {string} description
 * @property {string} type
 * @property {string} inputType
 * @property {boolean} readOnly
 * @property {string} mapping
 * @property {Object.<string, any>} validations
 * @property {string} refEndpoint
 */
module.exports = {};

/**
 * @typedef {Object} Coordinate
 * @property {string} lat
 * @property {string} lng
 */
module.exports = {};

/**
 * @typedef {Object} Currency
 * @property {string} name
 * @property {string} code
 */
module.exports = {};

/**
 * @typedef {Object} CountryData
 * @property {string} country
 * @property {string} iso_2
 * @property {string} iso_3
 * @property {string} country_code
 * @property {string} region
 * @property {string} region_code
 * @property {string} unit_measurement
 * @property {Coordinate} coordinate
 * @property {Currency} currency
 */
module.exports = {};