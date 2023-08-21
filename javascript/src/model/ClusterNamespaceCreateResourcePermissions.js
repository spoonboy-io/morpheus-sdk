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
import ClusterUpdatePermissionsResourcePermissionsSites from './ClusterUpdatePermissionsResourcePermissionsSites';

/**
 * The ClusterNamespaceCreateResourcePermissions model module.
 * @module model/ClusterNamespaceCreateResourcePermissions
 * @version 6.2.1
 */
class ClusterNamespaceCreateResourcePermissions {
    /**
     * Constructs a new <code>ClusterNamespaceCreateResourcePermissions</code>.
     * Map for namespace group and service plan permissions.
     * @alias module:model/ClusterNamespaceCreateResourcePermissions
     */
    constructor() { 
        
        ClusterNamespaceCreateResourcePermissions.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ClusterNamespaceCreateResourcePermissions</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ClusterNamespaceCreateResourcePermissions} obj Optional instance to populate.
     * @return {module:model/ClusterNamespaceCreateResourcePermissions} The populated <code>ClusterNamespaceCreateResourcePermissions</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ClusterNamespaceCreateResourcePermissions();

            if (data.hasOwnProperty('all')) {
                obj['all'] = ApiClient.convertToType(data['all'], 'Boolean');
            }
            if (data.hasOwnProperty('sites')) {
                obj['sites'] = ApiClient.convertToType(data['sites'], [ClusterUpdatePermissionsResourcePermissionsSites]);
            }
            if (data.hasOwnProperty('allPlans')) {
                obj['allPlans'] = ApiClient.convertToType(data['allPlans'], 'Boolean');
            }
            if (data.hasOwnProperty('plans')) {
                obj['plans'] = ApiClient.convertToType(data['plans'], [ClusterUpdatePermissionsResourcePermissionsSites]);
            }
        }
        return obj;
    }


}

/**
 * Pass true to allow access to all groups
 * @member {Boolean} all
 */
ClusterNamespaceCreateResourcePermissions.prototype['all'] = undefined;

/**
 * Array of groups that are allowed access
 * @member {Array.<module:model/ClusterUpdatePermissionsResourcePermissionsSites>} sites
 */
ClusterNamespaceCreateResourcePermissions.prototype['sites'] = undefined;

/**
 * Pass true to allow access to all plans
 * @member {Boolean} allPlans
 */
ClusterNamespaceCreateResourcePermissions.prototype['allPlans'] = undefined;

/**
 * Array of plans that are allowed access
 * @member {Array.<module:model/ClusterUpdatePermissionsResourcePermissionsSites>} plans
 */
ClusterNamespaceCreateResourcePermissions.prototype['plans'] = undefined;






export default ClusterNamespaceCreateResourcePermissions;
