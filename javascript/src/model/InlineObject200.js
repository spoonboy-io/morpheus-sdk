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
import ApiPricesPrice from './ApiPricesPrice';

/**
 * The InlineObject200 model module.
 * @module model/InlineObject200
 * @version 6.2.1
 */
class InlineObject200 {
    /**
     * Constructs a new <code>InlineObject200</code>.
     * @alias module:model/InlineObject200
     * @param price {module:model/ApiPricesPrice} 
     */
    constructor(price) { 
        
        InlineObject200.initialize(this, price);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, price) { 
        obj['price'] = price;
    }

    /**
     * Constructs a <code>InlineObject200</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject200} obj Optional instance to populate.
     * @return {module:model/InlineObject200} The populated <code>InlineObject200</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject200();

            if (data.hasOwnProperty('price')) {
                obj['price'] = ApiPricesPrice.constructFromObject(data['price']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiPricesPrice} price
 */
InlineObject200.prototype['price'] = undefined;






export default InlineObject200;

