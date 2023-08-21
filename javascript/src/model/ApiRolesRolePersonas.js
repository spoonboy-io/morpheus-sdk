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
 * The ApiRolesRolePersonas model module.
 * @module model/ApiRolesRolePersonas
 * @version 6.2.1
 */
class ApiRolesRolePersonas {
    /**
     * Constructs a new <code>ApiRolesRolePersonas</code>.
     * @alias module:model/ApiRolesRolePersonas
     * @param access {module:model/ApiRolesRolePersonas.AccessEnum} The new access level.
     */
    constructor(access) { 
        
        ApiRolesRolePersonas.initialize(this, access);
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
     * Constructs a <code>ApiRolesRolePersonas</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiRolesRolePersonas} obj Optional instance to populate.
     * @return {module:model/ApiRolesRolePersonas} The populated <code>ApiRolesRolePersonas</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiRolesRolePersonas();

            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('access')) {
                obj['access'] = ApiClient.convertToType(data['access'], 'String');
            }
        }
        return obj;
    }


}

/**
 * `code` of the persona
 * @member {module:model/ApiRolesRolePersonas.CodeEnum} code
 */
ApiRolesRolePersonas.prototype['code'] = undefined;

/**
 * The new access level.
 * @member {module:model/ApiRolesRolePersonas.AccessEnum} access
 */
ApiRolesRolePersonas.prototype['access'] = undefined;





/**
 * Allowed values for the <code>code</code> property.
 * @enum {String}
 * @readonly
 */
ApiRolesRolePersonas['CodeEnum'] = {

    /**
     * value: "standard"
     * @const
     */
    "standard": "standard",

    /**
     * value: "serviceCatalog"
     * @const
     */
    "serviceCatalog": "serviceCatalog",

    /**
     * value: "vdi"
     * @const
     */
    "vdi": "vdi"
};


/**
 * Allowed values for the <code>access</code> property.
 * @enum {String}
 * @readonly
 */
ApiRolesRolePersonas['AccessEnum'] = {

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



export default ApiRolesRolePersonas;

