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
import BlueprintKubernetesCreateKubernetesGit from './BlueprintKubernetesCreateKubernetesGit';

/**
 * The BlueprintKubernetesCreateKubernetes model module.
 * @module model/BlueprintKubernetesCreateKubernetes
 * @version 6.1.1
 */
class BlueprintKubernetesCreateKubernetes {
    /**
     * Constructs a new <code>BlueprintKubernetesCreateKubernetes</code>.
     * @alias module:model/BlueprintKubernetesCreateKubernetes
     * @param configType {module:model/BlueprintKubernetesCreateKubernetes.ConfigTypeEnum} Configuration Type
     */
    constructor(configType) { 
        
        BlueprintKubernetesCreateKubernetes.initialize(this, configType);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, configType) { 
        obj['configType'] = configType;
    }

    /**
     * Constructs a <code>BlueprintKubernetesCreateKubernetes</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BlueprintKubernetesCreateKubernetes} obj Optional instance to populate.
     * @return {module:model/BlueprintKubernetesCreateKubernetes} The populated <code>BlueprintKubernetesCreateKubernetes</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BlueprintKubernetesCreateKubernetes();

            if (data.hasOwnProperty('configType')) {
                obj['configType'] = ApiClient.convertToType(data['configType'], 'String');
            }
            if (data.hasOwnProperty('yaml')) {
                obj['yaml'] = ApiClient.convertToType(data['yaml'], 'String');
            }
            if (data.hasOwnProperty('git')) {
                obj['git'] = BlueprintKubernetesCreateKubernetesGit.constructFromObject(data['git']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>BlueprintKubernetesCreateKubernetes</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>BlueprintKubernetesCreateKubernetes</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of BlueprintKubernetesCreateKubernetes.RequiredProperties) {
            if (!data[property]) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // ensure the json data is a string
        if (data['configType'] && !(typeof data['configType'] === 'string' || data['configType'] instanceof String)) {
            throw new Error("Expected the field `configType` to be a primitive type in the JSON string but got " + data['configType']);
        }
        // ensure the json data is a string
        if (data['yaml'] && !(typeof data['yaml'] === 'string' || data['yaml'] instanceof String)) {
            throw new Error("Expected the field `yaml` to be a primitive type in the JSON string but got " + data['yaml']);
        }
        // validate the optional field `git`
        if (data['git']) { // data not null
          BlueprintKubernetesCreateKubernetesGit.validateJSON(data['git']);
        }

        return true;
    }


}

BlueprintKubernetesCreateKubernetes.RequiredProperties = ["configType"];

/**
 * Configuration Type
 * @member {module:model/BlueprintKubernetesCreateKubernetes.ConfigTypeEnum} configType
 */
BlueprintKubernetesCreateKubernetes.prototype['configType'] = undefined;

/**
 * Kubernetes Spec in YAML
 * @member {String} yaml
 */
BlueprintKubernetesCreateKubernetes.prototype['yaml'] = undefined;

/**
 * @member {module:model/BlueprintKubernetesCreateKubernetesGit} git
 */
BlueprintKubernetesCreateKubernetes.prototype['git'] = undefined;





/**
 * Allowed values for the <code>configType</code> property.
 * @enum {String}
 * @readonly
 */
BlueprintKubernetesCreateKubernetes['ConfigTypeEnum'] = {

    /**
     * value: "yaml"
     * @const
     */
    "yaml": "yaml",

    /**
     * value: "spec"
     * @const
     */
    "spec": "spec",

    /**
     * value: "git"
     * @const
     */
    "git": "git"
};



export default BlueprintKubernetesCreateKubernetes;

