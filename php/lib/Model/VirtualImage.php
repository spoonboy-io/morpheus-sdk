<?php
/**
 * VirtualImage
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
 * VirtualImage Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class VirtualImage implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'virtualImage';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'name' => 'string',
        'description' => 'string',
        'labels' => 'string[]',
        'owner_id' => 'int',
        'tenant' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance',
        'image_type' => 'string',
        'user_uploaded' => 'bool',
        'user_defined' => 'bool',
        'system_image' => 'bool',
        'is_cloud_init' => 'bool',
        'ssh_username' => 'string',
        'ssh_password' => 'string',
        'ssh_password_hash' => 'string',
        'ssh_key' => 'string',
        'os_type' => '\OpenAPI\Client\Model\VirtualImageOsType',
        'min_ram' => 'int',
        'min_ram_gb' => 'int',
        'min_disk' => 'string',
        'min_disk_gb' => 'string',
        'raw_size' => 'int',
        'raw_size_gb' => 'float',
        'trial_version' => 'bool',
        'virtio_supported' => 'bool',
        'uefi' => 'string',
        'is_auto_join_domain' => 'bool',
        'vm_tools_installed' => 'bool',
        'install_agent' => 'bool',
        'is_force_customization' => 'bool',
        'is_sysprep' => 'bool',
        'fips_enabled' => 'bool',
        'user_data' => 'string',
        'console_keymap' => 'string',
        'storage_provider' => 'string',
        'external_id' => 'string',
        'visibility' => 'string',
        'accounts' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance[]',
        'config' => '\OpenAPI\Client\Model\VirtualImageConfig',
        'volumes' => 'object[]',
        'storage_controllers' => 'object[]',
        'network_interfaces' => 'object[]',
        'tags' => 'object[]',
        'locations' => 'object[]',
        'date_created' => '\DateTime',
        'last_updated' => '\DateTime',
        'status' => 'string'
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
        'description' => null,
        'labels' => null,
        'owner_id' => 'int64',
        'tenant' => null,
        'image_type' => null,
        'user_uploaded' => null,
        'user_defined' => null,
        'system_image' => null,
        'is_cloud_init' => null,
        'ssh_username' => null,
        'ssh_password' => null,
        'ssh_password_hash' => null,
        'ssh_key' => null,
        'os_type' => null,
        'min_ram' => 'int64',
        'min_ram_gb' => 'int64',
        'min_disk' => null,
        'min_disk_gb' => null,
        'raw_size' => 'int64',
        'raw_size_gb' => null,
        'trial_version' => null,
        'virtio_supported' => null,
        'uefi' => null,
        'is_auto_join_domain' => null,
        'vm_tools_installed' => null,
        'install_agent' => null,
        'is_force_customization' => null,
        'is_sysprep' => null,
        'fips_enabled' => null,
        'user_data' => null,
        'console_keymap' => null,
        'storage_provider' => null,
        'external_id' => null,
        'visibility' => null,
        'accounts' => null,
        'config' => null,
        'volumes' => null,
        'storage_controllers' => null,
        'network_interfaces' => null,
        'tags' => null,
        'locations' => null,
        'date_created' => 'date-time',
        'last_updated' => 'date-time',
        'status' => null
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
        'description' => 'description',
        'labels' => 'labels',
        'owner_id' => 'ownerId',
        'tenant' => 'tenant',
        'image_type' => 'imageType',
        'user_uploaded' => 'userUploaded',
        'user_defined' => 'userDefined',
        'system_image' => 'systemImage',
        'is_cloud_init' => 'isCloudInit',
        'ssh_username' => 'sshUsername',
        'ssh_password' => 'sshPassword',
        'ssh_password_hash' => 'sshPasswordHash',
        'ssh_key' => 'sshKey',
        'os_type' => 'osType',
        'min_ram' => 'minRam',
        'min_ram_gb' => 'minRamGB',
        'min_disk' => 'minDisk',
        'min_disk_gb' => 'minDiskGB',
        'raw_size' => 'rawSize',
        'raw_size_gb' => 'rawSizeGB',
        'trial_version' => 'trialVersion',
        'virtio_supported' => 'virtioSupported',
        'uefi' => 'uefi',
        'is_auto_join_domain' => 'isAutoJoinDomain',
        'vm_tools_installed' => 'vmToolsInstalled',
        'install_agent' => 'installAgent',
        'is_force_customization' => 'isForceCustomization',
        'is_sysprep' => 'isSysprep',
        'fips_enabled' => 'fipsEnabled',
        'user_data' => 'userData',
        'console_keymap' => 'consoleKeymap',
        'storage_provider' => 'storageProvider',
        'external_id' => 'externalId',
        'visibility' => 'visibility',
        'accounts' => 'accounts',
        'config' => 'config',
        'volumes' => 'volumes',
        'storage_controllers' => 'storageControllers',
        'network_interfaces' => 'networkInterfaces',
        'tags' => 'tags',
        'locations' => 'locations',
        'date_created' => 'dateCreated',
        'last_updated' => 'lastUpdated',
        'status' => 'status'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'name' => 'setName',
        'description' => 'setDescription',
        'labels' => 'setLabels',
        'owner_id' => 'setOwnerId',
        'tenant' => 'setTenant',
        'image_type' => 'setImageType',
        'user_uploaded' => 'setUserUploaded',
        'user_defined' => 'setUserDefined',
        'system_image' => 'setSystemImage',
        'is_cloud_init' => 'setIsCloudInit',
        'ssh_username' => 'setSshUsername',
        'ssh_password' => 'setSshPassword',
        'ssh_password_hash' => 'setSshPasswordHash',
        'ssh_key' => 'setSshKey',
        'os_type' => 'setOsType',
        'min_ram' => 'setMinRam',
        'min_ram_gb' => 'setMinRamGb',
        'min_disk' => 'setMinDisk',
        'min_disk_gb' => 'setMinDiskGb',
        'raw_size' => 'setRawSize',
        'raw_size_gb' => 'setRawSizeGb',
        'trial_version' => 'setTrialVersion',
        'virtio_supported' => 'setVirtioSupported',
        'uefi' => 'setUefi',
        'is_auto_join_domain' => 'setIsAutoJoinDomain',
        'vm_tools_installed' => 'setVmToolsInstalled',
        'install_agent' => 'setInstallAgent',
        'is_force_customization' => 'setIsForceCustomization',
        'is_sysprep' => 'setIsSysprep',
        'fips_enabled' => 'setFipsEnabled',
        'user_data' => 'setUserData',
        'console_keymap' => 'setConsoleKeymap',
        'storage_provider' => 'setStorageProvider',
        'external_id' => 'setExternalId',
        'visibility' => 'setVisibility',
        'accounts' => 'setAccounts',
        'config' => 'setConfig',
        'volumes' => 'setVolumes',
        'storage_controllers' => 'setStorageControllers',
        'network_interfaces' => 'setNetworkInterfaces',
        'tags' => 'setTags',
        'locations' => 'setLocations',
        'date_created' => 'setDateCreated',
        'last_updated' => 'setLastUpdated',
        'status' => 'setStatus'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'name' => 'getName',
        'description' => 'getDescription',
        'labels' => 'getLabels',
        'owner_id' => 'getOwnerId',
        'tenant' => 'getTenant',
        'image_type' => 'getImageType',
        'user_uploaded' => 'getUserUploaded',
        'user_defined' => 'getUserDefined',
        'system_image' => 'getSystemImage',
        'is_cloud_init' => 'getIsCloudInit',
        'ssh_username' => 'getSshUsername',
        'ssh_password' => 'getSshPassword',
        'ssh_password_hash' => 'getSshPasswordHash',
        'ssh_key' => 'getSshKey',
        'os_type' => 'getOsType',
        'min_ram' => 'getMinRam',
        'min_ram_gb' => 'getMinRamGb',
        'min_disk' => 'getMinDisk',
        'min_disk_gb' => 'getMinDiskGb',
        'raw_size' => 'getRawSize',
        'raw_size_gb' => 'getRawSizeGb',
        'trial_version' => 'getTrialVersion',
        'virtio_supported' => 'getVirtioSupported',
        'uefi' => 'getUefi',
        'is_auto_join_domain' => 'getIsAutoJoinDomain',
        'vm_tools_installed' => 'getVmToolsInstalled',
        'install_agent' => 'getInstallAgent',
        'is_force_customization' => 'getIsForceCustomization',
        'is_sysprep' => 'getIsSysprep',
        'fips_enabled' => 'getFipsEnabled',
        'user_data' => 'getUserData',
        'console_keymap' => 'getConsoleKeymap',
        'storage_provider' => 'getStorageProvider',
        'external_id' => 'getExternalId',
        'visibility' => 'getVisibility',
        'accounts' => 'getAccounts',
        'config' => 'getConfig',
        'volumes' => 'getVolumes',
        'storage_controllers' => 'getStorageControllers',
        'network_interfaces' => 'getNetworkInterfaces',
        'tags' => 'getTags',
        'locations' => 'getLocations',
        'date_created' => 'getDateCreated',
        'last_updated' => 'getLastUpdated',
        'status' => 'getStatus'
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
        $this->container['description'] = $data['description'] ?? null;
        $this->container['labels'] = $data['labels'] ?? null;
        $this->container['owner_id'] = $data['owner_id'] ?? null;
        $this->container['tenant'] = $data['tenant'] ?? null;
        $this->container['image_type'] = $data['image_type'] ?? null;
        $this->container['user_uploaded'] = $data['user_uploaded'] ?? null;
        $this->container['user_defined'] = $data['user_defined'] ?? null;
        $this->container['system_image'] = $data['system_image'] ?? null;
        $this->container['is_cloud_init'] = $data['is_cloud_init'] ?? null;
        $this->container['ssh_username'] = $data['ssh_username'] ?? null;
        $this->container['ssh_password'] = $data['ssh_password'] ?? null;
        $this->container['ssh_password_hash'] = $data['ssh_password_hash'] ?? null;
        $this->container['ssh_key'] = $data['ssh_key'] ?? null;
        $this->container['os_type'] = $data['os_type'] ?? null;
        $this->container['min_ram'] = $data['min_ram'] ?? null;
        $this->container['min_ram_gb'] = $data['min_ram_gb'] ?? null;
        $this->container['min_disk'] = $data['min_disk'] ?? null;
        $this->container['min_disk_gb'] = $data['min_disk_gb'] ?? null;
        $this->container['raw_size'] = $data['raw_size'] ?? null;
        $this->container['raw_size_gb'] = $data['raw_size_gb'] ?? null;
        $this->container['trial_version'] = $data['trial_version'] ?? null;
        $this->container['virtio_supported'] = $data['virtio_supported'] ?? null;
        $this->container['uefi'] = $data['uefi'] ?? null;
        $this->container['is_auto_join_domain'] = $data['is_auto_join_domain'] ?? null;
        $this->container['vm_tools_installed'] = $data['vm_tools_installed'] ?? null;
        $this->container['install_agent'] = $data['install_agent'] ?? null;
        $this->container['is_force_customization'] = $data['is_force_customization'] ?? null;
        $this->container['is_sysprep'] = $data['is_sysprep'] ?? null;
        $this->container['fips_enabled'] = $data['fips_enabled'] ?? null;
        $this->container['user_data'] = $data['user_data'] ?? null;
        $this->container['console_keymap'] = $data['console_keymap'] ?? null;
        $this->container['storage_provider'] = $data['storage_provider'] ?? null;
        $this->container['external_id'] = $data['external_id'] ?? null;
        $this->container['visibility'] = $data['visibility'] ?? null;
        $this->container['accounts'] = $data['accounts'] ?? null;
        $this->container['config'] = $data['config'] ?? null;
        $this->container['volumes'] = $data['volumes'] ?? null;
        $this->container['storage_controllers'] = $data['storage_controllers'] ?? null;
        $this->container['network_interfaces'] = $data['network_interfaces'] ?? null;
        $this->container['tags'] = $data['tags'] ?? null;
        $this->container['locations'] = $data['locations'] ?? null;
        $this->container['date_created'] = $data['date_created'] ?? null;
        $this->container['last_updated'] = $data['last_updated'] ?? null;
        $this->container['status'] = $data['status'] ?? null;
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
     * Gets description
     *
     * @return string|null
     */
    public function getDescription()
    {
        return $this->container['description'];
    }

    /**
     * Sets description
     *
     * @param string|null $description description
     *
     * @return self
     */
    public function setDescription($description)
    {
        $this->container['description'] = $description;

        return $this;
    }

    /**
     * Gets labels
     *
     * @return string[]|null
     */
    public function getLabels()
    {
        return $this->container['labels'];
    }

    /**
     * Sets labels
     *
     * @param string[]|null $labels labels
     *
     * @return self
     */
    public function setLabels($labels)
    {
        $this->container['labels'] = $labels;

        return $this;
    }

    /**
     * Gets owner_id
     *
     * @return int|null
     */
    public function getOwnerId()
    {
        return $this->container['owner_id'];
    }

    /**
     * Sets owner_id
     *
     * @param int|null $owner_id owner_id
     *
     * @return self
     */
    public function setOwnerId($owner_id)
    {
        $this->container['owner_id'] = $owner_id;

        return $this;
    }

    /**
     * Gets tenant
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null
     */
    public function getTenant()
    {
        return $this->container['tenant'];
    }

    /**
     * Sets tenant
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null $tenant tenant
     *
     * @return self
     */
    public function setTenant($tenant)
    {
        $this->container['tenant'] = $tenant;

        return $this;
    }

    /**
     * Gets image_type
     *
     * @return string|null
     */
    public function getImageType()
    {
        return $this->container['image_type'];
    }

    /**
     * Sets image_type
     *
     * @param string|null $image_type image_type
     *
     * @return self
     */
    public function setImageType($image_type)
    {
        $this->container['image_type'] = $image_type;

        return $this;
    }

    /**
     * Gets user_uploaded
     *
     * @return bool|null
     */
    public function getUserUploaded()
    {
        return $this->container['user_uploaded'];
    }

    /**
     * Sets user_uploaded
     *
     * @param bool|null $user_uploaded user_uploaded
     *
     * @return self
     */
    public function setUserUploaded($user_uploaded)
    {
        $this->container['user_uploaded'] = $user_uploaded;

        return $this;
    }

    /**
     * Gets user_defined
     *
     * @return bool|null
     */
    public function getUserDefined()
    {
        return $this->container['user_defined'];
    }

    /**
     * Sets user_defined
     *
     * @param bool|null $user_defined user_defined
     *
     * @return self
     */
    public function setUserDefined($user_defined)
    {
        $this->container['user_defined'] = $user_defined;

        return $this;
    }

    /**
     * Gets system_image
     *
     * @return bool|null
     */
    public function getSystemImage()
    {
        return $this->container['system_image'];
    }

    /**
     * Sets system_image
     *
     * @param bool|null $system_image system_image
     *
     * @return self
     */
    public function setSystemImage($system_image)
    {
        $this->container['system_image'] = $system_image;

        return $this;
    }

    /**
     * Gets is_cloud_init
     *
     * @return bool|null
     */
    public function getIsCloudInit()
    {
        return $this->container['is_cloud_init'];
    }

    /**
     * Sets is_cloud_init
     *
     * @param bool|null $is_cloud_init is_cloud_init
     *
     * @return self
     */
    public function setIsCloudInit($is_cloud_init)
    {
        $this->container['is_cloud_init'] = $is_cloud_init;

        return $this;
    }

    /**
     * Gets ssh_username
     *
     * @return string|null
     */
    public function getSshUsername()
    {
        return $this->container['ssh_username'];
    }

    /**
     * Sets ssh_username
     *
     * @param string|null $ssh_username ssh_username
     *
     * @return self
     */
    public function setSshUsername($ssh_username)
    {
        $this->container['ssh_username'] = $ssh_username;

        return $this;
    }

    /**
     * Gets ssh_password
     *
     * @return string|null
     */
    public function getSshPassword()
    {
        return $this->container['ssh_password'];
    }

    /**
     * Sets ssh_password
     *
     * @param string|null $ssh_password ssh_password
     *
     * @return self
     */
    public function setSshPassword($ssh_password)
    {
        $this->container['ssh_password'] = $ssh_password;

        return $this;
    }

    /**
     * Gets ssh_password_hash
     *
     * @return string|null
     */
    public function getSshPasswordHash()
    {
        return $this->container['ssh_password_hash'];
    }

    /**
     * Sets ssh_password_hash
     *
     * @param string|null $ssh_password_hash ssh_password_hash
     *
     * @return self
     */
    public function setSshPasswordHash($ssh_password_hash)
    {
        $this->container['ssh_password_hash'] = $ssh_password_hash;

        return $this;
    }

    /**
     * Gets ssh_key
     *
     * @return string|null
     */
    public function getSshKey()
    {
        return $this->container['ssh_key'];
    }

    /**
     * Sets ssh_key
     *
     * @param string|null $ssh_key ssh_key
     *
     * @return self
     */
    public function setSshKey($ssh_key)
    {
        $this->container['ssh_key'] = $ssh_key;

        return $this;
    }

    /**
     * Gets os_type
     *
     * @return \OpenAPI\Client\Model\VirtualImageOsType|null
     */
    public function getOsType()
    {
        return $this->container['os_type'];
    }

    /**
     * Sets os_type
     *
     * @param \OpenAPI\Client\Model\VirtualImageOsType|null $os_type os_type
     *
     * @return self
     */
    public function setOsType($os_type)
    {
        $this->container['os_type'] = $os_type;

        return $this;
    }

    /**
     * Gets min_ram
     *
     * @return int|null
     */
    public function getMinRam()
    {
        return $this->container['min_ram'];
    }

    /**
     * Sets min_ram
     *
     * @param int|null $min_ram min_ram
     *
     * @return self
     */
    public function setMinRam($min_ram)
    {
        $this->container['min_ram'] = $min_ram;

        return $this;
    }

    /**
     * Gets min_ram_gb
     *
     * @return int|null
     */
    public function getMinRamGb()
    {
        return $this->container['min_ram_gb'];
    }

    /**
     * Sets min_ram_gb
     *
     * @param int|null $min_ram_gb min_ram_gb
     *
     * @return self
     */
    public function setMinRamGb($min_ram_gb)
    {
        $this->container['min_ram_gb'] = $min_ram_gb;

        return $this;
    }

    /**
     * Gets min_disk
     *
     * @return string|null
     */
    public function getMinDisk()
    {
        return $this->container['min_disk'];
    }

    /**
     * Sets min_disk
     *
     * @param string|null $min_disk min_disk
     *
     * @return self
     */
    public function setMinDisk($min_disk)
    {
        $this->container['min_disk'] = $min_disk;

        return $this;
    }

    /**
     * Gets min_disk_gb
     *
     * @return string|null
     */
    public function getMinDiskGb()
    {
        return $this->container['min_disk_gb'];
    }

    /**
     * Sets min_disk_gb
     *
     * @param string|null $min_disk_gb min_disk_gb
     *
     * @return self
     */
    public function setMinDiskGb($min_disk_gb)
    {
        $this->container['min_disk_gb'] = $min_disk_gb;

        return $this;
    }

    /**
     * Gets raw_size
     *
     * @return int|null
     */
    public function getRawSize()
    {
        return $this->container['raw_size'];
    }

    /**
     * Sets raw_size
     *
     * @param int|null $raw_size raw_size
     *
     * @return self
     */
    public function setRawSize($raw_size)
    {
        $this->container['raw_size'] = $raw_size;

        return $this;
    }

    /**
     * Gets raw_size_gb
     *
     * @return float|null
     */
    public function getRawSizeGb()
    {
        return $this->container['raw_size_gb'];
    }

    /**
     * Sets raw_size_gb
     *
     * @param float|null $raw_size_gb raw_size_gb
     *
     * @return self
     */
    public function setRawSizeGb($raw_size_gb)
    {
        $this->container['raw_size_gb'] = $raw_size_gb;

        return $this;
    }

    /**
     * Gets trial_version
     *
     * @return bool|null
     */
    public function getTrialVersion()
    {
        return $this->container['trial_version'];
    }

    /**
     * Sets trial_version
     *
     * @param bool|null $trial_version trial_version
     *
     * @return self
     */
    public function setTrialVersion($trial_version)
    {
        $this->container['trial_version'] = $trial_version;

        return $this;
    }

    /**
     * Gets virtio_supported
     *
     * @return bool|null
     */
    public function getVirtioSupported()
    {
        return $this->container['virtio_supported'];
    }

    /**
     * Sets virtio_supported
     *
     * @param bool|null $virtio_supported virtio_supported
     *
     * @return self
     */
    public function setVirtioSupported($virtio_supported)
    {
        $this->container['virtio_supported'] = $virtio_supported;

        return $this;
    }

    /**
     * Gets uefi
     *
     * @return string|null
     */
    public function getUefi()
    {
        return $this->container['uefi'];
    }

    /**
     * Sets uefi
     *
     * @param string|null $uefi uefi
     *
     * @return self
     */
    public function setUefi($uefi)
    {
        $this->container['uefi'] = $uefi;

        return $this;
    }

    /**
     * Gets is_auto_join_domain
     *
     * @return bool|null
     */
    public function getIsAutoJoinDomain()
    {
        return $this->container['is_auto_join_domain'];
    }

    /**
     * Sets is_auto_join_domain
     *
     * @param bool|null $is_auto_join_domain is_auto_join_domain
     *
     * @return self
     */
    public function setIsAutoJoinDomain($is_auto_join_domain)
    {
        $this->container['is_auto_join_domain'] = $is_auto_join_domain;

        return $this;
    }

    /**
     * Gets vm_tools_installed
     *
     * @return bool|null
     */
    public function getVmToolsInstalled()
    {
        return $this->container['vm_tools_installed'];
    }

    /**
     * Sets vm_tools_installed
     *
     * @param bool|null $vm_tools_installed vm_tools_installed
     *
     * @return self
     */
    public function setVmToolsInstalled($vm_tools_installed)
    {
        $this->container['vm_tools_installed'] = $vm_tools_installed;

        return $this;
    }

    /**
     * Gets install_agent
     *
     * @return bool|null
     */
    public function getInstallAgent()
    {
        return $this->container['install_agent'];
    }

    /**
     * Sets install_agent
     *
     * @param bool|null $install_agent install_agent
     *
     * @return self
     */
    public function setInstallAgent($install_agent)
    {
        $this->container['install_agent'] = $install_agent;

        return $this;
    }

    /**
     * Gets is_force_customization
     *
     * @return bool|null
     */
    public function getIsForceCustomization()
    {
        return $this->container['is_force_customization'];
    }

    /**
     * Sets is_force_customization
     *
     * @param bool|null $is_force_customization is_force_customization
     *
     * @return self
     */
    public function setIsForceCustomization($is_force_customization)
    {
        $this->container['is_force_customization'] = $is_force_customization;

        return $this;
    }

    /**
     * Gets is_sysprep
     *
     * @return bool|null
     */
    public function getIsSysprep()
    {
        return $this->container['is_sysprep'];
    }

    /**
     * Sets is_sysprep
     *
     * @param bool|null $is_sysprep is_sysprep
     *
     * @return self
     */
    public function setIsSysprep($is_sysprep)
    {
        $this->container['is_sysprep'] = $is_sysprep;

        return $this;
    }

    /**
     * Gets fips_enabled
     *
     * @return bool|null
     */
    public function getFipsEnabled()
    {
        return $this->container['fips_enabled'];
    }

    /**
     * Sets fips_enabled
     *
     * @param bool|null $fips_enabled fips_enabled
     *
     * @return self
     */
    public function setFipsEnabled($fips_enabled)
    {
        $this->container['fips_enabled'] = $fips_enabled;

        return $this;
    }

    /**
     * Gets user_data
     *
     * @return string|null
     */
    public function getUserData()
    {
        return $this->container['user_data'];
    }

    /**
     * Sets user_data
     *
     * @param string|null $user_data user_data
     *
     * @return self
     */
    public function setUserData($user_data)
    {
        $this->container['user_data'] = $user_data;

        return $this;
    }

    /**
     * Gets console_keymap
     *
     * @return string|null
     */
    public function getConsoleKeymap()
    {
        return $this->container['console_keymap'];
    }

    /**
     * Sets console_keymap
     *
     * @param string|null $console_keymap console_keymap
     *
     * @return self
     */
    public function setConsoleKeymap($console_keymap)
    {
        $this->container['console_keymap'] = $console_keymap;

        return $this;
    }

    /**
     * Gets storage_provider
     *
     * @return string|null
     */
    public function getStorageProvider()
    {
        return $this->container['storage_provider'];
    }

    /**
     * Sets storage_provider
     *
     * @param string|null $storage_provider storage_provider
     *
     * @return self
     */
    public function setStorageProvider($storage_provider)
    {
        $this->container['storage_provider'] = $storage_provider;

        return $this;
    }

    /**
     * Gets external_id
     *
     * @return string|null
     */
    public function getExternalId()
    {
        return $this->container['external_id'];
    }

    /**
     * Sets external_id
     *
     * @param string|null $external_id external_id
     *
     * @return self
     */
    public function setExternalId($external_id)
    {
        $this->container['external_id'] = $external_id;

        return $this;
    }

    /**
     * Gets visibility
     *
     * @return string|null
     */
    public function getVisibility()
    {
        return $this->container['visibility'];
    }

    /**
     * Sets visibility
     *
     * @param string|null $visibility visibility
     *
     * @return self
     */
    public function setVisibility($visibility)
    {
        $this->container['visibility'] = $visibility;

        return $this;
    }

    /**
     * Gets accounts
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance[]|null
     */
    public function getAccounts()
    {
        return $this->container['accounts'];
    }

    /**
     * Sets accounts
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance[]|null $accounts accounts
     *
     * @return self
     */
    public function setAccounts($accounts)
    {
        $this->container['accounts'] = $accounts;

        return $this;
    }

    /**
     * Gets config
     *
     * @return \OpenAPI\Client\Model\VirtualImageConfig|null
     */
    public function getConfig()
    {
        return $this->container['config'];
    }

    /**
     * Sets config
     *
     * @param \OpenAPI\Client\Model\VirtualImageConfig|null $config config
     *
     * @return self
     */
    public function setConfig($config)
    {
        $this->container['config'] = $config;

        return $this;
    }

    /**
     * Gets volumes
     *
     * @return object[]|null
     */
    public function getVolumes()
    {
        return $this->container['volumes'];
    }

    /**
     * Sets volumes
     *
     * @param object[]|null $volumes volumes
     *
     * @return self
     */
    public function setVolumes($volumes)
    {
        $this->container['volumes'] = $volumes;

        return $this;
    }

    /**
     * Gets storage_controllers
     *
     * @return object[]|null
     */
    public function getStorageControllers()
    {
        return $this->container['storage_controllers'];
    }

    /**
     * Sets storage_controllers
     *
     * @param object[]|null $storage_controllers storage_controllers
     *
     * @return self
     */
    public function setStorageControllers($storage_controllers)
    {
        $this->container['storage_controllers'] = $storage_controllers;

        return $this;
    }

    /**
     * Gets network_interfaces
     *
     * @return object[]|null
     */
    public function getNetworkInterfaces()
    {
        return $this->container['network_interfaces'];
    }

    /**
     * Sets network_interfaces
     *
     * @param object[]|null $network_interfaces network_interfaces
     *
     * @return self
     */
    public function setNetworkInterfaces($network_interfaces)
    {
        $this->container['network_interfaces'] = $network_interfaces;

        return $this;
    }

    /**
     * Gets tags
     *
     * @return object[]|null
     */
    public function getTags()
    {
        return $this->container['tags'];
    }

    /**
     * Sets tags
     *
     * @param object[]|null $tags tags
     *
     * @return self
     */
    public function setTags($tags)
    {
        $this->container['tags'] = $tags;

        return $this;
    }

    /**
     * Gets locations
     *
     * @return object[]|null
     */
    public function getLocations()
    {
        return $this->container['locations'];
    }

    /**
     * Sets locations
     *
     * @param object[]|null $locations locations
     *
     * @return self
     */
    public function setLocations($locations)
    {
        $this->container['locations'] = $locations;

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


