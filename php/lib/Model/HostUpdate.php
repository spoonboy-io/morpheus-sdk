<?php
/**
 * HostUpdate
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
 * HostUpdate Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class HostUpdate implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'hostUpdate';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'name' => 'string',
        'description' => 'string',
        'ssh_username' => 'string',
        'ssh_password' => 'string',
        'power_schedule_type' => 'int',
        'labels' => 'string[]',
        'tags' => '\OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]',
        'add_tags' => '\OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]',
        'remove_tags' => '\OpenAPI\Client\Model\InstanceUpdateInstanceRemoveTags[]',
        'guest_console_type' => 'string',
        'guest_console_username' => 'string',
        'guest_console_password' => 'string',
        'guest_console_port' => 'string',
        'guest_console_preferred' => 'bool'
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
        'ssh_username' => null,
        'ssh_password' => null,
        'power_schedule_type' => 'int64',
        'labels' => null,
        'tags' => null,
        'add_tags' => null,
        'remove_tags' => null,
        'guest_console_type' => null,
        'guest_console_username' => null,
        'guest_console_password' => null,
        'guest_console_port' => null,
        'guest_console_preferred' => null
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
        'ssh_username' => 'sshUsername',
        'ssh_password' => 'sshPassword',
        'power_schedule_type' => 'powerScheduleType',
        'labels' => 'labels',
        'tags' => 'tags',
        'add_tags' => 'addTags',
        'remove_tags' => 'removeTags',
        'guest_console_type' => 'guestConsoleType',
        'guest_console_username' => 'guestConsoleUsername',
        'guest_console_password' => 'guestConsolePassword',
        'guest_console_port' => 'guestConsolePort',
        'guest_console_preferred' => 'guestConsolePreferred'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'name' => 'setName',
        'description' => 'setDescription',
        'ssh_username' => 'setSshUsername',
        'ssh_password' => 'setSshPassword',
        'power_schedule_type' => 'setPowerScheduleType',
        'labels' => 'setLabels',
        'tags' => 'setTags',
        'add_tags' => 'setAddTags',
        'remove_tags' => 'setRemoveTags',
        'guest_console_type' => 'setGuestConsoleType',
        'guest_console_username' => 'setGuestConsoleUsername',
        'guest_console_password' => 'setGuestConsolePassword',
        'guest_console_port' => 'setGuestConsolePort',
        'guest_console_preferred' => 'setGuestConsolePreferred'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'name' => 'getName',
        'description' => 'getDescription',
        'ssh_username' => 'getSshUsername',
        'ssh_password' => 'getSshPassword',
        'power_schedule_type' => 'getPowerScheduleType',
        'labels' => 'getLabels',
        'tags' => 'getTags',
        'add_tags' => 'getAddTags',
        'remove_tags' => 'getRemoveTags',
        'guest_console_type' => 'getGuestConsoleType',
        'guest_console_username' => 'getGuestConsoleUsername',
        'guest_console_password' => 'getGuestConsolePassword',
        'guest_console_port' => 'getGuestConsolePort',
        'guest_console_preferred' => 'getGuestConsolePreferred'
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
        $this->container['ssh_username'] = $data['ssh_username'] ?? null;
        $this->container['ssh_password'] = $data['ssh_password'] ?? null;
        $this->container['power_schedule_type'] = $data['power_schedule_type'] ?? null;
        $this->container['labels'] = $data['labels'] ?? null;
        $this->container['tags'] = $data['tags'] ?? null;
        $this->container['add_tags'] = $data['add_tags'] ?? null;
        $this->container['remove_tags'] = $data['remove_tags'] ?? null;
        $this->container['guest_console_type'] = $data['guest_console_type'] ?? null;
        $this->container['guest_console_username'] = $data['guest_console_username'] ?? null;
        $this->container['guest_console_password'] = $data['guest_console_password'] ?? null;
        $this->container['guest_console_port'] = $data['guest_console_port'] ?? null;
        $this->container['guest_console_preferred'] = $data['guest_console_preferred'] ?? true;
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
     * @param string|null $name Unique name scoped to your account for the server.
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
     * @param string|null $description Optional description field.
     *
     * @return self
     */
    public function setDescription($description)
    {
        $this->container['description'] = $description;

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
     * @param string|null $ssh_username SSH Username
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
     * @param string|null $ssh_password SSH Password
     *
     * @return self
     */
    public function setSshPassword($ssh_password)
    {
        $this->container['ssh_password'] = $ssh_password;

        return $this;
    }

    /**
     * Gets power_schedule_type
     *
     * @return int|null
     */
    public function getPowerScheduleType()
    {
        return $this->container['power_schedule_type'];
    }

    /**
     * Sets power_schedule_type
     *
     * @param int|null $power_schedule_type Power schedule ID.
     *
     * @return self
     */
    public function setPowerScheduleType($power_schedule_type)
    {
        $this->container['power_schedule_type'] = $power_schedule_type;

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
     * Gets tags
     *
     * @return \OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]|null
     */
    public function getTags()
    {
        return $this->container['tags'];
    }

    /**
     * Sets tags
     *
     * @param \OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]|null $tags Metadata tags, Array of objects having a name and value.
     *
     * @return self
     */
    public function setTags($tags)
    {
        $this->container['tags'] = $tags;

        return $this;
    }

    /**
     * Gets add_tags
     *
     * @return \OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]|null
     */
    public function getAddTags()
    {
        return $this->container['add_tags'];
    }

    /**
     * Sets add_tags
     *
     * @param \OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]|null $add_tags Add or update value of Metadata tags, Array of objects having a name and value.
     *
     * @return self
     */
    public function setAddTags($add_tags)
    {
        $this->container['add_tags'] = $add_tags;

        return $this;
    }

    /**
     * Gets remove_tags
     *
     * @return \OpenAPI\Client\Model\InstanceUpdateInstanceRemoveTags[]|null
     */
    public function getRemoveTags()
    {
        return $this->container['remove_tags'];
    }

    /**
     * Sets remove_tags
     *
     * @param \OpenAPI\Client\Model\InstanceUpdateInstanceRemoveTags[]|null $remove_tags Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed.
     *
     * @return self
     */
    public function setRemoveTags($remove_tags)
    {
        $this->container['remove_tags'] = $remove_tags;

        return $this;
    }

    /**
     * Gets guest_console_type
     *
     * @return string|null
     */
    public function getGuestConsoleType()
    {
        return $this->container['guest_console_type'];
    }

    /**
     * Sets guest_console_type
     *
     * @param string|null $guest_console_type The Type of guest console this server provides such as disabled, vnc, rdp, ssh
     *
     * @return self
     */
    public function setGuestConsoleType($guest_console_type)
    {
        $this->container['guest_console_type'] = $guest_console_type;

        return $this;
    }

    /**
     * Gets guest_console_username
     *
     * @return string|null
     */
    public function getGuestConsoleUsername()
    {
        return $this->container['guest_console_username'];
    }

    /**
     * Sets guest_console_username
     *
     * @param string|null $guest_console_username The optional guest console username if you don't want to use the user defaults
     *
     * @return self
     */
    public function setGuestConsoleUsername($guest_console_username)
    {
        $this->container['guest_console_username'] = $guest_console_username;

        return $this;
    }

    /**
     * Gets guest_console_password
     *
     * @return string|null
     */
    public function getGuestConsolePassword()
    {
        return $this->container['guest_console_password'];
    }

    /**
     * Sets guest_console_password
     *
     * @param string|null $guest_console_password The optional guest console password if not using the accessing users creds
     *
     * @return self
     */
    public function setGuestConsolePassword($guest_console_password)
    {
        $this->container['guest_console_password'] = $guest_console_password;

        return $this;
    }

    /**
     * Gets guest_console_port
     *
     * @return string|null
     */
    public function getGuestConsolePort()
    {
        return $this->container['guest_console_port'];
    }

    /**
     * Sets guest_console_port
     *
     * @param string|null $guest_console_port The port the guest console is being accessed from
     *
     * @return self
     */
    public function setGuestConsolePort($guest_console_port)
    {
        $this->container['guest_console_port'] = $guest_console_port;

        return $this;
    }

    /**
     * Gets guest_console_preferred
     *
     * @return bool|null
     */
    public function getGuestConsolePreferred()
    {
        return $this->container['guest_console_preferred'];
    }

    /**
     * Sets guest_console_preferred
     *
     * @param bool|null $guest_console_preferred Can turn off guest console preferences on server in favor of hypervisor console
     *
     * @return self
     */
    public function setGuestConsolePreferred($guest_console_preferred)
    {
        $this->container['guest_console_preferred'] = $guest_console_preferred;

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

