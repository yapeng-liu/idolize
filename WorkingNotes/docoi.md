# 梦游社项目

[![CN doc](https://img.shields.io/badge/文档-中文版-blue.svg)](docoi.md)

- [1.2.2版本](#122版本)
- [1.3.0版本](#130版本)
- [1.3.1版本](#131版本)
- [1.4.0版本](#140版本)
- [2.4.0版本](#240版本)
- [2.4.1版本](#241版本)
- [2.5.0版本](#250版本)
---

### 1.2.2版本
*  时间：20230731
*  涉及组件：mysql（server_task任务表、user_task_progress用户任务进度表）
  *  任务模块：
     * 任务的创建，sql语句创建
     * 任务的展示，先查询获取server_task任务表所有任务，根据任务ID和用户ID查询user_task_progress用户任务进度表，获取不到用户进度则默认返回空
     * 任务的刷新，创建定时任务，刷新所有任务当天任务进度，根据任务类型（日常任务），每日定时刷新任务进度表中的任务进度
       * 任务的触发，在固定的事件下添加任务触发事件的代码，目前以2倍数数值与运算是否大于0来判断任务是否触发
       ~~~
         enum event {
             EventNone       = 0 ; //  占位,无意义
             SignIn          = 1 ; //  积分机器人签到事件
             ActivityLike    = 2 ; //  动态点赞事件
             ActivityComment = 4 ; //  动态评论或者回复评论事件
             ActivityCreate  = 8 ; //  动态创建事件
          }
       ~~~
*  总结：
   
### 1.3.0版本
*  时间：20230817
*  涉及组件：mysql、mangodb、mqtt
*  圈子频道-谁可发布动态：
   * 圈子频道表，设置可发布字段(release_mode 0全部成员、1指定成员)
   * 谁可以发布权限表，设置为指定成员可发布时需要查询表中圈子频道设置的权限，在圈子频道创建时默认添加不可修改的两个身份组
   * 身份组添加，操作谁可以发布权限表
   * 身份组删除，操作谁可以发布权限表
   * 发布时查询：
      * 首先查看圈子频道设置
      * 设置为全体成员则直接返回可发布
      * 设置为指定成员则需要查看和比较用户身份组信息和谁可以发布权限表，如果比对数据中存在相同数据，返回可发布，否则不可发布
*  圈子频道-修改动态所属频道：
   *  修改动态所属频道字段
*  圈子频道-增加动态仅频道可见设置：
   *  修改mangodb存储的动态 可见属性
   *  查找动态增加过滤选项
### 1.3.1版本
*  时间：20230913
*  涉及组件：mysql
*  身份组信息调整：
   *  查询请求的客户端版本进行低版本单独处理，版本隔离
*  离线推送设置
   *  腾讯云IM官网：https://cloud.tencent.com/document/product/269/2720#.E7.A6.BB.E7.BA.BF.E6.8E.A8.E9.80.81-offlinepushinfo-.E8.AF.B4.E6.98.8E
   *  配置OfflinePushInfo 离线推送 JSON 对象和前端对接
   *  视频审核由第三方数美进行审核

### 1.4.0版本
*  时间：20230817
*  涉及组件：mysql

### 2.4.0版本
~~~
    
~~~


### 2.4.1版本
#### 1.版本开发概要
~~~
    个人:
        圈子官方号群聊龙王消息定时早上9点发送
        动态自动加入推荐列表、动态操作显示规则调整
        管理后台礼物和商品列表接口
~~~

#### 2.优化线上数据库备份方案

* 备份文件压缩 tar -zcvf
* 降低备份程序IO占用优先级 nice、ionice、cpulimit

#### 3.转移服务器主
* 1.修改servers表中server所属用户ID
* 2.查看identity_groups表服务器所有身份组，找出服务器主身份ID
* 3.修改identity_group_members表，用户ID的最高身份ID为服务器主身份ID

#### 4.用户在线时长分析
[用户在线时长的统计.docx](%E7%94%A8%E6%88%B7%E5%9C%A8%E7%BA%BF%E6%97%B6%E9%95%BF%E7%9A%84%E7%BB%9F%E8%AE%A1.docx)

#### 5.自动替换证书流程

* https://github.com/certd/certd

### 2.5.0版本
#### 版本开发概要
~~~
 
 个人：
 任务模块增加 任务领取方式配置、任务完成进度提示
 动态助手消息（点赞、评论）功能重做，之前的实现方式是IM消息,改为本地查询mongo数据，主要实现（消息列表查询、消息旧数据兼容、消息MQTT实时通知、消息红点通知）
~~~

#### 版本重点功能记录
* 有关mongodb的bson文档 
* D: 有序的BSON文档,由一系列的 bson.E 元素构成
~~~
// D is an ordered representation of a BSON document. This type should be used when the order of the elements matters,
// such as MongoDB command documents. If the order of the elements does not matter, an M should be used instead.
//
// A D should not be constructed with duplicate key names, as that can cause undefined server behavior.
//
// Example usage:
//
//	bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}}
type D = primitive.D
~~~
* E: BSON文档中一个元素
~~~
// E represents a BSON element for a D. It is usually used inside a D.
type E = primitive.E
~~~
* M: 无序的BSON文档，Map类型，bson.M 通常用于构建 bson.D
~~~
// M is an unordered representation of a BSON document. This type should be used when the order of the elements does not
// matter. This type is handled as a regular map[string]interface{} when encoding and decoding. Elements will be
// serialized in an undefined, random order. If the order of the elements matters, a D should be used instead.
//
// Example usage:
//
//	bson.M{"foo": "bar", "hello": "world", "pi": 3.14159}
type M = primitive.M
~~~
* A: BSON 文档中的数组,使用“$and”,“$or”等
~~~
// An A is an ordered representation of a BSON array.
//
// Example usage:
//
//	bson.A{"bar", "world", 3.14159, bson.D{{"qux", 12345}}}
type A = primitive.A
~~~
* 查找评论
~~~
    var comments []Comment
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"$or", bson.A{
				bson.D{{"ask_user_ids", bson.M{"$in": []int64{toUserId}}}},
				bson.D{{"to_user_id", toUserId}},
			}}},
			bson.D{{"user_id", bson.M{"$ne": toUserId}}},
		}},
	}
	if sortId != "" {
		objectId, err := primitive.ObjectIDFromHex(sortId)
		if err != nil {
			r.log.Error("GetUserBeCommentedByPage ObjectIDFromHex err: ", err)
			return nil, errors.WithStack(err)
		}
		filter = append(filter, bson.E{Key: "_id", Value: bson.M{"$lt": objectId}})
	}
	var opts []*options.FindOptions
	opts = append(opts, options.Find().SetLimit(limit))
	opts = append(opts, options.Find().SetSort(bson.M{"_id": -1}))
	cur, err := r.commentColl.Find(ctx, filter, opts...)
		if err = cur.All(ctx, &comments); err != nil {
		r.log.Error("GetUserBeCommentedByPage All err: ", err)
		return nil, err
	}
	if err = cur.All(ctx, &comments); err != nil {
		r.log.Error("GetUserBeCommentedByPage All err: ", err)
		return nil, err
	}
~~~
* 查找点赞
~~~
	var likes []Like
	filter := bson.D{{Key: "to_user", Value: toUserId}}
	filter = append(filter, bson.E{Key: "user_id", Value: bson.M{"$ne": toUserId}})
	filter = append(filter, bson.E{Key: "object_type", Value: bson.M{"$in": objectTypes}})
	if sortId != "" {
		objectId, err := primitive.ObjectIDFromHex(sortId)
		if err != nil {
			r.log.Error("GetUserBeCommentedByPage ObjectIDFromHex err: ", err)
			return nil, errors.WithStack(err)
		}
		filter = append(filter, bson.E{Key: "_id", Value: bson.M{"$lt": objectId}})
	}
	var opts []*options.FindOptions
	if limit > 0 {
		opts = append(opts, options.Find().SetLimit(limit))
	}
	opts = append(opts, options.Find().SetSort(bson.M{"_id": -1}))
	cur, err := r.likeColl.Find(ctx, filter, opts...)
	if err != nil {
		r.log.Error("GetUserBeLikedByPage err: ", err)
		return nil, err
	}
	if err = cur.All(ctx, &likes); err != nil {
		r.log.Error("GetUserBeLikedByPage err: ", err)
		return nil, err
	}
~~~

#### SQL优化
* 场景
    获取服务器成员的服务器昵称
* 单个查找改为批量查找
~~~
	err := r.data.DB(ctx).Where("circle_id = ? and user_id = ?", cid, uid).First(&serverMember).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		r.log.Error("GetServerMemberByUidCid err :", err)
		return nil, err
	}
    优化：
    uIdsCIds [][]interface{}
    uIdsCIds = append(uIdsCIds, []interface{}{uid, cid})
	if err := r.data.DB(ctx).Model(&ServerMemberPo{}).Where("(user_id,circle_id) IN ?", uIdsCIds).Find(&cms).Error; err != nil {
		r.log.Errorf("get circleUserIds(%v) err(%v)", uIdsCIds, err)
		return nil, err
	}

~~~

