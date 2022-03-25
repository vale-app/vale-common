package repositories

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/vale-app/vale-common/models"
	"github.com/vale-app/vale-common/utils"
)

type TaxonomiesService struct {
	client *resty.Client
}

func NewTaxonomiesService(client *resty.Client) *TaxonomiesService {
	return &TaxonomiesService{client: client}
}

func (s *TaxonomiesService) GetCategoryByID(categoryID uint) (*models.Category, error) {
	res, err := s.client.R().Get(fmt.Sprintf("/taxonomies/categories/%d", categoryID))
	if err != nil {
		return nil, err
	}

	var c models.Category
	err = utils.JsonToStruct(string(res.Body()), &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (s *TaxonomiesService) GetSubCategoryByID(subCategoryID uint) (*models.SubCategory, error) {
	res, err := s.client.R().Get(fmt.Sprintf("/taxonomies/subcategories/%d", subCategoryID))
	if err != nil {
		return nil, err
	}

	var c models.SubCategory
	err = utils.JsonToStruct(string(res.Body()), c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
