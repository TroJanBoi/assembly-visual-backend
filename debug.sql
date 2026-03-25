-- Check the member records for the user
SELECT '--- Member Records ---';
SELECT * FROM member WHERE user_id = 3;

-- Check the classes they are members of
SELECT '--- Member Classes ---';
SELECT id, owner_id FROM classroom WHERE id IN (SELECT class_id FROM member WHERE user_id = 3);

-- Check if any owner of those classes is missing!
SELECT '--- Missing Owners for Member Classes ---';
SELECT id, owner_id FROM classroom c WHERE id IN (SELECT class_id FROM member WHERE user_id = 3) AND NOT EXISTS (SELECT 1 FROM "user" u WHERE u.id = c.owner_id AND u.deleted_at IS NULL);

-- Check classes owned by the user
SELECT '--- Owned Classes ---';
SELECT id, owner_id FROM classroom WHERE owner_id = 3;

-- Check if the IN query for empty classes fails
SELECT '--- Check soft deleted classrooms ---';
SELECT id, owner_id, deleted_at FROM classroom WHERE deleted_at IS NOT NULL;

-- Check if there are any members pointing to soft deleted classrooms
SELECT '--- Members pointing to deleted classrooms ---';
SELECT m.user_id, m.class_id, c.deleted_at FROM member m JOIN classroom c ON m.class_id = c.id WHERE c.deleted_at IS NOT NULL AND m.user_id = 3;
