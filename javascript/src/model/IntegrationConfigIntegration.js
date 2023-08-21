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
import OneOfobjectobject from './OneOfobjectobject';

/**
 * The IntegrationConfigIntegration model module.
 * @module model/IntegrationConfigIntegration
 * @version 6.2.1
 */
class IntegrationConfigIntegration {
    /**
     * Constructs a new <code>IntegrationConfigIntegration</code>.
     * @alias module:model/IntegrationConfigIntegration
     * @param name {String} Name, a unique identifier for the integration
     * @param type {module:model/IntegrationConfigIntegration.TypeEnum} Integration Type Code
     */
    constructor(name, type) { 
        
        IntegrationConfigIntegration.initialize(this, name, type);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name, type) { 
        obj['name'] = name;
        obj['type'] = type;
    }

    /**
     * Constructs a <code>IntegrationConfigIntegration</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/IntegrationConfigIntegration} obj Optional instance to populate.
     * @return {module:model/IntegrationConfigIntegration} The populated <code>IntegrationConfigIntegration</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new IntegrationConfigIntegration();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('refresh')) {
                obj['refresh'] = ApiClient.convertToType(data['refresh'], 'Boolean');
            }
            if (data.hasOwnProperty('credential')) {
                obj['credential'] = ApiClient.convertToType(data['credential'], OneOfobjectobject);
            }
        }
        return obj;
    }


}

/**
 * Name, a unique identifier for the integration
 * @member {String} name
 */
IntegrationConfigIntegration.prototype['name'] = undefined;

/**
 * Integration Type Code
 * @member {module:model/IntegrationConfigIntegration.TypeEnum} type
 */
IntegrationConfigIntegration.prototype['type'] = undefined;

/**
 * Set `true` to enable integration
 * @member {Boolean} enabled
 */
IntegrationConfigIntegration.prototype['enabled'] = undefined;

/**
 * Pass `false` to skip refresh.  By default, refresh is done on update, when it is supported by the integration type. 
 * @member {Boolean} refresh
 * @default true
 */
IntegrationConfigIntegration.prototype['refresh'] = true;

/**
 * Map containing Credential ID or the default {\"type\": \"local\"}  which means use the values set in the local task options username and password instead of associating a credential. 
 * @member {module:model/OneOfobjectobject} credential
 */
IntegrationConfigIntegration.prototype['credential'] = undefined;





/**
 * Allowed values for the <code>type</code> property.
 * @enum {String}
 * @readonly
 */
IntegrationConfigIntegration['TypeEnum'] = {

    /**
     * value: "ansible"
     * @const
     */
    "ansible": "ansible",

    /**
     * value: "ansibleTower"
     * @const
     */
    "ansibleTower": "ansibleTower",

    /**
     * value: "bindDns"
     * @const
     */
    "bindDns": "bindDns",

    /**
     * value: "chef"
     * @const
     */
    "chef": "chef",

    /**
     * value: "cherwell"
     * @const
     */
    "cherwell": "cherwell",

    /**
     * value: "cypher"
     * @const
     */
    "cypher": "cypher",

    /**
     * value: "docker.registry"
     * @const
     */
    "docker.registry": "docker.registry",

    /**
     * value: "git"
     * @const
     */
    "git": "git",

    /**
     * value: "github"
     * @const
     */
    "github": "github",

    /**
     * value: "microsoft.dns"
     * @const
     */
    "microsoft.dns": "microsoft.dns",

    /**
     * value: "powerDns"
     * @const
     */
    "powerDns": "powerDns",

    /**
     * value: "puppet"
     * @const
     */
    "puppet": "puppet",

    /**
     * value: "remedy"
     * @const
     */
    "remedy": "remedy",

    /**
     * value: "amazonDns"
     * @const
     */
    "amazonDns": "amazonDns",

    /**
     * value: "saltMaster"
     * @const
     */
    "saltMaster": "saltMaster",

    /**
     * value: "serviceNow"
     * @const
     */
    "serviceNow": "serviceNow",

    /**
     * value: "vro"
     * @const
     */
    "vro": "vro"
};



export default IntegrationConfigIntegration;
