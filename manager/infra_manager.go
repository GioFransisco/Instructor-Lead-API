package manager

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/config"
)

type InfraManager interface {
	Connection() *sql.DB
}

type infraManager struct {
	cfg *config.Config
	db  *sql.DB
}

func (i *infraManager) openConnection() error {
	// password=%s , i.cfg.DbPass
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable", i.cfg.DbHost, i.cfg.DbUser, i.cfg.DbName, i.cfg.DbPort)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	i.db = db

	return nil
}

// Connection implements InfraManager.
func (i *infraManager) Connection() *sql.DB {
	return i.db
}

func NewInfraManager(cfg *config.Config) InfraManager {
	infra := &infraManager{cfg: cfg}

	if err := infra.openConnection(); err != nil {
		panic(err)
	}

	return infra
}
