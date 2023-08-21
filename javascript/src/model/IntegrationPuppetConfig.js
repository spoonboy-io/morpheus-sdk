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
 * The IntegrationPuppetConfig model module.
 * @module model/IntegrationPuppetConfig
 * @version 6.2.1
 */
class IntegrationPuppetConfig {
    /**
     * Constructs a new <code>IntegrationPuppetConfig</code>.
     * @alias module:model/IntegrationPuppetConfig
     */
    constructor() { 
        
        IntegrationPuppetConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>IntegrationPuppetConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/IntegrationPuppetConfig} obj Optional instance to populate.
     * @return {module:model/IntegrationPuppetConfig} The populated <code>IntegrationPuppetConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new IntegrationPuppetConfig();

            if (data.hasOwnProperty('puppetMaster')) {
                obj['puppetMaster'] = ApiClient.convertToType(data['puppetMaster'], 'String');
            }
            if (data.hasOwnProperty('puppetFireNow')) {
                obj['puppetFireNow'] = ApiClient.convertToType(data['puppetFireNow'], 'String');
            }
            if (data.hasOwnProperty('puppetSshUser')) {
                obj['puppetSshUser'] = ApiClient.convertToType(data['puppetSshUser'], 'String');
            }
            if (data.hasOwnProperty('puppetSshPassword')) {
                obj['puppetSshPassword'] = ApiClient.convertToType(data['puppetSshPassword'], 'String');
            }
            if (data.hasOwnProperty('puppetSshPasswordHash')) {
                obj['puppetSshPasswordHash'] = ApiClient.convertToType(data['puppetSshPasswordHash'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} puppetMaster
 */
IntegrationPuppetConfig.prototype['puppetMaster'] = undefined;

/**
 * @member {String} puppetFireNow
 */
IntegrationPuppetConfig.prototype['puppetFireNow'] = undefined;

/**
 * @member {String} puppetSshUser
 */
IntegrationPuppetConfig.prototype['puppetSshUser'] = undefined;

/**
 * @member {String} puppetSshPassword
 */
IntegrationPuppetConfig.prototype['puppetSshPassword'] = undefined;

/**
 * @member {String} puppetSshPasswordHash
 */
IntegrationPuppetConfig.prototype['puppetSshPasswordHash'] = undefined;






export default IntegrationPuppetConfig;

