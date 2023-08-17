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
 * The BillingZonesInnerVirtualImages model module.
 * @module model/BillingZonesInnerVirtualImages
 * @version 6.1.1
 */
class BillingZonesInnerVirtualImages {
    /**
     * Constructs a new <code>BillingZonesInnerVirtualImages</code>.
     * @alias module:model/BillingZonesInnerVirtualImages
     */
    constructor() { 
        
        BillingZonesInnerVirtualImages.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BillingZonesInnerVirtualImages</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BillingZonesInnerVirtualImages} obj Optional instance to populate.
     * @return {module:model/BillingZonesInnerVirtualImages} The populated <code>BillingZonesInnerVirtualImages</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BillingZonesInnerVirtualImages();

            if (data.hasOwnProperty('price')) {
                obj['price'] = ApiClient.convertToType(data['price'], 'Number');
            }
            if (data.hasOwnProperty('cost')) {
                obj['cost'] = ApiClient.convertToType(data['cost'], 'Number');
            }
            if (data.hasOwnProperty('virtualImages')) {
                obj['virtualImages'] = ApiClient.convertToType(data['virtualImages'], [Object]);
            }
            if (data.hasOwnProperty('count')) {
                obj['count'] = ApiClient.convertToType(data['count'], 'Number');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>BillingZonesInnerVirtualImages</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>BillingZonesInnerVirtualImages</code>.
     */
    static validateJSON(data) {
        // ensure the json data is an array
        if (!Array.isArray(data['virtualImages'])) {
            throw new Error("Expected the field `virtualImages` to be an array in the JSON data but got " + data['virtualImages']);
        }

        return true;
    }


}



/**
 * @member {Number} price
 */
BillingZonesInnerVirtualImages.prototype['price'] = undefined;

/**
 * @member {Number} cost
 */
BillingZonesInnerVirtualImages.prototype['cost'] = undefined;

/**
 * @member {Array.<Object>} virtualImages
 */
BillingZonesInnerVirtualImages.prototype['virtualImages'] = undefined;

/**
 * @member {Number} count
 */
BillingZonesInnerVirtualImages.prototype['count'] = undefined;






export default BillingZonesInnerVirtualImages;

