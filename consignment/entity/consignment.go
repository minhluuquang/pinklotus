package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	pb "github.com/minhluuquang/pinklotus/consignment/proto/consignment"
)

// Consignment is a struct for domain consignment
type Consignment struct {
	ID               string     `json:"id,omitempty"`
	VesselID         *string    `json:"vessel_id,omitempty"`
	ContainersNumber int32      `json:"containers_number,omitempty"`
	Weight           int32      `json:"weight,omitempty"`
	Description      string     `json:"description,omitempty"`
	CreatedAt        time.Time  `json:"created_at,omitempty"`
	UpdatedAt        time.Time  `json:"updated_at,omitempty"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

func ConsignmentProtoToDomain(c pb.Consignment) Consignment {
	return Consignment{
		ID:               c.Id,
		VesselID:         &c.VesselId,
		ContainersNumber: c.ContainersNumber,
		Weight:           c.Weight,
		Description:      c.Description,
	}
}

func ConsignmentDomainToProto(c Consignment) pb.Consignment {
	return pb.Consignment{
		Id:               c.ID,
		VesselId:         *c.VesselID,
		ContainersNumber: c.ContainersNumber,
		Weight:           c.Weight,
		Description:      c.Description,
	}
}

// BeforeCreate prepare data before create data
func (d *Consignment) BeforeCreate(scope *gorm.Scope) error {
	d.ID = uuid.New().String()
	return nil
}
