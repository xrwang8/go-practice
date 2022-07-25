package host

import (
	"context"
)

type Service interface {
	CreateHost(context.Context, *Host) (*Host, error)

	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)

	DescribeHost(context.Context, *DescribeHostRequest) (*Host, error)

	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)

	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
}
