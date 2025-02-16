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
	CategoryPreloadTag   string
)

type Product struct {
	Base
	StoreId     uint32 `gorm:"index:store_id;not null"`
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

type Dao struct {
	ctx context.Context
	db  *gorm.DB
}

func NewDao(ctx context.Context, db *gorm.DB) *Dao {
	return &Dao{
		ctx: ctx,
		db:  db,
	}
}

func (q *Dao) GetProductById(id uint32) (product *Product, err error) {
	if id == 0 {
		return nil, errors.New("product id can't be empty")
	}
	err = q.db.WithContext(q.ctx).Model(&Product{}).Preload(CategoryPreloadTag).First(&product, id).Error
	return
}

func (q *Dao) GetProductsByQuery(query string) (products []*Product, err error) {
	err = q.db.WithContext(q.ctx).Where("MATCH(name, description) AGAINST(?)", query).Find(&products).Error
	return
}

func (q *Dao) GetProductPage(page, pagesize uint32) (products []Product, err error) {
	err = q.db.WithContext(q.ctx).Model(&Product{}).Order(Id).Limit(int(pagesize)).Offset(int((page - 1) * pagesize)).Find(&products).Error
	return
}

func (q *Dao) GetProductPageByCategory(page, pagesize uint32, category *Category) (products []Product, err error) {
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

func (q *Dao) CreateProduct(product *Product) uint32 {
	if product == nil || product.StoreId == 0 || product.Name == "" || product.Description == "" {
		return 0
	}
	if q.db.WithContext(q.ctx).Create(product).Error != nil {
		return 0
	}
	return product.ID
}

func (q *Dao) DeleteProduct(product *Product) error {
	if product == nil {
		return errors.New("product is required")
	}
	return q.db.WithContext(q.ctx).Delete(product).Error
}

func (q *Dao) UpdateProduct(product *Product) error {
	if product == nil || product.StoreId == 0 || product.Name == "" || product.Description == "" {
		return errors.New("product is required")
	}
	err := q.db.WithContext(q.ctx).Model(&product).Association(CategoryPreloadTag).Replace(product.Categories)
	if err != nil {
		return err
	}
	return q.db.WithContext(q.ctx).Save(product).Error
}

type CacheDao struct {
	Dao
	cache *redis.Client
}

func NewCacheDao(ctx context.Context, db *gorm.DB, cache *redis.Client) *CacheDao {
	return &CacheDao{
		Dao:   *NewDao(ctx, db),
		cache: cache,
	}
}

func (cq *CacheDao) GetCachedProductKey(id uint32) string {
	return fmt.Sprintf("product:id:%d", id)
}

func (cq *CacheDao) GetCachedProduct(id uint32) (product *Product, err error) {
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

func (cq *CacheDao) SetCachedProduct(id uint32, prdBytes []byte) (err error) {
	key := cq.GetCachedProductKey(id)
	return cq.cache.Set(cq.ctx, key, prdBytes, time.Hour).Err()
}

func (cq *CacheDao) DelCachedProduct(id uint32) error {
	key := cq.GetCachedProductKey(id)
	return cq.cache.Del(cq.ctx, key).Err()
}

func (cq *CacheDao) GetProductById(id uint32) (product *Product, err error) {
	product, err = cq.GetCachedProduct(id)
	// miss cached, then find in db
	if err != nil {
		product, err = cq.Dao.GetProductById(id)
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

func (cq *CacheDao) GetProductsByQuery(query string) (products []*Product, err error) {
	return cq.Dao.GetProductsByQuery(query)
}

func (cq *CacheDao) GetCachedProductPage(page, pagesize uint32, categoryName string) (products []Product, err error) {
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

func (cq *CacheDao) SetCachedProductPage(page, pagesize uint32, categoryName string, prdBytes []byte) (err error) {
	key := fmt.Sprintf("product:category:%s:page:%d:pagesize:%d", categoryName, page, pagesize)
	err = cq.cache.Set(cq.ctx, key, prdBytes, time.Hour).Err()
	return
}

func (cq *CacheDao) GetProductPageByCategory(page, pagesize uint32, category *Category) (products []Product, err error) {
	if page <= 0 || pagesize <= 0 {
		return nil, errors.New("page and pagesize must be positive")
	}
	if category != nil {
		products, err = cq.GetCachedProductPage(page, pagesize, category.Name)
	}
	if category == nil || err != nil {
		products, err = cq.Dao.GetProductPageByCategory(page, pagesize, category)
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

func (cq *CacheDao) CreateProduct(product *Product) (id uint32) {
	id = cq.Dao.CreateProduct(product)
	if id == 0 {
		return id
	}
	prdBytes, err := json.Marshal(product)
	if err == nil {
		cq.SetCachedProduct(id, prdBytes)
	}
	return id
}

func (cq *CacheDao) DeleteProduct(product *Product) error {
	err := cq.Dao.DeleteProduct(product)
	if err != nil {
		return err
	}
	if product != nil {
		cq.DelCachedProduct(product.ID)
	}
	return err
}

func (cq *CacheDao) UpdateProduct(product *Product) error {
	err := cq.Dao.UpdateProduct(product)
	if err != nil {
		return err
	}
	if product != nil {
		prdBytes, err := json.Marshal(*product)
		if err != nil {
			return nil
		}
		cq.SetCachedProduct(product.ID, prdBytes)
	}
	return err
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
				CategoryPreloadTag = prdType.Field(i).Name
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
