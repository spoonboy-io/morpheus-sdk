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
import OneOfblueprintARMCreateSuccessblueprintCFTCreateSuccessblueprintHelmCreateSuccessblueprintKubernetesCreateSuccessblueprintMorpheusCreateSuccessblueprintTerraformCreateSuccess from './OneOfblueprintARMCreateSuccessblueprintCFTCreateSuccessblueprintHelmCreateSuccessblueprintKubernetesCreateSuccessblueprintMorpheusCreateSuccessblueprintTerraformCreateSuccess';

/**
 * The BlueprintCreateSuccess model module.
 * @module model/BlueprintCreateSuccess
 * @version 6.2.1
 */
class BlueprintCreateSuccess {
    /**
     * Constructs a new <code>BlueprintCreateSuccess</code>.
     * @alias module:model/BlueprintCreateSuccess
     */
    constructor() { 
        
        BlueprintCreateSuccess.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BlueprintCreateSuccess</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BlueprintCreateSuccess} obj Optional instance to populate.
     * @return {module:model/BlueprintCreateSuccess} The populated <code>BlueprintCreateSuccess</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BlueprintCreateSuccess();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], OneOfblueprintARMCreateSuccessblueprintCFTCreateSuccessblueprintHelmCreateSuccessblueprintKubernetesCreateSuccessblueprintMorpheusCreateSuccessblueprintTerraformCreateSuccess);
            }
        }
        return obj;
    }


}

/**
 * Blueprint ID
 * @member {Number} id
 */
BlueprintCreateSuccess.prototype['id'] = undefined;

/**
 * A name for the blueprint
 * @member {String} name
 */
BlueprintCreateSuccess.prototype['name'] = undefined;

/**
 * A description for the blueprint
 * @member {String} description
 */
BlueprintCreateSuccess.prototype['description'] = undefined;

/**
 * @member {Array.<String>} labels
 */
BlueprintCreateSuccess.prototype['labels'] = undefined;

/**
 * Category
 * @member {String} category
 */
BlueprintCreateSuccess.prototype['category'] = undefined;

/**
 * @member {module:model/OneOfblueprintARMCreateSuccessblueprintCFTCreateSuccessblueprintHelmCreateSuccessblueprintKubernetesCreateSuccessblueprintMorpheusCreateSuccessblueprintTerraformCreateSuccess} config
 */
BlueprintCreateSuccess.prototype['config'] = undefined;






export default BlueprintCreateSuccess;

