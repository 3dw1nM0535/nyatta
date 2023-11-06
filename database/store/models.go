// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package store

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type UnitState string

const (
	UnitStateVacant      UnitState = "vacant"
	UnitStateUnavailable UnitState = "unavailable"
	UnitStateOccupied    UnitState = "occupied"
)

func (e *UnitState) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UnitState(s)
	case string:
		*e = UnitState(s)
	default:
		return fmt.Errorf("unsupported scan type for UnitState: %T", src)
	}
	return nil
}

type NullUnitState struct {
	UnitState UnitState
	Valid     bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUnitState) Scan(value interface{}) error {
	if value == nil {
		ns.UnitState, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UnitState.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUnitState) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.UnitState, nil
}

type UploadCategory string

const (
	UploadCategoryProfileImg   UploadCategory = "profile_img"
	UploadCategoryUnitImages   UploadCategory = "unit_images"
	UploadCategoryCaretakerImg UploadCategory = "caretaker_img"
)

func (e *UploadCategory) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UploadCategory(s)
	case string:
		*e = UploadCategory(s)
	default:
		return fmt.Errorf("unsupported scan type for UploadCategory: %T", src)
	}
	return nil
}

type NullUploadCategory struct {
	UploadCategory UploadCategory
	Valid          bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUploadCategory) Scan(value interface{}) error {
	if value == nil {
		ns.UploadCategory, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UploadCategory.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUploadCategory) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.UploadCategory, nil
}

type Amenity struct {
	ID             int64         `json:"id"`
	Name           string        `json:"name"`
	Provider       string        `json:"provider"`
	Category       string        `json:"category"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	PropertyUnitID sql.NullInt64 `json:"property_unit_id"`
}

type Bedroom struct {
	ID             int64     `json:"id"`
	BedroomNumber  int32     `json:"bedroom_number"`
	EnSuite        bool      `json:"en_suite"`
	Master         bool      `json:"master"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	PropertyUnitID int64     `json:"property_unit_id"`
}

type Caretaker struct {
	ID        int64          `json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Phone     sql.NullString `json:"phone"`
	Verified  bool           `json:"verified"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type Mailing struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

type Property struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	Location    interface{}   `json:"location"`
	Type        string        `json:"type"`
	Status      string        `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	CreatedBy   sql.NullInt64 `json:"created_by"`
	CaretakerID sql.NullInt64 `json:"caretaker_id"`
}

type PropertyUnit struct {
	ID         int64         `json:"id"`
	Name       string        `json:"name"`
	Type       string        `json:"type"`
	State      UnitState     `json:"state"`
	Location   interface{}   `json:"location"`
	Price      int32         `json:"price"`
	Bathrooms  int32         `json:"bathrooms"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
	PropertyID sql.NullInt64 `json:"property_id"`
}

type Shoot struct {
	ID             int64     `json:"id"`
	ShootDate      time.Time `json:"shoot_date"`
	PropertyID     int64     `json:"property_id"`
	PropertyUnitID int64     `json:"property_unit_id"`
	Status         string    `json:"status"`
	CaretakerID    int64     `json:"caretaker_id"`
}

type Tenant struct {
	ID             int64         `json:"id"`
	StartDate      time.Time     `json:"start_date"`
	EndDate        sql.NullTime  `json:"end_date"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	PropertyUnitID sql.NullInt64 `json:"property_unit_id"`
	UserID         sql.NullInt64 `json:"user_id"`
}

type Upload struct {
	ID             int64         `json:"id"`
	Upload         string        `json:"upload"`
	Category       string        `json:"category"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	PropertyUnitID sql.NullInt64 `json:"property_unit_id"`
	PropertyID     sql.NullInt64 `json:"property_id"`
	UserID         sql.NullInt64 `json:"user_id"`
	CaretakerID    sql.NullInt64 `json:"caretaker_id"`
}

type User struct {
	ID         int64          `json:"id"`
	Email      sql.NullString `json:"email"`
	FirstName  sql.NullString `json:"first_name"`
	LastName   sql.NullString `json:"last_name"`
	Phone      string         `json:"phone"`
	Onboarding sql.NullBool   `json:"onboarding"`
	IsLandlord sql.NullBool   `json:"is_landlord"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}
