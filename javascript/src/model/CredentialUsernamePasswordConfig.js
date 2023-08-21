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
import CredentialAccessSecretKeyConfigIntegration from './CredentialAccessSecretKeyConfigIntegration';

/**
 * The CredentialUsernamePasswordConfig model module.
 * @module model/CredentialUsernamePasswordConfig
 * @version 6.2.1
 */
class CredentialUsernamePasswordConfig {
    /**
     * Constructs a new <code>CredentialUsernamePasswordConfig</code>.
     * @alias module:model/CredentialUsernamePasswordConfig
     * @param type {module:model/CredentialUsernamePasswordConfig.TypeEnum} Credential Type Code
     * @param name {String} A unique name scoped to your account for the credential
     * @param username {String} Username
     * @param password {String} Password
     */
    constructor(type, name, username, password) { 
        
        CredentialUsernamePasswordConfig.initialize(this, type, name, username, password);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, type, name, username, password) { 
        obj['type'] = type;
        obj['name'] = name;
        obj['username'] = username;
        obj['password'] = password;
    }

    /**
     * Constructs a <code>CredentialUsernamePasswordConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CredentialUsernamePasswordConfig} obj Optional instance to populate.
     * @return {module:model/CredentialUsernamePasswordConfig} The populated <code>CredentialUsernamePasswordConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CredentialUsernamePasswordConfig();

            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('integration')) {
                obj['integration'] = CredentialAccessSecretKeyConfigIntegration.constructFromObject(data['integration']);
            }
            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
            }
            if (data.hasOwnProperty('password')) {
                obj['password'] = ApiClient.convertToType(data['password'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Credential Type Code
 * @member {module:model/CredentialUsernamePasswordConfig.TypeEnum} type
 */
CredentialUsernamePasswordConfig.prototype['type'] = undefined;

/**
 * A unique name scoped to your account for the credential
 * @member {String} name
 */
CredentialUsernamePasswordConfig.prototype['name'] = undefined;

/**
 * Optional Description
 * @member {String} description
 */
CredentialUsernamePasswordConfig.prototype['description'] = undefined;

/**
 * Credential enabled
 * @member {Boolean} enabled
 * @default true
 */
CredentialUsernamePasswordConfig.prototype['enabled'] = true;

/**
 * @member {module:model/CredentialAccessSecretKeyConfigIntegration} integration
 */
CredentialUsernamePasswordConfig.prototype['integration'] = undefined;

/**
 * Username
 * @member {String} username
 */
CredentialUsernamePasswordConfig.prototype['username'] = undefined;

/**
 * Password
 * @member {String} password
 */
CredentialUsernamePasswordConfig.prototype['password'] = undefined;





/**
 * Allowed values for the <code>type</code> property.
 * @enum {String}
 * @readonly
 */
CredentialUsernamePasswordConfig['TypeEnum'] = {

    /**
     * value: "username-password"
     * @const
     */
    "username-password": "username-password"
};



export default CredentialUsernamePasswordConfig;
