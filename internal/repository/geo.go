package repository

import "context"

type GeoRepository interface {
	CheckLocation(ctx context.Context, latitude float64, longitude float64) (bool, error)
}
