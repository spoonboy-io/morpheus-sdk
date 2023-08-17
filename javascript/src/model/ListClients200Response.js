/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.1.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import Client from './Client';
import Meta from './Meta';
import MetaMeta from './MetaMeta';

/**
 * The ListClients200Response model module.
 * @module model/ListClients200Response
 * @version 6.1.1
 */
class ListClients200Response {
    /**
     * Constructs a new <code>ListClients200Response</code>.
     * @alias module:model/ListClients200Response
     * @implements module:model/Meta
     */
    constructor() { 
        Meta.initialize(this);
        ListClients200Response.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ListClients200Response</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ListClients200Response} obj Optional instance to populate.
     * @return {module:model/ListClients200Response} The populated <code>ListClients200Response</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ListClients200Response();
            Meta.constructFromObject(data, obj);

            if (data.hasOwnProperty('meta')) {
                obj['meta'] = MetaMeta.constructFromObject(data['meta']);
            }
            if (data.hasOwnProperty('clients')) {
                obj['clients'] = ApiClient.convertToType(data['clients'], [Client]);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>ListClients200Response</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>ListClients200Response</code>.
     */
    static validateJSON(data) {
        // validate the optional field `meta`
        if (data['meta']) { // data not null
          MetaMeta.validateJSON(data['meta']);
        }
        if (data['clients']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['clients'])) {
                throw new Error("Expected the field `clients` to be an array in the JSON data but got " + data['clients']);
            }
            // validate the optional field `clients` (array)
            for (const item of data['clients']) {
                Client.validateJSON(item);
            };
        }

        return true;
    }


}



/**
 * @member {module:model/MetaMeta} meta
 */
ListClients200Response.prototype['meta'] = undefined;

/**
 * @member {Array.<module:model/Client>} clients
 */
ListClients200Response.prototype['clients'] = undefined;


// Implement Meta interface:
/**
 * @member {module:model/MetaMeta} meta
 */
Meta.prototype['meta'] = undefined;




export default ListClients200Response;
