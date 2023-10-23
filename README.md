# juniorTask

Для настройки программы необходимо применить миграции, находящиеся в папке /migration,
например утилитой goose
Прежде всего, нужно установить утилиту goose командой
go install github.com/pressly/goose/v3/cmd/goose@latest
Затем запустит команду в корне проекта для применения миграций (указывая свои данные для подключения к PostgreSQL)
goose -dir migration postgres "host=localhost user=postgres password=mysecretpassword dbname=postgres sslmode=disable" up