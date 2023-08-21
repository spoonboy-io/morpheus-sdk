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
import InlineResponse20079LoadBalancerMonitorLoadBalancerType from './InlineResponse20079LoadBalancerMonitorLoadBalancerType';
import IntegrationSNOWConfig from './IntegrationSNOWConfig';

/**
 * The IntegrationSNOW model module.
 * @module model/IntegrationSNOW
 * @version 6.2.1
 */
class IntegrationSNOW {
    /**
     * Constructs a new <code>IntegrationSNOW</code>.
     * @alias module:model/IntegrationSNOW
     */
    constructor() { 
        
        IntegrationSNOW.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>IntegrationSNOW</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/IntegrationSNOW} obj Optional instance to populate.
     * @return {module:model/IntegrationSNOW} The populated <code>IntegrationSNOW</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new IntegrationSNOW();

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
            if (data.hasOwnProperty('url')) {
                obj['url'] = ApiClient.convertToType(data['url'], 'String');
            }
            if (data.hasOwnProperty('isPlugin')) {
                obj['isPlugin'] = ApiClient.convertToType(data['isPlugin'], 'Boolean');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = IntegrationSNOWConfig.constructFromObject(data['config']);
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
IntegrationSNOW.prototype['id'] = undefined;

/**
 * @member {String} name
 */
IntegrationSNOW.prototype['name'] = undefined;

/**
 * @member {Boolean} enabled
 */
IntegrationSNOW.prototype['enabled'] = undefined;

/**
 * @member {module:model/IntegrationSNOW.TypeEnum} type
 */
IntegrationSNOW.prototype['type'] = undefined;

/**
 * @member {module:model/InlineResponse20079LoadBalancerMonitorLoadBalancerType} integrationType
 */
IntegrationSNOW.prototype['integrationType'] = undefined;

/**
 * @member {String} url
 */
IntegrationSNOW.prototype['url'] = undefined;

/**
 * @member {Boolean} isPlugin
 */
IntegrationSNOW.prototype['isPlugin'] = undefined;

/**
 * @member {module:model/IntegrationSNOWConfig} config
 */
IntegrationSNOW.prototype['config'] = undefined;

/**
 * @member {String} status
 */
IntegrationSNOW.prototype['status'] = undefined;

/**
 * @member {Date} statusDate
 */
IntegrationSNOW.prototype['statusDate'] = undefined;

/**
 * @member {String} statusMessage
 */
IntegrationSNOW.prototype['statusMessage'] = undefined;

/**
 * @member {String} lastSync
 */
IntegrationSNOW.prototype['lastSync'] = undefined;

/**
 * @member {String} lastSyncDuration
 */
IntegrationSNOW.prototype['lastSyncDuration'] = undefined;

/**
 * @member {module:model/Creds} credential
 */
IntegrationSNOW.prototype['credential'] = undefined;





/**
 * Allowed values for the <code>type</code> property.
 * @enum {String}
 * @readonly
 */
IntegrationSNOW['TypeEnum'] = {

    /**
     * value: "serviceNow"
     * @const
     */
    "serviceNow": "serviceNow"
};



export default IntegrationSNOW;
