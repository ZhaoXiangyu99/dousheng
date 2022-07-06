package request

// =============================================封装发起请求的参数=========================================================

// ==============================================用户相关的参数===========================================================

// UserParam 登录注册的参数
type UserParam struct {
	UserName string `form:"username"` //此处不能用json,必须要用form,因为post方法的参数是在form里面
	PassWord string `form:"password"`
}

// ==============================================关注相关的参数===========================================================

// FollowActionParam 添加或取消关注的参数
type FollowActionParam struct {
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id"`
	ActionType int64  `form:"action_type"`
}

// RelationParam 获取关注列表和粉丝列表的参数
type RelationParam struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

// ==================================================视频相关参数=========================================================

// FeedParam 视频流相关参数
type FeedParam struct {
	LatestTime int64  `form:"latest_time"`
	Token      string `form:"token"`
}

// PublishListParam 获取视频发布列表的参数
type PublishListParam struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

// PublishParam 视频投稿相关参数
type PublishParam struct {
	Token string `form:"token" binding:"required"`
	Title string `form:"title" binding:"required"`
}

// ==================================================评论相关参数=========================================================

// CommentActionParam 添加或删除评论相关参数
type CommentActionParam struct {
	UserId      int64  `form:"user_id"`
	Token       string `form:"token"`
	VideoId     int64  `form:"video_id"`
	ActionType  int64  `form:"action_type"`
	CommentText string `form:"comment_text"`
	DeleteId    int64  `form:"comment_id"`
}

// CommentListParam 评论列表相关参数
type CommentListParam struct {
	Token   string `form:"token"`
	VideoId int64  `form:"video_id"`
}

// ====================================================点赞相关参数=======================================================

// FavoriteActionParam 点赞或取消点赞相关参数
type FavoriteActionParam struct {
	Token      string `form:"token"`
	VideoId    int64  `form:"video_id"`
	ActionType int64  `form:"action_type"`
}

// FavoriteListParam 点赞列表相关参数
type FavoriteListParam struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}
