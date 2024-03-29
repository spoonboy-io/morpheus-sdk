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
 * The InstanceTypeCreatePriceSets model module.
 * @module model/InstanceTypeCreatePriceSets
 * @version 6.2.1
 */
class InstanceTypeCreatePriceSets {
    /**
     * Constructs a new <code>InstanceTypeCreatePriceSets</code>.
     * @alias module:model/InstanceTypeCreatePriceSets
     * @param id {Number} Price Set ID
     */
    constructor(id) { 
        
        InstanceTypeCreatePriceSets.initialize(this, id);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, id) { 
        obj['id'] = id;
    }

    /**
     * Constructs a <code>InstanceTypeCreatePriceSets</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceTypeCreatePriceSets} obj Optional instance to populate.
     * @return {module:model/InstanceTypeCreatePriceSets} The populated <code>InstanceTypeCreatePriceSets</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceTypeCreatePriceSets();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * Price Set ID
 * @member {Number} id
 */
InstanceTypeCreatePriceSets.prototype['id'] = undefined;






export default InstanceTypeCreatePriceSets;

