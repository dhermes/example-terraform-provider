module "books" {
  source = "../../modules/postgresql-container"

  port           = "22411"
  container_name = "dev-postgres-books"
  network_name   = docker_network.books.name

  providers = {
    docker = docker.local
  }
}
