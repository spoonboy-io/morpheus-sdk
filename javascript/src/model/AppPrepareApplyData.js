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
import AppPrepareApplyDataGroup from './AppPrepareApplyDataGroup';
import AppPrepareApplyDataTerraform from './AppPrepareApplyDataTerraform';

/**
 * The AppPrepareApplyData model module.
 * @module model/AppPrepareApplyData
 * @version 6.1.1
 */
class AppPrepareApplyData {
    /**
     * Constructs a new <code>AppPrepareApplyData</code>.
     * @alias module:model/AppPrepareApplyData
     */
    constructor() { 
        
        AppPrepareApplyData.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>AppPrepareApplyData</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AppPrepareApplyData} obj Optional instance to populate.
     * @return {module:model/AppPrepareApplyData} The populated <code>AppPrepareApplyData</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AppPrepareApplyData();

            if (data.hasOwnProperty('image')) {
                obj['image'] = ApiClient.convertToType(data['image'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('autoValidate')) {
                obj['autoValidate'] = ApiClient.convertToType(data['autoValidate'], 'Boolean');
            }
            if (data.hasOwnProperty('terraform')) {
                obj['terraform'] = AppPrepareApplyDataTerraform.constructFromObject(data['terraform']);
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], Object);
            }
            if (data.hasOwnProperty('blueprintName')) {
                obj['blueprintName'] = ApiClient.convertToType(data['blueprintName'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('templateId')) {
                obj['templateId'] = ApiClient.convertToType(data['templateId'], 'Number');
            }
            if (data.hasOwnProperty('blueprintId')) {
                obj['blueprintId'] = ApiClient.convertToType(data['blueprintId'], 'Number');
            }
            if (data.hasOwnProperty('group')) {
                obj['group'] = AppPrepareApplyDataGroup.constructFromObject(data['group']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>AppPrepareApplyData</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>AppPrepareApplyData</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['image'] && !(typeof data['image'] === 'string' || data['image'] instanceof String)) {
            throw new Error("Expected the field `image` to be a primitive type in the JSON string but got " + data['image']);
        }
        // ensure the json data is a string
        if (data['name'] && !(typeof data['name'] === 'string' || data['name'] instanceof String)) {
            throw new Error("Expected the field `name` to be a primitive type in the JSON string but got " + data['name']);
        }
        // validate the optional field `terraform`
        if (data['terraform']) { // data not null
          AppPrepareApplyDataTerraform.validateJSON(data['terraform']);
        }
        // ensure the json data is a string
        if (data['type'] && !(typeof data['type'] === 'string' || data['type'] instanceof String)) {
            throw new Error("Expected the field `type` to be a primitive type in the JSON string but got " + data['type']);
        }
        // ensure the json data is a string
        if (data['blueprintName'] && !(typeof data['blueprintName'] === 'string' || data['blueprintName'] instanceof String)) {
            throw new Error("Expected the field `blueprintName` to be a primitive type in the JSON string but got " + data['blueprintName']);
        }
        // ensure the json data is a string
        if (data['description'] && !(typeof data['description'] === 'string' || data['description'] instanceof String)) {
            throw new Error("Expected the field `description` to be a primitive type in the JSON string but got " + data['description']);
        }
        // validate the optional field `group`
        if (data['group']) { // data not null
          AppPrepareApplyDataGroup.validateJSON(data['group']);
        }

        return true;
    }


}



/**
 * @member {String} image
 */
AppPrepareApplyData.prototype['image'] = undefined;

/**
 * @member {String} name
 */
AppPrepareApplyData.prototype['name'] = undefined;

/**
 * @member {Boolean} autoValidate
 */
AppPrepareApplyData.prototype['autoValidate'] = undefined;

/**
 * @member {module:model/AppPrepareApplyDataTerraform} terraform
 */
AppPrepareApplyData.prototype['terraform'] = undefined;

/**
 * @member {String} type
 */
AppPrepareApplyData.prototype['type'] = undefined;

/**
 * @member {Object} config
 */
AppPrepareApplyData.prototype['config'] = undefined;

/**
 * @member {String} blueprintName
 */
AppPrepareApplyData.prototype['blueprintName'] = undefined;

/**
 * @member {String} description
 */
AppPrepareApplyData.prototype['description'] = undefined;

/**
 * @member {Number} templateId
 */
AppPrepareApplyData.prototype['templateId'] = undefined;

/**
 * @member {Number} blueprintId
 */
AppPrepareApplyData.prototype['blueprintId'] = undefined;

/**
 * @member {module:model/AppPrepareApplyDataGroup} group
 */
AppPrepareApplyData.prototype['group'] = undefined;






export default AppPrepareApplyData;
