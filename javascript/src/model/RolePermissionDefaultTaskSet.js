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
 * The RolePermissionDefaultTaskSet model module.
 * @module model/RolePermissionDefaultTaskSet
 * @version 6.2.1
 */
class RolePermissionDefaultTaskSet {
    /**
     * Constructs a new <code>RolePermissionDefaultTaskSet</code>.
     * @alias module:model/RolePermissionDefaultTaskSet
     * @param permissionCode {module:model/RolePermissionDefaultTaskSet.PermissionCodeEnum} `TaskSet` is the code for Default Workflow Access
     * @param access {module:model/RolePermissionDefaultTaskSet.AccessEnum} The new access level.
     */
    constructor(permissionCode, access) { 
        
        RolePermissionDefaultTaskSet.initialize(this, permissionCode, access);
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
     * Constructs a <code>RolePermissionDefaultTaskSet</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/RolePermissionDefaultTaskSet} obj Optional instance to populate.
     * @return {module:model/RolePermissionDefaultTaskSet} The populated <code>RolePermissionDefaultTaskSet</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RolePermissionDefaultTaskSet();

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
 * `TaskSet` is the code for Default Workflow Access
 * @member {module:model/RolePermissionDefaultTaskSet.PermissionCodeEnum} permissionCode
 */
RolePermissionDefaultTaskSet.prototype['permissionCode'] = undefined;

/**
 * The new access level.
 * @member {module:model/RolePermissionDefaultTaskSet.AccessEnum} access
 */
RolePermissionDefaultTaskSet.prototype['access'] = undefined;





/**
 * Allowed values for the <code>permissionCode</code> property.
 * @enum {String}
 * @readonly
 */
RolePermissionDefaultTaskSet['PermissionCodeEnum'] = {

    /**
     * value: "TaskSet"
     * @const
     */
    "TaskSet": "TaskSet"
};


/**
 * Allowed values for the <code>access</code> property.
 * @enum {String}
 * @readonly
 */
RolePermissionDefaultTaskSet['AccessEnum'] = {

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



export default RolePermissionDefaultTaskSet;

