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
import OneOfobject from './OneOfobject';

/**
 * The InlineObject259 model module.
 * @module model/InlineObject259
 * @version 6.2.1
 */
class InlineObject259 {
    /**
     * Constructs a new <code>InlineObject259</code>.
     * @alias module:model/InlineObject259
     * @param vdiGateway {module:model/OneOfobject} 
     */
    constructor(vdiGateway) { 
        
        InlineObject259.initialize(this, vdiGateway);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, vdiGateway) { 
        obj['vdiGateway'] = vdiGateway;
    }

    /**
     * Constructs a <code>InlineObject259</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject259} obj Optional instance to populate.
     * @return {module:model/InlineObject259} The populated <code>InlineObject259</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject259();

            if (data.hasOwnProperty('vdiGateway')) {
                obj['vdiGateway'] = ApiClient.convertToType(data['vdiGateway'], OneOfobject);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/OneOfobject} vdiGateway
 */
InlineObject259.prototype['vdiGateway'] = undefined;






export default InlineObject259;

