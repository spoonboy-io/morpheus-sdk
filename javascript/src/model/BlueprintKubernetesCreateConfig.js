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
import BlueprintKubernetesCreateConfigSpecs from './BlueprintKubernetesCreateConfigSpecs';

/**
 * The BlueprintKubernetesCreateConfig model module.
 * @module model/BlueprintKubernetesCreateConfig
 * @version 6.2.1
 */
class BlueprintKubernetesCreateConfig {
    /**
     * Constructs a new <code>BlueprintKubernetesCreateConfig</code>.
     * @alias module:model/BlueprintKubernetesCreateConfig
     */
    constructor() { 
        
        BlueprintKubernetesCreateConfig.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BlueprintKubernetesCreateConfig</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BlueprintKubernetesCreateConfig} obj Optional instance to populate.
     * @return {module:model/BlueprintKubernetesCreateConfig} The populated <code>BlueprintKubernetesCreateConfig</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BlueprintKubernetesCreateConfig();

            if (data.hasOwnProperty('specs')) {
                obj['specs'] = ApiClient.convertToType(data['specs'], [BlueprintKubernetesCreateConfigSpecs]);
            }
        }
        return obj;
    }


}

/**
 * Array of Kubernetes specs in Morpheus
 * @member {Array.<module:model/BlueprintKubernetesCreateConfigSpecs>} specs
 */
BlueprintKubernetesCreateConfig.prototype['specs'] = undefined;






export default BlueprintKubernetesCreateConfig;

