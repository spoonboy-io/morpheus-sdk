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
import ApiStorageVolumesStorageVolumeStorageServer from './ApiStorageVolumesStorageVolumeStorageServer';

/**
 * The ApiStorageVolumesStorageVolume model module.
 * @module model/ApiStorageVolumesStorageVolume
 * @version 6.2.1
 */
class ApiStorageVolumesStorageVolume {
    /**
     * Constructs a new <code>ApiStorageVolumesStorageVolume</code>.
     * @alias module:model/ApiStorageVolumesStorageVolume
     * @param name {String} A unique name scoped to your account for the storage volume
     * @param type {String} Storage Type Code or ID
     * @param storageServer {module:model/ApiStorageVolumesStorageVolumeStorageServer} 
     * @param storageGroup {module:model/ApiStorageVolumesStorageVolumeStorageServer} 
     */
    constructor(name, type, storageServer, storageGroup) { 
        
        ApiStorageVolumesStorageVolume.initialize(this, name, type, storageServer, storageGroup);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name, type, storageServer, storageGroup) { 
        obj['name'] = name;
        obj['type'] = type;
        obj['storageServer'] = storageServer;
        obj['storageGroup'] = storageGroup;
    }

    /**
     * Constructs a <code>ApiStorageVolumesStorageVolume</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiStorageVolumesStorageVolume} obj Optional instance to populate.
     * @return {module:model/ApiStorageVolumesStorageVolume} The populated <code>ApiStorageVolumesStorageVolume</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiStorageVolumesStorageVolume();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], Object);
            }
            if (data.hasOwnProperty('storageServer')) {
                obj['storageServer'] = ApiStorageVolumesStorageVolumeStorageServer.constructFromObject(data['storageServer']);
            }
            if (data.hasOwnProperty('storageGroup')) {
                obj['storageGroup'] = ApiStorageVolumesStorageVolumeStorageServer.constructFromObject(data['storageGroup']);
            }
        }
        return obj;
    }


}

/**
 * A unique name scoped to your account for the storage volume
 * @member {String} name
 */
ApiStorageVolumesStorageVolume.prototype['name'] = undefined;

/**
 * Storage Type Code or ID
 * @member {String} type
 */
ApiStorageVolumesStorageVolume.prototype['type'] = undefined;

/**
 * Configuration object with parameters that vary by `type`.
 * @member {Object} config
 */
ApiStorageVolumesStorageVolume.prototype['config'] = undefined;

/**
 * @member {module:model/ApiStorageVolumesStorageVolumeStorageServer} storageServer
 */
ApiStorageVolumesStorageVolume.prototype['storageServer'] = undefined;

/**
 * @member {module:model/ApiStorageVolumesStorageVolumeStorageServer} storageGroup
 */
ApiStorageVolumesStorageVolume.prototype['storageGroup'] = undefined;






export default ApiStorageVolumesStorageVolume;

