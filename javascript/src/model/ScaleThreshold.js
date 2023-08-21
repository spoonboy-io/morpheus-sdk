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
 * The ScaleThreshold model module.
 * @module model/ScaleThreshold
 * @version 6.2.1
 */
class ScaleThreshold {
    /**
     * Constructs a new <code>ScaleThreshold</code>.
     * @alias module:model/ScaleThreshold
     */
    constructor() { 
        
        ScaleThreshold.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ScaleThreshold</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ScaleThreshold} obj Optional instance to populate.
     * @return {module:model/ScaleThreshold} The populated <code>ScaleThreshold</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ScaleThreshold();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
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
            if (data.hasOwnProperty('scaleIncrement')) {
                obj['scaleIncrement'] = ApiClient.convertToType(data['scaleIncrement'], 'Number');
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
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ScaleThreshold.prototype['id'] = undefined;

/**
 * @member {String} name
 */
ScaleThreshold.prototype['name'] = undefined;

/**
 * @member {String} type
 */
ScaleThreshold.prototype['type'] = undefined;

/**
 * @member {Boolean} autoUp
 */
ScaleThreshold.prototype['autoUp'] = undefined;

/**
 * @member {Boolean} autoDown
 */
ScaleThreshold.prototype['autoDown'] = undefined;

/**
 * @member {Number} minCount
 */
ScaleThreshold.prototype['minCount'] = undefined;

/**
 * @member {Number} maxCount
 */
ScaleThreshold.prototype['maxCount'] = undefined;

/**
 * @member {Number} scaleIncrement
 */
ScaleThreshold.prototype['scaleIncrement'] = undefined;

/**
 * @member {Boolean} cpuEnabled
 */
ScaleThreshold.prototype['cpuEnabled'] = undefined;

/**
 * @member {Number} minCpu
 */
ScaleThreshold.prototype['minCpu'] = undefined;

/**
 * @member {Number} maxCpu
 */
ScaleThreshold.prototype['maxCpu'] = undefined;

/**
 * @member {Boolean} memoryEnabled
 */
ScaleThreshold.prototype['memoryEnabled'] = undefined;

/**
 * @member {Number} minMemory
 */
ScaleThreshold.prototype['minMemory'] = undefined;

/**
 * @member {Number} maxMemory
 */
ScaleThreshold.prototype['maxMemory'] = undefined;

/**
 * @member {Boolean} diskEnabled
 */
ScaleThreshold.prototype['diskEnabled'] = undefined;

/**
 * @member {Number} minDisk
 */
ScaleThreshold.prototype['minDisk'] = undefined;

/**
 * @member {Number} maxDisk
 */
ScaleThreshold.prototype['maxDisk'] = undefined;

/**
 * @member {Date} dateCreated
 */
ScaleThreshold.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
ScaleThreshold.prototype['lastUpdated'] = undefined;






export default ScaleThreshold;

