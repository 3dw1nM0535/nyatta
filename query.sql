-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  phone
) VALUES (
  $1
)
RETURNING *;

-- name: CreateProperty :one
INSERT INTO properties (
  name, type, created_by, location
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetProperty :one
SELECT * FROM properties
WHERE id = $1 LIMIT 1;

-- name: FindByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: PropertiesCreatedBy :many
SELECT * FROM properties
WHERE created_by = $1;

-- name: CreateAmenity :one
INSERT INTO amenities (
  name, category, property_unit_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: CreateTenant :one
INSERT INTO tenants (
  start_date, property_unit_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: CreatePropertyUnit :one
INSERT INTO property_units (
  property_id, bathrooms, name, type, price, location
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: CreateUnitBedroom :one
INSERT INTO bedrooms (
  property_unit_id, bedroom_number, en_suite, master
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUnitBedrooms :many
SELECT * FROM bedrooms
WHERE property_unit_id = $1;

-- name: GetUnitTenancy :many
SELECT * FROM tenants
WHERE property_unit_id = $1;

-- name: GetPropertyUnits :many
SELECT * FROM property_units
WHERE property_id = $1;

-- name: FindUserByPhone :one
SELECT * FROM users
WHERE phone = $1;

-- name: CreateCaretaker :one
INSERT INTO caretakers (
  first_name, last_name, phone
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: CreateShootSchedule :one
INSERT INTO shoots (
  shoot_date, property_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: OnboardUser :one
UPDATE users
SET onboarding = $1
WHERE email = $2
RETURNING *;

-- name: SaveMail :one
INSERT INTO mailings (
  email
) VALUES (
  $1
)
RETURNING *;

-- name: MailingExists :one
SELECT EXISTS(
  SELECT * FROM mailings
  WHERE email = $1
);

-- name: PropertyUnitsCount :one
SELECT COUNT(*) FROM property_units
WHERE property_id = $1;

-- name: OccupiedUnitsCount :one
SELECT COUNT(*) FROM property_units
WHERE property_id = $1 AND state = 'occupied';

-- name: VacantUnitsCount :one
SELECT COUNT(*) FROM property_units
WHERE property_id = $1 AND state = 'vacant';

-- name: UnitAmenityCount :one
SELECT COUNT(*) from amenities
WHERE property_unit_id = $1;
