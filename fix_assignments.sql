-- Update all assignments to be due next week so they are not "overdue"
UPDATE assignment SET due_date = CURRENT_TIMESTAMP + INTERVAL '7 days';

-- Update assignment titles to be distinct based on the classroom topic
-- We will extract the number from the current title
UPDATE assignment 
SET title = c.topic || ' - Task #' || SUBSTRING(assignment.title FROM '[0-9]+')
FROM classroom c 
WHERE assignment.class_id = c.id 
  AND assignment.title LIKE 'Assignment%';

-- Also make a few randomly overdue to look realistic
UPDATE assignment 
SET due_date = CURRENT_TIMESTAMP - INTERVAL '2 days'
WHERE id % 5 = 0;
