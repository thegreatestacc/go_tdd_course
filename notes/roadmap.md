# Roadmap: Learn Go with Tests

Курс: `course/learn-go-with-tests/`  
Источник оглавления: [SUMMARY.md](../course/learn-go-with-tests/SUMMARY.md)

Шкала сложности: **1** — базовый старт, **5** — продвинутый уровень.

---

## Go fundamentals

### 1. Install Go
- **Основные темы:** установка Go, `GOPATH` / модули, настройка редактора, `go run`, `go test`, структура проекта
- **Сложность:** 1
- **Для собеседований:** знать команды `go mod init`, `go build`, `go test`, `go vet`; понимать layout модуля (`go.mod`, `main.go`). Редко спрашивают детально, но без этого не пройти практику.

### 2. Hello, World
- **Основные темы:** `package main`, `func main`, импорты, разделение domain/side-effects, первый тест, `testing` package, subtests, closures
- **Сложность:** 1
- **Для собеседований:** структура Go-программы, идиома «тестируем логику, не `fmt.Println`», синтаксис `t.Run`, table-driven подход закладывается здесь.

### 3. Integers
- **Основные темы:** объявление функций, типы (`int`), арифметика, godoc-комментарии, именованные возвращаемые значения
- **Сложность:** 1
- **Для собеседований:** базовый синтаксис функций, zero values, разница `int` / `int64`; умение писать читаемый и документированный код.

### 4. Iteration
- **Основные темы:** цикл `for`, `range`, бенчмарки (`Benchmark*`), `testing.B`, сравнение производительности
- **Сложность:** 2
- **Для собеседований:** единственный цикл в Go — `for`; уметь читать и писать benchmark; на middle+ иногда спрашивают про `b.N`, `b.ResetTimer()`.

### 5. Arrays and Slices
- **Основные темы:** массивы vs слайсы, `len`, `cap`, `append`, varargs, `range`, `make`, test coverage (`go test -cover`)
- **Сложность:** 2
- **Для собеседований:** **очень частая тема** — отличие array/slice, как работает `append`, sharing underlying array, slice expressions `s[low:high]`, capacity growth.

### 6. Structs, Methods & Interfaces
- **Основные темы:** `struct`, методы с value/pointer receiver, интерфейсы (implicit satisfaction), table-driven tests, композиция
- **Сложность:** 3
- **Для собеседований:** **ключевая глава** — интерфейсы в Go, когда pointer receiver, embedding, table-driven tests; основа большинства Go-вопросов на собесах.

### 7. Pointers & Errors
- **Основные темы:** указатели (`*T`, `&x`), `error` interface, создание и обработка ошибок, `errors.New`, идиома `if err != nil`
- **Сложность:** 3
- **Для собеседований:** **обязательно** — nil pointer, pointer vs value semantics, error handling без exceptions; часто дают задачи на исправление багов с указателями.

### 8. Maps
- **Основные темы:** `map[K]V`, инициализация (`make` / literal), чтение/запись, проверка наличия ключа (`ok` idiom), итерация
- **Сложность:** 2
- **Для собеседований:** maps не упорядочены, zero value `nil`, нельзя брать адрес элемента map; частые задачи на подсчёт частот, группировку, dedup.

### 9. Dependency Injection
- **Основные темы:** DI через интерфейсы, `io.Writer` / `io.Reader`, инверсия зависимостей, тестируемый дизайн
- **Сложность:** 3
- **Для собеседований:** умение проектировать зависимости через интерфейсы; `io.Writer` как абстракция вывода — стандартный паттерн в Go-кодовых базах.

### 10. Mocking
- **Основные темы:** тестовые doubles (stubs/mocks), DI для изоляции, тестирование без реальных HTTP-вызовов, `httptest`
- **Сложность:** 3
- **Для собеседований:** как тестировать код с внешними зависимостями; разница mock/stub/fake; на backend-собесах часто просят спроектировать тестируемый HTTP handler.

### 11. Concurrency
- **Основные темы:** goroutines, channels, `go` keyword, параллельная обработка, race conditions (вводятся)
- **Сложность:** 4
- **Для собеседований:** **топ-тема** — goroutines vs threads, channel basics, когда параллелизм ускоряет, а когда нет; подготовка к вопросам про data races.

### 12. Select
- **Основные темы:** `select` для multiplexing каналов, таймауты, non-blocking receive, координация goroutines
- **Сложность:** 4
- **Для собеседований:** `select` с `default`, паттерн timeout через `time.After`, graceful coordination; часто в задачах на concurrency.

### 13. Reflection
- **Основные темы:** `reflect` package, `TypeOf`, `ValueOf`, runtime type inspection, осторожное использование
- **Сложность:** 4
- **Для собеседований:** знать, что reflection есть и дорогой; где встречается (JSON, ORM, теги struct); на собесах реже пишут reflection вручную, но спрашивают trade-offs.

### 14. Sync
- **Основные темы:** `sync.WaitGroup`, `sync.Mutex`, `sync.RWMutex`, синхронизация доступа к shared state
- **Сложность:** 4
- **Для собеседований:** **часто** — когда mutex vs channel, deadlock, WaitGroup pattern; must-know для любого Go backend interview.

### 15. Context
- **Основные темы:** `context.Context`, cancellation, deadlines, `WithCancel`, `WithTimeout`, propagation по вызовам
- **Сложность:** 4
- **Для собеседований:** **обязательно для production Go** — передача context в HTTP/DB/gRPC, отмена долгих операций, не хранить context в struct.

### 16. Intro to Property Based Tests (Roman Numerals)
- **Основные темы:** property-based testing, `rapid` / quickcheck-стиль, kata Roman Numerals, инварианты, round-trip тесты
- **Сложность:** 3
- **Для собеседований:** показывает зрелость в тестировании; полезно упоминать property-based подход для парсеров/кодировщиков; сама kata — хорошая live-coding задача.

### 17. Maths
- **Основные темы:** пакет `math`, генерация SVG, работа с координатами, sin/cos, форматирование вывода
- **Сложность:** 2
- **Для собеседований:** низкий приоритет; полезно как пример работы со `strings.Builder` и файловым выводом, но math/SVG редко спрашивают.

### 18. Reading Files
- **Основные темы:** `os.Open`, `io.ReadAll`, `bufio`, обработка файлов, тестирование file I/O через `strings.Reader` / temp files
- **Сложность:** 3
- **Для собеседований:** работа с `io` abstractions, обработка ошибок при I/O, `defer file.Close()`; типичные задачи на парсинг логов/CSV.

### 19. Templating
- **Основные темы:** `html/template`, рендеринг HTML из данных, XSS-safe templates, approval testing
- **Сложность:** 3
- **Для собеседований:** для backend — знать `html/template` vs `text/template`; auto-escaping; на fullstack Go — рендеринг страниц, layout templates.

### 20. Generics
- **Основные темы:** type parameters, constraints, generic functions и data structures, type inference, `[T any]`
- **Сложность:** 4
- **Для собеседований:** с Go 1.18+ **спрашивают** — когда generics уместны, constraints, отличие от interfaces; не перегибать с generics в Go-стиле.

### 21. Revisiting Arrays and Slices with Generics
- **Основные темы:** generic `Reduce`, `Stack[T]`, обобщённые коллекции, устранение дублирования кода
- **Сложность:** 4
- **Для собеседований:** умение применить generics к collections; хорошая практика для задач «напиши generic Stack/Queue/Filter».

---

## Testing fundamentals

### 22. Introduction to Acceptance Tests
- **Основные темы:** acceptance/end-to-end тесты, graceful shutdown HTTP-сервера, SIGTERM/SIGKILL, `httptest.Server`, black-box testing
- **Сложность:** 4
- **Для собеседований:** разница unit/integration/acceptance; graceful shutdown — реальный production topic; показывает понимание жизненного цикла сервиса.

### 23. Scaling Acceptance Tests
- **Основные темы:** сложность e2e-тестов, test pyramid, contract tests, управление flakiness, тестовая инфраструктура
- **Сложность:** 4
- **Для собеседований:** для senior — **архитектура тестирования**, когда e2e не масштабируются, contract testing как альтернатива mocks.

### 24. Working without Mocks
- **Основные темы:** fakes, contract tests, тестирование через реальные границы, maintainable tests без spy/stub hell
- **Сложность:** 4
- **Для собеседований:** зрелый взгляд на тестирование; отличие fake vs mock; на lead/senior собесах — как ты строишь тестовую стратегию в команде.

### 25. Refactoring Checklist
- **Основные темы:** безопасный рефакторинг под тестами, code smells, TDD refactor phase, чеклист изменений
- **Сложность:** 2
- **Для собеседований:** демонстрирует инженерную культуру; полезно говорить про red-green-refactor на собеседовании; напрямую задачи редко дают.

---

## Build an application

### 26. Intro (Build an Application)
- **Основные темы:** итеративная разработка приложения, TDD на реальном проекте, наращивание функциональности по главам
- **Сложность:** 2
- **Для собеседований:** meta-глава; полезна для рассказа о подходе к разработке (incremental delivery, TDD workflow).

### 27. HTTP Server
- **Основные темы:** `net/http`, handlers, routing, `httptest`, dependency injection в handlers, JSON responses, middleware basics
- **Сложность:** 3
- **Для собеседований:** **must-have** — `http.Handler` / `HandlerFunc`, status codes, тестирование handlers через `httptest.ResponseRecorder`; основа Go backend.

### 28. JSON, Routing and Embedding
- **Основные темы:** `encoding/json`, marshal/unmarshal, struct tags, HTTP routing, embedding types, composition в API
- **Сложность:** 3
- **Для собеседований:** JSON tags, обработка неизвестных полей, embedded structs; стандартные вопросы на API-разработку.

### 29. IO and Sorting
- **Основные темы:** персистентность на диск, `sort` package, `io` interfaces, file-based storage, тестирование persistence layer
- **Сложность:** 3
- **Для собеседований:** `sort.Interface` / `slices.Sort`, работа с файлами, separation of concerns между storage и business logic.

### 30. Command Line & Package Structure
- **Основные темы:** `flag` / CLI args, multi-binary project layout, пакетная структура, `main` packages
- **Сложность:** 3
- **Для собеседований:** организация Go-проекта (`cmd/`, `internal/`, `pkg/`), cobra/flag basics; на middle+ спрашивают project layout.

### 31. Time
- **Основные темы:** `time` package, scheduling, tickers, `Sleep`, тестирование time-dependent code
- **Сложность:** 3
- **Для собеседований:** `time.Time` vs `time.Duration`, monotonic clock, тестирование через injection clock; задачи на rate limiting, timeouts.

### 32. WebSockets
- **Основные темы:** WebSocket server, `gorilla/websocket` или аналог, real-time communication, тестирование WS
- **Сложность:** 4
- **Для собеседований:** реже базовых HTTP-вопросов, но встречается в real-time проектах; понимание upgrade protocol, concurrent connections.

---

## Questions and Answers

### 33. OS Exec
- **Основные темы:** `os/exec`, вызов внешних команд, тестирование через DI, изоляция side effects
- **Сложность:** 3
- **Для собеседований:** безопасность (`exec.Command` с аргументами, не shell injection), тестируемость обёрток над CLI-утилитами.

### 34. Error Types
- **Основные темы:** custom error types, error wrapping (`fmt.Errorf` + `%w`), `errors.Is`, `errors.As`, typed errors
- **Сложность:** 4
- **Для собеседований:** **очень важно** — sentinel errors vs custom types, wrapping/unwrapping, проверка цепочки ошибок; стандарт с Go 1.13+.

### 35. Context-aware Reader
- **Основные темы:** расширение `io.Reader` через context, cancellation при чтении, composable interfaces
- **Сложность:** 5
- **Для собеседований:** продвинутая композиция интерфейсов; показывает глубокое понимание `io` + `context`; на senior — отличный differentiator.

### 36. Revisiting HTTP Handlers
- **Основные темы:** правильный дизайн HTTP handlers, тестируемость, separation of concerns, common handler anti-patterns
- **Сложность:** 4
- **Для собеседований:** **практично** — как не смешивать routing, validation, business logic; структура handler'ов в production code.

---

## Meta

### 37. Why Unit Tests and TDD
- **Основные темы:** мотивация TDD, ценность unit-тестов, feedback loop, уверенность при рефакторинге
- **Сложность:** 1
- **Для собеседований:** soft-skill / culture fit вопросы: «как ты тестируешь», «зачем TDD»; подготовить 2–3 конкретных примера из практики.

### 38. Anti-patterns
- **Основные темы:** TDD anti-patterns, over-mocking, testing implementation details, brittle tests
- **Сложность:** 2
- **Для собеседований:** полезно для senior-дискуссий; показывает, что ты знаешь не только «как тестировать», но и «как не надо».

---

## Рекомендуемый порядок изучения

```
База (1–8) → Интерфейсы и DI (9–10) → Concurrency (11–15) → Generics (20–21)
     ↓
HTTP-приложение (27–32) → Acceptance testing (22–24) → Q&A по мере необходимости
```

## Приоритет для собеседований Go Backend

| Приоритет | Главы |
|-----------|-------|
| **Критично** | 5, 6, 7, 8, 11, 14, 15, 27, 34 |
| **Высокий** | 4, 9, 10, 12, 16, 20, 28, 30, 36 |
| **Средний** | 18, 19, 22, 29, 31, 33, 35 |
| **По желанию** | 13, 17, 23, 24, 32, 37, 38 |

---

*Сгенерировано на основе структуры `course/learn-go-with-tests/`.*
