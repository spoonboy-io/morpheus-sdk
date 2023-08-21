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
 * The Deployments model module.
 * @module model/Deployments
 * @version 6.2.1
 */
class Deployments {
    /**
     * Constructs a new <code>Deployments</code>.
     * @alias module:model/Deployments
     */
    constructor() { 
        
        Deployments.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Deployments</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Deployments} obj Optional instance to populate.
     * @return {module:model/Deployments} The populated <code>Deployments</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Deployments();

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
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
Deployments.prototype['id'] = undefined;

/**
 * @member {String} name
 */
Deployments.prototype['name'] = undefined;

/**
 * @member {String} description
 */
Deployments.prototype['description'] = undefined;

/**
 * @member {Number} accountId
 */
Deployments.prototype['accountId'] = undefined;

/**
 * @member {String} externalId
 */
Deployments.prototype['externalId'] = undefined;

/**
 * @member {Date} dateCreated
 */
Deployments.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
Deployments.prototype['lastUpdated'] = undefined;

/**
 * @member {Number} versionCount
 */
Deployments.prototype['versionCount'] = undefined;






export default Deployments;
