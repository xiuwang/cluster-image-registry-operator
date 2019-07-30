package powerbiembeddedapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/powerbiembedded/mgmt/2016-01-29/powerbiembedded"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	GetAvailableOperations(ctx context.Context) (result powerbiembedded.OperationList, err error)
}

var _ BaseClientAPI = (*powerbiembedded.BaseClient)(nil)

// WorkspaceCollectionsClientAPI contains the set of methods on the WorkspaceCollectionsClient type.
type WorkspaceCollectionsClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, body powerbiembedded.CheckNameRequest) (result powerbiembedded.CheckNameResponse, err error)
	Create(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body powerbiembedded.CreateWorkspaceCollectionRequest) (result powerbiembedded.WorkspaceCollection, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceCollectionName string) (result powerbiembedded.WorkspaceCollectionsDeleteFuture, err error)
	GetAccessKeys(ctx context.Context, resourceGroupName string, workspaceCollectionName string) (result powerbiembedded.WorkspaceCollectionAccessKeys, err error)
	GetByName(ctx context.Context, resourceGroupName string, workspaceCollectionName string) (result powerbiembedded.WorkspaceCollection, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result powerbiembedded.WorkspaceCollectionList, err error)
	ListBySubscription(ctx context.Context) (result powerbiembedded.WorkspaceCollectionList, err error)
	Migrate(ctx context.Context, resourceGroupName string, body powerbiembedded.MigrateWorkspaceCollectionRequest) (result autorest.Response, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body powerbiembedded.WorkspaceCollectionAccessKey) (result powerbiembedded.WorkspaceCollectionAccessKeys, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body powerbiembedded.UpdateWorkspaceCollectionRequest) (result powerbiembedded.WorkspaceCollection, err error)
}

var _ WorkspaceCollectionsClientAPI = (*powerbiembedded.WorkspaceCollectionsClient)(nil)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceCollectionName string) (result powerbiembedded.WorkspaceList, err error)
}

var _ WorkspacesClientAPI = (*powerbiembedded.WorkspacesClient)(nil)