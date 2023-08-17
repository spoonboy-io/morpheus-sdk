/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.1.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The BlueprintKubernetesCreateConfigSpecsInner model module.
 * @module model/BlueprintKubernetesCreateConfigSpecsInner
 * @version 6.1.1
 */
class BlueprintKubernetesCreateConfigSpecsInner {
    /**
     * Constructs a new <code>BlueprintKubernetesCreateConfigSpecsInner</code>.
     * @alias module:model/BlueprintKubernetesCreateConfigSpecsInner
     */
    constructor() { 
        
        BlueprintKubernetesCreateConfigSpecsInner.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BlueprintKubernetesCreateConfigSpecsInner</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BlueprintKubernetesCreateConfigSpecsInner} obj Optional instance to populate.
     * @return {module:model/BlueprintKubernetesCreateConfigSpecsInner} The populated <code>BlueprintKubernetesCreateConfigSpecsInner</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BlueprintKubernetesCreateConfigSpecsInner();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('value')) {
                obj['value'] = ApiClient.convertToType(data['value'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>BlueprintKubernetesCreateConfigSpecsInner</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>BlueprintKubernetesCreateConfigSpecsInner</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['name'] && !(typeof data['name'] === 'string' || data['name'] instanceof String)) {
            throw new Error("Expected the field `name` to be a primitive type in the JSON string but got " + data['name']);
        }

        return true;
    }


}



/**
 * @member {Number} id
 */
BlueprintKubernetesCreateConfigSpecsInner.prototype['id'] = undefined;

/**
 * @member {Number} value
 */
BlueprintKubernetesCreateConfigSpecsInner.prototype['value'] = undefined;

/**
 * @member {String} name
 */
BlueprintKubernetesCreateConfigSpecsInner.prototype['name'] = undefined;






export default BlueprintKubernetesCreateConfigSpecsInner;

