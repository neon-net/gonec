[![GitHub issues](https://img.shields.io/github/issues/covrom/gonec.svg)](https://github.com/covrom/gonec/issues) [![Travis](https://travis-ci.org/covrom/gonec.svg?branch=master)](https://github.com/covrom/gonec/releases)

[![Gonec Logo](/extra/gonec.png)](https://github.com/covrom/gonec/releases)

[![Presentation](/extra/button_ppt.png)](https://gitpitch.com/covrom/gonec)
[![Demo site](/extra/button_play.png)](https://gonec.herokuapp.com/)

[![Download](/extra/button_down.png)](https://github.com/covrom/gonec/releases)
[![Docs](/extra/button_doc.png)](https://github.com/covrom/gonec/wiki)

[![Chat](/extra/button_chat.png)](https://gitter.im/gonec/Lobby)
[![Blog](/extra/button_blog.png)](https://www.facebook.com/gonecplatform/)

## Цели

Платформа `Гонец:Микросервисы` создана для решения программистами 1С задач, связанных с высокопроизводительными распределенными вычислениями, создания микросервисов, вэб-сервисов и вэб-порталов для работы тысяч пользователей, работы с высокоэффективными базами данных, с использованием синтаксиса языка, похожего, но не ограниченного возможностями языка 1С.

Еще никогда программистам 1С не были так легко доступны возможности:
* Создать микросервис с произвольным сетевым протоколом, развернуть его на linux, в docker контейнере или кластере kubernetes
* Выполнить сложную многопоточную вычислительную задачу для десятков тысяч подключающихся пользователей за миллисекунды
* Взаимодействовать с пользователем через web-браузер с минимальным трафиком
* Сохранять и получать данные с максимально доступной скоростью в key-value базах данных

Более подробно см. в [презентации](https://gitpitch.com/covrom/gonec)

## Описание синтаксиса языка и примеры использования интерпретатора

[Документация находится здесь](https://github.com/covrom/gonec/wiki)

## Масштабируемость языка и платформы
Язык Гонец расширяется путем изменения правил синтаксиса в формате YACC, а так же написания собственных высокоэффективных библиотек структур и функций на Го, которые могут быть доступны как объекты метаданных в языке Гонец.

Посмотреть на использование интерпретатора в роли микросервиса можно по [ссылке](https://gonec.herokuapp.com/) выше.
В этой реализации в интерпретатор встроена простая система запуска кода через обычный браузер, которая работает на технологии ajax, общающаяся с микросервисом сессий исполнения кода интерпретатором.

## Какова производительность интерпретатора?
Производительность выше, чем 1С, в десятки раз, и выше, чем у интерпретатора языка Python при работе с большими данными и сетевыми соединениями.
Скорость интерпретации кода соответствует скорости программ на Go и скорости работы библиотек, написанных на Go.

На платформе Гонец возможна реализация как обработки больших объемов данных, так и быстрой обработки часто поступающих клиентских запросов.

Интерпретатор языка использует повторное выделение памяти в синхронизированном пуле, что сокращает расход памяти даже при выполнении глубоких рекурсивных алгоритмов.

Пример сравнения производительности цикла без тела, перебор значений от 1 до 1 млн.
Участники сравнения:
* Гонец с регистровой виртуальной машиной
* 1С:Предприятие 8.3.9.2170 (файловая)

![PerfVs1C](/extra/perf1c.gif)

Производительность одновременного запуска 1000 серверных и 1000 клиентских подключений, по протоколам TCP и HTTP, на 4-ядерном Core-i5 3570

![PerfConnect](/extra/http_perfomance.png)

## Какой статус разработки интерпретатора?
Интерпретатор работает стабильно, протестирован и находится в стадии разработки стандартной библиотеки.
