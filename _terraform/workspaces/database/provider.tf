provider "postgresql" {
  alias     = "docker"
  host      = "127.0.0.1"
  port      = 22411
  database  = "superuser_db"
  username  = "superuser"
  password  = "testpassword_superuser"
  sslmode   = "disable"
  superuser = false
}
