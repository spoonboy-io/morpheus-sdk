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
import BillingInstancesPrices from './BillingInstancesPrices';

/**
 * The BillingInstancesApplicablePrices model module.
 * @module model/BillingInstancesApplicablePrices
 * @version 6.2.1
 */
class BillingInstancesApplicablePrices {
    /**
     * Constructs a new <code>BillingInstancesApplicablePrices</code>.
     * @alias module:model/BillingInstancesApplicablePrices
     */
    constructor() { 
        
        BillingInstancesApplicablePrices.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BillingInstancesApplicablePrices</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BillingInstancesApplicablePrices} obj Optional instance to populate.
     * @return {module:model/BillingInstancesApplicablePrices} The populated <code>BillingInstancesApplicablePrices</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BillingInstancesApplicablePrices();

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
                obj['prices'] = ApiClient.convertToType(data['prices'], [BillingInstancesPrices]);
            }
        }
        return obj;
    }


}

/**
 * @member {Date} startDate
 */
BillingInstancesApplicablePrices.prototype['startDate'] = undefined;

/**
 * @member {Date} endDate
 */
BillingInstancesApplicablePrices.prototype['endDate'] = undefined;

/**
 * @member {Number} numUnits
 */
BillingInstancesApplicablePrices.prototype['numUnits'] = undefined;

/**
 * @member {Number} cost
 */
BillingInstancesApplicablePrices.prototype['cost'] = undefined;

/**
 * @member {Number} price
 */
BillingInstancesApplicablePrices.prototype['price'] = undefined;

/**
 * @member {String} currency
 */
BillingInstancesApplicablePrices.prototype['currency'] = undefined;

/**
 * @member {Array.<module:model/BillingInstancesPrices>} prices
 */
BillingInstancesApplicablePrices.prototype['prices'] = undefined;






export default BillingInstancesApplicablePrices;
