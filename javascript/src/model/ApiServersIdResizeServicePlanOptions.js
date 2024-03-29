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
 * The ApiServersIdResizeServicePlanOptions model module.
 * @module model/ApiServersIdResizeServicePlanOptions
 * @version 6.2.1
 */
class ApiServersIdResizeServicePlanOptions {
    /**
     * Constructs a new <code>ApiServersIdResizeServicePlanOptions</code>.
     * @alias module:model/ApiServersIdResizeServicePlanOptions
     */
    constructor() { 
        
        ApiServersIdResizeServicePlanOptions.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiServersIdResizeServicePlanOptions</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiServersIdResizeServicePlanOptions} obj Optional instance to populate.
     * @return {module:model/ApiServersIdResizeServicePlanOptions} The populated <code>ApiServersIdResizeServicePlanOptions</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiServersIdResizeServicePlanOptions();

            if (data.hasOwnProperty('maxCores')) {
                obj['maxCores'] = ApiClient.convertToType(data['maxCores'], 'Number');
            }
            if (data.hasOwnProperty('coresPerSocket')) {
                obj['coresPerSocket'] = ApiClient.convertToType(data['coresPerSocket'], 'Number');
            }
            if (data.hasOwnProperty('maxMemory')) {
                obj['maxMemory'] = ApiClient.convertToType(data['maxMemory'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * Core Count
 * @member {Number} maxCores
 */
ApiServersIdResizeServicePlanOptions.prototype['maxCores'] = undefined;

/**
 * Cores Per Socket
 * @member {Number} coresPerSocket
 */
ApiServersIdResizeServicePlanOptions.prototype['coresPerSocket'] = undefined;

/**
 * Memory in bytes For backwards compatability, values less than 1048576 are treated as being in MB and will be converted to bytes
 * @member {Number} maxMemory
 */
ApiServersIdResizeServicePlanOptions.prototype['maxMemory'] = undefined;






export default ApiServersIdResizeServicePlanOptions;

