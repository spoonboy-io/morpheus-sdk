<?php
/**
 * InlineResponse2004
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
 * InlineResponse2004 Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class InlineResponse2004 implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'inline_response_200_4';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'archive_bucket' => '\OpenAPI\Client\Model\ArchiveBucket',
        'is_owner' => 'bool',
        'parent_directory' => 'string',
        'archive_files' => '\OpenAPI\Client\Model\ArchiveBucketFile[]',
        'archive_file_count' => 'int'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'archive_bucket' => null,
        'is_owner' => null,
        'parent_directory' => null,
        'archive_files' => null,
        'archive_file_count' => 'int64'
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
        'archive_bucket' => 'archiveBucket',
        'is_owner' => 'isOwner',
        'parent_directory' => 'parentDirectory',
        'archive_files' => 'archiveFiles',
        'archive_file_count' => 'archiveFileCount'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'archive_bucket' => 'setArchiveBucket',
        'is_owner' => 'setIsOwner',
        'parent_directory' => 'setParentDirectory',
        'archive_files' => 'setArchiveFiles',
        'archive_file_count' => 'setArchiveFileCount'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'archive_bucket' => 'getArchiveBucket',
        'is_owner' => 'getIsOwner',
        'parent_directory' => 'getParentDirectory',
        'archive_files' => 'getArchiveFiles',
        'archive_file_count' => 'getArchiveFileCount'
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
        $this->container['archive_bucket'] = $data['archive_bucket'] ?? null;
        $this->container['is_owner'] = $data['is_owner'] ?? null;
        $this->container['parent_directory'] = $data['parent_directory'] ?? null;
        $this->container['archive_files'] = $data['archive_files'] ?? null;
        $this->container['archive_file_count'] = $data['archive_file_count'] ?? null;
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
     * Gets archive_bucket
     *
     * @return \OpenAPI\Client\Model\ArchiveBucket|null
     */
    public function getArchiveBucket()
    {
        return $this->container['archive_bucket'];
    }

    /**
     * Sets archive_bucket
     *
     * @param \OpenAPI\Client\Model\ArchiveBucket|null $archive_bucket archive_bucket
     *
     * @return self
     */
    public function setArchiveBucket($archive_bucket)
    {
        $this->container['archive_bucket'] = $archive_bucket;

        return $this;
    }

    /**
     * Gets is_owner
     *
     * @return bool|null
     */
    public function getIsOwner()
    {
        return $this->container['is_owner'];
    }

    /**
     * Sets is_owner
     *
     * @param bool|null $is_owner is_owner
     *
     * @return self
     */
    public function setIsOwner($is_owner)
    {
        $this->container['is_owner'] = $is_owner;

        return $this;
    }

    /**
     * Gets parent_directory
     *
     * @return string|null
     */
    public function getParentDirectory()
    {
        return $this->container['parent_directory'];
    }

    /**
     * Sets parent_directory
     *
     * @param string|null $parent_directory parent_directory
     *
     * @return self
     */
    public function setParentDirectory($parent_directory)
    {
        $this->container['parent_directory'] = $parent_directory;

        return $this;
    }

    /**
     * Gets archive_files
     *
     * @return \OpenAPI\Client\Model\ArchiveBucketFile[]|null
     */
    public function getArchiveFiles()
    {
        return $this->container['archive_files'];
    }

    /**
     * Sets archive_files
     *
     * @param \OpenAPI\Client\Model\ArchiveBucketFile[]|null $archive_files archive_files
     *
     * @return self
     */
    public function setArchiveFiles($archive_files)
    {
        $this->container['archive_files'] = $archive_files;

        return $this;
    }

    /**
     * Gets archive_file_count
     *
     * @return int|null
     */
    public function getArchiveFileCount()
    {
        return $this->container['archive_file_count'];
    }

    /**
     * Sets archive_file_count
     *
     * @param int|null $archive_file_count archive_file_count
     *
     * @return self
     */
    public function setArchiveFileCount($archive_file_count)
    {
        $this->container['archive_file_count'] = $archive_file_count;

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


