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
 * The ClusterUpdatePermissionsResourcePermissions model module.
 * @module model/ClusterUpdatePermissionsResourcePermissions
 * @version 6.2.1
 */
class ClusterUpdatePermissionsResourcePermissions {
    /**
     * Constructs a new <code>ClusterUpdatePermissionsResourcePermissions</code>.
     * @alias module:model/ClusterUpdatePermissionsResourcePermissions
     */
    constructor() { 
        
        ClusterUpdatePermissionsResourcePermissions.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ClusterUpdatePermissionsResourcePermissions</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ClusterUpdatePermissionsResourcePermissions} obj Optional instance to populate.
     * @return {module:model/ClusterUpdatePermissionsResourcePermissions} The populated <code>ClusterUpdatePermissionsResourcePermissions</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ClusterUpdatePermissionsResourcePermissions();

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
ClusterUpdatePermissionsResourcePermissions.prototype['all'] = undefined;

/**
 * Array of groups that are allowed access
 * @member {Array.<module:model/ClusterUpdatePermissionsResourcePermissionsSites>} sites
 */
ClusterUpdatePermissionsResourcePermissions.prototype['sites'] = undefined;

/**
 * Pass true to allow access to all plans
 * @member {Boolean} allPlans
 */
ClusterUpdatePermissionsResourcePermissions.prototype['allPlans'] = undefined;

/**
 * Array of plans that are allowed access
 * @member {Array.<module:model/ClusterUpdatePermissionsResourcePermissionsSites>} plans
 */
ClusterUpdatePermissionsResourcePermissions.prototype['plans'] = undefined;






export default ClusterUpdatePermissionsResourcePermissions;

