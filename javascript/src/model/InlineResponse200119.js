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
import NetworkService from './NetworkService';

/**
 * The InlineResponse200119 model module.
 * @module model/InlineResponse200119
 * @version 6.2.1
 */
class InlineResponse200119 {
    /**
     * Constructs a new <code>InlineResponse200119</code>.
     * @alias module:model/InlineResponse200119
     */
    constructor() { 
        
        InlineResponse200119.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse200119</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse200119} obj Optional instance to populate.
     * @return {module:model/InlineResponse200119} The populated <code>InlineResponse200119</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse200119();

            if (data.hasOwnProperty('networkServices')) {
                obj['networkServices'] = ApiClient.convertToType(data['networkServices'], [NetworkService]);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/NetworkService>} networkServices
 */
InlineResponse200119.prototype['networkServices'] = undefined;






export default InlineResponse200119;

