package main

import (
	"context"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	micro "github.com/micro/go-micro"
	"github.com/minhluuquang/pinklotus/vessel/db"
	"github.com/minhluuquang/pinklotus/vessel/entity"
	pb "github.com/minhluuquang/pinklotus/vessel/proto/vessel"
)

type VesselRepository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type Repository struct {
	DB *gorm.DB
}

func (v *Repository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	vessel := entity.Vessel{}
	if err := v.DB.Where("capacity >= ?", spec.Capacity).Where("max_weight >= ?", spec.MaxWeight).First(&vessel).Error; err != nil {
		return nil, err
	}
	result := entity.VesselDomainToProto(vessel)
	return &result, nil
}

type service struct {
	repo VesselRepository
}

func (s *service) FindAvailable(ctx context.Context, in *pb.Specification, out *pb.Response) error {
	vessel, err := s.repo.FindAvailable(in)
	if err != nil {
		return err
	}

	out.Vessel = vessel
	return nil
}

var (
	// db
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbName = os.Getenv("DB_NAME")
	dbPass = os.Getenv("DB_PASS")
	dbSsl  = os.Getenv("DB_SSL")
)

func main() {
	db, err := db.GetDB(dbHost, dbPort, dbUser, dbPass, dbName, dbSsl)
	if err != nil {
		log.Fatalf("cannot start db: %v", err)
	}

	repo := &Repository{DB: db}

	// New Service
	s := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	// Initialise service
	s.Init()

	// register to auto-genterated code
	pb.RegisterVesselServiceHandler(s.Server(), &service{repo})

	// Run service
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
