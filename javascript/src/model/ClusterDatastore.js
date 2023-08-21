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
import ClusterDatastorePermissions from './ClusterDatastorePermissions';
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';

/**
 * The ClusterDatastore model module.
 * @module model/ClusterDatastore
 * @version 6.2.1
 */
class ClusterDatastore {
    /**
     * Constructs a new <code>ClusterDatastore</code>.
     * @alias module:model/ClusterDatastore
     */
    constructor() { 
        
        ClusterDatastore.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ClusterDatastore</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ClusterDatastore} obj Optional instance to populate.
     * @return {module:model/ClusterDatastore} The populated <code>ClusterDatastore</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ClusterDatastore();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('storageSize')) {
                obj['storageSize'] = ApiClient.convertToType(data['storageSize'], 'Number');
            }
            if (data.hasOwnProperty('freeSpace')) {
                obj['freeSpace'] = ApiClient.convertToType(data['freeSpace'], 'Number');
            }
            if (data.hasOwnProperty('drsEnabled')) {
                obj['drsEnabled'] = ApiClient.convertToType(data['drsEnabled'], 'Boolean');
            }
            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
            if (data.hasOwnProperty('allowWrite')) {
                obj['allowWrite'] = ApiClient.convertToType(data['allowWrite'], 'Boolean');
            }
            if (data.hasOwnProperty('defaultStore')) {
                obj['defaultStore'] = ApiClient.convertToType(data['defaultStore'], 'Boolean');
            }
            if (data.hasOwnProperty('online')) {
                obj['online'] = ApiClient.convertToType(data['online'], 'Boolean');
            }
            if (data.hasOwnProperty('allowRead')) {
                obj['allowRead'] = ApiClient.convertToType(data['allowRead'], 'Boolean');
            }
            if (data.hasOwnProperty('allowProvision')) {
                obj['allowProvision'] = ApiClient.convertToType(data['allowProvision'], 'Boolean');
            }
            if (data.hasOwnProperty('refType')) {
                obj['refType'] = ApiClient.convertToType(data['refType'], 'String');
            }
            if (data.hasOwnProperty('refId')) {
                obj['refId'] = ApiClient.convertToType(data['refId'], 'Number');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('zone')) {
                obj['zone'] = ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.constructFromObject(data['zone']);
            }
            if (data.hasOwnProperty('zonePool')) {
                obj['zonePool'] = ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.constructFromObject(data['zonePool']);
            }
            if (data.hasOwnProperty('owner')) {
                obj['owner'] = ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.constructFromObject(data['owner']);
            }
            if (data.hasOwnProperty('tenants')) {
                obj['tenants'] = ApiClient.convertToType(data['tenants'], [InlineResponse20040AppDeployInstance]);
            }
            if (data.hasOwnProperty('permissions')) {
                obj['permissions'] = ClusterDatastorePermissions.constructFromObject(data['permissions']);
            }
            if (data.hasOwnProperty('datastores')) {
                obj['datastores'] = ApiClient.convertToType(data['datastores'], [Object]);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ClusterDatastore.prototype['id'] = undefined;

/**
 * @member {String} name
 */
ClusterDatastore.prototype['name'] = undefined;

/**
 * @member {String} code
 */
ClusterDatastore.prototype['code'] = undefined;

/**
 * @member {String} type
 */
ClusterDatastore.prototype['type'] = undefined;

/**
 * @member {String} visibility
 */
ClusterDatastore.prototype['visibility'] = undefined;

/**
 * @member {Number} storageSize
 */
ClusterDatastore.prototype['storageSize'] = undefined;

/**
 * @member {Number} freeSpace
 */
ClusterDatastore.prototype['freeSpace'] = undefined;

/**
 * @member {Boolean} drsEnabled
 */
ClusterDatastore.prototype['drsEnabled'] = undefined;

/**
 * @member {Boolean} active
 */
ClusterDatastore.prototype['active'] = undefined;

/**
 * @member {Boolean} allowWrite
 */
ClusterDatastore.prototype['allowWrite'] = undefined;

/**
 * @member {Boolean} defaultStore
 */
ClusterDatastore.prototype['defaultStore'] = undefined;

/**
 * @member {Boolean} online
 */
ClusterDatastore.prototype['online'] = undefined;

/**
 * @member {Boolean} allowRead
 */
ClusterDatastore.prototype['allowRead'] = undefined;

/**
 * @member {Boolean} allowProvision
 */
ClusterDatastore.prototype['allowProvision'] = undefined;

/**
 * @member {String} refType
 */
ClusterDatastore.prototype['refType'] = undefined;

/**
 * @member {Number} refId
 */
ClusterDatastore.prototype['refId'] = undefined;

/**
 * @member {String} externalId
 */
ClusterDatastore.prototype['externalId'] = undefined;

/**
 * @member {module:model/ApiBlueprintsIdUpdatePermissionsResourcePermissionSites} zone
 */
ClusterDatastore.prototype['zone'] = undefined;

/**
 * @member {module:model/ApiBlueprintsIdUpdatePermissionsResourcePermissionSites} zonePool
 */
ClusterDatastore.prototype['zonePool'] = undefined;

/**
 * @member {module:model/ApiBlueprintsIdUpdatePermissionsResourcePermissionSites} owner
 */
ClusterDatastore.prototype['owner'] = undefined;

/**
 * @member {Array.<module:model/InlineResponse20040AppDeployInstance>} tenants
 */
ClusterDatastore.prototype['tenants'] = undefined;

/**
 * @member {module:model/ClusterDatastorePermissions} permissions
 */
ClusterDatastore.prototype['permissions'] = undefined;

/**
 * @member {Array.<Object>} datastores
 */
ClusterDatastore.prototype['datastores'] = undefined;






export default ClusterDatastore;

