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
import ApiSecurityGroupsSecurityGroup from './ApiSecurityGroupsSecurityGroup';

/**
 * The InlineObject213 model module.
 * @module model/InlineObject213
 * @version 6.2.1
 */
class InlineObject213 {
    /**
     * Constructs a new <code>InlineObject213</code>.
     * @alias module:model/InlineObject213
     * @param securityGroup {module:model/ApiSecurityGroupsSecurityGroup} 
     */
    constructor(securityGroup) { 
        
        InlineObject213.initialize(this, securityGroup);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, securityGroup) { 
        obj['securityGroup'] = securityGroup;
    }

    /**
     * Constructs a <code>InlineObject213</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject213} obj Optional instance to populate.
     * @return {module:model/InlineObject213} The populated <code>InlineObject213</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject213();

            if (data.hasOwnProperty('securityGroup')) {
                obj['securityGroup'] = ApiSecurityGroupsSecurityGroup.constructFromObject(data['securityGroup']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiSecurityGroupsSecurityGroup} securityGroup
 */
InlineObject213.prototype['securityGroup'] = undefined;






export default InlineObject213;

