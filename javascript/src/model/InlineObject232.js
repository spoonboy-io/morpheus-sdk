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
import ApiStorageBucketsStorageBucket from './ApiStorageBucketsStorageBucket';

/**
 * The InlineObject232 model module.
 * @module model/InlineObject232
 * @version 6.2.1
 */
class InlineObject232 {
    /**
     * Constructs a new <code>InlineObject232</code>.
     * @alias module:model/InlineObject232
     * @param storageBucket {module:model/ApiStorageBucketsStorageBucket} 
     */
    constructor(storageBucket) { 
        
        InlineObject232.initialize(this, storageBucket);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, storageBucket) { 
        obj['storageBucket'] = storageBucket;
    }

    /**
     * Constructs a <code>InlineObject232</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject232} obj Optional instance to populate.
     * @return {module:model/InlineObject232} The populated <code>InlineObject232</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject232();

            if (data.hasOwnProperty('storageBucket')) {
                obj['storageBucket'] = ApiStorageBucketsStorageBucket.constructFromObject(data['storageBucket']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiStorageBucketsStorageBucket} storageBucket
 */
InlineObject232.prototype['storageBucket'] = undefined;






export default InlineObject232;

