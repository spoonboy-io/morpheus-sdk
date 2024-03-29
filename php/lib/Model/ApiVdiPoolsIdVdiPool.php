<?php
/**
 * ApiVdiPoolsIdVdiPool
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
 * ApiVdiPoolsIdVdiPool Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ApiVdiPoolsIdVdiPool implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = '_api_vdi_pools__id__vdiPool';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'name' => 'string',
        'description' => 'string',
        'owner' => 'int',
        'min_idle' => 'float',
        'initial_pool_size' => 'float',
        'max_idle' => 'float',
        'max_pool_size' => 'float',
        'allocation_timeout_minutes' => 'float',
        'persistent_user' => 'bool',
        'recyclable' => 'bool',
        'allow_copy' => 'bool',
        'allow_printer' => 'bool',
        'allow_fileshare' => 'bool',
        'allow_hypervisor_console' => 'bool',
        'auto_create_local_user_on_reservation' => 'bool',
        'enabled' => 'bool',
        'icon_path' => 'string',
        'apps' => 'int[]',
        'gateway' => 'int',
        'instance_config' => 'string',
        'config' => '\OpenAPI\Client\Model\ApiVdiPoolsIdVdiPoolConfig',
        'guest_console_jump_host' => 'string',
        'guest_console_jump_port' => 'int',
        'guest_console_jump_username' => 'string',
        'guest_console_jump_password' => 'string',
        'guest_console_jump_keypair' => 'int'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'name' => null,
        'description' => null,
        'owner' => 'int64',
        'min_idle' => null,
        'initial_pool_size' => null,
        'max_idle' => null,
        'max_pool_size' => null,
        'allocation_timeout_minutes' => null,
        'persistent_user' => null,
        'recyclable' => null,
        'allow_copy' => null,
        'allow_printer' => null,
        'allow_fileshare' => null,
        'allow_hypervisor_console' => null,
        'auto_create_local_user_on_reservation' => null,
        'enabled' => null,
        'icon_path' => null,
        'apps' => 'int64',
        'gateway' => 'int64',
        'instance_config' => null,
        'config' => null,
        'guest_console_jump_host' => null,
        'guest_console_jump_port' => 'int64',
        'guest_console_jump_username' => null,
        'guest_console_jump_password' => null,
        'guest_console_jump_keypair' => 'int64'
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
        'name' => 'name',
        'description' => 'description',
        'owner' => 'owner',
        'min_idle' => 'minIdle',
        'initial_pool_size' => 'initialPoolSize',
        'max_idle' => 'maxIdle',
        'max_pool_size' => 'maxPoolSize',
        'allocation_timeout_minutes' => 'allocationTimeoutMinutes',
        'persistent_user' => 'persistentUser',
        'recyclable' => 'recyclable',
        'allow_copy' => 'allowCopy',
        'allow_printer' => 'allowPrinter',
        'allow_fileshare' => 'allowFileshare',
        'allow_hypervisor_console' => 'allowHypervisorConsole',
        'auto_create_local_user_on_reservation' => 'autoCreateLocalUserOnReservation',
        'enabled' => 'enabled',
        'icon_path' => 'iconPath',
        'apps' => 'apps',
        'gateway' => 'gateway',
        'instance_config' => 'instanceConfig',
        'config' => 'config',
        'guest_console_jump_host' => 'guestConsoleJumpHost',
        'guest_console_jump_port' => 'guestConsoleJumpPort',
        'guest_console_jump_username' => 'guestConsoleJumpUsername',
        'guest_console_jump_password' => 'guestConsoleJumpPassword',
        'guest_console_jump_keypair' => 'guestConsoleJumpKeypair'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'name' => 'setName',
        'description' => 'setDescription',
        'owner' => 'setOwner',
        'min_idle' => 'setMinIdle',
        'initial_pool_size' => 'setInitialPoolSize',
        'max_idle' => 'setMaxIdle',
        'max_pool_size' => 'setMaxPoolSize',
        'allocation_timeout_minutes' => 'setAllocationTimeoutMinutes',
        'persistent_user' => 'setPersistentUser',
        'recyclable' => 'setRecyclable',
        'allow_copy' => 'setAllowCopy',
        'allow_printer' => 'setAllowPrinter',
        'allow_fileshare' => 'setAllowFileshare',
        'allow_hypervisor_console' => 'setAllowHypervisorConsole',
        'auto_create_local_user_on_reservation' => 'setAutoCreateLocalUserOnReservation',
        'enabled' => 'setEnabled',
        'icon_path' => 'setIconPath',
        'apps' => 'setApps',
        'gateway' => 'setGateway',
        'instance_config' => 'setInstanceConfig',
        'config' => 'setConfig',
        'guest_console_jump_host' => 'setGuestConsoleJumpHost',
        'guest_console_jump_port' => 'setGuestConsoleJumpPort',
        'guest_console_jump_username' => 'setGuestConsoleJumpUsername',
        'guest_console_jump_password' => 'setGuestConsoleJumpPassword',
        'guest_console_jump_keypair' => 'setGuestConsoleJumpKeypair'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'name' => 'getName',
        'description' => 'getDescription',
        'owner' => 'getOwner',
        'min_idle' => 'getMinIdle',
        'initial_pool_size' => 'getInitialPoolSize',
        'max_idle' => 'getMaxIdle',
        'max_pool_size' => 'getMaxPoolSize',
        'allocation_timeout_minutes' => 'getAllocationTimeoutMinutes',
        'persistent_user' => 'getPersistentUser',
        'recyclable' => 'getRecyclable',
        'allow_copy' => 'getAllowCopy',
        'allow_printer' => 'getAllowPrinter',
        'allow_fileshare' => 'getAllowFileshare',
        'allow_hypervisor_console' => 'getAllowHypervisorConsole',
        'auto_create_local_user_on_reservation' => 'getAutoCreateLocalUserOnReservation',
        'enabled' => 'getEnabled',
        'icon_path' => 'getIconPath',
        'apps' => 'getApps',
        'gateway' => 'getGateway',
        'instance_config' => 'getInstanceConfig',
        'config' => 'getConfig',
        'guest_console_jump_host' => 'getGuestConsoleJumpHost',
        'guest_console_jump_port' => 'getGuestConsoleJumpPort',
        'guest_console_jump_username' => 'getGuestConsoleJumpUsername',
        'guest_console_jump_password' => 'getGuestConsoleJumpPassword',
        'guest_console_jump_keypair' => 'getGuestConsoleJumpKeypair'
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
        $this->container['name'] = $data['name'] ?? null;
        $this->container['description'] = $data['description'] ?? null;
        $this->container['owner'] = $data['owner'] ?? null;
        $this->container['min_idle'] = $data['min_idle'] ?? null;
        $this->container['initial_pool_size'] = $data['initial_pool_size'] ?? null;
        $this->container['max_idle'] = $data['max_idle'] ?? null;
        $this->container['max_pool_size'] = $data['max_pool_size'] ?? null;
        $this->container['allocation_timeout_minutes'] = $data['allocation_timeout_minutes'] ?? null;
        $this->container['persistent_user'] = $data['persistent_user'] ?? false;
        $this->container['recyclable'] = $data['recyclable'] ?? false;
        $this->container['allow_copy'] = $data['allow_copy'] ?? false;
        $this->container['allow_printer'] = $data['allow_printer'] ?? false;
        $this->container['allow_fileshare'] = $data['allow_fileshare'] ?? false;
        $this->container['allow_hypervisor_console'] = $data['allow_hypervisor_console'] ?? false;
        $this->container['auto_create_local_user_on_reservation'] = $data['auto_create_local_user_on_reservation'] ?? false;
        $this->container['enabled'] = $data['enabled'] ?? true;
        $this->container['icon_path'] = $data['icon_path'] ?? null;
        $this->container['apps'] = $data['apps'] ?? null;
        $this->container['gateway'] = $data['gateway'] ?? null;
        $this->container['instance_config'] = $data['instance_config'] ?? null;
        $this->container['config'] = $data['config'] ?? null;
        $this->container['guest_console_jump_host'] = $data['guest_console_jump_host'] ?? null;
        $this->container['guest_console_jump_port'] = $data['guest_console_jump_port'] ?? null;
        $this->container['guest_console_jump_username'] = $data['guest_console_jump_username'] ?? null;
        $this->container['guest_console_jump_password'] = $data['guest_console_jump_password'] ?? null;
        $this->container['guest_console_jump_keypair'] = $data['guest_console_jump_keypair'] ?? null;
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
     * @param string|null $name Virtual Desktop name
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
     * @param string|null $description Virtual Desktop description
     *
     * @return self
     */
    public function setDescription($description)
    {
        $this->container['description'] = $description;

        return $this;
    }

    /**
     * Gets owner
     *
     * @return int|null
     */
    public function getOwner()
    {
        return $this->container['owner'];
    }

    /**
     * Sets owner
     *
     * @param int|null $owner Owner (User) ID
     *
     * @return self
     */
    public function setOwner($owner)
    {
        $this->container['owner'] = $owner;

        return $this;
    }

    /**
     * Gets min_idle
     *
     * @return float|null
     */
    public function getMinIdle()
    {
        return $this->container['min_idle'];
    }

    /**
     * Sets min_idle
     *
     * @param float|null $min_idle Min Idle - Sets the minimum number of idle instances on standby in the pool. The pool will always try to maintain this number of available instances on standby.
     *
     * @return self
     */
    public function setMinIdle($min_idle)
    {
        $this->container['min_idle'] = $min_idle;

        return $this;
    }

    /**
     * Gets initial_pool_size
     *
     * @return float|null
     */
    public function getInitialPoolSize()
    {
        return $this->container['initial_pool_size'];
    }

    /**
     * Sets initial_pool_size
     *
     * @param float|null $initial_pool_size The initial size of the pool to be allocated on creation
     *
     * @return self
     */
    public function setInitialPoolSize($initial_pool_size)
    {
        $this->container['initial_pool_size'] = $initial_pool_size;

        return $this;
    }

    /**
     * Gets max_idle
     *
     * @return float|null
     */
    public function getMaxIdle()
    {
        return $this->container['max_idle'];
    }

    /**
     * Sets max_idle
     *
     * @param float|null $max_idle Sets the maximum number of idle instances on standby in the pool. If the number of idle instances supersedes this, the pool will start removing instances
     *
     * @return self
     */
    public function setMaxIdle($max_idle)
    {
        $this->container['max_idle'] = $max_idle;

        return $this;
    }

    /**
     * Gets max_pool_size
     *
     * @return float|null
     */
    public function getMaxPoolSize()
    {
        return $this->container['max_pool_size'];
    }

    /**
     * Sets max_pool_size
     *
     * @param float|null $max_pool_size Max limit on number of allocations and instances within the pool.
     *
     * @return self
     */
    public function setMaxPoolSize($max_pool_size)
    {
        $this->container['max_pool_size'] = $max_pool_size;

        return $this;
    }

    /**
     * Gets allocation_timeout_minutes
     *
     * @return float|null
     */
    public function getAllocationTimeoutMinutes()
    {
        return $this->container['allocation_timeout_minutes'];
    }

    /**
     * Sets allocation_timeout_minutes
     *
     * @param float|null $allocation_timeout_minutes Time (in minutes) after a user disconnects before an allocation is recycled or shutdown depending on persistence.
     *
     * @return self
     */
    public function setAllocationTimeoutMinutes($allocation_timeout_minutes)
    {
        $this->container['allocation_timeout_minutes'] = $allocation_timeout_minutes;

        return $this;
    }

    /**
     * Gets persistent_user
     *
     * @return bool|null
     */
    public function getPersistentUser()
    {
        return $this->container['persistent_user'];
    }

    /**
     * Sets persistent_user
     *
     * @param bool|null $persistent_user Persistent Desktop Pool
     *
     * @return self
     */
    public function setPersistentUser($persistent_user)
    {
        $this->container['persistent_user'] = $persistent_user;

        return $this;
    }

    /**
     * Gets recyclable
     *
     * @return bool|null
     */
    public function getRecyclable()
    {
        return $this->container['recyclable'];
    }

    /**
     * Sets recyclable
     *
     * @param bool|null $recyclable Recyclable VDI Pools only work with cloud types that support snapshot management (i.e. Vmware, Nutanix, VCD)
     *
     * @return self
     */
    public function setRecyclable($recyclable)
    {
        $this->container['recyclable'] = $recyclable;

        return $this;
    }

    /**
     * Gets allow_copy
     *
     * @return bool|null
     */
    public function getAllowCopy()
    {
        return $this->container['allow_copy'];
    }

    /**
     * Sets allow_copy
     *
     * @param bool|null $allow_copy Allow copy from desktop
     *
     * @return self
     */
    public function setAllowCopy($allow_copy)
    {
        $this->container['allow_copy'] = $allow_copy;

        return $this;
    }

    /**
     * Gets allow_printer
     *
     * @return bool|null
     */
    public function getAllowPrinter()
    {
        return $this->container['allow_printer'];
    }

    /**
     * Sets allow_printer
     *
     * @param bool|null $allow_printer Allow local printers from Desktop
     *
     * @return self
     */
    public function setAllowPrinter($allow_printer)
    {
        $this->container['allow_printer'] = $allow_printer;

        return $this;
    }

    /**
     * Gets allow_fileshare
     *
     * @return bool|null
     */
    public function getAllowFileshare()
    {
        return $this->container['allow_fileshare'];
    }

    /**
     * Sets allow_fileshare
     *
     * @param bool|null $allow_fileshare Allow File Share
     *
     * @return self
     */
    public function setAllowFileshare($allow_fileshare)
    {
        $this->container['allow_fileshare'] = $allow_fileshare;

        return $this;
    }

    /**
     * Gets allow_hypervisor_console
     *
     * @return bool|null
     */
    public function getAllowHypervisorConsole()
    {
        return $this->container['allow_hypervisor_console'];
    }

    /**
     * Sets allow_hypervisor_console
     *
     * @param bool|null $allow_hypervisor_console Allow hypervisor console
     *
     * @return self
     */
    public function setAllowHypervisorConsole($allow_hypervisor_console)
    {
        $this->container['allow_hypervisor_console'] = $allow_hypervisor_console;

        return $this;
    }

    /**
     * Gets auto_create_local_user_on_reservation
     *
     * @return bool|null
     */
    public function getAutoCreateLocalUserOnReservation()
    {
        return $this->container['auto_create_local_user_on_reservation'];
    }

    /**
     * Sets auto_create_local_user_on_reservation
     *
     * @param bool|null $auto_create_local_user_on_reservation Auto create local user upon reservation
     *
     * @return self
     */
    public function setAutoCreateLocalUserOnReservation($auto_create_local_user_on_reservation)
    {
        $this->container['auto_create_local_user_on_reservation'] = $auto_create_local_user_on_reservation;

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
     * @param bool|null $enabled Can be used to enable or disable the VDI pool
     *
     * @return self
     */
    public function setEnabled($enabled)
    {
        $this->container['enabled'] = $enabled;

        return $this;
    }

    /**
     * Gets icon_path
     *
     * @return string|null
     */
    public function getIconPath()
    {
        return $this->container['icon_path'];
    }

    /**
     * Sets icon_path
     *
     * @param string|null $icon_path The relative location of an icon image
     *
     * @return self
     */
    public function setIconPath($icon_path)
    {
        $this->container['icon_path'] = $icon_path;

        return $this;
    }

    /**
     * Gets apps
     *
     * @return int[]|null
     */
    public function getApps()
    {
        return $this->container['apps'];
    }

    /**
     * Sets apps
     *
     * @param int[]|null $apps Array of VDI App IDs
     *
     * @return self
     */
    public function setApps($apps)
    {
        $this->container['apps'] = $apps;

        return $this;
    }

    /**
     * Gets gateway
     *
     * @return int|null
     */
    public function getGateway()
    {
        return $this->container['gateway'];
    }

    /**
     * Sets gateway
     *
     * @param int|null $gateway VDI Gateway ID
     *
     * @return self
     */
    public function setGateway($gateway)
    {
        $this->container['gateway'] = $gateway;

        return $this;
    }

    /**
     * Gets instance_config
     *
     * @return string|null
     */
    public function getInstanceConfig()
    {
        return $this->container['instance_config'];
    }

    /**
     * Sets instance_config
     *
     * @param string|null $instance_config Instance Config JSON. Passing as a string will preserve property order.  See `config` object for required values.
     *
     * @return self
     */
    public function setInstanceConfig($instance_config)
    {
        $this->container['instance_config'] = $instance_config;

        return $this;
    }

    /**
     * Gets config
     *
     * @return \OpenAPI\Client\Model\ApiVdiPoolsIdVdiPoolConfig|null
     */
    public function getConfig()
    {
        return $this->container['config'];
    }

    /**
     * Sets config
     *
     * @param \OpenAPI\Client\Model\ApiVdiPoolsIdVdiPoolConfig|null $config config
     *
     * @return self
     */
    public function setConfig($config)
    {
        $this->container['config'] = $config;

        return $this;
    }

    /**
     * Gets guest_console_jump_host
     *
     * @return string|null
     */
    public function getGuestConsoleJumpHost()
    {
        return $this->container['guest_console_jump_host'];
    }

    /**
     * Sets guest_console_jump_host
     *
     * @param string|null $guest_console_jump_host Guest Console Jump Host
     *
     * @return self
     */
    public function setGuestConsoleJumpHost($guest_console_jump_host)
    {
        $this->container['guest_console_jump_host'] = $guest_console_jump_host;

        return $this;
    }

    /**
     * Gets guest_console_jump_port
     *
     * @return int|null
     */
    public function getGuestConsoleJumpPort()
    {
        return $this->container['guest_console_jump_port'];
    }

    /**
     * Sets guest_console_jump_port
     *
     * @param int|null $guest_console_jump_port Guest Console Jump Port
     *
     * @return self
     */
    public function setGuestConsoleJumpPort($guest_console_jump_port)
    {
        $this->container['guest_console_jump_port'] = $guest_console_jump_port;

        return $this;
    }

    /**
     * Gets guest_console_jump_username
     *
     * @return string|null
     */
    public function getGuestConsoleJumpUsername()
    {
        return $this->container['guest_console_jump_username'];
    }

    /**
     * Sets guest_console_jump_username
     *
     * @param string|null $guest_console_jump_username Guest Console Jump Username
     *
     * @return self
     */
    public function setGuestConsoleJumpUsername($guest_console_jump_username)
    {
        $this->container['guest_console_jump_username'] = $guest_console_jump_username;

        return $this;
    }

    /**
     * Gets guest_console_jump_password
     *
     * @return string|null
     */
    public function getGuestConsoleJumpPassword()
    {
        return $this->container['guest_console_jump_password'];
    }

    /**
     * Sets guest_console_jump_password
     *
     * @param string|null $guest_console_jump_password Guest Console Jump Password
     *
     * @return self
     */
    public function setGuestConsoleJumpPassword($guest_console_jump_password)
    {
        $this->container['guest_console_jump_password'] = $guest_console_jump_password;

        return $this;
    }

    /**
     * Gets guest_console_jump_keypair
     *
     * @return int|null
     */
    public function getGuestConsoleJumpKeypair()
    {
        return $this->container['guest_console_jump_keypair'];
    }

    /**
     * Sets guest_console_jump_keypair
     *
     * @param int|null $guest_console_jump_keypair Guest Console Jump Key Pair. see `Key Pair`
     *
     * @return self
     */
    public function setGuestConsoleJumpKeypair($guest_console_jump_keypair)
    {
        $this->container['guest_console_jump_keypair'] = $guest_console_jump_keypair;

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


