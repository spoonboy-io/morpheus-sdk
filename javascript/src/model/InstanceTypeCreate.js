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
import ClusterLayoutCreateEnvironmentVariables from './ClusterLayoutCreateEnvironmentVariables';
import InstanceTypeCreatePriceSets from './InstanceTypeCreatePriceSets';

/**
 * The InstanceTypeCreate model module.
 * @module model/InstanceTypeCreate
 * @version 6.2.1
 */
class InstanceTypeCreate {
    /**
     * Constructs a new <code>InstanceTypeCreate</code>.
     * @alias module:model/InstanceTypeCreate
     * @param name {String} Instance type name
     */
    constructor(name) { 
        
        InstanceTypeCreate.initialize(this, name);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name) { 
        obj['name'] = name;
    }

    /**
     * Constructs a <code>InstanceTypeCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceTypeCreate} obj Optional instance to populate.
     * @return {module:model/InstanceTypeCreate} The populated <code>InstanceTypeCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceTypeCreate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('featured')) {
                obj['featured'] = ApiClient.convertToType(data['featured'], 'Boolean');
            }
            if (data.hasOwnProperty('hasSettings')) {
                obj['hasSettings'] = ApiClient.convertToType(data['hasSettings'], 'Boolean');
            }
            if (data.hasOwnProperty('hasAutoScale')) {
                obj['hasAutoScale'] = ApiClient.convertToType(data['hasAutoScale'], 'Boolean');
            }
            if (data.hasOwnProperty('hasDeployment')) {
                obj['hasDeployment'] = ApiClient.convertToType(data['hasDeployment'], 'Boolean');
            }
            if (data.hasOwnProperty('environmentPrefix')) {
                obj['environmentPrefix'] = ApiClient.convertToType(data['environmentPrefix'], 'String');
            }
            if (data.hasOwnProperty('environmentVariables')) {
                obj['environmentVariables'] = ApiClient.convertToType(data['environmentVariables'], [ClusterLayoutCreateEnvironmentVariables]);
            }
            if (data.hasOwnProperty('priceSets')) {
                obj['priceSets'] = ApiClient.convertToType(data['priceSets'], [InstanceTypeCreatePriceSets]);
            }
            if (data.hasOwnProperty('optionTypes')) {
                obj['optionTypes'] = ApiClient.convertToType(data['optionTypes'], ['Number']);
            }
        }
        return obj;
    }


}

/**
 * Instance type name
 * @member {String} name
 */
InstanceTypeCreate.prototype['name'] = undefined;

/**
 * Array of label strings, can be used for filtering.
 * @member {Array.<String>} labels
 */
InstanceTypeCreate.prototype['labels'] = undefined;

/**
 * Instance type description
 * @member {String} description
 */
InstanceTypeCreate.prototype['description'] = undefined;

/**
 * Instance type code
 * @member {String} code
 */
InstanceTypeCreate.prototype['code'] = undefined;

/**
 * Category
 * @member {String} category
 */
InstanceTypeCreate.prototype['category'] = undefined;

/**
 * Visibility
 * @member {module:model/InstanceTypeCreate.VisibilityEnum} visibility
 * @default 'private'
 */
InstanceTypeCreate.prototype['visibility'] = 'private';

/**
 * Featured
 * @member {Boolean} featured
 */
InstanceTypeCreate.prototype['featured'] = undefined;

/**
 * Enable Settings
 * @member {Boolean} hasSettings
 */
InstanceTypeCreate.prototype['hasSettings'] = undefined;

/**
 * Enable Scaling (Horizontal)
 * @member {Boolean} hasAutoScale
 */
InstanceTypeCreate.prototype['hasAutoScale'] = undefined;

/**
 * Supports Deployments
 * @member {Boolean} hasDeployment
 */
InstanceTypeCreate.prototype['hasDeployment'] = undefined;

/**
 * Environment Prefix, can be used to make exported evars unique.
 * @member {String} environmentPrefix
 */
InstanceTypeCreate.prototype['environmentPrefix'] = undefined;

/**
 * Array of instance type env variables.
 * @member {Array.<module:model/ClusterLayoutCreateEnvironmentVariables>} environmentVariables
 */
InstanceTypeCreate.prototype['environmentVariables'] = undefined;

/**
 * Array of price set objects
 * @member {Array.<module:model/InstanceTypeCreatePriceSets>} priceSets
 */
InstanceTypeCreate.prototype['priceSets'] = undefined;

/**
 * Array of instance type option type IDs
 * @member {Array.<Number>} optionTypes
 */
InstanceTypeCreate.prototype['optionTypes'] = undefined;





/**
 * Allowed values for the <code>visibility</code> property.
 * @enum {String}
 * @readonly
 */
InstanceTypeCreate['VisibilityEnum'] = {

    /**
     * value: "private"
     * @const
     */
    "private": "private",

    /**
     * value: "public"
     * @const
     */
    "public": "public"
};



export default InstanceTypeCreate;

