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
 * The InstanceCreatePorts model module.
 * @module model/InstanceCreatePorts
 * @version 6.2.1
 */
class InstanceCreatePorts {
    /**
     * Constructs a new <code>InstanceCreatePorts</code>.
     * @alias module:model/InstanceCreatePorts
     * @param port {Number} Port number.
     */
    constructor(port) { 
        
        InstanceCreatePorts.initialize(this, port);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, port) { 
        obj['port'] = port;
    }

    /**
     * Constructs a <code>InstanceCreatePorts</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceCreatePorts} obj Optional instance to populate.
     * @return {module:model/InstanceCreatePorts} The populated <code>InstanceCreatePorts</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceCreatePorts();

            if (data.hasOwnProperty('port')) {
                obj['port'] = ApiClient.convertToType(data['port'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('lb')) {
                obj['lb'] = ApiClient.convertToType(data['lb'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Port number.
 * @member {Number} port
 */
InstanceCreatePorts.prototype['port'] = undefined;

/**
 * A name for the port.
 * @member {String} name
 */
InstanceCreatePorts.prototype['name'] = undefined;

/**
 * The load balancer protocol. HTTP, HTTPS, or TCP.
 * @member {String} lb
 */
InstanceCreatePorts.prototype['lb'] = undefined;






export default InstanceCreatePorts;

