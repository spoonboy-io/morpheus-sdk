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

/**
 * The BlueprintCFTCreateCloudFormationGit model module.
 * @module model/BlueprintCFTCreateCloudFormationGit
 * @version 6.1.1
 */
class BlueprintCFTCreateCloudFormationGit {
    /**
     * Constructs a new <code>BlueprintCFTCreateCloudFormationGit</code>.
     * @alias module:model/BlueprintCFTCreateCloudFormationGit
     * @param repoId {Number} Morpheus SCM Repository ID
     * @param path {String} Path to CloudFormation Files in the Repository
     * @param integrationId {Number} Morpheus SCM Integration ID
     * @param branch {String} Branch Name
     */
    constructor(repoId, path, integrationId, branch) { 
        
        BlueprintCFTCreateCloudFormationGit.initialize(this, repoId, path, integrationId, branch);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, repoId, path, integrationId, branch) { 
        obj['repoId'] = repoId;
        obj['path'] = path;
        obj['integrationId'] = integrationId;
        obj['branch'] = branch;
    }

    /**
     * Constructs a <code>BlueprintCFTCreateCloudFormationGit</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BlueprintCFTCreateCloudFormationGit} obj Optional instance to populate.
     * @return {module:model/BlueprintCFTCreateCloudFormationGit} The populated <code>BlueprintCFTCreateCloudFormationGit</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BlueprintCFTCreateCloudFormationGit();

            if (data.hasOwnProperty('repoId')) {
                obj['repoId'] = ApiClient.convertToType(data['repoId'], 'Number');
            }
            if (data.hasOwnProperty('path')) {
                obj['path'] = ApiClient.convertToType(data['path'], 'String');
            }
            if (data.hasOwnProperty('integrationId')) {
                obj['integrationId'] = ApiClient.convertToType(data['integrationId'], 'Number');
            }
            if (data.hasOwnProperty('branch')) {
                obj['branch'] = ApiClient.convertToType(data['branch'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>BlueprintCFTCreateCloudFormationGit</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>BlueprintCFTCreateCloudFormationGit</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of BlueprintCFTCreateCloudFormationGit.RequiredProperties) {
            if (!data[property]) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }
        // ensure the json data is a string
        if (data['path'] && !(typeof data['path'] === 'string' || data['path'] instanceof String)) {
            throw new Error("Expected the field `path` to be a primitive type in the JSON string but got " + data['path']);
        }
        // ensure the json data is a string
        if (data['branch'] && !(typeof data['branch'] === 'string' || data['branch'] instanceof String)) {
            throw new Error("Expected the field `branch` to be a primitive type in the JSON string but got " + data['branch']);
        }

        return true;
    }


}

BlueprintCFTCreateCloudFormationGit.RequiredProperties = ["repoId", "path", "integrationId", "branch"];

/**
 * Morpheus SCM Repository ID
 * @member {Number} repoId
 */
BlueprintCFTCreateCloudFormationGit.prototype['repoId'] = undefined;

/**
 * Path to CloudFormation Files in the Repository
 * @member {String} path
 */
BlueprintCFTCreateCloudFormationGit.prototype['path'] = undefined;

/**
 * Morpheus SCM Integration ID
 * @member {Number} integrationId
 */
BlueprintCFTCreateCloudFormationGit.prototype['integrationId'] = undefined;

/**
 * Branch Name
 * @member {String} branch
 */
BlueprintCFTCreateCloudFormationGit.prototype['branch'] = undefined;






export default BlueprintCFTCreateCloudFormationGit;

