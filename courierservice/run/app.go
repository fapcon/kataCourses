package run

import (
	"context"
	"courierservice/cache"
	"courierservice/geo"
	cservice "courierservice/module/courier/service"
	cstorage "courierservice/module/courier/storage"
	"courierservice/module/courierfacade/controller"
	cfservice "courierservice/module/courierfacade/service"
	oservice "courierservice/module/order/service"
	ostorage "courierservice/module/order/storage"
	"courierservice/router"
	"courierservice/server"
	"courierservice/workers/order"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() error {
	// получение хоста и порта redis
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	// инициализация клиента redis
	rclient := cache.NewRedisClient(host, port)

	// инициализация контекста с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// проверка доступности redis
	_, err := rclient.Ping(ctx).Result()
	if err != nil {
		return err
	}

	fmt.Println("redis inintialized")

	// инициализация разрешенной зоны
	allowedZone := geo.NewAllowedZone()
	fmt.Println("allowed zone inintialized")
	// инициализация запрещенных зон
	disAllowedZones := []geo.PolygonChecker{geo.NewDisAllowedZone1(), geo.NewDisAllowedZone2()}
	fmt.Println("disabled zones inintialized")

	// инициализация хранилища заказов
	orderStorage := ostorage.NewOrderStorage(rclient)
	fmt.Println("order storage inintialized")
	// инициализация сервиса заказов
	orderService := oservice.NewOrderService(orderStorage, allowedZone, disAllowedZones)
	fmt.Println("order service inintialized")

	orderGenerator := order.NewOrderGenerator(orderService)
	orderGenerator.Run()
	fmt.Println("order generator run")

	oldOrderCleaner := order.NewOrderCleaner(orderService)
	oldOrderCleaner.Run()
	fmt.Println("order cleaner run")

	// инициализация хранилища курьеров
	courierStorage := cstorage.NewCourierStorage(rclient)
	fmt.Println("courier storage initialized")
	// инициализация сервиса курьеров
	courierSevice := cservice.NewCourierService(courierStorage, allowedZone, disAllowedZones)
	fmt.Println("courier service initialized")

	// инициализация фасада сервиса курьеров
	courierFacade := cfservice.NewCourierFacade(courierSevice, orderService)
	fmt.Println("courier facade initialized")

	// инициализация контроллера курьеров
	courierController := controller.NewCourierController(courierFacade)
	fmt.Println("courier controller initialized")

	// инициализация роутера
	routes := router.NewRouter(courierController)
	fmt.Println("router initialized")
	// инициализация сервера
	r := server.NewHTTPServer()
	fmt.Println("geogrpc initialized")
	// инициализация группы роутов
	api := r.Group("/api")

	// инициализация роутов
	routes.CourierAPI(api)
	fmt.Println("api group routes initialized")

	mainRoute := r.Group("/")

	routes.Swagger(mainRoute)
	fmt.Println("swagger route initialized")

	// инициализация статических файлов
	r.NoRoute(gin.WrapH(http.FileServer(http.Dir("public"))))
	fmt.Println("static routes initialized")

	// запуск сервера
	serverPort := os.Getenv("SERVER_PORT")

	if os.Getenv("ENV") == "prod" {
		certFile := "/app/certs/cert.pem"
		keyFile := "/app/certs/private.pem"
		return r.RunTLS(":443", certFile, keyFile)
	}

	fmt.Println("launching geogrpc...")
	return r.Run(":" + serverPort)
}
