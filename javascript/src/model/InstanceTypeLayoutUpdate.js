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
import ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions from './ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions';
import ClusterLayoutCreateEnvironmentVariables from './ClusterLayoutCreateEnvironmentVariables';
import InstanceTypeCreatePriceSets from './InstanceTypeCreatePriceSets';

/**
 * The InstanceTypeLayoutUpdate model module.
 * @module model/InstanceTypeLayoutUpdate
 * @version 6.2.1
 */
class InstanceTypeLayoutUpdate {
    /**
     * Constructs a new <code>InstanceTypeLayoutUpdate</code>.
     * @alias module:model/InstanceTypeLayoutUpdate
     */
    constructor() { 
        
        InstanceTypeLayoutUpdate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InstanceTypeLayoutUpdate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceTypeLayoutUpdate} obj Optional instance to populate.
     * @return {module:model/InstanceTypeLayoutUpdate} The populated <code>InstanceTypeLayoutUpdate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceTypeLayoutUpdate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('instanceVersion')) {
                obj['instanceVersion'] = ApiClient.convertToType(data['instanceVersion'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('creatable')) {
                obj['creatable'] = ApiClient.convertToType(data['creatable'], 'Boolean');
            }
            if (data.hasOwnProperty('provisionTypeCode')) {
                obj['provisionTypeCode'] = ApiClient.convertToType(data['provisionTypeCode'], 'String');
            }
            if (data.hasOwnProperty('memoryRequirement')) {
                obj['memoryRequirement'] = ApiClient.convertToType(data['memoryRequirement'], 'String');
            }
            if (data.hasOwnProperty('hasAutoScale')) {
                obj['hasAutoScale'] = ApiClient.convertToType(data['hasAutoScale'], 'Boolean');
            }
            if (data.hasOwnProperty('supportsConvertToManaged')) {
                obj['supportsConvertToManaged'] = ApiClient.convertToType(data['supportsConvertToManaged'], 'Boolean');
            }
            if (data.hasOwnProperty('containerTypes')) {
                obj['containerTypes'] = ApiClient.convertToType(data['containerTypes'], ['Number']);
            }
            if (data.hasOwnProperty('optionTypes')) {
                obj['optionTypes'] = ApiClient.convertToType(data['optionTypes'], ['Number']);
            }
            if (data.hasOwnProperty('specTemplates')) {
                obj['specTemplates'] = ApiClient.convertToType(data['specTemplates'], ['Number']);
            }
            if (data.hasOwnProperty('environmentVariables')) {
                obj['environmentVariables'] = ApiClient.convertToType(data['environmentVariables'], [ClusterLayoutCreateEnvironmentVariables]);
            }
            if (data.hasOwnProperty('priceSets')) {
                obj['priceSets'] = ApiClient.convertToType(data['priceSets'], [InstanceTypeCreatePriceSets]);
            }
            if (data.hasOwnProperty('permissions')) {
                obj['permissions'] = ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions.constructFromObject(data['permissions']);
            }
        }
        return obj;
    }


}

/**
 * Layout name
 * @member {String} name
 */
InstanceTypeLayoutUpdate.prototype['name'] = undefined;

/**
 * @member {Array.<String>} labels
 */
InstanceTypeLayoutUpdate.prototype['labels'] = undefined;

/**
 * Version of the layout
 * @member {String} instanceVersion
 */
InstanceTypeLayoutUpdate.prototype['instanceVersion'] = undefined;

/**
 * Layout description
 * @member {String} description
 */
InstanceTypeLayoutUpdate.prototype['description'] = undefined;

/**
 * Can be used to enable / disable the creatability of the layout.
 * @member {Boolean} creatable
 * @default true
 */
InstanceTypeLayoutUpdate.prototype['creatable'] = true;

/**
 * Provision type code
 * @member {String} provisionTypeCode
 */
InstanceTypeLayoutUpdate.prototype['provisionTypeCode'] = undefined;

/**
 * Memory requirement in megabytes
 * @member {String} memoryRequirement
 */
InstanceTypeLayoutUpdate.prototype['memoryRequirement'] = undefined;

/**
 * Can be used to enable / disable the horizontal scaling.
 * @member {Boolean} hasAutoScale
 * @default false
 */
InstanceTypeLayoutUpdate.prototype['hasAutoScale'] = false;

/**
 * Can be used to enable / disable the supports convert to managed.
 * @member {Boolean} supportsConvertToManaged
 * @default false
 */
InstanceTypeLayoutUpdate.prototype['supportsConvertToManaged'] = false;

/**
 * Array of layout node type IDs
 * @member {Array.<Number>} containerTypes
 */
InstanceTypeLayoutUpdate.prototype['containerTypes'] = undefined;

/**
 * Array of layout option type IDs
 * @member {Array.<Number>} optionTypes
 */
InstanceTypeLayoutUpdate.prototype['optionTypes'] = undefined;

/**
 * Array of layout spec template IDs
 * @member {Array.<Number>} specTemplates
 */
InstanceTypeLayoutUpdate.prototype['specTemplates'] = undefined;

/**
 * The environmentVariables parameter is array of env objects
 * @member {Array.<module:model/ClusterLayoutCreateEnvironmentVariables>} environmentVariables
 */
InstanceTypeLayoutUpdate.prototype['environmentVariables'] = undefined;

/**
 * Array of price set objects
 * @member {Array.<module:model/InstanceTypeCreatePriceSets>} priceSets
 */
InstanceTypeLayoutUpdate.prototype['priceSets'] = undefined;

/**
 * @member {module:model/ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions} permissions
 */
InstanceTypeLayoutUpdate.prototype['permissions'] = undefined;






export default InstanceTypeLayoutUpdate;

