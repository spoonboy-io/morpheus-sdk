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
import ResourcePoolGroupsCreateInput from './ResourcePoolGroupsCreateInput';

/**
 * The InlineObject207 model module.
 * @module model/InlineObject207
 * @version 6.2.1
 */
class InlineObject207 {
    /**
     * Constructs a new <code>InlineObject207</code>.
     * @alias module:model/InlineObject207
     */
    constructor() { 
        
        InlineObject207.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineObject207</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject207} obj Optional instance to populate.
     * @return {module:model/InlineObject207} The populated <code>InlineObject207</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject207();

            if (data.hasOwnProperty('resourcePoolGroup')) {
                obj['resourcePoolGroup'] = ResourcePoolGroupsCreateInput.constructFromObject(data['resourcePoolGroup']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ResourcePoolGroupsCreateInput} resourcePoolGroup
 */
InlineObject207.prototype['resourcePoolGroup'] = undefined;






export default InlineObject207;

