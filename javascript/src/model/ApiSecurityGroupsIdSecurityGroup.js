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
import ApiSecurityGroupsSecurityGroupTenantPermissions from './ApiSecurityGroupsSecurityGroupTenantPermissions';
import ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions from './ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions';

/**
 * The ApiSecurityGroupsIdSecurityGroup model module.
 * @module model/ApiSecurityGroupsIdSecurityGroup
 * @version 6.2.1
 */
class ApiSecurityGroupsIdSecurityGroup {
    /**
     * Constructs a new <code>ApiSecurityGroupsIdSecurityGroup</code>.
     * @alias module:model/ApiSecurityGroupsIdSecurityGroup
     */
    constructor() { 
        
        ApiSecurityGroupsIdSecurityGroup.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiSecurityGroupsIdSecurityGroup</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiSecurityGroupsIdSecurityGroup} obj Optional instance to populate.
     * @return {module:model/ApiSecurityGroupsIdSecurityGroup} The populated <code>ApiSecurityGroupsIdSecurityGroup</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiSecurityGroupsIdSecurityGroup();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
            if (data.hasOwnProperty('tenantPermissions')) {
                obj['tenantPermissions'] = ApiClient.convertToType(data['tenantPermissions'], [ApiSecurityGroupsSecurityGroupTenantPermissions]);
            }
            if (data.hasOwnProperty('resourcePermissions')) {
                obj['resourcePermissions'] = ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions.constructFromObject(data['resourcePermissions']);
            }
        }
        return obj;
    }


}

/**
 * Name for your security group
 * @member {String} name
 */
ApiSecurityGroupsIdSecurityGroup.prototype['name'] = undefined;

/**
 * Optional description field
 * @member {String} description
 */
ApiSecurityGroupsIdSecurityGroup.prototype['description'] = undefined;

/**
 * Set to `false` to disable a security group.
 * @member {Boolean} active
 */
ApiSecurityGroupsIdSecurityGroup.prototype['active'] = undefined;

/**
 * @member {Array.<module:model/ApiSecurityGroupsSecurityGroupTenantPermissions>} tenantPermissions
 */
ApiSecurityGroupsIdSecurityGroup.prototype['tenantPermissions'] = undefined;

/**
 * @member {module:model/ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions} resourcePermissions
 */
ApiSecurityGroupsIdSecurityGroup.prototype['resourcePermissions'] = undefined;






export default ApiSecurityGroupsIdSecurityGroup;

