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
import InstanceUpdateInstance from './InstanceUpdateInstance';
import InstancesConfigCustomOptions from './InstancesConfigCustomOptions';

/**
 * The InstanceUpdate model module.
 * @module model/InstanceUpdate
 * @version 6.2.1
 */
class InstanceUpdate {
    /**
     * Constructs a new <code>InstanceUpdate</code>.
     * @alias module:model/InstanceUpdate
     */
    constructor() { 
        
        InstanceUpdate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InstanceUpdate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceUpdate} obj Optional instance to populate.
     * @return {module:model/InstanceUpdate} The populated <code>InstanceUpdate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceUpdate();

            if (data.hasOwnProperty('instance')) {
                obj['instance'] = InstanceUpdateInstance.constructFromObject(data['instance']);
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = InstancesConfigCustomOptions.constructFromObject(data['config']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/InstanceUpdateInstance} instance
 */
InstanceUpdate.prototype['instance'] = undefined;

/**
 * @member {module:model/InstancesConfigCustomOptions} config
 */
InstanceUpdate.prototype['config'] = undefined;






export default InstanceUpdate;
