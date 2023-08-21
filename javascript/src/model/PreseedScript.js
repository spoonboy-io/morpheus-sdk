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
import ArchiveBucketFileCreatedBy from './ArchiveBucketFileCreatedBy';
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';

/**
 * The PreseedScript model module.
 * @module model/PreseedScript
 * @version 6.2.1
 */
class PreseedScript {
    /**
     * Constructs a new <code>PreseedScript</code>.
     * @alias module:model/PreseedScript
     */
    constructor() { 
        
        PreseedScript.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>PreseedScript</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/PreseedScript} obj Optional instance to populate.
     * @return {module:model/PreseedScript} The populated <code>PreseedScript</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new PreseedScript();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('account')) {
                obj['account'] = InlineResponse20040AppDeployInstance.constructFromObject(data['account']);
            }
            if (data.hasOwnProperty('fileName')) {
                obj['fileName'] = ApiClient.convertToType(data['fileName'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('content')) {
                obj['content'] = ApiClient.convertToType(data['content'], 'String');
            }
            if (data.hasOwnProperty('createdBy')) {
                obj['createdBy'] = ArchiveBucketFileCreatedBy.constructFromObject(data['createdBy']);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
PreseedScript.prototype['id'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} account
 */
PreseedScript.prototype['account'] = undefined;

/**
 * @member {String} fileName
 */
PreseedScript.prototype['fileName'] = undefined;

/**
 * @member {String} description
 */
PreseedScript.prototype['description'] = undefined;

/**
 * @member {String} content
 */
PreseedScript.prototype['content'] = undefined;

/**
 * @member {module:model/ArchiveBucketFileCreatedBy} createdBy
 */
PreseedScript.prototype['createdBy'] = undefined;






export default PreseedScript;
