package postgres

import (
	"context"
	pb "gardenManagement/genproto/GardenManagementService"
	"time"

	"github.com/jmoiron/sqlx"
)

type GardenRepo struct {
	db *sqlx.DB
}

func NewGardenRepo(db *sqlx.DB) *GardenRepo {
	return &GardenRepo{db: db}
}

func (g *GardenRepo) DoesGardenExist(ctx context.Context, in *pb.IdRequest) (*pb.DoesGardenExistResponse, error) {
	gardenID := in.GetId()

	query := "SELECT EXISTS (SELECT 1 FROM gardens WHERE id = $1)"

	var exists bool
	err := g.db.QueryRowContext(ctx, query, gardenID).Scan(&exists)
	if err != nil {
		return &pb.DoesGardenExistResponse{Exists: false}, err
	}

	return &pb.DoesGardenExistResponse{Exists: true}, nil
}

// 1
func (g *GardenRepo) CreateGarden(ctx context.Context, in *pb.GardenRequest) (*pb.GardenResponse, error) {
	query := `
		INSERT INTO gardens (
			user_id,
			name,
			type,
			area_sqm)
		VALUES 
			($1, $2, $3, $4)
		RETURNING 
			id,
			user_id,
			name,
			type,
			area_sqm,
			created_at,
			updated_at
	`
	var (
		gardenRespone pb.GardenResponse
		created_at    time.Time
		updated_at    time.Time
	)

	row := g.db.QueryRowContext(ctx, query, in.UserId, in.Name, in.Type, in.Area)
	if err := row.Scan(
		&gardenRespone.Id,
		&gardenRespone.UserId,
		&gardenRespone.Name,
		&gardenRespone.Type,
		&gardenRespone.Area,
		&created_at,
		&updated_at,
	); err != nil {
		return nil, err
	}

	gardenRespone.CreatedAt = created_at.Format("2006-01-02 15:04:05")
	gardenRespone.UpdatedAt = updated_at.Format("2006-01-02 15:04:05")

	return &gardenRespone, nil
}

// 2
func (g *GardenRepo) GetGardenByID(ctx context.Context, in *pb.IdRequest) (*pb.GardenResponse, error) {
	query := `
		SELECT 
			id,
			user_id,
			name,
			type,
			area_sqm,
			created_at,
			updated_at
		FROM gardens
		WHERE id = $1 
	`

	var (
		gardenRespone pb.GardenResponse
		created_at    time.Time
		updated_at    time.Time
	)

	row := g.db.QueryRowContext(ctx, query, in.Id)
	if err := row.Scan(
		&gardenRespone.Id,
		&gardenRespone.UserId,
		&gardenRespone.Name,
		&gardenRespone.Type,
		&gardenRespone.Area,
		&created_at,
		&updated_at,
	); err != nil {
		return nil, err
	}

	gardenRespone.CreatedAt = created_at.Format("2006-01-02 15:04:05")
	gardenRespone.UpdatedAt = updated_at.Format("2006-01-02 15:04:05")

	return &gardenRespone, nil
}

// 3
func (g *GardenRepo) UpdateGardenByID(ctx context.Context, in *pb.UpdateGardenRequest) (*pb.GardenResponse, error) {
	query := `
		UPDATE 
			gardens
		SET
			user_id = $1,
			name = $2,
			type = $3,
			area_sqm = $4,
			updated_at = now()
		WHERE 
			id = $5 AND deleted_at IS NULL
		RETURNING
			id,
			user_id,
			name,
			type,
			area_sqm,
			created_at,
			updated_at
	`

	var (
		gardenRespone pb.GardenResponse
		created_at    time.Time
		updated_at    time.Time
	)

	row := g.db.QueryRowContext(ctx, query, in.UserId, in.Name, in.Type, in.Area, in.UserId)
	if err := row.Scan(
		&gardenRespone.Id,
		&gardenRespone.UserId,
		&gardenRespone.Name,
		&gardenRespone.Type,
		&gardenRespone.Area,
		&created_at,
		&updated_at,
	); err != nil {
		return nil, err
	}

	gardenRespone.CreatedAt = created_at.Format("2006-01-02 15:04:05")
	gardenRespone.UpdatedAt = updated_at.Format("2006-01-02 15:04:05")

	return &gardenRespone, nil
}

// 4
func (g *GardenRepo) DeleteGardenByID(ctx context.Context, in *pb.IdRequest) (*pb.DateResponse, error) {
	query := `
		UPDATE 
			gardens
		SET
			deleted_at = now()
		WHERE
			id = $1
		RETURNING
			deleted_at
	`

	var deleted_at time.Time

	row := g.db.QueryRowContext(ctx, query, in.Id)
	if err := row.Scan(&deleted_at); err != nil {
		return nil, err
	}

	return &pb.DateResponse{Message: deleted_at.Format("2006-01-02 15:04:05")}, nil
}

// 5
func (g *GardenRepo) GetGardensByUserID(ctx context.Context, in *pb.IdRequest) (*pb.Gardens, error) {
	query := `
		SELECT
			id,
			user_id,
			name,
			type,
			area_sqm,
			created_at,
			updated_at
		FROM 
			gardens
		WHERE 
		user_id = $1
		AND deleted_at IS NULL
	`

	rows, err := g.db.QueryContext(ctx, query, in.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plants []*pb.GardenResponse
	for rows.Next() {
		var (
			gardenRespone pb.GardenResponse
			created_at    time.Time
			updated_at    time.Time
		)
		err := rows.Scan(
			&gardenRespone.Id,
			&gardenRespone.UserId,
			&gardenRespone.Name,
			&gardenRespone.Type,
			&gardenRespone.Area,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		gardenRespone.CreatedAt = created_at.Format("2006-01-02 15:04:05")
		gardenRespone.UpdatedAt = updated_at.Format("2006-01-02 15:04:05")
		plants = append(plants, &gardenRespone)
	}

	return &pb.Gardens{Gardens: plants}, nil
}

// 6
func (g *GardenRepo) CreatePlantByGardenID(ctx context.Context, in *pb.PlantRequest) (*pb.PlantResponse, error) {
	query := `
		INSERT INTO plants (
			garden_id,
			species,
			quantity,
			planting_date,
			status)
		VALUES
			($1, $2, $3, now(), $4)
		RETURNING
			id,
			garden_id,
			species,
			quantity,
			planting_date,
			status,
			created_at,
			updated_at
	`

	var (
		plantRespone  pb.PlantResponse
		created_at    time.Time
		updated_at    time.Time
		planting_date time.Time
	)

	row := g.db.QueryRowContext(ctx, query, in.GardenId, in.Species, in.Quantity, in.Status)
	if err := row.Scan(
		&plantRespone.Id,
		&plantRespone.GardenId,
		&plantRespone.Species,
		&plantRespone.Quantity,
		&planting_date,
		&plantRespone.Status,
		&created_at,
		&updated_at,
	); err != nil {
		return nil, err
	}

	plantRespone.CreatedAt = created_at.Format("2006-01-02 15:04:05")
	plantRespone.UpdatedAt = updated_at.Format("2006-01-02 15:04:05")
	plantRespone.PlantingDate = planting_date.Format("2006-01-02 15:04:05")

	return &plantRespone, nil
}

// 7
func (g *GardenRepo) GetPlantsByGardenID(ctx context.Context, in *pb.IdRequest) (*pb.Plants, error) {
	query := `
		SELECT
			id,
			garden_id,
			species,
			quantity,
			planting_date,
			status,
			created_at,
			updated_at
		FROM
			plants
		WHERE
			garden_id = $1
	`

	rows, err := g.db.QueryContext(ctx, query, in.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		plants       []*pb.PlantResponse
		created_at   time.Time
		updated_at   time.Time
		plating_date time.Time
	)

	for rows.Next() {
		var plantRespone pb.PlantResponse
		err := rows.Scan(
			&plantRespone.Id,
			&plantRespone.GardenId,
			&plantRespone.Species,
			&plantRespone.Quantity,
			&plating_date,
			&plantRespone.Status,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		plantRespone.CreatedAt = created_at.Format("2006-01-02 15:04:05")
		plantRespone.UpdatedAt = updated_at.Format("2006-01-02 15:04:05")
		plantRespone.PlantingDate = plating_date.Format("2006-01-02 15:04:05")
		plants = append(plants, &plantRespone)
	}

	return &pb.Plants{Plants: plants}, nil
}

// 8
func (g *GardenRepo) UpdatePlantByPlantsID(ctx context.Context, in *pb.PlantRequest) (*pb.PlantResponse, error) {
	query := `
		UPDATE
			plants
		SET
			garden_id = $1,
			species = $2,
			quantity = $3,
			planting_date = now(),
			status = $5,
			updated_at = now()
		WHERE
			id = $6 AND deleted_at IS NULL
		RETURNING
			id,
			garden_id,
			species,
			quantity,
			planting_date,
			status,
			created_at,
			updated_at
	`

	var (
		plantRespone  pb.PlantResponse
		created_at    time.Time
		updated_at    time.Time
		planting_date time.Time
	)

	row := g.db.QueryRowContext(ctx, query, in.GardenId, in.Species, in.Quantity, in.Status)
	if err := row.Scan(
		&plantRespone.Id,
		&plantRespone.GardenId,
		&plantRespone.Species,
		&plantRespone.Quantity,
		&planting_date,
		&plantRespone.Status,
		&created_at,
		&updated_at,
	); err != nil {
		return nil, err
	}

	plantRespone.CreatedAt = created_at.Format("2006-01-02 15:04:05")
	plantRespone.UpdatedAt = updated_at.Format("2006-01-02 15:04:05")
	plantRespone.PlantingDate = planting_date.Format("2006-01-02 15:04:05")

	return &plantRespone, nil
}

// 9
func (g *GardenRepo) DeletePlantByPlantsID(ctx context.Context, in *pb.IdRequest) (*pb.DateResponse, error) {
	query := `
		UPDATE
			plants
		SET
			deleted_at = now()
		WHERE
		 	id = $1 and deleted_at IS NULL
		RETURNING
			deleted_at
	`

	var deleted_at time.Time

	row := g.db.QueryRowContext(ctx, query, in.Id)
	if err := row.Scan(
		&deleted_at,
	); err != nil {
		return nil, err
	}

	return &pb.DateResponse{Message: deleted_at.Format("2006-01-02 15:04:05")}, nil
}

// 10
func (g *GardenRepo) CreateCareLogByPlantID(ctx context.Context, in *pb.CareLogs) (*pb.CareLogsResponse, error) {
	query := `
		INSERT INTO care_logs (
			plant_id,
			action,
			notes,
			logged_at)	
		VALUES
			($1, $2, $3, now())
		RETURNING
			id,
			plant_id,
			action,
			notes,
			logged_at
	`

	var (
		careLogsRespone pb.CareLogsResponse
		logged_at       time.Time
	)

	row := g.db.QueryRowContext(ctx, query, in.PlantId, in.Action, in.Notes)
	if err := row.Scan(
		&careLogsRespone.Id,
		&careLogsRespone.PlantId,
		&careLogsRespone.Action,
		&careLogsRespone.Notes,
		&logged_at,
	); err != nil {
		return nil, err
	}

	careLogsRespone.LoggedAt = logged_at.Format("2006-01-02 15:04:05")

	return &careLogsRespone, nil
}

// 11
func (g *GardenRepo) GetCareLogsByPlantID(ctx context.Context, in *pb.IdRequest) (*pb.CareLogsByPlantID, error) {
	query := `
		SELECT
			id,
			plant_id,
			action,
			notes,
			logged_at
		FROM 	
			care_logs
		WHERE
			plant_id = $1
	`

	rows, err := g.db.QueryContext(ctx, query, in.Id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var careLogs []*pb.CareLogsResponse

	for rows.Next() {
		var careLogsRespone pb.CareLogsResponse
		var logged_at time.Time
		err := rows.Scan(
			&careLogsRespone.Id,
			&careLogsRespone.PlantId,
			&careLogsRespone.Action,
			&careLogsRespone.Notes,
			&logged_at,
		)
		if err != nil {
			return nil, err
		}
		careLogsRespone.LoggedAt = logged_at.Format("2006-01-02 15:04:05")
		careLogs = append(careLogs, &careLogsRespone)
	}

	return &pb.CareLogsByPlantID{CareLogs: careLogs}, nil
}
