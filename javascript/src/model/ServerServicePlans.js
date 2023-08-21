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
import InstanceServicePlanAutoOptions from './InstanceServicePlanAutoOptions';
import ServerServicePlansDatastores from './ServerServicePlansDatastores';

/**
 * The ServerServicePlans model module.
 * @module model/ServerServicePlans
 * @version 6.2.1
 */
class ServerServicePlans {
    /**
     * Constructs a new <code>ServerServicePlans</code>.
     * @alias module:model/ServerServicePlans
     */
    constructor() { 
        
        ServerServicePlans.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ServerServicePlans</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ServerServicePlans} obj Optional instance to populate.
     * @return {module:model/ServerServicePlans} The populated <code>ServerServicePlans</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ServerServicePlans();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('value')) {
                obj['value'] = ApiClient.convertToType(data['value'], 'Number');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('maxStorage')) {
                obj['maxStorage'] = ApiClient.convertToType(data['maxStorage'], 'Number');
            }
            if (data.hasOwnProperty('maxMemory')) {
                obj['maxMemory'] = ApiClient.convertToType(data['maxMemory'], 'Number');
            }
            if (data.hasOwnProperty('maxCpu')) {
                obj['maxCpu'] = ApiClient.convertToType(data['maxCpu'], 'Number');
            }
            if (data.hasOwnProperty('maxCores')) {
                obj['maxCores'] = ApiClient.convertToType(data['maxCores'], 'Number');
            }
            if (data.hasOwnProperty('maxDataStorage')) {
                obj['maxDataStorage'] = ApiClient.convertToType(data['maxDataStorage'], 'Number');
            }
            if (data.hasOwnProperty('customCpu')) {
                obj['customCpu'] = ApiClient.convertToType(data['customCpu'], 'Boolean');
            }
            if (data.hasOwnProperty('customMaxMemory')) {
                obj['customMaxMemory'] = ApiClient.convertToType(data['customMaxMemory'], 'Boolean');
            }
            if (data.hasOwnProperty('customMaxStorage')) {
                obj['customMaxStorage'] = ApiClient.convertToType(data['customMaxStorage'], 'Boolean');
            }
            if (data.hasOwnProperty('customMaxDataStorage')) {
                obj['customMaxDataStorage'] = ApiClient.convertToType(data['customMaxDataStorage'], 'Boolean');
            }
            if (data.hasOwnProperty('customCoresPerSocket')) {
                obj['customCoresPerSocket'] = ApiClient.convertToType(data['customCoresPerSocket'], 'Boolean');
            }
            if (data.hasOwnProperty('coresPerSocket')) {
                obj['coresPerSocket'] = ApiClient.convertToType(data['coresPerSocket'], 'Number');
            }
            if (data.hasOwnProperty('storageTypes')) {
                obj['storageTypes'] = ApiClient.convertToType(data['storageTypes'], [Object]);
            }
            if (data.hasOwnProperty('rootStorageTypes')) {
                obj['rootStorageTypes'] = ApiClient.convertToType(data['rootStorageTypes'], [Object]);
            }
            if (data.hasOwnProperty('addVolumes')) {
                obj['addVolumes'] = ApiClient.convertToType(data['addVolumes'], 'Boolean');
            }
            if (data.hasOwnProperty('customizeVolume')) {
                obj['customizeVolume'] = ApiClient.convertToType(data['customizeVolume'], 'Boolean');
            }
            if (data.hasOwnProperty('rootDiskCustomizable')) {
                obj['rootDiskCustomizable'] = ApiClient.convertToType(data['rootDiskCustomizable'], 'Boolean');
            }
            if (data.hasOwnProperty('hostDiskMode')) {
                obj['hostDiskMode'] = ApiClient.convertToType(data['hostDiskMode'], 'String');
            }
            if (data.hasOwnProperty('hasDatastore')) {
                obj['hasDatastore'] = ApiClient.convertToType(data['hasDatastore'], 'String');
            }
            if (data.hasOwnProperty('lvmSupported')) {
                obj['lvmSupported'] = ApiClient.convertToType(data['lvmSupported'], 'String');
            }
            if (data.hasOwnProperty('minDisk')) {
                obj['minDisk'] = ApiClient.convertToType(data['minDisk'], 'String');
            }
            if (data.hasOwnProperty('maxDisk')) {
                obj['maxDisk'] = ApiClient.convertToType(data['maxDisk'], 'String');
            }
            if (data.hasOwnProperty('datastores')) {
                obj['datastores'] = ServerServicePlansDatastores.constructFromObject(data['datastores']);
            }
            if (data.hasOwnProperty('supportsAutoDatastore')) {
                obj['supportsAutoDatastore'] = ApiClient.convertToType(data['supportsAutoDatastore'], 'Boolean');
            }
            if (data.hasOwnProperty('autoOptions')) {
                obj['autoOptions'] = ApiClient.convertToType(data['autoOptions'], [InstanceServicePlanAutoOptions]);
            }
            if (data.hasOwnProperty('cpuOptions')) {
                obj['cpuOptions'] = ApiClient.convertToType(data['cpuOptions'], [Object]);
            }
            if (data.hasOwnProperty('memoryOptions')) {
                obj['memoryOptions'] = ApiClient.convertToType(data['memoryOptions'], [Object]);
            }
            if (data.hasOwnProperty('rootCustomSizeOptions')) {
                obj['rootCustomSizeOptions'] = ApiClient.convertToType(data['rootCustomSizeOptions'], Object);
            }
            if (data.hasOwnProperty('customSizeOptions')) {
                obj['customSizeOptions'] = ApiClient.convertToType(data['customSizeOptions'], Object);
            }
            if (data.hasOwnProperty('customCores')) {
                obj['customCores'] = ApiClient.convertToType(data['customCores'], 'Boolean');
            }
            if (data.hasOwnProperty('maxDisks')) {
                obj['maxDisks'] = ApiClient.convertToType(data['maxDisks'], 'String');
            }
            if (data.hasOwnProperty('memorySizeType')) {
                obj['memorySizeType'] = ApiClient.convertToType(data['memorySizeType'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ServerServicePlans.prototype['id'] = undefined;

/**
 * @member {String} name
 */
ServerServicePlans.prototype['name'] = undefined;

/**
 * @member {Number} value
 */
ServerServicePlans.prototype['value'] = undefined;

/**
 * @member {String} code
 */
ServerServicePlans.prototype['code'] = undefined;

/**
 * @member {Number} maxStorage
 */
ServerServicePlans.prototype['maxStorage'] = undefined;

/**
 * @member {Number} maxMemory
 */
ServerServicePlans.prototype['maxMemory'] = undefined;

/**
 * @member {Number} maxCpu
 */
ServerServicePlans.prototype['maxCpu'] = undefined;

/**
 * @member {Number} maxCores
 */
ServerServicePlans.prototype['maxCores'] = undefined;

/**
 * @member {Number} maxDataStorage
 */
ServerServicePlans.prototype['maxDataStorage'] = undefined;

/**
 * @member {Boolean} customCpu
 */
ServerServicePlans.prototype['customCpu'] = undefined;

/**
 * @member {Boolean} customMaxMemory
 */
ServerServicePlans.prototype['customMaxMemory'] = undefined;

/**
 * @member {Boolean} customMaxStorage
 */
ServerServicePlans.prototype['customMaxStorage'] = undefined;

/**
 * @member {Boolean} customMaxDataStorage
 */
ServerServicePlans.prototype['customMaxDataStorage'] = undefined;

/**
 * @member {Boolean} customCoresPerSocket
 */
ServerServicePlans.prototype['customCoresPerSocket'] = undefined;

/**
 * @member {Number} coresPerSocket
 */
ServerServicePlans.prototype['coresPerSocket'] = undefined;

/**
 * @member {Array.<Object>} storageTypes
 */
ServerServicePlans.prototype['storageTypes'] = undefined;

/**
 * @member {Array.<Object>} rootStorageTypes
 */
ServerServicePlans.prototype['rootStorageTypes'] = undefined;

/**
 * @member {Boolean} addVolumes
 */
ServerServicePlans.prototype['addVolumes'] = undefined;

/**
 * @member {Boolean} customizeVolume
 */
ServerServicePlans.prototype['customizeVolume'] = undefined;

/**
 * @member {Boolean} rootDiskCustomizable
 */
ServerServicePlans.prototype['rootDiskCustomizable'] = undefined;

/**
 * @member {String} hostDiskMode
 */
ServerServicePlans.prototype['hostDiskMode'] = undefined;

/**
 * @member {String} hasDatastore
 */
ServerServicePlans.prototype['hasDatastore'] = undefined;

/**
 * @member {String} lvmSupported
 */
ServerServicePlans.prototype['lvmSupported'] = undefined;

/**
 * @member {String} minDisk
 */
ServerServicePlans.prototype['minDisk'] = undefined;

/**
 * @member {String} maxDisk
 */
ServerServicePlans.prototype['maxDisk'] = undefined;

/**
 * @member {module:model/ServerServicePlansDatastores} datastores
 */
ServerServicePlans.prototype['datastores'] = undefined;

/**
 * @member {Boolean} supportsAutoDatastore
 */
ServerServicePlans.prototype['supportsAutoDatastore'] = undefined;

/**
 * @member {Array.<module:model/InstanceServicePlanAutoOptions>} autoOptions
 */
ServerServicePlans.prototype['autoOptions'] = undefined;

/**
 * @member {Array.<Object>} cpuOptions
 */
ServerServicePlans.prototype['cpuOptions'] = undefined;

/**
 * @member {Array.<Object>} memoryOptions
 */
ServerServicePlans.prototype['memoryOptions'] = undefined;

/**
 * @member {Object} rootCustomSizeOptions
 */
ServerServicePlans.prototype['rootCustomSizeOptions'] = undefined;

/**
 * @member {Object} customSizeOptions
 */
ServerServicePlans.prototype['customSizeOptions'] = undefined;

/**
 * @member {Boolean} customCores
 */
ServerServicePlans.prototype['customCores'] = undefined;

/**
 * @member {String} maxDisks
 */
ServerServicePlans.prototype['maxDisks'] = undefined;

/**
 * @member {String} memorySizeType
 */
ServerServicePlans.prototype['memorySizeType'] = undefined;






export default ServerServicePlans;
