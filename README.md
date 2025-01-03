# Финальный проект 1 семестра

REST API сервис для загрузки и выгрузки данных о ценах.

## Требования к системе

- go 1.23.3
- postgres 13.3+ (предполагается, что база данных уже создана и доступна по адресу `localhost:5432`)

## Установка и запуск

Как установить и запустить приложение локально?

Сценарий 1 (с нуля):

```
chmod +x scripts/prepare.sh
./scripts/start.sh
```

Сценарий 2 (с существующей инфраструктурой):

Компиляция

```
chmod +x scripts/prepare.sh
./scripts/prepare.sh
```

Запуск

```
chmod +x scripts/run.sh
./scripts/run.sh
```


## Тестирование

Директория `sample_data` - это пример директории, которая является разархивированной версией файла `sample_data.zip`

Тестировние
```
chmod +x scripts/tests.sh
for idx in {1..3}; do ./scripts/tests.sh $idx; done
```

## Контакт

К кому можно обращаться в случае вопросов?

`petrkrotov2001@mail.ru` - email

`peter40127` - TG
