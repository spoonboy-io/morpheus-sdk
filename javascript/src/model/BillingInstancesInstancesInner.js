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
import BillingInstancesInstancesInnerContainersInner from './BillingInstancesInstancesInnerContainersInner';

/**
 * The BillingInstancesInstancesInner model module.
 * @module model/BillingInstancesInstancesInner
 * @version 6.1.1
 */
class BillingInstancesInstancesInner {
    /**
     * Constructs a new <code>BillingInstancesInstancesInner</code>.
     * @alias module:model/BillingInstancesInstancesInner
     */
    constructor() { 
        
        BillingInstancesInstancesInner.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BillingInstancesInstancesInner</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BillingInstancesInstancesInner} obj Optional instance to populate.
     * @return {module:model/BillingInstancesInstancesInner} The populated <code>BillingInstancesInstancesInner</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BillingInstancesInstancesInner();

            if (data.hasOwnProperty('instanceId')) {
                obj['instanceId'] = ApiClient.convertToType(data['instanceId'], 'Number');
            }
            if (data.hasOwnProperty('instanceUUID')) {
                obj['instanceUUID'] = ApiClient.convertToType(data['instanceUUID'], 'String');
            }
            if (data.hasOwnProperty('startDate')) {
                obj['startDate'] = ApiClient.convertToType(data['startDate'], 'Date');
            }
            if (data.hasOwnProperty('endDate')) {
                obj['endDate'] = ApiClient.convertToType(data['endDate'], 'Date');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('price')) {
                obj['price'] = ApiClient.convertToType(data['price'], 'Number');
            }
            if (data.hasOwnProperty('cost')) {
                obj['cost'] = ApiClient.convertToType(data['cost'], 'Number');
            }
            if (data.hasOwnProperty('currency')) {
                obj['currency'] = ApiClient.convertToType(data['currency'], 'String');
            }
            if (data.hasOwnProperty('containers')) {
                obj['containers'] = ApiClient.convertToType(data['containers'], [BillingInstancesInstancesInnerContainersInner]);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>BillingInstancesInstancesInner</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>BillingInstancesInstancesInner</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['instanceUUID'] && !(typeof data['instanceUUID'] === 'string' || data['instanceUUID'] instanceof String)) {
            throw new Error("Expected the field `instanceUUID` to be a primitive type in the JSON string but got " + data['instanceUUID']);
        }
        // ensure the json data is a string
        if (data['name'] && !(typeof data['name'] === 'string' || data['name'] instanceof String)) {
            throw new Error("Expected the field `name` to be a primitive type in the JSON string but got " + data['name']);
        }
        // ensure the json data is a string
        if (data['currency'] && !(typeof data['currency'] === 'string' || data['currency'] instanceof String)) {
            throw new Error("Expected the field `currency` to be a primitive type in the JSON string but got " + data['currency']);
        }
        if (data['containers']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['containers'])) {
                throw new Error("Expected the field `containers` to be an array in the JSON data but got " + data['containers']);
            }
            // validate the optional field `containers` (array)
            for (const item of data['containers']) {
                BillingInstancesInstancesInnerContainersInner.validateJSON(item);
            };
        }

        return true;
    }


}



/**
 * @member {Number} instanceId
 */
BillingInstancesInstancesInner.prototype['instanceId'] = undefined;

/**
 * @member {String} instanceUUID
 */
BillingInstancesInstancesInner.prototype['instanceUUID'] = undefined;

/**
 * @member {Date} startDate
 */
BillingInstancesInstancesInner.prototype['startDate'] = undefined;

/**
 * @member {Date} endDate
 */
BillingInstancesInstancesInner.prototype['endDate'] = undefined;

/**
 * @member {String} name
 */
BillingInstancesInstancesInner.prototype['name'] = undefined;

/**
 * @member {Number} price
 */
BillingInstancesInstancesInner.prototype['price'] = undefined;

/**
 * @member {Number} cost
 */
BillingInstancesInstancesInner.prototype['cost'] = undefined;

/**
 * @member {String} currency
 */
BillingInstancesInstancesInner.prototype['currency'] = undefined;

/**
 * @member {Array.<module:model/BillingInstancesInstancesInnerContainersInner>} containers
 */
BillingInstancesInstancesInner.prototype['containers'] = undefined;






export default BillingInstancesInstancesInner;

