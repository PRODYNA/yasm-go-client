# \TemplatesAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/graph/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTemplates**](TemplatesAPI.md#GetTemplates) | **Get** /templates | list of existing templates



## GetTemplates

> []string GetTemplates(ctx).Execute()

list of existing templates

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesAPI.GetTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplates`: []string
    fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

