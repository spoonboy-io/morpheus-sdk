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
import ZoneVcenterConfigNetworkServer from './ZoneVcenterConfigNetworkServer';

/**
 * The ZoneVcenterConfig model module.
 * @module model/ZoneVcenterConfig
 * @version 6.2.1
 */
class ZoneVcenterConfig {
    /**
     * Constructs a new <code>ZoneVcenterConfig</code>.
     * @alias module:model/ZoneVcenterConfig
     */
    constructor() { 
        
        ZoneVcenterConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ZoneVcenterConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ZoneVcenterConfig} obj Optional instance to populate.
     * @return {module:model/ZoneVcenterConfig} The populated <code>ZoneVcenterConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ZoneVcenterConfig();

            if (data.hasOwnProperty('apiUrl')) {
                obj['apiUrl'] = ApiClient.convertToType(data['apiUrl'], 'String');
            }
            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
            }
            if (data.hasOwnProperty('password')) {
                obj['password'] = ApiClient.convertToType(data['password'], 'String');
            }
            if (data.hasOwnProperty('datacenter')) {
                obj['datacenter'] = ApiClient.convertToType(data['datacenter'], 'String');
            }
            if (data.hasOwnProperty('cluster')) {
                obj['cluster'] = ApiClient.convertToType(data['cluster'], 'String');
            }
            if (data.hasOwnProperty('resourcePoolId')) {
                obj['resourcePoolId'] = ApiClient.convertToType(data['resourcePoolId'], 'String');
            }
            if (data.hasOwnProperty('resourcePool')) {
                obj['resourcePool'] = ApiClient.convertToType(data['resourcePool'], 'String');
            }
            if (data.hasOwnProperty('rpcMode')) {
                obj['rpcMode'] = ApiClient.convertToType(data['rpcMode'], 'String');
            }
            if (data.hasOwnProperty('_hideHostSelection')) {
                obj['_hideHostSelection'] = ApiClient.convertToType(data['_hideHostSelection'], 'String');
            }
            if (data.hasOwnProperty('hideHostSelection')) {
                obj['hideHostSelection'] = ApiClient.convertToType(data['hideHostSelection'], 'String');
            }
            if (data.hasOwnProperty('_importExisting')) {
                obj['_importExisting'] = ApiClient.convertToType(data['_importExisting'], 'String');
            }
            if (data.hasOwnProperty('importExisting')) {
                obj['importExisting'] = ApiClient.convertToType(data['importExisting'], 'String');
            }
            if (data.hasOwnProperty('_enableVnc')) {
                obj['_enableVnc'] = ApiClient.convertToType(data['_enableVnc'], 'String');
            }
            if (data.hasOwnProperty('enableVnc')) {
                obj['enableVnc'] = ApiClient.convertToType(data['enableVnc'], 'String');
            }
            if (data.hasOwnProperty('_enableDiskTypeSelection')) {
                obj['_enableDiskTypeSelection'] = ApiClient.convertToType(data['_enableDiskTypeSelection'], 'String');
            }
            if (data.hasOwnProperty('_enableNetworkTypeSelection')) {
                obj['_enableNetworkTypeSelection'] = ApiClient.convertToType(data['_enableNetworkTypeSelection'], 'String');
            }
            if (data.hasOwnProperty('diskStorageType')) {
                obj['diskStorageType'] = ApiClient.convertToType(data['diskStorageType'], 'String');
            }
            if (data.hasOwnProperty('applianceUrl')) {
                obj['applianceUrl'] = ApiClient.convertToType(data['applianceUrl'], 'String');
            }
            if (data.hasOwnProperty('datacenterName')) {
                obj['datacenterName'] = ApiClient.convertToType(data['datacenterName'], 'String');
            }
            if (data.hasOwnProperty('networkServer.id')) {
                obj['networkServer.id'] = ApiClient.convertToType(data['networkServer.id'], 'String');
            }
            if (data.hasOwnProperty('networkServer')) {
                obj['networkServer'] = ZoneVcenterConfigNetworkServer.constructFromObject(data['networkServer']);
            }
            if (data.hasOwnProperty('securityMode')) {
                obj['securityMode'] = ApiClient.convertToType(data['securityMode'], 'String');
            }
            if (data.hasOwnProperty('certificateProvider')) {
                obj['certificateProvider'] = ApiClient.convertToType(data['certificateProvider'], 'String');
            }
            if (data.hasOwnProperty('backupMode')) {
                obj['backupMode'] = ApiClient.convertToType(data['backupMode'], 'String');
            }
            if (data.hasOwnProperty('replicationMode')) {
                obj['replicationMode'] = ApiClient.convertToType(data['replicationMode'], 'String');
            }
            if (data.hasOwnProperty('dnsIntegrationId')) {
                obj['dnsIntegrationId'] = ApiClient.convertToType(data['dnsIntegrationId'], 'String');
            }
            if (data.hasOwnProperty('configCmdbId')) {
                obj['configCmdbId'] = ApiClient.convertToType(data['configCmdbId'], 'String');
            }
            if (data.hasOwnProperty('configManagementId')) {
                obj['configManagementId'] = ApiClient.convertToType(data['configManagementId'], 'String');
            }
            if (data.hasOwnProperty('configCmId')) {
                obj['configCmId'] = ApiClient.convertToType(data['configCmId'], 'String');
            }
            if (data.hasOwnProperty('securityServer')) {
                obj['securityServer'] = ApiClient.convertToType(data['securityServer'], 'String');
            }
            if (data.hasOwnProperty('serviceRegistryId')) {
                obj['serviceRegistryId'] = ApiClient.convertToType(data['serviceRegistryId'], 'String');
            }
            if (data.hasOwnProperty('enableDiskTypeSelection')) {
                obj['enableDiskTypeSelection'] = ApiClient.convertToType(data['enableDiskTypeSelection'], 'String');
            }
            if (data.hasOwnProperty('kubeUrl')) {
                obj['kubeUrl'] = ApiClient.convertToType(data['kubeUrl'], 'String');
            }
            if (data.hasOwnProperty('apiVersion')) {
                obj['apiVersion'] = ApiClient.convertToType(data['apiVersion'], 'String');
            }
            if (data.hasOwnProperty('datacenterId')) {
                obj['datacenterId'] = ApiClient.convertToType(data['datacenterId'], 'String');
            }
            if (data.hasOwnProperty('configCmdbDiscovery')) {
                obj['configCmdbDiscovery'] = ApiClient.convertToType(data['configCmdbDiscovery'], 'Boolean');
            }
            if (data.hasOwnProperty('distributedWorkerId')) {
                obj['distributedWorkerId'] = ApiClient.convertToType(data['distributedWorkerId'], 'String');
            }
            if (data.hasOwnProperty('passwordHash')) {
                obj['passwordHash'] = ApiClient.convertToType(data['passwordHash'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} apiUrl
 */
ZoneVcenterConfig.prototype['apiUrl'] = undefined;

/**
 * @member {String} username
 */
ZoneVcenterConfig.prototype['username'] = undefined;

/**
 * @member {String} password
 */
ZoneVcenterConfig.prototype['password'] = undefined;

/**
 * @member {String} datacenter
 */
ZoneVcenterConfig.prototype['datacenter'] = undefined;

/**
 * @member {String} cluster
 */
ZoneVcenterConfig.prototype['cluster'] = undefined;

/**
 * @member {String} resourcePoolId
 */
ZoneVcenterConfig.prototype['resourcePoolId'] = undefined;

/**
 * @member {String} resourcePool
 */
ZoneVcenterConfig.prototype['resourcePool'] = undefined;

/**
 * @member {String} rpcMode
 */
ZoneVcenterConfig.prototype['rpcMode'] = undefined;

/**
 * @member {String} _hideHostSelection
 */
ZoneVcenterConfig.prototype['_hideHostSelection'] = undefined;

/**
 * @member {String} hideHostSelection
 */
ZoneVcenterConfig.prototype['hideHostSelection'] = undefined;

/**
 * @member {String} _importExisting
 */
ZoneVcenterConfig.prototype['_importExisting'] = undefined;

/**
 * @member {String} importExisting
 */
ZoneVcenterConfig.prototype['importExisting'] = undefined;

/**
 * @member {String} _enableVnc
 */
ZoneVcenterConfig.prototype['_enableVnc'] = undefined;

/**
 * @member {String} enableVnc
 */
ZoneVcenterConfig.prototype['enableVnc'] = undefined;

/**
 * @member {String} _enableDiskTypeSelection
 */
ZoneVcenterConfig.prototype['_enableDiskTypeSelection'] = undefined;

/**
 * @member {String} _enableNetworkTypeSelection
 */
ZoneVcenterConfig.prototype['_enableNetworkTypeSelection'] = undefined;

/**
 * @member {String} diskStorageType
 */
ZoneVcenterConfig.prototype['diskStorageType'] = undefined;

/**
 * @member {String} applianceUrl
 */
ZoneVcenterConfig.prototype['applianceUrl'] = undefined;

/**
 * @member {String} datacenterName
 */
ZoneVcenterConfig.prototype['datacenterName'] = undefined;

/**
 * @member {String} networkServer.id
 */
ZoneVcenterConfig.prototype['networkServer.id'] = undefined;

/**
 * @member {module:model/ZoneVcenterConfigNetworkServer} networkServer
 */
ZoneVcenterConfig.prototype['networkServer'] = undefined;

/**
 * @member {String} securityMode
 */
ZoneVcenterConfig.prototype['securityMode'] = undefined;

/**
 * @member {String} certificateProvider
 */
ZoneVcenterConfig.prototype['certificateProvider'] = undefined;

/**
 * @member {String} backupMode
 */
ZoneVcenterConfig.prototype['backupMode'] = undefined;

/**
 * @member {String} replicationMode
 */
ZoneVcenterConfig.prototype['replicationMode'] = undefined;

/**
 * @member {String} dnsIntegrationId
 */
ZoneVcenterConfig.prototype['dnsIntegrationId'] = undefined;

/**
 * @member {String} configCmdbId
 */
ZoneVcenterConfig.prototype['configCmdbId'] = undefined;

/**
 * @member {String} configManagementId
 */
ZoneVcenterConfig.prototype['configManagementId'] = undefined;

/**
 * @member {String} configCmId
 */
ZoneVcenterConfig.prototype['configCmId'] = undefined;

/**
 * @member {String} securityServer
 */
ZoneVcenterConfig.prototype['securityServer'] = undefined;

/**
 * @member {String} serviceRegistryId
 */
ZoneVcenterConfig.prototype['serviceRegistryId'] = undefined;

/**
 * @member {String} enableDiskTypeSelection
 */
ZoneVcenterConfig.prototype['enableDiskTypeSelection'] = undefined;

/**
 * @member {String} kubeUrl
 */
ZoneVcenterConfig.prototype['kubeUrl'] = undefined;

/**
 * @member {String} apiVersion
 */
ZoneVcenterConfig.prototype['apiVersion'] = undefined;

/**
 * @member {String} datacenterId
 */
ZoneVcenterConfig.prototype['datacenterId'] = undefined;

/**
 * @member {Boolean} configCmdbDiscovery
 */
ZoneVcenterConfig.prototype['configCmdbDiscovery'] = undefined;

/**
 * @member {String} distributedWorkerId
 */
ZoneVcenterConfig.prototype['distributedWorkerId'] = undefined;

/**
 * @member {String} passwordHash
 */
ZoneVcenterConfig.prototype['passwordHash'] = undefined;






export default ZoneVcenterConfig;

