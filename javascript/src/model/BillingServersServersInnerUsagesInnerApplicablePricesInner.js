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
import BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner from './BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner';

/**
 * The BillingServersServersInnerUsagesInnerApplicablePricesInner model module.
 * @module model/BillingServersServersInnerUsagesInnerApplicablePricesInner
 * @version 6.1.1
 */
class BillingServersServersInnerUsagesInnerApplicablePricesInner {
    /**
     * Constructs a new <code>BillingServersServersInnerUsagesInnerApplicablePricesInner</code>.
     * @alias module:model/BillingServersServersInnerUsagesInnerApplicablePricesInner
     */
    constructor() { 
        
        BillingServersServersInnerUsagesInnerApplicablePricesInner.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BillingServersServersInnerUsagesInnerApplicablePricesInner</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BillingServersServersInnerUsagesInnerApplicablePricesInner} obj Optional instance to populate.
     * @return {module:model/BillingServersServersInnerUsagesInnerApplicablePricesInner} The populated <code>BillingServersServersInnerUsagesInnerApplicablePricesInner</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BillingServersServersInnerUsagesInnerApplicablePricesInner();

            if (data.hasOwnProperty('startDate')) {
                obj['startDate'] = ApiClient.convertToType(data['startDate'], 'Date');
            }
            if (data.hasOwnProperty('endDate')) {
                obj['endDate'] = ApiClient.convertToType(data['endDate'], 'Date');
            }
            if (data.hasOwnProperty('numUnits')) {
                obj['numUnits'] = ApiClient.convertToType(data['numUnits'], 'Number');
            }
            if (data.hasOwnProperty('cost')) {
                obj['cost'] = ApiClient.convertToType(data['cost'], 'Number');
            }
            if (data.hasOwnProperty('price')) {
                obj['price'] = ApiClient.convertToType(data['price'], 'Number');
            }
            if (data.hasOwnProperty('currency')) {
                obj['currency'] = ApiClient.convertToType(data['currency'], 'String');
            }
            if (data.hasOwnProperty('prices')) {
                obj['prices'] = ApiClient.convertToType(data['prices'], [BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner]);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>BillingServersServersInnerUsagesInnerApplicablePricesInner</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>BillingServersServersInnerUsagesInnerApplicablePricesInner</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['currency'] && !(typeof data['currency'] === 'string' || data['currency'] instanceof String)) {
            throw new Error("Expected the field `currency` to be a primitive type in the JSON string but got " + data['currency']);
        }
        if (data['prices']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['prices'])) {
                throw new Error("Expected the field `prices` to be an array in the JSON data but got " + data['prices']);
            }
            // validate the optional field `prices` (array)
            for (const item of data['prices']) {
                BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner.validateJSON(item);
            };
        }

        return true;
    }


}



/**
 * @member {Date} startDate
 */
BillingServersServersInnerUsagesInnerApplicablePricesInner.prototype['startDate'] = undefined;

/**
 * @member {Date} endDate
 */
BillingServersServersInnerUsagesInnerApplicablePricesInner.prototype['endDate'] = undefined;

/**
 * @member {Number} numUnits
 */
BillingServersServersInnerUsagesInnerApplicablePricesInner.prototype['numUnits'] = undefined;

/**
 * @member {Number} cost
 */
BillingServersServersInnerUsagesInnerApplicablePricesInner.prototype['cost'] = undefined;

/**
 * @member {Number} price
 */
BillingServersServersInnerUsagesInnerApplicablePricesInner.prototype['price'] = undefined;

/**
 * @member {String} currency
 */
BillingServersServersInnerUsagesInnerApplicablePricesInner.prototype['currency'] = undefined;

/**
 * @member {Array.<module:model/BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner>} prices
 */
BillingServersServersInnerUsagesInnerApplicablePricesInner.prototype['prices'] = undefined;






export default BillingServersServersInnerUsagesInnerApplicablePricesInner;
