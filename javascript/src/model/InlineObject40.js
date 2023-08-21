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
 * The InlineObject40 model module.
 * @module model/InlineObject40
 * @version 6.2.1
 */
class InlineObject40 {
    /**
     * Constructs a new <code>InlineObject40</code>.
     * @alias module:model/InlineObject40
     * @param muted {Boolean} Set to false to unmute
     */
    constructor(muted) { 
        
        InlineObject40.initialize(this, muted);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, muted) { 
        obj['muted'] = muted || true;
    }

    /**
     * Constructs a <code>InlineObject40</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject40} obj Optional instance to populate.
     * @return {module:model/InlineObject40} The populated <code>InlineObject40</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject40();

            if (data.hasOwnProperty('muted')) {
                obj['muted'] = ApiClient.convertToType(data['muted'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * Set to false to unmute
 * @member {Boolean} muted
 * @default true
 */
InlineObject40.prototype['muted'] = true;






export default InlineObject40;

