package feedBackService

import (
	"todoBackend/app/models"
	"todoBackend/utils"
)

// AddComment 将评论添加到数据库
func AddComment(comments []models.Comment) {
	db := utils.ConnectDB()
	// 创建新的评论记录
	for _, comment := range comments {
		db.Create(&comment)
	}
}

// IncrementHelpful 帮助计数加一
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

// IncrementHelpLess 无帮助计数加一
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
