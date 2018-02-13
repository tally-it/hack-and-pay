package sql

import (
	"context"

	"github.com/marove2000/hack-and-pay/contract"
	"github.com/marove2000/hack-and-pay/errors"
	"github.com/marove2000/hack-and-pay/repository/sql/models"

)

func (m *Mysql) AddCategory(ctx context.Context, r contract.AddCategoryBody) (int, error) {
	logger := pkgLogger.ForFunc(ctx, "AddProduct")
	logger.Debug("enter repo")

	category := models.Category{
		Name:         r.Name,
	}

	if r.IsVisible == true {
		category.IsVisible = boolToString(true)
	} else {
		category.IsVisible = boolToString(false)
	}


	if r.IsActive == true {
		category.IsActive = boolToString(true)
	} else {
		category.IsActive = boolToString(false)
	}

	if r.IsRoot == true {
		category.IsRoot = boolToString(true)
	} else {
		category.IsRoot = boolToString(false)
	}


	err := category.Insert(m.db)
	if err != nil {
		logger.WithError(err).Error("failed to insert category")
		return 0, errors.InternalServerError("db error", err)
	}

	return category.CategoryID, nil
}