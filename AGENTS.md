# Book Service Training Protocol

## Project Goal

Build an MVP Go project for managing a list of books.

The service should eventually support:
- adding books;
- reading book data;
- updating book data;
- deleting books;
- listing books with pagination;
- searching and filtering books;
- a structure that is easy to extend horizontally and vertically.

The project is educational. The user is learning Go and should implement the project step by step.

## Collaboration Rules

- The assistant owns the hidden roadmap and decides the order of tasks.
- The assistant gives one task specification at a time.
- Task specifications must be written as requirements only.
- The assistant must not include hints, examples, code snippets, or implementation instructions in task specifications.
- The user implements each task independently.
- If the user does not understand something, they may ask other agents or ask a direct clarification.
- After the user reports that a task is done, the assistant reviews the result.
- Reviews should focus on correctness, Go best practices, extensibility, simplicity, naming, package boundaries, and tests.
- After review, the assistant either asks for corrections or gives the next task.
- The project should stay intentionally simple, but not careless.
- Prefer standard library first. Add dependencies only when there is a clear reason.

## Technical Direction

- Language: Go.
- Project type: backend service.
- Initial delivery style: small, reviewable increments.
- Architecture goal: clear boundaries between entrypoints, application logic, domain concepts, and infrastructure.
- MVP scope should grow incrementally from project skeleton to CRUD, persistence, pagination, search, validation, tests, and operational basics.

## Task Format

Each task should contain:
- goal;
- requirements;
- acceptance criteria;
- what to report back when done.

Each task should avoid:
- examples;
- ready-made code;
- package layout answers unless the task is specifically about choosing or creating structure;
- explanations of how to implement the requirements.
