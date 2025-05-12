
# WorkMate Microservices Architecture

This repository contains the microservices architecture for a **WorkMate**. The system is designed to be modular, scalable, and easy to maintain. Each service is self-contained and responsible for a specific business function. This architecture ensures flexibility, fault tolerance, and ease of development.

---

## Services Overview

### 1. **User Service**
   - **Description**: Manages user authentication, registration, login, and profile management. Handles roles like admin, user, etc.
   - **Key Functions**:
     - User registration and login.
     - User profile management.
     - Role management (admin, user).

### 2. **Task Service**
   - **Description**: Handles creating, updating, deleting, and retrieving tasks, as well as assigning them to specific users.
   - **Key Functions**:
     - Create, update, delete, and retrieve tasks.
     - Assign tasks to specific users.
     - Handle task statuses (pending, in-progress, completed).

### 3. **Notification Service**
   - **Description**: Sends notifications (via email, SMS, or in-app) when tasks are assigned, updated, or completed.
   - **Key Functions**:
     - Send task-related notifications.
     - Support multiple notification channels (email, SMS, in-app).

### 4. **Authentication Service**
   - **Description**: Manages **JWT authentication** and **authorization** for users to access different parts of the app securely.
   - **Key Functions**:
     - Handle login and JWT generation.
     - User authorization for protected routes.

### 5. **Audit/Logging Service**
   - **Description**: Logs all actions (like task updates, user logins) for **security**, **tracking**, and **debugging** purposes.
   - **Key Functions**:
     - Record user activities and task events.
     - Generate logs for debugging and security auditing.

### 6. **Search Service**
   - **Description**: Provides search functionality for tasks based on keywords, statuses, assignees, etc.
   - **Key Functions**:
     - Index and search tasks by various parameters (keywords, assignees, status).
     - Return filtered search results.

### 7. **File Service**
   - **Description**: Handles **uploading**, **storing**, and **retrieving** any file attachments for tasks (like documents or images).
   - **Key Functions**:
     - Upload and retrieve files related to tasks.
     - Support for file types like images, documents, and others.

### 8. **API Gateway**
   - **Description**: Acts as the **entry point** for all API requests, routing them to the appropriate service, and providing centralized logging and authentication checks.
   - **Key Functions**:
     - Route requests to corresponding microservices.
     - Handle authentication and logging.

### 9. **Analytics Service**
   - **Description**: Collects **usage data** (e.g., number of tasks completed, overdue tasks) for reporting or insights.
   - **Key Functions**:
     - Collect task data for reporting.
     - Provide insights into task completion rates, overdue tasks, etc.

---

## How They Work Together

- **User Service** interacts with **Authentication Service** to manage login.
- **Task Service** stores tasks and sends updates to **Notification Service**.
- **API Gateway** makes sure the requests hit the correct microservice (like routing a login request to **User Service**).
- **File Service** is used when users need to attach files to tasks.

---

## Setup and Installation

### Prerequisites
-

### Steps to Run the System:

1. Clone the repository:
   ```bash
   git clone https://github.com/likhithkp/WorkMate
   cd WorkMate
   ```

2. Build and run the services using Docker Compose:
   ```bash
   docker-compose up --build
   ```

3. Access the services through their respective endpoints:
   - API Gateway: `http://localhost:8080`
   - User Service: `http://localhost:8081`
   - Task Service: `http://localhost:8082`
   - Notification Service: `http://localhost:8083`
   - Authentication Service: `http://localhost:8084`
   - File Service: `http://localhost:8085`
   - Analytics Service: `http://localhost:8086`

---

## Contributing

1. Fork the repository.
2. Create a new branch for your changes.
3. Submit a pull request with a clear description of your changes.

---

## License

This project is licensed under the MIT License.

