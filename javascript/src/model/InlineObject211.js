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
import ApiScaleThresholdsScaleThreshold from './ApiScaleThresholdsScaleThreshold';

/**
 * The InlineObject211 model module.
 * @module model/InlineObject211
 * @version 6.2.1
 */
class InlineObject211 {
    /**
     * Constructs a new <code>InlineObject211</code>.
     * @alias module:model/InlineObject211
     * @param scaleThreshold {module:model/ApiScaleThresholdsScaleThreshold} 
     */
    constructor(scaleThreshold) { 
        
        InlineObject211.initialize(this, scaleThreshold);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, scaleThreshold) { 
        obj['scaleThreshold'] = scaleThreshold;
    }

    /**
     * Constructs a <code>InlineObject211</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject211} obj Optional instance to populate.
     * @return {module:model/InlineObject211} The populated <code>InlineObject211</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject211();

            if (data.hasOwnProperty('scaleThreshold')) {
                obj['scaleThreshold'] = ApiScaleThresholdsScaleThreshold.constructFromObject(data['scaleThreshold']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiScaleThresholdsScaleThreshold} scaleThreshold
 */
InlineObject211.prototype['scaleThreshold'] = undefined;






export default InlineObject211;

