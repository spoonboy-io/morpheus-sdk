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
import SpecTemplateUpdateConfig from './SpecTemplateUpdateConfig';
import SpecTemplateUpdateFile from './SpecTemplateUpdateFile';
import SpecTemplateUpdateType from './SpecTemplateUpdateType';

/**
 * The SpecTemplateUpdate model module.
 * @module model/SpecTemplateUpdate
 * @version 6.2.1
 */
class SpecTemplateUpdate {
    /**
     * Constructs a new <code>SpecTemplateUpdate</code>.
     * @alias module:model/SpecTemplateUpdate
     */
    constructor() { 
        
        SpecTemplateUpdate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>SpecTemplateUpdate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/SpecTemplateUpdate} obj Optional instance to populate.
     * @return {module:model/SpecTemplateUpdate} The populated <code>SpecTemplateUpdate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new SpecTemplateUpdate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = SpecTemplateUpdateType.constructFromObject(data['type']);
            }
            if (data.hasOwnProperty('file')) {
                obj['file'] = SpecTemplateUpdateFile.constructFromObject(data['file']);
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = SpecTemplateUpdateConfig.constructFromObject(data['config']);
            }
        }
        return obj;
    }


}

/**
 * Spec template name
 * @member {String} name
 */
SpecTemplateUpdate.prototype['name'] = undefined;

/**
 * Array of label strings, can be used for filtering.
 * @member {Array.<String>} labels
 */
SpecTemplateUpdate.prototype['labels'] = undefined;

/**
 * @member {module:model/SpecTemplateUpdateType} type
 */
SpecTemplateUpdate.prototype['type'] = undefined;

/**
 * @member {module:model/SpecTemplateUpdateFile} file
 */
SpecTemplateUpdate.prototype['file'] = undefined;

/**
 * @member {module:model/SpecTemplateUpdateConfig} config
 */
SpecTemplateUpdate.prototype['config'] = undefined;






export default SpecTemplateUpdate;

