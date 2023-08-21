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
import ClientUpdate from './ClientUpdate';

/**
 * The InlineObject28 model module.
 * @module model/InlineObject28
 * @version 6.2.1
 */
class InlineObject28 {
    /**
     * Constructs a new <code>InlineObject28</code>.
     * @alias module:model/InlineObject28
     * @param client {module:model/ClientUpdate} 
     */
    constructor(client) { 
        
        InlineObject28.initialize(this, client);
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
     * Constructs a <code>InlineObject28</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject28} obj Optional instance to populate.
     * @return {module:model/InlineObject28} The populated <code>InlineObject28</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject28();

            if (data.hasOwnProperty('client')) {
                obj['client'] = ClientUpdate.constructFromObject(data['client']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ClientUpdate} client
 */
InlineObject28.prototype['client'] = undefined;






export default InlineObject28;

