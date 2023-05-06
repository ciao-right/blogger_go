package routerGroup

import (
	serviceGroup "blogger/router/service"
	systemGroup "blogger/router/system"
)

type RouterGroup struct {
	System  systemGroup.SystemRouter
	Service serviceGroup.ServiceRouter
}
