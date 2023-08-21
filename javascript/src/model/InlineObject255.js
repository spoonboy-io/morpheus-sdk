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
import UserCreate from './UserCreate';

/**
 * The InlineObject255 model module.
 * @module model/InlineObject255
 * @version 6.2.1
 */
class InlineObject255 {
    /**
     * Constructs a new <code>InlineObject255</code>.
     * @alias module:model/InlineObject255
     * @param user {module:model/UserCreate} 
     */
    constructor(user) { 
        
        InlineObject255.initialize(this, user);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, user) { 
        obj['user'] = user;
    }

    /**
     * Constructs a <code>InlineObject255</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject255} obj Optional instance to populate.
     * @return {module:model/InlineObject255} The populated <code>InlineObject255</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject255();

            if (data.hasOwnProperty('user')) {
                obj['user'] = UserCreate.constructFromObject(data['user']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/UserCreate} user
 */
InlineObject255.prototype['user'] = undefined;






export default InlineObject255;
