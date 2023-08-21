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
 * The FileTemplateUpdate model module.
 * @module model/FileTemplateUpdate
 * @version 6.2.1
 */
class FileTemplateUpdate {
    /**
     * Constructs a new <code>FileTemplateUpdate</code>.
     * @alias module:model/FileTemplateUpdate
     */
    constructor() { 
        
        FileTemplateUpdate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>FileTemplateUpdate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/FileTemplateUpdate} obj Optional instance to populate.
     * @return {module:model/FileTemplateUpdate} The populated <code>FileTemplateUpdate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new FileTemplateUpdate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('fileName')) {
                obj['fileName'] = ApiClient.convertToType(data['fileName'], 'String');
            }
            if (data.hasOwnProperty('filePath')) {
                obj['filePath'] = ApiClient.convertToType(data['filePath'], 'String');
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('templatePhase')) {
                obj['templatePhase'] = ApiClient.convertToType(data['templatePhase'], 'String');
            }
            if (data.hasOwnProperty('template')) {
                obj['template'] = ApiClient.convertToType(data['template'], 'String');
            }
            if (data.hasOwnProperty('fileOwner')) {
                obj['fileOwner'] = ApiClient.convertToType(data['fileOwner'], 'Number');
            }
            if (data.hasOwnProperty('settingName')) {
                obj['settingName'] = ApiClient.convertToType(data['settingName'], 'String');
            }
            if (data.hasOwnProperty('settingCategory')) {
                obj['settingCategory'] = ApiClient.convertToType(data['settingCategory'], 'String');
            }
        }
        return obj;
    }


}

/**
 * File template name
 * @member {String} name
 */
FileTemplateUpdate.prototype['name'] = undefined;

/**
 * Array of label strings, can be used for filtering.
 * @member {Array.<String>} labels
 */
FileTemplateUpdate.prototype['labels'] = undefined;

/**
 * Filename for the file template
 * @member {String} fileName
 */
FileTemplateUpdate.prototype['fileName'] = undefined;

/**
 * Path for the file template
 * @member {String} filePath
 */
FileTemplateUpdate.prototype['filePath'] = undefined;

/**
 * Category
 * @member {String} category
 */
FileTemplateUpdate.prototype['category'] = undefined;

/**
 * Template Phase, provision, start, etc.
 * @member {String} templatePhase
 */
FileTemplateUpdate.prototype['templatePhase'] = undefined;

/**
 * Template content, that is, the file template content itself.
 * @member {String} template
 */
FileTemplateUpdate.prototype['template'] = undefined;

/**
 * File Owner
 * @member {Number} fileOwner
 */
FileTemplateUpdate.prototype['fileOwner'] = undefined;

/**
 * Setting Name
 * @member {String} settingName
 */
FileTemplateUpdate.prototype['settingName'] = undefined;

/**
 * Setting Category
 * @member {String} settingCategory
 */
FileTemplateUpdate.prototype['settingCategory'] = undefined;






export default FileTemplateUpdate;
