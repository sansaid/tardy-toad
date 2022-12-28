resource "fly_app" "tardy-toad" {
  name = "tardy-toad"
  org = local.fly_io_org
}