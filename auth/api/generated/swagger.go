// Package api_client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api_client

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xUwW7UQAz9lchwHDUL3HIDiQMS4gD0hDgMiZNMm4wH2+mqqvLvaCabzULTtBJF7Ckj",
	"23nP9rN9ByX1gTx6FSjuQMoWe5ue75mJ4yMwBWR1mMwlVRi/ehsQCnBesUGG0UCPIrY5dYqy8w2MowHG",
	"n4NjrKD4NkEs8d/NHE8/rrDUiPWVrtHfJ9fZvE0wha3hXgqu1IS9dV181MS9VSgOFvMnj4Hasegn26+V",
	"aaCzG06m7gnNWQhO4MwxoQTyUGUfqXH+r8oLVmRPXD2e5wxx/OOhpD5j40T/X9s3SnqqJinKPCbNRici",
	"oPM1RaoKpWQX1JGHAt4O2maCfOPKiKVOY0LwZW+bBjmzg7Zg4AZZpvhXF7uLXcycAnobHBTwJpkiv7ap",
	"r3kX5yBPDSfR+6RpTrIh7kICYhsdH6rZdTl5YgtQ9B1Vt9Pme0Wf4GwInSvTX/mVkF9OR3y9ZKyhgBf5",
	"clvyw2HJlzFNTfk9r+jMKqsWTvuvPGASRAJ5mUbn9W73bClNt2YlHbqGZKvt0Omz0U13dYUOZ8doIOfD",
	"1myoOC9W5nG/LuYc8Y/1PG742UiaCj4nRcfxVwAAAP//pyK442oHAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}