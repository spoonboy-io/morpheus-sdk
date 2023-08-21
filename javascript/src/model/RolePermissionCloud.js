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
 * The RolePermissionCloud model module.
 * @module model/RolePermissionCloud
 * @version 6.2.1
 */
class RolePermissionCloud {
    /**
     * Constructs a new <code>RolePermissionCloud</code>.
     * @alias module:model/RolePermissionCloud
     * @param cloudId {Number} `id` of the cloud (zone)
     * @param access {module:model/RolePermissionCloud.AccessEnum} The new access level.
     */
    constructor(cloudId, access) { 
        
        RolePermissionCloud.initialize(this, cloudId, access);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, cloudId, access) { 
        obj['cloudId'] = cloudId;
        obj['access'] = access;
    }

    /**
     * Constructs a <code>RolePermissionCloud</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/RolePermissionCloud} obj Optional instance to populate.
     * @return {module:model/RolePermissionCloud} The populated <code>RolePermissionCloud</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RolePermissionCloud();

            if (data.hasOwnProperty('cloudId')) {
                obj['cloudId'] = ApiClient.convertToType(data['cloudId'], 'Number');
            }
            if (data.hasOwnProperty('access')) {
                obj['access'] = ApiClient.convertToType(data['access'], 'String');
            }
        }
        return obj;
    }


}

/**
 * `id` of the cloud (zone)
 * @member {Number} cloudId
 */
RolePermissionCloud.prototype['cloudId'] = undefined;

/**
 * The new access level.
 * @member {module:model/RolePermissionCloud.AccessEnum} access
 */
RolePermissionCloud.prototype['access'] = undefined;





/**
 * Allowed values for the <code>access</code> property.
 * @enum {String}
 * @readonly
 */
RolePermissionCloud['AccessEnum'] = {

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



export default RolePermissionCloud;

