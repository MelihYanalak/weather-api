package repository

type IGeoRepository interface {
	CheckLocation(latitude float64, longitude float64) (bool, error)
	Initialize(filePath string) error
}
