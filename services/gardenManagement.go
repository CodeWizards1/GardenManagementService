package services

import (
	"context"
	pb "gardenManagement/genproto/GardenManagementService"
	user "gardenManagement/genproto/UserManagementService"
	"gardenManagement/storage/postgres"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(port string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type gardenManagementRepo struct {
	db         postgres.GardenRepo
	userClient user.UserManagementServiceClient
	pb.UnimplementedGardenManagementServiceServer
}

func NewGardenManagementRepo(db *sqlx.DB, userServiceAddress string) (*gardenManagementRepo, error) {
	conn, err := Connect(userServiceAddress)
	if err != nil {
		return nil, err
	}
	userClient := user.NewUserManagementServiceClient(conn)
	return &gardenManagementRepo{db: *postgres.NewGardenRepo(db), userClient: userClient}, nil
}

func (g *gardenManagementRepo) DoesGardenExist(ctx context.Context, in *pb.IdRequest) (*pb.DoesGardenExistResponse, error) {
	res, err := g.db.DoesGardenExist(ctx, in)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 1
func (g *gardenManagementRepo) CreateGarden(ctx context.Context, in *pb.GardenRequest) (*pb.GardenResponse, error) {
	_, err := g.userClient.DoesUserExists(ctx, &user.IdUserRequest{UserId: in.UserId})
	if err != nil {
		return nil, err
	}

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
