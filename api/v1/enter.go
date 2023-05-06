package v1

import (
	"blogger/api/v1/service"
	"blogger/api/v1/system"
)

type Api struct {
	System  system.SysApi
	Service service.SvcApi
}
