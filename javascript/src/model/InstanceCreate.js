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
import AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject from './AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject';
import ApiServersIdMakeManagedServerTags from './ApiServersIdMakeManagedServerTags';
import InstanceCreateInstance from './InstanceCreateInstance';
import InstanceCreateNetwork from './InstanceCreateNetwork';
import InstanceCreatePorts from './InstanceCreatePorts';
import InstanceCreateVolume from './InstanceCreateVolume';

/**
 * The InstanceCreate model module.
 * @module model/InstanceCreate
 * @version 6.2.1
 */
class InstanceCreate {
    /**
     * Constructs a new <code>InstanceCreate</code>.
     * @alias module:model/InstanceCreate
     * @param instance {module:model/InstanceCreateInstance} 
     * @param config {module:model/AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject} 
     */
    constructor(instance, config) { 
        
        InstanceCreate.initialize(this, instance, config);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, instance, config) { 
        obj['instance'] = instance;
        obj['config'] = config;
    }

    /**
     * Constructs a <code>InstanceCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceCreate} obj Optional instance to populate.
     * @return {module:model/InstanceCreate} The populated <code>InstanceCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceCreate();

            if (data.hasOwnProperty('instance')) {
                obj['instance'] = InstanceCreateInstance.constructFromObject(data['instance']);
            }
            if (data.hasOwnProperty('zoneId')) {
                obj['zoneId'] = ApiClient.convertToType(data['zoneId'], 'Number');
            }
            if (data.hasOwnProperty('evars')) {
                obj['evars'] = ApiClient.convertToType(data['evars'], [ApiServersIdMakeManagedServerTags]);
            }
            if (data.hasOwnProperty('copies')) {
                obj['copies'] = ApiClient.convertToType(data['copies'], 'Number');
            }
            if (data.hasOwnProperty('layoutSize')) {
                obj['layoutSize'] = ApiClient.convertToType(data['layoutSize'], 'Number');
            }
            if (data.hasOwnProperty('servicePlanOptions')) {
                obj['servicePlanOptions'] = ApiClient.convertToType(data['servicePlanOptions'], Object);
            }
            if (data.hasOwnProperty('securityGroups')) {
                obj['securityGroups'] = ApiClient.convertToType(data['securityGroups'], [Object]);
            }
            if (data.hasOwnProperty('volumes')) {
                obj['volumes'] = ApiClient.convertToType(data['volumes'], [InstanceCreateVolume]);
            }
            if (data.hasOwnProperty('networkInterfaces')) {
                obj['networkInterfaces'] = ApiClient.convertToType(data['networkInterfaces'], [InstanceCreateNetwork]);
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject);
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], [ApiServersIdMakeManagedServerTags]);
            }
            if (data.hasOwnProperty('metadata')) {
                obj['metadata'] = ApiClient.convertToType(data['metadata'], [ApiServersIdMakeManagedServerTags]);
            }
            if (data.hasOwnProperty('ports')) {
                obj['ports'] = ApiClient.convertToType(data['ports'], [InstanceCreatePorts]);
            }
            if (data.hasOwnProperty('taskSetId')) {
                obj['taskSetId'] = ApiClient.convertToType(data['taskSetId'], 'Number');
            }
            if (data.hasOwnProperty('taskSetName')) {
                obj['taskSetName'] = ApiClient.convertToType(data['taskSetName'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/InstanceCreateInstance} instance
 */
InstanceCreate.prototype['instance'] = undefined;

/**
 * The Cloud ID to provision the instance onto.
 * @member {Number} zoneId
 */
InstanceCreate.prototype['zoneId'] = undefined;

/**
 * Environment Variables, an array of objects that have name and value.
 * @member {Array.<module:model/ApiServersIdMakeManagedServerTags>} evars
 */
InstanceCreate.prototype['evars'] = undefined;

/**
 * Number of copies to provision.
 * @member {Number} copies
 * @default 1
 */
InstanceCreate.prototype['copies'] = 1;

/**
 * Apply a multiply factor of containers/vms within the instance.
 * @member {Number} layoutSize
 * @default 1
 */
InstanceCreate.prototype['layoutSize'] = 1;

/**
 * Map of custom options depending on selected service plan.
 * @member {Object} servicePlanOptions
 */
InstanceCreate.prototype['servicePlanOptions'] = undefined;

/**
 * Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to.
 * @member {Array.<Object>} securityGroups
 */
InstanceCreate.prototype['securityGroups'] = undefined;

/**
 * The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of
 * @member {Array.<module:model/InstanceCreateVolume>} volumes
 */
InstanceCreate.prototype['volumes'] = undefined;

/**
 * The networkInterfaces parameter is for network configuration.  The Options API `/api/options/zoneNetworkOptions?zoneId=5&provisionTypeId=10` can be used to see which options are available. 
 * @member {Array.<module:model/InstanceCreateNetwork>} networkInterfaces
 */
InstanceCreate.prototype['networkInterfaces'] = undefined;

/**
 * @member {module:model/AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject} config
 */
InstanceCreate.prototype['config'] = undefined;

/**
 * Array of strings (keywords).
 * @member {Array.<String>} labels
 */
InstanceCreate.prototype['labels'] = undefined;

/**
 * Metadata tags, Array of objects having a name and value.
 * @member {Array.<module:model/ApiServersIdMakeManagedServerTags>} tags
 */
InstanceCreate.prototype['tags'] = undefined;

/**
 * Alias for `tags`.
 * @member {Array.<module:model/ApiServersIdMakeManagedServerTags>} metadata
 */
InstanceCreate.prototype['metadata'] = undefined;

/**
 * The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened. 
 * @member {Array.<module:model/InstanceCreatePorts>} ports
 */
InstanceCreate.prototype['ports'] = undefined;

/**
 * The Workflow ID to execute.
 * @member {Number} taskSetId
 */
InstanceCreate.prototype['taskSetId'] = undefined;

/**
 * The Workflow Name to execute.
 * @member {String} taskSetName
 */
InstanceCreate.prototype['taskSetName'] = undefined;






export default InstanceCreate;

