package application

import (
	"context"

	"github.com/copchase/slapyou2/internal/client/dynamodb"
	slapHandler "github.com/copchase/slapyou2/internal/handler/slap"
	"github.com/copchase/slapyou2/internal/repository/slapstats"
	slapSvc "github.com/copchase/slapyou2/internal/service/slap"
	"github.com/julienschmidt/httprouter"
)

type application struct {
	ctx context.Context
}

func NewApplication(ctx context.Context) *application {

	// construct all clients
	dynamoClient := dynamodb.NewDynamoClient()

	// construct all repositories
	slapStatsRepo := slapstats.NewSlapStatsRepo(dynamoClient)

	// construct all services
	slapSvc := slapSvc.NewSlapService(slapStatsRepo)

	// construct all handlers
	slapHandler := slapHandler.NewSlapHandler(slapSvc)

	router := httprouter.New()
	slapHandler.RegisterRoutes(router)

	// http server turn on

	return &application{
		ctx: ctx,
	}
}

func (a *application) Run() {
}
