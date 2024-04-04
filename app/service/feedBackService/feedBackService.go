package feedBackService

import (
	"todoBackend/app/models"
	"todoBackend/utils"
)

func AddComment(comments []models.Comment) {
	db := utils.ConnectDB()
	// 创建新的评论记录
	for _, comment := range comments {
		db.Create(&comment)
	}
}
func IncrementHelpful() error {
	db := utils.ConnectDB()
	var feedback models.FeedBack
	if err := db.FirstOrCreate(&feedback).Error; err != nil {
		return err
	}
	feedback.Helpful++
	db.Save(&feedback)
	return nil
}
func IncrementHelpLess() error {
	db := utils.ConnectDB()
	var feedback models.FeedBack
	if err := db.FirstOrCreate(&feedback).Error; err != nil {
		return err
	}
	feedback.Helpless++
	db.Save(&feedback)
	return nil
}
