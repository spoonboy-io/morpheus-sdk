## Morpheus SDK

> Experimental. Based on version 6.2.1
> 
> Note: there is currently an unresolved issue re invalid 4XX and 5XX status codes in the SDKs. These will cause compile 
> time or run time errors.
> This appears to be a bug with OpenAPI Generator.

Using the Morpheus OpenAPI specification to generate SDK's for popular client languages.

Uses the [OpenAPI Generator](https://openapi-generator.tech/) to build the SDKs via Docker.

Uses v5.0.0 to work past some issues in OpenAPI Generator while not moving too far from the 3.0.3 spec.

Included language targets and their status below. ATM two targets fail to build using OpenAPI Generator v5.0.0
- Go
- JavaScript
- Java
- Groovy
- Python (does not build)
- PHP
- PowerShell (does not build)

### Using

This is experimental. At this time it runs only against the [main branch of the API specfication](https://github.com/gomorpheus/morpheus-openapi).

There's no automation on this currently, so to generate the SDKs use the `Makefile` which uses Docker.


#### Build bundled.yaml

To obtain the latest complete spec we need to clone [Morpheus OpenAPI repo](https://github.com/gomorpheus/morpheus-openapi) 
and build the `bundled.yaml`. A version of the compiled `bundled.yaml` is included in the with this repository

All the commands in the Makefile target `bundled.yaml`. This command clones the repo and builds the file.

```shell
make bundled
```

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