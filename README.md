### Parser сайта ФКТ-АЛТАЙ
Parser страниц списка вопросов рубрики вопрос-ответ https://xn----8sba0bbi0cdm.xn--p1ai/qa/question

Парсинг темы с сохранением в файл в формате json отформатированном
``` 
$ ./fct-parser -json -indent -file topic-44707.json "https://xn----8sba0b
bi0cdm.xn--p1ai/qa/question/view-44707"
2022/11/09 00:10:56 The file ./topic-44707.json was successeful writing
```
Парсинг темы с сохранением в файл в формате csv
``` 
$ ./fct-parser -file topic-44707.json "https://xn----8sba0bbi0cdm.xn--p1a
i/qa/question/view-44707"
2022/11/09 00:13:04 The file ./topic-44707.json was successeful writing
```