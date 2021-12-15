module "books_database" {
  source = "../../modules/postgresql-database"

  db_name        = "books"
  admin_password = "testpassword_admin"
  app_password   = "testpassword_app"

  providers = {
    postgresql = postgresql.docker
  }
}

module "books_grants" {
  source = "../../modules/postgresql-grants"

  schema     = "books"
  db_name    = module.books_database.db_name
  admin_role = module.books_database.admin_role
  app_role   = module.books_database.app_role

  providers = {
    postgresql = postgresql.docker
  }
}
