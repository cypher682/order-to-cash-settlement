package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("SERVICE_NAME", "")
	t.Setenv("LOG_LEVEL", "")
	t.Setenv("DATABASE_URL", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.Port != 8080 {
		t.Fatalf("Port = %d, want 8080", cfg.Port)
	}
	if cfg.ServiceName != "o2c" {
		t.Fatalf("ServiceName = %q, want o2c", cfg.ServiceName)
	}
}

func TestLoadRejectsInvalidPort(t *testing.T) {
	t.Setenv("PORT", "70000")

	if _, err := Load(); err == nil {
		t.Fatal("Load() = nil error, want invalid PORT error")
	}
}
