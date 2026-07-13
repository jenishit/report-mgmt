# Graph Report - .  (2026-07-13)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 678 nodes · 1380 edges · 37 communities (19 shown, 18 thin omitted)
- Extraction: 83% EXTRACTED · 17% INFERRED · 0% AMBIGUOUS · INFERRED: 231 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `596588e7`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- handleError
- Password
- LabTestsRepository
- main
- Result
- config.go
- Report
- ListVisit
- Order
- Patient
- GetProfileDetails
- Doctor
- LabSettings
- labTests.go
- DB
- 20260707080505_create_departments_table.sql
- LoginResponse
- 20260710073046_visit_result_table.sql
- 20260702101734_create_user_table.sql
- 20260709063003_other_users_table.sql
- 20260713074207_report_table.sql
- doctor.go
- lab.go
- labTests.go
- orderItem.go
- patient.go
- profile.go
- report.go
- result.go
- role.go
- users.go
- visit.go
- CreateRole
- 20260702094601_create_role_table.sql
- 20260706113407_create_lab_settings.sql
- github.com/jenish-brainztechs/go-backend

## God Nodes (most connected - your core abstractions)
1. `handleError()` - 57 edges
2. `handleSuccess()` - 56 edges
3. `main()` - 37 edges
4. `validationError()` - 36 edges
5. `LabTestsHandler` - 21 edges
6. `NewRouter()` - 20 edges
7. `LabTestsRepository` - 20 edges
8. `LabTestsService` - 20 edges
9. `Result` - 16 edges
10. `Order` - 14 edges

## Surprising Connections (you probably didn't know these)
- `main()` --calls--> `NewLabTestsHandler()`  [INFERRED]
  cmd/main.go → internal/adapter/handler/http/labTests.go
- `main()` --calls--> `NewPatientHandler()`  [INFERRED]
  cmd/main.go → internal/adapter/handler/http/patient.go
- `main()` --calls--> `NewReportHandler()`  [INFERRED]
  cmd/main.go → internal/adapter/handler/http/report.go
- `main()` --calls--> `NewResultHandler()`  [INFERRED]
  cmd/main.go → internal/adapter/handler/http/result.go
- `main()` --calls--> `NewVisitHandler()`  [INFERRED]
  cmd/main.go → internal/adapter/handler/http/visit.go

## Import Cycles
- None detected.

## Communities (37 total, 18 thin omitted)

### Community 0 - "handleError"
Cohesion: 0.06
Nodes (37): errorResponse, LabTestsHandler, PatientHandler, ReportHandler, response, ResultHandler, VisitHandler, Context (+29 more)

### Community 1 - "Password"
Cohesion: 0.05
Nodes (37): BasicDetails, Login, LoginResponse, Role, RoleUser, User, UserRole, CreateUser (+29 more)

### Community 2 - "LabTestsRepository"
Cohesion: 0.09
Nodes (18): Department, PanelComponent, Panels, ReferenceRange, TestCatalog, TestParameter, Context, DB (+10 more)

### Community 3 - "main"
Cohesion: 0.07
Nodes (34): AuthService, main(), DoctorService, Engine, HandlerFunc, AuthHandler, DoctorHandler, LabHandler (+26 more)

### Community 4 - "Result"
Cohesion: 0.08
Nodes (28): Flag, Result, BatchCreateResultRequest, BatchResultItem, CreateResultRequest, ResultResponse, UpdateResultRequest, Time (+20 more)

### Community 5 - "config.go"
Cohesion: 0.09
Nodes (35): App, Cache, Container, DB, HTTP, Redis, Refresh, Session (+27 more)

### Community 6 - "Report"
Cohesion: 0.10
Nodes (26): Report, ReportPrint, ReportStatus, CreateReportRequest, ReportPrintRequest, ReportPrintResponse, ReportResponse, UpdateReportRequest (+18 more)

### Community 7 - "ListVisit"
Cohesion: 0.11
Nodes (22): ListVisit, Status, Visit, CreateVisit, ListVisits, VisitRequest, Time, UUID (+14 more)

### Community 8 - "Order"
Cohesion: 0.11
Nodes (21): Order, OrderStatus, CreateOrderRequest, OrderResponse, UpdateOrderRequest, Time, UUID, OrderResponseFromDomain() (+13 more)

### Community 9 - "Patient"
Cohesion: 0.11
Nodes (20): Gender, Patient, CreatePatient, PatientResponse, Time, UUID, PtResponse(), PtResponses() (+12 more)

### Community 10 - "GetProfileDetails"
Cohesion: 0.11
Nodes (19): GetProfileDetails, Profile, ProfileResponse, UpdateProfileRequest, UUID, NewProfileResponse(), NewProfileResponses(), Context (+11 more)

### Community 11 - "Doctor"
Cohesion: 0.11
Nodes (18): DoctorRepository, Doctor, CreateDoctor, DoctorResponse, DocResponse(), DocResponses(), UUID, Context (+10 more)

### Community 12 - "LabSettings"
Cohesion: 0.11
Nodes (18): LabSettings, LabRequest, LabResponse, UUID, LabsResponse(), LabsResponses(), Context, DB (+10 more)

### Community 13 - "labTests.go"
Cohesion: 0.15
Nodes (25): DepartmentRequest, DepartmentResponse, PanelComponentRequest, PanelComponentResponse, PanelRequest, PanelResponse, ReferenceRangeRequest, ReferenceRangeResponse (+17 more)

### Community 14 - "DB"
Cohesion: 0.28
Nodes (5): Context, New(), Pool, DB, StatementBuilderType

### Community 15 - "20260707080505_create_departments_table.sql"
Cohesion: 0.62
Nodes (6): departments, panel_components, panels, reference_ranges, test_catalog, test_parameters

### Community 16 - "LoginResponse"
Cohesion: 0.50
Nodes (4): LoginRequest, LoginResponse, UUID, ToLoginResponse()

### Community 17 - "20260710073046_visit_result_table.sql"
Cohesion: 0.83
Nodes (3): order_item, result, visits

## Knowledge Gaps
- **34 isolated node(s):** `github.com/jenish-brainztechs/go-backend`, `LoginRequest`, `CreateDoctor`, `LabRequest`, `UpdateProfileRequest` (+29 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **18 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `main()` connect `main` to `handleError`, `Password`, `LabTestsRepository`, `Result`, `Report`, `ListVisit`, `Order`, `Patient`, `GetProfileDetails`, `Doctor`, `LabSettings`?**
  _High betweenness centrality (0.403) - this node is a cross-community bridge._
- **Why does `NewRouter()` connect `main` to `handleError`, `config.go`?**
  _High betweenness centrality (0.184) - this node is a cross-community bridge._
- **Are the 53 inferred relationships involving `handleError()` (e.g. with `.Login()` and `.CreateDoctor()`) actually correct?**
  _`handleError()` has 53 INFERRED edges - model-reasoned connections that need verification._
- **Are the 53 inferred relationships involving `handleSuccess()` (e.g. with `.Login()` and `.CreateDoctor()`) actually correct?**
  _`handleSuccess()` has 53 INFERRED edges - model-reasoned connections that need verification._
- **Are the 36 inferred relationships involving `main()` (e.g. with `NewAuthHandler()` and `NewDoctorHandler()`) actually correct?**
  _`main()` has 36 INFERRED edges - model-reasoned connections that need verification._
- **Are the 32 inferred relationships involving `validationError()` (e.g. with `.Login()` and `.CreateDoctor()`) actually correct?**
  _`validationError()` has 32 INFERRED edges - model-reasoned connections that need verification._
- **What connects `github.com/jenish-brainztechs/go-backend`, `LoginRequest`, `CreateDoctor` to the rest of the system?**
  _34 weakly-connected nodes found - possible documentation gaps or missing edges._