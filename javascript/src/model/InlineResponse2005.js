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
import ArchiveBucketFile from './ArchiveBucketFile';

/**
 * The InlineResponse2005 model module.
 * @module model/InlineResponse2005
 * @version 6.2.1
 */
class InlineResponse2005 {
    /**
     * Constructs a new <code>InlineResponse2005</code>.
     * @alias module:model/InlineResponse2005
     */
    constructor() { 
        
        InlineResponse2005.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse2005</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse2005} obj Optional instance to populate.
     * @return {module:model/InlineResponse2005} The populated <code>InlineResponse2005</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse2005();

            if (data.hasOwnProperty('archiveFile')) {
                obj['archiveFile'] = ArchiveBucketFile.constructFromObject(data['archiveFile']);
            }
            if (data.hasOwnProperty('isOwner')) {
                obj['isOwner'] = ApiClient.convertToType(data['isOwner'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ArchiveBucketFile} archiveFile
 */
InlineResponse2005.prototype['archiveFile'] = undefined;

/**
 * @member {Boolean} isOwner
 */
InlineResponse2005.prototype['isOwner'] = undefined;






export default InlineResponse2005;
