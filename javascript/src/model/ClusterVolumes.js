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
 * The ClusterVolumes model module.
 * @module model/ClusterVolumes
 * @version 6.2.1
 */
class ClusterVolumes {
    /**
     * Constructs a new <code>ClusterVolumes</code>.
     * @alias module:model/ClusterVolumes
     */
    constructor() { 
        
        ClusterVolumes.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ClusterVolumes</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ClusterVolumes} obj Optional instance to populate.
     * @return {module:model/ClusterVolumes} The populated <code>ClusterVolumes</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ClusterVolumes();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('internalId')) {
                obj['internalId'] = ApiClient.convertToType(data['internalId'], 'String');
            }
            if (data.hasOwnProperty('displayOrder')) {
                obj['displayOrder'] = ApiClient.convertToType(data['displayOrder'], 'Number');
            }
            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
            if (data.hasOwnProperty('usedStorage')) {
                obj['usedStorage'] = ApiClient.convertToType(data['usedStorage'], 'Number');
            }
            if (data.hasOwnProperty('provisionType')) {
                obj['provisionType'] = ApiClient.convertToType(data['provisionType'], 'String');
            }
            if (data.hasOwnProperty('resizeable')) {
                obj['resizeable'] = ApiClient.convertToType(data['resizeable'], 'Boolean');
            }
            if (data.hasOwnProperty('online')) {
                obj['online'] = ApiClient.convertToType(data['online'], 'Boolean');
            }
            if (data.hasOwnProperty('deviceDisplayName')) {
                obj['deviceDisplayName'] = ApiClient.convertToType(data['deviceDisplayName'], 'String');
            }
            if (data.hasOwnProperty('refType')) {
                obj['refType'] = ApiClient.convertToType(data['refType'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('datastoreOption')) {
                obj['datastoreOption'] = ApiClient.convertToType(data['datastoreOption'], 'String');
            }
            if (data.hasOwnProperty('claimName')) {
                obj['claimName'] = ApiClient.convertToType(data['claimName'], 'String');
            }
            if (data.hasOwnProperty('volumeType')) {
                obj['volumeType'] = ApiClient.convertToType(data['volumeType'], 'String');
            }
            if (data.hasOwnProperty('deviceName')) {
                obj['deviceName'] = ApiClient.convertToType(data['deviceName'], 'String');
            }
            if (data.hasOwnProperty('removable')) {
                obj['removable'] = ApiClient.convertToType(data['removable'], 'Boolean');
            }
            if (data.hasOwnProperty('poolName')) {
                obj['poolName'] = ApiClient.convertToType(data['poolName'], 'String');
            }
            if (data.hasOwnProperty('readOnly')) {
                obj['readOnly'] = ApiClient.convertToType(data['readOnly'], 'Boolean');
            }
            if (data.hasOwnProperty('sourceId')) {
                obj['sourceId'] = ApiClient.convertToType(data['sourceId'], 'String');
            }
            if (data.hasOwnProperty('zoneId')) {
                obj['zoneId'] = ApiClient.convertToType(data['zoneId'], 'Number');
            }
            if (data.hasOwnProperty('rootVolume')) {
                obj['rootVolume'] = ApiClient.convertToType(data['rootVolume'], 'Boolean');
            }
            if (data.hasOwnProperty('refId')) {
                obj['refId'] = ApiClient.convertToType(data['refId'], 'Number');
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('rawData')) {
                obj['rawData'] = ApiClient.convertToType(data['rawData'], 'String');
            }
            if (data.hasOwnProperty('maxStorage')) {
                obj['maxStorage'] = ApiClient.convertToType(data['maxStorage'], 'Number');
            }
            if (data.hasOwnProperty('account')) {
                obj['account'] = ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.constructFromObject(data['account']);
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.constructFromObject(data['type']);
            }
            if (data.hasOwnProperty('datastore')) {
                obj['datastore'] = ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.constructFromObject(data['datastore']);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ClusterVolumes.prototype['id'] = undefined;

/**
 * @member {String} internalId
 */
ClusterVolumes.prototype['internalId'] = undefined;

/**
 * @member {Number} displayOrder
 */
ClusterVolumes.prototype['displayOrder'] = undefined;

/**
 * @member {Boolean} active
 */
ClusterVolumes.prototype['active'] = undefined;

/**
 * @member {Number} usedStorage
 */
ClusterVolumes.prototype['usedStorage'] = undefined;

/**
 * @member {String} provisionType
 */
ClusterVolumes.prototype['provisionType'] = undefined;

/**
 * @member {Boolean} resizeable
 */
ClusterVolumes.prototype['resizeable'] = undefined;

/**
 * @member {Boolean} online
 */
ClusterVolumes.prototype['online'] = undefined;

/**
 * @member {String} deviceDisplayName
 */
ClusterVolumes.prototype['deviceDisplayName'] = undefined;

/**
 * @member {String} refType
 */
ClusterVolumes.prototype['refType'] = undefined;

/**
 * @member {String} name
 */
ClusterVolumes.prototype['name'] = undefined;

/**
 * @member {String} externalId
 */
ClusterVolumes.prototype['externalId'] = undefined;

/**
 * @member {String} datastoreOption
 */
ClusterVolumes.prototype['datastoreOption'] = undefined;

/**
 * @member {String} claimName
 */
ClusterVolumes.prototype['claimName'] = undefined;

/**
 * @member {String} volumeType
 */
ClusterVolumes.prototype['volumeType'] = undefined;

/**
 * @member {String} deviceName
 */
ClusterVolumes.prototype['deviceName'] = undefined;

/**
 * @member {Boolean} removable
 */
ClusterVolumes.prototype['removable'] = undefined;

/**
 * @member {String} poolName
 */
ClusterVolumes.prototype['poolName'] = undefined;

/**
 * @member {Boolean} readOnly
 */
ClusterVolumes.prototype['readOnly'] = undefined;

/**
 * @member {String} sourceId
 */
ClusterVolumes.prototype['sourceId'] = undefined;

/**
 * @member {Number} zoneId
 */
ClusterVolumes.prototype['zoneId'] = undefined;

/**
 * @member {Boolean} rootVolume
 */
ClusterVolumes.prototype['rootVolume'] = undefined;

/**
 * @member {Number} refId
 */
ClusterVolumes.prototype['refId'] = undefined;

/**
 * @member {String} category
 */
ClusterVolumes.prototype['category'] = undefined;

/**
 * @member {String} status
 */
ClusterVolumes.prototype['status'] = undefined;

/**
 * @member {String} rawData
 */
ClusterVolumes.prototype['rawData'] = undefined;

/**
 * @member {Number} maxStorage
 */
ClusterVolumes.prototype['maxStorage'] = undefined;

/**
 * @member {module:model/ApiBlueprintsIdUpdatePermissionsResourcePermissionSites} account
 */
ClusterVolumes.prototype['account'] = undefined;

/**
 * @member {module:model/ApiBlueprintsIdUpdatePermissionsResourcePermissionSites} type
 */
ClusterVolumes.prototype['type'] = undefined;

/**
 * @member {module:model/ApiBlueprintsIdUpdatePermissionsResourcePermissionSites} datastore
 */
ClusterVolumes.prototype['datastore'] = undefined;






export default ClusterVolumes;

