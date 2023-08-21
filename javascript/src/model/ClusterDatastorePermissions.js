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
import ClusterDatastorePermissionsResourcePermissions from './ClusterDatastorePermissionsResourcePermissions';
import InlineResponse20095NetworkRouterPermissionsTenantPermissions from './InlineResponse20095NetworkRouterPermissionsTenantPermissions';

/**
 * The ClusterDatastorePermissions model module.
 * @module model/ClusterDatastorePermissions
 * @version 6.2.1
 */
class ClusterDatastorePermissions {
    /**
     * Constructs a new <code>ClusterDatastorePermissions</code>.
     * @alias module:model/ClusterDatastorePermissions
     */
    constructor() { 
        
        ClusterDatastorePermissions.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ClusterDatastorePermissions</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ClusterDatastorePermissions} obj Optional instance to populate.
     * @return {module:model/ClusterDatastorePermissions} The populated <code>ClusterDatastorePermissions</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ClusterDatastorePermissions();

            if (data.hasOwnProperty('resourcePermissions')) {
                obj['resourcePermissions'] = ClusterDatastorePermissionsResourcePermissions.constructFromObject(data['resourcePermissions']);
            }
            if (data.hasOwnProperty('tenantPermissions')) {
                obj['tenantPermissions'] = InlineResponse20095NetworkRouterPermissionsTenantPermissions.constructFromObject(data['tenantPermissions']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ClusterDatastorePermissionsResourcePermissions} resourcePermissions
 */
ClusterDatastorePermissions.prototype['resourcePermissions'] = undefined;

/**
 * @member {module:model/InlineResponse20095NetworkRouterPermissionsTenantPermissions} tenantPermissions
 */
ClusterDatastorePermissions.prototype['tenantPermissions'] = undefined;






export default ClusterDatastorePermissions;
