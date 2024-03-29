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
import UserSourceCreateUserSource from './UserSourceCreateUserSource';

/**
 * The UserSourceCreate model module.
 * @module model/UserSourceCreate
 * @version 6.2.1
 */
class UserSourceCreate {
    /**
     * Constructs a new <code>UserSourceCreate</code>.
     * @alias module:model/UserSourceCreate
     * @param userSource {module:model/UserSourceCreateUserSource} 
     */
    constructor(userSource) { 
        
        UserSourceCreate.initialize(this, userSource);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, userSource) { 
        obj['userSource'] = userSource;
    }

    /**
     * Constructs a <code>UserSourceCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/UserSourceCreate} obj Optional instance to populate.
     * @return {module:model/UserSourceCreate} The populated <code>UserSourceCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new UserSourceCreate();

            if (data.hasOwnProperty('userSource')) {
                obj['userSource'] = UserSourceCreateUserSource.constructFromObject(data['userSource']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/UserSourceCreateUserSource} userSource
 */
UserSourceCreate.prototype['userSource'] = undefined;






export default UserSourceCreate;

