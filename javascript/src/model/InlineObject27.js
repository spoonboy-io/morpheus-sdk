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
import ApiClientsClient from './ApiClientsClient';

/**
 * The InlineObject27 model module.
 * @module model/InlineObject27
 * @version 6.2.1
 */
class InlineObject27 {
    /**
     * Constructs a new <code>InlineObject27</code>.
     * @alias module:model/InlineObject27
     * @param client {module:model/ApiClientsClient} 
     */
    constructor(client) { 
        
        InlineObject27.initialize(this, client);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, client) { 
        obj['client'] = client;
    }

    /**
     * Constructs a <code>InlineObject27</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject27} obj Optional instance to populate.
     * @return {module:model/InlineObject27} The populated <code>InlineObject27</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject27();

            if (data.hasOwnProperty('client')) {
                obj['client'] = ApiClientsClient.constructFromObject(data['client']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiClientsClient} client
 */
InlineObject27.prototype['client'] = undefined;






export default InlineObject27;

