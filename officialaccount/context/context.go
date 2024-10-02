package context

import (
	"github.com/23233/wechat_gzh/v2/credential"
	"github.com/23233/wechat_gzh/v2/officialaccount/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
