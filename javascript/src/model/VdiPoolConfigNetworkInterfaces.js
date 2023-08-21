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
import VdiPoolConfigNetwork from './VdiPoolConfigNetwork';

/**
 * The VdiPoolConfigNetworkInterfaces model module.
 * @module model/VdiPoolConfigNetworkInterfaces
 * @version 6.2.1
 */
class VdiPoolConfigNetworkInterfaces {
    /**
     * Constructs a new <code>VdiPoolConfigNetworkInterfaces</code>.
     * @alias module:model/VdiPoolConfigNetworkInterfaces
     */
    constructor() { 
        
        VdiPoolConfigNetworkInterfaces.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>VdiPoolConfigNetworkInterfaces</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/VdiPoolConfigNetworkInterfaces} obj Optional instance to populate.
     * @return {module:model/VdiPoolConfigNetworkInterfaces} The populated <code>VdiPoolConfigNetworkInterfaces</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new VdiPoolConfigNetworkInterfaces();

            if (data.hasOwnProperty('primaryInterface')) {
                obj['primaryInterface'] = ApiClient.convertToType(data['primaryInterface'], 'Boolean');
            }
            if (data.hasOwnProperty('network')) {
                obj['network'] = VdiPoolConfigNetwork.constructFromObject(data['network']);
            }
            if (data.hasOwnProperty('ipMode')) {
                obj['ipMode'] = ApiClient.convertToType(data['ipMode'], 'String');
            }
            if (data.hasOwnProperty('showNetworkPoolLabel')) {
                obj['showNetworkPoolLabel'] = ApiClient.convertToType(data['showNetworkPoolLabel'], 'Boolean');
            }
            if (data.hasOwnProperty('showNetworkDhcpLabel')) {
                obj['showNetworkDhcpLabel'] = ApiClient.convertToType(data['showNetworkDhcpLabel'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {Boolean} primaryInterface
 */
VdiPoolConfigNetworkInterfaces.prototype['primaryInterface'] = undefined;

/**
 * @member {module:model/VdiPoolConfigNetwork} network
 */
VdiPoolConfigNetworkInterfaces.prototype['network'] = undefined;

/**
 * @member {String} ipMode
 */
VdiPoolConfigNetworkInterfaces.prototype['ipMode'] = undefined;

/**
 * @member {Boolean} showNetworkPoolLabel
 */
VdiPoolConfigNetworkInterfaces.prototype['showNetworkPoolLabel'] = undefined;

/**
 * @member {Boolean} showNetworkDhcpLabel
 */
VdiPoolConfigNetworkInterfaces.prototype['showNetworkDhcpLabel'] = undefined;






export default VdiPoolConfigNetworkInterfaces;

