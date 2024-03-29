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
import ApiServersIdMakeManagedServerTags from './ApiServersIdMakeManagedServerTags';
import ClusterLayoutComputeServerType from './ClusterLayoutComputeServerType';
import ClusterMastersInterfaces from './ClusterMastersInterfaces';
import ClusterMastersStats from './ClusterMastersStats';
import ClusterMastersVolumes from './ClusterMastersVolumes';
import InlineResponse200107NetworkPoolCreatedBy from './InlineResponse200107NetworkPoolCreatedBy';
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';
import InlineResponse20079LoadBalancerMonitorLoadBalancerType from './InlineResponse20079LoadBalancerMonitorLoadBalancerType';

/**
 * The ClusterMasters model module.
 * @module model/ClusterMasters
 * @version 6.2.1
 */
class ClusterMasters {
    /**
     * Constructs a new <code>ClusterMasters</code>.
     * @alias module:model/ClusterMasters
     */
    constructor() { 
        
        ClusterMasters.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ClusterMasters</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ClusterMasters} obj Optional instance to populate.
     * @return {module:model/ClusterMasters} The populated <code>ClusterMasters</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ClusterMasters();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('uuid')) {
                obj['uuid'] = ApiClient.convertToType(data['uuid'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('internalId')) {
                obj['internalId'] = ApiClient.convertToType(data['internalId'], 'String');
            }
            if (data.hasOwnProperty('externalUniqueId')) {
                obj['externalUniqueId'] = ApiClient.convertToType(data['externalUniqueId'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('externalName')) {
                obj['externalName'] = ApiClient.convertToType(data['externalName'], 'String');
            }
            if (data.hasOwnProperty('hostname')) {
                obj['hostname'] = ApiClient.convertToType(data['hostname'], 'String');
            }
            if (data.hasOwnProperty('accountId')) {
                obj['accountId'] = ApiClient.convertToType(data['accountId'], 'Number');
            }
            if (data.hasOwnProperty('account')) {
                obj['account'] = InlineResponse20040AppDeployInstance.constructFromObject(data['account']);
            }
            if (data.hasOwnProperty('owner')) {
                obj['owner'] = InlineResponse200107NetworkPoolCreatedBy.constructFromObject(data['owner']);
            }
            if (data.hasOwnProperty('zone')) {
                obj['zone'] = InlineResponse20040AppDeployInstance.constructFromObject(data['zone']);
            }
            if (data.hasOwnProperty('plan')) {
                obj['plan'] = InlineResponse20079LoadBalancerMonitorLoadBalancerType.constructFromObject(data['plan']);
            }
            if (data.hasOwnProperty('computeServerType')) {
                obj['computeServerType'] = ClusterLayoutComputeServerType.constructFromObject(data['computeServerType']);
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('zoneId')) {
                obj['zoneId'] = ApiClient.convertToType(data['zoneId'], 'Number');
            }
            if (data.hasOwnProperty('siteId')) {
                obj['siteId'] = ApiClient.convertToType(data['siteId'], 'Number');
            }
            if (data.hasOwnProperty('resourcePoolId')) {
                obj['resourcePoolId'] = ApiClient.convertToType(data['resourcePoolId'], 'Number');
            }
            if (data.hasOwnProperty('folderId')) {
                obj['folderId'] = ApiClient.convertToType(data['folderId'], 'String');
            }
            if (data.hasOwnProperty('sshHost')) {
                obj['sshHost'] = ApiClient.convertToType(data['sshHost'], 'String');
            }
            if (data.hasOwnProperty('sshPort')) {
                obj['sshPort'] = ApiClient.convertToType(data['sshPort'], 'Number');
            }
            if (data.hasOwnProperty('externalIp')) {
                obj['externalIp'] = ApiClient.convertToType(data['externalIp'], 'String');
            }
            if (data.hasOwnProperty('internalIp')) {
                obj['internalIp'] = ApiClient.convertToType(data['internalIp'], 'String');
            }
            if (data.hasOwnProperty('volumeId')) {
                obj['volumeId'] = ApiClient.convertToType(data['volumeId'], 'String');
            }
            if (data.hasOwnProperty('platform')) {
                obj['platform'] = ApiClient.convertToType(data['platform'], 'String');
            }
            if (data.hasOwnProperty('platformVersion')) {
                obj['platformVersion'] = ApiClient.convertToType(data['platformVersion'], 'String');
            }
            if (data.hasOwnProperty('sshUsername')) {
                obj['sshUsername'] = ApiClient.convertToType(data['sshUsername'], 'String');
            }
            if (data.hasOwnProperty('sshPassword')) {
                obj['sshPassword'] = ApiClient.convertToType(data['sshPassword'], 'String');
            }
            if (data.hasOwnProperty('sshPasswordHash')) {
                obj['sshPasswordHash'] = ApiClient.convertToType(data['sshPasswordHash'], 'String');
            }
            if (data.hasOwnProperty('osDevice')) {
                obj['osDevice'] = ApiClient.convertToType(data['osDevice'], 'String');
            }
            if (data.hasOwnProperty('osType')) {
                obj['osType'] = ApiClient.convertToType(data['osType'], 'String');
            }
            if (data.hasOwnProperty('dataDevice')) {
                obj['dataDevice'] = ApiClient.convertToType(data['dataDevice'], 'String');
            }
            if (data.hasOwnProperty('lvmEnabled')) {
                obj['lvmEnabled'] = ApiClient.convertToType(data['lvmEnabled'], 'Boolean');
            }
            if (data.hasOwnProperty('apiKey')) {
                obj['apiKey'] = ApiClient.convertToType(data['apiKey'], 'String');
            }
            if (data.hasOwnProperty('softwareRaid')) {
                obj['softwareRaid'] = ApiClient.convertToType(data['softwareRaid'], 'Boolean');
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
            if (data.hasOwnProperty('stats')) {
                obj['stats'] = ClusterMastersStats.constructFromObject(data['stats']);
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('statusMessage')) {
                obj['statusMessage'] = ApiClient.convertToType(data['statusMessage'], 'String');
            }
            if (data.hasOwnProperty('errorMessage')) {
                obj['errorMessage'] = ApiClient.convertToType(data['errorMessage'], 'String');
            }
            if (data.hasOwnProperty('statusDate')) {
                obj['statusDate'] = ApiClient.convertToType(data['statusDate'], 'String');
            }
            if (data.hasOwnProperty('statusPercent')) {
                obj['statusPercent'] = ApiClient.convertToType(data['statusPercent'], 'String');
            }
            if (data.hasOwnProperty('statusEta')) {
                obj['statusEta'] = ApiClient.convertToType(data['statusEta'], 'String');
            }
            if (data.hasOwnProperty('powerState')) {
                obj['powerState'] = ApiClient.convertToType(data['powerState'], 'String');
            }
            if (data.hasOwnProperty('agentInstalled')) {
                obj['agentInstalled'] = ApiClient.convertToType(data['agentInstalled'], 'Boolean');
            }
            if (data.hasOwnProperty('lastAgentUpdate')) {
                obj['lastAgentUpdate'] = ApiClient.convertToType(data['lastAgentUpdate'], 'Date');
            }
            if (data.hasOwnProperty('agentVersion')) {
                obj['agentVersion'] = ApiClient.convertToType(data['agentVersion'], 'String');
            }
            if (data.hasOwnProperty('maxCores')) {
                obj['maxCores'] = ApiClient.convertToType(data['maxCores'], 'Number');
            }
            if (data.hasOwnProperty('coresPerSocket')) {
                obj['coresPerSocket'] = ApiClient.convertToType(data['coresPerSocket'], 'String');
            }
            if (data.hasOwnProperty('maxMemory')) {
                obj['maxMemory'] = ApiClient.convertToType(data['maxMemory'], 'Number');
            }
            if (data.hasOwnProperty('maxStorage')) {
                obj['maxStorage'] = ApiClient.convertToType(data['maxStorage'], 'Number');
            }
            if (data.hasOwnProperty('maxCpu')) {
                obj['maxCpu'] = ApiClient.convertToType(data['maxCpu'], 'Number');
            }
            if (data.hasOwnProperty('hourlyPrice')) {
                obj['hourlyPrice'] = ApiClient.convertToType(data['hourlyPrice'], 'Number');
            }
            if (data.hasOwnProperty('sourceImage')) {
                obj['sourceImage'] = InlineResponse20079LoadBalancerMonitorLoadBalancerType.constructFromObject(data['sourceImage']);
            }
            if (data.hasOwnProperty('serverOs')) {
                obj['serverOs'] = ApiClient.convertToType(data['serverOs'], 'String');
            }
            if (data.hasOwnProperty('volumes')) {
                obj['volumes'] = ApiClient.convertToType(data['volumes'], [ClusterMastersVolumes]);
            }
            if (data.hasOwnProperty('controllers')) {
                obj['controllers'] = ApiClient.convertToType(data['controllers'], [Object]);
            }
            if (data.hasOwnProperty('interfaces')) {
                obj['interfaces'] = ApiClient.convertToType(data['interfaces'], [ClusterMastersInterfaces]);
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], [ApiServersIdMakeManagedServerTags]);
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('tagCompliant')) {
                obj['tagCompliant'] = ApiClient.convertToType(data['tagCompliant'], 'String');
            }
            if (data.hasOwnProperty('containers')) {
                obj['containers'] = ApiClient.convertToType(data['containers'], ['Number']);
            }
            if (data.hasOwnProperty('guestConsolePreferred')) {
                obj['guestConsolePreferred'] = ApiClient.convertToType(data['guestConsolePreferred'], 'Boolean');
            }
            if (data.hasOwnProperty('guestConsoleType')) {
                obj['guestConsoleType'] = ApiClient.convertToType(data['guestConsoleType'], 'String');
            }
            if (data.hasOwnProperty('guestConsoleUsername')) {
                obj['guestConsoleUsername'] = ApiClient.convertToType(data['guestConsoleUsername'], 'String');
            }
            if (data.hasOwnProperty('guestConsolePassword')) {
                obj['guestConsolePassword'] = ApiClient.convertToType(data['guestConsolePassword'], 'String');
            }
            if (data.hasOwnProperty('guestConsolePasswordHash')) {
                obj['guestConsolePasswordHash'] = ApiClient.convertToType(data['guestConsolePasswordHash'], 'String');
            }
            if (data.hasOwnProperty('guestConsolePort')) {
                obj['guestConsolePort'] = ApiClient.convertToType(data['guestConsolePort'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ClusterMasters.prototype['id'] = undefined;

/**
 * @member {String} uuid
 */
ClusterMasters.prototype['uuid'] = undefined;

/**
 * @member {String} externalId
 */
ClusterMasters.prototype['externalId'] = undefined;

/**
 * @member {String} internalId
 */
ClusterMasters.prototype['internalId'] = undefined;

/**
 * @member {String} externalUniqueId
 */
ClusterMasters.prototype['externalUniqueId'] = undefined;

/**
 * @member {String} name
 */
ClusterMasters.prototype['name'] = undefined;

/**
 * @member {String} externalName
 */
ClusterMasters.prototype['externalName'] = undefined;

/**
 * @member {String} hostname
 */
ClusterMasters.prototype['hostname'] = undefined;

/**
 * @member {Number} accountId
 */
ClusterMasters.prototype['accountId'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} account
 */
ClusterMasters.prototype['account'] = undefined;

/**
 * @member {module:model/InlineResponse200107NetworkPoolCreatedBy} owner
 */
ClusterMasters.prototype['owner'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} zone
 */
ClusterMasters.prototype['zone'] = undefined;

/**
 * @member {module:model/InlineResponse20079LoadBalancerMonitorLoadBalancerType} plan
 */
ClusterMasters.prototype['plan'] = undefined;

/**
 * @member {module:model/ClusterLayoutComputeServerType} computeServerType
 */
ClusterMasters.prototype['computeServerType'] = undefined;

/**
 * @member {String} visibility
 */
ClusterMasters.prototype['visibility'] = undefined;

/**
 * @member {String} description
 */
ClusterMasters.prototype['description'] = undefined;

/**
 * @member {Number} zoneId
 */
ClusterMasters.prototype['zoneId'] = undefined;

/**
 * @member {Number} siteId
 */
ClusterMasters.prototype['siteId'] = undefined;

/**
 * @member {Number} resourcePoolId
 */
ClusterMasters.prototype['resourcePoolId'] = undefined;

/**
 * @member {String} folderId
 */
ClusterMasters.prototype['folderId'] = undefined;

/**
 * @member {String} sshHost
 */
ClusterMasters.prototype['sshHost'] = undefined;

/**
 * @member {Number} sshPort
 */
ClusterMasters.prototype['sshPort'] = undefined;

/**
 * @member {String} externalIp
 */
ClusterMasters.prototype['externalIp'] = undefined;

/**
 * @member {String} internalIp
 */
ClusterMasters.prototype['internalIp'] = undefined;

/**
 * @member {String} volumeId
 */
ClusterMasters.prototype['volumeId'] = undefined;

/**
 * @member {String} platform
 */
ClusterMasters.prototype['platform'] = undefined;

/**
 * @member {String} platformVersion
 */
ClusterMasters.prototype['platformVersion'] = undefined;

/**
 * @member {String} sshUsername
 */
ClusterMasters.prototype['sshUsername'] = undefined;

/**
 * @member {String} sshPassword
 */
ClusterMasters.prototype['sshPassword'] = undefined;

/**
 * @member {String} sshPasswordHash
 */
ClusterMasters.prototype['sshPasswordHash'] = undefined;

/**
 * @member {String} osDevice
 */
ClusterMasters.prototype['osDevice'] = undefined;

/**
 * @member {String} osType
 */
ClusterMasters.prototype['osType'] = undefined;

/**
 * @member {String} dataDevice
 */
ClusterMasters.prototype['dataDevice'] = undefined;

/**
 * @member {Boolean} lvmEnabled
 */
ClusterMasters.prototype['lvmEnabled'] = undefined;

/**
 * @member {String} apiKey
 */
ClusterMasters.prototype['apiKey'] = undefined;

/**
 * @member {Boolean} softwareRaid
 */
ClusterMasters.prototype['softwareRaid'] = undefined;

/**
 * @member {Date} dateCreated
 */
ClusterMasters.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
ClusterMasters.prototype['lastUpdated'] = undefined;

/**
 * @member {module:model/ClusterMastersStats} stats
 */
ClusterMasters.prototype['stats'] = undefined;

/**
 * @member {String} status
 */
ClusterMasters.prototype['status'] = undefined;

/**
 * @member {String} statusMessage
 */
ClusterMasters.prototype['statusMessage'] = undefined;

/**
 * @member {String} errorMessage
 */
ClusterMasters.prototype['errorMessage'] = undefined;

/**
 * @member {String} statusDate
 */
ClusterMasters.prototype['statusDate'] = undefined;

/**
 * @member {String} statusPercent
 */
ClusterMasters.prototype['statusPercent'] = undefined;

/**
 * @member {String} statusEta
 */
ClusterMasters.prototype['statusEta'] = undefined;

/**
 * @member {String} powerState
 */
ClusterMasters.prototype['powerState'] = undefined;

/**
 * @member {Boolean} agentInstalled
 */
ClusterMasters.prototype['agentInstalled'] = undefined;

/**
 * @member {Date} lastAgentUpdate
 */
ClusterMasters.prototype['lastAgentUpdate'] = undefined;

/**
 * @member {String} agentVersion
 */
ClusterMasters.prototype['agentVersion'] = undefined;

/**
 * @member {Number} maxCores
 */
ClusterMasters.prototype['maxCores'] = undefined;

/**
 * @member {String} coresPerSocket
 */
ClusterMasters.prototype['coresPerSocket'] = undefined;

/**
 * @member {Number} maxMemory
 */
ClusterMasters.prototype['maxMemory'] = undefined;

/**
 * @member {Number} maxStorage
 */
ClusterMasters.prototype['maxStorage'] = undefined;

/**
 * @member {Number} maxCpu
 */
ClusterMasters.prototype['maxCpu'] = undefined;

/**
 * @member {Number} hourlyPrice
 */
ClusterMasters.prototype['hourlyPrice'] = undefined;

/**
 * @member {module:model/InlineResponse20079LoadBalancerMonitorLoadBalancerType} sourceImage
 */
ClusterMasters.prototype['sourceImage'] = undefined;

/**
 * @member {String} serverOs
 */
ClusterMasters.prototype['serverOs'] = undefined;

/**
 * @member {Array.<module:model/ClusterMastersVolumes>} volumes
 */
ClusterMasters.prototype['volumes'] = undefined;

/**
 * @member {Array.<Object>} controllers
 */
ClusterMasters.prototype['controllers'] = undefined;

/**
 * @member {Array.<module:model/ClusterMastersInterfaces>} interfaces
 */
ClusterMasters.prototype['interfaces'] = undefined;

/**
 * @member {Array.<String>} labels
 */
ClusterMasters.prototype['labels'] = undefined;

/**
 * @member {Array.<module:model/ApiServersIdMakeManagedServerTags>} tags
 */
ClusterMasters.prototype['tags'] = undefined;

/**
 * @member {Boolean} enabled
 */
ClusterMasters.prototype['enabled'] = undefined;

/**
 * @member {String} tagCompliant
 */
ClusterMasters.prototype['tagCompliant'] = undefined;

/**
 * @member {Array.<Number>} containers
 */
ClusterMasters.prototype['containers'] = undefined;

/**
 * @member {Boolean} guestConsolePreferred
 */
ClusterMasters.prototype['guestConsolePreferred'] = undefined;

/**
 * @member {String} guestConsoleType
 */
ClusterMasters.prototype['guestConsoleType'] = undefined;

/**
 * @member {String} guestConsoleUsername
 */
ClusterMasters.prototype['guestConsoleUsername'] = undefined;

/**
 * @member {String} guestConsolePassword
 */
ClusterMasters.prototype['guestConsolePassword'] = undefined;

/**
 * @member {String} guestConsolePasswordHash
 */
ClusterMasters.prototype['guestConsolePasswordHash'] = undefined;

/**
 * @member {String} guestConsolePort
 */
ClusterMasters.prototype['guestConsolePort'] = undefined;






export default ClusterMasters;

