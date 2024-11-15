/*
Cloud Controller API

API specification for managing environment variable groups in Cloud Controller.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package capiclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


type AuditEventsAPI interface {

	/*
	V3AuditEventsGet List audit events

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV3AuditEventsGetRequest
	*/
	V3AuditEventsGet(ctx context.Context) ApiV3AuditEventsGetRequest

	// V3AuditEventsGetExecute executes the request
	//  @return AuditEventList
	V3AuditEventsGetExecute(r ApiV3AuditEventsGetRequest) (*AuditEventList, *http.Response, error)

	/*
	V3AuditEventsGuidGet Retrieve an audit event

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param guid Unique identifier for the event
	@return ApiV3AuditEventsGuidGetRequest
	*/
	V3AuditEventsGuidGet(ctx context.Context, guid string) ApiV3AuditEventsGuidGetRequest

	// V3AuditEventsGuidGetExecute executes the request
	//  @return AuditEvent
	V3AuditEventsGuidGetExecute(r ApiV3AuditEventsGuidGetRequest) (*AuditEvent, *http.Response, error)
}

// AuditEventsAPIService AuditEventsAPI service
type AuditEventsAPIService service

type ApiV3AuditEventsGetRequest struct {
	ctx context.Context
	ApiService AuditEventsAPI
	types *[]string
	targetGuids *[]string
	spaceGuids *[]string
	organizationGuids *[]string
	page *int32
	perPage *int32
	orderBy *string
	createdAts *[]time.Time
	updatedAts *[]time.Time
}

// Comma-delimited list of event types to filter by
func (r ApiV3AuditEventsGetRequest) Types(types []string) ApiV3AuditEventsGetRequest {
	r.types = &types
	return r
}

// Comma-delimited list of target guids to filter by. Also supports filtering by exclusion.
func (r ApiV3AuditEventsGetRequest) TargetGuids(targetGuids []string) ApiV3AuditEventsGetRequest {
	r.targetGuids = &targetGuids
	return r
}

// Comma-delimited list of space guids to filter by
func (r ApiV3AuditEventsGetRequest) SpaceGuids(spaceGuids []string) ApiV3AuditEventsGetRequest {
	r.spaceGuids = &spaceGuids
	return r
}

// Comma-delimited list of organization guids to filter by
func (r ApiV3AuditEventsGetRequest) OrganizationGuids(organizationGuids []string) ApiV3AuditEventsGetRequest {
	r.organizationGuids = &organizationGuids
	return r
}

// Page to display
func (r ApiV3AuditEventsGetRequest) Page(page int32) ApiV3AuditEventsGetRequest {
	r.page = &page
	return r
}

// Number of results per page
func (r ApiV3AuditEventsGetRequest) PerPage(perPage int32) ApiV3AuditEventsGetRequest {
	r.perPage = &perPage
	return r
}

// Value to sort by
func (r ApiV3AuditEventsGetRequest) OrderBy(orderBy string) ApiV3AuditEventsGetRequest {
	r.orderBy = &orderBy
	return r
}

// Timestamp to filter by
func (r ApiV3AuditEventsGetRequest) CreatedAts(createdAts []time.Time) ApiV3AuditEventsGetRequest {
	r.createdAts = &createdAts
	return r
}

// Timestamp to filter by
func (r ApiV3AuditEventsGetRequest) UpdatedAts(updatedAts []time.Time) ApiV3AuditEventsGetRequest {
	r.updatedAts = &updatedAts
	return r
}

func (r ApiV3AuditEventsGetRequest) Execute() (*AuditEventList, *http.Response, error) {
	return r.ApiService.V3AuditEventsGetExecute(r)
}

/*
V3AuditEventsGet List audit events

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV3AuditEventsGetRequest
*/
func (a *AuditEventsAPIService) V3AuditEventsGet(ctx context.Context) ApiV3AuditEventsGetRequest {
	return ApiV3AuditEventsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuditEventList
func (a *AuditEventsAPIService) V3AuditEventsGetExecute(r ApiV3AuditEventsGetRequest) (*AuditEventList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditEventList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditEventsAPIService.V3AuditEventsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/audit_events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.types != nil {
		t := *r.types
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "types", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "types", t, "form", "multi")
		}
	}
	if r.targetGuids != nil {
		t := *r.targetGuids
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "target_guids", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "target_guids", t, "form", "multi")
		}
	}
	if r.spaceGuids != nil {
		t := *r.spaceGuids
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "space_guids", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "space_guids", t, "form", "multi")
		}
	}
	if r.organizationGuids != nil {
		t := *r.organizationGuids
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "organization_guids", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "organization_guids", t, "form", "multi")
		}
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "form", "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "form", "")
	}
	if r.createdAts != nil {
		t := *r.createdAts
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "created_ats", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "created_ats", t, "form", "multi")
		}
	}
	if r.updatedAts != nil {
		t := *r.updatedAts
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "updated_ats", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "updated_ats", t, "form", "multi")
		}
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

type ApiV3AuditEventsGuidGetRequest struct {
	ctx context.Context
	ApiService AuditEventsAPI
	guid string
}

func (r ApiV3AuditEventsGuidGetRequest) Execute() (*AuditEvent, *http.Response, error) {
	return r.ApiService.V3AuditEventsGuidGetExecute(r)
}

/*
V3AuditEventsGuidGet Retrieve an audit event

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param guid Unique identifier for the event
 @return ApiV3AuditEventsGuidGetRequest
*/
func (a *AuditEventsAPIService) V3AuditEventsGuidGet(ctx context.Context, guid string) ApiV3AuditEventsGuidGetRequest {
	return ApiV3AuditEventsGuidGetRequest{
		ApiService: a,
		ctx: ctx,
		guid: guid,
	}
}

// Execute executes the request
//  @return AuditEvent
func (a *AuditEventsAPIService) V3AuditEventsGuidGetExecute(r ApiV3AuditEventsGuidGetRequest) (*AuditEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditEventsAPIService.V3AuditEventsGuidGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/audit_events/{guid}"
	localVarPath = strings.Replace(localVarPath, "{"+"guid"+"}", url.PathEscape(parameterValueToString(r.guid, "guid")), -1)

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
