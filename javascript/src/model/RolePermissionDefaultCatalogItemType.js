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
 * The RolePermissionDefaultCatalogItemType model module.
 * @module model/RolePermissionDefaultCatalogItemType
 * @version 6.2.1
 */
class RolePermissionDefaultCatalogItemType {
    /**
     * Constructs a new <code>RolePermissionDefaultCatalogItemType</code>.
     * @alias module:model/RolePermissionDefaultCatalogItemType
     * @param permissionCode {module:model/RolePermissionDefaultCatalogItemType.PermissionCodeEnum} `CatalogItemType` is the code for Default Catalog Item Type Access
     * @param access {module:model/RolePermissionDefaultCatalogItemType.AccessEnum} The new access level.
     */
    constructor(permissionCode, access) { 
        
        RolePermissionDefaultCatalogItemType.initialize(this, permissionCode, access);
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
     * Constructs a <code>RolePermissionDefaultCatalogItemType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/RolePermissionDefaultCatalogItemType} obj Optional instance to populate.
     * @return {module:model/RolePermissionDefaultCatalogItemType} The populated <code>RolePermissionDefaultCatalogItemType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RolePermissionDefaultCatalogItemType();

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
 * `CatalogItemType` is the code for Default Catalog Item Type Access
 * @member {module:model/RolePermissionDefaultCatalogItemType.PermissionCodeEnum} permissionCode
 */
RolePermissionDefaultCatalogItemType.prototype['permissionCode'] = undefined;

/**
 * The new access level.
 * @member {module:model/RolePermissionDefaultCatalogItemType.AccessEnum} access
 */
RolePermissionDefaultCatalogItemType.prototype['access'] = undefined;





/**
 * Allowed values for the <code>permissionCode</code> property.
 * @enum {String}
 * @readonly
 */
RolePermissionDefaultCatalogItemType['PermissionCodeEnum'] = {

    /**
     * value: "CatalogItemType"
     * @const
     */
    "CatalogItemType": "CatalogItemType"
};


/**
 * Allowed values for the <code>access</code> property.
 * @enum {String}
 * @readonly
 */
RolePermissionDefaultCatalogItemType['AccessEnum'] = {

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



export default RolePermissionDefaultCatalogItemType;

