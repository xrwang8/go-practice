package impl

import (
	"context"

	"github.com/go-practice/restful-api-demo/apps/host"
)

func (h *HostService) CreateHost(ctx context.Context, host *host.Host) (*host.Host, error) {
	h.l.Debug("记得记得就记得叫")
	return nil, nil
}

func (h *HostService) QueryHost(ctx context.Context, request *host.QueryHostRequest) (*host.HostSet, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HostService) DescribeHost(ctx context.Context, request *host.DescribeHostRequest) (*host.Host, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HostService) UpdateHost(ctx context.Context, request *host.UpdateHostRequest) (*host.Host, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HostService) DeleteHost(ctx context.Context, request *host.DeleteHostRequest) (*host.Host, error) {
	//TODO implement me
	panic("implement me")
}
