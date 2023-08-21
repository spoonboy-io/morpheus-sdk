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
 * The InstanceTypeLayoutCreate model module.
 * @module model/InstanceTypeLayoutCreate
 * @version 6.2.1
 */
class InstanceTypeLayoutCreate {
    /**
     * Constructs a new <code>InstanceTypeLayoutCreate</code>.
     * @alias module:model/InstanceTypeLayoutCreate
     * @param name {String} Layout name
     * @param instanceVersion {String} Version of the layout
     * @param provisionTypeCode {String} Provision type code
     */
    constructor(name, instanceVersion, provisionTypeCode) { 
        
        InstanceTypeLayoutCreate.initialize(this, name, instanceVersion, provisionTypeCode);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name, instanceVersion, provisionTypeCode) { 
        obj['name'] = name;
        obj['instanceVersion'] = instanceVersion;
        obj['provisionTypeCode'] = provisionTypeCode;
    }

    /**
     * Constructs a <code>InstanceTypeLayoutCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceTypeLayoutCreate} obj Optional instance to populate.
     * @return {module:model/InstanceTypeLayoutCreate} The populated <code>InstanceTypeLayoutCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceTypeLayoutCreate();

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
InstanceTypeLayoutCreate.prototype['name'] = undefined;

/**
 * @member {Array.<String>} labels
 */
InstanceTypeLayoutCreate.prototype['labels'] = undefined;

/**
 * Version of the layout
 * @member {String} instanceVersion
 */
InstanceTypeLayoutCreate.prototype['instanceVersion'] = undefined;

/**
 * Layout description
 * @member {String} description
 */
InstanceTypeLayoutCreate.prototype['description'] = undefined;

/**
 * Can be used to enable / disable the creatability of the layout.
 * @member {Boolean} creatable
 * @default true
 */
InstanceTypeLayoutCreate.prototype['creatable'] = true;

/**
 * Provision type code
 * @member {String} provisionTypeCode
 */
InstanceTypeLayoutCreate.prototype['provisionTypeCode'] = undefined;

/**
 * Memory requirement in megabytes
 * @member {String} memoryRequirement
 */
InstanceTypeLayoutCreate.prototype['memoryRequirement'] = undefined;

/**
 * Can be used to enable / disable the horizontal scaling.
 * @member {Boolean} hasAutoScale
 * @default false
 */
InstanceTypeLayoutCreate.prototype['hasAutoScale'] = false;

/**
 * Can be used to enable / disable the supports convert to managed.
 * @member {Boolean} supportsConvertToManaged
 * @default false
 */
InstanceTypeLayoutCreate.prototype['supportsConvertToManaged'] = false;

/**
 * Array of layout node type IDs
 * @member {Array.<Number>} containerTypes
 */
InstanceTypeLayoutCreate.prototype['containerTypes'] = undefined;

/**
 * Array of layout option type IDs
 * @member {Array.<Number>} optionTypes
 */
InstanceTypeLayoutCreate.prototype['optionTypes'] = undefined;

/**
 * Array of layout spec template IDs
 * @member {Array.<Number>} specTemplates
 */
InstanceTypeLayoutCreate.prototype['specTemplates'] = undefined;

/**
 * The environmentVariables parameter is array of env objects
 * @member {Array.<module:model/ClusterLayoutCreateEnvironmentVariables>} environmentVariables
 */
InstanceTypeLayoutCreate.prototype['environmentVariables'] = undefined;

/**
 * Array of price set objects
 * @member {Array.<module:model/InstanceTypeCreatePriceSets>} priceSets
 */
InstanceTypeLayoutCreate.prototype['priceSets'] = undefined;

/**
 * @member {module:model/ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions} permissions
 */
InstanceTypeLayoutCreate.prototype['permissions'] = undefined;






export default InstanceTypeLayoutCreate;
