package main

import (
	"ERC1155/controller"
	"ERC1155/model"
	"ERC1155/router"
	"ERC1155/services"
	"ERC1155/util"
	"context"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

func main(){
	// load config vars ...
	config, err := util.LoadConfig(".")
	if err != nil{
		panic(err)
	}

	// database ...
	db, err := gorm.Open(postgres.Open(config.DBSource))
	db.AutoMigrate(&model.Block{})
	db.AutoMigrate(&model.Transaction{})

	blockService := services.NewBlockService(db, config)

	// block tracker ...
	blockService.TrackBlock()
	blockController := controller.NewBlockController(blockService)


	sm := mux.NewRouter()
	sm.Use(util.Middleware)
	route := router.NewRouter(sm)


	route.GET("/transactions", blockController.GetTransaction)
	route.GET("/", func(w http.ResponseWriter, r *http.Request) {
		util.SuccessResponse(w, 200, "ERC1155", nil)
	})

	server := http.Server{
		Addr: config.PORT,
		Handler: sm,

	}

	go func() {
		log.Printf("Server is listening to %v", config.PORT)
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Gracefully shutdown of server
	signChan := make(chan os.Signal)
	signChan <- os.Kill
	signChan <- os.Interrupt
	ctx, _ := context.WithTimeout(context.Background(), 20 * time.Second)
	server.Shutdown(ctx)
}

