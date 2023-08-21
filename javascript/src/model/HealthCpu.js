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
 * The HealthCpu model module.
 * @module model/HealthCpu
 * @version 6.2.1
 */
class HealthCpu {
    /**
     * Constructs a new <code>HealthCpu</code>.
     * @alias module:model/HealthCpu
     */
    constructor() { 
        
        HealthCpu.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>HealthCpu</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/HealthCpu} obj Optional instance to populate.
     * @return {module:model/HealthCpu} The populated <code>HealthCpu</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new HealthCpu();

            if (data.hasOwnProperty('success')) {
                obj['success'] = ApiClient.convertToType(data['success'], 'Boolean');
            }
            if (data.hasOwnProperty('cpuLoad')) {
                obj['cpuLoad'] = ApiClient.convertToType(data['cpuLoad'], 'Number');
            }
            if (data.hasOwnProperty('cpuTotalLoad')) {
                obj['cpuTotalLoad'] = ApiClient.convertToType(data['cpuTotalLoad'], 'Number');
            }
            if (data.hasOwnProperty('processorCount')) {
                obj['processorCount'] = ApiClient.convertToType(data['processorCount'], 'Number');
            }
            if (data.hasOwnProperty('processTime')) {
                obj['processTime'] = ApiClient.convertToType(data['processTime'], 'Number');
            }
            if (data.hasOwnProperty('systemLoad')) {
                obj['systemLoad'] = ApiClient.convertToType(data['systemLoad'], 'Number');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Boolean} success
 */
HealthCpu.prototype['success'] = undefined;

/**
 * @member {Number} cpuLoad
 */
HealthCpu.prototype['cpuLoad'] = undefined;

/**
 * @member {Number} cpuTotalLoad
 */
HealthCpu.prototype['cpuTotalLoad'] = undefined;

/**
 * @member {Number} processorCount
 */
HealthCpu.prototype['processorCount'] = undefined;

/**
 * @member {Number} processTime
 */
HealthCpu.prototype['processTime'] = undefined;

/**
 * @member {Number} systemLoad
 */
HealthCpu.prototype['systemLoad'] = undefined;

/**
 * @member {String} status
 */
HealthCpu.prototype['status'] = undefined;






export default HealthCpu;
