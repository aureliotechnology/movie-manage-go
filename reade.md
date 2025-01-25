# Movie Catalog - Robust and Modular Film Management System

## **Overview**
The Movie Catalog project is a robust and modular system for managing a collection of movies. It is designed to showcase best practices in software architecture, including **Domain-Driven Design (DDD)**, **Hexagonal Architecture**, **Test-Driven Development (TDD)**, and **Clean Code** principles. This project is also integrated with **SonarQube** to ensure high-quality code and maintainability.

---

## **Features**

### **Core Functionalities**
- [ ] **Movies**:
  - Title, director, release date, duration, synopsis, and rating.
  - Support for multiple genres and tags.
  - Related movies.
  - File attachments (e.g., posters and trailers).

- [ ] **Users**:
  - User authentication and registration.
  - Favorites and history tracking.

- [ ] **Comments**:
  - Add, view, and moderate comments on movies.

- [ ] **Genres and Tags**:
  - Predefined genres with the ability to extend.
  - Tags for categorizing movies.

- [ ] **Recommendations**:
  - Suggest movies based on user preferences and history.

- [ ] **Administrative Tools**:
  - Manage movies, users, and comments.
  - Generate usage and activity reports.

---

## **Architecture**

The project leverages a layered and modular architecture for scalability and maintainability:

### **Hexagonal Architecture**
- **Core (Domain)**:
  - Contains the business logic and entities.
  - Independent of frameworks and external dependencies.

- **Application**:
  - Handles use cases and orchestrates the interaction between the domain and adapters.

- **Adapters**:
  - Ports and adapters for input (e.g., API handlers) and output (e.g., database repositories).

### **Domain-Driven Design (DDD)**
- Rich domain models for entities like `Movie`, `Genre`, `User`, and `Comment`.
- Value objects for encapsulating domain-specific logic (e.g., `Duration`, `ReleaseDate`).

### **Infrastructure**
- PostgreSQL for relational data storage.
- SonarQube for code quality analysis.
- Docker and Docker Compose for local development.

---

## **Project Structure**
```plaintext
movie-catalog/
├── api/                        # Entry points (REST/GraphQL)
├── cmd/                        # Application entry point
├── internal/                   # Core business logic
│   ├── app/                    # Application layer (use cases)
│   ├── core/                   # Domain layer
│   │   ├── domain/             # Domain entities
│   │   ├── service/            # Domain services
│   │   └── port/               # Input and output ports
│   ├── adapter/                # Adapters for external systems
│   │   ├── repository/         # Database repositories
│   │   ├── external/           # External integrations (e.g., S3)
│   │   └── db/                 # Database connection
├── tests/                      # Centralized tests (mocks and suites)
├── docker-compose.yml          # Local deployment configuration
├── Dockerfile                  # Application Dockerfile
├── sonar-project.properties    # SonarQube configuration
└── README.md                   # Project documentation
```

---

## **Technologies Used**

### **Backend**
- **Language**: Go
- **Framework**: Echo (or Gin)
- **Database**: PostgreSQL
- **ORM**: GORM (or direct use of `database/sql`).
- **Testing**: Built-in Go testing framework with Ginkgo/Gomega for BDD.

### **Infrastructure**
- **Containerization**: Docker and Docker Compose
- **Code Quality**: SonarQube
- **Logging**: Zap or Logrus
- **Metrics**: Prometheus and Grafana (optional)

---

## **Setup Instructions**

### **Prerequisites**
- Docker and Docker Compose installed.
- Go (latest version).
- PostgreSQL installed locally (optional if using Docker).

### **Local Development**
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/movie-catalog.git
   cd movie-catalog
   ```

2. Start the services with Docker Compose:
   ```bash
   docker-compose up --build
   ```

3. Access the API at: `http://localhost:3000`.

4. (Optional) Access SonarQube at: `http://localhost:9000`.

### **Run Tests**
- [ ] Execute the test suite:
```bash
go test ./...
```

---

## **Code Quality and Coverage**

### **SonarQube Configuration**
1. Start SonarQube using Docker:
   ```bash
   docker-compose up -d sonarqube
   ```

2. Analyze the project:
   ```bash
   sonar-scanner
   ```

### **Coverage Reports**
- [ ] Generate and visualize test coverage using Go tools.
  ```bash
  go test ./... -coverprofile=coverage.out
  go tool cover -html=coverage.out
  ```

---

## **Future Improvements**
- Add support for GraphQL.
- Implement AI-based movie recommendations.
- Enhance observability with detailed metrics and tracing.
- Add caching for frequently accessed endpoints.

---

## **Contributing**
1. Fork the repository.
2. Create a feature branch: `git checkout -b feature-name`.
3. Commit changes: `git commit -m 'Add new feature'`.
4. Push to the branch: `git push origin feature-name`.
5. Submit a pull request.

---

## **License**
This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## **Contact**
For questions or feedback, please contact:
- **Name**: Your Name
- **Email**: your.email@example.com
- **GitHub**: [your-username](https://github.com/your-username)
