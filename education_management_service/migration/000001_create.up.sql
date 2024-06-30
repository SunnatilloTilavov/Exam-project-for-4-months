-- Enable UUID extension (for PostgreSQL)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create tables
CREATE TABLE branch (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL UNIQUE,
    address VARCHAR(255),
    phone VARCHAR(20) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
;

CREATE TABLE support_teacher (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    login VARCHAR(255) UNIQUE NOT NULL,
    fullname VARCHAR(255),
    phone VARCHAR(20) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    salary INT,
    ieltsScore DECIMAL(10, 2),
    ieltsAttemptCount INT,
    branchId UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (branchId) REFERENCES branch(id)
);

CREATE TABLE teacher (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    login VARCHAR(255) UNIQUE NOT NULL,
    fullname VARCHAR(255),
    phone VARCHAR(20) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    salary INT,
    ieltsScore DECIMAL(10, 2),
    ieltsAttemptCount INT,
    supportTeacherId UUID,
    branchId UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (supportTeacherId) REFERENCES support_teacher(id),
    FOREIGN KEY (branchId) REFERENCES branch(id)
);

CREATE TABLE administration (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    login VARCHAR(255) UNIQUE NOT NULL,
    fullname VARCHAR(255),
    phone VARCHAR(20),
    password VARCHAR(255) NOT NULL,
    salary INT,
    ieltsScore DECIMAL(10, 2),
    branchId UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (branchId) REFERENCES branch(id)
);

CREATE TABLE student (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    login VARCHAR(255) UNIQUE NOT NULL,
    fullname VARCHAR(255),
    phone VARCHAR(20) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    groupName VARCHAR(255),
    branchId UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (branchId) REFERENCES branch(id)
);

CREATE TABLE manager (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    login VARCHAR(255) UNIQUE NOT NULL,
    fullname VARCHAR(255),
    salary DECIMAL(10, 2),
    phone VARCHAR(20) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    branchId UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (branchId) REFERENCES branch(id)
);

CREATE TABLE "group" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) UNIQUE NOT NULL,
    teacherId UUID,
    supportTeacherId UUID,
    branchId UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    type VARCHAR(50) CHECK (type IN ('beginner', 'elementary', 'intermediate', 'ielts')),
    FOREIGN KEY (teacherId) REFERENCES teacher(id),
    FOREIGN KEY (supportTeacherId) REFERENCES support_teacher(id),
    FOREIGN KEY (branchId) REFERENCES branch(id)
);


CREATE TABLE journal (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    groupId UUID,
    fromDate DATE,
    toDate DATE,
    studentsCount INT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (groupId) REFERENCES "group"(id)
);

CREATE TABLE schedule (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    journalId UUID,
    date DATE,
    startTime TIME,
    endTime TIME,
    lesson VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (journalId) REFERENCES journal(id)
);

CREATE TABLE task (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    scheduleId UUID,
    label VARCHAR(255),
    deadlineDate DATE,
    deadlineTime TIME,
    score INT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (scheduleId) REFERENCES schedule(id)
);

CREATE TABLE student_task (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    taskId UUID,
    studentId UUID,
    score INTEGER, 
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (taskId) REFERENCES task(id),
    FOREIGN KEY (studentId) REFERENCES student(id)
);

CREATE TABLE event (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    assignStudent VARCHAR(255),
    topic VARCHAR(255),
    startTime TIME,
    date DATE,
    branchId UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (branchId) REFERENCES branch(id)
);

CREATE TABLE event_student (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    eventId UUID,
    studentId UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (eventId) REFERENCES event(id),
    FOREIGN KEY (studentId) REFERENCES student(id)
);

CREATE TABLE student_payment (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    studentId UUID,
    groupId UUID,
    paidSum DECIMAL(10, 2),
    administrationId UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (studentId) REFERENCES student(id),
    FOREIGN KEY (groupId) REFERENCES "group"(id),
    FOREIGN KEY (administrationId) REFERENCES administration(id)
);
