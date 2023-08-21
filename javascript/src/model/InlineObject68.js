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
import DeploymentCreate from './DeploymentCreate';

/**
 * The InlineObject68 model module.
 * @module model/InlineObject68
 * @version 6.2.1
 */
class InlineObject68 {
    /**
     * Constructs a new <code>InlineObject68</code>.
     * @alias module:model/InlineObject68
     */
    constructor() { 
        
        InlineObject68.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineObject68</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject68} obj Optional instance to populate.
     * @return {module:model/InlineObject68} The populated <code>InlineObject68</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject68();

            if (data.hasOwnProperty('deployment')) {
                obj['deployment'] = DeploymentCreate.constructFromObject(data['deployment']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/DeploymentCreate} deployment
 */
InlineObject68.prototype['deployment'] = undefined;






export default InlineObject68;
