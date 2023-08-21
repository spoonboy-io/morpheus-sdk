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
import InstanceNetwork from './InstanceNetwork';

/**
 * The InstanceInterfaces model module.
 * @module model/InstanceInterfaces
 * @version 6.2.1
 */
class InstanceInterfaces {
    /**
     * Constructs a new <code>InstanceInterfaces</code>.
     * @alias module:model/InstanceInterfaces
     */
    constructor() { 
        
        InstanceInterfaces.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InstanceInterfaces</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceInterfaces} obj Optional instance to populate.
     * @return {module:model/InstanceInterfaces} The populated <code>InstanceInterfaces</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceInterfaces();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('network')) {
                obj['network'] = InstanceNetwork.constructFromObject(data['network']);
            }
            if (data.hasOwnProperty('ipAddress')) {
                obj['ipAddress'] = ApiClient.convertToType(data['ipAddress'], 'String');
            }
            if (data.hasOwnProperty('networkInterfaceTypeId')) {
                obj['networkInterfaceTypeId'] = ApiClient.convertToType(data['networkInterfaceTypeId'], 'Number');
            }
            if (data.hasOwnProperty('ipMode')) {
                obj['ipMode'] = ApiClient.convertToType(data['ipMode'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} id
 */
InstanceInterfaces.prototype['id'] = undefined;

/**
 * @member {module:model/InstanceNetwork} network
 */
InstanceInterfaces.prototype['network'] = undefined;

/**
 * @member {String} ipAddress
 */
InstanceInterfaces.prototype['ipAddress'] = undefined;

/**
 * @member {Number} networkInterfaceTypeId
 */
InstanceInterfaces.prototype['networkInterfaceTypeId'] = undefined;

/**
 * @member {String} ipMode
 */
InstanceInterfaces.prototype['ipMode'] = undefined;






export default InstanceInterfaces;
