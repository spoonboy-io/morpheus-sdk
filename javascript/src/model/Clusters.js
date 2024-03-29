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
import ClustersLayout from './ClustersLayout';
import ClustersServers from './ClustersServers';
import ClustersWorkerStats from './ClustersWorkerStats';
import ClustersZone from './ClustersZone';
import InlineResponse200107NetworkPoolCreatedBy from './InlineResponse200107NetworkPoolCreatedBy';
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';

/**
 * The Clusters model module.
 * @module model/Clusters
 * @version 6.2.1
 */
class Clusters {
    /**
     * Constructs a new <code>Clusters</code>.
     * @alias module:model/Clusters
     */
    constructor() { 
        
        Clusters.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Clusters</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Clusters} obj Optional instance to populate.
     * @return {module:model/Clusters} The populated <code>Clusters</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Clusters();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('location')) {
                obj['location'] = ApiClient.convertToType(data['location'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('serviceUrl')) {
                obj['serviceUrl'] = ApiClient.convertToType(data['serviceUrl'], 'String');
            }
            if (data.hasOwnProperty('serviceHost')) {
                obj['serviceHost'] = ApiClient.convertToType(data['serviceHost'], 'String');
            }
            if (data.hasOwnProperty('servicePath')) {
                obj['servicePath'] = ApiClient.convertToType(data['servicePath'], 'String');
            }
            if (data.hasOwnProperty('serviceHostname')) {
                obj['serviceHostname'] = ApiClient.convertToType(data['serviceHostname'], 'String');
            }
            if (data.hasOwnProperty('servicePort')) {
                obj['servicePort'] = ApiClient.convertToType(data['servicePort'], 'Number');
            }
            if (data.hasOwnProperty('serviceUsername')) {
                obj['serviceUsername'] = ApiClient.convertToType(data['serviceUsername'], 'String');
            }
            if (data.hasOwnProperty('servicePassword')) {
                obj['servicePassword'] = ApiClient.convertToType(data['servicePassword'], 'String');
            }
            if (data.hasOwnProperty('servicePasswordHash')) {
                obj['servicePasswordHash'] = ApiClient.convertToType(data['servicePasswordHash'], 'String');
            }
            if (data.hasOwnProperty('serviceToken')) {
                obj['serviceToken'] = ApiClient.convertToType(data['serviceToken'], 'String');
            }
            if (data.hasOwnProperty('serviceTokenHash')) {
                obj['serviceTokenHash'] = ApiClient.convertToType(data['serviceTokenHash'], 'String');
            }
            if (data.hasOwnProperty('serviceAccess')) {
                obj['serviceAccess'] = ApiClient.convertToType(data['serviceAccess'], 'String');
            }
            if (data.hasOwnProperty('serviceAccessHash')) {
                obj['serviceAccessHash'] = ApiClient.convertToType(data['serviceAccessHash'], 'String');
            }
            if (data.hasOwnProperty('serviceCert')) {
                obj['serviceCert'] = ApiClient.convertToType(data['serviceCert'], 'String');
            }
            if (data.hasOwnProperty('serviceCertHash')) {
                obj['serviceCertHash'] = ApiClient.convertToType(data['serviceCertHash'], 'String');
            }
            if (data.hasOwnProperty('serviceVersion')) {
                obj['serviceVersion'] = ApiClient.convertToType(data['serviceVersion'], 'String');
            }
            if (data.hasOwnProperty('searchDomains')) {
                obj['searchDomains'] = ApiClient.convertToType(data['searchDomains'], 'String');
            }
            if (data.hasOwnProperty('enableInternalDns')) {
                obj['enableInternalDns'] = ApiClient.convertToType(data['enableInternalDns'], 'Boolean');
            }
            if (data.hasOwnProperty('internalId')) {
                obj['internalId'] = ApiClient.convertToType(data['internalId'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('datacenterId')) {
                obj['datacenterId'] = ApiClient.convertToType(data['datacenterId'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('statusDate')) {
                obj['statusDate'] = ApiClient.convertToType(data['statusDate'], 'Date');
            }
            if (data.hasOwnProperty('statusMessage')) {
                obj['statusMessage'] = ApiClient.convertToType(data['statusMessage'], 'String');
            }
            if (data.hasOwnProperty('inventoryLevel')) {
                obj['inventoryLevel'] = ApiClient.convertToType(data['inventoryLevel'], 'String');
            }
            if (data.hasOwnProperty('lastSync')) {
                obj['lastSync'] = ApiClient.convertToType(data['lastSync'], 'Date');
            }
            if (data.hasOwnProperty('nextRunDate')) {
                obj['nextRunDate'] = ApiClient.convertToType(data['nextRunDate'], 'Date');
            }
            if (data.hasOwnProperty('lastSyncDuration')) {
                obj['lastSyncDuration'] = ApiClient.convertToType(data['lastSyncDuration'], 'Number');
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
            if (data.hasOwnProperty('managed')) {
                obj['managed'] = ApiClient.convertToType(data['managed'], 'Boolean');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('serviceEntry')) {
                obj['serviceEntry'] = ApiClient.convertToType(data['serviceEntry'], 'String');
            }
            if (data.hasOwnProperty('createdBy')) {
                obj['createdBy'] = InlineResponse200107NetworkPoolCreatedBy.constructFromObject(data['createdBy']);
            }
            if (data.hasOwnProperty('userGroup')) {
                obj['userGroup'] = ApiClient.convertToType(data['userGroup'], 'String');
            }
            if (data.hasOwnProperty('layout')) {
                obj['layout'] = ClustersLayout.constructFromObject(data['layout']);
            }
            if (data.hasOwnProperty('owner')) {
                obj['owner'] = InlineResponse20040AppDeployInstance.constructFromObject(data['owner']);
            }
            if (data.hasOwnProperty('servers')) {
                obj['servers'] = ApiClient.convertToType(data['servers'], [ClustersServers]);
            }
            if (data.hasOwnProperty('accounts')) {
                obj['accounts'] = ApiClient.convertToType(data['accounts'], [Object]);
            }
            if (data.hasOwnProperty('integrations')) {
                obj['integrations'] = ApiClient.convertToType(data['integrations'], [Object]);
            }
            if (data.hasOwnProperty('site')) {
                obj['site'] = InlineResponse20040AppDeployInstance.constructFromObject(data['site']);
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = InlineResponse20040AppDeployInstance.constructFromObject(data['type']);
            }
            if (data.hasOwnProperty('zone')) {
                obj['zone'] = ClustersZone.constructFromObject(data['zone']);
            }
            if (data.hasOwnProperty('workerStats')) {
                obj['workerStats'] = ClustersWorkerStats.constructFromObject(data['workerStats']);
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], Object);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
Clusters.prototype['id'] = undefined;

/**
 * @member {String} name
 */
Clusters.prototype['name'] = undefined;

/**
 * @member {String} code
 */
Clusters.prototype['code'] = undefined;

/**
 * @member {String} category
 */
Clusters.prototype['category'] = undefined;

/**
 * @member {String} visibility
 */
Clusters.prototype['visibility'] = undefined;

/**
 * @member {String} description
 */
Clusters.prototype['description'] = undefined;

/**
 * @member {String} location
 */
Clusters.prototype['location'] = undefined;

/**
 * @member {Boolean} enabled
 */
Clusters.prototype['enabled'] = undefined;

/**
 * @member {String} serviceUrl
 */
Clusters.prototype['serviceUrl'] = undefined;

/**
 * @member {String} serviceHost
 */
Clusters.prototype['serviceHost'] = undefined;

/**
 * @member {String} servicePath
 */
Clusters.prototype['servicePath'] = undefined;

/**
 * @member {String} serviceHostname
 */
Clusters.prototype['serviceHostname'] = undefined;

/**
 * @member {Number} servicePort
 */
Clusters.prototype['servicePort'] = undefined;

/**
 * @member {String} serviceUsername
 */
Clusters.prototype['serviceUsername'] = undefined;

/**
 * @member {String} servicePassword
 */
Clusters.prototype['servicePassword'] = undefined;

/**
 * @member {String} servicePasswordHash
 */
Clusters.prototype['servicePasswordHash'] = undefined;

/**
 * @member {String} serviceToken
 */
Clusters.prototype['serviceToken'] = undefined;

/**
 * @member {String} serviceTokenHash
 */
Clusters.prototype['serviceTokenHash'] = undefined;

/**
 * @member {String} serviceAccess
 */
Clusters.prototype['serviceAccess'] = undefined;

/**
 * @member {String} serviceAccessHash
 */
Clusters.prototype['serviceAccessHash'] = undefined;

/**
 * @member {String} serviceCert
 */
Clusters.prototype['serviceCert'] = undefined;

/**
 * @member {String} serviceCertHash
 */
Clusters.prototype['serviceCertHash'] = undefined;

/**
 * @member {String} serviceVersion
 */
Clusters.prototype['serviceVersion'] = undefined;

/**
 * @member {String} searchDomains
 */
Clusters.prototype['searchDomains'] = undefined;

/**
 * @member {Boolean} enableInternalDns
 */
Clusters.prototype['enableInternalDns'] = undefined;

/**
 * @member {String} internalId
 */
Clusters.prototype['internalId'] = undefined;

/**
 * @member {String} externalId
 */
Clusters.prototype['externalId'] = undefined;

/**
 * @member {String} datacenterId
 */
Clusters.prototype['datacenterId'] = undefined;

/**
 * @member {String} status
 */
Clusters.prototype['status'] = undefined;

/**
 * @member {Date} statusDate
 */
Clusters.prototype['statusDate'] = undefined;

/**
 * @member {String} statusMessage
 */
Clusters.prototype['statusMessage'] = undefined;

/**
 * @member {String} inventoryLevel
 */
Clusters.prototype['inventoryLevel'] = undefined;

/**
 * @member {Date} lastSync
 */
Clusters.prototype['lastSync'] = undefined;

/**
 * @member {Date} nextRunDate
 */
Clusters.prototype['nextRunDate'] = undefined;

/**
 * @member {Number} lastSyncDuration
 */
Clusters.prototype['lastSyncDuration'] = undefined;

/**
 * @member {Date} dateCreated
 */
Clusters.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
Clusters.prototype['lastUpdated'] = undefined;

/**
 * @member {Boolean} managed
 */
Clusters.prototype['managed'] = undefined;

/**
 * @member {Array.<String>} labels
 */
Clusters.prototype['labels'] = undefined;

/**
 * @member {String} serviceEntry
 */
Clusters.prototype['serviceEntry'] = undefined;

/**
 * @member {module:model/InlineResponse200107NetworkPoolCreatedBy} createdBy
 */
Clusters.prototype['createdBy'] = undefined;

/**
 * @member {String} userGroup
 */
Clusters.prototype['userGroup'] = undefined;

/**
 * @member {module:model/ClustersLayout} layout
 */
Clusters.prototype['layout'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} owner
 */
Clusters.prototype['owner'] = undefined;

/**
 * @member {Array.<module:model/ClustersServers>} servers
 */
Clusters.prototype['servers'] = undefined;

/**
 * @member {Array.<Object>} accounts
 */
Clusters.prototype['accounts'] = undefined;

/**
 * @member {Array.<Object>} integrations
 */
Clusters.prototype['integrations'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} site
 */
Clusters.prototype['site'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} type
 */
Clusters.prototype['type'] = undefined;

/**
 * @member {module:model/ClustersZone} zone
 */
Clusters.prototype['zone'] = undefined;

/**
 * @member {module:model/ClustersWorkerStats} workerStats
 */
Clusters.prototype['workerStats'] = undefined;

/**
 * @member {Object} config
 */
Clusters.prototype['config'] = undefined;






export default Clusters;

