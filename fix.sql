DELETE FROM member WHERE user_id NOT IN (SELECT id FROM "user");
DELETE FROM classroom WHERE owner_id NOT IN (SELECT id FROM "user");
DELETE FROM member WHERE class_id NOT IN (SELECT id FROM classroom);
DELETE FROM assignment WHERE class_id NOT IN (SELECT id FROM classroom);
DELETE FROM playground WHERE assignment_id NOT IN (SELECT id FROM assignment);
DELETE FROM submission WHERE assignment_id NOT IN (SELECT id FROM assignment);
UPDATE assignment SET condition = '{"type": "register_eq", "register": "A", "value": 0}' WHERE condition::text LIKE '[%';
