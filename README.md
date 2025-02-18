# api-students
API for students enrolled in the company

Routes:
    - GET /students - List all students
    - POST /students - Creat student
    - GET /students/:id - Get infos from a specific student
    - PUT /students/:id - Update student
    - DELETE /student/:id - Delete student

Struct Student:
    - Name (string)
    - CPF (int)
    - Email (string)
    - Age (int)
    - Active (bool)