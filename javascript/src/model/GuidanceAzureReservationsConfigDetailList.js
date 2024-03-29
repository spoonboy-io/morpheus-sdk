/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The GuidanceAzureReservationsConfigDetailList model module.
 * @module model/GuidanceAzureReservationsConfigDetailList
 * @version 6.2.1
 */
class GuidanceAzureReservationsConfigDetailList {
    /**
     * Constructs a new <code>GuidanceAzureReservationsConfigDetailList</code>.
     * @alias module:model/GuidanceAzureReservationsConfigDetailList
     */
    constructor() { 
        
        GuidanceAzureReservationsConfigDetailList.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>GuidanceAzureReservationsConfigDetailList</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/GuidanceAzureReservationsConfigDetailList} obj Optional instance to populate.
     * @return {module:model/GuidanceAzureReservationsConfigDetailList} The populated <code>GuidanceAzureReservationsConfigDetailList</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new GuidanceAzureReservationsConfigDetailList();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('apiName')) {
                obj['apiName'] = ApiClient.convertToType(data['apiName'], 'String');
            }
            if (data.hasOwnProperty('apiType')) {
                obj['apiType'] = ApiClient.convertToType(data['apiType'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('period')) {
                obj['period'] = ApiClient.convertToType(data['period'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('size')) {
                obj['size'] = ApiClient.convertToType(data['size'], 'String');
            }
            if (data.hasOwnProperty('region')) {
                obj['region'] = ApiClient.convertToType(data['region'], 'String');
            }
            if (data.hasOwnProperty('term')) {
                obj['term'] = ApiClient.convertToType(data['term'], 'String');
            }
            if (data.hasOwnProperty('meterId')) {
                obj['meterId'] = ApiClient.convertToType(data['meterId'], 'String');
            }
            if (data.hasOwnProperty('onDemandCount')) {
                obj['onDemandCount'] = ApiClient.convertToType(data['onDemandCount'], 'Number');
            }
            if (data.hasOwnProperty('onDemandCost')) {
                obj['onDemandCost'] = ApiClient.convertToType(data['onDemandCost'], 'Number');
            }
            if (data.hasOwnProperty('reservedCount')) {
                obj['reservedCount'] = ApiClient.convertToType(data['reservedCount'], 'Number');
            }
            if (data.hasOwnProperty('reservedCost')) {
                obj['reservedCost'] = ApiClient.convertToType(data['reservedCost'], 'Number');
            }
            if (data.hasOwnProperty('recommendedCount')) {
                obj['recommendedCount'] = ApiClient.convertToType(data['recommendedCount'], 'Number');
            }
            if (data.hasOwnProperty('recommendedCost')) {
                obj['recommendedCost'] = ApiClient.convertToType(data['recommendedCost'], 'Number');
            }
            if (data.hasOwnProperty('totalSavings')) {
                obj['totalSavings'] = ApiClient.convertToType(data['totalSavings'], 'Number');
            }
            if (data.hasOwnProperty('totalSavingsPercent')) {
                obj['totalSavingsPercent'] = ApiClient.convertToType(data['totalSavingsPercent'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {String} id
 */
GuidanceAzureReservationsConfigDetailList.prototype['id'] = undefined;

/**
 * @member {String} apiName
 */
GuidanceAzureReservationsConfigDetailList.prototype['apiName'] = undefined;

/**
 * @member {String} apiType
 */
GuidanceAzureReservationsConfigDetailList.prototype['apiType'] = undefined;

/**
 * @member {String} externalId
 */
GuidanceAzureReservationsConfigDetailList.prototype['externalId'] = undefined;

/**
 * @member {String} period
 */
GuidanceAzureReservationsConfigDetailList.prototype['period'] = undefined;

/**
 * @member {String} name
 */
GuidanceAzureReservationsConfigDetailList.prototype['name'] = undefined;

/**
 * @member {String} type
 */
GuidanceAzureReservationsConfigDetailList.prototype['type'] = undefined;

/**
 * @member {String} category
 */
GuidanceAzureReservationsConfigDetailList.prototype['category'] = undefined;

/**
 * @member {String} size
 */
GuidanceAzureReservationsConfigDetailList.prototype['size'] = undefined;

/**
 * @member {String} region
 */
GuidanceAzureReservationsConfigDetailList.prototype['region'] = undefined;

/**
 * @member {String} term
 */
GuidanceAzureReservationsConfigDetailList.prototype['term'] = undefined;

/**
 * @member {String} meterId
 */
GuidanceAzureReservationsConfigDetailList.prototype['meterId'] = undefined;

/**
 * @member {Number} onDemandCount
 */
GuidanceAzureReservationsConfigDetailList.prototype['onDemandCount'] = undefined;

/**
 * @member {Number} onDemandCost
 */
GuidanceAzureReservationsConfigDetailList.prototype['onDemandCost'] = undefined;

/**
 * @member {Number} reservedCount
 */
GuidanceAzureReservationsConfigDetailList.prototype['reservedCount'] = undefined;

/**
 * @member {Number} reservedCost
 */
GuidanceAzureReservationsConfigDetailList.prototype['reservedCost'] = undefined;

/**
 * @member {Number} recommendedCount
 */
GuidanceAzureReservationsConfigDetailList.prototype['recommendedCount'] = undefined;

/**
 * @member {Number} recommendedCost
 */
GuidanceAzureReservationsConfigDetailList.prototype['recommendedCost'] = undefined;

/**
 * @member {Number} totalSavings
 */
GuidanceAzureReservationsConfigDetailList.prototype['totalSavings'] = undefined;

/**
 * @member {Number} totalSavingsPercent
 */
GuidanceAzureReservationsConfigDetailList.prototype['totalSavingsPercent'] = undefined;






export default GuidanceAzureReservationsConfigDetailList;

