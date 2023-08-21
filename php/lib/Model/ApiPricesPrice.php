<?php
/**
 * ApiPricesPrice
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
 * ApiPricesPrice Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ApiPricesPrice implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = '_api_prices_price';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'name' => 'string',
        'code' => 'string',
        'account' => '\OpenAPI\Client\Model\ApiPricesPriceAccount',
        'price_type' => 'string',
        'price_unit' => 'string',
        'incur_charges' => 'string',
        'currency' => 'string',
        'cost' => 'float',
        'markup_type' => 'string',
        'markup' => 'float',
        'markup_percent' => 'float',
        'custom_price' => 'float',
        'platform' => 'string',
        'software' => 'string',
        'volume_type' => '\OpenAPI\Client\Model\ApiPricesPriceVolumeType',
        'datastore' => '\OpenAPI\Client\Model\ApiPricesPriceDatastore',
        'cross_cloud_apply' => 'bool'
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
        'code' => null,
        'account' => null,
        'price_type' => null,
        'price_unit' => null,
        'incur_charges' => null,
        'currency' => null,
        'cost' => null,
        'markup_type' => null,
        'markup' => null,
        'markup_percent' => null,
        'custom_price' => null,
        'platform' => null,
        'software' => null,
        'volume_type' => null,
        'datastore' => null,
        'cross_cloud_apply' => null
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
        'code' => 'code',
        'account' => 'account',
        'price_type' => 'priceType',
        'price_unit' => 'priceUnit',
        'incur_charges' => 'incurCharges',
        'currency' => 'currency',
        'cost' => 'cost',
        'markup_type' => 'markupType',
        'markup' => 'markup',
        'markup_percent' => 'markupPercent',
        'custom_price' => 'customPrice',
        'platform' => 'platform',
        'software' => 'software',
        'volume_type' => 'volumeType',
        'datastore' => 'datastore',
        'cross_cloud_apply' => 'crossCloudApply'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'name' => 'setName',
        'code' => 'setCode',
        'account' => 'setAccount',
        'price_type' => 'setPriceType',
        'price_unit' => 'setPriceUnit',
        'incur_charges' => 'setIncurCharges',
        'currency' => 'setCurrency',
        'cost' => 'setCost',
        'markup_type' => 'setMarkupType',
        'markup' => 'setMarkup',
        'markup_percent' => 'setMarkupPercent',
        'custom_price' => 'setCustomPrice',
        'platform' => 'setPlatform',
        'software' => 'setSoftware',
        'volume_type' => 'setVolumeType',
        'datastore' => 'setDatastore',
        'cross_cloud_apply' => 'setCrossCloudApply'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'name' => 'getName',
        'code' => 'getCode',
        'account' => 'getAccount',
        'price_type' => 'getPriceType',
        'price_unit' => 'getPriceUnit',
        'incur_charges' => 'getIncurCharges',
        'currency' => 'getCurrency',
        'cost' => 'getCost',
        'markup_type' => 'getMarkupType',
        'markup' => 'getMarkup',
        'markup_percent' => 'getMarkupPercent',
        'custom_price' => 'getCustomPrice',
        'platform' => 'getPlatform',
        'software' => 'getSoftware',
        'volume_type' => 'getVolumeType',
        'datastore' => 'getDatastore',
        'cross_cloud_apply' => 'getCrossCloudApply'
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

    const PRICE_TYPE_FIXED = 'fixed';
    const PRICE_TYPE_COMPUTE = 'compute';
    const PRICE_TYPE_MEMORY = 'memory';
    const PRICE_TYPE_CORES = 'cores';
    const PRICE_TYPE_STORAGE = 'storage';
    const PRICE_TYPE_DATASTORE = 'datastore';
    const PRICE_TYPE_PLATFORM = 'platform';
    const PRICE_TYPE_SOFTWARE = 'software';
    const PRICE_TYPE_LOAD_BALANCER = 'load_balancer';
    const PRICE_TYPE_LOAD_BALANCER_VIRTUAL_SERVER = 'load_balancer_virtual_server';
    const PRICE_UNIT_MINUTE = 'minute';
    const PRICE_UNIT_HOUR = 'hour';
    const PRICE_UNIT_DAY = 'day';
    const PRICE_UNIT_MONTH = 'month';
    const PRICE_UNIT_YEAR = 'year';
    const PRICE_UNIT_TWO_YEAR = 'two year';
    const PRICE_UNIT_THREE_YEAR = 'three year';
    const PRICE_UNIT_FOUR_YEAR = 'four year';
    const PRICE_UNIT_FIVE_YEAR = 'five year';
    const INCUR_CHARGES_RUNNING = 'running';
    const INCUR_CHARGES_STOPPED = 'stopped';
    const INCUR_CHARGES_ALWAYS = 'always';
    const MARKUP_TYPE_FIXED = 'fixed';
    const MARKUP_TYPE_PERCENT = 'percent';
    const MARKUP_TYPE_CUSTOM = 'custom';
    

    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getPriceTypeAllowableValues()
    {
        return [
            self::PRICE_TYPE_FIXED,
            self::PRICE_TYPE_COMPUTE,
            self::PRICE_TYPE_MEMORY,
            self::PRICE_TYPE_CORES,
            self::PRICE_TYPE_STORAGE,
            self::PRICE_TYPE_DATASTORE,
            self::PRICE_TYPE_PLATFORM,
            self::PRICE_TYPE_SOFTWARE,
            self::PRICE_TYPE_LOAD_BALANCER,
            self::PRICE_TYPE_LOAD_BALANCER_VIRTUAL_SERVER,
        ];
    }
    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getPriceUnitAllowableValues()
    {
        return [
            self::PRICE_UNIT_MINUTE,
            self::PRICE_UNIT_HOUR,
            self::PRICE_UNIT_DAY,
            self::PRICE_UNIT_MONTH,
            self::PRICE_UNIT_YEAR,
            self::PRICE_UNIT_TWO_YEAR,
            self::PRICE_UNIT_THREE_YEAR,
            self::PRICE_UNIT_FOUR_YEAR,
            self::PRICE_UNIT_FIVE_YEAR,
        ];
    }
    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getIncurChargesAllowableValues()
    {
        return [
            self::INCUR_CHARGES_RUNNING,
            self::INCUR_CHARGES_STOPPED,
            self::INCUR_CHARGES_ALWAYS,
        ];
    }
    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getMarkupTypeAllowableValues()
    {
        return [
            self::MARKUP_TYPE_FIXED,
            self::MARKUP_TYPE_PERCENT,
            self::MARKUP_TYPE_CUSTOM,
        ];
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
        $this->container['code'] = $data['code'] ?? null;
        $this->container['account'] = $data['account'] ?? null;
        $this->container['price_type'] = $data['price_type'] ?? null;
        $this->container['price_unit'] = $data['price_unit'] ?? null;
        $this->container['incur_charges'] = $data['incur_charges'] ?? null;
        $this->container['currency'] = $data['currency'] ?? null;
        $this->container['cost'] = $data['cost'] ?? null;
        $this->container['markup_type'] = $data['markup_type'] ?? null;
        $this->container['markup'] = $data['markup'] ?? null;
        $this->container['markup_percent'] = $data['markup_percent'] ?? null;
        $this->container['custom_price'] = $data['custom_price'] ?? null;
        $this->container['platform'] = $data['platform'] ?? null;
        $this->container['software'] = $data['software'] ?? null;
        $this->container['volume_type'] = $data['volume_type'] ?? null;
        $this->container['datastore'] = $data['datastore'] ?? null;
        $this->container['cross_cloud_apply'] = $data['cross_cloud_apply'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['name'] === null) {
            $invalidProperties[] = "'name' can't be null";
        }
        if ($this->container['code'] === null) {
            $invalidProperties[] = "'code' can't be null";
        }
        if ($this->container['price_type'] === null) {
            $invalidProperties[] = "'price_type' can't be null";
        }
        $allowedValues = $this->getPriceTypeAllowableValues();
        if (!is_null($this->container['price_type']) && !in_array($this->container['price_type'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'price_type', must be one of '%s'",
                $this->container['price_type'],
                implode("', '", $allowedValues)
            );
        }

        if ($this->container['price_unit'] === null) {
            $invalidProperties[] = "'price_unit' can't be null";
        }
        $allowedValues = $this->getPriceUnitAllowableValues();
        if (!is_null($this->container['price_unit']) && !in_array($this->container['price_unit'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'price_unit', must be one of '%s'",
                $this->container['price_unit'],
                implode("', '", $allowedValues)
            );
        }

        if ($this->container['incur_charges'] === null) {
            $invalidProperties[] = "'incur_charges' can't be null";
        }
        $allowedValues = $this->getIncurChargesAllowableValues();
        if (!is_null($this->container['incur_charges']) && !in_array($this->container['incur_charges'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'incur_charges', must be one of '%s'",
                $this->container['incur_charges'],
                implode("', '", $allowedValues)
            );
        }

        if ($this->container['currency'] === null) {
            $invalidProperties[] = "'currency' can't be null";
        }
        if ($this->container['cost'] === null) {
            $invalidProperties[] = "'cost' can't be null";
        }
        $allowedValues = $this->getMarkupTypeAllowableValues();
        if (!is_null($this->container['markup_type']) && !in_array($this->container['markup_type'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'markup_type', must be one of '%s'",
                $this->container['markup_type'],
                implode("', '", $allowedValues)
            );
        }

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
     * @return string
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     *
     * @param string $name Price name
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
     * @return string
     */
    public function getCode()
    {
        return $this->container['code'];
    }

    /**
     * Sets code
     *
     * @param string $code Price code, must be unique
     *
     * @return self
     */
    public function setCode($code)
    {
        $this->container['code'] = $code;

        return $this;
    }

    /**
     * Gets account
     *
     * @return \OpenAPI\Client\Model\ApiPricesPriceAccount|null
     */
    public function getAccount()
    {
        return $this->container['account'];
    }

    /**
     * Sets account
     *
     * @param \OpenAPI\Client\Model\ApiPricesPriceAccount|null $account account
     *
     * @return self
     */
    public function setAccount($account)
    {
        $this->container['account'] = $account;

        return $this;
    }

    /**
     * Gets price_type
     *
     * @return string
     */
    public function getPriceType()
    {
        return $this->container['price_type'];
    }

    /**
     * Sets price_type
     *
     * @param string $price_type Restricts query to only load only prices with specified priceType. * `fixed` - Everything * `compute` - Memory + CPU * `memory` - Memory * `cores` - Cores * `storage` - Storage * `datastore` - Datastore * `platform` - Platform * `software` - Software * `load_balancer` - Load Balancer * `load_balancer_virtual_server` - Load Balancer Virtual Server
     *
     * @return self
     */
    public function setPriceType($price_type)
    {
        $allowedValues = $this->getPriceTypeAllowableValues();
        if (!in_array($price_type, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'price_type', must be one of '%s'",
                    $price_type,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['price_type'] = $price_type;

        return $this;
    }

    /**
     * Gets price_unit
     *
     * @return string
     */
    public function getPriceUnit()
    {
        return $this->container['price_unit'];
    }

    /**
     * Sets price_unit
     *
     * @param string $price_unit The unit of pricing
     *
     * @return self
     */
    public function setPriceUnit($price_unit)
    {
        $allowedValues = $this->getPriceUnitAllowableValues();
        if (!in_array($price_unit, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'price_unit', must be one of '%s'",
                    $price_unit,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['price_unit'] = $price_unit;

        return $this;
    }

    /**
     * Gets incur_charges
     *
     * @return string
     */
    public function getIncurCharges()
    {
        return $this->container['incur_charges'];
    }

    /**
     * Sets incur_charges
     *
     * @param string $incur_charges Indicates when to incur charge
     *
     * @return self
     */
    public function setIncurCharges($incur_charges)
    {
        $allowedValues = $this->getIncurChargesAllowableValues();
        if (!in_array($incur_charges, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'incur_charges', must be one of '%s'",
                    $incur_charges,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['incur_charges'] = $incur_charges;

        return $this;
    }

    /**
     * Gets currency
     *
     * @return string
     */
    public function getCurrency()
    {
        return $this->container['currency'];
    }

    /**
     * Sets currency
     *
     * @param string $currency ISO Currency code
     *
     * @return self
     */
    public function setCurrency($currency)
    {
        $this->container['currency'] = $currency;

        return $this;
    }

    /**
     * Gets cost
     *
     * @return float
     */
    public function getCost()
    {
        return $this->container['cost'];
    }

    /**
     * Sets cost
     *
     * @param float $cost Cost
     *
     * @return self
     */
    public function setCost($cost)
    {
        $this->container['cost'] = $cost;

        return $this;
    }

    /**
     * Gets markup_type
     *
     * @return string|null
     */
    public function getMarkupType()
    {
        return $this->container['markup_type'];
    }

    /**
     * Sets markup_type
     *
     * @param string|null $markup_type Price adjustment type
     *
     * @return self
     */
    public function setMarkupType($markup_type)
    {
        $allowedValues = $this->getMarkupTypeAllowableValues();
        if (!is_null($markup_type) && !in_array($markup_type, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'markup_type', must be one of '%s'",
                    $markup_type,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['markup_type'] = $markup_type;

        return $this;
    }

    /**
     * Gets markup
     *
     * @return float|null
     */
    public function getMarkup()
    {
        return $this->container['markup'];
    }

    /**
     * Sets markup
     *
     * @param float|null $markup Amount for `fixed` price adjustment type
     *
     * @return self
     */
    public function setMarkup($markup)
    {
        $this->container['markup'] = $markup;

        return $this;
    }

    /**
     * Gets markup_percent
     *
     * @return float|null
     */
    public function getMarkupPercent()
    {
        return $this->container['markup_percent'];
    }

    /**
     * Sets markup_percent
     *
     * @param float|null $markup_percent Percent for `percent` price adjustment type
     *
     * @return self
     */
    public function setMarkupPercent($markup_percent)
    {
        $this->container['markup_percent'] = $markup_percent;

        return $this;
    }

    /**
     * Gets custom_price
     *
     * @return float|null
     */
    public function getCustomPrice()
    {
        return $this->container['custom_price'];
    }

    /**
     * Sets custom_price
     *
     * @param float|null $custom_price Custom price for `custom` price adjustment type
     *
     * @return self
     */
    public function setCustomPrice($custom_price)
    {
        $this->container['custom_price'] = $custom_price;

        return $this;
    }

    /**
     * Gets platform
     *
     * @return string|null
     */
    public function getPlatform()
    {
        return $this->container['platform'];
    }

    /**
     * Sets platform
     *
     * @param string|null $platform Platform.  Required for `platform` price type
     *
     * @return self
     */
    public function setPlatform($platform)
    {
        $this->container['platform'] = $platform;

        return $this;
    }

    /**
     * Gets software
     *
     * @return string|null
     */
    public function getSoftware()
    {
        return $this->container['software'];
    }

    /**
     * Sets software
     *
     * @param string|null $software Software.  Required for software price type
     *
     * @return self
     */
    public function setSoftware($software)
    {
        $this->container['software'] = $software;

        return $this;
    }

    /**
     * Gets volume_type
     *
     * @return \OpenAPI\Client\Model\ApiPricesPriceVolumeType|null
     */
    public function getVolumeType()
    {
        return $this->container['volume_type'];
    }

    /**
     * Sets volume_type
     *
     * @param \OpenAPI\Client\Model\ApiPricesPriceVolumeType|null $volume_type volume_type
     *
     * @return self
     */
    public function setVolumeType($volume_type)
    {
        $this->container['volume_type'] = $volume_type;

        return $this;
    }

    /**
     * Gets datastore
     *
     * @return \OpenAPI\Client\Model\ApiPricesPriceDatastore|null
     */
    public function getDatastore()
    {
        return $this->container['datastore'];
    }

    /**
     * Sets datastore
     *
     * @param \OpenAPI\Client\Model\ApiPricesPriceDatastore|null $datastore datastore
     *
     * @return self
     */
    public function setDatastore($datastore)
    {
        $this->container['datastore'] = $datastore;

        return $this;
    }

    /**
     * Gets cross_cloud_apply
     *
     * @return bool|null
     */
    public function getCrossCloudApply()
    {
        return $this->container['cross_cloud_apply'];
    }

    /**
     * Sets cross_cloud_apply
     *
     * @param bool|null $cross_cloud_apply Apply price across clouds, optional true/false flag for datastore price type
     *
     * @return self
     */
    public function setCrossCloudApply($cross_cloud_apply)
    {
        $this->container['cross_cloud_apply'] = $cross_cloud_apply;

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


