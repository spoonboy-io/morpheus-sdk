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
import GuidanceSettings from './GuidanceSettings';

/**
 * The InlineObject79 model module.
 * @module model/InlineObject79
 * @version 6.2.1
 */
class InlineObject79 {
    /**
     * Constructs a new <code>InlineObject79</code>.
     * @alias module:model/InlineObject79
     */
    constructor() { 
        
        InlineObject79.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineObject79</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject79} obj Optional instance to populate.
     * @return {module:model/InlineObject79} The populated <code>InlineObject79</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject79();

            if (data.hasOwnProperty('guidanceSettings')) {
                obj['guidanceSettings'] = GuidanceSettings.constructFromObject(data['guidanceSettings']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/GuidanceSettings} guidanceSettings
 */
InlineObject79.prototype['guidanceSettings'] = undefined;






export default InlineObject79;

