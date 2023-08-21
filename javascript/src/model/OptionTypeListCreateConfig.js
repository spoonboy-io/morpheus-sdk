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
import OptionTypeListCreateConfigSourceHeaders from './OptionTypeListCreateConfigSourceHeaders';

/**
 * The OptionTypeListCreateConfig model module.
 * @module model/OptionTypeListCreateConfig
 * @version 6.2.1
 */
class OptionTypeListCreateConfig {
    /**
     * Constructs a new <code>OptionTypeListCreateConfig</code>.
     * Array of source headers to use when requesting data
     * @alias module:model/OptionTypeListCreateConfig
     */
    constructor() { 
        
        OptionTypeListCreateConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>OptionTypeListCreateConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/OptionTypeListCreateConfig} obj Optional instance to populate.
     * @return {module:model/OptionTypeListCreateConfig} The populated <code>OptionTypeListCreateConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new OptionTypeListCreateConfig();

            if (data.hasOwnProperty('sourceHeaders')) {
                obj['sourceHeaders'] = ApiClient.convertToType(data['sourceHeaders'], [OptionTypeListCreateConfigSourceHeaders]);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/OptionTypeListCreateConfigSourceHeaders>} sourceHeaders
 */
OptionTypeListCreateConfig.prototype['sourceHeaders'] = undefined;






export default OptionTypeListCreateConfig;
