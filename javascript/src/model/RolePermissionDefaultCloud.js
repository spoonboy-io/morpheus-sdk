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
 * The RolePermissionDefaultCloud model module.
 * @module model/RolePermissionDefaultCloud
 * @version 6.2.1
 */
class RolePermissionDefaultCloud {
    /**
     * Constructs a new <code>RolePermissionDefaultCloud</code>.
     * @alias module:model/RolePermissionDefaultCloud
     * @param permissionCode {module:model/RolePermissionDefaultCloud.PermissionCodeEnum} `ComputeZone` is the code for Default Cloud Access
     * @param access {module:model/RolePermissionDefaultCloud.AccessEnum} The new access level.
     */
    constructor(permissionCode, access) { 
        
        RolePermissionDefaultCloud.initialize(this, permissionCode, access);
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
     * Constructs a <code>RolePermissionDefaultCloud</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/RolePermissionDefaultCloud} obj Optional instance to populate.
     * @return {module:model/RolePermissionDefaultCloud} The populated <code>RolePermissionDefaultCloud</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RolePermissionDefaultCloud();

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
 * `ComputeZone` is the code for Default Cloud Access
 * @member {module:model/RolePermissionDefaultCloud.PermissionCodeEnum} permissionCode
 */
RolePermissionDefaultCloud.prototype['permissionCode'] = undefined;

/**
 * The new access level.
 * @member {module:model/RolePermissionDefaultCloud.AccessEnum} access
 */
RolePermissionDefaultCloud.prototype['access'] = undefined;





/**
 * Allowed values for the <code>permissionCode</code> property.
 * @enum {String}
 * @readonly
 */
RolePermissionDefaultCloud['PermissionCodeEnum'] = {

    /**
     * value: "ComputeZone"
     * @const
     */
    "ComputeZone": "ComputeZone"
};


/**
 * Allowed values for the <code>access</code> property.
 * @enum {String}
 * @readonly
 */
RolePermissionDefaultCloud['AccessEnum'] = {

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



export default RolePermissionDefaultCloud;
