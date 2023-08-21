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
import NetworkDhcpServerCreateConfig from './NetworkDhcpServerCreateConfig';

/**
 * The NetworkDhcpServerCreate model module.
 * @module model/NetworkDhcpServerCreate
 * @version 6.2.1
 */
class NetworkDhcpServerCreate {
    /**
     * Constructs a new <code>NetworkDhcpServerCreate</code>.
     * @alias module:model/NetworkDhcpServerCreate
     * @param serverIpAddress {String} Server Address for the DHCP Server
     * @param leaseTime {Number} Optional lease time for the DHCP Server
     * @param name {String} Name
     * @param config {module:model/NetworkDhcpServerCreateConfig} 
     */
    constructor(serverIpAddress, leaseTime, name, config) { 
        
        NetworkDhcpServerCreate.initialize(this, serverIpAddress, leaseTime, name, config);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, serverIpAddress, leaseTime, name, config) { 
        obj['serverIpAddress'] = serverIpAddress;
        obj['leaseTime'] = leaseTime || 86400;
        obj['name'] = name;
        obj['config'] = config;
    }

    /**
     * Constructs a <code>NetworkDhcpServerCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NetworkDhcpServerCreate} obj Optional instance to populate.
     * @return {module:model/NetworkDhcpServerCreate} The populated <code>NetworkDhcpServerCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NetworkDhcpServerCreate();

            if (data.hasOwnProperty('serverIpAddress')) {
                obj['serverIpAddress'] = ApiClient.convertToType(data['serverIpAddress'], 'String');
            }
            if (data.hasOwnProperty('leaseTime')) {
                obj['leaseTime'] = ApiClient.convertToType(data['leaseTime'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = NetworkDhcpServerCreateConfig.constructFromObject(data['config']);
            }
        }
        return obj;
    }


}

/**
 * Server Address for the DHCP Server
 * @member {String} serverIpAddress
 */
NetworkDhcpServerCreate.prototype['serverIpAddress'] = undefined;

/**
 * Optional lease time for the DHCP Server
 * @member {Number} leaseTime
 * @default 86400
 */
NetworkDhcpServerCreate.prototype['leaseTime'] = 86400;

/**
 * Name
 * @member {String} name
 */
NetworkDhcpServerCreate.prototype['name'] = undefined;

/**
 * @member {module:model/NetworkDhcpServerCreateConfig} config
 */
NetworkDhcpServerCreate.prototype['config'] = undefined;






export default NetworkDhcpServerCreate;

