// Code generated by hertz generator. DO NOT EDIT.

package product

import (
	product "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/handler/product"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.POST("/product", append(_createproductMw(), product.CreateProduct)...)
	_product := root.Group("/product", _productMw()...)
	_product.DELETE("/:id", append(_deleteproductMw(), product.DeleteProduct)...)
	_product.PUT("/:id", append(_updateproductMw(), product.UpdateProduct)...)
	root.GET("/products", append(_getproductlistMw(), product.GetProductList)...)
	{
		_product0 := root.Group("/product", _product0Mw()...)
		_product0.GET("/:id", append(_getproductMw(), product.GetProduct)...)
	}
	{
		_products := root.Group("/products", _productsMw()...)
		_products.GET("/search", append(_searchproductsMw(), product.SearchProducts)...)
	}
}
