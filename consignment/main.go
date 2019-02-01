package main

import (
	"context"
	"log"

	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/micro/go-micro"

	"github.com/minhluuquang/pinklotus/consignment/db"
	"github.com/minhluuquang/pinklotus/consignment/entity"
	pb "github.com/minhluuquang/pinklotus/consignment/proto/consignment"
	vesselPb "github.com/minhluuquang/pinklotus/vessel/proto/vessel"
)

// ConsignmentReposityory --
type ConsignmentRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

type Repository struct {
	DB *gorm.DB
}

// Create a consignment in db
func (r *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	c := entity.ConsignmentProtoToDomain(*consignment)
	if err := r.DB.Create(&c).Error; err != nil {
		return nil, err
	}
	result := entity.ConsignmentDomainToProto(c)
	return &result, nil
}

type service struct {
	repo         ConsignmentRepository
	vesselClient vesselPb.VesselService
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	vessel, err := s.vesselClient.FindAvailable(ctx, &vesselPb.Specification{
		MaxWeight: req.Weight,
		Capacity:  req.ContainersNumber,
	})
	if err != nil {
		return err
	}

	req.VesselId = vessel.Vessel.Id

	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	resp.Consigment = consignment

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
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselPb.NewVesselService("go.micro.srv.vessel", s.Client())

	// Initialise service
	s.Init()

	// register to auto-genterated code
	pb.RegisterShippingServiceHandler(s.Server(), &service{repo, vesselClient})

	// Run service
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
