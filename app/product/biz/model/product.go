package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const (
	ProductTable = "product"
	Id           = "id"
)

var (
	ProductIdColumn      string
	CategoryIdColumn     string
	ProductCategoryTable string
	ProductJoinQuery     string
	CategoryIdWhereQuery string
)

type Product struct {
	Base
	Name        string `gorm:"index:query,class:FULLTEXT;index:name;type:varchar(64);not null"`
	Description string `gorm:"index:query,class:FULLTEXT;type:TEXT"`
	Picture     string
	Price       uint32     `gorm:"not null;default:0"` // use integer to avoid accuracy loss; represent float with a decimal precision of 2
	Stock       uint32     `gorm:"not null;default:0"`
	Categories  []Category `gorm:"many2many:product_category"`
}

func (Product) TableName() string {
	return ProductTable
}

type Query struct {
	ctx context.Context
	db  *gorm.DB
}

func NewQuery(ctx context.Context, db *gorm.DB) *Query {
	return &Query{
		ctx: ctx,
		db:  db,
	}
}

func (q *Query) GetProductById(id uint32) (product *Product, err error) {
	if id == 0 {
		return nil, errors.New("product id can't be empty")
	}
	err = q.db.WithContext(q.ctx).First(&product, id).Error
	return
}

func (q *Query) GetProductsByQuery(query string) (products []*Product, err error) {
	err = q.db.WithContext(q.ctx).Where("MATCH(name, description) AGAINST(?)", query).Find(&products).Error
	return
}

func (q *Query) GetProductPage(page, pagesize uint32) (products []Product, err error) {
	err = q.db.WithContext(q.ctx).Model(&Product{}).Order(Id).Limit(int(pagesize)).Offset(int((page - 1) * pagesize)).Find(&products).Error
	return
}

func (q *Query) GetProductPageByCategory(page, pagesize uint32, category *Category) (products []Product, err error) {
	if category == nil || category.Name == "" {
		products, err = q.GetProductPage(page, pagesize)
		return
	}
	err = q.db.WithContext(q.ctx).Joins(ProductJoinQuery).
		Where(CategoryIdWhereQuery, category.ID).
		Limit(int(pagesize)).
		Offset(int((page - 1) * pagesize)).
		Find(&products).Error
	return
}

func (q *Query) CreateProduct(product *Product) uint32 {
	if product == nil {
		return 0
	}
	if q.db.WithContext(q.ctx).Create(product).Error != nil {
		return 0
	}
	return product.ID
}

func (q *Query) DeleteProduct(product *Product) bool {
	if product == nil {
		return false
	}
	return q.db.WithContext(q.ctx).Delete(product).Error == nil
}

func (q *Query) UpdateProduct(product *Product) bool {
	if product == nil {
		return false
	}
	return q.db.WithContext(q.ctx).Save(product).Error == nil
}

type CachedQuery struct {
	Query
	cache *redis.Client
}

func NewCachedQuery(ctx context.Context, db *gorm.DB, cache *redis.Client) *CachedQuery {
	return &CachedQuery{
		Query: *NewQuery(ctx, db),
		cache: cache,
	}
}

func (cq *CachedQuery) GetCachedProductKey(id uint32) string {
	return fmt.Sprintf("product:id:%d", id)
}

func (cq *CachedQuery) GetCachedProduct(id uint32) (product *Product, err error) {
	key := cq.GetCachedProductKey(id)
	results := cq.cache.Get(cq.ctx, key)
	err = results.Err()
	if err != nil {
		return nil, err
	}
	resultsBytes, err := results.Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resultsBytes, &product)
	return
}

func (cq *CachedQuery) SetCachedProduct(id uint32, prdBytes []byte) (err error) {
	key := cq.GetCachedProductKey(id)
	return cq.cache.Set(cq.ctx, key, prdBytes, time.Hour).Err()
}

func (cq *CachedQuery) DelCachedProduct(id uint32) error {
	key := cq.GetCachedProductKey(id)
	return cq.cache.Del(cq.ctx, key).Err()
}

func (cq *CachedQuery) GetProductById(id uint32) (product *Product, err error) {
	product, err = cq.GetCachedProduct(id)
	// miss cached, then find in db
	if err != nil {
		product, err = cq.Query.GetProductById(id)
		if err != nil {
			return nil, err
		}
		prdBytes, err := json.Marshal(product)
		if err != nil {
			return product, nil
		}
		cq.SetCachedProduct(id, prdBytes)
	}

	return product, nil
}

func (cq *CachedQuery) GetProductsByQuery(query string) (products []*Product, err error) {
	return cq.Query.GetProductsByQuery(query)
}

func (cq *CachedQuery) GetCachedProductPage(page, pagesize uint32, categoryName string) (products []Product, err error) {
	key := fmt.Sprintf("product:category:%s:page:%d:pagesize:%d", categoryName, page, pagesize)

	results := cq.cache.Get(cq.ctx, key)
	err = results.Err()
	if err != nil {
		return nil, err
	}
	resultsBytes, err := results.Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resultsBytes, &products)
	return
}

func (cq *CachedQuery) SetCachedProductPage(page, pagesize uint32, categoryName string, prdBytes []byte) (err error) {
	key := fmt.Sprintf("product:category:%s:page:%d:pagesize:%d", categoryName, page, pagesize)
	err = cq.cache.Set(cq.ctx, key, prdBytes, time.Hour).Err()
	return
}

func (cq *CachedQuery) GetProductPageByCategory(page, pagesize uint32, category *Category) (products []Product, err error) {
	if page <= 0 || pagesize <= 0 {
		return nil, errors.New("page and pagesize must be positive")
	}
	if category != nil {
		products, err = cq.GetCachedProductPage(page, pagesize, category.Name)
	}
	if category == nil || err != nil {
		products, err = cq.Query.GetProductPageByCategory(page, pagesize, category)
		if err != nil {
			return nil, err
		}
		prdBytes, err := json.Marshal(products)
		if err != nil {
			return products, nil
		}
		categoryName := ""
		if category != nil {
			categoryName = category.Name
		}
		cq.SetCachedProductPage(page, pagesize, categoryName, prdBytes)
	}
	return products, nil
}

func (cq *CachedQuery) CreateProduct(product *Product) (id uint32) {
	id = cq.Query.CreateProduct(product)
	if id == 0 {
		return id
	}
	prdBytes, err := json.Marshal(product)
	if err == nil {
		cq.SetCachedProduct(id, prdBytes)
	}
	return id
}

func (cq *CachedQuery) DeleteProduct(product *Product) bool {
	flag := cq.Query.DeleteProduct(product)
	if !flag {
		return flag
	}
	if product != nil {
		cq.DelCachedProduct(product.ID)
	}
	return flag
}

func (cq *CachedQuery) UpdateProduct(product *Product) bool {
	flag := cq.Query.UpdateProduct(product)
	if !flag {
		return flag
	}
	if product != nil {
		cq.DelCachedProduct(product.ID)
	}
	return flag
}

func init() {
	ProductIdColumn = fmt.Sprintf("%s_id", ProductTable)
	CategoryIdColumn = fmt.Sprintf("%s_id", CategoryTable)
	// get product_category table name
	prdType := reflect.TypeOf(Product{})
	for i := range prdType.NumField() {
		gormTag := prdType.Field(i).Tag.Get("gorm")
		for _, field := range strings.Split(gormTag, ";") {
			field = strings.TrimSpace(field)
			if s, ok := strings.CutPrefix(field, "many2many:"); ok {
				ProductCategoryTable = strings.TrimSpace(s)
				break
			}
		}
	}
	if ProductCategoryTable == "" {
		panic("failed to get product_category join table")
	}

	// generate query
	ProductJoinQuery = fmt.Sprintf("JOIN %s ON %s.%s = %s.%s",
		ProductCategoryTable,
		ProductCategoryTable, ProductIdColumn,
		ProductTable, Id,
	)
	CategoryIdWhereQuery = fmt.Sprintf("%s.%s = ?", ProductCategoryTable, CategoryIdColumn)
}
