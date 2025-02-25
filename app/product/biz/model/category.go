package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	CategoryTable = "category"
)

var categoryNotFoundErr = errors.New("category name is required")

type Category struct {
	Base
	Name     string    `gorm:"uniqueIndex;type:varchar(32);not null"`
	Products []Product `gorm:"many2many:product_category"`
}

func (Category) TableName() string {
	return CategoryTable
}

func (q *Dao) GetCategoryByName(categoryName string) (category *Category, err error) {
	if categoryName == "" {
		return nil, categoryNotFoundErr
	}
	err = q.db.WithContext(q.ctx).First(&category, &Category{Name: categoryName}).Error
	return category, err
}

func (q *Dao) GetOrCreateCategoryByName(categoryName string) (category *Category, err error) {
	if categoryName == "" {
		return nil, categoryNotFoundErr
	}

	err = q.db.WithContext(q.ctx).FirstOrCreate(&category, &Category{Name: categoryName}).Error
	return
}

func (q *Dao) DelUnusedCategory(category *Category) error {
	if category == nil {
		return categoryNotFoundErr
	}
	var count int64
	err := q.db.WithContext(q.ctx).Table(ProductCategoryTable).Where(CategoryIdWhereQuery, category.ID).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		err := q.db.WithContext(q.ctx).Delete(category).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (cq *CacheDao) GetCachedCategoryKey(categoryName string) string {
	return fmt.Sprintf("category:name:%s", categoryName)
}

func (cq *CacheDao) GetCachedCategoryByName(categoryName string) (category *Category, err error) {
	key := cq.GetCachedCategoryKey(categoryName)
	results := cq.cache.Get(cq.ctx, key)

	err = results.Err()
	if err != nil {
		return nil, err
	}
	resultsBytes, err := results.Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resultsBytes, &category)
	return
}

func (cq *CacheDao) SetCachedCategoryByName(categoryName string, categoryBytes []byte) (err error) {
	key := cq.GetCachedCategoryKey(categoryName)
	return cq.cache.Set(cq.ctx, key, categoryBytes, time.Hour).Err()
}

func (cq *CacheDao) DelCachedCategory(categoryName string) error {
	key := cq.GetCachedCategoryKey(categoryName)
	return cq.cache.Del(cq.ctx, key).Err()
}


func (cq *CacheDao) GetCategoryByName(categoryName string) (category *Category, err error) {
	if categoryName == "" {
		return nil, errors.New("product id can't be empty")
	}

	key := fmt.Sprintf("category:name:%s", categoryName)
	results := cq.cache.Get(cq.ctx, key)

	// find in cache
	err = func() error {
		err = results.Err()
		if err != nil {
			return err
		}
		resultsBytes, err := results.Bytes()
		if err != nil {
			return err
		}
		return json.Unmarshal(resultsBytes, &category)
	}()
	// miss cached, then find in db
	if err != nil {
		category, err = cq.Dao.GetCategoryByName(categoryName)
		if err != nil {
			return nil, err
		}
		prdBytes, err := json.Marshal(category)
		if err != nil {
			return category, nil
		}
		cq.cache.Set(cq.ctx, key, prdBytes, time.Hour)
	}

	return category, nil
}

func (cq *CacheDao) GetOrCreateCategoryByName(categoryName string) (category *Category, err error) {
	if categoryName == "" {
		return nil, categoryNotFoundErr
	}

	category, err = cq.GetCachedCategoryByName(categoryName)
	// miss cached, then find in db
	if err != nil {
		category, err = cq.Dao.GetOrCreateCategoryByName(categoryName)
		if err != nil {
			return nil, err
		}
		categoryBytes, err := json.Marshal(category)
		if err != nil {
			return category, nil
		}
		cq.SetCachedCategoryByName(categoryName, categoryBytes)
	}

	return category, nil
}

func (cq *CacheDao) DelUnusedCategory(category *Category) error {
	err := cq.Dao.DelUnusedCategory(category)
	if err != nil {
		return err
	}
	if category != nil {
		cq.DelCachedCategory(category.Name)
	}
	return nil
}