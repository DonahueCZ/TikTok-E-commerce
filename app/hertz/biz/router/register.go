// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	product "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/router/product"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	product.Register(r)
}
