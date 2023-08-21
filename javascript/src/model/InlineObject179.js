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
import NetworkTransportZoneCreate from './NetworkTransportZoneCreate';

/**
 * The InlineObject179 model module.
 * @module model/InlineObject179
 * @version 6.2.1
 */
class InlineObject179 {
    /**
     * Constructs a new <code>InlineObject179</code>.
     * The parameters for update a Network Transport Zone is type dependent. The following lists the common parameters. Get a specific network type to list available options for the network relay type. 
     * @alias module:model/InlineObject179
     */
    constructor() { 
        
        InlineObject179.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineObject179</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject179} obj Optional instance to populate.
     * @return {module:model/InlineObject179} The populated <code>InlineObject179</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject179();

            if (data.hasOwnProperty('networkScope')) {
                obj['networkScope'] = NetworkTransportZoneCreate.constructFromObject(data['networkScope']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/NetworkTransportZoneCreate} networkScope
 */
InlineObject179.prototype['networkScope'] = undefined;






export default InlineObject179;

