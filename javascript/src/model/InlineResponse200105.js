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
import NetworkRouterRoute from './NetworkRouterRoute';

/**
 * The InlineResponse200105 model module.
 * @module model/InlineResponse200105
 * @version 6.2.1
 */
class InlineResponse200105 {
    /**
     * Constructs a new <code>InlineResponse200105</code>.
     * @alias module:model/InlineResponse200105
     */
    constructor() { 
        
        InlineResponse200105.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse200105</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse200105} obj Optional instance to populate.
     * @return {module:model/InlineResponse200105} The populated <code>InlineResponse200105</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse200105();

            if (data.hasOwnProperty('networkRoute')) {
                obj['networkRoute'] = NetworkRouterRoute.constructFromObject(data['networkRoute']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/NetworkRouterRoute} networkRoute
 */
InlineResponse200105.prototype['networkRoute'] = undefined;






export default InlineResponse200105;
