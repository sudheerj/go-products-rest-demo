package storage

import (
	log "github.com/sirupsen/logrus"
	"github.com/sudheerj/go-rest.git/model"
)


func (storeImpl *storeImpl) GetReviews(product *model.Product) (*[]model.Review, error) {
	log.WithFields(
		log.Fields{
			"ProductName": product.Name,
		},
	).Info("Get all reviews of a product")

	reviews := make([]model.Review, 0)
	if err := storeImpl.db.Model(&product).Association("Reviews").Error; err != nil {
		return nil, err
	}
	storeImpl.db.Model(&product).Association("Reviews").Find(&reviews)
	return &reviews, nil
}

func (storeImpl *storeImpl) GetReview(id int) (*model.Review, error) {
	log.WithFields(
		log.Fields{
			"ReviewID": id,
		},
	).Info("Get specific review")

	review := model.Review{}
	if err := storeImpl.db.First(&review, id).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (storeImpl *storeImpl) CreateReview(newReview *model.Review) error {
	log.Println("Create a new review")
	if err := storeImpl.db.Save(&newReview).Error; err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) UpdateReview(review *model.Review) error {
	log.WithFields(
		log.Fields{
			"ProductName": review.Title,
		},
	).Info("Update review of a product")

	if err := storeImpl.db.Save(&review).Error; err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) DeleteReview(review *model.Review) error {
	log.WithFields(
		log.Fields{
			"ProductName": review.Title,
		},
	).Info("Delete review of a product")

	if err := storeImpl.db.Delete(&review).Error; err != nil {
		return err
	}
	return nil
}

