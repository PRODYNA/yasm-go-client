/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.68.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ShoppingCartAPIService ShoppingCartAPI service
type ShoppingCartAPIService service

type ShoppingCartAPIAddPersonToShoppingCartRequest struct {
	ctx context.Context
	ApiService *ShoppingCartAPIService
	shoppingCartId string
	personId string
}

func (r ShoppingCartAPIAddPersonToShoppingCartRequest) Execute() (*ShoppingCartDetail, *http.Response, error) {
	return r.ApiService.AddPersonToShoppingCartExecute(r)
}

/*
AddPersonToShoppingCart Add a Person to a ShoppingCart

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shoppingCartId
 @param personId
 @return ShoppingCartAPIAddPersonToShoppingCartRequest
*/
func (a *ShoppingCartAPIService) AddPersonToShoppingCart(ctx context.Context, shoppingCartId string, personId string) ShoppingCartAPIAddPersonToShoppingCartRequest {
	return ShoppingCartAPIAddPersonToShoppingCartRequest{
		ApiService: a,
		ctx: ctx,
		shoppingCartId: shoppingCartId,
		personId: personId,
	}
}

// Execute executes the request
//  @return ShoppingCartDetail
func (a *ShoppingCartAPIService) AddPersonToShoppingCartExecute(r ShoppingCartAPIAddPersonToShoppingCartRequest) (*ShoppingCartDetail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ShoppingCartDetail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShoppingCartAPIService.AddPersonToShoppingCart")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shopping-carts/{shoppingCartId}/persons/{personId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shoppingCartId"+"}", url.PathEscape(parameterValueToString(r.shoppingCartId, "shoppingCartId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"personId"+"}", url.PathEscape(parameterValueToString(r.personId, "personId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ShoppingCartAPICreateShoppingCartRequest struct {
	ctx context.Context
	ApiService *ShoppingCartAPIService
	shoppingCart *ShoppingCart
}

func (r ShoppingCartAPICreateShoppingCartRequest) ShoppingCart(shoppingCart ShoppingCart) ShoppingCartAPICreateShoppingCartRequest {
	r.shoppingCart = &shoppingCart
	return r
}

func (r ShoppingCartAPICreateShoppingCartRequest) Execute() (*ShoppingCartDetail, *http.Response, error) {
	return r.ApiService.CreateShoppingCartExecute(r)
}

/*
CreateShoppingCart Create a ShoppingCart

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ShoppingCartAPICreateShoppingCartRequest
*/
func (a *ShoppingCartAPIService) CreateShoppingCart(ctx context.Context) ShoppingCartAPICreateShoppingCartRequest {
	return ShoppingCartAPICreateShoppingCartRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ShoppingCartDetail
func (a *ShoppingCartAPIService) CreateShoppingCartExecute(r ShoppingCartAPICreateShoppingCartRequest) (*ShoppingCartDetail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ShoppingCartDetail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShoppingCartAPIService.CreateShoppingCart")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shopping-carts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shoppingCart == nil {
		return localVarReturnValue, nil, reportError("shoppingCart is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.shoppingCart
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ShoppingCartAPIDeleteShoppingCartRequest struct {
	ctx context.Context
	ApiService *ShoppingCartAPIService
	shoppingCartId string
}

func (r ShoppingCartAPIDeleteShoppingCartRequest) Execute() (*Status, *http.Response, error) {
	return r.ApiService.DeleteShoppingCartExecute(r)
}

/*
DeleteShoppingCart Delete a ShoppingCart

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shoppingCartId
 @return ShoppingCartAPIDeleteShoppingCartRequest
*/
func (a *ShoppingCartAPIService) DeleteShoppingCart(ctx context.Context, shoppingCartId string) ShoppingCartAPIDeleteShoppingCartRequest {
	return ShoppingCartAPIDeleteShoppingCartRequest{
		ApiService: a,
		ctx: ctx,
		shoppingCartId: shoppingCartId,
	}
}

// Execute executes the request
//  @return Status
func (a *ShoppingCartAPIService) DeleteShoppingCartExecute(r ShoppingCartAPIDeleteShoppingCartRequest) (*Status, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Status
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShoppingCartAPIService.DeleteShoppingCart")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shopping-carts/{shoppingCartId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shoppingCartId"+"}", url.PathEscape(parameterValueToString(r.shoppingCartId, "shoppingCartId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ShoppingCartAPIGetShoppingCartRequest struct {
	ctx context.Context
	ApiService *ShoppingCartAPIService
	shoppingCartId string
}

func (r ShoppingCartAPIGetShoppingCartRequest) Execute() (*ShoppingCartDetail, *http.Response, error) {
	return r.ApiService.GetShoppingCartExecute(r)
}

/*
GetShoppingCart Get details about a ShoppingCart

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shoppingCartId
 @return ShoppingCartAPIGetShoppingCartRequest
*/
func (a *ShoppingCartAPIService) GetShoppingCart(ctx context.Context, shoppingCartId string) ShoppingCartAPIGetShoppingCartRequest {
	return ShoppingCartAPIGetShoppingCartRequest{
		ApiService: a,
		ctx: ctx,
		shoppingCartId: shoppingCartId,
	}
}

// Execute executes the request
//  @return ShoppingCartDetail
func (a *ShoppingCartAPIService) GetShoppingCartExecute(r ShoppingCartAPIGetShoppingCartRequest) (*ShoppingCartDetail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ShoppingCartDetail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShoppingCartAPIService.GetShoppingCart")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shopping-carts/{shoppingCartId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shoppingCartId"+"}", url.PathEscape(parameterValueToString(r.shoppingCartId, "shoppingCartId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ShoppingCartAPIRemovePersonFromShoppingCartRequest struct {
	ctx context.Context
	ApiService *ShoppingCartAPIService
	shoppingCartId string
	personId string
}

func (r ShoppingCartAPIRemovePersonFromShoppingCartRequest) Execute() (*ShoppingCartDetail, *http.Response, error) {
	return r.ApiService.RemovePersonFromShoppingCartExecute(r)
}

/*
RemovePersonFromShoppingCart Remove a Person from a ShoppingCart

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shoppingCartId
 @param personId
 @return ShoppingCartAPIRemovePersonFromShoppingCartRequest
*/
func (a *ShoppingCartAPIService) RemovePersonFromShoppingCart(ctx context.Context, shoppingCartId string, personId string) ShoppingCartAPIRemovePersonFromShoppingCartRequest {
	return ShoppingCartAPIRemovePersonFromShoppingCartRequest{
		ApiService: a,
		ctx: ctx,
		shoppingCartId: shoppingCartId,
		personId: personId,
	}
}

// Execute executes the request
//  @return ShoppingCartDetail
func (a *ShoppingCartAPIService) RemovePersonFromShoppingCartExecute(r ShoppingCartAPIRemovePersonFromShoppingCartRequest) (*ShoppingCartDetail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ShoppingCartDetail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShoppingCartAPIService.RemovePersonFromShoppingCart")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shopping-carts/{shoppingCartId}/persons/{personId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shoppingCartId"+"}", url.PathEscape(parameterValueToString(r.shoppingCartId, "shoppingCartId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"personId"+"}", url.PathEscape(parameterValueToString(r.personId, "personId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ShoppingCartAPISearchShoppingCartsRequest struct {
	ctx context.Context
	ApiService *ShoppingCartAPIService
	shoppingCartSearch *ShoppingCartSearch
}

func (r ShoppingCartAPISearchShoppingCartsRequest) ShoppingCartSearch(shoppingCartSearch ShoppingCartSearch) ShoppingCartAPISearchShoppingCartsRequest {
	r.shoppingCartSearch = &shoppingCartSearch
	return r
}

func (r ShoppingCartAPISearchShoppingCartsRequest) Execute() (*PagedShoppingCarts, *http.Response, error) {
	return r.ApiService.SearchShoppingCartsExecute(r)
}

/*
SearchShoppingCarts Complex search over shopping cart entities

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ShoppingCartAPISearchShoppingCartsRequest
*/
func (a *ShoppingCartAPIService) SearchShoppingCarts(ctx context.Context) ShoppingCartAPISearchShoppingCartsRequest {
	return ShoppingCartAPISearchShoppingCartsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PagedShoppingCarts
func (a *ShoppingCartAPIService) SearchShoppingCartsExecute(r ShoppingCartAPISearchShoppingCartsRequest) (*PagedShoppingCarts, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PagedShoppingCarts
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShoppingCartAPIService.SearchShoppingCarts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shopping-carts/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.shoppingCartSearch
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ShoppingCartAPIShareShoppingCartRequest struct {
	ctx context.Context
	ApiService *ShoppingCartAPIService
	shoppingCartId string
	personId string
}

func (r ShoppingCartAPIShareShoppingCartRequest) Execute() (*ShoppingCartDetail, *http.Response, error) {
	return r.ApiService.ShareShoppingCartExecute(r)
}

/*
ShareShoppingCart Share a ShoppingCart with a Person

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shoppingCartId
 @param personId
 @return ShoppingCartAPIShareShoppingCartRequest
*/
func (a *ShoppingCartAPIService) ShareShoppingCart(ctx context.Context, shoppingCartId string, personId string) ShoppingCartAPIShareShoppingCartRequest {
	return ShoppingCartAPIShareShoppingCartRequest{
		ApiService: a,
		ctx: ctx,
		shoppingCartId: shoppingCartId,
		personId: personId,
	}
}

// Execute executes the request
//  @return ShoppingCartDetail
func (a *ShoppingCartAPIService) ShareShoppingCartExecute(r ShoppingCartAPIShareShoppingCartRequest) (*ShoppingCartDetail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ShoppingCartDetail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShoppingCartAPIService.ShareShoppingCart")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shopping-carts/{shoppingCartId}/shared-with/{personId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shoppingCartId"+"}", url.PathEscape(parameterValueToString(r.shoppingCartId, "shoppingCartId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"personId"+"}", url.PathEscape(parameterValueToString(r.personId, "personId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ShoppingCartAPIUnshareShoppingCartRequest struct {
	ctx context.Context
	ApiService *ShoppingCartAPIService
	shoppingCartId string
	personId string
}

func (r ShoppingCartAPIUnshareShoppingCartRequest) Execute() (*ShoppingCartDetail, *http.Response, error) {
	return r.ApiService.UnshareShoppingCartExecute(r)
}

/*
UnshareShoppingCart Unshare a ShoppingCart with a Person

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shoppingCartId
 @param personId
 @return ShoppingCartAPIUnshareShoppingCartRequest
*/
func (a *ShoppingCartAPIService) UnshareShoppingCart(ctx context.Context, shoppingCartId string, personId string) ShoppingCartAPIUnshareShoppingCartRequest {
	return ShoppingCartAPIUnshareShoppingCartRequest{
		ApiService: a,
		ctx: ctx,
		shoppingCartId: shoppingCartId,
		personId: personId,
	}
}

// Execute executes the request
//  @return ShoppingCartDetail
func (a *ShoppingCartAPIService) UnshareShoppingCartExecute(r ShoppingCartAPIUnshareShoppingCartRequest) (*ShoppingCartDetail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ShoppingCartDetail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShoppingCartAPIService.UnshareShoppingCart")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shopping-carts/{shoppingCartId}/shared-with/{personId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shoppingCartId"+"}", url.PathEscape(parameterValueToString(r.shoppingCartId, "shoppingCartId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"personId"+"}", url.PathEscape(parameterValueToString(r.personId, "personId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ShoppingCartAPIUpdateShoppingCartRequest struct {
	ctx context.Context
	ApiService *ShoppingCartAPIService
	shoppingCartId string
	shoppingCart *ShoppingCart
}

func (r ShoppingCartAPIUpdateShoppingCartRequest) ShoppingCart(shoppingCart ShoppingCart) ShoppingCartAPIUpdateShoppingCartRequest {
	r.shoppingCart = &shoppingCart
	return r
}

func (r ShoppingCartAPIUpdateShoppingCartRequest) Execute() (*ShoppingCartDetail, *http.Response, error) {
	return r.ApiService.UpdateShoppingCartExecute(r)
}

/*
UpdateShoppingCart Update a ShoppingCart

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shoppingCartId
 @return ShoppingCartAPIUpdateShoppingCartRequest
*/
func (a *ShoppingCartAPIService) UpdateShoppingCart(ctx context.Context, shoppingCartId string) ShoppingCartAPIUpdateShoppingCartRequest {
	return ShoppingCartAPIUpdateShoppingCartRequest{
		ApiService: a,
		ctx: ctx,
		shoppingCartId: shoppingCartId,
	}
}

// Execute executes the request
//  @return ShoppingCartDetail
func (a *ShoppingCartAPIService) UpdateShoppingCartExecute(r ShoppingCartAPIUpdateShoppingCartRequest) (*ShoppingCartDetail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ShoppingCartDetail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShoppingCartAPIService.UpdateShoppingCart")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shopping-carts/{shoppingCartId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shoppingCartId"+"}", url.PathEscape(parameterValueToString(r.shoppingCartId, "shoppingCartId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shoppingCart == nil {
		return localVarReturnValue, nil, reportError("shoppingCart is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.shoppingCart
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
