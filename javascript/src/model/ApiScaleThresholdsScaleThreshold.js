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
 * The ApiScaleThresholdsScaleThreshold model module.
 * @module model/ApiScaleThresholdsScaleThreshold
 * @version 6.2.1
 */
class ApiScaleThresholdsScaleThreshold {
    /**
     * Constructs a new <code>ApiScaleThresholdsScaleThreshold</code>.
     * @alias module:model/ApiScaleThresholdsScaleThreshold
     * @param name {String} A name for the scale threshold
     */
    constructor(name) { 
        
        ApiScaleThresholdsScaleThreshold.initialize(this, name);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name) { 
        obj['name'] = name;
    }

    /**
     * Constructs a <code>ApiScaleThresholdsScaleThreshold</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiScaleThresholdsScaleThreshold} obj Optional instance to populate.
     * @return {module:model/ApiScaleThresholdsScaleThreshold} The populated <code>ApiScaleThresholdsScaleThreshold</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiScaleThresholdsScaleThreshold();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('autoUp')) {
                obj['autoUp'] = ApiClient.convertToType(data['autoUp'], 'Boolean');
            }
            if (data.hasOwnProperty('autoDown')) {
                obj['autoDown'] = ApiClient.convertToType(data['autoDown'], 'Boolean');
            }
            if (data.hasOwnProperty('minCount')) {
                obj['minCount'] = ApiClient.convertToType(data['minCount'], 'Number');
            }
            if (data.hasOwnProperty('maxCount')) {
                obj['maxCount'] = ApiClient.convertToType(data['maxCount'], 'Number');
            }
            if (data.hasOwnProperty('cpuEnabled')) {
                obj['cpuEnabled'] = ApiClient.convertToType(data['cpuEnabled'], 'Boolean');
            }
            if (data.hasOwnProperty('minCpu')) {
                obj['minCpu'] = ApiClient.convertToType(data['minCpu'], 'Number');
            }
            if (data.hasOwnProperty('maxCpu')) {
                obj['maxCpu'] = ApiClient.convertToType(data['maxCpu'], 'Number');
            }
            if (data.hasOwnProperty('memoryEnabled')) {
                obj['memoryEnabled'] = ApiClient.convertToType(data['memoryEnabled'], 'Boolean');
            }
            if (data.hasOwnProperty('minMemory')) {
                obj['minMemory'] = ApiClient.convertToType(data['minMemory'], 'Number');
            }
            if (data.hasOwnProperty('maxMemory')) {
                obj['maxMemory'] = ApiClient.convertToType(data['maxMemory'], 'Number');
            }
            if (data.hasOwnProperty('diskEnabled')) {
                obj['diskEnabled'] = ApiClient.convertToType(data['diskEnabled'], 'Boolean');
            }
            if (data.hasOwnProperty('minDisk')) {
                obj['minDisk'] = ApiClient.convertToType(data['minDisk'], 'Number');
            }
            if (data.hasOwnProperty('maxDisk')) {
                obj['maxDisk'] = ApiClient.convertToType(data['maxDisk'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * A name for the scale threshold
 * @member {String} name
 */
ApiScaleThresholdsScaleThreshold.prototype['name'] = undefined;

/**
 * Auto Upscale
 * @member {Boolean} autoUp
 * @default false
 */
ApiScaleThresholdsScaleThreshold.prototype['autoUp'] = false;

/**
 * Auto Downscale
 * @member {Boolean} autoDown
 * @default false
 */
ApiScaleThresholdsScaleThreshold.prototype['autoDown'] = false;

/**
 * The minimum number of nodes to scale down to
 * @member {Number} minCount
 */
ApiScaleThresholdsScaleThreshold.prototype['minCount'] = undefined;

/**
 * The maximum number of nodes to scale up to
 * @member {Number} maxCount
 */
ApiScaleThresholdsScaleThreshold.prototype['maxCount'] = undefined;

/**
 * Enable CPU Threshold
 * @member {Boolean} cpuEnabled
 * @default false
 */
ApiScaleThresholdsScaleThreshold.prototype['cpuEnabled'] = false;

/**
 * Min CPU (%)
 * @member {Number} minCpu
 * @default 0
 */
ApiScaleThresholdsScaleThreshold.prototype['minCpu'] = 0;

/**
 * Max CPU (%)
 * @member {Number} maxCpu
 * @default 0
 */
ApiScaleThresholdsScaleThreshold.prototype['maxCpu'] = 0;

/**
 * Enable Memory Threshold
 * @member {Boolean} memoryEnabled
 * @default false
 */
ApiScaleThresholdsScaleThreshold.prototype['memoryEnabled'] = false;

/**
 * Min Memory (%)
 * @member {Number} minMemory
 * @default 0
 */
ApiScaleThresholdsScaleThreshold.prototype['minMemory'] = 0;

/**
 * Max Memory (%)
 * @member {Number} maxMemory
 * @default 0
 */
ApiScaleThresholdsScaleThreshold.prototype['maxMemory'] = 0;

/**
 * Enable Disk Threshold
 * @member {Boolean} diskEnabled
 * @default false
 */
ApiScaleThresholdsScaleThreshold.prototype['diskEnabled'] = false;

/**
 * Min Disk (%)
 * @member {Number} minDisk
 * @default 0
 */
ApiScaleThresholdsScaleThreshold.prototype['minDisk'] = 0;

/**
 * Max Disk (%)
 * @member {Number} maxDisk
 * @default 0
 */
ApiScaleThresholdsScaleThreshold.prototype['maxDisk'] = 0;






export default ApiScaleThresholdsScaleThreshold;

