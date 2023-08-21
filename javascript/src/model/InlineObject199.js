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
import ApiPriceSetsIdPriceSet from './ApiPriceSetsIdPriceSet';

/**
 * The InlineObject199 model module.
 * @module model/InlineObject199
 * @version 6.2.1
 */
class InlineObject199 {
    /**
     * Constructs a new <code>InlineObject199</code>.
     * @alias module:model/InlineObject199
     * @param priceSet {module:model/ApiPriceSetsIdPriceSet} 
     */
    constructor(priceSet) { 
        
        InlineObject199.initialize(this, priceSet);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, priceSet) { 
        obj['priceSet'] = priceSet;
    }

    /**
     * Constructs a <code>InlineObject199</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject199} obj Optional instance to populate.
     * @return {module:model/InlineObject199} The populated <code>InlineObject199</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject199();

            if (data.hasOwnProperty('priceSet')) {
                obj['priceSet'] = ApiPriceSetsIdPriceSet.constructFromObject(data['priceSet']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiPriceSetsIdPriceSet} priceSet
 */
InlineObject199.prototype['priceSet'] = undefined;






export default InlineObject199;

