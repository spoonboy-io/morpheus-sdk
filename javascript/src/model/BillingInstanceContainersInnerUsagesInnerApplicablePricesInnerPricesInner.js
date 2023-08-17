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

/**
 * The BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner model module.
 * @module model/BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner
 * @version 6.1.1
 */
class BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner {
    /**
     * Constructs a new <code>BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner</code>.
     * @alias module:model/BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner
     */
    constructor() { 
        
        BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner} obj Optional instance to populate.
     * @return {module:model/BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner} The populated <code>BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner();

            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('pricePerUnit')) {
                obj['pricePerUnit'] = ApiClient.convertToType(data['pricePerUnit'], 'Number');
            }
            if (data.hasOwnProperty('costPerUnit')) {
                obj['costPerUnit'] = ApiClient.convertToType(data['costPerUnit'], 'Number');
            }
            if (data.hasOwnProperty('cost')) {
                obj['cost'] = ApiClient.convertToType(data['cost'], 'Number');
            }
            if (data.hasOwnProperty('price')) {
                obj['price'] = ApiClient.convertToType(data['price'], 'Number');
            }
            if (data.hasOwnProperty('quantity')) {
                obj['quantity'] = ApiClient.convertToType(data['quantity'], 'Number');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['type'] && !(typeof data['type'] === 'string' || data['type'] instanceof String)) {
            throw new Error("Expected the field `type` to be a primitive type in the JSON string but got " + data['type']);
        }

        return true;
    }


}



/**
 * @member {String} type
 */
BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner.prototype['type'] = undefined;

/**
 * @member {Number} pricePerUnit
 */
BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner.prototype['pricePerUnit'] = undefined;

/**
 * @member {Number} costPerUnit
 */
BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner.prototype['costPerUnit'] = undefined;

/**
 * @member {Number} cost
 */
BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner.prototype['cost'] = undefined;

/**
 * @member {Number} price
 */
BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner.prototype['price'] = undefined;

/**
 * @member {Number} quantity
 */
BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner.prototype['quantity'] = undefined;






export default BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner;
