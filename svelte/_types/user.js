/**
 * @typedef {Object} User
 * @property {string | number} id
 * @property {string} email
 * @property {string} password
 * @property {number} createdAt
 * @property {number} createdBy
 * @property {number} updatedAt
 * @property {number} updatedBy
 * @property {number} deletedAt
 * @property {number} passwordSetAt
 * @property {string} secretCode
 * @property {number} secretCodeAt
 * @property {number} verificationSentAt
 * @property {number} verifiedAt
 * @property {number} lastLoginAt
 * @property {string} fullName
 * @property {string} userName
 * @property {string} country
 * @property {string} language
 */
module.exports = {}

/**
 * @typedef {Object} Session
 * @property {string} sessionToken
 * @property {number | string} userId
 * @property {number} expiredAt
 * @property {string} device
 * @property {number} loginAt
 * @property {string} loginIPs
 */
module.exports = {}

/**
 * @typedef {Object} MostLoggedInUser
 * @property {string} time_period
 * @property {string} email
 * @property {string} full_name
 * @property {number} total
 */
module.exports = {}