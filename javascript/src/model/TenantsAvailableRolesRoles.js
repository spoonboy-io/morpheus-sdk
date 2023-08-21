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
 * The TenantsAvailableRolesRoles model module.
 * @module model/TenantsAvailableRolesRoles
 * @version 6.2.1
 */
class TenantsAvailableRolesRoles {
    /**
     * Constructs a new <code>TenantsAvailableRolesRoles</code>.
     * @alias module:model/TenantsAvailableRolesRoles
     */
    constructor() { 
        
        TenantsAvailableRolesRoles.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TenantsAvailableRolesRoles</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TenantsAvailableRolesRoles} obj Optional instance to populate.
     * @return {module:model/TenantsAvailableRolesRoles} The populated <code>TenantsAvailableRolesRoles</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TenantsAvailableRolesRoles();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('authority')) {
                obj['authority'] = ApiClient.convertToType(data['authority'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('roleType')) {
                obj['roleType'] = ApiClient.convertToType(data['roleType'], 'String');
            }
            if (data.hasOwnProperty('owner')) {
                obj['owner'] = ApiClient.convertToType(data['owner'], Object);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
TenantsAvailableRolesRoles.prototype['id'] = undefined;

/**
 * @member {String} authority
 */
TenantsAvailableRolesRoles.prototype['authority'] = undefined;

/**
 * @member {String} description
 */
TenantsAvailableRolesRoles.prototype['description'] = undefined;

/**
 * @member {String} roleType
 */
TenantsAvailableRolesRoles.prototype['roleType'] = undefined;

/**
 * @member {Object} owner
 */
TenantsAvailableRolesRoles.prototype['owner'] = undefined;






export default TenantsAvailableRolesRoles;
