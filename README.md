# Финальный проект 1 семестра

REST API сервис для загрузки и выгрузки данных о ценах.

## Требования к системе

- go 1.23.3
- postgres 13.3+ (предполагается, что база данных уже создана и доступна по адресу `localhost:5432`)

## Установка и запуск

При наличии инфраструктуры
```
//Компиляция

chmod +x scripts/prepare.sh
./scripts/prepare.sh

//Запуск

chmod +x scripts/run.sh
./scripts/run.sh
```

или

```

chmod +x scripts/start.sh
./scripts/start.sh

```


## Тестирование

Директория `sample_data` - это пример директории, которая является разархивированной версией файла `sample_data.zip`

```

chmod +x scripts/tests.sh
for idx in {1..3}; do ./scripts/tests.sh $idx; done

```

## Контакт

К кому можно обращаться в случае вопросов?

tg @peter40127 или по контактам, указанным в Github

https://github.com/KrotovPetr

Задание на уровень 2