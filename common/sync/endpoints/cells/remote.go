/*
 * Copyright (c) 2019. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

package cells

import (
	"context"
	"fmt"
	"strings"

	servicecontext "github.com/pydio/cells/common/service/context"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	microgrpc "github.com/micro/go-plugins/client/grpc"
	"github.com/pborman/uuid"

	sdk "github.com/pydio/cells-sdk-go"
	"github.com/pydio/cells-sdk-go/transport"
	"github.com/pydio/cells/common"
	"github.com/pydio/cells/common/proto/tree"
	"github.com/pydio/cells/common/sync/model"
)

// Remote connect to a remove Cells server using the GRPC gateway.
type Remote struct {
	abstract
	config *sdk.SdkConfig
}

func NewRemote(config *sdk.SdkConfig, root string, options Options) *Remote {
	c := &Remote{
		abstract: abstract{
			root:       strings.TrimLeft(root, "/"),
			options:    options,
			clientUUID: uuid.New(),
		},
		config: config,
	}
	c.factory = &remoteClientFactory{
		config:   config,
		registry: NewDynamicRegistry(config),
	}
	c.source = c
	logCtx := context.Background()
	logCtx = servicecontext.WithServiceName(logCtx, "endpoint.cells.remote")
	logCtx = servicecontext.WithServiceColor(logCtx, servicecontext.ServiceColorGrpc)
	c.globalCtx = logCtx
	return c
}

func (c *Remote) GetEndpointInfo() model.EndpointInfo {
	return model.EndpointInfo{
		URI: fmt.Sprintf("%s/%s", c.config.Url, c.root),
		RequiresNormalization: false,
		RequiresFoldersRescan: false,
		SupportsTargetEcho:    true,
		Ignores:               []string{common.PYDIO_SYNC_HIDDEN_FILE_META},
	}

}

type remoteClientFactory struct {
	config   *sdk.SdkConfig
	registry *DynamicRegistry
}

func (f *remoteClientFactory) GetNodeProviderClient(ctx context.Context) (context.Context, tree.NodeProviderClient, error) {
	ctx, cli, e := f.getClient(ctx)
	if e != nil {
		return nil, nil, e
	}
	return ctx, tree.NewNodeProviderClient(transport.TargetServiceName, cli), nil
}

func (f *remoteClientFactory) GetNodeReceiverClient(ctx context.Context) (context.Context, tree.NodeReceiverClient, error) {
	ctx, cli, e := f.getClient(ctx)
	if e != nil {
		return nil, nil, e
	}
	return ctx, tree.NewNodeReceiverClient(transport.TargetServiceName, cli), nil
}

func (f *remoteClientFactory) GetNodeChangesStreamClient(ctx context.Context) (context.Context, tree.NodeChangesStreamerClient, error) {
	ctx, cli, e := f.getClient(ctx)
	if e != nil {
		return nil, nil, e
	}
	return ctx, tree.NewNodeChangesStreamerClient(transport.TargetServiceName, cli), nil
}

func (f *remoteClientFactory) GetObjectsClient(ctx context.Context) (context.Context, objectsClient, error) {
	return ctx, transport.NewS3Client(f.config), nil

}

func (f *remoteClientFactory) getClient(ctx context.Context) (context.Context, client.Client, error) {
	jwt, err := transport.RetrieveToken(f.config)
	if err != nil {
		return nil, nil, err
	}
	// create registry
	microClient := microgrpc.NewClient(
		client.Registry(f.registry.Micro),
		client.Wrap(func(i client.Client) client.Client {
			return &RegistryRefreshClient{w: i, r: f.registry}
		}),
	)
	var md metadata.Metadata
	if m, ok := metadata.FromContext(ctx); ok {
		md = m
	} else {
		md = metadata.Metadata{}
	}
	md["x-pydio-bearer"] = jwt
	ctx = metadata.NewContext(ctx, md)
	return ctx, microClient, nil

}
