## Parser страниц вопросов ФКТ-АЛТАЙ
Утилита командной строки для синтаксического анализа страниц списка вопросов к рубрике «Вопрос-ответ» и сохранением в файл в формате csv или json.

Для получения файла с последним обсуждением событий СВОДД в формате csv просто запустите утилиту<br>
```
./fct-parser 
```

По умолчанию для удобства чтения в текстовых редакторах из текста файла вырезаются все html теги, но если необходимо получить файл с html тегами, то при вызове утилиты укажите опцию `-h`
```
./fct-parser -h 
```

Для получения всех файлов с обсуждениями событий с начала СВОДД в формате csv запустите утилиту с флагом `-a`<br>
После запуска утилита последовательно сделает запросы по url адресам из файла конфигурации config.json и сохранит результаты в файлы csv
```
./fct-parser -a
```

Для получения списка ссылок с обсуждениями событий с начала СВОДД
```
./fct-parser -l
```

Для получения ссылки текущего активного обсуждения событий с начала СВОДД
```
./fct-parser -с
```
```
https://фкт-алтай.рф/qa/question/view-41574
```

Для получения файла в формате csv с любым вопросом надо передать url страницы, например
```
./fct-parser https://xn----8sba0bbi0cdm.xn--p1ai/qa/question/view-35030
```
```
2022/11/20 21:27:37 fetching config file https://raw.githubusercontent.com/audetv/fct-parser/main/config.json
2022/11/20 21:27:38 parse https://фкт-алтай.рф/qa/question/view-35030
2022/11/20 21:27:38 всего комментариев 101
2022/11/20 21:27:38 file ./qa-question-view-35030.csv was successful writing
```
Опции командной строки
----------------------

```
  -a, --all                     сохранение всего списка обсуждений событий с начала СВОДД в отдельные файлы
  -c, --current                 вывод в консоль адреса ссылки текущего активного обсуждения событий с начала СВОДД  
  -h, --html-tags               вывод с сохранение с html тегов
  -j, --json                    вывод в формате json (по умолчанию "csv")
  -i, --json-indent             форматированный вывод json с отступами и переносами строк
  -l, --list                    вывод в консоль списка адресов страниц с обсуждениями событий с начала СВОДД
      --help                    вывод справки по командам утилиты 
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

Получение данных нескольких тем по url адресам указанным в командной строке и сохранение каждой темы в отдельный файл 
```
./fct-parser -i "https://xn----8sba0bbi0cdm.xn--p1ai/qa/question/view-44707" "https://xn----8sba0bbi0cdm.xn--p1ai/qa/question/view-41574"
```
Пример вывода при успешном выполнении команды и получении данных по всем заданным ссылкам
```
2022/11/10 16:40:56 parse https://xn----8sba0bbi0cdm.xn--p1ai/qa/question/view-44707
2022/11/10 16:40:56 file qa-question-view-44707.json was successful writing
2022/11/10 16:40:58 parse https://xn----8sba0bbi0cdm.xn--p1ai/qa/question/view-41574
2022/11/10 16:40:58 file qa-question-view-41574.json was successful writing
2022/11/10 16:40:58 все запросы выполнены
```

Для получения всего списка адресов вопросов с обсуждениями
```
./fct-parser -l
```

```
https://фкт-алтай.рф/qa/question/view-44538
https://фкт-алтай.рф/qa/question/view-44612
https://фкт-алтай.рф/qa/question/view-44707
https://фкт-алтай.рф/qa/question/view-44757
https://фкт-алтай.рф/qa/question/view-44883
https://фкт-алтай.рф/qa/question/view-44962
https://фкт-алтай.рф/qa/question/view-45044
https://фкт-алтай.рф/qa/question/view-35650
https://фкт-алтай.рф/qa/question/view-35298
https://фкт-алтай.рф/qa/question/view-4604
https://фкт-алтай.рф/qa/question/view-7533
https://фкт-алтай.рф/qa/question/view-23174
https://фкт-алтай.рф/qa/question/view-37945
https://фкт-алтай.рф/qa/question/view-12422
https://фкт-алтай.рф/qa/question/view-25867
https://фкт-алтай.рф/qa/question/view-14365
https://фкт-алтай.рф/qa/question/view-34312
https://фкт-алтай.рф/qa/question/view-37694
https://фкт-алтай.рф/qa/question/view-7279
https://фкт-алтай.рф/qa/question/view-2656
https://фкт-алтай.рф/qa/question/view-12734
https://фкт-алтай.рф/qa/question/view-3893
https://фкт-алтай.рф/qa/question/view-4910
https://фкт-алтай.рф/qa/question/view-3467
https://фкт-алтай.рф/qa/question/view-21294
https://фкт-алтай.рф/qa/question/view-41574
```