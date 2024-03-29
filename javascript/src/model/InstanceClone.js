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
import InstanceCloneGroup from './InstanceCloneGroup';

/**
 * The InstanceClone model module.
 * @module model/InstanceClone
 * @version 6.2.1
 */
class InstanceClone {
    /**
     * Constructs a new <code>InstanceClone</code>.
     * @alias module:model/InstanceClone
     */
    constructor() { 
        
        InstanceClone.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InstanceClone</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceClone} obj Optional instance to populate.
     * @return {module:model/InstanceClone} The populated <code>InstanceClone</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceClone();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('group')) {
                obj['group'] = InstanceCloneGroup.constructFromObject(data['group']);
            }
        }
        return obj;
    }


}

/**
 * A name for the new cloned instance. If none is specified the existing name will be duplicated with the 'clone' suffix added.
 * @member {String} name
 */
InstanceClone.prototype['name'] = undefined;

/**
 * @member {module:model/InstanceCloneGroup} group
 */
InstanceClone.prototype['group'] = undefined;






export default InstanceClone;

