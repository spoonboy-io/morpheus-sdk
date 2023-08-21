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
import Creds from './Creds';
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';
import InlineResponse20079LoadBalancerMonitorLoadBalancerType from './InlineResponse20079LoadBalancerMonitorLoadBalancerType';
import IntegrationGitRepoConfig from './IntegrationGitRepoConfig';

/**
 * The IntegrationGitHub model module.
 * @module model/IntegrationGitHub
 * @version 6.2.1
 */
class IntegrationGitHub {
    /**
     * Constructs a new <code>IntegrationGitHub</code>.
     * @alias module:model/IntegrationGitHub
     */
    constructor() { 
        
        IntegrationGitHub.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>IntegrationGitHub</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/IntegrationGitHub} obj Optional instance to populate.
     * @return {module:model/IntegrationGitHub} The populated <code>IntegrationGitHub</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new IntegrationGitHub();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('integrationType')) {
                obj['integrationType'] = InlineResponse20079LoadBalancerMonitorLoadBalancerType.constructFromObject(data['integrationType']);
            }
            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
            }
            if (data.hasOwnProperty('token')) {
                obj['token'] = ApiClient.convertToType(data['token'], 'String');
            }
            if (data.hasOwnProperty('tokenHash')) {
                obj['tokenHash'] = ApiClient.convertToType(data['tokenHash'], 'String');
            }
            if (data.hasOwnProperty('serviceKey')) {
                obj['serviceKey'] = InlineResponse20040AppDeployInstance.constructFromObject(data['serviceKey']);
            }
            if (data.hasOwnProperty('isPlugin')) {
                obj['isPlugin'] = ApiClient.convertToType(data['isPlugin'], 'Boolean');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = IntegrationGitRepoConfig.constructFromObject(data['config']);
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('statusDate')) {
                obj['statusDate'] = ApiClient.convertToType(data['statusDate'], 'Date');
            }
            if (data.hasOwnProperty('statusMessage')) {
                obj['statusMessage'] = ApiClient.convertToType(data['statusMessage'], 'String');
            }
            if (data.hasOwnProperty('lastSync')) {
                obj['lastSync'] = ApiClient.convertToType(data['lastSync'], 'String');
            }
            if (data.hasOwnProperty('lastSyncDuration')) {
                obj['lastSyncDuration'] = ApiClient.convertToType(data['lastSyncDuration'], 'String');
            }
            if (data.hasOwnProperty('credential')) {
                obj['credential'] = Creds.constructFromObject(data['credential']);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
IntegrationGitHub.prototype['id'] = undefined;

/**
 * @member {String} name
 */
IntegrationGitHub.prototype['name'] = undefined;

/**
 * @member {Boolean} enabled
 */
IntegrationGitHub.prototype['enabled'] = undefined;

/**
 * @member {module:model/IntegrationGitHub.TypeEnum} type
 */
IntegrationGitHub.prototype['type'] = undefined;

/**
 * @member {module:model/InlineResponse20079LoadBalancerMonitorLoadBalancerType} integrationType
 */
IntegrationGitHub.prototype['integrationType'] = undefined;

/**
 * @member {String} username
 */
IntegrationGitHub.prototype['username'] = undefined;

/**
 * @member {String} token
 */
IntegrationGitHub.prototype['token'] = undefined;

/**
 * @member {String} tokenHash
 */
IntegrationGitHub.prototype['tokenHash'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} serviceKey
 */
IntegrationGitHub.prototype['serviceKey'] = undefined;

/**
 * @member {Boolean} isPlugin
 */
IntegrationGitHub.prototype['isPlugin'] = undefined;

/**
 * @member {module:model/IntegrationGitRepoConfig} config
 */
IntegrationGitHub.prototype['config'] = undefined;

/**
 * @member {String} status
 */
IntegrationGitHub.prototype['status'] = undefined;

/**
 * @member {Date} statusDate
 */
IntegrationGitHub.prototype['statusDate'] = undefined;

/**
 * @member {String} statusMessage
 */
IntegrationGitHub.prototype['statusMessage'] = undefined;

/**
 * @member {String} lastSync
 */
IntegrationGitHub.prototype['lastSync'] = undefined;

/**
 * @member {String} lastSyncDuration
 */
IntegrationGitHub.prototype['lastSyncDuration'] = undefined;

/**
 * @member {module:model/Creds} credential
 */
IntegrationGitHub.prototype['credential'] = undefined;





/**
 * Allowed values for the <code>type</code> property.
 * @enum {String}
 * @readonly
 */
IntegrationGitHub['TypeEnum'] = {

    /**
     * value: "github"
     * @const
     */
    "github": "github"
};



export default IntegrationGitHub;
