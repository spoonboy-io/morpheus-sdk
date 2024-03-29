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
import ClusterNamespaceCreateResourcePermissions from './ClusterNamespaceCreateResourcePermissions';

/**
 * The ClusterNamespaceUpdatePermissions model module.
 * @module model/ClusterNamespaceUpdatePermissions
 * @version 6.2.1
 */
class ClusterNamespaceUpdatePermissions {
    /**
     * Constructs a new <code>ClusterNamespaceUpdatePermissions</code>.
     * @alias module:model/ClusterNamespaceUpdatePermissions
     */
    constructor() { 
        
        ClusterNamespaceUpdatePermissions.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ClusterNamespaceUpdatePermissions</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ClusterNamespaceUpdatePermissions} obj Optional instance to populate.
     * @return {module:model/ClusterNamespaceUpdatePermissions} The populated <code>ClusterNamespaceUpdatePermissions</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ClusterNamespaceUpdatePermissions();

            if (data.hasOwnProperty('resourcePermissions')) {
                obj['resourcePermissions'] = ClusterNamespaceCreateResourcePermissions.constructFromObject(data['resourcePermissions']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ClusterNamespaceCreateResourcePermissions} resourcePermissions
 */
ClusterNamespaceUpdatePermissions.prototype['resourcePermissions'] = undefined;






export default ClusterNamespaceUpdatePermissions;

