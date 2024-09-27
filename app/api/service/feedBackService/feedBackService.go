package feedBackService

import (
	"todoBackend/app/models/feedback_model"
	"todoBackend/utils/db"
)

// AddComment 将评论添加到数据库
func AddComment(comments []feedback_model.Comment) {
	// 创建新的评论记录
	for _, comment := range comments {
		db.DB.Create(&comment)
	}
}

// IncrementHelpful 帮助计数加一
func IncrementHelpful() error {
	var feedback feedback_model.FeedBack
	if err := db.DB.FirstOrCreate(&feedback).Error; err != nil {
		return err
	}
	feedback.Helpful++
	db.DB.Save(&feedback)
	return nil
}

// IncrementHelpLess 无帮助计数加一
func IncrementHelpLess() error {
	var feedback feedback_model.FeedBack
	if err := db.DB.FirstOrCreate(&feedback).Error; err != nil {
		return err
	}
	feedback.Helpless++
	db.DB.Save(&feedback)
	return nil
}
