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
 * The GuidanceVmwareSizingPlanBeforeActionProvisionType model module.
 * @module model/GuidanceVmwareSizingPlanBeforeActionProvisionType
 * @version 6.2.1
 */
class GuidanceVmwareSizingPlanBeforeActionProvisionType {
    /**
     * Constructs a new <code>GuidanceVmwareSizingPlanBeforeActionProvisionType</code>.
     * @alias module:model/GuidanceVmwareSizingPlanBeforeActionProvisionType
     */
    constructor() { 
        
        GuidanceVmwareSizingPlanBeforeActionProvisionType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>GuidanceVmwareSizingPlanBeforeActionProvisionType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/GuidanceVmwareSizingPlanBeforeActionProvisionType} obj Optional instance to populate.
     * @return {module:model/GuidanceVmwareSizingPlanBeforeActionProvisionType} The populated <code>GuidanceVmwareSizingPlanBeforeActionProvisionType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new GuidanceVmwareSizingPlanBeforeActionProvisionType();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('rootDiskCustomizable')) {
                obj['rootDiskCustomizable'] = ApiClient.convertToType(data['rootDiskCustomizable'], 'Boolean');
            }
            if (data.hasOwnProperty('addVolumes')) {
                obj['addVolumes'] = ApiClient.convertToType(data['addVolumes'], 'Boolean');
            }
            if (data.hasOwnProperty('customizeVolume')) {
                obj['customizeVolume'] = ApiClient.convertToType(data['customizeVolume'], 'Boolean');
            }
            if (data.hasOwnProperty('hasConfigurableCpuSockets')) {
                obj['hasConfigurableCpuSockets'] = ApiClient.convertToType(data['hasConfigurableCpuSockets'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
GuidanceVmwareSizingPlanBeforeActionProvisionType.prototype['id'] = undefined;

/**
 * @member {String} name
 */
GuidanceVmwareSizingPlanBeforeActionProvisionType.prototype['name'] = undefined;

/**
 * @member {String} code
 */
GuidanceVmwareSizingPlanBeforeActionProvisionType.prototype['code'] = undefined;

/**
 * @member {Boolean} rootDiskCustomizable
 */
GuidanceVmwareSizingPlanBeforeActionProvisionType.prototype['rootDiskCustomizable'] = undefined;

/**
 * @member {Boolean} addVolumes
 */
GuidanceVmwareSizingPlanBeforeActionProvisionType.prototype['addVolumes'] = undefined;

/**
 * @member {Boolean} customizeVolume
 */
GuidanceVmwareSizingPlanBeforeActionProvisionType.prototype['customizeVolume'] = undefined;

/**
 * @member {Boolean} hasConfigurableCpuSockets
 */
GuidanceVmwareSizingPlanBeforeActionProvisionType.prototype['hasConfigurableCpuSockets'] = undefined;






export default GuidanceVmwareSizingPlanBeforeActionProvisionType;
