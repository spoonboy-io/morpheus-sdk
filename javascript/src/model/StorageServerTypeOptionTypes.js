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

/**
 * The StorageServerTypeOptionTypes model module.
 * @module model/StorageServerTypeOptionTypes
 * @version 6.2.1
 */
class StorageServerTypeOptionTypes {
    /**
     * Constructs a new <code>StorageServerTypeOptionTypes</code>.
     * @alias module:model/StorageServerTypeOptionTypes
     */
    constructor() { 
        
        StorageServerTypeOptionTypes.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>StorageServerTypeOptionTypes</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/StorageServerTypeOptionTypes} obj Optional instance to populate.
     * @return {module:model/StorageServerTypeOptionTypes} The populated <code>StorageServerTypeOptionTypes</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new StorageServerTypeOptionTypes();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('fieldName')) {
                obj['fieldName'] = ApiClient.convertToType(data['fieldName'], 'String');
            }
            if (data.hasOwnProperty('fieldLabel')) {
                obj['fieldLabel'] = ApiClient.convertToType(data['fieldLabel'], 'String');
            }
            if (data.hasOwnProperty('fieldCode')) {
                obj['fieldCode'] = ApiClient.convertToType(data['fieldCode'], 'String');
            }
            if (data.hasOwnProperty('fieldContext')) {
                obj['fieldContext'] = ApiClient.convertToType(data['fieldContext'], 'String');
            }
            if (data.hasOwnProperty('fieldGroup')) {
                obj['fieldGroup'] = ApiClient.convertToType(data['fieldGroup'], 'String');
            }
            if (data.hasOwnProperty('fieldClass')) {
                obj['fieldClass'] = ApiClient.convertToType(data['fieldClass'], 'String');
            }
            if (data.hasOwnProperty('fieldAddOn')) {
                obj['fieldAddOn'] = ApiClient.convertToType(data['fieldAddOn'], 'String');
            }
            if (data.hasOwnProperty('fieldComponent')) {
                obj['fieldComponent'] = ApiClient.convertToType(data['fieldComponent'], 'String');
            }
            if (data.hasOwnProperty('fieldInput')) {
                obj['fieldInput'] = ApiClient.convertToType(data['fieldInput'], 'String');
            }
            if (data.hasOwnProperty('placeHolder')) {
                obj['placeHolder'] = ApiClient.convertToType(data['placeHolder'], 'String');
            }
            if (data.hasOwnProperty('verifyPattern')) {
                obj['verifyPattern'] = ApiClient.convertToType(data['verifyPattern'], 'String');
            }
            if (data.hasOwnProperty('helpBlock')) {
                obj['helpBlock'] = ApiClient.convertToType(data['helpBlock'], 'String');
            }
            if (data.hasOwnProperty('helpBlockFieldCode')) {
                obj['helpBlockFieldCode'] = ApiClient.convertToType(data['helpBlockFieldCode'], 'String');
            }
            if (data.hasOwnProperty('defaultValue')) {
                obj['defaultValue'] = ApiClient.convertToType(data['defaultValue'], 'String');
            }
            if (data.hasOwnProperty('optionSource')) {
                obj['optionSource'] = ApiClient.convertToType(data['optionSource'], 'String');
            }
            if (data.hasOwnProperty('optionSourceType')) {
                obj['optionSourceType'] = ApiClient.convertToType(data['optionSourceType'], 'String');
            }
            if (data.hasOwnProperty('optionList')) {
                obj['optionList'] = ApiClient.convertToType(data['optionList'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('advanced')) {
                obj['advanced'] = ApiClient.convertToType(data['advanced'], 'Boolean');
            }
            if (data.hasOwnProperty('required')) {
                obj['required'] = ApiClient.convertToType(data['required'], 'Boolean');
            }
            if (data.hasOwnProperty('exportMeta')) {
                obj['exportMeta'] = ApiClient.convertToType(data['exportMeta'], 'Boolean');
            }
            if (data.hasOwnProperty('editable')) {
                obj['editable'] = ApiClient.convertToType(data['editable'], 'Boolean');
            }
            if (data.hasOwnProperty('creatable')) {
                obj['creatable'] = ApiClient.convertToType(data['creatable'], 'Boolean');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], Object);
            }
            if (data.hasOwnProperty('displayOrder')) {
                obj['displayOrder'] = ApiClient.convertToType(data['displayOrder'], 'Number');
            }
            if (data.hasOwnProperty('wrapperClass')) {
                obj['wrapperClass'] = ApiClient.convertToType(data['wrapperClass'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('noBlank')) {
                obj['noBlank'] = ApiClient.convertToType(data['noBlank'], 'Boolean');
            }
            if (data.hasOwnProperty('dependsOnCode')) {
                obj['dependsOnCode'] = ApiClient.convertToType(data['dependsOnCode'], 'String');
            }
            if (data.hasOwnProperty('visibleOnCode')) {
                obj['visibleOnCode'] = ApiClient.convertToType(data['visibleOnCode'], 'String');
            }
            if (data.hasOwnProperty('requireOnCode')) {
                obj['requireOnCode'] = ApiClient.convertToType(data['requireOnCode'], 'String');
            }
            if (data.hasOwnProperty('contextualDefault')) {
                obj['contextualDefault'] = ApiClient.convertToType(data['contextualDefault'], 'Boolean');
            }
            if (data.hasOwnProperty('displayValueOnDetails')) {
                obj['displayValueOnDetails'] = ApiClient.convertToType(data['displayValueOnDetails'], 'Boolean');
            }
            if (data.hasOwnProperty('showOnCreate')) {
                obj['showOnCreate'] = ApiClient.convertToType(data['showOnCreate'], 'Boolean');
            }
            if (data.hasOwnProperty('showOnEdit')) {
                obj['showOnEdit'] = ApiClient.convertToType(data['showOnEdit'], 'Boolean');
            }
            if (data.hasOwnProperty('localCredential')) {
                obj['localCredential'] = ApiClient.convertToType(data['localCredential'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
StorageServerTypeOptionTypes.prototype['id'] = undefined;

/**
 * @member {String} name
 */
StorageServerTypeOptionTypes.prototype['name'] = undefined;

/**
 * @member {String} description
 */
StorageServerTypeOptionTypes.prototype['description'] = undefined;

/**
 * @member {String} code
 */
StorageServerTypeOptionTypes.prototype['code'] = undefined;

/**
 * @member {String} fieldName
 */
StorageServerTypeOptionTypes.prototype['fieldName'] = undefined;

/**
 * @member {String} fieldLabel
 */
StorageServerTypeOptionTypes.prototype['fieldLabel'] = undefined;

/**
 * @member {String} fieldCode
 */
StorageServerTypeOptionTypes.prototype['fieldCode'] = undefined;

/**
 * @member {String} fieldContext
 */
StorageServerTypeOptionTypes.prototype['fieldContext'] = undefined;

/**
 * @member {String} fieldGroup
 */
StorageServerTypeOptionTypes.prototype['fieldGroup'] = undefined;

/**
 * @member {String} fieldClass
 */
StorageServerTypeOptionTypes.prototype['fieldClass'] = undefined;

/**
 * @member {String} fieldAddOn
 */
StorageServerTypeOptionTypes.prototype['fieldAddOn'] = undefined;

/**
 * @member {String} fieldComponent
 */
StorageServerTypeOptionTypes.prototype['fieldComponent'] = undefined;

/**
 * @member {String} fieldInput
 */
StorageServerTypeOptionTypes.prototype['fieldInput'] = undefined;

/**
 * @member {String} placeHolder
 */
StorageServerTypeOptionTypes.prototype['placeHolder'] = undefined;

/**
 * @member {String} verifyPattern
 */
StorageServerTypeOptionTypes.prototype['verifyPattern'] = undefined;

/**
 * @member {String} helpBlock
 */
StorageServerTypeOptionTypes.prototype['helpBlock'] = undefined;

/**
 * @member {String} helpBlockFieldCode
 */
StorageServerTypeOptionTypes.prototype['helpBlockFieldCode'] = undefined;

/**
 * @member {String} defaultValue
 */
StorageServerTypeOptionTypes.prototype['defaultValue'] = undefined;

/**
 * @member {String} optionSource
 */
StorageServerTypeOptionTypes.prototype['optionSource'] = undefined;

/**
 * @member {String} optionSourceType
 */
StorageServerTypeOptionTypes.prototype['optionSourceType'] = undefined;

/**
 * @member {String} optionList
 */
StorageServerTypeOptionTypes.prototype['optionList'] = undefined;

/**
 * @member {String} type
 */
StorageServerTypeOptionTypes.prototype['type'] = undefined;

/**
 * @member {Boolean} advanced
 */
StorageServerTypeOptionTypes.prototype['advanced'] = undefined;

/**
 * @member {Boolean} required
 */
StorageServerTypeOptionTypes.prototype['required'] = undefined;

/**
 * @member {Boolean} exportMeta
 */
StorageServerTypeOptionTypes.prototype['exportMeta'] = undefined;

/**
 * @member {Boolean} editable
 */
StorageServerTypeOptionTypes.prototype['editable'] = undefined;

/**
 * @member {Boolean} creatable
 */
StorageServerTypeOptionTypes.prototype['creatable'] = undefined;

/**
 * @member {Object} config
 */
StorageServerTypeOptionTypes.prototype['config'] = undefined;

/**
 * @member {Number} displayOrder
 */
StorageServerTypeOptionTypes.prototype['displayOrder'] = undefined;

/**
 * @member {String} wrapperClass
 */
StorageServerTypeOptionTypes.prototype['wrapperClass'] = undefined;

/**
 * @member {Boolean} enabled
 */
StorageServerTypeOptionTypes.prototype['enabled'] = undefined;

/**
 * @member {Boolean} noBlank
 */
StorageServerTypeOptionTypes.prototype['noBlank'] = undefined;

/**
 * @member {String} dependsOnCode
 */
StorageServerTypeOptionTypes.prototype['dependsOnCode'] = undefined;

/**
 * @member {String} visibleOnCode
 */
StorageServerTypeOptionTypes.prototype['visibleOnCode'] = undefined;

/**
 * @member {String} requireOnCode
 */
StorageServerTypeOptionTypes.prototype['requireOnCode'] = undefined;

/**
 * @member {Boolean} contextualDefault
 */
StorageServerTypeOptionTypes.prototype['contextualDefault'] = undefined;

/**
 * @member {Boolean} displayValueOnDetails
 */
StorageServerTypeOptionTypes.prototype['displayValueOnDetails'] = undefined;

/**
 * @member {Boolean} showOnCreate
 */
StorageServerTypeOptionTypes.prototype['showOnCreate'] = undefined;

/**
 * @member {Boolean} showOnEdit
 */
StorageServerTypeOptionTypes.prototype['showOnEdit'] = undefined;

/**
 * @member {Boolean} localCredential
 */
StorageServerTypeOptionTypes.prototype['localCredential'] = undefined;






export default StorageServerTypeOptionTypes;

