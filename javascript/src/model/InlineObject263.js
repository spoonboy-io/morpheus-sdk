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
import VirtualImageCreate from './VirtualImageCreate';

/**
 * The InlineObject263 model module.
 * @module model/InlineObject263
 * @version 6.2.1
 */
class InlineObject263 {
    /**
     * Constructs a new <code>InlineObject263</code>.
     * @alias module:model/InlineObject263
     * @param virtualImage {module:model/VirtualImageCreate} 
     */
    constructor(virtualImage) { 
        
        InlineObject263.initialize(this, virtualImage);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, virtualImage) { 
        obj['virtualImage'] = virtualImage;
    }

    /**
     * Constructs a <code>InlineObject263</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject263} obj Optional instance to populate.
     * @return {module:model/InlineObject263} The populated <code>InlineObject263</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject263();

            if (data.hasOwnProperty('virtualImage')) {
                obj['virtualImage'] = VirtualImageCreate.constructFromObject(data['virtualImage']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/VirtualImageCreate} virtualImage
 */
InlineObject263.prototype['virtualImage'] = undefined;






export default InlineObject263;

