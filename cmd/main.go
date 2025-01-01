package main
import (
	"project_sem/internal/app"
	"project_sem/internal/config"
	"log"
)
func main() {
	cfg, err := config.Load("config.yaml", "yaml")
	if err != nil {
		log.Fatalf("failed to load configuration with error %s", err)
	}
	instance := app.New(cfg)
	instance.Run()
}