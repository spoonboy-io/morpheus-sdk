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
import OneOfstringlong from './OneOfstringlong';

/**
 * The CredentialAccessSecretKeyConfigIntegration model module.
 * @module model/CredentialAccessSecretKeyConfigIntegration
 * @version 6.2.1
 */
class CredentialAccessSecretKeyConfigIntegration {
    /**
     * Constructs a new <code>CredentialAccessSecretKeyConfigIntegration</code>.
     * Credential Store. ID of a Credential Integration. This can be set to store the credential in an external store. 
     * @alias module:model/CredentialAccessSecretKeyConfigIntegration
     */
    constructor() { 
        
        CredentialAccessSecretKeyConfigIntegration.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CredentialAccessSecretKeyConfigIntegration</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CredentialAccessSecretKeyConfigIntegration} obj Optional instance to populate.
     * @return {module:model/CredentialAccessSecretKeyConfigIntegration} The populated <code>CredentialAccessSecretKeyConfigIntegration</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CredentialAccessSecretKeyConfigIntegration();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], OneOfstringlong);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/OneOfstringlong} id
 */
CredentialAccessSecretKeyConfigIntegration.prototype['id'] = undefined;






export default CredentialAccessSecretKeyConfigIntegration;

