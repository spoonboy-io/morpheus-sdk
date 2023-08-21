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
import PolicyCloudUpdatePolicyType from './PolicyCloudUpdatePolicyType';

/**
 * The PolicyCloudUpdate model module.
 * @module model/PolicyCloudUpdate
 * @version 6.2.1
 */
class PolicyCloudUpdate {
    /**
     * Constructs a new <code>PolicyCloudUpdate</code>.
     * @alias module:model/PolicyCloudUpdate
     */
    constructor() { 
        
        PolicyCloudUpdate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>PolicyCloudUpdate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/PolicyCloudUpdate} obj Optional instance to populate.
     * @return {module:model/PolicyCloudUpdate} The populated <code>PolicyCloudUpdate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new PolicyCloudUpdate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('policyType')) {
                obj['policyType'] = PolicyCloudUpdatePolicyType.constructFromObject(data['policyType']);
            }
        }
        return obj;
    }


}

/**
 * A name for the policy
 * @member {String} name
 */
PolicyCloudUpdate.prototype['name'] = undefined;

/**
 * A description for the policy
 * @member {String} description
 */
PolicyCloudUpdate.prototype['description'] = undefined;

/**
 * @member {module:model/PolicyCloudUpdatePolicyType} policyType
 */
PolicyCloudUpdate.prototype['policyType'] = undefined;






export default PolicyCloudUpdate;
