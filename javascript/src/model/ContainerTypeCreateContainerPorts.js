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
 * The ContainerTypeCreateContainerPorts model module.
 * @module model/ContainerTypeCreateContainerPorts
 * @version 6.2.1
 */
class ContainerTypeCreateContainerPorts {
    /**
     * Constructs a new <code>ContainerTypeCreateContainerPorts</code>.
     * @alias module:model/ContainerTypeCreateContainerPorts
     * @param name {String} 
     * @param port {Number} 
     */
    constructor(name, port) { 
        
        ContainerTypeCreateContainerPorts.initialize(this, name, port);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name, port) { 
        obj['name'] = name;
        obj['port'] = port;
    }

    /**
     * Constructs a <code>ContainerTypeCreateContainerPorts</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ContainerTypeCreateContainerPorts} obj Optional instance to populate.
     * @return {module:model/ContainerTypeCreateContainerPorts} The populated <code>ContainerTypeCreateContainerPorts</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ContainerTypeCreateContainerPorts();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('port')) {
                obj['port'] = ApiClient.convertToType(data['port'], 'Number');
            }
            if (data.hasOwnProperty('loadBalanceProtocol')) {
                obj['loadBalanceProtocol'] = ApiClient.convertToType(data['loadBalanceProtocol'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} name
 */
ContainerTypeCreateContainerPorts.prototype['name'] = undefined;

/**
 * @member {Number} port
 */
ContainerTypeCreateContainerPorts.prototype['port'] = undefined;

/**
 * @member {String} loadBalanceProtocol
 */
ContainerTypeCreateContainerPorts.prototype['loadBalanceProtocol'] = undefined;






export default ContainerTypeCreateContainerPorts;

