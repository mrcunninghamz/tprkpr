package data

import (
	"github.com/mrcunninghamz/tprkpr/platform/data/models"
	"github.com/mrcunninghamz/tprkpr/platform/data/seeds"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DataContext struct {
	*gorm.DB
}

func Connect() (*DataContext, error) {
	connectionString := os.Getenv("POSTGRESSQL_CONNECTION")
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {

		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.Payday{}, &models.Bill{}, &models.Worksheet{}, &models.WorksheetItem{})
	if err != nil {

		return nil, err
	}

	RunCustomSql(db)

	for _, seed := range seeds.All() {
		if err := seed.Run(db); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	return &DataContext{
		DB: db,
	}, nil
}

func RunCustomSql(db *gorm.DB) {
	db.Exec(`
		CREATE OR REPLACE FUNCTION update_worksheet() RETURNS TRIGGER AS $update_worksheet$
			BEGIN
				IF (TG_OP = 'DELETE') THEN
					UPDATE public.worksheets
					SET total = COALESCE((select sum(amount) from public.worksheet_items where worksheet_id = OLD.worksheet_id), 0),
					total_paid = COALESCE((select sum(amount) from public.worksheet_items where worksheet_id = OLD.worksheet_id and is_paid = true),0)
					WHERE id = OLD.worksheet_id;
					RETURN NULL;
				END IF;
		
				UPDATE public.worksheets
				SET total = COALESCE((select sum(amount) from public.worksheet_items where worksheet_id = NEW.worksheet_id),0),
				total_paid = COALESCE((select sum(amount) from public.worksheet_items where worksheet_id = NEW.worksheet_id and is_paid = true),0)
				WHERE id = NEW.worksheet_id;
		
				RETURN NULL;
			END;
		$update_worksheet$ LANGUAGE plpgsql;
		
		CREATE OR REPLACE TRIGGER update_worksheet
		AFTER INSERT OR UPDATE OR DELETE ON public.worksheet_items
    	FOR EACH ROW EXECUTE FUNCTION update_worksheet();
	`)
}
