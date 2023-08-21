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
 * The IntegrationGitRepoConfigIntegrationConfig model module.
 * @module model/IntegrationGitRepoConfigIntegrationConfig
 * @version 6.2.1
 */
class IntegrationGitRepoConfigIntegrationConfig {
    /**
     * Constructs a new <code>IntegrationGitRepoConfigIntegrationConfig</code>.
     * @alias module:model/IntegrationGitRepoConfigIntegrationConfig
     */
    constructor() { 
        
        IntegrationGitRepoConfigIntegrationConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>IntegrationGitRepoConfigIntegrationConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/IntegrationGitRepoConfigIntegrationConfig} obj Optional instance to populate.
     * @return {module:model/IntegrationGitRepoConfigIntegrationConfig} The populated <code>IntegrationGitRepoConfigIntegrationConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new IntegrationGitRepoConfigIntegrationConfig();

            if (data.hasOwnProperty('defaultBranch')) {
                obj['defaultBranch'] = ApiClient.convertToType(data['defaultBranch'], 'String');
            }
            if (data.hasOwnProperty('cacheEnabled')) {
                obj['cacheEnabled'] = ApiClient.convertToType(data['cacheEnabled'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * Default Branch
 * @member {module:model/IntegrationGitRepoConfigIntegrationConfig.DefaultBranchEnum} defaultBranch
 */
IntegrationGitRepoConfigIntegrationConfig.prototype['defaultBranch'] = undefined;

/**
 * Enable Git Repository Caching
 * @member {Boolean} cacheEnabled
 */
IntegrationGitRepoConfigIntegrationConfig.prototype['cacheEnabled'] = undefined;





/**
 * Allowed values for the <code>defaultBranch</code> property.
 * @enum {String}
 * @readonly
 */
IntegrationGitRepoConfigIntegrationConfig['DefaultBranchEnum'] = {

    /**
     * value: "main"
     * @const
     */
    "main": "main",

    /**
     * value: "master"
     * @const
     */
    "master": "master"
};



export default IntegrationGitRepoConfigIntegrationConfig;

