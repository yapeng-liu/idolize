package main

import "fmt"

func main() {
	servers := []int64{1, 2, 3, 4, 5}
	whiteServers := []int64{1, 2}

	//过滤白名单服务器
	filteredServerIds := make([]int64, 0)
	whiteServerMap := make(map[int64]struct{})
	for _, serverId := range whiteServers {
		fmt.Println("123")
		whiteServerMap[serverId] = struct{}{}
	}
	for _, serverId := range servers {
		if _, ok := whiteServerMap[serverId]; !ok {
			filteredServerIds = append(filteredServerIds, serverId)
		}
	}
	fmt.Println(filteredServerIds)
}

//func (s *securityRepo) InsertActionLog(ctx context.Context, userID, actionId int64, actionDetail string) error {
//	if _, err := s.actionLogsColl.
//		InsertOne(ctx, bson.D{
//			{"user_id", userID},
//			{"action_id", actionId},
//			{"action_detail", actionDetail},
//			{"created_at", time.Now().Unix()},
//		}); err != nil {
//		s.log.Error("insert action log mongodb err: ", err)
//	}
//	return nil
//}

type ActionLog struct {
	ID           string `bson:"_id,omitempty"`
	UserId       int64  `bson:"user_id"`       // 用户id
	ActionId     int64  `bson:"action_id"`     // 事件id
	ActionDetail string `bson:"action_detail"` // 事件详情
	CreatedAt    int64  `bson:"created_at"`    // 创建时间
}
