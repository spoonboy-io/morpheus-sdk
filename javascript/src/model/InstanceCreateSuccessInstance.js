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
import InstanceCreateInstanceInstanceType from './InstanceCreateInstanceInstanceType';
import InstanceCreateInstanceLayout from './InstanceCreateInstanceLayout';
import InstanceCreateInstancePlan from './InstanceCreateInstancePlan';
import InstanceCreateInstanceSite from './InstanceCreateInstanceSite';

/**
 * The InstanceCreateSuccessInstance model module.
 * @module model/InstanceCreateSuccessInstance
 * @version 6.2.1
 */
class InstanceCreateSuccessInstance {
    /**
     * Constructs a new <code>InstanceCreateSuccessInstance</code>.
     * Key for name, site, instanceType layout, and plan.
     * @alias module:model/InstanceCreateSuccessInstance
     * @param name {String} Name of the instance to be created.
     * @param site {module:model/InstanceCreateInstanceSite} 
     * @param instanceType {module:model/InstanceCreateInstanceInstanceType} 
     * @param layout {module:model/InstanceCreateInstanceLayout} 
     * @param plan {module:model/InstanceCreateInstancePlan} 
     */
    constructor(name, site, instanceType, layout, plan) { 
        
        InstanceCreateSuccessInstance.initialize(this, name, site, instanceType, layout, plan);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name, site, instanceType, layout, plan) { 
        obj['name'] = name;
        obj['site'] = site;
        obj['instanceType'] = instanceType;
        obj['layout'] = layout;
        obj['plan'] = plan;
    }

    /**
     * Constructs a <code>InstanceCreateSuccessInstance</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceCreateSuccessInstance} obj Optional instance to populate.
     * @return {module:model/InstanceCreateSuccessInstance} The populated <code>InstanceCreateSuccessInstance</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceCreateSuccessInstance();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('site')) {
                obj['site'] = InstanceCreateInstanceSite.constructFromObject(data['site']);
            }
            if (data.hasOwnProperty('instanceType')) {
                obj['instanceType'] = InstanceCreateInstanceInstanceType.constructFromObject(data['instanceType']);
            }
            if (data.hasOwnProperty('layout')) {
                obj['layout'] = InstanceCreateInstanceLayout.constructFromObject(data['layout']);
            }
            if (data.hasOwnProperty('plan')) {
                obj['plan'] = InstanceCreateInstancePlan.constructFromObject(data['plan']);
            }
        }
        return obj;
    }


}

/**
 * Name of the instance to be created.
 * @member {String} name
 */
InstanceCreateSuccessInstance.prototype['name'] = undefined;

/**
 * @member {module:model/InstanceCreateInstanceSite} site
 */
InstanceCreateSuccessInstance.prototype['site'] = undefined;

/**
 * @member {module:model/InstanceCreateInstanceInstanceType} instanceType
 */
InstanceCreateSuccessInstance.prototype['instanceType'] = undefined;

/**
 * @member {module:model/InstanceCreateInstanceLayout} layout
 */
InstanceCreateSuccessInstance.prototype['layout'] = undefined;

/**
 * @member {module:model/InstanceCreateInstancePlan} plan
 */
InstanceCreateSuccessInstance.prototype['plan'] = undefined;






export default InstanceCreateSuccessInstance;

