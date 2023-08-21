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
 * The RolePermissionInstanceTypeAll model module.
 * @module model/RolePermissionInstanceTypeAll
 * @version 6.2.1
 */
class RolePermissionInstanceTypeAll {
    /**
     * Constructs a new <code>RolePermissionInstanceTypeAll</code>.
     * @alias module:model/RolePermissionInstanceTypeAll
     * @param access {module:model/RolePermissionInstanceTypeAll.AccessEnum} The new access level.
     */
    constructor(access) { 
        
        RolePermissionInstanceTypeAll.initialize(this, access);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, access) { 
        obj['access'] = access;
    }

    /**
     * Constructs a <code>RolePermissionInstanceTypeAll</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/RolePermissionInstanceTypeAll} obj Optional instance to populate.
     * @return {module:model/RolePermissionInstanceTypeAll} The populated <code>RolePermissionInstanceTypeAll</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RolePermissionInstanceTypeAll();

            if (data.hasOwnProperty('allInstanceTypes')) {
                obj['allInstanceTypes'] = ApiClient.convertToType(data['allInstanceTypes'], 'Boolean');
            }
            if (data.hasOwnProperty('access')) {
                obj['access'] = ApiClient.convertToType(data['access'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Apply to all instance types
 * @member {Boolean} allInstanceTypes
 */
RolePermissionInstanceTypeAll.prototype['allInstanceTypes'] = undefined;

/**
 * The new access level.
 * @member {module:model/RolePermissionInstanceTypeAll.AccessEnum} access
 */
RolePermissionInstanceTypeAll.prototype['access'] = undefined;





/**
 * Allowed values for the <code>access</code> property.
 * @enum {String}
 * @readonly
 */
RolePermissionInstanceTypeAll['AccessEnum'] = {

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



export default RolePermissionInstanceTypeAll;

