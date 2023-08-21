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
import ClusterLayoutComputeServerType from './ClusterLayoutComputeServerType';
import GuidanceVmwareSizingResourceControllers from './GuidanceVmwareSizingResourceControllers';
import GuidanceVmwareSizingResourceInterfaces from './GuidanceVmwareSizingResourceInterfaces';
import GuidanceVmwareSizingResourceServerOs from './GuidanceVmwareSizingResourceServerOs';
import GuidanceVmwareSizingResourceStats from './GuidanceVmwareSizingResourceStats';
import GuidanceVmwareSizingResourceVolumes from './GuidanceVmwareSizingResourceVolumes';
import InlineResponse200107NetworkPoolCreatedBy from './InlineResponse200107NetworkPoolCreatedBy';
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';
import InlineResponse20079LoadBalancerMonitorLoadBalancerType from './InlineResponse20079LoadBalancerMonitorLoadBalancerType';

/**
 * The GuidanceVmwareSizingResource model module.
 * @module model/GuidanceVmwareSizingResource
 * @version 6.2.1
 */
class GuidanceVmwareSizingResource {
    /**
     * Constructs a new <code>GuidanceVmwareSizingResource</code>.
     * @alias module:model/GuidanceVmwareSizingResource
     */
    constructor() { 
        
        GuidanceVmwareSizingResource.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>GuidanceVmwareSizingResource</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/GuidanceVmwareSizingResource} obj Optional instance to populate.
     * @return {module:model/GuidanceVmwareSizingResource} The populated <code>GuidanceVmwareSizingResource</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new GuidanceVmwareSizingResource();

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
            if (data.hasOwnProperty('parentServer')) {
                obj['parentServer'] = InlineResponse20040AppDeployInstance.constructFromObject(data['parentServer']);
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
                obj['folderId'] = ApiClient.convertToType(data['folderId'], 'Number');
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
                obj['stats'] = GuidanceVmwareSizingResourceStats.constructFromObject(data['stats']);
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
                obj['statusDate'] = ApiClient.convertToType(data['statusDate'], 'Date');
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
                obj['coresPerSocket'] = ApiClient.convertToType(data['coresPerSocket'], 'Number');
            }
            if (data.hasOwnProperty('maxMemory')) {
                obj['maxMemory'] = ApiClient.convertToType(data['maxMemory'], 'Number');
            }
            if (data.hasOwnProperty('maxStorage')) {
                obj['maxStorage'] = ApiClient.convertToType(data['maxStorage'], 'Number');
            }
            if (data.hasOwnProperty('maxCpu')) {
                obj['maxCpu'] = ApiClient.convertToType(data['maxCpu'], 'String');
            }
            if (data.hasOwnProperty('hourlyPrice')) {
                obj['hourlyPrice'] = ApiClient.convertToType(data['hourlyPrice'], 'Number');
            }
            if (data.hasOwnProperty('sourceImage')) {
                obj['sourceImage'] = InlineResponse20079LoadBalancerMonitorLoadBalancerType.constructFromObject(data['sourceImage']);
            }
            if (data.hasOwnProperty('serverOs')) {
                obj['serverOs'] = GuidanceVmwareSizingResourceServerOs.constructFromObject(data['serverOs']);
            }
            if (data.hasOwnProperty('volumes')) {
                obj['volumes'] = ApiClient.convertToType(data['volumes'], [GuidanceVmwareSizingResourceVolumes]);
            }
            if (data.hasOwnProperty('controllers')) {
                obj['controllers'] = ApiClient.convertToType(data['controllers'], [GuidanceVmwareSizingResourceControllers]);
            }
            if (data.hasOwnProperty('interfaces')) {
                obj['interfaces'] = ApiClient.convertToType(data['interfaces'], [GuidanceVmwareSizingResourceInterfaces]);
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], [Object]);
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], [Object]);
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
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
GuidanceVmwareSizingResource.prototype['id'] = undefined;

/**
 * @member {String} uuid
 */
GuidanceVmwareSizingResource.prototype['uuid'] = undefined;

/**
 * @member {String} externalId
 */
GuidanceVmwareSizingResource.prototype['externalId'] = undefined;

/**
 * @member {String} internalId
 */
GuidanceVmwareSizingResource.prototype['internalId'] = undefined;

/**
 * @member {String} externalUniqueId
 */
GuidanceVmwareSizingResource.prototype['externalUniqueId'] = undefined;

/**
 * @member {String} name
 */
GuidanceVmwareSizingResource.prototype['name'] = undefined;

/**
 * @member {String} externalName
 */
GuidanceVmwareSizingResource.prototype['externalName'] = undefined;

/**
 * @member {String} hostname
 */
GuidanceVmwareSizingResource.prototype['hostname'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} parentServer
 */
GuidanceVmwareSizingResource.prototype['parentServer'] = undefined;

/**
 * @member {Number} accountId
 */
GuidanceVmwareSizingResource.prototype['accountId'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} account
 */
GuidanceVmwareSizingResource.prototype['account'] = undefined;

/**
 * @member {module:model/InlineResponse200107NetworkPoolCreatedBy} owner
 */
GuidanceVmwareSizingResource.prototype['owner'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} zone
 */
GuidanceVmwareSizingResource.prototype['zone'] = undefined;

/**
 * @member {module:model/InlineResponse20079LoadBalancerMonitorLoadBalancerType} plan
 */
GuidanceVmwareSizingResource.prototype['plan'] = undefined;

/**
 * @member {module:model/ClusterLayoutComputeServerType} computeServerType
 */
GuidanceVmwareSizingResource.prototype['computeServerType'] = undefined;

/**
 * @member {String} visibility
 */
GuidanceVmwareSizingResource.prototype['visibility'] = undefined;

/**
 * @member {String} description
 */
GuidanceVmwareSizingResource.prototype['description'] = undefined;

/**
 * @member {Number} zoneId
 */
GuidanceVmwareSizingResource.prototype['zoneId'] = undefined;

/**
 * @member {Number} siteId
 */
GuidanceVmwareSizingResource.prototype['siteId'] = undefined;

/**
 * @member {Number} resourcePoolId
 */
GuidanceVmwareSizingResource.prototype['resourcePoolId'] = undefined;

/**
 * @member {Number} folderId
 */
GuidanceVmwareSizingResource.prototype['folderId'] = undefined;

/**
 * @member {String} sshHost
 */
GuidanceVmwareSizingResource.prototype['sshHost'] = undefined;

/**
 * @member {Number} sshPort
 */
GuidanceVmwareSizingResource.prototype['sshPort'] = undefined;

/**
 * @member {String} externalIp
 */
GuidanceVmwareSizingResource.prototype['externalIp'] = undefined;

/**
 * @member {String} internalIp
 */
GuidanceVmwareSizingResource.prototype['internalIp'] = undefined;

/**
 * @member {String} volumeId
 */
GuidanceVmwareSizingResource.prototype['volumeId'] = undefined;

/**
 * @member {String} platform
 */
GuidanceVmwareSizingResource.prototype['platform'] = undefined;

/**
 * @member {String} platformVersion
 */
GuidanceVmwareSizingResource.prototype['platformVersion'] = undefined;

/**
 * @member {String} sshUsername
 */
GuidanceVmwareSizingResource.prototype['sshUsername'] = undefined;

/**
 * @member {String} sshPassword
 */
GuidanceVmwareSizingResource.prototype['sshPassword'] = undefined;

/**
 * @member {String} sshPasswordHash
 */
GuidanceVmwareSizingResource.prototype['sshPasswordHash'] = undefined;

/**
 * @member {String} osDevice
 */
GuidanceVmwareSizingResource.prototype['osDevice'] = undefined;

/**
 * @member {String} osType
 */
GuidanceVmwareSizingResource.prototype['osType'] = undefined;

/**
 * @member {String} dataDevice
 */
GuidanceVmwareSizingResource.prototype['dataDevice'] = undefined;

/**
 * @member {Boolean} lvmEnabled
 */
GuidanceVmwareSizingResource.prototype['lvmEnabled'] = undefined;

/**
 * @member {String} apiKey
 */
GuidanceVmwareSizingResource.prototype['apiKey'] = undefined;

/**
 * @member {Boolean} softwareRaid
 */
GuidanceVmwareSizingResource.prototype['softwareRaid'] = undefined;

/**
 * @member {Date} dateCreated
 */
GuidanceVmwareSizingResource.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
GuidanceVmwareSizingResource.prototype['lastUpdated'] = undefined;

/**
 * @member {module:model/GuidanceVmwareSizingResourceStats} stats
 */
GuidanceVmwareSizingResource.prototype['stats'] = undefined;

/**
 * @member {String} status
 */
GuidanceVmwareSizingResource.prototype['status'] = undefined;

/**
 * @member {String} statusMessage
 */
GuidanceVmwareSizingResource.prototype['statusMessage'] = undefined;

/**
 * @member {String} errorMessage
 */
GuidanceVmwareSizingResource.prototype['errorMessage'] = undefined;

/**
 * @member {Date} statusDate
 */
GuidanceVmwareSizingResource.prototype['statusDate'] = undefined;

/**
 * @member {String} statusPercent
 */
GuidanceVmwareSizingResource.prototype['statusPercent'] = undefined;

/**
 * @member {String} statusEta
 */
GuidanceVmwareSizingResource.prototype['statusEta'] = undefined;

/**
 * @member {String} powerState
 */
GuidanceVmwareSizingResource.prototype['powerState'] = undefined;

/**
 * @member {Boolean} agentInstalled
 */
GuidanceVmwareSizingResource.prototype['agentInstalled'] = undefined;

/**
 * @member {Date} lastAgentUpdate
 */
GuidanceVmwareSizingResource.prototype['lastAgentUpdate'] = undefined;

/**
 * @member {String} agentVersion
 */
GuidanceVmwareSizingResource.prototype['agentVersion'] = undefined;

/**
 * @member {Number} maxCores
 */
GuidanceVmwareSizingResource.prototype['maxCores'] = undefined;

/**
 * @member {Number} coresPerSocket
 */
GuidanceVmwareSizingResource.prototype['coresPerSocket'] = undefined;

/**
 * @member {Number} maxMemory
 */
GuidanceVmwareSizingResource.prototype['maxMemory'] = undefined;

/**
 * @member {Number} maxStorage
 */
GuidanceVmwareSizingResource.prototype['maxStorage'] = undefined;

/**
 * @member {String} maxCpu
 */
GuidanceVmwareSizingResource.prototype['maxCpu'] = undefined;

/**
 * @member {Number} hourlyPrice
 */
GuidanceVmwareSizingResource.prototype['hourlyPrice'] = undefined;

/**
 * @member {module:model/InlineResponse20079LoadBalancerMonitorLoadBalancerType} sourceImage
 */
GuidanceVmwareSizingResource.prototype['sourceImage'] = undefined;

/**
 * @member {module:model/GuidanceVmwareSizingResourceServerOs} serverOs
 */
GuidanceVmwareSizingResource.prototype['serverOs'] = undefined;

/**
 * @member {Array.<module:model/GuidanceVmwareSizingResourceVolumes>} volumes
 */
GuidanceVmwareSizingResource.prototype['volumes'] = undefined;

/**
 * @member {Array.<module:model/GuidanceVmwareSizingResourceControllers>} controllers
 */
GuidanceVmwareSizingResource.prototype['controllers'] = undefined;

/**
 * @member {Array.<module:model/GuidanceVmwareSizingResourceInterfaces>} interfaces
 */
GuidanceVmwareSizingResource.prototype['interfaces'] = undefined;

/**
 * @member {Array.<Object>} labels
 */
GuidanceVmwareSizingResource.prototype['labels'] = undefined;

/**
 * @member {Array.<Object>} tags
 */
GuidanceVmwareSizingResource.prototype['tags'] = undefined;

/**
 * @member {Boolean} enabled
 */
GuidanceVmwareSizingResource.prototype['enabled'] = undefined;

/**
 * @member {String} tagCompliant
 */
GuidanceVmwareSizingResource.prototype['tagCompliant'] = undefined;

/**
 * @member {Array.<Number>} containers
 */
GuidanceVmwareSizingResource.prototype['containers'] = undefined;






export default GuidanceVmwareSizingResource;

