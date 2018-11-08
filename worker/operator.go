package worker

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/coreos/etcd/clientv3"
	"github.com/cybozu-go/log"
	"github.com/cybozu-go/neco"
	"github.com/cybozu-go/neco/storage"
)

// Operator installs or updates programs
type Operator interface {
	// UpdateNeco updates neco package.
	UpdateNeco(ctx context.Context, req *neco.UpdateRequest) error

	// FinalStep is the step number of the final operation.
	FinalStep() int

	// RunStep executes operations for given step.
	RunStep(ctx context.Context, req *neco.UpdateRequest, step int) error
}

type operator struct {
	mylrn       int
	ec          *clientv3.Client
	storage     storage.Storage
	proxyClient *http.Client
	localClient *http.Client
}

// NewOperator creates an Operator
func NewOperator(ctx context.Context, ec *clientv3.Client, mylrn int) (Operator, error) {
	st := storage.NewStorage(ec)
	localClient := localHTTPClient()
	proxyClient := localClient

	proxy, err := st.GetProxyConfig(ctx)
	if err != nil {
		if err != storage.ErrNotFound {
			return nil, err
		}
	} else {
		if len(proxy) > 0 {
			proxyURL, err := url.Parse(proxy)
			if err != nil {
				return nil, err
			}
			proxyClient = proxyHTTPClient(proxyURL)
		}
	}

	return &operator{
		mylrn:       mylrn,
		ec:          ec,
		storage:     st,
		proxyClient: proxyClient,
		localClient: localClient,
	}, nil
}

func (o *operator) UpdateNeco(ctx context.Context, req *neco.UpdateRequest) error {
	deb := &neco.DebianPackage{
		Name:       "neco",
		Repository: neco.GitHubRepoName,
		Owner:      neco.GitHubRepoOwner,
		Release:    req.Version,
	}

	log.Info("update neco", map[string]interface{}{
		"version": req.Version,
	})
	return InstallDebianPackage(ctx, o.proxyClient, deb)
}

func (o *operator) FinalStep() int {
	return 2
}

func (o *operator) RunStep(ctx context.Context, req *neco.UpdateRequest, step int) error {
	switch step {
	case 1:
		return o.UpdateEtcd(ctx, req)
	case 2:
		return o.UpdateVault(ctx, req)
	}

	return fmt.Errorf("invalid step: %d", step)
}