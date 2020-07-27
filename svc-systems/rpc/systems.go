//(C) Copyright [2020] Hewlett Packard Enterprise Development LP
//
//Licensed under the Apache License, Version 2.0 (the "License"); you may
//not use this file except in compliance with the License. You may obtain
//a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//License for the specific language governing permissions and limitations
// under the License.

//Package rpc ...
package rpc

import (
	"context"
	"log"
	"net/http"

	"github.com/bharath-b-hpe/odimra/lib-utilities/common"
	systemsproto "github.com/bharath-b-hpe/odimra/lib-utilities/proto/systems"
	"github.com/bharath-b-hpe/odimra/svc-plugin-rest-client/pmbhandle"
	"github.com/bharath-b-hpe/odimra/svc-systems/systems"
)

// Systems struct helps to register service
type Systems struct {
	IsAuthorizedRPC func(sessionToken string, privileges, oemPrivileges []string) (int32, string)
}

//GetSystemResource defines the operations which handles the RPC request response
// for the getting the system resource  of systems micro service.
// The functionality retrives the request and return backs the response to
// RPC according to the protoc file defined in the util-lib package.
// The function uses IsAuthorized of util-lib to validate the session
// which is present in the request.
func (s *Systems) GetSystemResource(ctx context.Context, req *systemsproto.GetSystemsRequest, resp *systemsproto.SystemsResponse) error {
	sessionToken := req.SessionToken
	authStatusCode, authStatusMessage := s.IsAuthorizedRPC(sessionToken, []string{common.PrivilegeLogin}, []string{})
	if authStatusCode != http.StatusOK {
		errorMessage := "error while trying to authenticate session"
		resp.StatusCode = authStatusCode
		resp.StatusMessage = authStatusMessage
		rpcResp := common.GeneralError(authStatusCode, authStatusMessage, errorMessage, nil, nil)
		resp.Body = generateResponse(rpcResp.Body)
		resp.Header = rpcResp.Header
		log.Printf(errorMessage)
		return nil
	}
	var pc = systems.PluginContact{
		ContactClient:  pmbhandle.ContactPlugin,
		DevicePassword: common.DecryptWithPrivateKey,
	}
	data := pc.GetSystemResource(req)
	resp.Header = data.Header
	resp.StatusCode = data.StatusCode
	resp.StatusMessage = data.StatusMessage
	resp.Body = generateResponse(data.Body)
	return nil
}

// GetSystemsCollection defines the operation which has the RPC request
// for getting the systems data from odimra.
// Retrieves all the keys with table name systems collection and create the response
// to send back to requested user.
func (s *Systems) GetSystemsCollection(ctx context.Context, req *systemsproto.GetSystemsRequest, resp *systemsproto.SystemsResponse) error {
	sessionToken := req.SessionToken
	authStatusCode, authStatusMessage := s.IsAuthorizedRPC(sessionToken, []string{common.PrivilegeLogin}, []string{})
	if authStatusCode != http.StatusOK {
		errorMessage := "error while trying to authenticate session"
		resp.StatusCode = authStatusCode
		resp.StatusMessage = authStatusMessage
		rpcResp := common.GeneralError(authStatusCode, authStatusMessage, errorMessage, nil, nil)
		resp.Body = generateResponse(rpcResp.Body)
		resp.Header = rpcResp.Header
		log.Printf(errorMessage)
		return nil
	}
	data := systems.GetSystemsCollection(req)
	resp.Header = data.Header
	resp.StatusCode = data.StatusCode
	resp.StatusMessage = data.StatusMessage
	resp.Body = generateResponse(data.Body)
	return nil
}

//GetSystems defines the operations which handles the RPC request response
// for the getting the system resource  of systems micro service.
// The functionality retrives the request and return backs the response to
// RPC according to the protoc file defined in the util-lib package.
// The function uses IsAuthorized of util-lib to validate the session
// which is present in the request.
func (s *Systems) GetSystems(ctx context.Context, req *systemsproto.GetSystemsRequest, resp *systemsproto.SystemsResponse) error {
	sessionToken := req.SessionToken
	authStatusCode, authStatusMessage := s.IsAuthorizedRPC(sessionToken, []string{common.PrivilegeLogin}, []string{})
	if authStatusCode != http.StatusOK {
		errorMessage := "error while trying to authenticate session"
		resp.StatusCode = authStatusCode
		resp.StatusMessage = authStatusMessage
		rpcResp := common.GeneralError(authStatusCode, authStatusMessage, errorMessage, nil, nil)
		resp.Body = generateResponse(rpcResp.Body)
		resp.Header = rpcResp.Header
		log.Printf(errorMessage)
		return nil
	}
	var pc = systems.PluginContact{
		ContactClient:  pmbhandle.ContactPlugin,
		DevicePassword: common.DecryptWithPrivateKey,
	}
	data := pc.GetSystems(req)
	resp.Header = data.Header
	resp.StatusCode = data.StatusCode
	resp.StatusMessage = data.StatusMessage
	resp.Body = generateResponse(data.Body)
	return nil
}

// ComputerSystemReset defines the operations which handles the RPC request response
// for the ComputerSystemReset service of systems micro service.
// The functionality retrives the request and return backs the response to
// RPC according to the protoc file defined in the lib-utilities package.
// The function also checks for the session time out of the token
// which is present in the request.
func (s *Systems) ComputerSystemReset(ctx context.Context, req *systemsproto.ComputerSystemResetRequest, resp *systemsproto.SystemsResponse) error {
	sessionToken := req.SessionToken
	authStatusCode, authStatusMessage := s.IsAuthorizedRPC(sessionToken, []string{common.PrivilegeConfigureComponents}, []string{})
	if authStatusCode != http.StatusOK {
		errorMessage := "error while trying to authenticate session"
		resp.StatusCode = authStatusCode
		resp.StatusMessage = authStatusMessage
		rpcResp := common.GeneralError(authStatusCode, authStatusMessage, errorMessage, nil, nil)
		resp.Body = generateResponse(rpcResp.Body)
		resp.Header = rpcResp.Header
		log.Printf(errorMessage)
		return nil
	}
	var pc = systems.PluginContact{
		ContactClient:  pmbhandle.ContactPlugin,
		DevicePassword: common.DecryptWithPrivateKey,
	}
	data := pc.ComputerSystemReset(req.SystemID, req.ResetType)
	resp.StatusCode = data.StatusCode
	resp.StatusMessage = data.StatusMessage
	resp.Header = data.Header
	resp.Body = generateResponse(data.Body)
	return nil
}

// SetDefaultBootOrder defines the operations which handles the RPC request response
// for the SetDefaultBootOrder service of systems micro service.
// The functionality retrives the request and return backs the response to
// RPC according to the protoc file defined in the lib-utilities package.
// The function also checks for the session time out of the token
// which is present in the request.
func (s *Systems) SetDefaultBootOrder(ctx context.Context, req *systemsproto.DefaultBootOrderRequest, resp *systemsproto.SystemsResponse) error {
	sessionToken := req.SessionToken
	authStatusCode, authStatusMessage := s.IsAuthorizedRPC(sessionToken, []string{common.PrivilegeConfigureComponents}, []string{})
	if authStatusCode != http.StatusOK {
		errorMessage := "error while trying to authenticate session"
		resp.StatusCode = authStatusCode
		resp.StatusMessage = authStatusMessage
		rpcResp := common.GeneralError(authStatusCode, authStatusMessage, errorMessage, nil, nil)
		resp.Body = generateResponse(rpcResp.Body)
		resp.Header = rpcResp.Header
		log.Printf(errorMessage)
		return nil
	}
	var pc = systems.PluginContact{
		ContactClient:  pmbhandle.ContactPlugin,
		DevicePassword: common.DecryptWithPrivateKey,
	}
	data := pc.SetDefaultBootOrder(req.SystemID)
	resp.StatusCode = data.StatusCode
	resp.StatusMessage = data.StatusMessage
	resp.Header = data.Header
	resp.Body = generateResponse(data.Body)
	return nil
}

// ChangeBiosSettings defines the operations which handles the RPC request response
// for the ChangeBiosSettings service of systems micro service.
// The functionality retrives the request and return backs the response to
// RPC according to the protoc file defined in the lib-utilities package.
// The function also checks for the session time out of the token
// which is present in the request.
func (s *Systems) ChangeBiosSettings(ctx context.Context, req *systemsproto.BiosSettingsRequest, resp *systemsproto.SystemsResponse) error {
	sessionToken := req.SessionToken
	authStatusCode, authStatusMessage := s.IsAuthorizedRPC(sessionToken, []string{common.PrivilegeConfigureComponents}, []string{})
	if authStatusCode != http.StatusOK {
		errorMessage := "error while trying to authenticate session"
		resp.StatusCode = authStatusCode
		resp.StatusMessage = authStatusMessage
		rpcResp := common.GeneralError(authStatusCode, authStatusMessage, errorMessage, nil, nil)
		resp.Body = generateResponse(rpcResp.Body)
		resp.Header = rpcResp.Header
		log.Printf(errorMessage)
		return nil
	}
	var pc = systems.PluginContact{
		ContactClient:  pmbhandle.ContactPlugin,
		DevicePassword: common.DecryptWithPrivateKey,
	}
	data := pc.ChangeBiosSettings(req)
	resp.StatusCode = data.StatusCode
	resp.StatusMessage = data.StatusMessage
	resp.Header = data.Header
	resp.Body = generateResponse(data.Body)
	return nil
}

// ChangeBootOrderSettings defines the operations which handles the RPC request response
// for the ChangeBootOrderSettings service of systems micro service.
// The functionality retrives the request and return backs the response to
// RPC according to the protoc file defined in the lib-utilities package.
// The function also checks for the session time out of the token
// which is present in the request.
func (s *Systems) ChangeBootOrderSettings(ctx context.Context, req *systemsproto.BootOrderSettingsRequest, resp *systemsproto.SystemsResponse) error {
	sessionToken := req.SessionToken
	authStatusCode, authStatusMessage := s.IsAuthorizedRPC(sessionToken, []string{common.PrivilegeConfigureComponents}, []string{})
	if authStatusCode != http.StatusOK {
		errorMessage := "error while trying to authenticate session"
		resp.StatusCode = authStatusCode
		resp.StatusMessage = authStatusMessage
		rpcResp := common.GeneralError(authStatusCode, authStatusMessage, errorMessage, nil, nil)
		resp.Body = generateResponse(rpcResp.Body)
		resp.Header = rpcResp.Header
		log.Printf(errorMessage)
		return nil
	}
	var pc = systems.PluginContact{
		ContactClient:  pmbhandle.ContactPlugin,
		DevicePassword: common.DecryptWithPrivateKey,
	}
	data := pc.ChangeBootOrderSettings(req)
	resp.StatusCode = data.StatusCode
	resp.StatusMessage = data.StatusMessage
	resp.Header = data.Header
	resp.Body = generateResponse(data.Body)
	return nil
}