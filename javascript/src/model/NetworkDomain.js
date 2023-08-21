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
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';

/**
 * The NetworkDomain model module.
 * @module model/NetworkDomain
 * @version 6.2.1
 */
class NetworkDomain {
    /**
     * Constructs a new <code>NetworkDomain</code>.
     * @alias module:model/NetworkDomain
     */
    constructor() { 
        
        NetworkDomain.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>NetworkDomain</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NetworkDomain} obj Optional instance to populate.
     * @return {module:model/NetworkDomain} The populated <code>NetworkDomain</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NetworkDomain();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
            if (data.hasOwnProperty('fqdn')) {
                obj['fqdn'] = ApiClient.convertToType(data['fqdn'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('domainController')) {
                obj['domainController'] = ApiClient.convertToType(data['domainController'], 'Boolean');
            }
            if (data.hasOwnProperty('publicZone')) {
                obj['publicZone'] = ApiClient.convertToType(data['publicZone'], 'Boolean');
            }
            if (data.hasOwnProperty('domainUsername')) {
                obj['domainUsername'] = ApiClient.convertToType(data['domainUsername'], 'String');
            }
            if (data.hasOwnProperty('domainPassword')) {
                obj['domainPassword'] = ApiClient.convertToType(data['domainPassword'], 'String');
            }
            if (data.hasOwnProperty('refType')) {
                obj['refType'] = ApiClient.convertToType(data['refType'], 'String');
            }
            if (data.hasOwnProperty('refId')) {
                obj['refId'] = ApiClient.convertToType(data['refId'], 'Number');
            }
            if (data.hasOwnProperty('refSource')) {
                obj['refSource'] = ApiClient.convertToType(data['refSource'], 'String');
            }
            if (data.hasOwnProperty('internalId')) {
                obj['internalId'] = ApiClient.convertToType(data['internalId'], 'String');
            }
            if (data.hasOwnProperty('ouPath')) {
                obj['ouPath'] = ApiClient.convertToType(data['ouPath'], 'String');
            }
            if (data.hasOwnProperty('dcServer')) {
                obj['dcServer'] = ApiClient.convertToType(data['dcServer'], 'String');
            }
            if (data.hasOwnProperty('zoneType')) {
                obj['zoneType'] = ApiClient.convertToType(data['zoneType'], 'String');
            }
            if (data.hasOwnProperty('dnssec')) {
                obj['dnssec'] = ApiClient.convertToType(data['dnssec'], 'String');
            }
            if (data.hasOwnProperty('domainSerial')) {
                obj['domainSerial'] = ApiClient.convertToType(data['domainSerial'], 'String');
            }
            if (data.hasOwnProperty('account')) {
                obj['account'] = InlineResponse20040AppDeployInstance.constructFromObject(data['account']);
            }
            if (data.hasOwnProperty('owner')) {
                obj['owner'] = InlineResponse20040AppDeployInstance.constructFromObject(data['owner']);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
NetworkDomain.prototype['id'] = undefined;

/**
 * @member {String} name
 */
NetworkDomain.prototype['name'] = undefined;

/**
 * @member {Boolean} active
 */
NetworkDomain.prototype['active'] = undefined;

/**
 * @member {String} fqdn
 */
NetworkDomain.prototype['fqdn'] = undefined;

/**
 * @member {String} description
 */
NetworkDomain.prototype['description'] = undefined;

/**
 * @member {String} visibility
 */
NetworkDomain.prototype['visibility'] = undefined;

/**
 * @member {Boolean} domainController
 */
NetworkDomain.prototype['domainController'] = undefined;

/**
 * @member {Boolean} publicZone
 */
NetworkDomain.prototype['publicZone'] = undefined;

/**
 * @member {String} domainUsername
 */
NetworkDomain.prototype['domainUsername'] = undefined;

/**
 * @member {String} domainPassword
 */
NetworkDomain.prototype['domainPassword'] = undefined;

/**
 * @member {String} refType
 */
NetworkDomain.prototype['refType'] = undefined;

/**
 * @member {Number} refId
 */
NetworkDomain.prototype['refId'] = undefined;

/**
 * @member {String} refSource
 */
NetworkDomain.prototype['refSource'] = undefined;

/**
 * @member {String} internalId
 */
NetworkDomain.prototype['internalId'] = undefined;

/**
 * @member {String} ouPath
 */
NetworkDomain.prototype['ouPath'] = undefined;

/**
 * @member {String} dcServer
 */
NetworkDomain.prototype['dcServer'] = undefined;

/**
 * @member {String} zoneType
 */
NetworkDomain.prototype['zoneType'] = undefined;

/**
 * @member {String} dnssec
 */
NetworkDomain.prototype['dnssec'] = undefined;

/**
 * @member {String} domainSerial
 */
NetworkDomain.prototype['domainSerial'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} account
 */
NetworkDomain.prototype['account'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} owner
 */
NetworkDomain.prototype['owner'] = undefined;






export default NetworkDomain;
