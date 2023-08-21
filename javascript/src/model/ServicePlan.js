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
import GuidanceVmwareSizingPlanBeforeActionPriceSets from './GuidanceVmwareSizingPlanBeforeActionPriceSets';
import GuidanceVmwareSizingPlanBeforeActionProvisionType from './GuidanceVmwareSizingPlanBeforeActionProvisionType';
import InlineResponse20094Network from './InlineResponse20094Network';
import ServicePlanConfig from './ServicePlanConfig';
import ServicePlanPermissions from './ServicePlanPermissions';

/**
 * The ServicePlan model module.
 * @module model/ServicePlan
 * @version 6.2.1
 */
class ServicePlan {
    /**
     * Constructs a new <code>ServicePlan</code>.
     * @alias module:model/ServicePlan
     */
    constructor() { 
        
        ServicePlan.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ServicePlan</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ServicePlan} obj Optional instance to populate.
     * @return {module:model/ServicePlan} The populated <code>ServicePlan</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ServicePlan();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
            if (data.hasOwnProperty('sortOrder')) {
                obj['sortOrder'] = ApiClient.convertToType(data['sortOrder'], 'Number');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('maxStorage')) {
                obj['maxStorage'] = ApiClient.convertToType(data['maxStorage'], 'Number');
            }
            if (data.hasOwnProperty('maxMemory')) {
                obj['maxMemory'] = ApiClient.convertToType(data['maxMemory'], 'Number');
            }
            if (data.hasOwnProperty('maxCpu')) {
                obj['maxCpu'] = ApiClient.convertToType(data['maxCpu'], 'Number');
            }
            if (data.hasOwnProperty('maxCores')) {
                obj['maxCores'] = ApiClient.convertToType(data['maxCores'], 'Number');
            }
            if (data.hasOwnProperty('maxDisks')) {
                obj['maxDisks'] = ApiClient.convertToType(data['maxDisks'], 'Number');
            }
            if (data.hasOwnProperty('coresPerSocket')) {
                obj['coresPerSocket'] = ApiClient.convertToType(data['coresPerSocket'], 'Number');
            }
            if (data.hasOwnProperty('customCpu')) {
                obj['customCpu'] = ApiClient.convertToType(data['customCpu'], 'Boolean');
            }
            if (data.hasOwnProperty('customCores')) {
                obj['customCores'] = ApiClient.convertToType(data['customCores'], 'Boolean');
            }
            if (data.hasOwnProperty('customMaxStorage')) {
                obj['customMaxStorage'] = ApiClient.convertToType(data['customMaxStorage'], 'Boolean');
            }
            if (data.hasOwnProperty('customMaxDataStorage')) {
                obj['customMaxDataStorage'] = ApiClient.convertToType(data['customMaxDataStorage'], 'Boolean');
            }
            if (data.hasOwnProperty('customMaxMemory')) {
                obj['customMaxMemory'] = ApiClient.convertToType(data['customMaxMemory'], 'Boolean');
            }
            if (data.hasOwnProperty('addVolumes')) {
                obj['addVolumes'] = ApiClient.convertToType(data['addVolumes'], 'Boolean');
            }
            if (data.hasOwnProperty('memoryOptionSource')) {
                obj['memoryOptionSource'] = ApiClient.convertToType(data['memoryOptionSource'], 'String');
            }
            if (data.hasOwnProperty('cpuOptionSource')) {
                obj['cpuOptionSource'] = ApiClient.convertToType(data['cpuOptionSource'], 'String');
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
            if (data.hasOwnProperty('regionCode')) {
                obj['regionCode'] = ApiClient.convertToType(data['regionCode'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('editable')) {
                obj['editable'] = ApiClient.convertToType(data['editable'], 'Boolean');
            }
            if (data.hasOwnProperty('provisionType')) {
                obj['provisionType'] = GuidanceVmwareSizingPlanBeforeActionProvisionType.constructFromObject(data['provisionType']);
            }
            if (data.hasOwnProperty('tenants')) {
                obj['tenants'] = ApiClient.convertToType(data['tenants'], 'String');
            }
            if (data.hasOwnProperty('priceSets')) {
                obj['priceSets'] = ApiClient.convertToType(data['priceSets'], [GuidanceVmwareSizingPlanBeforeActionPriceSets]);
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ServicePlanConfig.constructFromObject(data['config']);
            }
            if (data.hasOwnProperty('zones')) {
                obj['zones'] = ApiClient.convertToType(data['zones'], [InlineResponse20094Network]);
            }
            if (data.hasOwnProperty('permissions')) {
                obj['permissions'] = ServicePlanPermissions.constructFromObject(data['permissions']);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ServicePlan.prototype['id'] = undefined;

/**
 * @member {String} name
 */
ServicePlan.prototype['name'] = undefined;

/**
 * @member {String} code
 */
ServicePlan.prototype['code'] = undefined;

/**
 * @member {Boolean} active
 */
ServicePlan.prototype['active'] = undefined;

/**
 * @member {Number} sortOrder
 */
ServicePlan.prototype['sortOrder'] = undefined;

/**
 * @member {String} description
 */
ServicePlan.prototype['description'] = undefined;

/**
 * @member {Number} maxStorage
 */
ServicePlan.prototype['maxStorage'] = undefined;

/**
 * @member {Number} maxMemory
 */
ServicePlan.prototype['maxMemory'] = undefined;

/**
 * @member {Number} maxCpu
 */
ServicePlan.prototype['maxCpu'] = undefined;

/**
 * @member {Number} maxCores
 */
ServicePlan.prototype['maxCores'] = undefined;

/**
 * @member {Number} maxDisks
 */
ServicePlan.prototype['maxDisks'] = undefined;

/**
 * @member {Number} coresPerSocket
 */
ServicePlan.prototype['coresPerSocket'] = undefined;

/**
 * @member {Boolean} customCpu
 */
ServicePlan.prototype['customCpu'] = undefined;

/**
 * @member {Boolean} customCores
 */
ServicePlan.prototype['customCores'] = undefined;

/**
 * @member {Boolean} customMaxStorage
 */
ServicePlan.prototype['customMaxStorage'] = undefined;

/**
 * @member {Boolean} customMaxDataStorage
 */
ServicePlan.prototype['customMaxDataStorage'] = undefined;

/**
 * @member {Boolean} customMaxMemory
 */
ServicePlan.prototype['customMaxMemory'] = undefined;

/**
 * @member {Boolean} addVolumes
 */
ServicePlan.prototype['addVolumes'] = undefined;

/**
 * @member {String} memoryOptionSource
 */
ServicePlan.prototype['memoryOptionSource'] = undefined;

/**
 * @member {String} cpuOptionSource
 */
ServicePlan.prototype['cpuOptionSource'] = undefined;

/**
 * @member {Date} dateCreated
 */
ServicePlan.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
ServicePlan.prototype['lastUpdated'] = undefined;

/**
 * @member {String} regionCode
 */
ServicePlan.prototype['regionCode'] = undefined;

/**
 * @member {String} visibility
 */
ServicePlan.prototype['visibility'] = undefined;

/**
 * @member {Boolean} editable
 */
ServicePlan.prototype['editable'] = undefined;

/**
 * @member {module:model/GuidanceVmwareSizingPlanBeforeActionProvisionType} provisionType
 */
ServicePlan.prototype['provisionType'] = undefined;

/**
 * @member {String} tenants
 */
ServicePlan.prototype['tenants'] = undefined;

/**
 * @member {Array.<module:model/GuidanceVmwareSizingPlanBeforeActionPriceSets>} priceSets
 */
ServicePlan.prototype['priceSets'] = undefined;

/**
 * @member {module:model/ServicePlanConfig} config
 */
ServicePlan.prototype['config'] = undefined;

/**
 * @member {Array.<module:model/InlineResponse20094Network>} zones
 */
ServicePlan.prototype['zones'] = undefined;

/**
 * @member {module:model/ServicePlanPermissions} permissions
 */
ServicePlan.prototype['permissions'] = undefined;






export default ServicePlan;

