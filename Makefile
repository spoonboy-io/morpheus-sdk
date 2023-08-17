all-sdk: go-sdk js-sdk powershell-sdk java-sdk grooy-sdk python-sdk php-sdk

validate:
	docker run --rm \
	  	-v ${PWD}:/local openapitools/openapi-generator-cli validate \
      	-i https://raw.githubusercontent.com/gomorpheus/morpheus-openapi/main/openapi.yaml

list-generators:
	docker run --rm \
	  	-v ${PWD}:/local openapitools/openapi-generator-cli list

go-sdk:
	echo "Creating Go SDK"
	docker run --rm \
  		-v ${PWD}:/local openapitools/openapi-generator-cli generate --skip-validate-spec \
  		-i https://raw.githubusercontent.com/gomorpheus/morpheus-openapi/main/openapi.yaml \
  		-g go \
  		-o /local/go

js-sdk:
	echo "Creating JavaScript SDK"
	docker run --rm \
  		-v ${PWD}:/local openapitools/openapi-generator-cli generate --skip-validate-spec \
  		-i https://raw.githubusercontent.com/gomorpheus/morpheus-openapi/main/openapi.yaml \
  		-g javascript \
  		-o /local/javascript

python-sdk:
	echo "Creating Python SDK"
	docker run --rm \
  		-v ${PWD}:/local openapitools/openapi-generator-cli generate --skip-validate-spec \
  		-i https://raw.githubusercontent.com/gomorpheus/morpheus-openapi/main/openapi.yaml \
  		-g python \
  		-o /local/python

php-sdk:
	echo "Creating PHP SDK"
	docker run --rm \
  		-v ${PWD}:/local openapitools/openapi-generator-cli generate --skip-validate-spec \
  		-i https://raw.githubusercontent.com/gomorpheus/morpheus-openapi/main/openapi.yaml \
  		-g php \
  		-o /local/php

powershell-sdk:
	echo "Creating PowerShell SDK"
	docker run --rm \
  		-v ${PWD}:/local openapitools/openapi-generator-cli generate --skip-validate-spec \
  		-i https://raw.githubusercontent.com/gomorpheus/morpheus-openapi/main/openapi.yaml \
  		-g powershell \
  		-o /local/powershell

java-sdk:
	echo "Creating Java SDK"
	docker run --rm \
  		-v ${PWD}:/local openapitools/openapi-generator-cli generate --skip-validate-spec \
  		-i https://raw.githubusercontent.com/gomorpheus/morpheus-openapi/main/openapi.yaml \
  		-g java \
  		-o /local/java

groovy-sdk:
	echo "Creating Groovy SDK"
	docker run --rm \
  		-v ${PWD}:/local openapitools/openapi-generator-cli generate --skip-validate-spec \
  		-i https://raw.githubusercontent.com/gomorpheus/morpheus-openapi/main/openapi.yaml \
  		-g groovy \
  		-o /local/groovy
