/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"io"
	"net/http"

	"github.com/openshift-online/ocm-sdk-go/helpers"
)

func readRegistryCredentialsAddRequest(request *RegistryCredentialsAddServerRequest, r *http.Request) error {
	var err error
	request.body, err = UnmarshalRegistryCredential(r.Body)
	return err
}
func writeRegistryCredentialsAddRequest(request *RegistryCredentialsAddRequest, writer io.Writer) error {
	return MarshalRegistryCredential(request.body, writer)
}
func readRegistryCredentialsAddResponse(response *RegistryCredentialsAddResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalRegistryCredential(reader)
	return err
}
func writeRegistryCredentialsAddResponse(response *RegistryCredentialsAddServerResponse, w http.ResponseWriter) error {
	return MarshalRegistryCredential(response.body, w)
}
func readRegistryCredentialsListRequest(request *RegistryCredentialsListServerRequest, r *http.Request) error {
	var err error
	query := r.URL.Query()
	request.order, err = helpers.ParseString(query, "order")
	if err != nil {
		return err
	}
	request.page, err = helpers.ParseInteger(query, "page")
	if err != nil {
		return err
	}
	if request.page == nil {
		request.page = helpers.NewInteger(1)
	}
	request.search, err = helpers.ParseString(query, "search")
	if err != nil {
		return err
	}
	request.size, err = helpers.ParseInteger(query, "size")
	if err != nil {
		return err
	}
	if request.size == nil {
		request.size = helpers.NewInteger(100)
	}
	return nil
}
func writeRegistryCredentialsListRequest(request *RegistryCredentialsListRequest, writer io.Writer) error {
	return nil
}
func readRegistryCredentialsListResponse(response *RegistryCredentialsListResponse, reader io.Reader) error {
	iterator, err := helpers.NewIterator(reader)
	if err != nil {
		return err
	}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "page":
			value := iterator.ReadInt()
			response.page = &value
		case "size":
			value := iterator.ReadInt()
			response.size = &value
		case "total":
			value := iterator.ReadInt()
			response.total = &value
		case "items":
			items := readRegistryCredentialList(iterator)
			response.items = &RegistryCredentialList{
				items: items,
			}
		default:
			iterator.ReadAny()
		}
	}
	return iterator.Error
}
func writeRegistryCredentialsListResponse(response *RegistryCredentialsListServerResponse, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.status)
	stream := helpers.NewStream(w)
	stream.WriteObjectStart()
	stream.WriteObjectField("kind")
	count := 1
	stream.WriteString(RegistryCredentialListKind)
	if response.items != nil && response.items.href != nil {
		stream.WriteMore()
		stream.WriteObjectField("href")
		stream.WriteString(*response.items.href)
		count++
	}
	if response.page != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("page")
		stream.WriteInt(*response.page)
		count++
	}
	if response.size != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("size")
		stream.WriteInt(*response.size)
		count++
	}
	if response.total != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("total")
		stream.WriteInt(*response.total)
		count++
	}
	if response.items != nil {
		if response.items.items != nil {
			if count > 0 {
				stream.WriteMore()
			}
			stream.WriteObjectField("items")
			writeRegistryCredentialList(response.items.items, stream)
			count++
		}
	}
	stream.WriteObjectEnd()
	stream.Flush()
	return stream.Error
}