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
import ApiAccountsAccountIdGroupsGroup from './ApiAccountsAccountIdGroupsGroup';

/**
 * The InlineObject240 model module.
 * @module model/InlineObject240
 * @version 6.2.1
 */
class InlineObject240 {
    /**
     * Constructs a new <code>InlineObject240</code>.
     * @alias module:model/InlineObject240
     * @param group {module:model/ApiAccountsAccountIdGroupsGroup} 
     */
    constructor(group) { 
        
        InlineObject240.initialize(this, group);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, group) { 
        obj['group'] = group;
    }

    /**
     * Constructs a <code>InlineObject240</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject240} obj Optional instance to populate.
     * @return {module:model/InlineObject240} The populated <code>InlineObject240</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject240();

            if (data.hasOwnProperty('group')) {
                obj['group'] = ApiAccountsAccountIdGroupsGroup.constructFromObject(data['group']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiAccountsAccountIdGroupsGroup} group
 */
InlineObject240.prototype['group'] = undefined;






export default InlineObject240;
