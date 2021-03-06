/*
 * TheGamesDB API
 *
 * API Documentation
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gamesdb

import (
	_context "context"
	"github.com/antihax/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// GamesApiService GamesApi service
type GamesApiService service

// GamesByGameIDOpts Optional parameters for the method 'GamesByGameID'
type GamesByGameIDOpts struct {
	Fields  optional.String
	Include optional.String
	Page    optional.Int32
}

/*
GamesByGameID Fetch game(s) by id
can request additional information can be requestes through &#x60;fields&#x60; and &#x60;include&#x60; params
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param apikey (Required)
 * @param id (Required) - supports `,` delimited list
 * @param optional nil or *GamesByGameIDOpts - Optional Parameters:
 * @param "Fields" (optional.String) -  (Optional) - valid `,` delimited options: `players`, `publishers`, `genres`, `overview`, `last_updated`, `rating`, `platform`, `coop`, `youtube`, `os`, `processor`, `ram`, `hdd`, `video`, `sound`, `alternates`
 * @param "Include" (optional.String) -  (Optional) - valid `,` delimited options: `boxart`, `platform`
 * @param "Page" (optional.Int32) -  (Optional) - results page offset to return
@return GamesByGameId
*/
func (a *GamesApiService) GamesByGameID(ctx _context.Context, apikey string, id string, localVarOptionals *GamesByGameIDOpts) (GamesByGameId, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GamesByGameId
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/Games/ByGameID"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("apikey", parameterToString(apikey, ""))
	localVarQueryParams.Add("id", parameterToString(id, ""))
	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Include.IsSet() {
		localVarQueryParams.Add("include", parameterToString(localVarOptionals.Include.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v GamesByGameId
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GamesByGameNameOpts Optional parameters for the method 'GamesByGameName'
type GamesByGameNameOpts struct {
	Fields         optional.String
	FilterPlatform optional.String
	Include        optional.String
	Page           optional.Int32
}

/*
GamesByGameName Fetch game(s) by name
can request additional information can be requestes through &#x60;fields&#x60; and &#x60;include&#x60; params
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param apikey (Required)
 * @param name (Required) - Search term
 * @param optional nil or *GamesByGameNameOpts - Optional Parameters:
 * @param "Fields" (optional.String) -  (Optional) - valid `,` delimited options: `players`, `publishers`, `genres`, `overview`, `last_updated`, `rating`, `platform`, `coop`, `youtube`, `os`, `processor`, `ram`, `hdd`, `video`, `sound`, `alternates`
 * @param "FilterPlatform" (optional.String) -  (Optional) - platform `id` can be obtain from the platforms api below, supports `,` delimited list
 * @param "Include" (optional.String) -  (Optional) - valid `,` delimited options: `boxart`, `platform`
 * @param "Page" (optional.Int32) -  (Optional) - results page offset to return
@return GamesByGameId
*/
func (a *GamesApiService) GamesByGameName(ctx _context.Context, apikey string, name string, localVarOptionals *GamesByGameNameOpts) (GamesByGameId, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GamesByGameId
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1.1/Games/ByGameName"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("apikey", parameterToString(apikey, ""))
	localVarQueryParams.Add("name", parameterToString(name, ""))
	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterPlatform.IsSet() {
		localVarQueryParams.Add("filter[platform]", parameterToString(localVarOptionals.FilterPlatform.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Include.IsSet() {
		localVarQueryParams.Add("include", parameterToString(localVarOptionals.Include.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v GamesByGameId
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GamesByGameNameV1Opts Optional parameters for the method 'GamesByGameNameV1'
type GamesByGameNameV1Opts struct {
	Fields         optional.String
	FilterPlatform optional.String
	Include        optional.String
	Page           optional.Int32
}

/*
GamesByGameNameV1 Fetch game(s) by name
can request additional information can be requestes through &#x60;fields&#x60; and &#x60;include&#x60; params
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param apikey (Required)
 * @param name (Required) - Search term
 * @param optional nil or *GamesByGameNameV1Opts - Optional Parameters:
 * @param "Fields" (optional.String) -  (Optional) - valid `,` delimited options: `players`, `publishers`, `genres`, `overview`, `last_updated`, `rating`, `platform`, `coop`, `youtube`, `os`, `processor`, `ram`, `hdd`, `video`, `sound`, `alternates`
 * @param "FilterPlatform" (optional.String) -  (Optional) - platform `id` can be obtain from the platforms api below, supports `,` delimited list
 * @param "Include" (optional.String) -  (Optional) - valid `,` delimited options: `boxart`, `platform`
 * @param "Page" (optional.Int32) -  (Optional) - results page offset to return
@return GamesByGameIdV1
*/
func (a *GamesApiService) GamesByGameNameV1(ctx _context.Context, apikey string, name string, localVarOptionals *GamesByGameNameV1Opts) (GamesByGameIdV1, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GamesByGameIdV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/Games/ByGameName"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("apikey", parameterToString(apikey, ""))
	localVarQueryParams.Add("name", parameterToString(name, ""))
	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterPlatform.IsSet() {
		localVarQueryParams.Add("filter[platform]", parameterToString(localVarOptionals.FilterPlatform.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Include.IsSet() {
		localVarQueryParams.Add("include", parameterToString(localVarOptionals.Include.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v GamesByGameIdV1
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GamesByPlatformIDOpts Optional parameters for the method 'GamesByPlatformID'
type GamesByPlatformIDOpts struct {
	Fields  optional.String
	Include optional.String
	Page    optional.Int32
}

/*
GamesByPlatformID Fetch game(s) by platform id
can request additional information can be requested through &#x60;fields&#x60; and &#x60;include&#x60; params
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param apikey (Required)
 * @param id (Required) - platform `id` can be obtain from the platforms api below, supports `,` delimited list
 * @param optional nil or *GamesByPlatformIDOpts - Optional Parameters:
 * @param "Fields" (optional.String) -  (Optional) - valid `,` delimited options: `players`, `publishers`, `genres`, `overview`, `last_updated`, `rating`, `platform`, `coop`, `youtube`, `os`, `processor`, `ram`, `hdd`, `video`, `sound`, `alternates`
 * @param "Include" (optional.String) -  (Optional) - valid `,` delimited options: `boxart`, `platform`
 * @param "Page" (optional.Int32) -  (Optional) - results page offset to return
@return GamesByGameId
*/
func (a *GamesApiService) GamesByPlatformID(ctx _context.Context, apikey string, id string, localVarOptionals *GamesByPlatformIDOpts) (GamesByGameId, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GamesByGameId
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/Games/ByPlatformID"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("apikey", parameterToString(apikey, ""))
	localVarQueryParams.Add("id", parameterToString(id, ""))
	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Include.IsSet() {
		localVarQueryParams.Add("include", parameterToString(localVarOptionals.Include.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v GamesByGameId
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GamesImagesOpts Optional parameters for the method 'GamesImages'
type GamesImagesOpts struct {
	FilterType optional.String
	Page       optional.Int32
}

/*
GamesImages Fetch game(s) images by game(s) id
results can be filtered with &#x60;filter[type]&#x60; param
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param apikey (Required)
 * @param gamesId (Required) - game(s) `id` can be obtain from the above games api, supports `,` delimited list
 * @param optional nil or *GamesImagesOpts - Optional Parameters:
 * @param "FilterType" (optional.String) -  (Optional) - valid `,` delimited options: `fanart`, `banner`, `boxart`, `screenshot`, `clearlogo`
 * @param "Page" (optional.Int32) -  (Optional) - results page offset to return
@return GamesImages
*/
func (a *GamesApiService) GamesImages(ctx _context.Context, apikey string, gamesId string, localVarOptionals *GamesImagesOpts) (GamesImages, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GamesImages
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/Games/Images"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("apikey", parameterToString(apikey, ""))
	localVarQueryParams.Add("games_id", parameterToString(gamesId, ""))
	if localVarOptionals != nil && localVarOptionals.FilterType.IsSet() {
		localVarQueryParams.Add("filter[type]", parameterToString(localVarOptionals.FilterType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v GamesImages
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GamesUpdatesOpts Optional parameters for the method 'GamesUpdates'
type GamesUpdatesOpts struct {
	Time optional.Int32
	Page optional.Int32
}

/*
GamesUpdates Fetch games update
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param apikey (Required)
 * @param lastEditId (Required)
 * @param optional nil or *GamesUpdatesOpts - Optional Parameters:
 * @param "Time" (optional.Int32) -  (Optional)
 * @param "Page" (optional.Int32) -  (Optional) - results page offset to return
@return GamesUpdates
*/
func (a *GamesApiService) GamesUpdates(ctx _context.Context, apikey string, lastEditId int32, localVarOptionals *GamesUpdatesOpts) (GamesUpdates, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GamesUpdates
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/Games/Updates"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("apikey", parameterToString(apikey, ""))
	localVarQueryParams.Add("last_edit_id", parameterToString(lastEditId, ""))
	if localVarOptionals != nil && localVarOptionals.Time.IsSet() {
		localVarQueryParams.Add("time", parameterToString(localVarOptionals.Time.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
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
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v GamesUpdates
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
