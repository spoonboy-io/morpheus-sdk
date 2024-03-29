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

/**
 * The ClustersWorkerStats model module.
 * @module model/ClustersWorkerStats
 * @version 6.2.1
 */
class ClustersWorkerStats {
    /**
     * Constructs a new <code>ClustersWorkerStats</code>.
     * @alias module:model/ClustersWorkerStats
     */
    constructor() { 
        
        ClustersWorkerStats.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ClustersWorkerStats</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ClustersWorkerStats} obj Optional instance to populate.
     * @return {module:model/ClustersWorkerStats} The populated <code>ClustersWorkerStats</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ClustersWorkerStats();

            if (data.hasOwnProperty('usedStorage')) {
                obj['usedStorage'] = ApiClient.convertToType(data['usedStorage'], 'Number');
            }
            if (data.hasOwnProperty('maxStorage')) {
                obj['maxStorage'] = ApiClient.convertToType(data['maxStorage'], 'Number');
            }
            if (data.hasOwnProperty('usedMemory')) {
                obj['usedMemory'] = ApiClient.convertToType(data['usedMemory'], 'Number');
            }
            if (data.hasOwnProperty('maxMemory')) {
                obj['maxMemory'] = ApiClient.convertToType(data['maxMemory'], 'Number');
            }
            if (data.hasOwnProperty('usedCpu')) {
                obj['usedCpu'] = ApiClient.convertToType(data['usedCpu'], 'Number');
            }
            if (data.hasOwnProperty('cpuUsage')) {
                obj['cpuUsage'] = ApiClient.convertToType(data['cpuUsage'], 'Number');
            }
            if (data.hasOwnProperty('cpuUsagePeak')) {
                obj['cpuUsagePeak'] = ApiClient.convertToType(data['cpuUsagePeak'], 'Number');
            }
            if (data.hasOwnProperty('cpuUsageAvg')) {
                obj['cpuUsageAvg'] = ApiClient.convertToType(data['cpuUsageAvg'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} usedStorage
 */
ClustersWorkerStats.prototype['usedStorage'] = undefined;

/**
 * @member {Number} maxStorage
 */
ClustersWorkerStats.prototype['maxStorage'] = undefined;

/**
 * @member {Number} usedMemory
 */
ClustersWorkerStats.prototype['usedMemory'] = undefined;

/**
 * @member {Number} maxMemory
 */
ClustersWorkerStats.prototype['maxMemory'] = undefined;

/**
 * @member {Number} usedCpu
 */
ClustersWorkerStats.prototype['usedCpu'] = undefined;

/**
 * @member {Number} cpuUsage
 */
ClustersWorkerStats.prototype['cpuUsage'] = undefined;

/**
 * @member {Number} cpuUsagePeak
 */
ClustersWorkerStats.prototype['cpuUsagePeak'] = undefined;

/**
 * @member {Number} cpuUsageAvg
 */
ClustersWorkerStats.prototype['cpuUsageAvg'] = undefined;






export default ClustersWorkerStats;

