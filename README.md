## Morpheus SDK

> Experimental. Based on version 6.2.1
> 
> Some SDKs can not be generated.
>
> There are currently many issues with the generated SDKs that have been built.

<hr>

Using the Morpheus OpenAPI specification to generate SDK's for popular client languages.

Uses the [OpenAPI Generator](https://openapi-generator.tech/) to build the SDKs via Docker.

Uses v5.0.0 to work past some issues, v6.0.0 works allows Python to be generated and Powershell partially.

Included language targets and their status below. 

| Language   | Generates | Tested | Good | Comment                                      |
|------------|---------|--------|-----|----------------------------------------------|
| Go         | ✅       | ✅      | ❌   | Issue #2, Issue #3                           |
| JavaScript | ✅       | ❌      |     |                                              |
| Java       | ✅       | ❌      |     |                                              |
| Groovy     |    ✅     | ❌      |     |                                              | 
| Python     | ✅        |   ❌     |     | Builds with 6.0.0                            |   
| PHP        | ✅       | ❌      |     | Issue #2                                     |
| PowerShell | 🤷        |    ❌     |     | OpenAPI 6.0.0 outputs much of it, but errors |

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