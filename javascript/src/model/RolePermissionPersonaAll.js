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
 * The RolePermissionPersonaAll model module.
 * @module model/RolePermissionPersonaAll
 * @version 6.2.1
 */
class RolePermissionPersonaAll {
    /**
     * Constructs a new <code>RolePermissionPersonaAll</code>.
     * @alias module:model/RolePermissionPersonaAll
     * @param allPersonas {Boolean} Apply to all personas
     * @param access {module:model/RolePermissionPersonaAll.AccessEnum} The new access level.
     */
    constructor(allPersonas, access) { 
        
        RolePermissionPersonaAll.initialize(this, allPersonas, access);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, allPersonas, access) { 
        obj['allPersonas'] = allPersonas;
        obj['access'] = access;
    }

    /**
     * Constructs a <code>RolePermissionPersonaAll</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/RolePermissionPersonaAll} obj Optional instance to populate.
     * @return {module:model/RolePermissionPersonaAll} The populated <code>RolePermissionPersonaAll</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RolePermissionPersonaAll();

            if (data.hasOwnProperty('allPersonas')) {
                obj['allPersonas'] = ApiClient.convertToType(data['allPersonas'], 'Boolean');
            }
            if (data.hasOwnProperty('access')) {
                obj['access'] = ApiClient.convertToType(data['access'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Apply to all personas
 * @member {Boolean} allPersonas
 */
RolePermissionPersonaAll.prototype['allPersonas'] = undefined;

/**
 * The new access level.
 * @member {module:model/RolePermissionPersonaAll.AccessEnum} access
 */
RolePermissionPersonaAll.prototype['access'] = undefined;





/**
 * Allowed values for the <code>access</code> property.
 * @enum {String}
 * @readonly
 */
RolePermissionPersonaAll['AccessEnum'] = {

    /**
     * value: "default"
     * @const
     */
    "default": "default",

    /**
     * value: "full"
     * @const
     */
    "full": "full",

    /**
     * value: "none"
     * @const
     */
    "none": "none"
};



export default RolePermissionPersonaAll;
