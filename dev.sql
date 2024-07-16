-- define a select statement to get all students enrolled in a course
-- write an index to improve the performance of the query
CREATE INDEX idx_students_location ON courses.students (city, state);

-- create an index on registration_date column
CREATE INDEX idx_registrations_registration_date ON courses.registrations (registration_date);

-- define a stored procedure to get course enrollment by location
CREATE PROCEDURE courses.GetCourseEnrollmentByLocation
    @city VARCHAR(100),
    @state VARCHAR(100)
AS
BEGIN
    SELECT COUNT(*) AS enrollment
    FROM courses.students s
    JOIN courses.registrations r ON s.student_id = r.student_id
    JOIN courses.courses c ON r.course_id = c.course_id
    WHERE s.city = @city AND s.state = @state;
END;

-- select registrations for September 2023
SELECT *
FROM courses.registrations
WHERE YEAR(registration_date) = 2023 AND MONTH(registration_date) = 9;