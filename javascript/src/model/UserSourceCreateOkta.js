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
 * The UserSourceCreateOkta model module.
 * @module model/UserSourceCreateOkta
 * @version 6.2.1
 */
class UserSourceCreateOkta {
    /**
     * Constructs a new <code>UserSourceCreateOkta</code>.
     * Okta Configuration
     * @alias module:model/UserSourceCreateOkta
     */
    constructor() { 
        
        UserSourceCreateOkta.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>UserSourceCreateOkta</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/UserSourceCreateOkta} obj Optional instance to populate.
     * @return {module:model/UserSourceCreateOkta} The populated <code>UserSourceCreateOkta</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new UserSourceCreateOkta();

            if (data.hasOwnProperty('url')) {
                obj['url'] = ApiClient.convertToType(data['url'], 'String');
            }
            if (data.hasOwnProperty('administratorAPIToken')) {
                obj['administratorAPIToken'] = ApiClient.convertToType(data['administratorAPIToken'], 'String');
            }
            if (data.hasOwnProperty('requiredGroup')) {
                obj['requiredGroup'] = ApiClient.convertToType(data['requiredGroup'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Okta URL
 * @member {String} url
 */
UserSourceCreateOkta.prototype['url'] = undefined;

/**
 * Administrator API Token
 * @member {String} administratorAPIToken
 */
UserSourceCreateOkta.prototype['administratorAPIToken'] = undefined;

/**
 * Required Group
 * @member {String} requiredGroup
 */
UserSourceCreateOkta.prototype['requiredGroup'] = undefined;






export default UserSourceCreateOkta;

