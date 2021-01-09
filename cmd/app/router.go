package app

func (s *Server) InitRoutes() {
	//products
	s.router.POST("/app/product/list", s.GetProductAll)
	s.router.POST("/app/product", s.AddProducts)
	s.router.PUT("/app/product", s.EditProducts)
	s.router.DELETE("/app/product/:id", s.RemoveProducts)

	//customers
	s.router.POST("/app/customers/list", s.GetCustomersAll)
	s.router.POST("/app/customers", s.AddCustomer)
	s.router.PUT("/app/customers", s.EditCustomer)
	s.router.DELETE("/app/customers/:id", s.RemoveCustomer)

	//bill
	s.router.POST("/app/bill/list", s.GetBillAll)
	s.router.POST("/app/bill", s.AddBill)


}