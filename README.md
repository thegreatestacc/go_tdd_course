# Learn Go with Tests: progress

Мои решения упражнений из курса
[Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests).

> Последнее обновление: 13 июня 2026 года

## Текущий результат

| Метрика | Результат |
|---|---:|
| Пройдено основной программы | **4 из 31 глав (12,9%)** |
| Пройдено раздела Go fundamentals | **4 из 20 глав (20%)** |
| Качество текущих решений | **98 / 100** |
| Решено практических этапов | **7** |
| Минимальная версия Go | **1.24** |
| Проверено на Go | **1.24.0** |

Прогресс основной программы:

```text
[####---------------------------] 12,9%
```

## Пройденные главы

| № | Глава | Статус | Реализовано |
|---:|---|---|---|
| 1 | Hello, World | Завершена | Приветствия на нескольких языках, helper для тестов |
| 2 | Integers | Завершена | `Add` |
| 3 | Iteration | Завершена | `Repeat`, benchmark |
| 4 | Arrays and slices | Завершена | `Sum`, `SumAll`, `SumAllTails` |
| 5 | Structs, methods & interfaces | Следующая | — |

Решения находятся в каталоге [`solutions`](solutions).

## Оценка качества

Оценка отражает текущее состояние решений, а не только факт прохождения
существующих тестов.

| Критерий | Балл | Комментарий |
|---|---:|---|
| Корректность | 40 / 40 | Все реализованные функции корректно обрабатывают основные и граничные сценарии |
| Тесты | 29 / 30 | Все пакеты имеют 100% statement coverage; используются subtests и общие test helpers |
| Go-стиль | 19 / 20 | Код проходит `gofmt` и `go vet`, использует idiomatic range, `strings.Builder` и `slices.Equal` |
| Структура проекта | 10 / 10 | Production-код отделён от тестов, а каталоги и пакеты имеют понятные имена |
| **Итого** | **98 / 100** | Известные дефекты устранены, все проверки проходят |

### Что нужно улучшить

Известных дефектов в текущих решениях нет. После прохождения главы
**Structs, methods & interfaces** существующие наборы однотипных подтестов можно
переписать как table-driven tests.

### Сильные стороны

- Все тесты проходят, а statement coverage каждого пакета составляет 100%.
- Используются subtests, общие helpers, `t.Helper()`, benchmark и пакет `slices`.
- Функции корректно обрабатывают пустые входы и другие граничные случаи.
- Production-код находится в обычных `.go` файлах, отдельно от тестов.
- Код проходит `gofmt`, `go test`, `go build` и `go vet`.

## Проверка

Команда для запуска всех решений:

```powershell
go test ./solutions/...
```

Полная проверка:

```powershell
gofmt -w ./solutions
go test ./solutions/...
go test -cover ./solutions/...
go build ./solutions/...
go vet ./solutions/...
```

Все пять пакетов проходят проверки на Go 1.24.0 и имеют 100% statement
coverage. Минимальная версия проекта — Go 1.24, поскольку benchmark использует
добавленный в этой версии `b.Loop()`.

Проверенный результат:

```text
ok LearnGoWithTests/solutions/hello_world
ok LearnGoWithTests/solutions/integers
ok LearnGoWithTests/solutions/iteration
ok LearnGoWithTests/solutions/arrays_and_slices/arrays
ok LearnGoWithTests/solutions/arrays_and_slices/slices
```

## Как считается прогресс

В основную программу включены практические главы из разделов
**Go fundamentals**, **Testing fundamentals** и **Build an application**:
всего 31 глава. Установочные, справочные и meta-главы в процент прогресса
не входят.
