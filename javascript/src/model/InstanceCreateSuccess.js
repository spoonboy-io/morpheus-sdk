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
import InstanceCreateSuccessInstance from './InstanceCreateSuccessInstance';

/**
 * The InstanceCreateSuccess model module.
 * @module model/InstanceCreateSuccess
 * @version 6.2.1
 */
class InstanceCreateSuccess {
    /**
     * Constructs a new <code>InstanceCreateSuccess</code>.
     * @alias module:model/InstanceCreateSuccess
     * @param instance {module:model/InstanceCreateSuccessInstance} 
     * @param zoneId {Number} The Cloud ID to provision the instance onto.
     */
    constructor(instance, zoneId) { 
        
        InstanceCreateSuccess.initialize(this, instance, zoneId);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, instance, zoneId) { 
        obj['instance'] = instance;
        obj['zoneId'] = zoneId;
    }

    /**
     * Constructs a <code>InstanceCreateSuccess</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceCreateSuccess} obj Optional instance to populate.
     * @return {module:model/InstanceCreateSuccess} The populated <code>InstanceCreateSuccess</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceCreateSuccess();

            if (data.hasOwnProperty('instance')) {
                obj['instance'] = InstanceCreateSuccessInstance.constructFromObject(data['instance']);
            }
            if (data.hasOwnProperty('zoneId')) {
                obj['zoneId'] = ApiClient.convertToType(data['zoneId'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/InstanceCreateSuccessInstance} instance
 */
InstanceCreateSuccess.prototype['instance'] = undefined;

/**
 * The Cloud ID to provision the instance onto.
 * @member {Number} zoneId
 */
InstanceCreateSuccess.prototype['zoneId'] = undefined;






export default InstanceCreateSuccess;

