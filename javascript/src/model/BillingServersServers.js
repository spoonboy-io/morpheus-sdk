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
import BillingServersUsages from './BillingServersUsages';

/**
 * The BillingServersServers model module.
 * @module model/BillingServersServers
 * @version 6.2.1
 */
class BillingServersServers {
    /**
     * Constructs a new <code>BillingServersServers</code>.
     * @alias module:model/BillingServersServers
     */
    constructor() { 
        
        BillingServersServers.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BillingServersServers</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BillingServersServers} obj Optional instance to populate.
     * @return {module:model/BillingServersServers} The populated <code>BillingServersServers</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BillingServersServers();

            if (data.hasOwnProperty('refType')) {
                obj['refType'] = ApiClient.convertToType(data['refType'], 'String');
            }
            if (data.hasOwnProperty('refUUID')) {
                obj['refUUID'] = ApiClient.convertToType(data['refUUID'], 'String');
            }
            if (data.hasOwnProperty('refId')) {
                obj['refId'] = ApiClient.convertToType(data['refId'], 'Number');
            }
            if (data.hasOwnProperty('startDate')) {
                obj['startDate'] = ApiClient.convertToType(data['startDate'], 'Date');
            }
            if (data.hasOwnProperty('endDate')) {
                obj['endDate'] = ApiClient.convertToType(data['endDate'], 'Date');
            }
            if (data.hasOwnProperty('cost')) {
                obj['cost'] = ApiClient.convertToType(data['cost'], 'Number');
            }
            if (data.hasOwnProperty('price')) {
                obj['price'] = ApiClient.convertToType(data['price'], 'Number');
            }
            if (data.hasOwnProperty('numUnits')) {
                obj['numUnits'] = ApiClient.convertToType(data['numUnits'], 'Number');
            }
            if (data.hasOwnProperty('unit')) {
                obj['unit'] = ApiClient.convertToType(data['unit'], 'String');
            }
            if (data.hasOwnProperty('currency')) {
                obj['currency'] = ApiClient.convertToType(data['currency'], 'String');
            }
            if (data.hasOwnProperty('usages')) {
                obj['usages'] = ApiClient.convertToType(data['usages'], [BillingServersUsages]);
            }
            if (data.hasOwnProperty('numUsages')) {
                obj['numUsages'] = ApiClient.convertToType(data['numUsages'], 'Number');
            }
            if (data.hasOwnProperty('totalUsages')) {
                obj['totalUsages'] = ApiClient.convertToType(data['totalUsages'], 'Number');
            }
            if (data.hasOwnProperty('hasMoreUsages')) {
                obj['hasMoreUsages'] = ApiClient.convertToType(data['hasMoreUsages'], 'Boolean');
            }
            if (data.hasOwnProperty('foundPricing')) {
                obj['foundPricing'] = ApiClient.convertToType(data['foundPricing'], 'Boolean');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('serverUUID')) {
                obj['serverUUID'] = ApiClient.convertToType(data['serverUUID'], 'String');
            }
            if (data.hasOwnProperty('serverUniqueId')) {
                obj['serverUniqueId'] = ApiClient.convertToType(data['serverUniqueId'], 'String');
            }
            if (data.hasOwnProperty('serverExternalId')) {
                obj['serverExternalId'] = ApiClient.convertToType(data['serverExternalId'], 'String');
            }
            if (data.hasOwnProperty('serverInternalId')) {
                obj['serverInternalId'] = ApiClient.convertToType(data['serverInternalId'], 'String');
            }
            if (data.hasOwnProperty('resourcePoolId')) {
                obj['resourcePoolId'] = ApiClient.convertToType(data['resourcePoolId'], 'String');
            }
            if (data.hasOwnProperty('resourcePoolName')) {
                obj['resourcePoolName'] = ApiClient.convertToType(data['resourcePoolName'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} refType
 */
BillingServersServers.prototype['refType'] = undefined;

/**
 * @member {String} refUUID
 */
BillingServersServers.prototype['refUUID'] = undefined;

/**
 * @member {Number} refId
 */
BillingServersServers.prototype['refId'] = undefined;

/**
 * @member {Date} startDate
 */
BillingServersServers.prototype['startDate'] = undefined;

/**
 * @member {Date} endDate
 */
BillingServersServers.prototype['endDate'] = undefined;

/**
 * @member {Number} cost
 */
BillingServersServers.prototype['cost'] = undefined;

/**
 * @member {Number} price
 */
BillingServersServers.prototype['price'] = undefined;

/**
 * @member {Number} numUnits
 */
BillingServersServers.prototype['numUnits'] = undefined;

/**
 * @member {String} unit
 */
BillingServersServers.prototype['unit'] = undefined;

/**
 * @member {String} currency
 */
BillingServersServers.prototype['currency'] = undefined;

/**
 * @member {Array.<module:model/BillingServersUsages>} usages
 */
BillingServersServers.prototype['usages'] = undefined;

/**
 * @member {Number} numUsages
 */
BillingServersServers.prototype['numUsages'] = undefined;

/**
 * @member {Number} totalUsages
 */
BillingServersServers.prototype['totalUsages'] = undefined;

/**
 * @member {Boolean} hasMoreUsages
 */
BillingServersServers.prototype['hasMoreUsages'] = undefined;

/**
 * @member {Boolean} foundPricing
 */
BillingServersServers.prototype['foundPricing'] = undefined;

/**
 * @member {String} name
 */
BillingServersServers.prototype['name'] = undefined;

/**
 * @member {String} serverUUID
 */
BillingServersServers.prototype['serverUUID'] = undefined;

/**
 * @member {String} serverUniqueId
 */
BillingServersServers.prototype['serverUniqueId'] = undefined;

/**
 * @member {String} serverExternalId
 */
BillingServersServers.prototype['serverExternalId'] = undefined;

/**
 * @member {String} serverInternalId
 */
BillingServersServers.prototype['serverInternalId'] = undefined;

/**
 * @member {String} resourcePoolId
 */
BillingServersServers.prototype['resourcePoolId'] = undefined;

/**
 * @member {String} resourcePoolName
 */
BillingServersServers.prototype['resourcePoolName'] = undefined;






export default BillingServersServers;

