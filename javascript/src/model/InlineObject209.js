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
import ApiRolesIdRole from './ApiRolesIdRole';

/**
 * The InlineObject209 model module.
 * @module model/InlineObject209
 * @version 6.2.1
 */
class InlineObject209 {
    /**
     * Constructs a new <code>InlineObject209</code>.
     * @alias module:model/InlineObject209
     * @param role {module:model/ApiRolesIdRole} 
     */
    constructor(role) { 
        
        InlineObject209.initialize(this, role);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, role) { 
        obj['role'] = role;
    }

    /**
     * Constructs a <code>InlineObject209</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject209} obj Optional instance to populate.
     * @return {module:model/InlineObject209} The populated <code>InlineObject209</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject209();

            if (data.hasOwnProperty('role')) {
                obj['role'] = ApiRolesIdRole.constructFromObject(data['role']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiRolesIdRole} role
 */
InlineObject209.prototype['role'] = undefined;






export default InlineObject209;

