terraform {
  required_version = ">= 1.3"

  required_providers {
    google = ">= 3.3"
  }
}

provider "google" {
  project = var.project_id
}

resource "google_cloud_run_service" "scheduler-service" {
  name     = "dnd-scheduler-bot"
  location = "europe-west2"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
