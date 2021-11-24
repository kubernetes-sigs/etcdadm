package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeStorageSets invokes the ecs.DescribeStorageSets API synchronously
func (client *Client) DescribeStorageSets(request *DescribeStorageSetsRequest) (response *DescribeStorageSetsResponse, err error) {
	response = CreateDescribeStorageSetsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStorageSetsWithChan invokes the ecs.DescribeStorageSets API asynchronously
func (client *Client) DescribeStorageSetsWithChan(request *DescribeStorageSetsRequest) (<-chan *DescribeStorageSetsResponse, <-chan error) {
	responseChan := make(chan *DescribeStorageSetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStorageSets(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeStorageSetsWithCallback invokes the ecs.DescribeStorageSets API asynchronously
func (client *Client) DescribeStorageSetsWithCallback(request *DescribeStorageSetsRequest, callback func(response *DescribeStorageSetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStorageSetsResponse
		var err error
		defer close(result)
		response, err = client.DescribeStorageSets(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeStorageSetsRequest is the request struct for api DescribeStorageSets
type DescribeStorageSetsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	StorageSetIds        string           `position:"Query" name:"StorageSetIds"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	StorageSetName       string           `position:"Query" name:"StorageSetName"`
}

// DescribeStorageSetsResponse is the response struct for api DescribeStorageSets
type DescribeStorageSetsResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	TotalCount  int         `json:"TotalCount" xml:"TotalCount"`
	PageNumber  int         `json:"PageNumber" xml:"PageNumber"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	StorageSets StorageSets `json:"StorageSets" xml:"StorageSets"`
}

// CreateDescribeStorageSetsRequest creates a request to invoke DescribeStorageSets API
func CreateDescribeStorageSetsRequest() (request *DescribeStorageSetsRequest) {
	request = &DescribeStorageSetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeStorageSets", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeStorageSetsResponse creates a response to parse from DescribeStorageSets response
func CreateDescribeStorageSetsResponse() (response *DescribeStorageSetsResponse) {
	response = &DescribeStorageSetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
