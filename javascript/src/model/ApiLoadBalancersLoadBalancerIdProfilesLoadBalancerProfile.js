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

/**
 * The ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile model module.
 * @module model/ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile
 * @version 6.2.1
 */
class ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile {
    /**
     * Constructs a new <code>ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile</code>.
     * @alias module:model/ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile
     */
    constructor() { 
        
        ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile} obj Optional instance to populate.
     * @return {module:model/ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile} The populated <code>ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('serviceType')) {
                obj['serviceType'] = ApiClient.convertToType(data['serviceType'], 'String');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], Object);
            }
        }
        return obj;
    }


}

/**
 * Name
 * @member {String} name
 */
ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile.prototype['name'] = undefined;

/**
 * Description
 * @member {String} description
 */
ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile.prototype['description'] = undefined;

/**
 * Service Type
 * @member {String} serviceType
 */
ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile.prototype['serviceType'] = undefined;

/**
 * Configuration object with parameters that vary by type.
 * @member {Object} config
 */
ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile.prototype['config'] = undefined;






export default ApiLoadBalancersLoadBalancerIdProfilesLoadBalancerProfile;

