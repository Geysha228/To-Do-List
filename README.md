## Используемые пакеты
- go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest (для SQL-миграций)
Для создания новой миграции, используем следующую команду:
migrate create -ext sql -dir migrations -seq <имя_миграции>
