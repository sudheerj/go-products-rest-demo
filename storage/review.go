package storage
import (
	"errors"
	"github.com/sudheerj/go-rest.git/model"
)

func (storeImpl *storeImpl) GetReviews(productId, limit, offset int) (*[]model.Review, error) {
	reviews := []model.Review{}
	product, _ := DBStore.GetProduct(productId)
	if err := storeImpl.db.Model(&product).Association("Reviews").Error; err != nil {
		return nil, err
	}
	return &reviews, nil
}

func (storeImpl *storeImpl) GetReview(id int) (*model.Review, error) {
	review := model.Review{}
	if err := storeImpl.db.First(&review, model.Review{ProductID: id}).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (storeImpl *storeImpl) CreateReview(review *model.Review) error {
	if review == nil {
		return errors.New("no review found with given product Id")
	}
	if err := storeImpl.db.Save(&review).Error; err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) UpdateReview(review *model.Review) error {
	if review == nil {
		return errors.New("no review found with given product Id")
	}
	if err := storeImpl.db.Save(&review).Error; err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) DeleteReview(id int) error {
	review, err := storeImpl.GetProduct(id)
	if err != nil {
		return errors.New("no review found with given product Id")
	}
	if err := storeImpl.db.Delete(&review).Error; err != nil {
		return err
	}
	return nil
}

