## Morpheus SDK

Experimental. Based on version 6.1.1

Using the Morpheus OpenAPI specification to generate SDK's for popular client languages.

Uses the [OpenAPI Generator](https://openapi-generator.tech/) to build the SDKs via Docker.

Note that ATM the validation of the file API fails and two targets (Python and PHP fail to build)

- Go
- JavaScript
- Java
- Groovy
- Python (does not build)
- PHP (does not build)
- PowerShell

### Using

This is experimental. At this time it runs only against the [main branch of the API specfication](https://github.com/gomorpheus/morpheus-openapi).

There's no automation on this currently, so to generate the SDKs use the `Makefile` which uses Docker.

#### Validate

```shell
make validate
```

#### List Generators

Returns list of generators and their status. Use to create new SDK builds in the `Makefile`

```shell
make list-generators
```

#### Make SDK
Consult the Makefile
```shell
## Create the Javascript SDK
make javascript-sdk
```


### Todo
- Test the SDKs
- Ability to target a branch and build API for specific version of Morpheus, not just latest