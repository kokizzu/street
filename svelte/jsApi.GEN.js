
import axios from "axios";
/**
 * @typedef {Object} GuestLoginIn
 * @property {String} Email
 * @property {String} Password
 */
const GuestLoginIn = {
  Email: '', // string
  Password: '', // string
}
/**
 * @typedef {Object} GuestLoginOut
 * @property {number} User.Id
 * @property {String} User.Email
 * @property {String} User.Password
 * @property {number} User.CreatedAt
 * @property {number} User.CreatedBy
 * @property {number} User.UpdatedAt
 * @property {number} User.UpdatedBy
 * @property {number} User.DeletedAt
 * @property {number} User.PasswordSetAt
 * @property {String} User.SecretCode
 * @property {number} User.SecretCodeAt
 * @property {number} User.VerificationSentAt
 * @property {number} User.VerifiedAt
 * @property {number} User.LastLoginAt
 */
const GuestLoginOut = {
  User: { // rqAuth.Users
    Id: 0, // uint64
    Email: '', // string
    Password: '', // string
    CreatedAt: 0, // int64
    CreatedBy: 0, // uint64
    UpdatedAt: 0, // int64
    UpdatedBy: 0, // uint64
    DeletedAt: 0, // int64
    PasswordSetAt: 0, // int64
    SecretCode: '', // string
    SecretCodeAt: 0, // int64
    VerificationSentAt: 0, // int64
    VerifiedAt: 0, // int64
    LastLoginAt: 0, // int64
  }, // rqAuth.Users
}
/**
 * @callback GuestLoginCallback
 * @param {GuestLoginOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestLoginIn} i
 * @param {GuestLoginCallback} cb
 * @returns {Promise}
 */
async function GuestLogin( i, cb ) {
  return await axios.post( '/guest/login', i ).then( cb )
}

/**
 * @typedef {Object} GuestRegisterIn
 * @property {String} Email
 * @property {String} Password
 */
const GuestRegisterIn = {
  Email: '', // string
  Password: '', // string
}
/**
 * @typedef {Object} GuestRegisterOut
 * @property {number} User.Id
 * @property {String} User.Email
 * @property {String} User.Password
 * @property {number} User.CreatedAt
 * @property {number} User.CreatedBy
 * @property {number} User.UpdatedAt
 * @property {number} User.UpdatedBy
 * @property {number} User.DeletedAt
 * @property {number} User.PasswordSetAt
 * @property {String} User.SecretCode
 * @property {number} User.SecretCodeAt
 * @property {number} User.VerificationSentAt
 * @property {number} User.VerifiedAt
 * @property {number} User.LastLoginAt
 */
const GuestRegisterOut = {
  User: { // rqAuth.Users
    Id: 0, // uint64
    Email: '', // string
    Password: '', // string
    CreatedAt: 0, // int64
    CreatedBy: 0, // uint64
    UpdatedAt: 0, // int64
    UpdatedBy: 0, // uint64
    DeletedAt: 0, // int64
    PasswordSetAt: 0, // int64
    SecretCode: '', // string
    SecretCodeAt: 0, // int64
    VerificationSentAt: 0, // int64
    VerifiedAt: 0, // int64
    LastLoginAt: 0, // int64
  }, // rqAuth.Users
}
/**
 * @callback GuestRegisterCallback
 * @param {GuestRegisterOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestRegisterIn} i
 * @param {GuestRegisterCallback} cb
 * @returns {Promise}
 */
async function GuestRegister( i, cb ) {
  return await axios.post( '/guest/register', i ).then( cb )
}

