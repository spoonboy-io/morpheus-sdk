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
import NetworkType from './NetworkType';

/**
 * The InlineResponse20086 model module.
 * @module model/InlineResponse20086
 * @version 6.2.1
 */
class InlineResponse20086 {
    /**
     * Constructs a new <code>InlineResponse20086</code>.
     * @alias module:model/InlineResponse20086
     */
    constructor() { 
        
        InlineResponse20086.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse20086</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse20086} obj Optional instance to populate.
     * @return {module:model/InlineResponse20086} The populated <code>InlineResponse20086</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse20086();

            if (data.hasOwnProperty('networkType')) {
                obj['networkType'] = NetworkType.constructFromObject(data['networkType']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/NetworkType} networkType
 */
InlineResponse20086.prototype['networkType'] = undefined;






export default InlineResponse20086;
