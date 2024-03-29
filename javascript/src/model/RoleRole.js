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
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';

/**
 * The RoleRole model module.
 * @module model/RoleRole
 * @version 6.2.1
 */
class RoleRole {
    /**
     * Constructs a new <code>RoleRole</code>.
     * @alias module:model/RoleRole
     */
    constructor() { 
        
        RoleRole.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>RoleRole</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/RoleRole} obj Optional instance to populate.
     * @return {module:model/RoleRole} The populated <code>RoleRole</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RoleRole();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('authority')) {
                obj['authority'] = ApiClient.convertToType(data['authority'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('scope')) {
                obj['scope'] = ApiClient.convertToType(data['scope'], 'String');
            }
            if (data.hasOwnProperty('roleType')) {
                obj['roleType'] = ApiClient.convertToType(data['roleType'], 'String');
            }
            if (data.hasOwnProperty('multitenant')) {
                obj['multitenant'] = ApiClient.convertToType(data['multitenant'], 'Boolean');
            }
            if (data.hasOwnProperty('multitenantLocked')) {
                obj['multitenantLocked'] = ApiClient.convertToType(data['multitenantLocked'], 'Boolean');
            }
            if (data.hasOwnProperty('parentRoleId')) {
                obj['parentRoleId'] = ApiClient.convertToType(data['parentRoleId'], 'String');
            }
            if (data.hasOwnProperty('diverged')) {
                obj['diverged'] = ApiClient.convertToType(data['diverged'], 'Boolean');
            }
            if (data.hasOwnProperty('ownerId')) {
                obj['ownerId'] = ApiClient.convertToType(data['ownerId'], 'Number');
            }
            if (data.hasOwnProperty('owner')) {
                obj['owner'] = InlineResponse20040AppDeployInstance.constructFromObject(data['owner']);
            }
            if (data.hasOwnProperty('defaultPersona')) {
                obj['defaultPersona'] = ApiClient.convertToType(data['defaultPersona'], 'String');
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
RoleRole.prototype['id'] = undefined;

/**
 * a unique name of the role
 * @member {String} name
 */
RoleRole.prototype['name'] = undefined;

/**
 * Alias for name
 * @member {String} authority
 */
RoleRole.prototype['authority'] = undefined;

/**
 * @member {String} description
 */
RoleRole.prototype['description'] = undefined;

/**
 * @member {String} scope
 */
RoleRole.prototype['scope'] = undefined;

/**
 * @member {String} roleType
 */
RoleRole.prototype['roleType'] = undefined;

/**
 * @member {Boolean} multitenant
 */
RoleRole.prototype['multitenant'] = undefined;

/**
 * @member {Boolean} multitenantLocked
 */
RoleRole.prototype['multitenantLocked'] = undefined;

/**
 * @member {String} parentRoleId
 */
RoleRole.prototype['parentRoleId'] = undefined;

/**
 * @member {Boolean} diverged
 */
RoleRole.prototype['diverged'] = undefined;

/**
 * @member {Number} ownerId
 */
RoleRole.prototype['ownerId'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} owner
 */
RoleRole.prototype['owner'] = undefined;

/**
 * @member {String} defaultPersona
 */
RoleRole.prototype['defaultPersona'] = undefined;

/**
 * @member {Date} dateCreated
 */
RoleRole.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
RoleRole.prototype['lastUpdated'] = undefined;






export default RoleRole;

