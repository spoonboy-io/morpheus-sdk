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
import NetworkCreate from './NetworkCreate';

/**
 * The InlineObject142 model module.
 * @module model/InlineObject142
 * @version 6.2.1
 */
class InlineObject142 {
    /**
     * Constructs a new <code>InlineObject142</code>.
     * @alias module:model/InlineObject142
     */
    constructor() { 
        
        InlineObject142.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineObject142</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject142} obj Optional instance to populate.
     * @return {module:model/InlineObject142} The populated <code>InlineObject142</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject142();

            if (data.hasOwnProperty('network')) {
                obj['network'] = NetworkCreate.constructFromObject(data['network']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/NetworkCreate} network
 */
InlineObject142.prototype['network'] = undefined;






export default InlineObject142;

