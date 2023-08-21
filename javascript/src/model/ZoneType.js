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
import ZoneTypeOptionTypes1 from './ZoneTypeOptionTypes1';
import ZoneTypeServerTypes from './ZoneTypeServerTypes';

/**
 * The ZoneType model module.
 * @module model/ZoneType
 * @version 6.2.1
 */
class ZoneType {
    /**
     * Constructs a new <code>ZoneType</code>.
     * @alias module:model/ZoneType
     */
    constructor() { 
        
        ZoneType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ZoneType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ZoneType} obj Optional instance to populate.
     * @return {module:model/ZoneType} The populated <code>ZoneType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ZoneType();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('provision')) {
                obj['provision'] = ApiClient.convertToType(data['provision'], 'Boolean');
            }
            if (data.hasOwnProperty('autoCapacity')) {
                obj['autoCapacity'] = ApiClient.convertToType(data['autoCapacity'], 'Boolean');
            }
            if (data.hasOwnProperty('migrationTarget')) {
                obj['migrationTarget'] = ApiClient.convertToType(data['migrationTarget'], 'Boolean');
            }
            if (data.hasOwnProperty('hasDatastores')) {
                obj['hasDatastores'] = ApiClient.convertToType(data['hasDatastores'], 'Boolean');
            }
            if (data.hasOwnProperty('hasNetworks')) {
                obj['hasNetworks'] = ApiClient.convertToType(data['hasNetworks'], 'Boolean');
            }
            if (data.hasOwnProperty('hasResourcePools')) {
                obj['hasResourcePools'] = ApiClient.convertToType(data['hasResourcePools'], 'Boolean');
            }
            if (data.hasOwnProperty('hasSecurityGroups')) {
                obj['hasSecurityGroups'] = ApiClient.convertToType(data['hasSecurityGroups'], 'Boolean');
            }
            if (data.hasOwnProperty('hasContainers')) {
                obj['hasContainers'] = ApiClient.convertToType(data['hasContainers'], 'Boolean');
            }
            if (data.hasOwnProperty('hasBareMetal')) {
                obj['hasBareMetal'] = ApiClient.convertToType(data['hasBareMetal'], 'Boolean');
            }
            if (data.hasOwnProperty('hasServices')) {
                obj['hasServices'] = ApiClient.convertToType(data['hasServices'], 'Boolean');
            }
            if (data.hasOwnProperty('hasFunctions')) {
                obj['hasFunctions'] = ApiClient.convertToType(data['hasFunctions'], 'Boolean');
            }
            if (data.hasOwnProperty('hasJobs')) {
                obj['hasJobs'] = ApiClient.convertToType(data['hasJobs'], 'Boolean');
            }
            if (data.hasOwnProperty('hasDiscovery')) {
                obj['hasDiscovery'] = ApiClient.convertToType(data['hasDiscovery'], 'Boolean');
            }
            if (data.hasOwnProperty('hasCloudInit')) {
                obj['hasCloudInit'] = ApiClient.convertToType(data['hasCloudInit'], 'Boolean');
            }
            if (data.hasOwnProperty('hasFolders')) {
                obj['hasFolders'] = ApiClient.convertToType(data['hasFolders'], 'Boolean');
            }
            if (data.hasOwnProperty('hasFloatingIps')) {
                obj['hasFloatingIps'] = ApiClient.convertToType(data['hasFloatingIps'], 'Boolean');
            }
            if (data.hasOwnProperty('hasMarketplace')) {
                obj['hasMarketplace'] = ApiClient.convertToType(data['hasMarketplace'], 'Boolean');
            }
            if (data.hasOwnProperty('canCreateResourcePools')) {
                obj['canCreateResourcePools'] = ApiClient.convertToType(data['canCreateResourcePools'], 'Boolean');
            }
            if (data.hasOwnProperty('canDeleteResourcePools')) {
                obj['canDeleteResourcePools'] = ApiClient.convertToType(data['canDeleteResourcePools'], 'Boolean');
            }
            if (data.hasOwnProperty('canCreateDatastores')) {
                obj['canCreateDatastores'] = ApiClient.convertToType(data['canCreateDatastores'], 'Boolean');
            }
            if (data.hasOwnProperty('canCreateNetworks')) {
                obj['canCreateNetworks'] = ApiClient.convertToType(data['canCreateNetworks'], 'Boolean');
            }
            if (data.hasOwnProperty('canChooseContainerMode')) {
                obj['canChooseContainerMode'] = ApiClient.convertToType(data['canChooseContainerMode'], 'Boolean');
            }
            if (data.hasOwnProperty('provisionRequiresResourcePool')) {
                obj['provisionRequiresResourcePool'] = ApiClient.convertToType(data['provisionRequiresResourcePool'], 'Boolean');
            }
            if (data.hasOwnProperty('supportsDistributedWorker')) {
                obj['supportsDistributedWorker'] = ApiClient.convertToType(data['supportsDistributedWorker'], 'Boolean');
            }
            if (data.hasOwnProperty('cloud')) {
                obj['cloud'] = ApiClient.convertToType(data['cloud'], 'String');
            }
            if (data.hasOwnProperty('provisionTypes')) {
                obj['provisionTypes'] = ApiClient.convertToType(data['provisionTypes'], ['Number']);
            }
            if (data.hasOwnProperty('zoneInstanceTypeLayoutId')) {
                obj['zoneInstanceTypeLayoutId'] = ApiClient.convertToType(data['zoneInstanceTypeLayoutId'], 'Number');
            }
            if (data.hasOwnProperty('serverTypes')) {
                obj['serverTypes'] = ApiClient.convertToType(data['serverTypes'], [ZoneTypeServerTypes]);
            }
            if (data.hasOwnProperty('optionTypes')) {
                obj['optionTypes'] = ApiClient.convertToType(data['optionTypes'], [ZoneTypeOptionTypes1]);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ZoneType.prototype['id'] = undefined;

/**
 * @member {String} name
 */
ZoneType.prototype['name'] = undefined;

/**
 * @member {String} code
 */
ZoneType.prototype['code'] = undefined;

/**
 * @member {Boolean} enabled
 */
ZoneType.prototype['enabled'] = undefined;

/**
 * @member {Boolean} provision
 */
ZoneType.prototype['provision'] = undefined;

/**
 * @member {Boolean} autoCapacity
 */
ZoneType.prototype['autoCapacity'] = undefined;

/**
 * @member {Boolean} migrationTarget
 */
ZoneType.prototype['migrationTarget'] = undefined;

/**
 * @member {Boolean} hasDatastores
 */
ZoneType.prototype['hasDatastores'] = undefined;

/**
 * @member {Boolean} hasNetworks
 */
ZoneType.prototype['hasNetworks'] = undefined;

/**
 * @member {Boolean} hasResourcePools
 */
ZoneType.prototype['hasResourcePools'] = undefined;

/**
 * @member {Boolean} hasSecurityGroups
 */
ZoneType.prototype['hasSecurityGroups'] = undefined;

/**
 * @member {Boolean} hasContainers
 */
ZoneType.prototype['hasContainers'] = undefined;

/**
 * @member {Boolean} hasBareMetal
 */
ZoneType.prototype['hasBareMetal'] = undefined;

/**
 * @member {Boolean} hasServices
 */
ZoneType.prototype['hasServices'] = undefined;

/**
 * @member {Boolean} hasFunctions
 */
ZoneType.prototype['hasFunctions'] = undefined;

/**
 * @member {Boolean} hasJobs
 */
ZoneType.prototype['hasJobs'] = undefined;

/**
 * @member {Boolean} hasDiscovery
 */
ZoneType.prototype['hasDiscovery'] = undefined;

/**
 * @member {Boolean} hasCloudInit
 */
ZoneType.prototype['hasCloudInit'] = undefined;

/**
 * @member {Boolean} hasFolders
 */
ZoneType.prototype['hasFolders'] = undefined;

/**
 * @member {Boolean} hasFloatingIps
 */
ZoneType.prototype['hasFloatingIps'] = undefined;

/**
 * @member {Boolean} hasMarketplace
 */
ZoneType.prototype['hasMarketplace'] = undefined;

/**
 * @member {Boolean} canCreateResourcePools
 */
ZoneType.prototype['canCreateResourcePools'] = undefined;

/**
 * @member {Boolean} canDeleteResourcePools
 */
ZoneType.prototype['canDeleteResourcePools'] = undefined;

/**
 * @member {Boolean} canCreateDatastores
 */
ZoneType.prototype['canCreateDatastores'] = undefined;

/**
 * @member {Boolean} canCreateNetworks
 */
ZoneType.prototype['canCreateNetworks'] = undefined;

/**
 * @member {Boolean} canChooseContainerMode
 */
ZoneType.prototype['canChooseContainerMode'] = undefined;

/**
 * @member {Boolean} provisionRequiresResourcePool
 */
ZoneType.prototype['provisionRequiresResourcePool'] = undefined;

/**
 * @member {Boolean} supportsDistributedWorker
 */
ZoneType.prototype['supportsDistributedWorker'] = undefined;

/**
 * @member {String} cloud
 */
ZoneType.prototype['cloud'] = undefined;

/**
 * @member {Array.<Number>} provisionTypes
 */
ZoneType.prototype['provisionTypes'] = undefined;

/**
 * @member {Number} zoneInstanceTypeLayoutId
 */
ZoneType.prototype['zoneInstanceTypeLayoutId'] = undefined;

/**
 * @member {Array.<module:model/ZoneTypeServerTypes>} serverTypes
 */
ZoneType.prototype['serverTypes'] = undefined;

/**
 * @member {Array.<module:model/ZoneTypeOptionTypes1>} optionTypes
 */
ZoneType.prototype['optionTypes'] = undefined;






export default ZoneType;

