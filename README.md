# Asynchronous Job Scheduler

This is a asynchronous task scheduler service implemented in Golang. It allows you to start, pause, resume, terminate, and check the status of asynchronous tasks.

## Features

- Start a task with a specified target value.
- Pause a running task.
- Resume a paused task.
- Terminate a running task.
- Check the status of a task, including its current status and cursor position.

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

   ```shell
   git clone <repository_url>
   cd simple-task-management-service
   go build -o async_scheduler api/main.go
    ./AsyncJobScheduler
   ```

The service will start running on http://localhost:8080.

### API Endpoints

- POST /start?n=<target> - Start a new task with the specified target value (n).
- PATCH /pause?id=<task_id> - Pause a running task with the given task ID.
- PATCH /resume?id=<task_id> - Resume a paused task with the given task ID.
- PATCH /terminate?id=<task_id> - Terminate a running task with the given task ID.
- GET /status?id=<task_id> - Get the status of a task with the given task ID.


###Contributing

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.



   
