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
 * The RolePermissionDefaultBlueprint model module.
 * @module model/RolePermissionDefaultBlueprint
 * @version 6.2.1
 */
class RolePermissionDefaultBlueprint {
    /**
     * Constructs a new <code>RolePermissionDefaultBlueprint</code>.
     * @alias module:model/RolePermissionDefaultBlueprint
     * @param permissionCode {module:model/RolePermissionDefaultBlueprint.PermissionCodeEnum} `AppTemplate` is the code for Default Blueprint Access
     * @param access {module:model/RolePermissionDefaultBlueprint.AccessEnum} The new access level.
     */
    constructor(permissionCode, access) { 
        
        RolePermissionDefaultBlueprint.initialize(this, permissionCode, access);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, permissionCode, access) { 
        obj['permissionCode'] = permissionCode;
        obj['access'] = access;
    }

    /**
     * Constructs a <code>RolePermissionDefaultBlueprint</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/RolePermissionDefaultBlueprint} obj Optional instance to populate.
     * @return {module:model/RolePermissionDefaultBlueprint} The populated <code>RolePermissionDefaultBlueprint</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RolePermissionDefaultBlueprint();

            if (data.hasOwnProperty('permissionCode')) {
                obj['permissionCode'] = ApiClient.convertToType(data['permissionCode'], 'String');
            }
            if (data.hasOwnProperty('access')) {
                obj['access'] = ApiClient.convertToType(data['access'], 'String');
            }
        }
        return obj;
    }


}

/**
 * `AppTemplate` is the code for Default Blueprint Access
 * @member {module:model/RolePermissionDefaultBlueprint.PermissionCodeEnum} permissionCode
 */
RolePermissionDefaultBlueprint.prototype['permissionCode'] = undefined;

/**
 * The new access level.
 * @member {module:model/RolePermissionDefaultBlueprint.AccessEnum} access
 */
RolePermissionDefaultBlueprint.prototype['access'] = undefined;





/**
 * Allowed values for the <code>permissionCode</code> property.
 * @enum {String}
 * @readonly
 */
RolePermissionDefaultBlueprint['PermissionCodeEnum'] = {

    /**
     * value: "AppTemplate"
     * @const
     */
    "AppTemplate": "AppTemplate"
};


/**
 * Allowed values for the <code>access</code> property.
 * @enum {String}
 * @readonly
 */
RolePermissionDefaultBlueprint['AccessEnum'] = {

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



export default RolePermissionDefaultBlueprint;

