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

/**
 * The BillingServersPrices model module.
 * @module model/BillingServersPrices
 * @version 6.2.1
 */
class BillingServersPrices {
    /**
     * Constructs a new <code>BillingServersPrices</code>.
     * @alias module:model/BillingServersPrices
     */
    constructor() { 
        
        BillingServersPrices.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BillingServersPrices</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BillingServersPrices} obj Optional instance to populate.
     * @return {module:model/BillingServersPrices} The populated <code>BillingServersPrices</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BillingServersPrices();

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


}

/**
 * @member {String} type
 */
BillingServersPrices.prototype['type'] = undefined;

/**
 * @member {Number} pricePerUnit
 */
BillingServersPrices.prototype['pricePerUnit'] = undefined;

/**
 * @member {Number} costPerUnit
 */
BillingServersPrices.prototype['costPerUnit'] = undefined;

/**
 * @member {Number} cost
 */
BillingServersPrices.prototype['cost'] = undefined;

/**
 * @member {Number} price
 */
BillingServersPrices.prototype['price'] = undefined;

/**
 * @member {Number} quantity
 */
BillingServersPrices.prototype['quantity'] = undefined;






export default BillingServersPrices;

