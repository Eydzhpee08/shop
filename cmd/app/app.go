package app

import (
	"net"

	"github.com/Eydzhpee08/first-api/pkg/bills"
	"github.com/Eydzhpee08/first-api/pkg/products"
	"github.com/Eydzhpee08/first-api/pkg/customers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	productSvc *products.Service
	customerSvc *customers.Service
	billSvc *bills.Service
}

func NewServer(router *gin.Engine, productSvc *products.Service, customerSvc *customers.Service, billSvc *bills.Service) *Server{
	return &Server{router: router, productSvc: productSvc, customerSvc: customerSvc, billSvc: billSvc}
}

func (s *Server) Run(host, port string) error {
	return s.router.Run(net.JoinHostPort(host, port))
}