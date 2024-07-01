package services

import (
	"context"
	pb "gardenManagement/genproto/GardenManagementSevice/gardenManagementService"
	"gardenManagement/storage/postgres"

	"github.com/jmoiron/sqlx"
)

type gardenManagementRepo struct {
	db postgres.GardenRepo
	pb.UnimplementedGardenManagementServiceServer
}

func NewUserManagementRepo(db *sqlx.DB) *gardenManagementRepo {
	return &gardenManagementRepo{db: *postgres.NewGardenRepo(db)}
}

// 1
func (g *gardenManagementRepo) CreateGarden(ctx context.Context, in *pb.GardenRequest) (*pb.GardenResponse, error) {
	res, err := g.db.CreateGarden(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 2
func (g *gardenManagementRepo) GetGardenByID(ctx context.Context, in *pb.IdRequest) (*pb.GardenResponse, error) {
	res, err := g.db.GetGardenByID(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 3
func (g *gardenManagementRepo) UpdateGardenByID(ctx context.Context, in *pb.UpdateGardenRequest) (*pb.GardenResponse, error) {
	res, err := g.db.UpdateGardenByID(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 4
func (g *gardenManagementRepo) DeleteGardenByID(ctx context.Context, in *pb.IdRequest) (*pb.DateResponse, error) {
	res, err := g.db.DeleteGardenByID(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 5
func (g *gardenManagementRepo) GetGardensByUserID(ctx context.Context, in *pb.IdRequest) (*pb.Gardens, error) {
	res, err := g.db.GetGardensByUserID(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 6
func (g *gardenManagementRepo) CreatePlantByGardenID(ctx context.Context, in *pb.PlantRequest) (*pb.PlantResponse, error) {
	res, err := g.db.CreatePlantByGardenID(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 7
func (g *gardenManagementRepo) GetPlantsByGardenID(ctx context.Context, in *pb.IdRequest) (*pb.Plants, error) {
	res, err := g.db.GetPlantsByGardenID(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 8
func (g *gardenManagementRepo) UpdatePlantByPlantsID(ctx context.Context, in *pb.PlantRequest) (*pb.PlantResponse, error) {
	res, err := g.db.UpdatePlantByPlantsID(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 9
func (g *gardenManagementRepo) DeletePlantByPlantsID(ctx context.Context, in *pb.IdRequest) (*pb.DateResponse, error) {
	res, err := g.db.DeletePlantByPlantsID(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 10
func (g *gardenManagementRepo) CreateCareLogByPlantID(ctx context.Context, in *pb.CareLogs) (*pb.CareLogsResponse, error) {
	res, err := g.db.CreateCareLogByPlantID(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 11
func (g *gardenManagementRepo) GetCareLogsByPlantID(ctx context.Context, in *pb.IdRequest) (*pb.CareLogsByPlantID, error) {
	res, err := g.db.GetCareLogsByPlantID(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}
