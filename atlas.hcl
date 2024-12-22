env "local" {
  src="file://sql"
  url = "sqlite://local.db?_fk=1"
  dev = "sqlite://dev?mode=memory&_fk=1"
}
