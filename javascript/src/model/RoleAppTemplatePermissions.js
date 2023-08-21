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
 * The RoleAppTemplatePermissions model module.
 * @module model/RoleAppTemplatePermissions
 * @version 6.2.1
 */
class RoleAppTemplatePermissions {
    /**
     * Constructs a new <code>RoleAppTemplatePermissions</code>.
     * @alias module:model/RoleAppTemplatePermissions
     */
    constructor() { 
        
        RoleAppTemplatePermissions.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>RoleAppTemplatePermissions</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/RoleAppTemplatePermissions} obj Optional instance to populate.
     * @return {module:model/RoleAppTemplatePermissions} The populated <code>RoleAppTemplatePermissions</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RoleAppTemplatePermissions();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('access')) {
                obj['access'] = ApiClient.convertToType(data['access'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
RoleAppTemplatePermissions.prototype['id'] = undefined;

/**
 * @member {String} code
 */
RoleAppTemplatePermissions.prototype['code'] = undefined;

/**
 * @member {String} name
 */
RoleAppTemplatePermissions.prototype['name'] = undefined;

/**
 * @member {String} access
 */
RoleAppTemplatePermissions.prototype['access'] = undefined;






export default RoleAppTemplatePermissions;
