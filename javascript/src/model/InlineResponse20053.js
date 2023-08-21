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
import ImageBuild from './ImageBuild';
import ImageBuildExecution from './ImageBuildExecution';

/**
 * The InlineResponse20053 model module.
 * @module model/InlineResponse20053
 * @version 6.2.1
 */
class InlineResponse20053 {
    /**
     * Constructs a new <code>InlineResponse20053</code>.
     * @alias module:model/InlineResponse20053
     */
    constructor() { 
        
        InlineResponse20053.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse20053</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse20053} obj Optional instance to populate.
     * @return {module:model/InlineResponse20053} The populated <code>InlineResponse20053</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse20053();

            if (data.hasOwnProperty('imageBuild')) {
                obj['imageBuild'] = ImageBuild.constructFromObject(data['imageBuild']);
            }
            if (data.hasOwnProperty('imageBuildExecutions')) {
                obj['imageBuildExecutions'] = ApiClient.convertToType(data['imageBuildExecutions'], [ImageBuildExecution]);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ImageBuild} imageBuild
 */
InlineResponse20053.prototype['imageBuild'] = undefined;

/**
 * @member {Array.<module:model/ImageBuildExecution>} imageBuildExecutions
 */
InlineResponse20053.prototype['imageBuildExecutions'] = undefined;






export default InlineResponse20053;

