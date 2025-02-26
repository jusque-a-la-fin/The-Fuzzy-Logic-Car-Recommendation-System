# Веб-приложение, формирующее рекомендации по выбору автомобилей
Данный сервис помогает пользователю, непогруженному в мир автомобилей, найти машины, отсортированные по таким субъективным характеристикам, как **«Комфорт», «Экономичность», «Безопасность», «Динамика», «Управляемость»**. Для ранжирования автомобилей применяются алгоритм, использующий методы **«нечеткой логики»**. 
Веб-приложение построено по принципам «The Clean Architecture» by Robert C. Martin (Uncle Bob): https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html?clckid=b335004d <br><br>
Ниже необходимая документация:  
1) [Объяснение алгоритма](docs/algo.pdf)
2) [База знаний](docs/knowledge-base.pdf)

## Как запустить:
```bash
git clone https://github.com/jusque-a-la-fin/Car-Recommendation-System-with-Fuzzy-Logic.git && cd Car-Recommendation-System-with-Fuzzy-Logic && sudo docker compose up
```
В адресной строке браузера перейти по URL: http://localhost:8080/main

Для запуска линтеров выполнить в корневой директории проекта команду:
```
golangci-lint run
```
Инструкция по установке golangci-lint: https://golangci-lint.run/welcome/install/
