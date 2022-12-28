data "google_iam_policy" "tardy-toad-bot" {
  binding {
    role = "roles/iam.serviceAccountUser"

    members = local.admin_users
  }
}

resource "google_service_account" "tardy-toad-bot" {
  account_id   = "tardy-toad-bot"
  display_name = "tardy-toad-bot"

  description = "Service account for the tardy-toad Discord bot itself"
}

resource "google_service_account_iam_policy" "tardy-toad-bot" {
  service_account_id = google_service_account.tardy-toad-bot.name
  policy_data        = data.google_iam_policy.admin.policy_data
}