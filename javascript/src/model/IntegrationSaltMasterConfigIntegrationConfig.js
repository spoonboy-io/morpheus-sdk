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
 * The IntegrationSaltMasterConfigIntegrationConfig model module.
 * @module model/IntegrationSaltMasterConfigIntegrationConfig
 * @version 6.2.1
 */
class IntegrationSaltMasterConfigIntegrationConfig {
    /**
     * Constructs a new <code>IntegrationSaltMasterConfigIntegrationConfig</code>.
     * @alias module:model/IntegrationSaltMasterConfigIntegrationConfig
     */
    constructor() { 
        
        IntegrationSaltMasterConfigIntegrationConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>IntegrationSaltMasterConfigIntegrationConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/IntegrationSaltMasterConfigIntegrationConfig} obj Optional instance to populate.
     * @return {module:model/IntegrationSaltMasterConfigIntegrationConfig} The populated <code>IntegrationSaltMasterConfigIntegrationConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new IntegrationSaltMasterConfigIntegrationConfig();

            if (data.hasOwnProperty('saltApplyOnMinion')) {
                obj['saltApplyOnMinion'] = ApiClient.convertToType(data['saltApplyOnMinion'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * Apply state via Minion instead of Master (salt-call)
 * @member {Boolean} saltApplyOnMinion
 */
IntegrationSaltMasterConfigIntegrationConfig.prototype['saltApplyOnMinion'] = undefined;






export default IntegrationSaltMasterConfigIntegrationConfig;

