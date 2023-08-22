all-sdk: go-sdk js-sdk powershell-sdk java-sdk grooy-sdk python-sdk php-sdk

bundled:
	rm -Rf ./morpheus-openapi
	git clone https://github.com/gomorpheus/morpheus-openapi.git
	docker run --rm -v "${PWD}/morpheus-openapi":/spec redocly/openapi-cli bundle openapi.yaml > bundled.yaml

validate:
	cd morpheus-openapi
	docker run --rm \
	  	-v ${PWD}:/local openapitools/openapi-generator-cli:v5.0.0 validate \
      	-i local/bundled.yaml

get-version:
	docker run --rm \
	  	-v ${PWD}:/local openapitools/openapi-generator-cli:v5.0.0 version

list-generators:
	docker run --rm \
	  	-v ${PWD}:/local openapitools/openapi-generator-cli list

clean-code:
	##find go -type f -name "*.go" -exec sed -i '' -e 's/localVarHTTPResponse.StatusCode == 4XX/localVarHTTPResponse.StatusCode >= 400 \&\& localVarHTTPResponse.StatusCode < 500/g' {} +
	##find go -type f -name "*.go" -exec sed -i '' -e 's/localVarHTTPResponse.StatusCode == 5XX/localVarHTTPResponse.StatusCode >= 500/g' {} +

go-sdk:
	echo "Creating Go SDK"
	rm -Rf ./go
	mkdir go
	docker run --rm \
  		-v ${PWD}:/local openapitools/openapi-generator-cli:v5.0.0 generate --skip-validate-spec \
  		-i local/bundled.yaml \
  		-g go \
  		-o /local/go
  	## fix the 4XX and 5XX issues
	find go -type f -name "*.go" -exec sed -i '' -e 's/localVarHTTPResponse.StatusCode == 4XX/localVarHTTPResponse.StatusCode >= 400 \&\& localVarHTTPResponse.StatusCode < 500/g' {} +
	find go -type f -name "*.go" -exec sed -i '' -e 's/localVarHTTPResponse.StatusCode == 5XX/localVarHTTPResponse.StatusCode >= 500/g' {} +

js-sdk:
	echo "Creating JavaScript SDK"
	rm -Rf ./javascript
	mkdir javascript
	docker run --rm \
  		-v ${PWD}:/local openapitools/openapi-generator-cli:v5.0.0 generate --skip-validate-spec \
  		-i local/bundled.yaml  \
  		-g javascript \
  		-o /local/javascript

python-sdk:
	echo "Creating Python SDK"
	rm -Rf ./python
	mkdir python
	docker run --rm \
  		-v ${PWD}:/local openapitools/openapi-generator-cli:v5.0.0 generate --skip-validate-spec \
  		-i local/bundled.yaml \
  		-g python \
  		-o /local/python

php-sdk:
	echo "Creating PHP SDK"
	rm -Rf ./php
	mkdir php
	docker run --rm \
  		-v ${PWD}:/local openapitools/openapi-generator-cli:v5.0.0 generate --skip-validate-spec \
  		-i local/bundled.yaml \
  		-g php \
  		-o /local/php

powershell-sdk:
	echo "Creating PowerShell SDK"
	rm -Rf ./powershell
	mkdir powershell
	docker run --rm \
  		-v ${PWD}:/local openapitools/openapi-generator-cli:v5.0.0 generate --skip-validate-spec \
  		-i local/bundled.yaml \
  		-g powershell \
  		-o /local/powershell

java-sdk:
	echo "Creating Java SDK"
	rm -Rf ./java
	mkdir java
	docker run --rm \
  		-v ${PWD}:/local openapitools/openapi-generator-cli:v5.0.0 generate --skip-validate-spec \
  		-i local/bundled.yaml \
  		-g java \
  		-o /local/java

groovy-sdk:
	echo "Creating Groovy SDK"
	rm -Rf ./groovy
	mkdir groovy
	docker run --rm \
  		-v ${PWD}:/local openapitools/openapi-generator-cli:v5.0.0 generate --skip-validate-spec\
  		-i local/bundled.yaml \
  		-g groovy \
  		-o /local/groovy
