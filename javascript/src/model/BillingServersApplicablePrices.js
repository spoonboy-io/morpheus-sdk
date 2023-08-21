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
import BillingServersPrices from './BillingServersPrices';

/**
 * The BillingServersApplicablePrices model module.
 * @module model/BillingServersApplicablePrices
 * @version 6.2.1
 */
class BillingServersApplicablePrices {
    /**
     * Constructs a new <code>BillingServersApplicablePrices</code>.
     * @alias module:model/BillingServersApplicablePrices
     */
    constructor() { 
        
        BillingServersApplicablePrices.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BillingServersApplicablePrices</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BillingServersApplicablePrices} obj Optional instance to populate.
     * @return {module:model/BillingServersApplicablePrices} The populated <code>BillingServersApplicablePrices</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BillingServersApplicablePrices();

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
                obj['prices'] = ApiClient.convertToType(data['prices'], [BillingServersPrices]);
            }
        }
        return obj;
    }


}

/**
 * @member {Date} startDate
 */
BillingServersApplicablePrices.prototype['startDate'] = undefined;

/**
 * @member {Date} endDate
 */
BillingServersApplicablePrices.prototype['endDate'] = undefined;

/**
 * @member {Number} numUnits
 */
BillingServersApplicablePrices.prototype['numUnits'] = undefined;

/**
 * @member {Number} cost
 */
BillingServersApplicablePrices.prototype['cost'] = undefined;

/**
 * @member {Number} price
 */
BillingServersApplicablePrices.prototype['price'] = undefined;

/**
 * @member {String} currency
 */
BillingServersApplicablePrices.prototype['currency'] = undefined;

/**
 * @member {Array.<module:model/BillingServersPrices>} prices
 */
BillingServersApplicablePrices.prototype['prices'] = undefined;






export default BillingServersApplicablePrices;
