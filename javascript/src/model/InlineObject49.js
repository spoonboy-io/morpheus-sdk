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

/**
 * The InlineObject49 model module.
 * @module model/InlineObject49
 * @version 6.2.1
 */
class InlineObject49 {
    /**
     * Constructs a new <code>InlineObject49</code>.
     * @alias module:model/InlineObject49
     * @param securityGroupIds {Array.<Number>} 
     */
    constructor(securityGroupIds) { 
        
        InlineObject49.initialize(this, securityGroupIds);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, securityGroupIds) { 
        obj['securityGroupIds'] = securityGroupIds;
    }

    /**
     * Constructs a <code>InlineObject49</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject49} obj Optional instance to populate.
     * @return {module:model/InlineObject49} The populated <code>InlineObject49</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject49();

            if (data.hasOwnProperty('securityGroupIds')) {
                obj['securityGroupIds'] = ApiClient.convertToType(data['securityGroupIds'], ['Number']);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<Number>} securityGroupIds
 */
InlineObject49.prototype['securityGroupIds'] = undefined;






export default InlineObject49;
