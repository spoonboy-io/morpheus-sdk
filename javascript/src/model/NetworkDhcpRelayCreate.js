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
 * The NetworkDhcpRelayCreate model module.
 * @module model/NetworkDhcpRelayCreate
 * @version 6.2.1
 */
class NetworkDhcpRelayCreate {
    /**
     * Constructs a new <code>NetworkDhcpRelayCreate</code>.
     * @alias module:model/NetworkDhcpRelayCreate
     */
    constructor() { 
        
        NetworkDhcpRelayCreate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>NetworkDhcpRelayCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NetworkDhcpRelayCreate} obj Optional instance to populate.
     * @return {module:model/NetworkDhcpRelayCreate} The populated <code>NetworkDhcpRelayCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NetworkDhcpRelayCreate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('serverIpAddresses')) {
                obj['serverIpAddresses'] = ApiClient.convertToType(data['serverIpAddresses'], ['String']);
            }
        }
        return obj;
    }


}

/**
 * @member {String} name
 */
NetworkDhcpRelayCreate.prototype['name'] = undefined;

/**
 * @member {Array.<String>} serverIpAddresses
 */
NetworkDhcpRelayCreate.prototype['serverIpAddresses'] = undefined;






export default NetworkDhcpRelayCreate;

