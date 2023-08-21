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
 * The ClusterLayoutCreateMasters model module.
 * @module model/ClusterLayoutCreateMasters
 * @version 6.2.1
 */
class ClusterLayoutCreateMasters {
    /**
     * Constructs a new <code>ClusterLayoutCreateMasters</code>.
     * @alias module:model/ClusterLayoutCreateMasters
     * @param containerType {module:model/ApiStorageVolumesStorageVolumeStorageServer} 
     */
    constructor(containerType) { 
        
        ClusterLayoutCreateMasters.initialize(this, containerType);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, containerType) { 
        obj['containerType'] = containerType;
    }

    /**
     * Constructs a <code>ClusterLayoutCreateMasters</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ClusterLayoutCreateMasters} obj Optional instance to populate.
     * @return {module:model/ClusterLayoutCreateMasters} The populated <code>ClusterLayoutCreateMasters</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ClusterLayoutCreateMasters();

            if (data.hasOwnProperty('nodeCount')) {
                obj['nodeCount'] = ApiClient.convertToType(data['nodeCount'], 'Number');
            }
            if (data.hasOwnProperty('containerType')) {
                obj['containerType'] = ApiStorageVolumesStorageVolumeStorageServer.constructFromObject(data['containerType']);
            }
        }
        return obj;
    }


}

/**
 * Number of nodes
 * @member {Number} nodeCount
 * @default 1
 */
ClusterLayoutCreateMasters.prototype['nodeCount'] = 1;

/**
 * @member {module:model/ApiStorageVolumesStorageVolumeStorageServer} containerType
 */
ClusterLayoutCreateMasters.prototype['containerType'] = undefined;






export default ClusterLayoutCreateMasters;

