package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Crop struct {
	UID         uuid.UUID
	BatchID     string
	InitialArea Area
	CurrentArea Area
	Type        CropType
	PlantType   PlantType
	Container   CropContainer
	CreatedDate time.Time
}

// CropType defines type of crop
type CropType interface {
	Code() string
}

// Seeding implements CropType
type Seeding struct{}

func (s Seeding) Code() string { return "seeding" }

// Growing implements CropType
type Growing struct{}

func (g Growing) Code() string { return "growing" }

// CropContainer defines the container of a crop
type CropContainer struct {
	Quantity int
	Type     CropContainerType
}

// CropContainerType defines the type of a container
type CropContainerType interface {
	Code() string
}

// Tray implements CropContainerType
type Tray struct {
	Cell int
}

func (t Tray) Code() string { return "tray" }

// Pot implements CropContainerType
type Pot struct{}

func (p Pot) Code() string { return "pot" }

func CreateCropBatch(area Area) (Crop, error) {
	if area.UID == (uuid.UUID{}) {
		return Crop{}, CropError{Code: CropErrorInvalidArea}
	}

	return Crop{
		InitialArea: area,
		CurrentArea: area,
		CreatedDate: time.Now(),
	}, nil
}

func (c *Crop) ChangeCropType(cropType CropType) error {
	err := validateCropType(cropType)
	if err != nil {
		return err
	}

	c.Type = cropType

	return nil
}

func (c *Crop) ChangePlantType(plantType PlantType) error {
	err := validatePlantType(plantType)
	if err != nil {
		return err
	}

	batchID, err := generateBatchID(plantType)

	c.PlantType = plantType
	c.BatchID = batchID

	return nil
}

func (c *Crop) ChangeContainer(container CropContainer) error {
	err := validateCropContainer(container)
	if err != nil {
		return err
	}

	return nil
}

func generateBatchID(plantType PlantType) (string, error) {
	batchID := "DUMMY-BATCH-ID"

	return batchID, nil
}

func validateCropType(cropType CropType) error {
	switch cropType.(type) {
	case Seeding:
	case Growing:
	default:
		return CropError{Code: CropErrorInvalidCropType}
	}

	return nil
}

func validateCropContainer(container CropContainer) error {
	switch container.Type.(type) {
	case Tray:
	case Pot:
	default:
		return CropError{Code: CropContainerErrorInvalidType}
	}

	return nil
}
