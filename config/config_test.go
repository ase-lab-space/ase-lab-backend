package config_test

import (
	"testing"

	"github.com/ase-lab-space/ase-lab-backend/config"
)

func TestNew(t *testing.T) {
	cfg, err := config.New(".env.test")

	if err != nil {
		t.Fatal(err)
	}

	if cfg.PORT != 1234 {
		t.Fatalf("wrong PORT: expected %d, got %d", 1234, cfg.PORT)
	}
}
