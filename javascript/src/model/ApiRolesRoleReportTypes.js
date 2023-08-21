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
 * The ApiRolesRoleReportTypes model module.
 * @module model/ApiRolesRoleReportTypes
 * @version 6.2.1
 */
class ApiRolesRoleReportTypes {
    /**
     * Constructs a new <code>ApiRolesRoleReportTypes</code>.
     * @alias module:model/ApiRolesRoleReportTypes
     * @param code {String} `code` of the report type
     * @param access {module:model/ApiRolesRoleReportTypes.AccessEnum} The new access level.
     */
    constructor(code, access) { 
        
        ApiRolesRoleReportTypes.initialize(this, code, access);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, code, access) { 
        obj['code'] = code;
        obj['access'] = access;
    }

    /**
     * Constructs a <code>ApiRolesRoleReportTypes</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiRolesRoleReportTypes} obj Optional instance to populate.
     * @return {module:model/ApiRolesRoleReportTypes} The populated <code>ApiRolesRoleReportTypes</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiRolesRoleReportTypes();

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
 * `code` of the report type
 * @member {String} code
 */
ApiRolesRoleReportTypes.prototype['code'] = undefined;

/**
 * The new access level.
 * @member {module:model/ApiRolesRoleReportTypes.AccessEnum} access
 */
ApiRolesRoleReportTypes.prototype['access'] = undefined;





/**
 * Allowed values for the <code>access</code> property.
 * @enum {String}
 * @readonly
 */
ApiRolesRoleReportTypes['AccessEnum'] = {

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



export default ApiRolesRoleReportTypes;
