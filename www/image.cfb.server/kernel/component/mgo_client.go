package component

// 思路：
// 1. 此文件作为操作mongo的客户端组件
// 2. 然后提供操作的API给base_model进行二次封装

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"image-cfb-server/kernel/common"
	"time"
)

/**
	用于管理mongo的链接
 */
type mgoConnection struct {
	connections map[string]*mongo.Client
}

func (mgoConn *mgoConnection)getConnection(url string)  *mongo.Client{
	// 把uri 转换为 md5
	var uriMd5 string =common.H.Md5(url)

	if connection,ok := mgoConn.connections[url]; ok {
		return connection
	}

	// 链接
	mgoConn.setConnection()
}

func (mgoConn *mgoConnection)setConnection() *mongo.Client {
	var err error
	// 设置链接uri
	// "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(uri)
	// 设置连接池
	clientOptions.SetMaxPoolSize(100)
	// 设置最大空闲时间
	clientOptions.SetMaxConnIdleTime( 10 *time.Minute)
	// 设置最大链接时长，超过而认为链接失败
	clientOptions.SetConnectTimeout(3 * time.Second)

	mgo.ctx, mgo.cancel = context.WithTimeout()

	// 连接到MongoDB
	var mgoCli *mongo.Client
	mgoCli, err = mongo.Connect(mgo.ctx, clientOptions)
	if err != nil {
		panic("mongo链接失败："+ err.Error())
	}

	mgo.client = mgoCli

	return
}

type mgoClient struct {
	databaseName string	// 数据库名称
	collectionName string // 集合名称
	client *mongo.Client // mongo 链接
	collection *mongo.Collection // mongo 链接
	database *mongo.Database // mongo 链接
	ctx context.Context
	cancel context.CancelFunc
}

func NewMgoClient()  *mgoClient {
	return new(mgoClient)
}

/**
	链接，此处有个问题，每次实例化，都得重新链接一次，这不合理
 */
func (mgo *mgoClient)SetConnection(uri string) {

}

/**
	设置使用的集合
 */
func (mgo *mgoClient) find () {
	//collection := mgo.client.Database(mgo)
}
