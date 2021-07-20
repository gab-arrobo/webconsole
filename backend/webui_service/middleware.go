// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package webui_service

import (
	"github.com/gin-gonic/gin"

	"github.com/free5gc/path_util"
)

var PublicPath string

func init() {
	PublicPath = path_util.Free5gcPath("free5gc/webconsole/public")
}

func ReturnPublic() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		if method == "GET" {
			destPath := PublicPath + context.Request.RequestURI
			if destPath[len(destPath)-1] == '/' {
				destPath = destPath[:len(destPath)-1]
			}
			context.File(destPath)
		} else {
			context.Next()
		}
	}
}
