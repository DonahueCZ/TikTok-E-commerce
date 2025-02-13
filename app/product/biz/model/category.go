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

type Category struct {
	Base
	Name        string `gorm:"uniqueIndex;type:varchar(32);not null"`
	Products    []Product `gorm:"many2many:product_category"`
}

func (Category) TableName() string {
	return CategoryTable
}

func (q *Query) GetCategoryByName(categoryName string) (category *Category, err error) {
	if categoryName == "" {
		return nil, errors.New("category name is required")
	}
	err = q.db.WithContext(q.ctx).Where("name = ?", categoryName).First(&category).Error
	return category, err
}

func (cq *CachedQuery) GetCategoryByName(categoryName string) (category *Category, err error) {
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
		category, err = cq.Query.GetCategoryByName(categoryName)
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
