package feedback_model

import "gorm.io/gorm"

// FeedBack 反馈模型
type FeedBack struct {
	gorm.Model
	Helpful  int `json:"helpful"`  // 有帮助的
	Helpless int `json:"helpless"` // 无帮助的
}

// Comment 评论模型
type Comment struct {
	gorm.Model
	Text string `json:"text"` // 文本
}
