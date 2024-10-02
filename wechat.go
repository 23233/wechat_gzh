package wechat

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/23233/wechat_gzh/v2/cache"
	"github.com/23233/wechat_gzh/v2/officialaccount"
	offConfig "github.com/23233/wechat_gzh/v2/officialaccount/config"
	"github.com/23233/wechat_gzh/v2/pay"
	payConfig "github.com/23233/wechat_gzh/v2/pay/config"
	"github.com/23233/wechat_gzh/v2/util"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

// Wechat struct
type Wechat struct {
	cache cache.Cache
}

// NewWechat init
func NewWechat() *Wechat {
	return &Wechat{}
}

// SetCache 设置 cache
func (wc *Wechat) SetCache(cache cache.Cache) {
	wc.cache = cache
}

// GetOfficialAccount 获取微信公众号实例
func (wc *Wechat) GetOfficialAccount(cfg *offConfig.Config) *officialaccount.OfficialAccount {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return officialaccount.NewOfficialAccount(cfg)
}

// GetPay 获取微信支付的实例
func (wc *Wechat) GetPay(cfg *payConfig.Config) *pay.Pay {
	return pay.NewPay(cfg)
}

// SetHTTPClient  设置HTTPClient
func (wc *Wechat) SetHTTPClient(client *http.Client) {
	util.DefaultHTTPClient = client
}
