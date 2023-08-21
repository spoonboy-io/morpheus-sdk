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
 * The RolePermissionGroupAll model module.
 * @module model/RolePermissionGroupAll
 * @version 6.2.1
 */
class RolePermissionGroupAll {
    /**
     * Constructs a new <code>RolePermissionGroupAll</code>.
     * @alias module:model/RolePermissionGroupAll
     * @param allGroups {Boolean} Apply to all groups (site)
     * @param access {module:model/RolePermissionGroupAll.AccessEnum} The new access level.
     */
    constructor(allGroups, access) { 
        
        RolePermissionGroupAll.initialize(this, allGroups, access);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, allGroups, access) { 
        obj['allGroups'] = allGroups;
        obj['access'] = access;
    }

    /**
     * Constructs a <code>RolePermissionGroupAll</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/RolePermissionGroupAll} obj Optional instance to populate.
     * @return {module:model/RolePermissionGroupAll} The populated <code>RolePermissionGroupAll</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RolePermissionGroupAll();

            if (data.hasOwnProperty('allGroups')) {
                obj['allGroups'] = ApiClient.convertToType(data['allGroups'], 'Boolean');
            }
            if (data.hasOwnProperty('access')) {
                obj['access'] = ApiClient.convertToType(data['access'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Apply to all groups (site)
 * @member {Boolean} allGroups
 */
RolePermissionGroupAll.prototype['allGroups'] = undefined;

/**
 * The new access level.
 * @member {module:model/RolePermissionGroupAll.AccessEnum} access
 */
RolePermissionGroupAll.prototype['access'] = undefined;





/**
 * Allowed values for the <code>access</code> property.
 * @enum {String}
 * @readonly
 */
RolePermissionGroupAll['AccessEnum'] = {

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
     * value: "read"
     * @const
     */
    "read": "read",

    /**
     * value: "none"
     * @const
     */
    "none": "none"
};



export default RolePermissionGroupAll;

