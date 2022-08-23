package config

import (
	"testing"
)

func TestNew(t *testing.T) {
	cfg, err := New(".env.test")

	if err != nil {
		t.Fatal(err)
	}

	if cfg.PORT != 1234 {
		t.Errorf("wrong PORT: expected %d, got %d", 1234, cfg.PORT)
	}

	if cfg.GIN_MODE != "test" {
		t.Errorf("wrong GIN_MODE: expected %s, got %s", "test", cfg.GIN_MODE)
	}

	if cfg.CONTACT_FORM_NOTIFICATOR_ACCESS_TOKEN != "accesstoken" {
		t.Fatalf("wrong CONTACT_FORM_NOTIFICATOR_ACCESS_TOKEN: expected accesstoken, got %s", cfg.CONTACT_FORM_NOTIFICATOR_ACCESS_TOKEN)
	}
}
