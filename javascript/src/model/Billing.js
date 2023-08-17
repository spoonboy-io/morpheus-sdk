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
import BillingZonesInner from './BillingZonesInner';

/**
 * The Billing model module.
 * @module model/Billing
 * @version 6.1.1
 */
class Billing {
    /**
     * Constructs a new <code>Billing</code>.
     * @alias module:model/Billing
     */
    constructor() { 
        
        Billing.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Billing</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Billing} obj Optional instance to populate.
     * @return {module:model/Billing} The populated <code>Billing</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Billing();

            if (data.hasOwnProperty('accountId')) {
                obj['accountId'] = ApiClient.convertToType(data['accountId'], 'Number');
            }
            if (data.hasOwnProperty('accountUUID')) {
                obj['accountUUID'] = ApiClient.convertToType(data['accountUUID'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('startDate')) {
                obj['startDate'] = ApiClient.convertToType(data['startDate'], 'Date');
            }
            if (data.hasOwnProperty('endDate')) {
                obj['endDate'] = ApiClient.convertToType(data['endDate'], 'Date');
            }
            if (data.hasOwnProperty('priceUnit')) {
                obj['priceUnit'] = ApiClient.convertToType(data['priceUnit'], 'String');
            }
            if (data.hasOwnProperty('price')) {
                obj['price'] = ApiClient.convertToType(data['price'], 'Number');
            }
            if (data.hasOwnProperty('cost')) {
                obj['cost'] = ApiClient.convertToType(data['cost'], 'Number');
            }
            if (data.hasOwnProperty('zones')) {
                obj['zones'] = ApiClient.convertToType(data['zones'], [BillingZonesInner]);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>Billing</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>Billing</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['accountUUID'] && !(typeof data['accountUUID'] === 'string' || data['accountUUID'] instanceof String)) {
            throw new Error("Expected the field `accountUUID` to be a primitive type in the JSON string but got " + data['accountUUID']);
        }
        // ensure the json data is a string
        if (data['name'] && !(typeof data['name'] === 'string' || data['name'] instanceof String)) {
            throw new Error("Expected the field `name` to be a primitive type in the JSON string but got " + data['name']);
        }
        // ensure the json data is a string
        if (data['priceUnit'] && !(typeof data['priceUnit'] === 'string' || data['priceUnit'] instanceof String)) {
            throw new Error("Expected the field `priceUnit` to be a primitive type in the JSON string but got " + data['priceUnit']);
        }
        if (data['zones']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['zones'])) {
                throw new Error("Expected the field `zones` to be an array in the JSON data but got " + data['zones']);
            }
            // validate the optional field `zones` (array)
            for (const item of data['zones']) {
                BillingZonesInner.validateJSON(item);
            };
        }

        return true;
    }


}



/**
 * @member {Number} accountId
 */
Billing.prototype['accountId'] = undefined;

/**
 * @member {String} accountUUID
 */
Billing.prototype['accountUUID'] = undefined;

/**
 * @member {String} name
 */
Billing.prototype['name'] = undefined;

/**
 * @member {Date} startDate
 */
Billing.prototype['startDate'] = undefined;

/**
 * @member {Date} endDate
 */
Billing.prototype['endDate'] = undefined;

/**
 * @member {String} priceUnit
 */
Billing.prototype['priceUnit'] = undefined;

/**
 * @member {Number} price
 */
Billing.prototype['price'] = undefined;

/**
 * @member {Number} cost
 */
Billing.prototype['cost'] = undefined;

/**
 * @member {Array.<module:model/BillingZonesInner>} zones
 */
Billing.prototype['zones'] = undefined;






export default Billing;

