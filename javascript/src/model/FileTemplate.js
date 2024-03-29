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
import InlineResponse20082LoadBalancerInstanceSslCert from './InlineResponse20082LoadBalancerInstanceSslCert';

/**
 * The FileTemplate model module.
 * @module model/FileTemplate
 * @version 6.2.1
 */
class FileTemplate {
    /**
     * Constructs a new <code>FileTemplate</code>.
     * @alias module:model/FileTemplate
     */
    constructor() { 
        
        FileTemplate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>FileTemplate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/FileTemplate} obj Optional instance to populate.
     * @return {module:model/FileTemplate} The populated <code>FileTemplate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new FileTemplate();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('account')) {
                obj['account'] = InlineResponse20082LoadBalancerInstanceSslCert.constructFromObject(data['account']);
            }
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
            if (data.hasOwnProperty('templateType')) {
                obj['templateType'] = ApiClient.convertToType(data['templateType'], 'String');
            }
            if (data.hasOwnProperty('templatePhase')) {
                obj['templatePhase'] = ApiClient.convertToType(data['templatePhase'], 'String');
            }
            if (data.hasOwnProperty('template')) {
                obj['template'] = ApiClient.convertToType(data['template'], 'String');
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('settingCategory')) {
                obj['settingCategory'] = ApiClient.convertToType(data['settingCategory'], 'String');
            }
            if (data.hasOwnProperty('settingName')) {
                obj['settingName'] = ApiClient.convertToType(data['settingName'], 'String');
            }
            if (data.hasOwnProperty('autoRun')) {
                obj['autoRun'] = ApiClient.convertToType(data['autoRun'], 'Boolean');
            }
            if (data.hasOwnProperty('runOnScale')) {
                obj['runOnScale'] = ApiClient.convertToType(data['runOnScale'], 'Boolean');
            }
            if (data.hasOwnProperty('runOnDeploy')) {
                obj['runOnDeploy'] = ApiClient.convertToType(data['runOnDeploy'], 'Boolean');
            }
            if (data.hasOwnProperty('fileOwner')) {
                obj['fileOwner'] = ApiClient.convertToType(data['fileOwner'], 'String');
            }
            if (data.hasOwnProperty('fileGroup')) {
                obj['fileGroup'] = ApiClient.convertToType(data['fileGroup'], 'String');
            }
            if (data.hasOwnProperty('permissions')) {
                obj['permissions'] = ApiClient.convertToType(data['permissions'], 'String');
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
FileTemplate.prototype['id'] = undefined;

/**
 * @member {String} code
 */
FileTemplate.prototype['code'] = undefined;

/**
 * @member {module:model/InlineResponse20082LoadBalancerInstanceSslCert} account
 */
FileTemplate.prototype['account'] = undefined;

/**
 * @member {String} name
 */
FileTemplate.prototype['name'] = undefined;

/**
 * @member {Array.<String>} labels
 */
FileTemplate.prototype['labels'] = undefined;

/**
 * @member {String} fileName
 */
FileTemplate.prototype['fileName'] = undefined;

/**
 * @member {String} filePath
 */
FileTemplate.prototype['filePath'] = undefined;

/**
 * @member {String} templateType
 */
FileTemplate.prototype['templateType'] = undefined;

/**
 * @member {String} templatePhase
 */
FileTemplate.prototype['templatePhase'] = undefined;

/**
 * @member {String} template
 */
FileTemplate.prototype['template'] = undefined;

/**
 * @member {String} category
 */
FileTemplate.prototype['category'] = undefined;

/**
 * @member {String} settingCategory
 */
FileTemplate.prototype['settingCategory'] = undefined;

/**
 * @member {String} settingName
 */
FileTemplate.prototype['settingName'] = undefined;

/**
 * @member {Boolean} autoRun
 */
FileTemplate.prototype['autoRun'] = undefined;

/**
 * @member {Boolean} runOnScale
 */
FileTemplate.prototype['runOnScale'] = undefined;

/**
 * @member {Boolean} runOnDeploy
 */
FileTemplate.prototype['runOnDeploy'] = undefined;

/**
 * @member {String} fileOwner
 */
FileTemplate.prototype['fileOwner'] = undefined;

/**
 * @member {String} fileGroup
 */
FileTemplate.prototype['fileGroup'] = undefined;

/**
 * @member {String} permissions
 */
FileTemplate.prototype['permissions'] = undefined;

/**
 * @member {Date} dateCreated
 */
FileTemplate.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
FileTemplate.prototype['lastUpdated'] = undefined;






export default FileTemplate;

