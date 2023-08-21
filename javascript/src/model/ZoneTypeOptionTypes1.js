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
import ZoneTypeConfig from './ZoneTypeConfig';

/**
 * The ZoneTypeOptionTypes1 model module.
 * @module model/ZoneTypeOptionTypes1
 * @version 6.2.1
 */
class ZoneTypeOptionTypes1 {
    /**
     * Constructs a new <code>ZoneTypeOptionTypes1</code>.
     * @alias module:model/ZoneTypeOptionTypes1
     */
    constructor() { 
        
        ZoneTypeOptionTypes1.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ZoneTypeOptionTypes1</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ZoneTypeOptionTypes1} obj Optional instance to populate.
     * @return {module:model/ZoneTypeOptionTypes1} The populated <code>ZoneTypeOptionTypes1</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ZoneTypeOptionTypes1();

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
                obj['config'] = ZoneTypeConfig.constructFromObject(data['config']);
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
ZoneTypeOptionTypes1.prototype['id'] = undefined;

/**
 * @member {String} name
 */
ZoneTypeOptionTypes1.prototype['name'] = undefined;

/**
 * @member {String} description
 */
ZoneTypeOptionTypes1.prototype['description'] = undefined;

/**
 * @member {String} code
 */
ZoneTypeOptionTypes1.prototype['code'] = undefined;

/**
 * @member {String} fieldName
 */
ZoneTypeOptionTypes1.prototype['fieldName'] = undefined;

/**
 * @member {String} fieldLabel
 */
ZoneTypeOptionTypes1.prototype['fieldLabel'] = undefined;

/**
 * @member {String} fieldCode
 */
ZoneTypeOptionTypes1.prototype['fieldCode'] = undefined;

/**
 * @member {String} fieldContext
 */
ZoneTypeOptionTypes1.prototype['fieldContext'] = undefined;

/**
 * @member {String} fieldGroup
 */
ZoneTypeOptionTypes1.prototype['fieldGroup'] = undefined;

/**
 * @member {String} fieldClass
 */
ZoneTypeOptionTypes1.prototype['fieldClass'] = undefined;

/**
 * @member {String} fieldAddOn
 */
ZoneTypeOptionTypes1.prototype['fieldAddOn'] = undefined;

/**
 * @member {String} fieldComponent
 */
ZoneTypeOptionTypes1.prototype['fieldComponent'] = undefined;

/**
 * @member {String} fieldInput
 */
ZoneTypeOptionTypes1.prototype['fieldInput'] = undefined;

/**
 * @member {String} placeHolder
 */
ZoneTypeOptionTypes1.prototype['placeHolder'] = undefined;

/**
 * @member {String} verifyPattern
 */
ZoneTypeOptionTypes1.prototype['verifyPattern'] = undefined;

/**
 * @member {String} helpBlock
 */
ZoneTypeOptionTypes1.prototype['helpBlock'] = undefined;

/**
 * @member {String} helpBlockFieldCode
 */
ZoneTypeOptionTypes1.prototype['helpBlockFieldCode'] = undefined;

/**
 * @member {String} defaultValue
 */
ZoneTypeOptionTypes1.prototype['defaultValue'] = undefined;

/**
 * @member {String} optionSource
 */
ZoneTypeOptionTypes1.prototype['optionSource'] = undefined;

/**
 * @member {String} optionSourceType
 */
ZoneTypeOptionTypes1.prototype['optionSourceType'] = undefined;

/**
 * @member {String} optionList
 */
ZoneTypeOptionTypes1.prototype['optionList'] = undefined;

/**
 * @member {String} type
 */
ZoneTypeOptionTypes1.prototype['type'] = undefined;

/**
 * @member {Boolean} advanced
 */
ZoneTypeOptionTypes1.prototype['advanced'] = undefined;

/**
 * @member {Boolean} required
 */
ZoneTypeOptionTypes1.prototype['required'] = undefined;

/**
 * @member {Boolean} exportMeta
 */
ZoneTypeOptionTypes1.prototype['exportMeta'] = undefined;

/**
 * @member {Boolean} editable
 */
ZoneTypeOptionTypes1.prototype['editable'] = undefined;

/**
 * @member {Boolean} creatable
 */
ZoneTypeOptionTypes1.prototype['creatable'] = undefined;

/**
 * @member {module:model/ZoneTypeConfig} config
 */
ZoneTypeOptionTypes1.prototype['config'] = undefined;

/**
 * @member {Number} displayOrder
 */
ZoneTypeOptionTypes1.prototype['displayOrder'] = undefined;

/**
 * @member {String} wrapperClass
 */
ZoneTypeOptionTypes1.prototype['wrapperClass'] = undefined;

/**
 * @member {Boolean} enabled
 */
ZoneTypeOptionTypes1.prototype['enabled'] = undefined;

/**
 * @member {Boolean} noBlank
 */
ZoneTypeOptionTypes1.prototype['noBlank'] = undefined;

/**
 * @member {String} dependsOnCode
 */
ZoneTypeOptionTypes1.prototype['dependsOnCode'] = undefined;

/**
 * @member {String} visibleOnCode
 */
ZoneTypeOptionTypes1.prototype['visibleOnCode'] = undefined;

/**
 * @member {String} requireOnCode
 */
ZoneTypeOptionTypes1.prototype['requireOnCode'] = undefined;

/**
 * @member {Boolean} contextualDefault
 */
ZoneTypeOptionTypes1.prototype['contextualDefault'] = undefined;

/**
 * @member {Boolean} displayValueOnDetails
 */
ZoneTypeOptionTypes1.prototype['displayValueOnDetails'] = undefined;

/**
 * @member {Boolean} showOnCreate
 */
ZoneTypeOptionTypes1.prototype['showOnCreate'] = undefined;

/**
 * @member {Boolean} showOnEdit
 */
ZoneTypeOptionTypes1.prototype['showOnEdit'] = undefined;

/**
 * @member {Boolean} localCredential
 */
ZoneTypeOptionTypes1.prototype['localCredential'] = undefined;






export default ZoneTypeOptionTypes1;
