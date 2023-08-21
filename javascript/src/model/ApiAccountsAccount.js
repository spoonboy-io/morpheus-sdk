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
import ApiAccountsAccountRole from './ApiAccountsAccountRole';
import CurrencyCode from './CurrencyCode';

/**
 * The ApiAccountsAccount model module.
 * @module model/ApiAccountsAccount
 * @version 6.2.1
 */
class ApiAccountsAccount {
    /**
     * Constructs a new <code>ApiAccountsAccount</code>.
     * Payload for creating a new tenant
     * @alias module:model/ApiAccountsAccount
     * @param name {String} Name
     */
    constructor(name) { 
        
        ApiAccountsAccount.initialize(this, name);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name) { 
        obj['name'] = name;
    }

    /**
     * Constructs a <code>ApiAccountsAccount</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiAccountsAccount} obj Optional instance to populate.
     * @return {module:model/ApiAccountsAccount} The populated <code>ApiAccountsAccount</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiAccountsAccount();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('role')) {
                obj['role'] = ApiAccountsAccountRole.constructFromObject(data['role']);
            }
            if (data.hasOwnProperty('subdomain')) {
                obj['subdomain'] = ApiClient.convertToType(data['subdomain'], 'String');
            }
            if (data.hasOwnProperty('currency')) {
                obj['currency'] = CurrencyCode.constructFromObject(data['currency']);
            }
        }
        return obj;
    }


}

/**
 * Name
 * @member {String} name
 */
ApiAccountsAccount.prototype['name'] = undefined;

/**
 * Description
 * @member {String} description
 */
ApiAccountsAccount.prototype['description'] = undefined;

/**
 * @member {module:model/ApiAccountsAccountRole} role
 */
ApiAccountsAccount.prototype['role'] = undefined;

/**
 * The subdomain. This will be part of the login URL and username for sub tenant users.
 * @member {String} subdomain
 */
ApiAccountsAccount.prototype['subdomain'] = undefined;

/**
 * @member {module:model/CurrencyCode} currency
 */
ApiAccountsAccount.prototype['currency'] = undefined;






export default ApiAccountsAccount;
