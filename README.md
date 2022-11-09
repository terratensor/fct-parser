## Parser страниц вопросов ФКТ-АЛТАЙ
Утилита командной строки для синтаксического анализа страниц списка вопросов к рубрике «Вопрос-ответ» и сохранением в файл в формате csv или json

Для получения файла надо передать url страницы <br>
https://xn----8sba0bbi0cdm.xn--p1ai/qa/question/:topic-id

Опции командной строки
----------------------

```
-f, --file string[="topic"]   write to file name (default "topic")
-j, --json                    вывод в формате json (по умолчанию "csv")
-i, --json-indent             форматированный вывод json с отступами и переносами строк
```
### Примеры

Получение данных темы форума по-заданному url в формате csv с сохранением в файл topic.csv

``` 
./fct-parser "https://xn----8sba0bbi0cdm.xn--p1ai/qa/question/view-44707"
```

Получение данных темы в формате json без отступов и переносов строк с сохранением в файл topic.json
``` 
./fct-parser -j "https://xn----8sba0bbi0cdm.xn--p1ai/qa/question/view-44707"
```

Получение данных темы в формате json с отступами и переносами строк с сохранением в файл topic.json
``` 
./fct-parser -i "https://xn----8sba0bbi0cdm.xn--p1ai/qa/question/view-44707"
```

Получение данных темы в формате по умолчанию и сохранением в заданный файл <br>
``` 
./fct-parser -f=topic-44707.csv "https://xn----8sba0bbi0cdm.xn--p1ai/qa/question/view-44707"
```
При использовании опции `-f, --filename` расширение файла при необходимости указывайте самостоятельно

Получение данных нескольких тем по url адресам указанным в командной строке и сохранение каждой темы в отдельный файл 
```
./fct-parser -i "https://xn----8sba0bbi0cdm.xn--p1ai/qa/question/view-44707" "https://xn----8sba0bbi0cdm.xn--p1ai/qa/question/view-41574"
```
Пример вывода при успешном выполнении команды и получении данных по всем заданным ссылкам
```
2022/11/09 15:59:11 file topic-1.csv was successful writing
2022/11/09 15:59:13 file topic-2.csv was successful writing
```
