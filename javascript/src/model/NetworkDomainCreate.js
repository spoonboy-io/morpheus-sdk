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
 * The NetworkDomainCreate model module.
 * @module model/NetworkDomainCreate
 * @version 6.2.1
 */
class NetworkDomainCreate {
    /**
     * Constructs a new <code>NetworkDomainCreate</code>.
     * @alias module:model/NetworkDomainCreate
     */
    constructor() { 
        
        NetworkDomainCreate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>NetworkDomainCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NetworkDomainCreate} obj Optional instance to populate.
     * @return {module:model/NetworkDomainCreate} The populated <code>NetworkDomainCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NetworkDomainCreate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('displayName')) {
                obj['displayName'] = ApiClient.convertToType(data['displayName'], 'String');
            }
            if (data.hasOwnProperty('publicZone')) {
                obj['publicZone'] = ApiClient.convertToType(data['publicZone'], 'Boolean');
            }
            if (data.hasOwnProperty('taskSetId')) {
                obj['taskSetId'] = ApiClient.convertToType(data['taskSetId'], 'Number');
            }
            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
            if (data.hasOwnProperty('domainController')) {
                obj['domainController'] = ApiClient.convertToType(data['domainController'], 'Boolean');
            }
            if (data.hasOwnProperty('domainUsername')) {
                obj['domainUsername'] = ApiClient.convertToType(data['domainUsername'], 'String');
            }
            if (data.hasOwnProperty('domainPassword')) {
                obj['domainPassword'] = ApiClient.convertToType(data['domainPassword'], 'String');
            }
            if (data.hasOwnProperty('dcServer')) {
                obj['dcServer'] = ApiClient.convertToType(data['dcServer'], 'String');
            }
            if (data.hasOwnProperty('ouPath')) {
                obj['ouPath'] = ApiClient.convertToType(data['ouPath'], 'String');
            }
            if (data.hasOwnProperty('guestUsername')) {
                obj['guestUsername'] = ApiClient.convertToType(data['guestUsername'], 'String');
            }
            if (data.hasOwnProperty('guestPassword')) {
                obj['guestPassword'] = ApiClient.convertToType(data['guestPassword'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Name
 * @member {String} name
 */
NetworkDomainCreate.prototype['name'] = undefined;

/**
 * Description
 * @member {String} description
 */
NetworkDomainCreate.prototype['description'] = undefined;

/**
 * Overrides displayed name in domain selection components. Useful if using many OU Paths.
 * @member {String} displayName
 */
NetworkDomainCreate.prototype['displayName'] = undefined;

/**
 * Public Zone
 * @member {Boolean} publicZone
 * @default false
 */
NetworkDomainCreate.prototype['publicZone'] = false;

/**
 * Workflow ID. Workflows can be applied to an instance when associated with a domain. Useful for custom domain related scripting. (Important if wanting to ensure the computer is removed from the domain using teardown phased workflows.) 
 * @member {Number} taskSetId
 */
NetworkDomainCreate.prototype['taskSetId'] = undefined;

/**
 * Active
 * @member {Boolean} active
 */
NetworkDomainCreate.prototype['active'] = undefined;

/**
 * Join Domain Controller
 * @member {Boolean} domainController
 * @default true
 */
NetworkDomainCreate.prototype['domainController'] = true;

/**
 * Domain Username
 * @member {String} domainUsername
 */
NetworkDomainCreate.prototype['domainUsername'] = undefined;

/**
 * Domain Password
 * @member {String} domainPassword
 */
NetworkDomainCreate.prototype['domainPassword'] = undefined;

/**
 * DC Server. If specified, must be the server name and not an IP Address
 * @member {String} dcServer
 */
NetworkDomainCreate.prototype['dcServer'] = undefined;

/**
 * OU Path. (i.e. 'OU=staging,DC=ad,DC=yourdomain,DC=com')
 * @member {String} ouPath
 */
NetworkDomainCreate.prototype['ouPath'] = undefined;

/**
 * Guest Username. If set, will change the instances RPC Service User after joining a Domain.
 * @member {String} guestUsername
 */
NetworkDomainCreate.prototype['guestUsername'] = undefined;

/**
 * Guest Password
 * @member {String} guestPassword
 */
NetworkDomainCreate.prototype['guestPassword'] = undefined;






export default NetworkDomainCreate;

