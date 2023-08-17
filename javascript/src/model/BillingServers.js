/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.1.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import BillingServersServersInner from './BillingServersServersInner';

/**
 * The BillingServers model module.
 * @module model/BillingServers
 * @version 6.1.1
 */
class BillingServers {
    /**
     * Constructs a new <code>BillingServers</code>.
     * @alias module:model/BillingServers
     */
    constructor() { 
        
        BillingServers.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BillingServers</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BillingServers} obj Optional instance to populate.
     * @return {module:model/BillingServers} The populated <code>BillingServers</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BillingServers();

            if (data.hasOwnProperty('price')) {
                obj['price'] = ApiClient.convertToType(data['price'], 'Number');
            }
            if (data.hasOwnProperty('cost')) {
                obj['cost'] = ApiClient.convertToType(data['cost'], 'Number');
            }
            if (data.hasOwnProperty('startDate')) {
                obj['startDate'] = ApiClient.convertToType(data['startDate'], 'Date');
            }
            if (data.hasOwnProperty('endDate')) {
                obj['endDate'] = ApiClient.convertToType(data['endDate'], 'Date');
            }
            if (data.hasOwnProperty('servers')) {
                obj['servers'] = ApiClient.convertToType(data['servers'], [BillingServersServersInner]);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>BillingServers</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>BillingServers</code>.
     */
    static validateJSON(data) {
        if (data['servers']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['servers'])) {
                throw new Error("Expected the field `servers` to be an array in the JSON data but got " + data['servers']);
            }
            // validate the optional field `servers` (array)
            for (const item of data['servers']) {
                BillingServersServersInner.validateJSON(item);
            };
        }

        return true;
    }


}



/**
 * @member {Number} price
 */
BillingServers.prototype['price'] = undefined;

/**
 * @member {Number} cost
 */
BillingServers.prototype['cost'] = undefined;

/**
 * @member {Date} startDate
 */
BillingServers.prototype['startDate'] = undefined;

/**
 * @member {Date} endDate
 */
BillingServers.prototype['endDate'] = undefined;

/**
 * @member {Array.<module:model/BillingServersServersInner>} servers
 */
BillingServers.prototype['servers'] = undefined;






export default BillingServers;
