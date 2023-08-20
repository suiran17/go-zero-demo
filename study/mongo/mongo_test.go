package mongo

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/zeromicro/go-zero/core/stores/mon"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

func TestMongoDB(t *testing.T) {

	conn := mon.MustNewModel("mongodb://admin:admin@dev.in:27017", "db", "collection")
	// conn := mon.MustNewModel("mongodb://<admin>:<admin>@<dev.in>:<27017>", "db", "collection")
	// conn := mon.MustNewModel("mongodb://<user>:<password>@<host>:<port>", "db", "collection")
	ctx := context.Background()
	u := &User{
		ID:       primitive.ObjectID{},
		Username: "username1",
		Password: "password1",
		UpdateAt: time.Now(),
		CreateAt: time.Now(),
	}
	// insert one
	ret, err := conn.InsertOne(ctx, u)
	if err != nil {
		panic(err)
	}

	fmt.Println(ret)

	// 查询
	var newUser User
	err = conn.FindOne(ctx, &newUser, bson.M{"_id": ret.InsertedID})
	// err = conn.FindOne(ctx, &newUser, bson.M{"_id": u.ID})
	if err != nil {
		panic(err)
	}

	// 更新
	newUser.Username = "newUsername"
	_, err = conn.ReplaceOne(ctx, bson.M{"_id": newUser.ID}, newUser)
	if err != nil {
		panic(err)
	}

	// 删除
	_, err = conn.DeleteOne(ctx, bson.M{"_id": newUser.ID})
	if err != nil {
		panic(err)
	}
}
