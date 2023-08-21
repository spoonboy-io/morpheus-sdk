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
import NetworkPoolServerCreateBluecatCredential from './NetworkPoolServerCreateBluecatCredential';
import NetworkPoolServerUpdatePhpIpamConfig from './NetworkPoolServerUpdatePhpIpamConfig';

/**
 * The NetworkPoolServerUpdatePhpIpam model module.
 * @module model/NetworkPoolServerUpdatePhpIpam
 * @version 6.2.1
 */
class NetworkPoolServerUpdatePhpIpam {
    /**
     * Constructs a new <code>NetworkPoolServerUpdatePhpIpam</code>.
     * @alias module:model/NetworkPoolServerUpdatePhpIpam
     */
    constructor() { 
        
        NetworkPoolServerUpdatePhpIpam.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>NetworkPoolServerUpdatePhpIpam</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NetworkPoolServerUpdatePhpIpam} obj Optional instance to populate.
     * @return {module:model/NetworkPoolServerUpdatePhpIpam} The populated <code>NetworkPoolServerUpdatePhpIpam</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NetworkPoolServerUpdatePhpIpam();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('serviceUrl')) {
                obj['serviceUrl'] = ApiClient.convertToType(data['serviceUrl'], 'String');
            }
            if (data.hasOwnProperty('serviceUsername')) {
                obj['serviceUsername'] = ApiClient.convertToType(data['serviceUsername'], 'String');
            }
            if (data.hasOwnProperty('servicePassword')) {
                obj['servicePassword'] = ApiClient.convertToType(data['servicePassword'], 'String');
            }
            if (data.hasOwnProperty('serviceThrottleRate')) {
                obj['serviceThrottleRate'] = ApiClient.convertToType(data['serviceThrottleRate'], 'Number');
            }
            if (data.hasOwnProperty('ignoreSsl')) {
                obj['ignoreSsl'] = ApiClient.convertToType(data['ignoreSsl'], 'Boolean');
            }
            if (data.hasOwnProperty('networkFilter')) {
                obj['networkFilter'] = ApiClient.convertToType(data['networkFilter'], 'String');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = NetworkPoolServerUpdatePhpIpamConfig.constructFromObject(data['config']);
            }
            if (data.hasOwnProperty('credential')) {
                obj['credential'] = NetworkPoolServerCreateBluecatCredential.constructFromObject(data['credential']);
            }
        }
        return obj;
    }


}

/**
 * Name
 * @member {String} name
 */
NetworkPoolServerUpdatePhpIpam.prototype['name'] = undefined;

/**
 * Can be used to enable / disable the network pool server.
 * @member {Boolean} enabled
 * @default true
 */
NetworkPoolServerUpdatePhpIpam.prototype['enabled'] = true;

/**
 * URL
 * @member {String} serviceUrl
 */
NetworkPoolServerUpdatePhpIpam.prototype['serviceUrl'] = undefined;

/**
 * Username
 * @member {String} serviceUsername
 */
NetworkPoolServerUpdatePhpIpam.prototype['serviceUsername'] = undefined;

/**
 * Password
 * @member {String} servicePassword
 */
NetworkPoolServerUpdatePhpIpam.prototype['servicePassword'] = undefined;

/**
 * Throttle Rate
 * @member {Number} serviceThrottleRate
 * @default 0
 */
NetworkPoolServerUpdatePhpIpam.prototype['serviceThrottleRate'] = 0;

/**
 * Disable SSL SNI Verification
 * @member {Boolean} ignoreSsl
 */
NetworkPoolServerUpdatePhpIpam.prototype['ignoreSsl'] = undefined;

/**
 * Network Filter
 * @member {String} networkFilter
 */
NetworkPoolServerUpdatePhpIpam.prototype['networkFilter'] = undefined;

/**
 * @member {module:model/NetworkPoolServerUpdatePhpIpamConfig} config
 */
NetworkPoolServerUpdatePhpIpam.prototype['config'] = undefined;

/**
 * @member {module:model/NetworkPoolServerCreateBluecatCredential} credential
 */
NetworkPoolServerUpdatePhpIpam.prototype['credential'] = undefined;






export default NetworkPoolServerUpdatePhpIpam;
