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
 * The GuidanceVmwareSizingConfig model module.
 * @module model/GuidanceVmwareSizingConfig
 * @version 6.2.1
 */
class GuidanceVmwareSizingConfig {
    /**
     * Constructs a new <code>GuidanceVmwareSizingConfig</code>.
     * @alias module:model/GuidanceVmwareSizingConfig
     */
    constructor() { 
        
        GuidanceVmwareSizingConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>GuidanceVmwareSizingConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/GuidanceVmwareSizingConfig} obj Optional instance to populate.
     * @return {module:model/GuidanceVmwareSizingConfig} The populated <code>GuidanceVmwareSizingConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new GuidanceVmwareSizingConfig();

            if (data.hasOwnProperty('exists')) {
                obj['exists'] = ApiClient.convertToType(data['exists'], 'Boolean');
            }
            if (data.hasOwnProperty('objectId')) {
                obj['objectId'] = ApiClient.convertToType(data['objectId'], 'Number');
            }
            if (data.hasOwnProperty('cpuTotalTimeCount')) {
                obj['cpuTotalTimeCount'] = ApiClient.convertToType(data['cpuTotalTimeCount'], 'Number');
            }
            if (data.hasOwnProperty('cpuTotalTimeMin')) {
                obj['cpuTotalTimeMin'] = ApiClient.convertToType(data['cpuTotalTimeMin'], 'Number');
            }
            if (data.hasOwnProperty('cpuTotalTimeMax')) {
                obj['cpuTotalTimeMax'] = ApiClient.convertToType(data['cpuTotalTimeMax'], 'Number');
            }
            if (data.hasOwnProperty('cpuTotalTimeAvg')) {
                obj['cpuTotalTimeAvg'] = ApiClient.convertToType(data['cpuTotalTimeAvg'], 'Number');
            }
            if (data.hasOwnProperty('cpuTotalTimeSum')) {
                obj['cpuTotalTimeSum'] = ApiClient.convertToType(data['cpuTotalTimeSum'], 'Number');
            }
            if (data.hasOwnProperty('cpuIdleTimeCount')) {
                obj['cpuIdleTimeCount'] = ApiClient.convertToType(data['cpuIdleTimeCount'], 'Number');
            }
            if (data.hasOwnProperty('cpuIdleTimeMin')) {
                obj['cpuIdleTimeMin'] = ApiClient.convertToType(data['cpuIdleTimeMin'], 'Number');
            }
            if (data.hasOwnProperty('cpuIdleTimeMax')) {
                obj['cpuIdleTimeMax'] = ApiClient.convertToType(data['cpuIdleTimeMax'], 'Number');
            }
            if (data.hasOwnProperty('cpuIdleTimeAvg')) {
                obj['cpuIdleTimeAvg'] = ApiClient.convertToType(data['cpuIdleTimeAvg'], 'Number');
            }
            if (data.hasOwnProperty('cpuIdleTimeSum')) {
                obj['cpuIdleTimeSum'] = ApiClient.convertToType(data['cpuIdleTimeSum'], 'Number');
            }
            if (data.hasOwnProperty('cpuUsageCount')) {
                obj['cpuUsageCount'] = ApiClient.convertToType(data['cpuUsageCount'], 'Number');
            }
            if (data.hasOwnProperty('cpuUsageMin')) {
                obj['cpuUsageMin'] = ApiClient.convertToType(data['cpuUsageMin'], 'Number');
            }
            if (data.hasOwnProperty('cpuUsageMax')) {
                obj['cpuUsageMax'] = ApiClient.convertToType(data['cpuUsageMax'], 'Number');
            }
            if (data.hasOwnProperty('cpuUsageAvg')) {
                obj['cpuUsageAvg'] = ApiClient.convertToType(data['cpuUsageAvg'], 'Number');
            }
            if (data.hasOwnProperty('cpuUsageSum')) {
                obj['cpuUsageSum'] = ApiClient.convertToType(data['cpuUsageSum'], 'Number');
            }
            if (data.hasOwnProperty('maxMemoryCount')) {
                obj['maxMemoryCount'] = ApiClient.convertToType(data['maxMemoryCount'], 'Number');
            }
            if (data.hasOwnProperty('maxMemoryMin')) {
                obj['maxMemoryMin'] = ApiClient.convertToType(data['maxMemoryMin'], 'Number');
            }
            if (data.hasOwnProperty('maxMemoryMax')) {
                obj['maxMemoryMax'] = ApiClient.convertToType(data['maxMemoryMax'], 'Number');
            }
            if (data.hasOwnProperty('maxMemoryAvg')) {
                obj['maxMemoryAvg'] = ApiClient.convertToType(data['maxMemoryAvg'], 'Number');
            }
            if (data.hasOwnProperty('maxMemorySum')) {
                obj['maxMemorySum'] = ApiClient.convertToType(data['maxMemorySum'], 'Number');
            }
            if (data.hasOwnProperty('cpuUserTimeCount')) {
                obj['cpuUserTimeCount'] = ApiClient.convertToType(data['cpuUserTimeCount'], 'Number');
            }
            if (data.hasOwnProperty('cpuUserTimeMin')) {
                obj['cpuUserTimeMin'] = ApiClient.convertToType(data['cpuUserTimeMin'], 'Number');
            }
            if (data.hasOwnProperty('cpuUserTimeMax')) {
                obj['cpuUserTimeMax'] = ApiClient.convertToType(data['cpuUserTimeMax'], 'Number');
            }
            if (data.hasOwnProperty('cpuUserTimeAvg')) {
                obj['cpuUserTimeAvg'] = ApiClient.convertToType(data['cpuUserTimeAvg'], 'Number');
            }
            if (data.hasOwnProperty('cpuUserTimeSum')) {
                obj['cpuUserTimeSum'] = ApiClient.convertToType(data['cpuUserTimeSum'], 'Number');
            }
            if (data.hasOwnProperty('cpuSystemTimeCount')) {
                obj['cpuSystemTimeCount'] = ApiClient.convertToType(data['cpuSystemTimeCount'], 'Number');
            }
            if (data.hasOwnProperty('cpuSystemTimeMin')) {
                obj['cpuSystemTimeMin'] = ApiClient.convertToType(data['cpuSystemTimeMin'], 'Number');
            }
            if (data.hasOwnProperty('cpuSystemTimeMax')) {
                obj['cpuSystemTimeMax'] = ApiClient.convertToType(data['cpuSystemTimeMax'], 'Number');
            }
            if (data.hasOwnProperty('cpuSystemTimeAvg')) {
                obj['cpuSystemTimeAvg'] = ApiClient.convertToType(data['cpuSystemTimeAvg'], 'Number');
            }
            if (data.hasOwnProperty('cpuSystemTimeSum')) {
                obj['cpuSystemTimeSum'] = ApiClient.convertToType(data['cpuSystemTimeSum'], 'Number');
            }
            if (data.hasOwnProperty('usedMemoryCount')) {
                obj['usedMemoryCount'] = ApiClient.convertToType(data['usedMemoryCount'], 'Number');
            }
            if (data.hasOwnProperty('usedMemoryMin')) {
                obj['usedMemoryMin'] = ApiClient.convertToType(data['usedMemoryMin'], 'Number');
            }
            if (data.hasOwnProperty('usedMemoryMax')) {
                obj['usedMemoryMax'] = ApiClient.convertToType(data['usedMemoryMax'], 'Number');
            }
            if (data.hasOwnProperty('usedMemoryAvg')) {
                obj['usedMemoryAvg'] = ApiClient.convertToType(data['usedMemoryAvg'], 'Number');
            }
            if (data.hasOwnProperty('usedMemorySum')) {
                obj['usedMemorySum'] = ApiClient.convertToType(data['usedMemorySum'], 'Number');
            }
            if (data.hasOwnProperty('freeMemoryCount')) {
                obj['freeMemoryCount'] = ApiClient.convertToType(data['freeMemoryCount'], 'Number');
            }
            if (data.hasOwnProperty('freeMemoryMin')) {
                obj['freeMemoryMin'] = ApiClient.convertToType(data['freeMemoryMin'], 'Number');
            }
            if (data.hasOwnProperty('freeMemoryMax')) {
                obj['freeMemoryMax'] = ApiClient.convertToType(data['freeMemoryMax'], 'Number');
            }
            if (data.hasOwnProperty('freeMemoryAvg')) {
                obj['freeMemoryAvg'] = ApiClient.convertToType(data['freeMemoryAvg'], 'Number');
            }
            if (data.hasOwnProperty('freeMemorySum')) {
                obj['freeMemorySum'] = ApiClient.convertToType(data['freeMemorySum'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Boolean} exists
 */
GuidanceVmwareSizingConfig.prototype['exists'] = undefined;

/**
 * @member {Number} objectId
 */
GuidanceVmwareSizingConfig.prototype['objectId'] = undefined;

/**
 * @member {Number} cpuTotalTimeCount
 */
GuidanceVmwareSizingConfig.prototype['cpuTotalTimeCount'] = undefined;

/**
 * @member {Number} cpuTotalTimeMin
 */
GuidanceVmwareSizingConfig.prototype['cpuTotalTimeMin'] = undefined;

/**
 * @member {Number} cpuTotalTimeMax
 */
GuidanceVmwareSizingConfig.prototype['cpuTotalTimeMax'] = undefined;

/**
 * @member {Number} cpuTotalTimeAvg
 */
GuidanceVmwareSizingConfig.prototype['cpuTotalTimeAvg'] = undefined;

/**
 * @member {Number} cpuTotalTimeSum
 */
GuidanceVmwareSizingConfig.prototype['cpuTotalTimeSum'] = undefined;

/**
 * @member {Number} cpuIdleTimeCount
 */
GuidanceVmwareSizingConfig.prototype['cpuIdleTimeCount'] = undefined;

/**
 * @member {Number} cpuIdleTimeMin
 */
GuidanceVmwareSizingConfig.prototype['cpuIdleTimeMin'] = undefined;

/**
 * @member {Number} cpuIdleTimeMax
 */
GuidanceVmwareSizingConfig.prototype['cpuIdleTimeMax'] = undefined;

/**
 * @member {Number} cpuIdleTimeAvg
 */
GuidanceVmwareSizingConfig.prototype['cpuIdleTimeAvg'] = undefined;

/**
 * @member {Number} cpuIdleTimeSum
 */
GuidanceVmwareSizingConfig.prototype['cpuIdleTimeSum'] = undefined;

/**
 * @member {Number} cpuUsageCount
 */
GuidanceVmwareSizingConfig.prototype['cpuUsageCount'] = undefined;

/**
 * @member {Number} cpuUsageMin
 */
GuidanceVmwareSizingConfig.prototype['cpuUsageMin'] = undefined;

/**
 * @member {Number} cpuUsageMax
 */
GuidanceVmwareSizingConfig.prototype['cpuUsageMax'] = undefined;

/**
 * @member {Number} cpuUsageAvg
 */
GuidanceVmwareSizingConfig.prototype['cpuUsageAvg'] = undefined;

/**
 * @member {Number} cpuUsageSum
 */
GuidanceVmwareSizingConfig.prototype['cpuUsageSum'] = undefined;

/**
 * @member {Number} maxMemoryCount
 */
GuidanceVmwareSizingConfig.prototype['maxMemoryCount'] = undefined;

/**
 * @member {Number} maxMemoryMin
 */
GuidanceVmwareSizingConfig.prototype['maxMemoryMin'] = undefined;

/**
 * @member {Number} maxMemoryMax
 */
GuidanceVmwareSizingConfig.prototype['maxMemoryMax'] = undefined;

/**
 * @member {Number} maxMemoryAvg
 */
GuidanceVmwareSizingConfig.prototype['maxMemoryAvg'] = undefined;

/**
 * @member {Number} maxMemorySum
 */
GuidanceVmwareSizingConfig.prototype['maxMemorySum'] = undefined;

/**
 * @member {Number} cpuUserTimeCount
 */
GuidanceVmwareSizingConfig.prototype['cpuUserTimeCount'] = undefined;

/**
 * @member {Number} cpuUserTimeMin
 */
GuidanceVmwareSizingConfig.prototype['cpuUserTimeMin'] = undefined;

/**
 * @member {Number} cpuUserTimeMax
 */
GuidanceVmwareSizingConfig.prototype['cpuUserTimeMax'] = undefined;

/**
 * @member {Number} cpuUserTimeAvg
 */
GuidanceVmwareSizingConfig.prototype['cpuUserTimeAvg'] = undefined;

/**
 * @member {Number} cpuUserTimeSum
 */
GuidanceVmwareSizingConfig.prototype['cpuUserTimeSum'] = undefined;

/**
 * @member {Number} cpuSystemTimeCount
 */
GuidanceVmwareSizingConfig.prototype['cpuSystemTimeCount'] = undefined;

/**
 * @member {Number} cpuSystemTimeMin
 */
GuidanceVmwareSizingConfig.prototype['cpuSystemTimeMin'] = undefined;

/**
 * @member {Number} cpuSystemTimeMax
 */
GuidanceVmwareSizingConfig.prototype['cpuSystemTimeMax'] = undefined;

/**
 * @member {Number} cpuSystemTimeAvg
 */
GuidanceVmwareSizingConfig.prototype['cpuSystemTimeAvg'] = undefined;

/**
 * @member {Number} cpuSystemTimeSum
 */
GuidanceVmwareSizingConfig.prototype['cpuSystemTimeSum'] = undefined;

/**
 * @member {Number} usedMemoryCount
 */
GuidanceVmwareSizingConfig.prototype['usedMemoryCount'] = undefined;

/**
 * @member {Number} usedMemoryMin
 */
GuidanceVmwareSizingConfig.prototype['usedMemoryMin'] = undefined;

/**
 * @member {Number} usedMemoryMax
 */
GuidanceVmwareSizingConfig.prototype['usedMemoryMax'] = undefined;

/**
 * @member {Number} usedMemoryAvg
 */
GuidanceVmwareSizingConfig.prototype['usedMemoryAvg'] = undefined;

/**
 * @member {Number} usedMemorySum
 */
GuidanceVmwareSizingConfig.prototype['usedMemorySum'] = undefined;

/**
 * @member {Number} freeMemoryCount
 */
GuidanceVmwareSizingConfig.prototype['freeMemoryCount'] = undefined;

/**
 * @member {Number} freeMemoryMin
 */
GuidanceVmwareSizingConfig.prototype['freeMemoryMin'] = undefined;

/**
 * @member {Number} freeMemoryMax
 */
GuidanceVmwareSizingConfig.prototype['freeMemoryMax'] = undefined;

/**
 * @member {Number} freeMemoryAvg
 */
GuidanceVmwareSizingConfig.prototype['freeMemoryAvg'] = undefined;

/**
 * @member {Number} freeMemorySum
 */
GuidanceVmwareSizingConfig.prototype['freeMemorySum'] = undefined;






export default GuidanceVmwareSizingConfig;
