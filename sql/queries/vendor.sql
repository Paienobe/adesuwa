-- name: RegisterVendor :one
INSERT INTO vendor (id, name, email, password)
VALUES($1, $2, $3, $4)
RETURNING *;

-- name: FindVendorById :one
SELECT * FROM vendor WHERE id = $1; 

-- name: FindVendorByEmail :one
SELECT * FROM vendor WHERE email = $1;

-- name: UpdateVendorProfilePicture :one
UPDATE vendor SET profile_image = $1
WHERE id = $2
RETURNING profile_image;

-- name: UpdateVendorDescription :one
UPDATE vendor SET description = $1
WHERE id = $2
RETURNING description;

-- name: UpdateVendorBanner :one
UPDATE vendor SET banner_image = $1
WHERE id = $2
RETURNING banner_image;