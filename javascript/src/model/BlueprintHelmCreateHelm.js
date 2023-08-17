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
import BlueprintHelmCreateHelmGit from './BlueprintHelmCreateHelmGit';

/**
 * The BlueprintHelmCreateHelm model module.
 * @module model/BlueprintHelmCreateHelm
 * @version 6.1.1
 */
class BlueprintHelmCreateHelm {
    /**
     * Constructs a new <code>BlueprintHelmCreateHelm</code>.
     * @alias module:model/BlueprintHelmCreateHelm
     * @param configType {module:model/BlueprintHelmCreateHelm.ConfigTypeEnum} Configuration Type
     */
    constructor(configType) { 
        
        BlueprintHelmCreateHelm.initialize(this, configType);
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
     * Constructs a <code>BlueprintHelmCreateHelm</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BlueprintHelmCreateHelm} obj Optional instance to populate.
     * @return {module:model/BlueprintHelmCreateHelm} The populated <code>BlueprintHelmCreateHelm</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BlueprintHelmCreateHelm();

            if (data.hasOwnProperty('configType')) {
                obj['configType'] = ApiClient.convertToType(data['configType'], 'String');
            }
            if (data.hasOwnProperty('git')) {
                obj['git'] = BlueprintHelmCreateHelmGit.constructFromObject(data['git']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>BlueprintHelmCreateHelm</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>BlueprintHelmCreateHelm</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of BlueprintHelmCreateHelm.RequiredProperties) {
            if (!data[property]) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // ensure the json data is a string
        if (data['configType'] && !(typeof data['configType'] === 'string' || data['configType'] instanceof String)) {
            throw new Error("Expected the field `configType` to be a primitive type in the JSON string but got " + data['configType']);
        }
        // validate the optional field `git`
        if (data['git']) { // data not null
          BlueprintHelmCreateHelmGit.validateJSON(data['git']);
        }

        return true;
    }


}

BlueprintHelmCreateHelm.RequiredProperties = ["configType"];

/**
 * Configuration Type
 * @member {module:model/BlueprintHelmCreateHelm.ConfigTypeEnum} configType
 */
BlueprintHelmCreateHelm.prototype['configType'] = undefined;

/**
 * @member {module:model/BlueprintHelmCreateHelmGit} git
 */
BlueprintHelmCreateHelm.prototype['git'] = undefined;





/**
 * Allowed values for the <code>configType</code> property.
 * @enum {String}
 * @readonly
 */
BlueprintHelmCreateHelm['ConfigTypeEnum'] = {

    /**
     * value: "git"
     * @const
     */
    "git": "git"
};



export default BlueprintHelmCreateHelm;
