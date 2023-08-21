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
 * The IdentitySourcesSAMLConfigProviderSettings model module.
 * @module model/IdentitySourcesSAMLConfigProviderSettings
 * @version 6.2.1
 */
class IdentitySourcesSAMLConfigProviderSettings {
    /**
     * Constructs a new <code>IdentitySourcesSAMLConfigProviderSettings</code>.
     * @alias module:model/IdentitySourcesSAMLConfigProviderSettings
     */
    constructor() { 
        
        IdentitySourcesSAMLConfigProviderSettings.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>IdentitySourcesSAMLConfigProviderSettings</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/IdentitySourcesSAMLConfigProviderSettings} obj Optional instance to populate.
     * @return {module:model/IdentitySourcesSAMLConfigProviderSettings} The populated <code>IdentitySourcesSAMLConfigProviderSettings</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new IdentitySourcesSAMLConfigProviderSettings();

            if (data.hasOwnProperty('entityId')) {
                obj['entityId'] = ApiClient.convertToType(data['entityId'], 'String');
            }
            if (data.hasOwnProperty('acsUrl')) {
                obj['acsUrl'] = ApiClient.convertToType(data['acsUrl'], 'String');
            }
            if (data.hasOwnProperty('spMetadata')) {
                obj['spMetadata'] = ApiClient.convertToType(data['spMetadata'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} entityId
 */
IdentitySourcesSAMLConfigProviderSettings.prototype['entityId'] = undefined;

/**
 * @member {String} acsUrl
 */
IdentitySourcesSAMLConfigProviderSettings.prototype['acsUrl'] = undefined;

/**
 * @member {String} spMetadata
 */
IdentitySourcesSAMLConfigProviderSettings.prototype['spMetadata'] = undefined;






export default IdentitySourcesSAMLConfigProviderSettings;
