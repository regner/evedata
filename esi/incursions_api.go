/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.3.2.dev3
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

import (
	"net/url"
	"strings"
	"time"
	"errors"
	"golang.org/x/net/context"
	"encoding/json"
)

var _ context.Context

type IncursionsApiService service


/**
 * List incursions
 * Return a list of current incursions  ---  Alternate route: &#x60;/v1/incursions/&#x60;  Alternate route: &#x60;/legacy/incursions/&#x60;  Alternate route: &#x60;/dev/incursions/&#x60;   ---  This route is cached for up to 300 seconds
 *
 * @param datasource(string) The server name you would like data from 
 * @return []GetIncursions200Ok
 */
func (a IncursionsApiService) GetIncursions(datasource interface{}) ([]GetIncursions200Ok,  time.Time, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GetIncursions200Ok
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/incursions/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return successPayload, time.Now(), err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}


	 r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	 if err != nil {
		  return successPayload, time.Now(), err
	 }

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, time.Now(), err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, time.Now(), errors.New(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, time.Now(), err
	}

	expires := cacheExpires(localVarHttpResponse)
	return successPayload, expires, err
}

