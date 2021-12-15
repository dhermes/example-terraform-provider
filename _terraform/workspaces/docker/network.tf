resource "docker_network" "books" {
  name     = "dev-network-books"
  internal = true
}
