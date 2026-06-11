package controller

import (
	"strings"

	"github.com/wack7788/new-api/common"
	"github.com/wack7788/new-api/setting/system_setting"
)

func paymentReturnPath(suffix string) string {
	base := strings.TrimRight(system_setting.ServerAddress, "/")
	return base + common.ThemeAwarePath(suffix)
}
