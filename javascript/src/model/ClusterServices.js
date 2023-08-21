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
 * The ClusterServices model module.
 * @module model/ClusterServices
 * @version 6.2.1
 */
class ClusterServices {
    /**
     * Constructs a new <code>ClusterServices</code>.
     * @alias module:model/ClusterServices
     */
    constructor() { 
        
        ClusterServices.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ClusterServices</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ClusterServices} obj Optional instance to populate.
     * @return {module:model/ClusterServices} The populated <code>ClusterServices</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ClusterServices();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('externalIp')) {
                obj['externalIp'] = ApiClient.convertToType(data['externalIp'], 'String');
            }
            if (data.hasOwnProperty('internalIp')) {
                obj['internalIp'] = ApiClient.convertToType(data['internalIp'], 'String');
            }
            if (data.hasOwnProperty('externalPort')) {
                obj['externalPort'] = ApiClient.convertToType(data['externalPort'], 'String');
            }
            if (data.hasOwnProperty('internalPort')) {
                obj['internalPort'] = ApiClient.convertToType(data['internalPort'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
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
ClusterServices.prototype['id'] = undefined;

/**
 * @member {String} name
 */
ClusterServices.prototype['name'] = undefined;

/**
 * @member {String} type
 */
ClusterServices.prototype['type'] = undefined;

/**
 * @member {String} code
 */
ClusterServices.prototype['code'] = undefined;

/**
 * @member {String} externalIp
 */
ClusterServices.prototype['externalIp'] = undefined;

/**
 * @member {String} internalIp
 */
ClusterServices.prototype['internalIp'] = undefined;

/**
 * @member {String} externalPort
 */
ClusterServices.prototype['externalPort'] = undefined;

/**
 * @member {String} internalPort
 */
ClusterServices.prototype['internalPort'] = undefined;

/**
 * @member {String} status
 */
ClusterServices.prototype['status'] = undefined;

/**
 * @member {Date} dateCreated
 */
ClusterServices.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
ClusterServices.prototype['lastUpdated'] = undefined;






export default ClusterServices;
