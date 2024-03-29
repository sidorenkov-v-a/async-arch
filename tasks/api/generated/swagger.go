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

	"H4sIAAAAAAAC/+xWy24TSxD9lVHdu2zHc+ObKJkdkCyyigTJKopQZabsdOzpbrrLCVE0EkJI/E6EQLBA",
	"fMP4j9B0z/g5JgkEiQU7u6vqnFN1usu+gVTnRitS7CC5AYMWc2Ky/hujGx7sVZ+kggQM8jkIUJgTJE1Q",
	"gKVXY2kpg4TtmAS49JxyrKr+tdSHBP7pzki6Ieq6dXlRFE2Fp9y3VluvxGpDliX541Rn5BVdm4pbKqYB",
	"WSgE5OQcDuaDjq1UA6iAZ9JOAsQs/1Q0+frsglKusA6Hq8yOkcfubvg6rw32CF0LMDonB4ooDLivbY4M",
	"CYzHMgOxTCYgtYRM2RNeyM6QqcMyp7aSC2kxoK+OzZLRlsnek33tFASw5BG1RsYme5jkpZHW+hsKMT+y",
	"hQ6m+ubHNM+/zpVnPvuXvcnIpVYallo9aERL/TZtzsMtNN3WxuyNLqiAg72o/Fzelh/L28n78ktUfogm",
	"78pvkzfl1/J28rb8FB0f+9HRa8xNJQ+2zmLc7vW3OpsZ7nb+7+2cdXa20+0OYpzGvbjX39pFEHeNo+pK",
	"qr5elVRJjdhiOiQbObKXMqWpuwm8uMLBgGxUpb2s00DAJVkX6v/biDfiqmdtSKGRkEDPHwm/mrxzfq+4",
	"rrdUO14VESyPFF15IvBoFqvoQTaNH4VQ5Q45fqqz67CFFJPyoGjMSKa+rHvhgu/3W3xz987PalFeFY0y",
	"ZFxZrP6yOKOVC3d0M44fVVObGj0M17uP4xE/GlvY8S101AQK0RhpKdz/Hzj6vE6JcDTynroVU5uUozr6",
	"IF9Xde7nhq+DS7/TlcM/1JObsHMKjzMipu7yX4eTdrJZyvTn/1Sse6c19ppXWkd/5p3+9TP4WRTfAwAA",
	"///gRNiu/gkAAA==",
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
