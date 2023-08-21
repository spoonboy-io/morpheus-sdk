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
import IntegrationGitHubConfigIntegrationConfig from './IntegrationGitHubConfigIntegrationConfig';

/**
 * The IntegrationGitHubConfigIntegration model module.
 * @module model/IntegrationGitHubConfigIntegration
 * @version 6.2.1
 */
class IntegrationGitHubConfigIntegration {
    /**
     * Constructs a new <code>IntegrationGitHubConfigIntegration</code>.
     * @alias module:model/IntegrationGitHubConfigIntegration
     * @param name {String} Name, a unique identifier for the integration
     * @param type {module:model/IntegrationGitHubConfigIntegration.TypeEnum} Integration Type Code
     * @param serviceUsername {String} Username
     */
    constructor(name, type, serviceUsername) { 
        
        IntegrationGitHubConfigIntegration.initialize(this, name, type, serviceUsername);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name, type, serviceUsername) { 
        obj['name'] = name;
        obj['type'] = type;
        obj['serviceUsername'] = serviceUsername;
    }

    /**
     * Constructs a <code>IntegrationGitHubConfigIntegration</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/IntegrationGitHubConfigIntegration} obj Optional instance to populate.
     * @return {module:model/IntegrationGitHubConfigIntegration} The populated <code>IntegrationGitHubConfigIntegration</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new IntegrationGitHubConfigIntegration();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('serviceUsername')) {
                obj['serviceUsername'] = ApiClient.convertToType(data['serviceUsername'], 'String');
            }
            if (data.hasOwnProperty('servicePassword')) {
                obj['servicePassword'] = ApiClient.convertToType(data['servicePassword'], 'String');
            }
            if (data.hasOwnProperty('serviceToken')) {
                obj['serviceToken'] = ApiClient.convertToType(data['serviceToken'], 'String');
            }
            if (data.hasOwnProperty('serviceKey')) {
                obj['serviceKey'] = ApiClient.convertToType(data['serviceKey'], 'Number');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = IntegrationGitHubConfigIntegrationConfig.constructFromObject(data['config']);
            }
        }
        return obj;
    }


}

/**
 * Name, a unique identifier for the integration
 * @member {String} name
 */
IntegrationGitHubConfigIntegration.prototype['name'] = undefined;

/**
 * Integration Type Code
 * @member {module:model/IntegrationGitHubConfigIntegration.TypeEnum} type
 */
IntegrationGitHubConfigIntegration.prototype['type'] = undefined;

/**
 * Username
 * @member {String} serviceUsername
 */
IntegrationGitHubConfigIntegration.prototype['serviceUsername'] = undefined;

/**
 * Password
 * @member {String} servicePassword
 */
IntegrationGitHubConfigIntegration.prototype['servicePassword'] = undefined;

/**
 * Access Token
 * @member {String} serviceToken
 */
IntegrationGitHubConfigIntegration.prototype['serviceToken'] = undefined;

/**
 * Key Pair ID
 * @member {Number} serviceKey
 */
IntegrationGitHubConfigIntegration.prototype['serviceKey'] = undefined;

/**
 * @member {module:model/IntegrationGitHubConfigIntegrationConfig} config
 */
IntegrationGitHubConfigIntegration.prototype['config'] = undefined;





/**
 * Allowed values for the <code>type</code> property.
 * @enum {String}
 * @readonly
 */
IntegrationGitHubConfigIntegration['TypeEnum'] = {

    /**
     * value: "github"
     * @const
     */
    "github": "github"
};



export default IntegrationGitHubConfigIntegration;
