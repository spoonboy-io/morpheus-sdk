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
import IntegrationChefConfigDatabags from './IntegrationChefConfigDatabags';

/**
 * The IntegrationChefConfig model module.
 * @module model/IntegrationChefConfig
 * @version 6.2.1
 */
class IntegrationChefConfig {
    /**
     * Constructs a new <code>IntegrationChefConfig</code>.
     * @alias module:model/IntegrationChefConfig
     */
    constructor() { 
        
        IntegrationChefConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>IntegrationChefConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/IntegrationChefConfig} obj Optional instance to populate.
     * @return {module:model/IntegrationChefConfig} The populated <code>IntegrationChefConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new IntegrationChefConfig();

            if (data.hasOwnProperty('databags')) {
                obj['databags'] = ApiClient.convertToType(data['databags'], [IntegrationChefConfigDatabags]);
            }
            if (data.hasOwnProperty('endpoint')) {
                obj['endpoint'] = ApiClient.convertToType(data['endpoint'], 'String');
            }
            if (data.hasOwnProperty('org')) {
                obj['org'] = ApiClient.convertToType(data['org'], 'String');
            }
            if (data.hasOwnProperty('chefUser')) {
                obj['chefUser'] = ApiClient.convertToType(data['chefUser'], 'String');
            }
            if (data.hasOwnProperty('userKey')) {
                obj['userKey'] = ApiClient.convertToType(data['userKey'], 'String');
            }
            if (data.hasOwnProperty('orgKey')) {
                obj['orgKey'] = ApiClient.convertToType(data['orgKey'], 'String');
            }
            if (data.hasOwnProperty('version')) {
                obj['version'] = ApiClient.convertToType(data['version'], 'String');
            }
            if (data.hasOwnProperty('chefUseFqdn')) {
                obj['chefUseFqdn'] = ApiClient.convertToType(data['chefUseFqdn'], 'Boolean');
            }
            if (data.hasOwnProperty('windowsVersion')) {
                obj['windowsVersion'] = ApiClient.convertToType(data['windowsVersion'], 'String');
            }
            if (data.hasOwnProperty('windowsInstallUrl')) {
                obj['windowsInstallUrl'] = ApiClient.convertToType(data['windowsInstallUrl'], 'String');
            }
            if (data.hasOwnProperty('userKeyHash')) {
                obj['userKeyHash'] = ApiClient.convertToType(data['userKeyHash'], 'String');
            }
            if (data.hasOwnProperty('orgKeyHash')) {
                obj['orgKeyHash'] = ApiClient.convertToType(data['orgKeyHash'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/IntegrationChefConfigDatabags>} databags
 */
IntegrationChefConfig.prototype['databags'] = undefined;

/**
 * @member {String} endpoint
 */
IntegrationChefConfig.prototype['endpoint'] = undefined;

/**
 * @member {String} org
 */
IntegrationChefConfig.prototype['org'] = undefined;

/**
 * @member {String} chefUser
 */
IntegrationChefConfig.prototype['chefUser'] = undefined;

/**
 * @member {String} userKey
 */
IntegrationChefConfig.prototype['userKey'] = undefined;

/**
 * @member {String} orgKey
 */
IntegrationChefConfig.prototype['orgKey'] = undefined;

/**
 * @member {String} version
 */
IntegrationChefConfig.prototype['version'] = undefined;

/**
 * @member {Boolean} chefUseFqdn
 */
IntegrationChefConfig.prototype['chefUseFqdn'] = undefined;

/**
 * @member {String} windowsVersion
 */
IntegrationChefConfig.prototype['windowsVersion'] = undefined;

/**
 * @member {String} windowsInstallUrl
 */
IntegrationChefConfig.prototype['windowsInstallUrl'] = undefined;

/**
 * @member {String} userKeyHash
 */
IntegrationChefConfig.prototype['userKeyHash'] = undefined;

/**
 * @member {String} orgKeyHash
 */
IntegrationChefConfig.prototype['orgKeyHash'] = undefined;






export default IntegrationChefConfig;

