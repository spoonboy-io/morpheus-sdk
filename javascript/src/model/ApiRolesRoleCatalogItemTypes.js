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
 * The ApiRolesRoleCatalogItemTypes model module.
 * @module model/ApiRolesRoleCatalogItemTypes
 * @version 6.2.1
 */
class ApiRolesRoleCatalogItemTypes {
    /**
     * Constructs a new <code>ApiRolesRoleCatalogItemTypes</code>.
     * @alias module:model/ApiRolesRoleCatalogItemTypes
     * @param id {Number} `id` of the catalog item type
     * @param access {module:model/ApiRolesRoleCatalogItemTypes.AccessEnum} The new access level.
     */
    constructor(id, access) { 
        
        ApiRolesRoleCatalogItemTypes.initialize(this, id, access);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, id, access) { 
        obj['id'] = id;
        obj['access'] = access;
    }

    /**
     * Constructs a <code>ApiRolesRoleCatalogItemTypes</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiRolesRoleCatalogItemTypes} obj Optional instance to populate.
     * @return {module:model/ApiRolesRoleCatalogItemTypes} The populated <code>ApiRolesRoleCatalogItemTypes</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiRolesRoleCatalogItemTypes();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('access')) {
                obj['access'] = ApiClient.convertToType(data['access'], 'String');
            }
        }
        return obj;
    }


}

/**
 * `id` of the catalog item type
 * @member {Number} id
 */
ApiRolesRoleCatalogItemTypes.prototype['id'] = undefined;

/**
 * The new access level.
 * @member {module:model/ApiRolesRoleCatalogItemTypes.AccessEnum} access
 */
ApiRolesRoleCatalogItemTypes.prototype['access'] = undefined;





/**
 * Allowed values for the <code>access</code> property.
 * @enum {String}
 * @readonly
 */
ApiRolesRoleCatalogItemTypes['AccessEnum'] = {

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



export default ApiRolesRoleCatalogItemTypes;

