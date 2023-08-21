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
 * The CredentialConfig model module.
 * @module model/CredentialConfig
 * @version 6.2.1
 */
class CredentialConfig {
    /**
     * Constructs a new <code>CredentialConfig</code>.
     * @alias module:model/CredentialConfig
     */
    constructor() { 
        
        CredentialConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CredentialConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CredentialConfig} obj Optional instance to populate.
     * @return {module:model/CredentialConfig} The populated <code>CredentialConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CredentialConfig();

            if (data.hasOwnProperty('clientSecret')) {
                obj['clientSecret'] = ApiClient.convertToType(data['clientSecret'], 'String');
            }
            if (data.hasOwnProperty('clientId')) {
                obj['clientId'] = ApiClient.convertToType(data['clientId'], 'String');
            }
            if (data.hasOwnProperty('clientAuth')) {
                obj['clientAuth'] = ApiClient.convertToType(data['clientAuth'], 'String');
            }
            if (data.hasOwnProperty('scope')) {
                obj['scope'] = ApiClient.convertToType(data['scope'], 'String');
            }
            if (data.hasOwnProperty('grantType')) {
                obj['grantType'] = ApiClient.convertToType(data['grantType'], 'String');
            }
            if (data.hasOwnProperty('accessTokenUrl')) {
                obj['accessTokenUrl'] = ApiClient.convertToType(data['accessTokenUrl'], 'String');
            }
            if (data.hasOwnProperty('clientSecretHash')) {
                obj['clientSecretHash'] = ApiClient.convertToType(data['clientSecretHash'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} clientSecret
 */
CredentialConfig.prototype['clientSecret'] = undefined;

/**
 * @member {String} clientId
 */
CredentialConfig.prototype['clientId'] = undefined;

/**
 * @member {String} clientAuth
 */
CredentialConfig.prototype['clientAuth'] = undefined;

/**
 * @member {String} scope
 */
CredentialConfig.prototype['scope'] = undefined;

/**
 * @member {String} grantType
 */
CredentialConfig.prototype['grantType'] = undefined;

/**
 * @member {String} accessTokenUrl
 */
CredentialConfig.prototype['accessTokenUrl'] = undefined;

/**
 * @member {String} clientSecretHash
 */
CredentialConfig.prototype['clientSecretHash'] = undefined;






export default CredentialConfig;

