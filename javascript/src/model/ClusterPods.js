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
import ApiBlueprintsIdUpdatePermissionsResourcePermissionSites from './ApiBlueprintsIdUpdatePermissionsResourcePermissionSites';

/**
 * The ClusterPods model module.
 * @module model/ClusterPods
 * @version 6.2.1
 */
class ClusterPods {
    /**
     * Constructs a new <code>ClusterPods</code>.
     * @alias module:model/ClusterPods
     */
    constructor() { 
        
        ClusterPods.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ClusterPods</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ClusterPods} obj Optional instance to populate.
     * @return {module:model/ClusterPods} The populated <code>ClusterPods</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ClusterPods();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('resourceLevel')) {
                obj['resourceLevel'] = ApiClient.convertToType(data['resourceLevel'], 'String');
            }
            if (data.hasOwnProperty('resourceType')) {
                obj['resourceType'] = ApiClient.convertToType(data['resourceType'], 'String');
            }
            if (data.hasOwnProperty('managed')) {
                obj['managed'] = ApiClient.convertToType(data['managed'], 'Boolean');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
            if (data.hasOwnProperty('owner')) {
                obj['owner'] = ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.constructFromObject(data['owner']);
            }
            if (data.hasOwnProperty('totalCpuUsage')) {
                obj['totalCpuUsage'] = ApiClient.convertToType(data['totalCpuUsage'], 'Number');
            }
            if (data.hasOwnProperty('stats')) {
                obj['stats'] = ApiClient.convertToType(data['stats'], Object);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ClusterPods.prototype['id'] = undefined;

/**
 * @member {String} name
 */
ClusterPods.prototype['name'] = undefined;

/**
 * @member {String} code
 */
ClusterPods.prototype['code'] = undefined;

/**
 * @member {String} description
 */
ClusterPods.prototype['description'] = undefined;

/**
 * @member {String} category
 */
ClusterPods.prototype['category'] = undefined;

/**
 * @member {String} resourceLevel
 */
ClusterPods.prototype['resourceLevel'] = undefined;

/**
 * @member {String} resourceType
 */
ClusterPods.prototype['resourceType'] = undefined;

/**
 * @member {Boolean} managed
 */
ClusterPods.prototype['managed'] = undefined;

/**
 * @member {String} status
 */
ClusterPods.prototype['status'] = undefined;

/**
 * @member {Date} lastUpdated
 */
ClusterPods.prototype['lastUpdated'] = undefined;

/**
 * @member {module:model/ApiBlueprintsIdUpdatePermissionsResourcePermissionSites} owner
 */
ClusterPods.prototype['owner'] = undefined;

/**
 * @member {Number} totalCpuUsage
 */
ClusterPods.prototype['totalCpuUsage'] = undefined;

/**
 * @member {Object} stats
 */
ClusterPods.prototype['stats'] = undefined;






export default ClusterPods;
