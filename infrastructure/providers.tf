terraform {
  cloud {
    organization = "sanyia"

    workspaces {
      name = "tardy-toad"
    }
  }

  required_providers {
    fly = {
      source = "fly-apps/fly"
      version = "~> 0.0.20"
    }

    google = {
      source = "hashicorp/google"
      version = "~> 4.47"
    }
  }
}