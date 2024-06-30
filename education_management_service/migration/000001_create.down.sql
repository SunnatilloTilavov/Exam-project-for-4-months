-- Drop tables in reverse order of their creation

DROP TABLE IF EXISTS student_payment;
DROP TABLE IF EXISTS event_student;
DROP TABLE IF EXISTS event;
DROP TABLE IF EXISTS student_task;
DROP TABLE IF EXISTS task;
DROP TABLE IF EXISTS schedule;
DROP TABLE IF EXISTS "group";
DROP TABLE IF EXISTS manager;
DROP TABLE IF EXISTS student;
DROP TABLE IF EXISTS administration;
DROP TABLE IF EXISTS teacher;
DROP TABLE IF EXISTS support_teacher;
DROP TABLE IF EXISTS journal;
DROP TABLE IF EXISTS branch;

-- Drop UUID extension if it exists
DROP EXTENSION IF EXISTS "uuid-ossp";
