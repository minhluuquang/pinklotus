package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	pb "github.com/minhluuquang/pinklotus/vessel/proto/vessel"
)

// Consignment is a struct for domain vessel
type Vessel struct {
	ID        string     `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Capacity  int32      `json:"capacity,omitempty"`
	MaxWeight int32      `json:"max_weight,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func VesselProtoToDomain(v pb.Vessel) Vessel {
	return Vessel{
		ID:        v.Id,
		Name:      v.Name,
		Capacity:  v.Capacity,
		MaxWeight: v.MaxWeight,
	}
}

func VesselDomainToProto(v Vessel) pb.Vessel {
	return pb.Vessel{
		Id:        v.ID,
		Name:      v.Name,
		Capacity:  v.Capacity,
		MaxWeight: v.MaxWeight,
	}
}

// BeforeCreate prepare data before create data
func (d *Vessel) BeforeCreate(scope *gorm.Scope) error {
	d.ID = uuid.New().String()
	return nil
}
