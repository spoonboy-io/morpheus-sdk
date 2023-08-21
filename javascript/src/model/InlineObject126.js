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
 * The InlineObject126 model module.
 * @module model/InlineObject126
 * @version 6.2.1
 */
class InlineObject126 {
    /**
     * Constructs a new <code>InlineObject126</code>.
     * @alias module:model/InlineObject126
     * @param license {String} License Key. This is a long and unique string of your Morpheus license. License keys can be requested from the [Morpheus Hub](https://morpheushub.com).
     */
    constructor(license) { 
        
        InlineObject126.initialize(this, license);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, license) { 
        obj['license'] = license;
    }

    /**
     * Constructs a <code>InlineObject126</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject126} obj Optional instance to populate.
     * @return {module:model/InlineObject126} The populated <code>InlineObject126</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject126();

            if (data.hasOwnProperty('license')) {
                obj['license'] = ApiClient.convertToType(data['license'], 'String');
            }
        }
        return obj;
    }


}

/**
 * License Key. This is a long and unique string of your Morpheus license. License keys can be requested from the [Morpheus Hub](https://morpheushub.com).
 * @member {String} license
 */
InlineObject126.prototype['license'] = undefined;






export default InlineObject126;
