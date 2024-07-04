package main

import (
	"database/sql"
	"net/http"

	"github.com/devfullcycle/imersao18/golang/internal/events/infra/repository"
	"github.com/devfullcycle/imersao18/golang/internal/events/infra/service"
	"github.com/devfullcycle/imersao18/golang/internal/events/usecase"

	httpHandler "github.com/devfullcycle/imersao18/golang/internal/events/infra/http"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao18")
	if err != nil {
		panic(err)
	}
	db.Close()

	eventRepo, err := repository.NewMysqlEventRepository(db)
	if err != nil {
		panic(err)
	}

	partnerBaseURLs := map[int]string{
		1: "http://localhost:9080/api1",
		2: "http://localhost:9080/api2",
	}

	partnerFactory := service.NewPartnerFactory(partnerBaseURLs)

	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)
	buyTicketUseCase := usecase.NewBuyTicketsUseCase(eventRepo, partnerFactory)

	eventsHandler := httpHandler.NewEventsHandler(
		listEventsUseCase,
		listSpotsUseCase,
		getEventUseCase,
		buyTicketUseCase,
	)

	r := http.NewServeMux()
	r.HandleFunc("/events", eventsHandler.ListEvents)
	r.HandleFunc("/events/{event_id}", eventsHandler.GetEvent)
	r.HandleFunc("/events/{event_id}/spots", eventsHandler.ListSpots)
	r.HandleFunc("POST /checkout", eventsHandler.BuyTickets)

	http.ListenAndServe(":8080", r)
}
