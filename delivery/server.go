package delivery

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/config"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/repository"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/usecase"
)

type Server struct {
	// semua usecase disini
	vhUC usecase.VehicleUseCase
}

func (s *Server) Run(menu string) {
	s.initServer(menu)
}

func (s *Server) initServer(menu string) {
	// Semua server di init disini
	NewVehicleCli(s.vhUC, menu)
}

func NewServer() *Server {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	newDbConn, err := config.NewDbConnection(c)
	db := newDbConn.Conn()

	// semua service disini
	vehicleRepo := repository.NewVehicleRepository(db)
	vehicleUseCase := usecase.NewVehicleUseCase(vehicleRepo)
	return &Server{
		vhUC: vehicleUseCase,
	}
}
