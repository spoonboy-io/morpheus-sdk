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
 * The GroupCreateConfig model module.
 * @module model/GroupCreateConfig
 * @version 6.2.1
 */
class GroupCreateConfig {
    /**
     * Constructs a new <code>GroupCreateConfig</code>.
     * @alias module:model/GroupCreateConfig
     */
    constructor() { 
        
        GroupCreateConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>GroupCreateConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/GroupCreateConfig} obj Optional instance to populate.
     * @return {module:model/GroupCreateConfig} The populated <code>GroupCreateConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new GroupCreateConfig();

            if (data.hasOwnProperty('dnsIntegrationId')) {
                obj['dnsIntegrationId'] = ApiClient.convertToType(data['dnsIntegrationId'], 'String');
            }
            if (data.hasOwnProperty('configCmdbId')) {
                obj['configCmdbId'] = ApiClient.convertToType(data['configCmdbId'], 'String');
            }
            if (data.hasOwnProperty('configCmId')) {
                obj['configCmId'] = ApiClient.convertToType(data['configCmId'], 'String');
            }
            if (data.hasOwnProperty('serviceRegistryId')) {
                obj['serviceRegistryId'] = ApiClient.convertToType(data['serviceRegistryId'], 'String');
            }
            if (data.hasOwnProperty('configManagementId')) {
                obj['configManagementId'] = ApiClient.convertToType(data['configManagementId'], 'String');
            }
            if (data.hasOwnProperty('configCmdbDiscovery')) {
                obj['configCmdbDiscovery'] = ApiClient.convertToType(data['configCmdbDiscovery'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * Optional DNS Integration ID
 * @member {String} dnsIntegrationId
 */
GroupCreateConfig.prototype['dnsIntegrationId'] = undefined;

/**
 * Optional CMDB Integration ID
 * @member {String} configCmdbId
 */
GroupCreateConfig.prototype['configCmdbId'] = undefined;

/**
 * Optional Change Management Integration ID
 * @member {String} configCmId
 */
GroupCreateConfig.prototype['configCmId'] = undefined;

/**
 * Optional Service Registry Integration ID
 * @member {String} serviceRegistryId
 */
GroupCreateConfig.prototype['serviceRegistryId'] = undefined;

/**
 * Optional Configuration Management Integration ID
 * @member {String} configManagementId
 */
GroupCreateConfig.prototype['configManagementId'] = undefined;

/**
 * Enable or disable CMDB Discovery
 * @member {Boolean} configCmdbDiscovery
 */
GroupCreateConfig.prototype['configCmdbDiscovery'] = undefined;






export default GroupCreateConfig;
