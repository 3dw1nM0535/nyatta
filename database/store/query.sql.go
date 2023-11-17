// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package store

import (
	"context"
	"database/sql"
	"time"
)

const createAmenity = `-- name: CreateAmenity :one
INSERT INTO amenities (
  name, category, property_unit_id
) VALUES (
  $1, $2, $3
)
RETURNING id, name, provider, category, created_at, updated_at, property_unit_id
`

type CreateAmenityParams struct {
	Name           string        `json:"name"`
	Category       string        `json:"category"`
	PropertyUnitID sql.NullInt64 `json:"property_unit_id"`
}

func (q *Queries) CreateAmenity(ctx context.Context, arg CreateAmenityParams) (Amenity, error) {
	row := q.db.QueryRowContext(ctx, createAmenity, arg.Name, arg.Category, arg.PropertyUnitID)
	var i Amenity
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Provider,
		&i.Category,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PropertyUnitID,
	)
	return i, err
}

const createCaretaker = `-- name: CreateCaretaker :one
INSERT INTO caretakers (
  first_name, last_name, phone
) VALUES (
  $1, $2, $3
)
RETURNING id, first_name, last_name, phone, verified, created_at, updated_at
`

type CreateCaretakerParams struct {
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Phone     sql.NullString `json:"phone"`
}

func (q *Queries) CreateCaretaker(ctx context.Context, arg CreateCaretakerParams) (Caretaker, error) {
	row := q.db.QueryRowContext(ctx, createCaretaker, arg.FirstName, arg.LastName, arg.Phone)
	var i Caretaker
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Verified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createInvoice = `-- name: CreateInvoice :one
INSERT INTO invoices (
  reference, phone, msid
) VALUES (
  $1, $2, $3
)
RETURNING id, msid, channel, currency, bank, auth_code, country_code, fees, amount, phone, status, reference, created_at, updated_at
`

type CreateInvoiceParams struct {
	Reference sql.NullString `json:"reference"`
	Phone     sql.NullString `json:"phone"`
	Msid      sql.NullString `json:"msid"`
}

func (q *Queries) CreateInvoice(ctx context.Context, arg CreateInvoiceParams) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, createInvoice, arg.Reference, arg.Phone, arg.Msid)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.Msid,
		&i.Channel,
		&i.Currency,
		&i.Bank,
		&i.AuthCode,
		&i.CountryCode,
		&i.Fees,
		&i.Amount,
		&i.Phone,
		&i.Status,
		&i.Reference,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createProperty = `-- name: CreateProperty :one
INSERT INTO properties (
  name, type, created_by, location
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, name, location, type, created_at, updated_at, created_by, caretaker_id
`

type CreatePropertyParams struct {
	Name      string        `json:"name"`
	Type      interface{}   `json:"type"`
	CreatedBy sql.NullInt64 `json:"created_by"`
	Location  interface{}   `json:"location"`
}

func (q *Queries) CreateProperty(ctx context.Context, arg CreatePropertyParams) (Property, error) {
	row := q.db.QueryRowContext(ctx, createProperty,
		arg.Name,
		arg.Type,
		arg.CreatedBy,
		arg.Location,
	)
	var i Property
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Location,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.CaretakerID,
	)
	return i, err
}

const createPropertyUnit = `-- name: CreatePropertyUnit :one
INSERT INTO property_units (
  property_id, bathrooms, name, type, price
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, name, type, state, price, bathrooms, created_at, updated_at, property_id
`

type CreatePropertyUnitParams struct {
	PropertyID sql.NullInt64 `json:"property_id"`
	Bathrooms  int32         `json:"bathrooms"`
	Name       string        `json:"name"`
	Type       string        `json:"type"`
	Price      int32         `json:"price"`
}

func (q *Queries) CreatePropertyUnit(ctx context.Context, arg CreatePropertyUnitParams) (PropertyUnit, error) {
	row := q.db.QueryRowContext(ctx, createPropertyUnit,
		arg.PropertyID,
		arg.Bathrooms,
		arg.Name,
		arg.Type,
		arg.Price,
	)
	var i PropertyUnit
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.State,
		&i.Price,
		&i.Bathrooms,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PropertyID,
	)
	return i, err
}

const createShootSchedule = `-- name: CreateShootSchedule :one
INSERT INTO shoots (
  shoot_date, property_id
) VALUES (
  $1, $2
)
RETURNING id, shoot_date, property_id, property_unit_id, status, caretaker_id, created_at, updated_at
`

type CreateShootScheduleParams struct {
	ShootDate  time.Time `json:"shoot_date"`
	PropertyID int64     `json:"property_id"`
}

func (q *Queries) CreateShootSchedule(ctx context.Context, arg CreateShootScheduleParams) (Shoot, error) {
	row := q.db.QueryRowContext(ctx, createShootSchedule, arg.ShootDate, arg.PropertyID)
	var i Shoot
	err := row.Scan(
		&i.ID,
		&i.ShootDate,
		&i.PropertyID,
		&i.PropertyUnitID,
		&i.Status,
		&i.CaretakerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createTenant = `-- name: CreateTenant :one
INSERT INTO tenants (
  start_date, property_unit_id
) VALUES (
  $1, $2
)
RETURNING id, start_date, end_date, created_at, updated_at, property_unit_id, user_id
`

type CreateTenantParams struct {
	StartDate      time.Time     `json:"start_date"`
	PropertyUnitID sql.NullInt64 `json:"property_unit_id"`
}

func (q *Queries) CreateTenant(ctx context.Context, arg CreateTenantParams) (Tenant, error) {
	row := q.db.QueryRowContext(ctx, createTenant, arg.StartDate, arg.PropertyUnitID)
	var i Tenant
	err := row.Scan(
		&i.ID,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PropertyUnitID,
		&i.UserID,
	)
	return i, err
}

const createUnitBedroom = `-- name: CreateUnitBedroom :one
INSERT INTO bedrooms (
  property_unit_id, bedroom_number, en_suite, master
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, bedroom_number, en_suite, master, property_unit_id, created_at, updated_at
`

type CreateUnitBedroomParams struct {
	PropertyUnitID int64 `json:"property_unit_id"`
	BedroomNumber  int32 `json:"bedroom_number"`
	EnSuite        bool  `json:"en_suite"`
	Master         bool  `json:"master"`
}

func (q *Queries) CreateUnitBedroom(ctx context.Context, arg CreateUnitBedroomParams) (Bedroom, error) {
	row := q.db.QueryRowContext(ctx, createUnitBedroom,
		arg.PropertyUnitID,
		arg.BedroomNumber,
		arg.EnSuite,
		arg.Master,
	)
	var i Bedroom
	err := row.Scan(
		&i.ID,
		&i.BedroomNumber,
		&i.EnSuite,
		&i.Master,
		&i.PropertyUnitID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  phone
) VALUES (
  $1
)
RETURNING id, email, first_name, next_renewal, last_name, phone, onboarding, is_landlord, created_at, updated_at
`

func (q *Queries) CreateUser(ctx context.Context, phone string) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.NextRenewal,
		&i.LastName,
		&i.Phone,
		&i.Onboarding,
		&i.IsLandlord,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findByEmail = `-- name: FindByEmail :one
SELECT id, email, first_name, next_renewal, last_name, phone, onboarding, is_landlord, created_at, updated_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) FindByEmail(ctx context.Context, email sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, findByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.NextRenewal,
		&i.LastName,
		&i.Phone,
		&i.Onboarding,
		&i.IsLandlord,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserByPhone = `-- name: FindUserByPhone :one
SELECT id, email, first_name, next_renewal, last_name, phone, onboarding, is_landlord, created_at, updated_at FROM users
WHERE phone = $1
`

func (q *Queries) FindUserByPhone(ctx context.Context, phone string) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByPhone, phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.NextRenewal,
		&i.LastName,
		&i.Phone,
		&i.Onboarding,
		&i.IsLandlord,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProperty = `-- name: GetProperty :one
SELECT id, name, location, type, created_at, updated_at, created_by, caretaker_id FROM properties
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProperty(ctx context.Context, id int64) (Property, error) {
	row := q.db.QueryRowContext(ctx, getProperty, id)
	var i Property
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Location,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.CaretakerID,
	)
	return i, err
}

const getPropertyUnits = `-- name: GetPropertyUnits :many
SELECT id, name, type, state, price, bathrooms, created_at, updated_at, property_id FROM property_units
WHERE property_id = $1
`

func (q *Queries) GetPropertyUnits(ctx context.Context, propertyID sql.NullInt64) ([]PropertyUnit, error) {
	rows, err := q.db.QueryContext(ctx, getPropertyUnits, propertyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PropertyUnit
	for rows.Next() {
		var i PropertyUnit
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.State,
			&i.Price,
			&i.Bathrooms,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PropertyID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUnitBedrooms = `-- name: GetUnitBedrooms :many
SELECT id, bedroom_number, en_suite, master, property_unit_id, created_at, updated_at FROM bedrooms
WHERE property_unit_id = $1
`

func (q *Queries) GetUnitBedrooms(ctx context.Context, propertyUnitID int64) ([]Bedroom, error) {
	rows, err := q.db.QueryContext(ctx, getUnitBedrooms, propertyUnitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Bedroom
	for rows.Next() {
		var i Bedroom
		if err := rows.Scan(
			&i.ID,
			&i.BedroomNumber,
			&i.EnSuite,
			&i.Master,
			&i.PropertyUnitID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUnitTenancy = `-- name: GetUnitTenancy :many
SELECT id, start_date, end_date, created_at, updated_at, property_unit_id, user_id FROM tenants
WHERE property_unit_id = $1
`

func (q *Queries) GetUnitTenancy(ctx context.Context, propertyUnitID sql.NullInt64) ([]Tenant, error) {
	rows, err := q.db.QueryContext(ctx, getUnitTenancy, propertyUnitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tenant
	for rows.Next() {
		var i Tenant
		if err := rows.Scan(
			&i.ID,
			&i.StartDate,
			&i.EndDate,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PropertyUnitID,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, email, first_name, next_renewal, last_name, phone, onboarding, is_landlord, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.NextRenewal,
		&i.LastName,
		&i.Phone,
		&i.Onboarding,
		&i.IsLandlord,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const mailingExists = `-- name: MailingExists :one
SELECT EXISTS(
  SELECT id, email, created_at, updated_at FROM mailings
  WHERE email = $1
)
`

func (q *Queries) MailingExists(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, mailingExists, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const occupiedUnitsCount = `-- name: OccupiedUnitsCount :one
SELECT COUNT(*) FROM property_units
WHERE property_id = $1 AND state = 'occupied'
`

func (q *Queries) OccupiedUnitsCount(ctx context.Context, propertyID sql.NullInt64) (int64, error) {
	row := q.db.QueryRowContext(ctx, occupiedUnitsCount, propertyID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const onboardUser = `-- name: OnboardUser :one
UPDATE users
SET onboarding = $1
WHERE email = $2
RETURNING id, email, first_name, next_renewal, last_name, phone, onboarding, is_landlord, created_at, updated_at
`

type OnboardUserParams struct {
	Onboarding sql.NullBool   `json:"onboarding"`
	Email      sql.NullString `json:"email"`
}

func (q *Queries) OnboardUser(ctx context.Context, arg OnboardUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, onboardUser, arg.Onboarding, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.NextRenewal,
		&i.LastName,
		&i.Phone,
		&i.Onboarding,
		&i.IsLandlord,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const propertiesCreatedBy = `-- name: PropertiesCreatedBy :many
SELECT id, name, location, type, created_at, updated_at, created_by, caretaker_id FROM properties
WHERE created_by = $1
`

func (q *Queries) PropertiesCreatedBy(ctx context.Context, createdBy sql.NullInt64) ([]Property, error) {
	rows, err := q.db.QueryContext(ctx, propertiesCreatedBy, createdBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Property
	for rows.Next() {
		var i Property
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Location,
			&i.Type,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.CaretakerID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const propertyUnitsCount = `-- name: PropertyUnitsCount :one
SELECT COUNT(*) FROM property_units
WHERE property_id = $1
`

func (q *Queries) PropertyUnitsCount(ctx context.Context, propertyID sql.NullInt64) (int64, error) {
	row := q.db.QueryRowContext(ctx, propertyUnitsCount, propertyID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const saveMail = `-- name: SaveMail :one
INSERT INTO mailings (
  email
) VALUES (
  $1
)
RETURNING id, email, created_at, updated_at
`

func (q *Queries) SaveMail(ctx context.Context, email string) (Mailing, error) {
	row := q.db.QueryRowContext(ctx, saveMail, email)
	var i Mailing
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const unitAmenityCount = `-- name: UnitAmenityCount :one
SELECT COUNT(*) from amenities
WHERE property_unit_id = $1
`

func (q *Queries) UnitAmenityCount(ctx context.Context, propertyUnitID sql.NullInt64) (int64, error) {
	row := q.db.QueryRowContext(ctx, unitAmenityCount, propertyUnitID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateInvoiceForMpesa = `-- name: UpdateInvoiceForMpesa :one
UPDATE invoices
SET channel = $1, status = $2, amount = $3, currency = $4, bank = $5, auth_code = $6, country_code = $7, fees = $8, created_at = $9, updated_at = $10
WHERE reference = $11
RETURNING id, msid, channel, currency, bank, auth_code, country_code, fees, amount, phone, status, reference, created_at, updated_at
`

type UpdateInvoiceForMpesaParams struct {
	Channel     sql.NullString `json:"channel"`
	Status      interface{}    `json:"status"`
	Amount      sql.NullString `json:"amount"`
	Currency    sql.NullString `json:"currency"`
	Bank        sql.NullString `json:"bank"`
	AuthCode    sql.NullString `json:"auth_code"`
	CountryCode sql.NullString `json:"country_code"`
	Fees        sql.NullString `json:"fees"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Reference   sql.NullString `json:"reference"`
}

func (q *Queries) UpdateInvoiceForMpesa(ctx context.Context, arg UpdateInvoiceForMpesaParams) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, updateInvoiceForMpesa,
		arg.Channel,
		arg.Status,
		arg.Amount,
		arg.Currency,
		arg.Bank,
		arg.AuthCode,
		arg.CountryCode,
		arg.Fees,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Reference,
	)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.Msid,
		&i.Channel,
		&i.Currency,
		&i.Bank,
		&i.AuthCode,
		&i.CountryCode,
		&i.Fees,
		&i.Amount,
		&i.Phone,
		&i.Status,
		&i.Reference,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateLandlord = `-- name: UpdateLandlord :one
UPDATE users
SET is_landlord = $1, next_renewal = $2
WHERE phone = $3
RETURNING id, email, first_name, next_renewal, last_name, phone, onboarding, is_landlord, created_at, updated_at
`

type UpdateLandlordParams struct {
	IsLandlord  sql.NullBool `json:"is_landlord"`
	NextRenewal time.Time    `json:"next_renewal"`
	Phone       string       `json:"phone"`
}

func (q *Queries) UpdateLandlord(ctx context.Context, arg UpdateLandlordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateLandlord, arg.IsLandlord, arg.NextRenewal, arg.Phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.NextRenewal,
		&i.LastName,
		&i.Phone,
		&i.Onboarding,
		&i.IsLandlord,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const vacantUnitsCount = `-- name: VacantUnitsCount :one
SELECT COUNT(*) FROM property_units
WHERE property_id = $1 AND state = 'vacant'
`

func (q *Queries) VacantUnitsCount(ctx context.Context, propertyID sql.NullInt64) (int64, error) {
	row := q.db.QueryRowContext(ctx, vacantUnitsCount, propertyID)
	var count int64
	err := row.Scan(&count)
	return count, err
}
