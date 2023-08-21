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
 * The LoadBalancerInstanceUpdate model module.
 * @module model/LoadBalancerInstanceUpdate
 * @version 6.2.1
 */
class LoadBalancerInstanceUpdate {
    /**
     * Constructs a new <code>LoadBalancerInstanceUpdate</code>.
     * @alias module:model/LoadBalancerInstanceUpdate
     */
    constructor() { 
        
        LoadBalancerInstanceUpdate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>LoadBalancerInstanceUpdate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/LoadBalancerInstanceUpdate} obj Optional instance to populate.
     * @return {module:model/LoadBalancerInstanceUpdate} The populated <code>LoadBalancerInstanceUpdate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new LoadBalancerInstanceUpdate();

            if (data.hasOwnProperty('vipName')) {
                obj['vipName'] = ApiClient.convertToType(data['vipName'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('vipAddress')) {
                obj['vipAddress'] = ApiClient.convertToType(data['vipAddress'], 'String');
            }
            if (data.hasOwnProperty('vipPort')) {
                obj['vipPort'] = ApiClient.convertToType(data['vipPort'], 'String');
            }
            if (data.hasOwnProperty('vipProtocol')) {
                obj['vipProtocol'] = ApiClient.convertToType(data['vipProtocol'], 'String');
            }
            if (data.hasOwnProperty('vipHostname')) {
                obj['vipHostname'] = ApiClient.convertToType(data['vipHostname'], 'String');
            }
            if (data.hasOwnProperty('pool')) {
                obj['pool'] = ApiClient.convertToType(data['pool'], 'Number');
            }
            if (data.hasOwnProperty('sslCert')) {
                obj['sslCert'] = ApiClient.convertToType(data['sslCert'], 'Number');
            }
            if (data.hasOwnProperty('sslServerCert')) {
                obj['sslServerCert'] = ApiClient.convertToType(data['sslServerCert'], 'Number');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], Object);
            }
        }
        return obj;
    }


}

/**
 * VIP Name
 * @member {String} vipName
 */
LoadBalancerInstanceUpdate.prototype['vipName'] = undefined;

/**
 * Description
 * @member {String} description
 */
LoadBalancerInstanceUpdate.prototype['description'] = undefined;

/**
 * VIP Address
 * @member {String} vipAddress
 */
LoadBalancerInstanceUpdate.prototype['vipAddress'] = undefined;

/**
 * VIP Port
 * @member {String} vipPort
 */
LoadBalancerInstanceUpdate.prototype['vipPort'] = undefined;

/**
 * VIP Protocol
 * @member {String} vipProtocol
 */
LoadBalancerInstanceUpdate.prototype['vipProtocol'] = undefined;

/**
 * VIP Hostname
 * @member {String} vipHostname
 */
LoadBalancerInstanceUpdate.prototype['vipHostname'] = undefined;

/**
 * @member {Number} pool
 */
LoadBalancerInstanceUpdate.prototype['pool'] = undefined;

/**
 * SSL Client Certificate ID
 * @member {Number} sslCert
 */
LoadBalancerInstanceUpdate.prototype['sslCert'] = undefined;

/**
 * SSL Server Certificate ID
 * @member {Number} sslServerCert
 */
LoadBalancerInstanceUpdate.prototype['sslServerCert'] = undefined;

/**
 * Configuration object with parameters that vary by type.
 * @member {Object} config
 */
LoadBalancerInstanceUpdate.prototype['config'] = undefined;






export default LoadBalancerInstanceUpdate;

