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
import NetworkRoutersCreate from './NetworkRoutersCreate';

/**
 * The InlineObject148 model module.
 * @module model/InlineObject148
 * @version 6.2.1
 */
class InlineObject148 {
    /**
     * Constructs a new <code>InlineObject148</code>.
     * The parameters for creating a network router is type dependent. The following lists the common parameters. See get a specific type to list available options for that network router type. Note: when creating a router on NSX v3.0+ some BGP configuration settings require BGP to be disabled during initial creation. The BGP feature can be enabled in a subsequent router update API call. 
     * @alias module:model/InlineObject148
     */
    constructor() { 
        
        InlineObject148.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineObject148</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject148} obj Optional instance to populate.
     * @return {module:model/InlineObject148} The populated <code>InlineObject148</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject148();

            if (data.hasOwnProperty('networkRouter')) {
                obj['networkRouter'] = NetworkRoutersCreate.constructFromObject(data['networkRouter']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/NetworkRoutersCreate} networkRouter
 */
InlineObject148.prototype['networkRouter'] = undefined;






export default InlineObject148;

