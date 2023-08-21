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
import Creds from './Creds';
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';
import InlineResponse20094Network from './InlineResponse20094Network';
import Owner from './Owner';
import ResourcePermissions from './ResourcePermissions';

/**
 * The LoadBalancer model module.
 * @module model/LoadBalancer
 * @version 6.2.1
 */
class LoadBalancer {
    /**
     * Constructs a new <code>LoadBalancer</code>.
     * @alias module:model/LoadBalancer
     */
    constructor() { 
        
        LoadBalancer.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>LoadBalancer</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/LoadBalancer} obj Optional instance to populate.
     * @return {module:model/LoadBalancer} The populated <code>LoadBalancer</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new LoadBalancer();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('uuid')) {
                obj['uuid'] = ApiClient.convertToType(data['uuid'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('accountId')) {
                obj['accountId'] = ApiClient.convertToType(data['accountId'], 'Number');
            }
            if (data.hasOwnProperty('cloud')) {
                obj['cloud'] = InlineResponse20040AppDeployInstance.constructFromObject(data['cloud']);
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = InlineResponse20094Network.constructFromObject(data['type']);
            }
            if (data.hasOwnProperty('owner')) {
                obj['owner'] = Owner.constructFromObject(data['owner']);
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('host')) {
                obj['host'] = ApiClient.convertToType(data['host'], 'String');
            }
            if (data.hasOwnProperty('port')) {
                obj['port'] = ApiClient.convertToType(data['port'], 'Number');
            }
            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
            }
            if (data.hasOwnProperty('ip')) {
                obj['ip'] = ApiClient.convertToType(data['ip'], 'String');
            }
            if (data.hasOwnProperty('internalIp')) {
                obj['internalIp'] = ApiClient.convertToType(data['internalIp'], 'String');
            }
            if (data.hasOwnProperty('externalIp')) {
                obj['externalIp'] = ApiClient.convertToType(data['externalIp'], 'String');
            }
            if (data.hasOwnProperty('apiPort')) {
                obj['apiPort'] = ApiClient.convertToType(data['apiPort'], 'String');
            }
            if (data.hasOwnProperty('adminPort')) {
                obj['adminPort'] = ApiClient.convertToType(data['adminPort'], 'String');
            }
            if (data.hasOwnProperty('sslEnabled')) {
                obj['sslEnabled'] = ApiClient.convertToType(data['sslEnabled'], 'Boolean');
            }
            if (data.hasOwnProperty('sslCert')) {
                obj['sslCert'] = ApiClient.convertToType(data['sslCert'], 'String');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], Object);
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
            if (data.hasOwnProperty('credential')) {
                obj['credential'] = Creds.constructFromObject(data['credential']);
            }
            if (data.hasOwnProperty('tenants')) {
                obj['tenants'] = ApiClient.convertToType(data['tenants'], [Object]);
            }
            if (data.hasOwnProperty('resourcePermission')) {
                obj['resourcePermission'] = ResourcePermissions.constructFromObject(data['resourcePermission']);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
LoadBalancer.prototype['id'] = undefined;

/**
 * @member {String} uuid
 */
LoadBalancer.prototype['uuid'] = undefined;

/**
 * @member {String} name
 */
LoadBalancer.prototype['name'] = undefined;

/**
 * @member {Number} accountId
 */
LoadBalancer.prototype['accountId'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} cloud
 */
LoadBalancer.prototype['cloud'] = undefined;

/**
 * @member {module:model/InlineResponse20094Network} type
 */
LoadBalancer.prototype['type'] = undefined;

/**
 * @member {module:model/Owner} owner
 */
LoadBalancer.prototype['owner'] = undefined;

/**
 * @member {String} visibility
 */
LoadBalancer.prototype['visibility'] = undefined;

/**
 * @member {String} description
 */
LoadBalancer.prototype['description'] = undefined;

/**
 * @member {String} host
 */
LoadBalancer.prototype['host'] = undefined;

/**
 * @member {Number} port
 */
LoadBalancer.prototype['port'] = undefined;

/**
 * @member {String} username
 */
LoadBalancer.prototype['username'] = undefined;

/**
 * @member {String} ip
 */
LoadBalancer.prototype['ip'] = undefined;

/**
 * @member {String} internalIp
 */
LoadBalancer.prototype['internalIp'] = undefined;

/**
 * @member {String} externalIp
 */
LoadBalancer.prototype['externalIp'] = undefined;

/**
 * @member {String} apiPort
 */
LoadBalancer.prototype['apiPort'] = undefined;

/**
 * @member {String} adminPort
 */
LoadBalancer.prototype['adminPort'] = undefined;

/**
 * @member {Boolean} sslEnabled
 */
LoadBalancer.prototype['sslEnabled'] = undefined;

/**
 * @member {String} sslCert
 */
LoadBalancer.prototype['sslCert'] = undefined;

/**
 * @member {Object} config
 */
LoadBalancer.prototype['config'] = undefined;

/**
 * @member {Date} dateCreated
 */
LoadBalancer.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
LoadBalancer.prototype['lastUpdated'] = undefined;

/**
 * @member {module:model/Creds} credential
 */
LoadBalancer.prototype['credential'] = undefined;

/**
 * @member {Array.<Object>} tenants
 */
LoadBalancer.prototype['tenants'] = undefined;

/**
 * @member {module:model/ResourcePermissions} resourcePermission
 */
LoadBalancer.prototype['resourcePermission'] = undefined;






export default LoadBalancer;

