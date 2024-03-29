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
import SpecTemplateCreateConfig from './SpecTemplateCreateConfig';
import SpecTemplateCreateFile from './SpecTemplateCreateFile';
import SpecTemplateCreateType from './SpecTemplateCreateType';

/**
 * The SpecTemplateCreate model module.
 * @module model/SpecTemplateCreate
 * @version 6.2.1
 */
class SpecTemplateCreate {
    /**
     * Constructs a new <code>SpecTemplateCreate</code>.
     * @alias module:model/SpecTemplateCreate
     * @param name {String} Spec template name
     * @param type {module:model/SpecTemplateCreateType} 
     * @param file {module:model/SpecTemplateCreateFile} 
     */
    constructor(name, type, file) { 
        
        SpecTemplateCreate.initialize(this, name, type, file);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name, type, file) { 
        obj['name'] = name;
        obj['type'] = type;
        obj['file'] = file;
    }

    /**
     * Constructs a <code>SpecTemplateCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/SpecTemplateCreate} obj Optional instance to populate.
     * @return {module:model/SpecTemplateCreate} The populated <code>SpecTemplateCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new SpecTemplateCreate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = SpecTemplateCreateType.constructFromObject(data['type']);
            }
            if (data.hasOwnProperty('file')) {
                obj['file'] = SpecTemplateCreateFile.constructFromObject(data['file']);
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = SpecTemplateCreateConfig.constructFromObject(data['config']);
            }
        }
        return obj;
    }


}

/**
 * Spec template name
 * @member {String} name
 */
SpecTemplateCreate.prototype['name'] = undefined;

/**
 * Array of label strings, can be used for filtering.
 * @member {Array.<String>} labels
 */
SpecTemplateCreate.prototype['labels'] = undefined;

/**
 * @member {module:model/SpecTemplateCreateType} type
 */
SpecTemplateCreate.prototype['type'] = undefined;

/**
 * @member {module:model/SpecTemplateCreateFile} file
 */
SpecTemplateCreate.prototype['file'] = undefined;

/**
 * @member {module:model/SpecTemplateCreateConfig} config
 */
SpecTemplateCreate.prototype['config'] = undefined;






export default SpecTemplateCreate;

