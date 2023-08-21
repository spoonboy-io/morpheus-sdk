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
import DeploymentVersions from './DeploymentVersions';

/**
 * The Deployment model module.
 * @module model/Deployment
 * @version 6.2.1
 */
class Deployment {
    /**
     * Constructs a new <code>Deployment</code>.
     * @alias module:model/Deployment
     */
    constructor() { 
        
        Deployment.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Deployment</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Deployment} obj Optional instance to populate.
     * @return {module:model/Deployment} The populated <code>Deployment</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Deployment();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('accountId')) {
                obj['accountId'] = ApiClient.convertToType(data['accountId'], 'Number');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
            if (data.hasOwnProperty('versionCount')) {
                obj['versionCount'] = ApiClient.convertToType(data['versionCount'], 'Number');
            }
            if (data.hasOwnProperty('versions')) {
                obj['versions'] = ApiClient.convertToType(data['versions'], [DeploymentVersions]);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
Deployment.prototype['id'] = undefined;

/**
 * @member {String} name
 */
Deployment.prototype['name'] = undefined;

/**
 * @member {String} description
 */
Deployment.prototype['description'] = undefined;

/**
 * @member {Number} accountId
 */
Deployment.prototype['accountId'] = undefined;

/**
 * @member {String} externalId
 */
Deployment.prototype['externalId'] = undefined;

/**
 * @member {Date} dateCreated
 */
Deployment.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
Deployment.prototype['lastUpdated'] = undefined;

/**
 * @member {Number} versionCount
 */
Deployment.prototype['versionCount'] = undefined;

/**
 * @member {Array.<module:model/DeploymentVersions>} versions
 */
Deployment.prototype['versions'] = undefined;






export default Deployment;
