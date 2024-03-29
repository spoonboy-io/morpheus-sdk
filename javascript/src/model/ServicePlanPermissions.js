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
import InlineResponse20095NetworkRouterPermissionsTenantPermissions from './InlineResponse20095NetworkRouterPermissionsTenantPermissions';
import ServicePlanPermissionsResourcePermissions from './ServicePlanPermissionsResourcePermissions';

/**
 * The ServicePlanPermissions model module.
 * @module model/ServicePlanPermissions
 * @version 6.2.1
 */
class ServicePlanPermissions {
    /**
     * Constructs a new <code>ServicePlanPermissions</code>.
     * @alias module:model/ServicePlanPermissions
     */
    constructor() { 
        
        ServicePlanPermissions.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ServicePlanPermissions</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ServicePlanPermissions} obj Optional instance to populate.
     * @return {module:model/ServicePlanPermissions} The populated <code>ServicePlanPermissions</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ServicePlanPermissions();

            if (data.hasOwnProperty('resourcePermissions')) {
                obj['resourcePermissions'] = ServicePlanPermissionsResourcePermissions.constructFromObject(data['resourcePermissions']);
            }
            if (data.hasOwnProperty('tenantPermissions')) {
                obj['tenantPermissions'] = InlineResponse20095NetworkRouterPermissionsTenantPermissions.constructFromObject(data['tenantPermissions']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ServicePlanPermissionsResourcePermissions} resourcePermissions
 */
ServicePlanPermissions.prototype['resourcePermissions'] = undefined;

/**
 * @member {module:model/InlineResponse20095NetworkRouterPermissionsTenantPermissions} tenantPermissions
 */
ServicePlanPermissions.prototype['tenantPermissions'] = undefined;






export default ServicePlanPermissions;

