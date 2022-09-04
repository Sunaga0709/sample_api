package main

import (
	"fmt"
	"strconv"

	"sample_api/model"

	"github.com/gin-gonic/gin"
)

func init() {
	model.InitDB()
}

func main() {
	// Ginインスタンスとルーティングを取得
	router := getRouter()

	// ポート8080番でサーバ起動
	router.Run()
}

type InputName struct {
	Name string `json:"name" binding:"required"`
}

// Ginインスタンスとルーティング設定
func getRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", getAllUser)        // 全ユーザ取得
	router.POST("/users", createUser)       // ユーザ作成
	router.GET("/users/:id", getOneUser)    // ユーザ取得
	router.PUT("/users/:id", updateUser)    // ユーザ情報の更新
	router.DELETE("/users/:id", deleteUser) // ユーザ削除

	return router
}

func getAllUser(ctx *gin.Context) {
	users, modelErr := model.GetAllUser()
	if modelErr != nil {
		ctx.JSON(404, gin.H{"message": "not found"})
	} else {
		ctx.JSON(200, users)
	}
}

func createUser(ctx *gin.Context) {
	var name InputName
	ctx.BindJSON(&name)
	model.CreateUser(name.Name)
	ctx.JSON(200, nil)
}

func getOneUser(ctx *gin.Context) {
	// パスからIDを取得
	id, strErr := strconv.Atoi(ctx.Param("id"))
	if strErr != nil {
		fmt.Println(strErr)
	}
	user, modelErr := model.GetOneUser(id)
	if modelErr != nil {
		ctx.JSON(404, gin.H{"message": "not found"})
	} else {
		ctx.JSON(200, user)
	}
}

func updateUser(ctx *gin.Context) {
	// パスからIDを取得
	id, strErr := strconv.Atoi(ctx.Param("id"))
	if strErr != nil {
		fmt.Println(strErr)
	}

	var name InputName
	ctx.BindJSON(&name)
	modelErr := model.UpdateUser(id, name.Name)
	if modelErr != nil {
		ctx.JSON(404, gin.H{"message": "not found"})
	} else {
		ctx.JSON(200, nil)
	}
}

func deleteUser(ctx *gin.Context) {
	// パスからIDを取得
	id, strErr := strconv.Atoi(ctx.Param("id"))
	if strErr != nil {
		fmt.Println(strErr)
	}

	modelErr := model.DeleteUser(id)
	if modelErr != nil {
		ctx.JSON(404, gin.H{"message": "not found"})
	} else {
		ctx.JSON(200, nil)
	}
}
