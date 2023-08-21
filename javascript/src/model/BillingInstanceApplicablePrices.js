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
import BillingInstancePrices from './BillingInstancePrices';

/**
 * The BillingInstanceApplicablePrices model module.
 * @module model/BillingInstanceApplicablePrices
 * @version 6.2.1
 */
class BillingInstanceApplicablePrices {
    /**
     * Constructs a new <code>BillingInstanceApplicablePrices</code>.
     * @alias module:model/BillingInstanceApplicablePrices
     */
    constructor() { 
        
        BillingInstanceApplicablePrices.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BillingInstanceApplicablePrices</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BillingInstanceApplicablePrices} obj Optional instance to populate.
     * @return {module:model/BillingInstanceApplicablePrices} The populated <code>BillingInstanceApplicablePrices</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BillingInstanceApplicablePrices();

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
                obj['prices'] = ApiClient.convertToType(data['prices'], [BillingInstancePrices]);
            }
        }
        return obj;
    }


}

/**
 * @member {Date} startDate
 */
BillingInstanceApplicablePrices.prototype['startDate'] = undefined;

/**
 * @member {Date} endDate
 */
BillingInstanceApplicablePrices.prototype['endDate'] = undefined;

/**
 * @member {Number} numUnits
 */
BillingInstanceApplicablePrices.prototype['numUnits'] = undefined;

/**
 * @member {Number} cost
 */
BillingInstanceApplicablePrices.prototype['cost'] = undefined;

/**
 * @member {Number} price
 */
BillingInstanceApplicablePrices.prototype['price'] = undefined;

/**
 * @member {String} currency
 */
BillingInstanceApplicablePrices.prototype['currency'] = undefined;

/**
 * @member {Array.<module:model/BillingInstancePrices>} prices
 */
BillingInstanceApplicablePrices.prototype['prices'] = undefined;






export default BillingInstanceApplicablePrices;
