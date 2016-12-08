/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.3.1.dev2
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

	"encoding/json"
	"fmt"
)

type CalendarApiService service


/**
 * List calendar event summaries
 * Get 50 event summaries from the calendar. If no event ID is given, the resource will return the next 50 chronological event summaries from now. If an event ID is specified, it will return the next 50 chronological event summaries from after that event.   ---  Alternate route: &#x60;/v1/characters/{character_id}/calendar/&#x60;  Alternate route: &#x60;/legacy/characters/{character_id}/calendar/&#x60;  Alternate route: &#x60;/dev/characters/{character_id}/calendar/&#x60;   ---  This route is cached for up to 5 seconds
 *
 * @param characterId The character to retrieve events from 
 * @param fromEvent(int32) The event ID to retrieve events from 
 * @param datasource(string) The server name you would like data from 
 * @return []GetCharactersCharacterIdCalendar200Ok
 */
func (a CalendarApiService) GetCharactersCharacterIdCalendar(ts TokenSource, characterId int64, fromEvent interface{}, datasource interface{}) ([]GetCharactersCharacterIdCalendar200Ok,  error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/characters/{character_id}/calendar/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(fromEvent, "int32", "fromEvent"); err != nil {
		return nil, err
	}
	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, err
	}
	if fromEvent != nil {
		localVarQueryParams.Add("from_event", a.client.parameterToString(fromEvent, ""))
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
	 var successPayload = new([]GetCharactersCharacterIdCalendar200Ok)

	 r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	 if err != nil {
		  return *successPayload, err
	 }

	if ts != nil {
		if t, err := ts.Token(); err != nil {
			return *successPayload, err
		} else if t != nil {
			t.SetAuthHeader(r)
		}
	}

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return *successPayload, err
	 }

	 defer localVarHttpResponse.Body.Close()
	 if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return nil, err
     }

	return *successPayload, err
}

/**
 * Get an event
 * Get all the information for a specific event  ---  Alternate route: &#x60;/v3/characters/{character_id}/calendar/{event_id}/&#x60;  Alternate route: &#x60;/dev/characters/{character_id}/calendar/{event_id}/&#x60;   ---  This route is cached for up to 5 seconds
 *
 * @param characterId The character id requesting the event 
 * @param eventId The id of the event requested 
 * @param datasource(string) The server name you would like data from 
 * @return *GetCharactersCharacterIdCalendarEventIdOk
 */
func (a CalendarApiService) GetCharactersCharacterIdCalendarEventId(ts TokenSource, characterId int64, eventId int32, datasource interface{}) (*GetCharactersCharacterIdCalendarEventIdOk,  error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/characters/{character_id}/calendar/{event_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"event_id"+"}", fmt.Sprintf("%v", eventId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, err
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
	 var successPayload = new(GetCharactersCharacterIdCalendarEventIdOk)

	 r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	 if err != nil {
		  return successPayload, err
	 }

	if ts != nil {
		if t, err := ts.Token(); err != nil {
			return successPayload, err
		} else if t != nil {
			t.SetAuthHeader(r)
		}
	}

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, err
	 }

	 defer localVarHttpResponse.Body.Close()
	 if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return nil, err
     }

	return successPayload, err
}

/**
 * Respond to an event
 * Set your response status to an event  ---  Alternate route: &#x60;/v3/characters/{character_id}/calendar/{event_id}/&#x60;  Alternate route: &#x60;/dev/characters/{character_id}/calendar/{event_id}/&#x60; 
 *
 * @param characterId The character ID requesting the event 
 * @param eventId The ID of the event requested 
 * @param response The response value to set, overriding current value. 
 * @param datasource(string) The server name you would like data from 
 * @return nil
 */
func (a CalendarApiService) PutCharactersCharacterIdCalendarEventId(ts TokenSource, characterId int32, eventId int32, response PutCharactersCharacterIdCalendarEventIdResponse, datasource interface{}) ( error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/characters/{character_id}/calendar/{event_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"event_id"+"}", fmt.Sprintf("%v", eventId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return err
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
	// body params
	 localVarPostBody = &response

	 r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	 if err != nil {
		  return err
	 }

	if ts != nil {
		if t, err := ts.Token(); err != nil {
			return err
		} else if t != nil {
			t.SetAuthHeader(r)
		}
	}

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return err
	 }
	return err
}

