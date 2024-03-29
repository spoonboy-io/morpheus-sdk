<?php
/**
 * ZoneType
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
 * ZoneType Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ZoneType implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'zoneType';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'name' => 'string',
        'code' => 'string',
        'enabled' => 'bool',
        'provision' => 'bool',
        'auto_capacity' => 'bool',
        'migration_target' => 'bool',
        'has_datastores' => 'bool',
        'has_networks' => 'bool',
        'has_resource_pools' => 'bool',
        'has_security_groups' => 'bool',
        'has_containers' => 'bool',
        'has_bare_metal' => 'bool',
        'has_services' => 'bool',
        'has_functions' => 'bool',
        'has_jobs' => 'bool',
        'has_discovery' => 'bool',
        'has_cloud_init' => 'bool',
        'has_folders' => 'bool',
        'has_floating_ips' => 'bool',
        'has_marketplace' => 'bool',
        'can_create_resource_pools' => 'bool',
        'can_delete_resource_pools' => 'bool',
        'can_create_datastores' => 'bool',
        'can_create_networks' => 'bool',
        'can_choose_container_mode' => 'bool',
        'provision_requires_resource_pool' => 'bool',
        'supports_distributed_worker' => 'bool',
        'cloud' => 'string',
        'provision_types' => 'int[]',
        'zone_instance_type_layout_id' => 'int',
        'server_types' => '\OpenAPI\Client\Model\ZoneTypeServerTypes[]',
        'option_types' => '\OpenAPI\Client\Model\ZoneTypeOptionTypes1[]'
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
        'name' => null,
        'code' => null,
        'enabled' => null,
        'provision' => null,
        'auto_capacity' => null,
        'migration_target' => null,
        'has_datastores' => null,
        'has_networks' => null,
        'has_resource_pools' => null,
        'has_security_groups' => null,
        'has_containers' => null,
        'has_bare_metal' => null,
        'has_services' => null,
        'has_functions' => null,
        'has_jobs' => null,
        'has_discovery' => null,
        'has_cloud_init' => null,
        'has_folders' => null,
        'has_floating_ips' => null,
        'has_marketplace' => null,
        'can_create_resource_pools' => null,
        'can_delete_resource_pools' => null,
        'can_create_datastores' => null,
        'can_create_networks' => null,
        'can_choose_container_mode' => null,
        'provision_requires_resource_pool' => null,
        'supports_distributed_worker' => null,
        'cloud' => null,
        'provision_types' => 'int64',
        'zone_instance_type_layout_id' => 'int64',
        'server_types' => null,
        'option_types' => null
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
        'name' => 'name',
        'code' => 'code',
        'enabled' => 'enabled',
        'provision' => 'provision',
        'auto_capacity' => 'autoCapacity',
        'migration_target' => 'migrationTarget',
        'has_datastores' => 'hasDatastores',
        'has_networks' => 'hasNetworks',
        'has_resource_pools' => 'hasResourcePools',
        'has_security_groups' => 'hasSecurityGroups',
        'has_containers' => 'hasContainers',
        'has_bare_metal' => 'hasBareMetal',
        'has_services' => 'hasServices',
        'has_functions' => 'hasFunctions',
        'has_jobs' => 'hasJobs',
        'has_discovery' => 'hasDiscovery',
        'has_cloud_init' => 'hasCloudInit',
        'has_folders' => 'hasFolders',
        'has_floating_ips' => 'hasFloatingIps',
        'has_marketplace' => 'hasMarketplace',
        'can_create_resource_pools' => 'canCreateResourcePools',
        'can_delete_resource_pools' => 'canDeleteResourcePools',
        'can_create_datastores' => 'canCreateDatastores',
        'can_create_networks' => 'canCreateNetworks',
        'can_choose_container_mode' => 'canChooseContainerMode',
        'provision_requires_resource_pool' => 'provisionRequiresResourcePool',
        'supports_distributed_worker' => 'supportsDistributedWorker',
        'cloud' => 'cloud',
        'provision_types' => 'provisionTypes',
        'zone_instance_type_layout_id' => 'zoneInstanceTypeLayoutId',
        'server_types' => 'serverTypes',
        'option_types' => 'optionTypes'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'name' => 'setName',
        'code' => 'setCode',
        'enabled' => 'setEnabled',
        'provision' => 'setProvision',
        'auto_capacity' => 'setAutoCapacity',
        'migration_target' => 'setMigrationTarget',
        'has_datastores' => 'setHasDatastores',
        'has_networks' => 'setHasNetworks',
        'has_resource_pools' => 'setHasResourcePools',
        'has_security_groups' => 'setHasSecurityGroups',
        'has_containers' => 'setHasContainers',
        'has_bare_metal' => 'setHasBareMetal',
        'has_services' => 'setHasServices',
        'has_functions' => 'setHasFunctions',
        'has_jobs' => 'setHasJobs',
        'has_discovery' => 'setHasDiscovery',
        'has_cloud_init' => 'setHasCloudInit',
        'has_folders' => 'setHasFolders',
        'has_floating_ips' => 'setHasFloatingIps',
        'has_marketplace' => 'setHasMarketplace',
        'can_create_resource_pools' => 'setCanCreateResourcePools',
        'can_delete_resource_pools' => 'setCanDeleteResourcePools',
        'can_create_datastores' => 'setCanCreateDatastores',
        'can_create_networks' => 'setCanCreateNetworks',
        'can_choose_container_mode' => 'setCanChooseContainerMode',
        'provision_requires_resource_pool' => 'setProvisionRequiresResourcePool',
        'supports_distributed_worker' => 'setSupportsDistributedWorker',
        'cloud' => 'setCloud',
        'provision_types' => 'setProvisionTypes',
        'zone_instance_type_layout_id' => 'setZoneInstanceTypeLayoutId',
        'server_types' => 'setServerTypes',
        'option_types' => 'setOptionTypes'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'name' => 'getName',
        'code' => 'getCode',
        'enabled' => 'getEnabled',
        'provision' => 'getProvision',
        'auto_capacity' => 'getAutoCapacity',
        'migration_target' => 'getMigrationTarget',
        'has_datastores' => 'getHasDatastores',
        'has_networks' => 'getHasNetworks',
        'has_resource_pools' => 'getHasResourcePools',
        'has_security_groups' => 'getHasSecurityGroups',
        'has_containers' => 'getHasContainers',
        'has_bare_metal' => 'getHasBareMetal',
        'has_services' => 'getHasServices',
        'has_functions' => 'getHasFunctions',
        'has_jobs' => 'getHasJobs',
        'has_discovery' => 'getHasDiscovery',
        'has_cloud_init' => 'getHasCloudInit',
        'has_folders' => 'getHasFolders',
        'has_floating_ips' => 'getHasFloatingIps',
        'has_marketplace' => 'getHasMarketplace',
        'can_create_resource_pools' => 'getCanCreateResourcePools',
        'can_delete_resource_pools' => 'getCanDeleteResourcePools',
        'can_create_datastores' => 'getCanCreateDatastores',
        'can_create_networks' => 'getCanCreateNetworks',
        'can_choose_container_mode' => 'getCanChooseContainerMode',
        'provision_requires_resource_pool' => 'getProvisionRequiresResourcePool',
        'supports_distributed_worker' => 'getSupportsDistributedWorker',
        'cloud' => 'getCloud',
        'provision_types' => 'getProvisionTypes',
        'zone_instance_type_layout_id' => 'getZoneInstanceTypeLayoutId',
        'server_types' => 'getServerTypes',
        'option_types' => 'getOptionTypes'
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
        $this->container['name'] = $data['name'] ?? null;
        $this->container['code'] = $data['code'] ?? null;
        $this->container['enabled'] = $data['enabled'] ?? null;
        $this->container['provision'] = $data['provision'] ?? null;
        $this->container['auto_capacity'] = $data['auto_capacity'] ?? null;
        $this->container['migration_target'] = $data['migration_target'] ?? null;
        $this->container['has_datastores'] = $data['has_datastores'] ?? null;
        $this->container['has_networks'] = $data['has_networks'] ?? null;
        $this->container['has_resource_pools'] = $data['has_resource_pools'] ?? null;
        $this->container['has_security_groups'] = $data['has_security_groups'] ?? null;
        $this->container['has_containers'] = $data['has_containers'] ?? null;
        $this->container['has_bare_metal'] = $data['has_bare_metal'] ?? null;
        $this->container['has_services'] = $data['has_services'] ?? null;
        $this->container['has_functions'] = $data['has_functions'] ?? null;
        $this->container['has_jobs'] = $data['has_jobs'] ?? null;
        $this->container['has_discovery'] = $data['has_discovery'] ?? null;
        $this->container['has_cloud_init'] = $data['has_cloud_init'] ?? null;
        $this->container['has_folders'] = $data['has_folders'] ?? null;
        $this->container['has_floating_ips'] = $data['has_floating_ips'] ?? null;
        $this->container['has_marketplace'] = $data['has_marketplace'] ?? null;
        $this->container['can_create_resource_pools'] = $data['can_create_resource_pools'] ?? null;
        $this->container['can_delete_resource_pools'] = $data['can_delete_resource_pools'] ?? null;
        $this->container['can_create_datastores'] = $data['can_create_datastores'] ?? null;
        $this->container['can_create_networks'] = $data['can_create_networks'] ?? null;
        $this->container['can_choose_container_mode'] = $data['can_choose_container_mode'] ?? null;
        $this->container['provision_requires_resource_pool'] = $data['provision_requires_resource_pool'] ?? null;
        $this->container['supports_distributed_worker'] = $data['supports_distributed_worker'] ?? null;
        $this->container['cloud'] = $data['cloud'] ?? null;
        $this->container['provision_types'] = $data['provision_types'] ?? null;
        $this->container['zone_instance_type_layout_id'] = $data['zone_instance_type_layout_id'] ?? null;
        $this->container['server_types'] = $data['server_types'] ?? null;
        $this->container['option_types'] = $data['option_types'] ?? null;
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
     * Gets code
     *
     * @return string|null
     */
    public function getCode()
    {
        return $this->container['code'];
    }

    /**
     * Sets code
     *
     * @param string|null $code code
     *
     * @return self
     */
    public function setCode($code)
    {
        $this->container['code'] = $code;

        return $this;
    }

    /**
     * Gets enabled
     *
     * @return bool|null
     */
    public function getEnabled()
    {
        return $this->container['enabled'];
    }

    /**
     * Sets enabled
     *
     * @param bool|null $enabled enabled
     *
     * @return self
     */
    public function setEnabled($enabled)
    {
        $this->container['enabled'] = $enabled;

        return $this;
    }

    /**
     * Gets provision
     *
     * @return bool|null
     */
    public function getProvision()
    {
        return $this->container['provision'];
    }

    /**
     * Sets provision
     *
     * @param bool|null $provision provision
     *
     * @return self
     */
    public function setProvision($provision)
    {
        $this->container['provision'] = $provision;

        return $this;
    }

    /**
     * Gets auto_capacity
     *
     * @return bool|null
     */
    public function getAutoCapacity()
    {
        return $this->container['auto_capacity'];
    }

    /**
     * Sets auto_capacity
     *
     * @param bool|null $auto_capacity auto_capacity
     *
     * @return self
     */
    public function setAutoCapacity($auto_capacity)
    {
        $this->container['auto_capacity'] = $auto_capacity;

        return $this;
    }

    /**
     * Gets migration_target
     *
     * @return bool|null
     */
    public function getMigrationTarget()
    {
        return $this->container['migration_target'];
    }

    /**
     * Sets migration_target
     *
     * @param bool|null $migration_target migration_target
     *
     * @return self
     */
    public function setMigrationTarget($migration_target)
    {
        $this->container['migration_target'] = $migration_target;

        return $this;
    }

    /**
     * Gets has_datastores
     *
     * @return bool|null
     */
    public function getHasDatastores()
    {
        return $this->container['has_datastores'];
    }

    /**
     * Sets has_datastores
     *
     * @param bool|null $has_datastores has_datastores
     *
     * @return self
     */
    public function setHasDatastores($has_datastores)
    {
        $this->container['has_datastores'] = $has_datastores;

        return $this;
    }

    /**
     * Gets has_networks
     *
     * @return bool|null
     */
    public function getHasNetworks()
    {
        return $this->container['has_networks'];
    }

    /**
     * Sets has_networks
     *
     * @param bool|null $has_networks has_networks
     *
     * @return self
     */
    public function setHasNetworks($has_networks)
    {
        $this->container['has_networks'] = $has_networks;

        return $this;
    }

    /**
     * Gets has_resource_pools
     *
     * @return bool|null
     */
    public function getHasResourcePools()
    {
        return $this->container['has_resource_pools'];
    }

    /**
     * Sets has_resource_pools
     *
     * @param bool|null $has_resource_pools has_resource_pools
     *
     * @return self
     */
    public function setHasResourcePools($has_resource_pools)
    {
        $this->container['has_resource_pools'] = $has_resource_pools;

        return $this;
    }

    /**
     * Gets has_security_groups
     *
     * @return bool|null
     */
    public function getHasSecurityGroups()
    {
        return $this->container['has_security_groups'];
    }

    /**
     * Sets has_security_groups
     *
     * @param bool|null $has_security_groups has_security_groups
     *
     * @return self
     */
    public function setHasSecurityGroups($has_security_groups)
    {
        $this->container['has_security_groups'] = $has_security_groups;

        return $this;
    }

    /**
     * Gets has_containers
     *
     * @return bool|null
     */
    public function getHasContainers()
    {
        return $this->container['has_containers'];
    }

    /**
     * Sets has_containers
     *
     * @param bool|null $has_containers has_containers
     *
     * @return self
     */
    public function setHasContainers($has_containers)
    {
        $this->container['has_containers'] = $has_containers;

        return $this;
    }

    /**
     * Gets has_bare_metal
     *
     * @return bool|null
     */
    public function getHasBareMetal()
    {
        return $this->container['has_bare_metal'];
    }

    /**
     * Sets has_bare_metal
     *
     * @param bool|null $has_bare_metal has_bare_metal
     *
     * @return self
     */
    public function setHasBareMetal($has_bare_metal)
    {
        $this->container['has_bare_metal'] = $has_bare_metal;

        return $this;
    }

    /**
     * Gets has_services
     *
     * @return bool|null
     */
    public function getHasServices()
    {
        return $this->container['has_services'];
    }

    /**
     * Sets has_services
     *
     * @param bool|null $has_services has_services
     *
     * @return self
     */
    public function setHasServices($has_services)
    {
        $this->container['has_services'] = $has_services;

        return $this;
    }

    /**
     * Gets has_functions
     *
     * @return bool|null
     */
    public function getHasFunctions()
    {
        return $this->container['has_functions'];
    }

    /**
     * Sets has_functions
     *
     * @param bool|null $has_functions has_functions
     *
     * @return self
     */
    public function setHasFunctions($has_functions)
    {
        $this->container['has_functions'] = $has_functions;

        return $this;
    }

    /**
     * Gets has_jobs
     *
     * @return bool|null
     */
    public function getHasJobs()
    {
        return $this->container['has_jobs'];
    }

    /**
     * Sets has_jobs
     *
     * @param bool|null $has_jobs has_jobs
     *
     * @return self
     */
    public function setHasJobs($has_jobs)
    {
        $this->container['has_jobs'] = $has_jobs;

        return $this;
    }

    /**
     * Gets has_discovery
     *
     * @return bool|null
     */
    public function getHasDiscovery()
    {
        return $this->container['has_discovery'];
    }

    /**
     * Sets has_discovery
     *
     * @param bool|null $has_discovery has_discovery
     *
     * @return self
     */
    public function setHasDiscovery($has_discovery)
    {
        $this->container['has_discovery'] = $has_discovery;

        return $this;
    }

    /**
     * Gets has_cloud_init
     *
     * @return bool|null
     */
    public function getHasCloudInit()
    {
        return $this->container['has_cloud_init'];
    }

    /**
     * Sets has_cloud_init
     *
     * @param bool|null $has_cloud_init has_cloud_init
     *
     * @return self
     */
    public function setHasCloudInit($has_cloud_init)
    {
        $this->container['has_cloud_init'] = $has_cloud_init;

        return $this;
    }

    /**
     * Gets has_folders
     *
     * @return bool|null
     */
    public function getHasFolders()
    {
        return $this->container['has_folders'];
    }

    /**
     * Sets has_folders
     *
     * @param bool|null $has_folders has_folders
     *
     * @return self
     */
    public function setHasFolders($has_folders)
    {
        $this->container['has_folders'] = $has_folders;

        return $this;
    }

    /**
     * Gets has_floating_ips
     *
     * @return bool|null
     */
    public function getHasFloatingIps()
    {
        return $this->container['has_floating_ips'];
    }

    /**
     * Sets has_floating_ips
     *
     * @param bool|null $has_floating_ips has_floating_ips
     *
     * @return self
     */
    public function setHasFloatingIps($has_floating_ips)
    {
        $this->container['has_floating_ips'] = $has_floating_ips;

        return $this;
    }

    /**
     * Gets has_marketplace
     *
     * @return bool|null
     */
    public function getHasMarketplace()
    {
        return $this->container['has_marketplace'];
    }

    /**
     * Sets has_marketplace
     *
     * @param bool|null $has_marketplace has_marketplace
     *
     * @return self
     */
    public function setHasMarketplace($has_marketplace)
    {
        $this->container['has_marketplace'] = $has_marketplace;

        return $this;
    }

    /**
     * Gets can_create_resource_pools
     *
     * @return bool|null
     */
    public function getCanCreateResourcePools()
    {
        return $this->container['can_create_resource_pools'];
    }

    /**
     * Sets can_create_resource_pools
     *
     * @param bool|null $can_create_resource_pools can_create_resource_pools
     *
     * @return self
     */
    public function setCanCreateResourcePools($can_create_resource_pools)
    {
        $this->container['can_create_resource_pools'] = $can_create_resource_pools;

        return $this;
    }

    /**
     * Gets can_delete_resource_pools
     *
     * @return bool|null
     */
    public function getCanDeleteResourcePools()
    {
        return $this->container['can_delete_resource_pools'];
    }

    /**
     * Sets can_delete_resource_pools
     *
     * @param bool|null $can_delete_resource_pools can_delete_resource_pools
     *
     * @return self
     */
    public function setCanDeleteResourcePools($can_delete_resource_pools)
    {
        $this->container['can_delete_resource_pools'] = $can_delete_resource_pools;

        return $this;
    }

    /**
     * Gets can_create_datastores
     *
     * @return bool|null
     */
    public function getCanCreateDatastores()
    {
        return $this->container['can_create_datastores'];
    }

    /**
     * Sets can_create_datastores
     *
     * @param bool|null $can_create_datastores can_create_datastores
     *
     * @return self
     */
    public function setCanCreateDatastores($can_create_datastores)
    {
        $this->container['can_create_datastores'] = $can_create_datastores;

        return $this;
    }

    /**
     * Gets can_create_networks
     *
     * @return bool|null
     */
    public function getCanCreateNetworks()
    {
        return $this->container['can_create_networks'];
    }

    /**
     * Sets can_create_networks
     *
     * @param bool|null $can_create_networks can_create_networks
     *
     * @return self
     */
    public function setCanCreateNetworks($can_create_networks)
    {
        $this->container['can_create_networks'] = $can_create_networks;

        return $this;
    }

    /**
     * Gets can_choose_container_mode
     *
     * @return bool|null
     */
    public function getCanChooseContainerMode()
    {
        return $this->container['can_choose_container_mode'];
    }

    /**
     * Sets can_choose_container_mode
     *
     * @param bool|null $can_choose_container_mode can_choose_container_mode
     *
     * @return self
     */
    public function setCanChooseContainerMode($can_choose_container_mode)
    {
        $this->container['can_choose_container_mode'] = $can_choose_container_mode;

        return $this;
    }

    /**
     * Gets provision_requires_resource_pool
     *
     * @return bool|null
     */
    public function getProvisionRequiresResourcePool()
    {
        return $this->container['provision_requires_resource_pool'];
    }

    /**
     * Sets provision_requires_resource_pool
     *
     * @param bool|null $provision_requires_resource_pool provision_requires_resource_pool
     *
     * @return self
     */
    public function setProvisionRequiresResourcePool($provision_requires_resource_pool)
    {
        $this->container['provision_requires_resource_pool'] = $provision_requires_resource_pool;

        return $this;
    }

    /**
     * Gets supports_distributed_worker
     *
     * @return bool|null
     */
    public function getSupportsDistributedWorker()
    {
        return $this->container['supports_distributed_worker'];
    }

    /**
     * Sets supports_distributed_worker
     *
     * @param bool|null $supports_distributed_worker supports_distributed_worker
     *
     * @return self
     */
    public function setSupportsDistributedWorker($supports_distributed_worker)
    {
        $this->container['supports_distributed_worker'] = $supports_distributed_worker;

        return $this;
    }

    /**
     * Gets cloud
     *
     * @return string|null
     */
    public function getCloud()
    {
        return $this->container['cloud'];
    }

    /**
     * Sets cloud
     *
     * @param string|null $cloud cloud
     *
     * @return self
     */
    public function setCloud($cloud)
    {
        $this->container['cloud'] = $cloud;

        return $this;
    }

    /**
     * Gets provision_types
     *
     * @return int[]|null
     */
    public function getProvisionTypes()
    {
        return $this->container['provision_types'];
    }

    /**
     * Sets provision_types
     *
     * @param int[]|null $provision_types provision_types
     *
     * @return self
     */
    public function setProvisionTypes($provision_types)
    {
        $this->container['provision_types'] = $provision_types;

        return $this;
    }

    /**
     * Gets zone_instance_type_layout_id
     *
     * @return int|null
     */
    public function getZoneInstanceTypeLayoutId()
    {
        return $this->container['zone_instance_type_layout_id'];
    }

    /**
     * Sets zone_instance_type_layout_id
     *
     * @param int|null $zone_instance_type_layout_id zone_instance_type_layout_id
     *
     * @return self
     */
    public function setZoneInstanceTypeLayoutId($zone_instance_type_layout_id)
    {
        $this->container['zone_instance_type_layout_id'] = $zone_instance_type_layout_id;

        return $this;
    }

    /**
     * Gets server_types
     *
     * @return \OpenAPI\Client\Model\ZoneTypeServerTypes[]|null
     */
    public function getServerTypes()
    {
        return $this->container['server_types'];
    }

    /**
     * Sets server_types
     *
     * @param \OpenAPI\Client\Model\ZoneTypeServerTypes[]|null $server_types server_types
     *
     * @return self
     */
    public function setServerTypes($server_types)
    {
        $this->container['server_types'] = $server_types;

        return $this;
    }

    /**
     * Gets option_types
     *
     * @return \OpenAPI\Client\Model\ZoneTypeOptionTypes1[]|null
     */
    public function getOptionTypes()
    {
        return $this->container['option_types'];
    }

    /**
     * Sets option_types
     *
     * @param \OpenAPI\Client\Model\ZoneTypeOptionTypes1[]|null $option_types option_types
     *
     * @return self
     */
    public function setOptionTypes($option_types)
    {
        $this->container['option_types'] = $option_types;

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


