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
 * The InlineObject178 model module.
 * @module model/InlineObject178
 * @version 6.2.1
 */
class InlineObject178 {
    /**
     * Constructs a new <code>InlineObject178</code>.
     * The parameters for creating a network transport zone is type dependent. The following lists the common parameters. See get a specific type to list available options for the network server type.
     * @alias module:model/InlineObject178
     */
    constructor() { 
        
        InlineObject178.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineObject178</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject178} obj Optional instance to populate.
     * @return {module:model/InlineObject178} The populated <code>InlineObject178</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject178();

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
InlineObject178.prototype['networkScope'] = undefined;






export default InlineObject178;

