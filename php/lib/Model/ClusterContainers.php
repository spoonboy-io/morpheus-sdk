<?php
/**
 * ClusterContainers
 *
 * PHP version 7.2
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.0.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace OpenAPI\Client\Model;

use \ArrayAccess;
use \OpenAPI\Client\ObjectSerializer;

/**
 * ClusterContainers Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ClusterContainers implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'clusterContainers';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'uuid' => 'string',
        'account_id' => 'int',
        'instance' => 'string',
        'container_type' => '\OpenAPI\Client\Model\ClusterContainersContainerType',
        'container_type_set' => '\OpenAPI\Client\Model\ClusterContainersContainerTypeSet',
        'server' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance',
        'cloud' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance',
        'name' => 'string',
        'ip' => 'string',
        'internal_ip' => 'string',
        'internal_hostname' => 'string',
        'external_hostname' => 'string',
        'external_domain' => 'string',
        'external_fqdn' => 'string',
        'ports' => 'object[]',
        'plan' => '\OpenAPI\Client\Model\ClusterContainersPlan',
        'date_created' => '\DateTime',
        'last_updated' => '\DateTime',
        'stats_enabled' => 'bool',
        'status' => 'string',
        'user_status' => 'string',
        'environment_prefix' => 'string',
        'config_group' => 'string',
        'config_id' => 'string',
        'config_role' => 'string',
        'stats' => '\OpenAPI\Client\Model\ClusterContainersStats',
        'runtime_info' => 'object',
        'container_version' => 'string',
        'repository_image' => 'string',
        'plan_category' => 'string',
        'hostname' => 'string',
        'domain_name' => 'string',
        'volume_created' => 'bool',
        'container_created' => 'bool',
        'max_storage' => 'string',
        'max_memory' => 'string',
        'max_cores' => 'string',
        'max_cpu' => 'string',
        'hourly_price' => 'float',
        'available_actions' => '\OpenAPI\Client\Model\ClusterContainersAvailableActions[]'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'id' => 'int64',
        'uuid' => null,
        'account_id' => 'int64',
        'instance' => null,
        'container_type' => null,
        'container_type_set' => null,
        'server' => null,
        'cloud' => null,
        'name' => null,
        'ip' => null,
        'internal_ip' => null,
        'internal_hostname' => null,
        'external_hostname' => null,
        'external_domain' => null,
        'external_fqdn' => null,
        'ports' => null,
        'plan' => null,
        'date_created' => 'date-time',
        'last_updated' => 'date-time',
        'stats_enabled' => null,
        'status' => null,
        'user_status' => null,
        'environment_prefix' => null,
        'config_group' => null,
        'config_id' => null,
        'config_role' => null,
        'stats' => null,
        'runtime_info' => null,
        'container_version' => null,
        'repository_image' => null,
        'plan_category' => null,
        'hostname' => null,
        'domain_name' => null,
        'volume_created' => null,
        'container_created' => null,
        'max_storage' => null,
        'max_memory' => null,
        'max_cores' => null,
        'max_cpu' => null,
        'hourly_price' => null,
        'available_actions' => null
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'id' => 'id',
        'uuid' => 'uuid',
        'account_id' => 'accountId',
        'instance' => 'instance',
        'container_type' => 'containerType',
        'container_type_set' => 'containerTypeSet',
        'server' => 'server',
        'cloud' => 'cloud',
        'name' => 'name',
        'ip' => 'ip',
        'internal_ip' => 'internalIp',
        'internal_hostname' => 'internalHostname',
        'external_hostname' => 'externalHostname',
        'external_domain' => 'externalDomain',
        'external_fqdn' => 'externalFqdn',
        'ports' => 'ports',
        'plan' => 'plan',
        'date_created' => 'dateCreated',
        'last_updated' => 'lastUpdated',
        'stats_enabled' => 'statsEnabled',
        'status' => 'status',
        'user_status' => 'userStatus',
        'environment_prefix' => 'environmentPrefix',
        'config_group' => 'configGroup',
        'config_id' => 'configId',
        'config_role' => 'configRole',
        'stats' => 'stats',
        'runtime_info' => 'runtimeInfo',
        'container_version' => 'containerVersion',
        'repository_image' => 'repositoryImage',
        'plan_category' => 'planCategory',
        'hostname' => 'hostname',
        'domain_name' => 'domainName',
        'volume_created' => 'volumeCreated',
        'container_created' => 'containerCreated',
        'max_storage' => 'maxStorage',
        'max_memory' => 'maxMemory',
        'max_cores' => 'maxCores',
        'max_cpu' => 'maxCpu',
        'hourly_price' => 'hourlyPrice',
        'available_actions' => 'availableActions'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'uuid' => 'setUuid',
        'account_id' => 'setAccountId',
        'instance' => 'setInstance',
        'container_type' => 'setContainerType',
        'container_type_set' => 'setContainerTypeSet',
        'server' => 'setServer',
        'cloud' => 'setCloud',
        'name' => 'setName',
        'ip' => 'setIp',
        'internal_ip' => 'setInternalIp',
        'internal_hostname' => 'setInternalHostname',
        'external_hostname' => 'setExternalHostname',
        'external_domain' => 'setExternalDomain',
        'external_fqdn' => 'setExternalFqdn',
        'ports' => 'setPorts',
        'plan' => 'setPlan',
        'date_created' => 'setDateCreated',
        'last_updated' => 'setLastUpdated',
        'stats_enabled' => 'setStatsEnabled',
        'status' => 'setStatus',
        'user_status' => 'setUserStatus',
        'environment_prefix' => 'setEnvironmentPrefix',
        'config_group' => 'setConfigGroup',
        'config_id' => 'setConfigId',
        'config_role' => 'setConfigRole',
        'stats' => 'setStats',
        'runtime_info' => 'setRuntimeInfo',
        'container_version' => 'setContainerVersion',
        'repository_image' => 'setRepositoryImage',
        'plan_category' => 'setPlanCategory',
        'hostname' => 'setHostname',
        'domain_name' => 'setDomainName',
        'volume_created' => 'setVolumeCreated',
        'container_created' => 'setContainerCreated',
        'max_storage' => 'setMaxStorage',
        'max_memory' => 'setMaxMemory',
        'max_cores' => 'setMaxCores',
        'max_cpu' => 'setMaxCpu',
        'hourly_price' => 'setHourlyPrice',
        'available_actions' => 'setAvailableActions'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'uuid' => 'getUuid',
        'account_id' => 'getAccountId',
        'instance' => 'getInstance',
        'container_type' => 'getContainerType',
        'container_type_set' => 'getContainerTypeSet',
        'server' => 'getServer',
        'cloud' => 'getCloud',
        'name' => 'getName',
        'ip' => 'getIp',
        'internal_ip' => 'getInternalIp',
        'internal_hostname' => 'getInternalHostname',
        'external_hostname' => 'getExternalHostname',
        'external_domain' => 'getExternalDomain',
        'external_fqdn' => 'getExternalFqdn',
        'ports' => 'getPorts',
        'plan' => 'getPlan',
        'date_created' => 'getDateCreated',
        'last_updated' => 'getLastUpdated',
        'stats_enabled' => 'getStatsEnabled',
        'status' => 'getStatus',
        'user_status' => 'getUserStatus',
        'environment_prefix' => 'getEnvironmentPrefix',
        'config_group' => 'getConfigGroup',
        'config_id' => 'getConfigId',
        'config_role' => 'getConfigRole',
        'stats' => 'getStats',
        'runtime_info' => 'getRuntimeInfo',
        'container_version' => 'getContainerVersion',
        'repository_image' => 'getRepositoryImage',
        'plan_category' => 'getPlanCategory',
        'hostname' => 'getHostname',
        'domain_name' => 'getDomainName',
        'volume_created' => 'getVolumeCreated',
        'container_created' => 'getContainerCreated',
        'max_storage' => 'getMaxStorage',
        'max_memory' => 'getMaxMemory',
        'max_cores' => 'getMaxCores',
        'max_cpu' => 'getMaxCpu',
        'hourly_price' => 'getHourlyPrice',
        'available_actions' => 'getAvailableActions'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }

    

    

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['id'] = $data['id'] ?? null;
        $this->container['uuid'] = $data['uuid'] ?? null;
        $this->container['account_id'] = $data['account_id'] ?? null;
        $this->container['instance'] = $data['instance'] ?? null;
        $this->container['container_type'] = $data['container_type'] ?? null;
        $this->container['container_type_set'] = $data['container_type_set'] ?? null;
        $this->container['server'] = $data['server'] ?? null;
        $this->container['cloud'] = $data['cloud'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['ip'] = $data['ip'] ?? null;
        $this->container['internal_ip'] = $data['internal_ip'] ?? null;
        $this->container['internal_hostname'] = $data['internal_hostname'] ?? null;
        $this->container['external_hostname'] = $data['external_hostname'] ?? null;
        $this->container['external_domain'] = $data['external_domain'] ?? null;
        $this->container['external_fqdn'] = $data['external_fqdn'] ?? null;
        $this->container['ports'] = $data['ports'] ?? null;
        $this->container['plan'] = $data['plan'] ?? null;
        $this->container['date_created'] = $data['date_created'] ?? null;
        $this->container['last_updated'] = $data['last_updated'] ?? null;
        $this->container['stats_enabled'] = $data['stats_enabled'] ?? null;
        $this->container['status'] = $data['status'] ?? null;
        $this->container['user_status'] = $data['user_status'] ?? null;
        $this->container['environment_prefix'] = $data['environment_prefix'] ?? null;
        $this->container['config_group'] = $data['config_group'] ?? null;
        $this->container['config_id'] = $data['config_id'] ?? null;
        $this->container['config_role'] = $data['config_role'] ?? null;
        $this->container['stats'] = $data['stats'] ?? null;
        $this->container['runtime_info'] = $data['runtime_info'] ?? null;
        $this->container['container_version'] = $data['container_version'] ?? null;
        $this->container['repository_image'] = $data['repository_image'] ?? null;
        $this->container['plan_category'] = $data['plan_category'] ?? null;
        $this->container['hostname'] = $data['hostname'] ?? null;
        $this->container['domain_name'] = $data['domain_name'] ?? null;
        $this->container['volume_created'] = $data['volume_created'] ?? null;
        $this->container['container_created'] = $data['container_created'] ?? null;
        $this->container['max_storage'] = $data['max_storage'] ?? null;
        $this->container['max_memory'] = $data['max_memory'] ?? null;
        $this->container['max_cores'] = $data['max_cores'] ?? null;
        $this->container['max_cpu'] = $data['max_cpu'] ?? null;
        $this->container['hourly_price'] = $data['hourly_price'] ?? null;
        $this->container['available_actions'] = $data['available_actions'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets id
     *
     * @return int|null
     */
    public function getId()
    {
        return $this->container['id'];
    }

    /**
     * Sets id
     *
     * @param int|null $id id
     *
     * @return self
     */
    public function setId($id)
    {
        $this->container['id'] = $id;

        return $this;
    }

    /**
     * Gets uuid
     *
     * @return string|null
     */
    public function getUuid()
    {
        return $this->container['uuid'];
    }

    /**
     * Sets uuid
     *
     * @param string|null $uuid uuid
     *
     * @return self
     */
    public function setUuid($uuid)
    {
        $this->container['uuid'] = $uuid;

        return $this;
    }

    /**
     * Gets account_id
     *
     * @return int|null
     */
    public function getAccountId()
    {
        return $this->container['account_id'];
    }

    /**
     * Sets account_id
     *
     * @param int|null $account_id account_id
     *
     * @return self
     */
    public function setAccountId($account_id)
    {
        $this->container['account_id'] = $account_id;

        return $this;
    }

    /**
     * Gets instance
     *
     * @return string|null
     */
    public function getInstance()
    {
        return $this->container['instance'];
    }

    /**
     * Sets instance
     *
     * @param string|null $instance instance
     *
     * @return self
     */
    public function setInstance($instance)
    {
        $this->container['instance'] = $instance;

        return $this;
    }

    /**
     * Gets container_type
     *
     * @return \OpenAPI\Client\Model\ClusterContainersContainerType|null
     */
    public function getContainerType()
    {
        return $this->container['container_type'];
    }

    /**
     * Sets container_type
     *
     * @param \OpenAPI\Client\Model\ClusterContainersContainerType|null $container_type container_type
     *
     * @return self
     */
    public function setContainerType($container_type)
    {
        $this->container['container_type'] = $container_type;

        return $this;
    }

    /**
     * Gets container_type_set
     *
     * @return \OpenAPI\Client\Model\ClusterContainersContainerTypeSet|null
     */
    public function getContainerTypeSet()
    {
        return $this->container['container_type_set'];
    }

    /**
     * Sets container_type_set
     *
     * @param \OpenAPI\Client\Model\ClusterContainersContainerTypeSet|null $container_type_set container_type_set
     *
     * @return self
     */
    public function setContainerTypeSet($container_type_set)
    {
        $this->container['container_type_set'] = $container_type_set;

        return $this;
    }

    /**
     * Gets server
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null
     */
    public function getServer()
    {
        return $this->container['server'];
    }

    /**
     * Sets server
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null $server server
     *
     * @return self
     */
    public function setServer($server)
    {
        $this->container['server'] = $server;

        return $this;
    }

    /**
     * Gets cloud
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null
     */
    public function getCloud()
    {
        return $this->container['cloud'];
    }

    /**
     * Sets cloud
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null $cloud cloud
     *
     * @return self
     */
    public function setCloud($cloud)
    {
        $this->container['cloud'] = $cloud;

        return $this;
    }

    /**
     * Gets name
     *
     * @return string|null
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     *
     * @param string|null $name name
     *
     * @return self
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets ip
     *
     * @return string|null
     */
    public function getIp()
    {
        return $this->container['ip'];
    }

    /**
     * Sets ip
     *
     * @param string|null $ip ip
     *
     * @return self
     */
    public function setIp($ip)
    {
        $this->container['ip'] = $ip;

        return $this;
    }

    /**
     * Gets internal_ip
     *
     * @return string|null
     */
    public function getInternalIp()
    {
        return $this->container['internal_ip'];
    }

    /**
     * Sets internal_ip
     *
     * @param string|null $internal_ip internal_ip
     *
     * @return self
     */
    public function setInternalIp($internal_ip)
    {
        $this->container['internal_ip'] = $internal_ip;

        return $this;
    }

    /**
     * Gets internal_hostname
     *
     * @return string|null
     */
    public function getInternalHostname()
    {
        return $this->container['internal_hostname'];
    }

    /**
     * Sets internal_hostname
     *
     * @param string|null $internal_hostname internal_hostname
     *
     * @return self
     */
    public function setInternalHostname($internal_hostname)
    {
        $this->container['internal_hostname'] = $internal_hostname;

        return $this;
    }

    /**
     * Gets external_hostname
     *
     * @return string|null
     */
    public function getExternalHostname()
    {
        return $this->container['external_hostname'];
    }

    /**
     * Sets external_hostname
     *
     * @param string|null $external_hostname external_hostname
     *
     * @return self
     */
    public function setExternalHostname($external_hostname)
    {
        $this->container['external_hostname'] = $external_hostname;

        return $this;
    }

    /**
     * Gets external_domain
     *
     * @return string|null
     */
    public function getExternalDomain()
    {
        return $this->container['external_domain'];
    }

    /**
     * Sets external_domain
     *
     * @param string|null $external_domain external_domain
     *
     * @return self
     */
    public function setExternalDomain($external_domain)
    {
        $this->container['external_domain'] = $external_domain;

        return $this;
    }

    /**
     * Gets external_fqdn
     *
     * @return string|null
     */
    public function getExternalFqdn()
    {
        return $this->container['external_fqdn'];
    }

    /**
     * Sets external_fqdn
     *
     * @param string|null $external_fqdn external_fqdn
     *
     * @return self
     */
    public function setExternalFqdn($external_fqdn)
    {
        $this->container['external_fqdn'] = $external_fqdn;

        return $this;
    }

    /**
     * Gets ports
     *
     * @return object[]|null
     */
    public function getPorts()
    {
        return $this->container['ports'];
    }

    /**
     * Sets ports
     *
     * @param object[]|null $ports ports
     *
     * @return self
     */
    public function setPorts($ports)
    {
        $this->container['ports'] = $ports;

        return $this;
    }

    /**
     * Gets plan
     *
     * @return \OpenAPI\Client\Model\ClusterContainersPlan|null
     */
    public function getPlan()
    {
        return $this->container['plan'];
    }

    /**
     * Sets plan
     *
     * @param \OpenAPI\Client\Model\ClusterContainersPlan|null $plan plan
     *
     * @return self
     */
    public function setPlan($plan)
    {
        $this->container['plan'] = $plan;

        return $this;
    }

    /**
     * Gets date_created
     *
     * @return \DateTime|null
     */
    public function getDateCreated()
    {
        return $this->container['date_created'];
    }

    /**
     * Sets date_created
     *
     * @param \DateTime|null $date_created date_created
     *
     * @return self
     */
    public function setDateCreated($date_created)
    {
        $this->container['date_created'] = $date_created;

        return $this;
    }

    /**
     * Gets last_updated
     *
     * @return \DateTime|null
     */
    public function getLastUpdated()
    {
        return $this->container['last_updated'];
    }

    /**
     * Sets last_updated
     *
     * @param \DateTime|null $last_updated last_updated
     *
     * @return self
     */
    public function setLastUpdated($last_updated)
    {
        $this->container['last_updated'] = $last_updated;

        return $this;
    }

    /**
     * Gets stats_enabled
     *
     * @return bool|null
     */
    public function getStatsEnabled()
    {
        return $this->container['stats_enabled'];
    }

    /**
     * Sets stats_enabled
     *
     * @param bool|null $stats_enabled stats_enabled
     *
     * @return self
     */
    public function setStatsEnabled($stats_enabled)
    {
        $this->container['stats_enabled'] = $stats_enabled;

        return $this;
    }

    /**
     * Gets status
     *
     * @return string|null
     */
    public function getStatus()
    {
        return $this->container['status'];
    }

    /**
     * Sets status
     *
     * @param string|null $status status
     *
     * @return self
     */
    public function setStatus($status)
    {
        $this->container['status'] = $status;

        return $this;
    }

    /**
     * Gets user_status
     *
     * @return string|null
     */
    public function getUserStatus()
    {
        return $this->container['user_status'];
    }

    /**
     * Sets user_status
     *
     * @param string|null $user_status user_status
     *
     * @return self
     */
    public function setUserStatus($user_status)
    {
        $this->container['user_status'] = $user_status;

        return $this;
    }

    /**
     * Gets environment_prefix
     *
     * @return string|null
     */
    public function getEnvironmentPrefix()
    {
        return $this->container['environment_prefix'];
    }

    /**
     * Sets environment_prefix
     *
     * @param string|null $environment_prefix environment_prefix
     *
     * @return self
     */
    public function setEnvironmentPrefix($environment_prefix)
    {
        $this->container['environment_prefix'] = $environment_prefix;

        return $this;
    }

    /**
     * Gets config_group
     *
     * @return string|null
     */
    public function getConfigGroup()
    {
        return $this->container['config_group'];
    }

    /**
     * Sets config_group
     *
     * @param string|null $config_group config_group
     *
     * @return self
     */
    public function setConfigGroup($config_group)
    {
        $this->container['config_group'] = $config_group;

        return $this;
    }

    /**
     * Gets config_id
     *
     * @return string|null
     */
    public function getConfigId()
    {
        return $this->container['config_id'];
    }

    /**
     * Sets config_id
     *
     * @param string|null $config_id config_id
     *
     * @return self
     */
    public function setConfigId($config_id)
    {
        $this->container['config_id'] = $config_id;

        return $this;
    }

    /**
     * Gets config_role
     *
     * @return string|null
     */
    public function getConfigRole()
    {
        return $this->container['config_role'];
    }

    /**
     * Sets config_role
     *
     * @param string|null $config_role config_role
     *
     * @return self
     */
    public function setConfigRole($config_role)
    {
        $this->container['config_role'] = $config_role;

        return $this;
    }

    /**
     * Gets stats
     *
     * @return \OpenAPI\Client\Model\ClusterContainersStats|null
     */
    public function getStats()
    {
        return $this->container['stats'];
    }

    /**
     * Sets stats
     *
     * @param \OpenAPI\Client\Model\ClusterContainersStats|null $stats stats
     *
     * @return self
     */
    public function setStats($stats)
    {
        $this->container['stats'] = $stats;

        return $this;
    }

    /**
     * Gets runtime_info
     *
     * @return object|null
     */
    public function getRuntimeInfo()
    {
        return $this->container['runtime_info'];
    }

    /**
     * Sets runtime_info
     *
     * @param object|null $runtime_info runtime_info
     *
     * @return self
     */
    public function setRuntimeInfo($runtime_info)
    {
        $this->container['runtime_info'] = $runtime_info;

        return $this;
    }

    /**
     * Gets container_version
     *
     * @return string|null
     */
    public function getContainerVersion()
    {
        return $this->container['container_version'];
    }

    /**
     * Sets container_version
     *
     * @param string|null $container_version container_version
     *
     * @return self
     */
    public function setContainerVersion($container_version)
    {
        $this->container['container_version'] = $container_version;

        return $this;
    }

    /**
     * Gets repository_image
     *
     * @return string|null
     */
    public function getRepositoryImage()
    {
        return $this->container['repository_image'];
    }

    /**
     * Sets repository_image
     *
     * @param string|null $repository_image repository_image
     *
     * @return self
     */
    public function setRepositoryImage($repository_image)
    {
        $this->container['repository_image'] = $repository_image;

        return $this;
    }

    /**
     * Gets plan_category
     *
     * @return string|null
     */
    public function getPlanCategory()
    {
        return $this->container['plan_category'];
    }

    /**
     * Sets plan_category
     *
     * @param string|null $plan_category plan_category
     *
     * @return self
     */
    public function setPlanCategory($plan_category)
    {
        $this->container['plan_category'] = $plan_category;

        return $this;
    }

    /**
     * Gets hostname
     *
     * @return string|null
     */
    public function getHostname()
    {
        return $this->container['hostname'];
    }

    /**
     * Sets hostname
     *
     * @param string|null $hostname hostname
     *
     * @return self
     */
    public function setHostname($hostname)
    {
        $this->container['hostname'] = $hostname;

        return $this;
    }

    /**
     * Gets domain_name
     *
     * @return string|null
     */
    public function getDomainName()
    {
        return $this->container['domain_name'];
    }

    /**
     * Sets domain_name
     *
     * @param string|null $domain_name domain_name
     *
     * @return self
     */
    public function setDomainName($domain_name)
    {
        $this->container['domain_name'] = $domain_name;

        return $this;
    }

    /**
     * Gets volume_created
     *
     * @return bool|null
     */
    public function getVolumeCreated()
    {
        return $this->container['volume_created'];
    }

    /**
     * Sets volume_created
     *
     * @param bool|null $volume_created volume_created
     *
     * @return self
     */
    public function setVolumeCreated($volume_created)
    {
        $this->container['volume_created'] = $volume_created;

        return $this;
    }

    /**
     * Gets container_created
     *
     * @return bool|null
     */
    public function getContainerCreated()
    {
        return $this->container['container_created'];
    }

    /**
     * Sets container_created
     *
     * @param bool|null $container_created container_created
     *
     * @return self
     */
    public function setContainerCreated($container_created)
    {
        $this->container['container_created'] = $container_created;

        return $this;
    }

    /**
     * Gets max_storage
     *
     * @return string|null
     */
    public function getMaxStorage()
    {
        return $this->container['max_storage'];
    }

    /**
     * Sets max_storage
     *
     * @param string|null $max_storage max_storage
     *
     * @return self
     */
    public function setMaxStorage($max_storage)
    {
        $this->container['max_storage'] = $max_storage;

        return $this;
    }

    /**
     * Gets max_memory
     *
     * @return string|null
     */
    public function getMaxMemory()
    {
        return $this->container['max_memory'];
    }

    /**
     * Sets max_memory
     *
     * @param string|null $max_memory max_memory
     *
     * @return self
     */
    public function setMaxMemory($max_memory)
    {
        $this->container['max_memory'] = $max_memory;

        return $this;
    }

    /**
     * Gets max_cores
     *
     * @return string|null
     */
    public function getMaxCores()
    {
        return $this->container['max_cores'];
    }

    /**
     * Sets max_cores
     *
     * @param string|null $max_cores max_cores
     *
     * @return self
     */
    public function setMaxCores($max_cores)
    {
        $this->container['max_cores'] = $max_cores;

        return $this;
    }

    /**
     * Gets max_cpu
     *
     * @return string|null
     */
    public function getMaxCpu()
    {
        return $this->container['max_cpu'];
    }

    /**
     * Sets max_cpu
     *
     * @param string|null $max_cpu max_cpu
     *
     * @return self
     */
    public function setMaxCpu($max_cpu)
    {
        $this->container['max_cpu'] = $max_cpu;

        return $this;
    }

    /**
     * Gets hourly_price
     *
     * @return float|null
     */
    public function getHourlyPrice()
    {
        return $this->container['hourly_price'];
    }

    /**
     * Sets hourly_price
     *
     * @param float|null $hourly_price hourly_price
     *
     * @return self
     */
    public function setHourlyPrice($hourly_price)
    {
        $this->container['hourly_price'] = $hourly_price;

        return $this;
    }

    /**
     * Gets available_actions
     *
     * @return \OpenAPI\Client\Model\ClusterContainersAvailableActions[]|null
     */
    public function getAvailableActions()
    {
        return $this->container['available_actions'];
    }

    /**
     * Sets available_actions
     *
     * @param \OpenAPI\Client\Model\ClusterContainersAvailableActions[]|null $available_actions available_actions
     *
     * @return self
     */
    public function setAvailableActions($available_actions)
    {
        $this->container['available_actions'] = $available_actions;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}


