/**
 * @typedef {[number, number]} Coord
 */

/**
 * @typedef {Object} Property
 * @property {string} id
 * @property {string} uniqPropKey
 * @property {string} serialNumber
 * @property {string} sizeM2
 * @property {string} mainUse
 * @property {string} mainBuildingMaterial
 * @property {string} constructCompletedDate
 * @property {string} numberOfFloors
 * @property {string} buildingLamination
 * @property {string} address
 * @property {string} district
 * @property {string} note
 * @property {Coord} coord
 * @property {number} createdAt
 * @property {string} createdBy
 * @property {number} updatedAt
 * @property {string} updatedBy
 * @property {number} deletedAt
 * @property {string} formattedAddress
 * @property {string} lastPrice
 * @property {any[]} priceHistoriesSell
 * @property {any[]} priceHistoriesRent
 * @property {string} purpose
 * @property {string} houseType
 * @property {string[]} images
 * @property {number} bedroom
 * @property {number} bathroom
 * @property {number} agencyFeePercent
 * @property {any[]} floorList
 * @property {string} version
 * @property {number} yearBuilt
 * @property {number} yearRenovated
 * @property {number} totalSqft
 * @property {string} countyName
 * @property {string} street
 * @property {string} city
 * @property {string} state
 * @property {string} zip
 * @property {number} propertyLastUpdatedDate
 * @property {string} approvalState
 * @property {string} countryCode
 * @property {number} livingroom
 * @property {number} altitude
 * @property {number} parking
 * @property {number} depositFee
 * @property {number} minimumDurationYear
 * @property {any[]} otherFees
 * @property {any[]} imageLabels
 */
module.exports = {};

/**
 * @typedef {Object} PropertyUS
 * @property {string} id
 * @property {string} uniqPropKey
 * @property {string} serialNumber
 * @property {string} sizeM2
 * @property {string} mainUse
 * @property {string} mainBuildingMaterial
 * @property {string} constructCompletedDate
 * @property {string} numberOfFloors
 * @property {string} buildingLamination
 * @property {string} address
 * @property {string} district
 * @property {string} note
 * @property {Coord} coord
 * @property {number} createdAt
 * @property {string} createdBy
 * @property {number} updatedAt
 * @property {string} updatedBy
 * @property {number} deletedAt
 * @property {string} formattedAddress
 * @property {string} lastPrice
 * @property {string[]} priceHistoriesSell
 * @property {any[]} priceHistoriesRent
 * @property {string} purpose
 * @property {string} houseType
 * @property {string[]} images
 * @property {number} bedroom
 * @property {number} bathroom
 * @property {number} agencyFeePercent
 * @property {any[]} floorList
 * @property {string} version
 * @property {number} yearBuilt
 * @property {number} yearRenovated
 * @property {number} totalSqft
 * @property {string} countyName
 * @property {string} street
 * @property {string} city
 * @property {string} state
 * @property {string} zip
 * @property {number} propertyLastUpdatedDate
 * @property {string} approvalState
 * @property {string} countryCode
 * @property {number} livingroom
 * @property {number} altitude
 * @property {number} parking
 * @property {number} depositFee
 * @property {number} minimumDurationYear
 * @property {any[]} otherFees
 * @property {any[]} imageLabels
 */
module.exports = {};

/**
 * @typedef {Object} PropertyExtraUS
 * @property {string} id
 * @property {string} propertyKey
 * @property {string} countyUrl
 * @property {boolean} countyIsActive
 * @property {string} zoneDataInfo
 * @property {string} taxInfo
 * @property {string} historyTaxInfo
 * @property {Record<string, any> | string} amenitySuperGroups
 * @property {Record<string, any> | string} mlsDisclaimerInfo
 * @property {Record<string, any> | string} facilityInfo
 * @property {string} riskInfo
 * @property {Record<string, any> | string} mediaSourceJson
 * @property {string} taxNote
 */
module.exports = {};

/**
 * @typedef {Object} MlsDisclaimerInfo
 * @property {string} ListingBrokerName
 * @property {string} ListingBrokerNumber
 * @property {string} ListingAgentName
 * @property {string} ListingAgentNumber
 */
module.exports = {};

/**
 * @typedef {Object} PropertyWithNote
 * @property {string} id
 * @property {string} uniqPropKey
 * @property {string} serialNumber
 * @property {string} sizeM2
 * @property {string} mainUse
 * @property {string} mainBuildingMaterial
 * @property {string} constructCompletedDate
 * @property {string} numberOfFloors
 * @property {string} buildingLamination
 * @property {string} address
 * @property {string} district
 * @property {string} note
 * @property {Coord} coord
 * @property {number} createdAt
 * @property {string} createdBy
 * @property {number} updatedAt
 * @property {string} updatedBy
 * @property {number} deletedAt
 * @property {string} formattedAddress
 * @property {string} lastPrice
 * @property {any[]} priceHistoriesSell
 * @property {any[]} priceHistoriesRent
 * @property {string} purpose
 * @property {string} houseType
 * @property {string[]} images
 * @property {number} bedroom
 * @property {number} bathroom
 * @property {number} agencyFeePercent
 * @property {any[]} floorList
 * @property {string} version
 * @property {number} yearBuilt
 * @property {number} yearRenovated
 * @property {number} totalSqft
 * @property {string} countyName
 * @property {string} street
 * @property {string} city
 * @property {string} state
 * @property {string} zip
 * @property {number} propertyLastUpdatedDate
 * @property {string} approvalState
 * @property {string} countryCode
 * @property {number} livingroom
 * @property {number} altitude
 * @property {number} parking
 * @property {number} depositFee
 * @property {number} minimumDurationYear
 * @property {any[]} otherFees
 * @property {any[]} imageLabels
 * @property {string} contactEmail
 * @property {string} contactPhone
 * @property {string} about
 * @property {string} image3dUrl
 * @property {number} lat?
 * @property {number} lng?
 * @property {number} distanceKM?
 * @property {boolean} liked?
 * @property {number} likeCount?
 */
module.exports = {};

/**
 * @typedef {Object} ScannedAreasToRender
 * @property {string} time_period
 * @property {number} views
 * @property {number} latitude
 * @property {number} longitude
 * @property {string} city
 * @property {string} state
 */
module.exports = {};

/**
 * @typedef {Object} ScannedPropertiesToRender
 * @property {string} time_period
 * @property {number} views
 * @property {string} price
 * @property {number} total_sqft
 * @property {string} city
 * @property {string} state
 * @property {string} address
 */
module.exports = {};