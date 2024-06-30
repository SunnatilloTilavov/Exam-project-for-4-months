INSERT INTO branch (id, name, address, phone, created_at, updated_at)
VALUES
   ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Branch A', '123 Main St, City A', '+1234567890', NOW(), NOW()),
   ('7c9e6679-7425-40de-944b-e07fc1f90ae7', 'Branch B', '456 Elm St, City B', '+9876543210', NOW(), NOW()),
   ('7b52009b-64fd-4e7a-b2e0-df9e37b79d83', 'Branch C', '789 Oak St, City C', '+2468135790', NOW(), NOW());

--INSERT INTO journal (id, fromDate, toDate, studentsCount, created_at, updated_at)
--VALUES
--    ('4c9e6679-7425-40de-944b-e07fc1f90ae7', '2024-06-01', '2024-06-30', 50, NOW(), NOW()),
--    ('5b52009b-64fd-4e7a-b2e0-df9e37b79d83', '2024-07-01', '2024-07-31', 60, NOW(), NOW());
--
--
INSERT INTO support_teacher (id, login, fullname, phone, password, salary, ieltsScore, ieltsAttemptCount, branchId, created_at, updated_at)
VALUES
   ('3aeebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'ST00001', 'Support Teacher 1', '+9876543210', 'password123', 2000, 7.5, 2, 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
   ('6b52009b-64fd-4e7a-b2e0-df9e37b79d83', 'ST00002', 'Support Teacher 2', '+1234567890', 'password456', 1800, 6.8, 1, '7c9e6679-7425-40de-944b-e07fc1f90ae7', NOW(), NOW());

INSERT INTO teacher (id, login, fullname, phone, password, salary, ieltsScore, ieltsAttemptCount, supportTeacherId, branchId, created_at, updated_at)
VALUES
   ('8aeebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'T00001', 'Teacher 1', '+9876543210', 'password123', 2500, 8.0, 3, '3aeebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
   ('9b52009b-64fd-4e7a-b2e0-df9e37b79d83', 'T00002', 'Teacher 2', '+2468135790', 'password789', 2300, 7.2, 2, '6b52009b-64fd-4e7a-b2e0-df9e37b79d83', '7c9e6679-7425-40de-944b-e07fc1f90ae7', NOW(), NOW());

--INSERT INTO administration (id, login, fullname, phone, password, salary, ieltsScore, branchId, created_at, updated_at)
--VALUES
--    ('1aeebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'A00001', 'Admin 1', '+9876543210', 'password123', 3000, 7.0, '7b52009b-64fd-4e7a-b2e0-df9e37b79d83', NOW(), NOW()),
--     ('2b52009b-64fd-4e7a-b2e0-df9e37b79d83', 'A00002', 'Admin 2', '+1234567890', 'password456', 2800, 6.5, 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW());
--
INSERT INTO student (id, login, fullname, phone, password, groupName, branchId, created_at, updated_at)
VALUES
   ('c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'S00001', 'Student 1', '+9876543210', 'password123', 'Group A', '7b52009b-64fd-4e7a-b2e0-df9e37b79d83', NOW(), NOW()),
   ('d1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'S00002', 'Student 2', '+1234567890', 'password456', 'Group B', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW());
 
--
--INSERT INTO manager (id, login, fullname, salary, phone, password, branchId, created_at, updated_at)
--VALUES
--    ('f0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'M00001', 'Manager 1', 4000, '+9876543210', 'password123', '7c9e6679-7425-40de-944b-e07fc1f90ae7', NOW(), NOW()),
--    ('d1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'M00002', 'Manager 2', 3800, '+2468135790', 'password456', '7b52009b-64fd-4e7a-b2e0-df9e37b79d83', NOW(), NOW());
--
--
--INSERT INTO "group" (id, teacherId, supportTeacherId, journalId, branchId, created_at, updated_at, type)
--VALUES
--    ('f0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', '8aeebc99-9c0b-4ef8-bb6d-6bb9bd380a11', '3aeebc99-9c0b-4ef8-bb6d-6bb9bd380a11', '4c9e6679-7425-40de-944b-e07fc1f90ae7', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW(), 'beginner'),
--    ('d1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', '9b52009b-64fd-4e7a-b2e0-df9e37b79d83', '6b52009b-64fd-4e7a-b2e0-df9e37b79d83', '5b52009b-64fd-4e7a-b2e0-df9e37b79d83', '7c9e6679-7425-40de-944b-e07fc1f90ae7', NOW(), NOW(), 'intermediate');
--
--
--INSERT INTO schedule (id, journalId, date, startTime, endTime, lesson, created_at, updated_at)
--VALUES
--    ('f0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', '4c9e6679-7425-40de-944b-e07fc1f90ae7', '2024-06-05', '10:00:00', '12:00:00', 'Mathematics', NOW(), NOW()),
--    ('d1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', '5b52009b-64fd-4e7a-b2e0-df9e37b79d83', '2024-07-10', '09:00:00', '11:00:00', 'English', NOW(), NOW());
--
--
--INSERT INTO task (id, scheduleId, label, deadlineDate, deadlineTime, score, created_at, updated_at)
--VALUES
--    ('f0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'f0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Homework 1', '2024-06-10', '23:59:00', 100, NOW(), NOW()),
--    ('d1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'd1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Essay', '2024-07-15', '18:00:00', 90, NOW(), NOW());
--
--
--INSERT INTO student_task (id, taskId, studentId, created_at, updated_at)
--VALUES
--    ('f0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'f0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
--    ('d1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'd1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'd1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW());
--
INSERT INTO event (id, assignStudent, topic, startTime, date, branchId, created_at, updated_at)
VALUES
   ('f0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'student1', 'Parent-Teacher Meeting', '14:00:00', '2024-06-15', '7b52009b-64fd-4e7a-b2e0-df9e37b79d83', NOW(), NOW()),
   ('d1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'student2', 'Field Trip', '09:00:00', '2024-07-20', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW());

--
INSERT INTO event_student (id, eventId, studentId, created_at, updated_at)
VALUES
   ('f0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'f0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
   ('d1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'd1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'd1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW());

--
--INSERT INTO student_payment (id, studentId, groupId, paidSum, administrationId, created_at, updated_at)
--VALUES
--    ('f0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'f0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 500.00, '1aeebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
--    ('d1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'd1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'd1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 700.00, '2b52009b-64fd-4e7a-b2e0-df9e37b79d83', NOW(), NOW());
--
--
--