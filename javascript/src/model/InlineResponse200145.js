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
import StorageBucket from './StorageBucket';

/**
 * The InlineResponse200145 model module.
 * @module model/InlineResponse200145
 * @version 6.2.1
 */
class InlineResponse200145 {
    /**
     * Constructs a new <code>InlineResponse200145</code>.
     * @alias module:model/InlineResponse200145
     */
    constructor() { 
        
        InlineResponse200145.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse200145</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse200145} obj Optional instance to populate.
     * @return {module:model/InlineResponse200145} The populated <code>InlineResponse200145</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse200145();

            if (data.hasOwnProperty('storageBucket')) {
                obj['storageBucket'] = StorageBucket.constructFromObject(data['storageBucket']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/StorageBucket} storageBucket
 */
InlineResponse200145.prototype['storageBucket'] = undefined;






export default InlineResponse200145;
