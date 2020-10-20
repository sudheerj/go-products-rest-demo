package storage
import "github.com/sudheerj/go-rest.git/model"


func (storeImpl *storeImpl) GetReviews(product *model.Product) (*[]model.Review, error) {
	reviews := make([]model.Review, 0)
	if err := storeImpl.db.Model(&product).Association("Reviews").Error; err != nil {
		return nil, err
	}
	storeImpl.db.Model(&product).Association("Reviews").Find(&reviews)
	return &reviews, nil
}

func (storeImpl *storeImpl) GetReview(id int) (*model.Review, error) {
	review := model.Review{}
	if err := storeImpl.db.First(&review, id).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (storeImpl *storeImpl) CreateReview(newReview *model.Review) error {
	if err := storeImpl.db.Save(&newReview).Error; err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) UpdateReview(review *model.Review) error {
	if err := storeImpl.db.Save(&review).Error; err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) DeleteReview(review *model.Review) error {
	if err := storeImpl.db.Delete(&review).Error; err != nil {
		return err
	}
	return nil
}

