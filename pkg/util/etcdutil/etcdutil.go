// Copyright 2016 The etcd-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package etcdutil

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/coreos/etcd-operator/pkg/util/constants"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func initClient(ctx context.Context, clientURLs []string, tc *tls.Config) (*clientv3.Client, error) {
	cfg := clientv3.Config{
		Endpoints:   clientURLs,
		DialTimeout: constants.DefaultDialTimeout,
		TLS:         tc,
		Context:     ctx,
	}
	return clientv3.New(cfg)
}

func AddMember(clientURLs []string, tc *tls.Config, peerURLs []string) (*clientv3.MemberAddResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultRequestTimeout)
	defer cancel()
	etcdcli, err := initClient(ctx, clientURLs, tc)
	defer etcdcli.Close()
	if err != nil {
		return nil, fmt.Errorf("add one member failed: creating etcd client failed %v", err)
	}
	resp, err := etcdcli.MemberAdd(ctx, peerURLs)
	etcdcli.Close()
	return resp, err
}

func ListMembers(clientURLs []string, tc *tls.Config) (*clientv3.MemberListResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultRequestTimeout)
	defer cancel()
	etcdcli, err := initClient(ctx, clientURLs, tc)
	if err != nil {
		return nil, fmt.Errorf("list members failed: creating etcd client failed: %v", err)
	}
	defer etcdcli.Close()
	resp, err := etcdcli.MemberList(ctx)
	etcdcli.Close()
	return resp, err
}

func RemoveMember(clientURLs []string, tc *tls.Config, id uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultRequestTimeout)
	defer cancel()
	etcdcli, err := initClient(ctx, clientURLs, tc)
	if err != nil {
		return fmt.Errorf("remove members failed: creating etcd client failed: %v", err)
	}
	defer etcdcli.Close()
	_, err = etcdcli.Cluster.MemberRemove(ctx, id)
	return err
}

func ClientWithMaxRev(ctx context.Context, endpoints []string, tc *tls.Config) (*clientv3.Client, int64, error) {
	mapEps := make(map[string]*clientv3.Client)
	var maxClient *clientv3.Client
	maxRev := int64(0)
	errors := make([]string, 0)
	for _, endpoint := range endpoints {
		etcdcli, err := initClient(ctx, []string{endpoint}, tc)
		if err != nil {
			errors = append(errors, fmt.Sprintf("failed to create etcd client for endpoint (%v): %v", endpoint, err))
			continue
		}
		mapEps[endpoint] = etcdcli

		resp, err := etcdcli.Get(ctx, "/", clientv3.WithSerializable())
		if err != nil {
			errors = append(errors, fmt.Sprintf("failed to get revision from endpoint (%s)", endpoint))
			continue
		}

		logrus.Infof("getMaxRev: endpoint %s revision (%d)", endpoint, resp.Header.Revision)
		if resp.Header.Revision > maxRev {
			maxRev = resp.Header.Revision
			maxClient = etcdcli
		}
	}

	// close all open clients that are not maxClient.
	for _, cli := range mapEps {
		if cli == maxClient {
			continue
		}
		cli.Close()
	}

	if maxClient == nil {
		return nil, 0, fmt.Errorf("could not create an etcd client for the max revision purpose from given endpoints (%v)", endpoints)
	}

	var err error
	if len(errors) > 0 {
		errorStr := ""
		for _, errStr := range errors {
			errorStr += errStr + "\n"
		}
		err = fmt.Errorf(errorStr)
	}

	return maxClient, maxRev, err
}
