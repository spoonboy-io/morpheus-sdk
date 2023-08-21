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
import NetworkPoolServerCreateInfobloxConfig from './NetworkPoolServerCreateInfobloxConfig';

/**
 * The NetworkPoolServerUpdateInfoblox model module.
 * @module model/NetworkPoolServerUpdateInfoblox
 * @version 6.2.1
 */
class NetworkPoolServerUpdateInfoblox {
    /**
     * Constructs a new <code>NetworkPoolServerUpdateInfoblox</code>.
     * @alias module:model/NetworkPoolServerUpdateInfoblox
     */
    constructor() { 
        
        NetworkPoolServerUpdateInfoblox.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>NetworkPoolServerUpdateInfoblox</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NetworkPoolServerUpdateInfoblox} obj Optional instance to populate.
     * @return {module:model/NetworkPoolServerUpdateInfoblox} The populated <code>NetworkPoolServerUpdateInfoblox</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NetworkPoolServerUpdateInfoblox();

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
            if (data.hasOwnProperty('zoneFilter')) {
                obj['zoneFilter'] = ApiClient.convertToType(data['zoneFilter'], 'String');
            }
            if (data.hasOwnProperty('tenantMatch')) {
                obj['tenantMatch'] = ApiClient.convertToType(data['tenantMatch'], 'String');
            }
            if (data.hasOwnProperty('serviceMode')) {
                obj['serviceMode'] = ApiClient.convertToType(data['serviceMode'], 'String');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = NetworkPoolServerCreateInfobloxConfig.constructFromObject(data['config']);
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
NetworkPoolServerUpdateInfoblox.prototype['name'] = undefined;

/**
 * Can be used to enable / disable the network pool server.
 * @member {Boolean} enabled
 * @default true
 */
NetworkPoolServerUpdateInfoblox.prototype['enabled'] = true;

/**
 * URL
 * @member {String} serviceUrl
 */
NetworkPoolServerUpdateInfoblox.prototype['serviceUrl'] = undefined;

/**
 * Username
 * @member {String} serviceUsername
 */
NetworkPoolServerUpdateInfoblox.prototype['serviceUsername'] = undefined;

/**
 * Password
 * @member {String} servicePassword
 */
NetworkPoolServerUpdateInfoblox.prototype['servicePassword'] = undefined;

/**
 * Throttle Rate
 * @member {Number} serviceThrottleRate
 * @default 0
 */
NetworkPoolServerUpdateInfoblox.prototype['serviceThrottleRate'] = 0;

/**
 * Disable SSL SNI Verification
 * @member {Boolean} ignoreSsl
 */
NetworkPoolServerUpdateInfoblox.prototype['ignoreSsl'] = undefined;

/**
 * Network Filter
 * @member {String} networkFilter
 */
NetworkPoolServerUpdateInfoblox.prototype['networkFilter'] = undefined;

/**
 * Zone Filter
 * @member {String} zoneFilter
 */
NetworkPoolServerUpdateInfoblox.prototype['zoneFilter'] = undefined;

/**
 * Tenant Match
 * @member {String} tenantMatch
 */
NetworkPoolServerUpdateInfoblox.prototype['tenantMatch'] = undefined;

/**
 * IP Mode
 * @member {module:model/NetworkPoolServerUpdateInfoblox.ServiceModeEnum} serviceMode
 * @default 'static'
 */
NetworkPoolServerUpdateInfoblox.prototype['serviceMode'] = 'static';

/**
 * @member {module:model/NetworkPoolServerCreateInfobloxConfig} config
 */
NetworkPoolServerUpdateInfoblox.prototype['config'] = undefined;

/**
 * @member {module:model/NetworkPoolServerCreateBluecatCredential} credential
 */
NetworkPoolServerUpdateInfoblox.prototype['credential'] = undefined;





/**
 * Allowed values for the <code>serviceMode</code> property.
 * @enum {String}
 * @readonly
 */
NetworkPoolServerUpdateInfoblox['ServiceModeEnum'] = {

    /**
     * value: "static"
     * @const
     */
    "static": "static",

    /**
     * value: "dhcp"
     * @const
     */
    "dhcp": "dhcp"
};



export default NetworkPoolServerUpdateInfoblox;

