
-- CREATE TABLE houses (
--     house_name TEXT NOT NULL,
--     head TEXT NOT NULL,
--     PRIMARY KEY(house_name)
-- );

CREATE TABLE students (
    id INTEGER,
    name TEXT NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE houses (
    id INTEGER,
    house_name TEXT NOT NULL,
    head TEXT NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE assignment (
    student_id INTEGER,
    house_id INTEGER,
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (house_id) REFERENCES houses(id)
);

DROP TABLE houses;
DROP TABLE assignment;

SELECT house, head FROM students;

/* Inserting new rows with data provided by a SELECT statement */
INSERT INTO houses (house_name,head) SELECT house, head FROM (SELECT DISTINCT house, head FROM students);

DELETE FROM houses WHERE id = 2 OR id = 3;

SELECT house, head INTO newTable FROM (SELECT DISTINCT house, head FROM students);

/* prepare to copy to TABLE assignment */
SELECT students.id, houses.id FROM students INNER JOIN houses ON students.house = houses.house_name;

INSERT INTO assignment (student_id, house_id) SELECT st_id, ho_id FROM (SELECT students.id AS st_id, houses.id AS ho_id FROM students INNER JOIN houses ON students.house = houses.house_name);

/* delete column */
ALTER TABLE students DROP COLUMN house;
ALTER TABLE students DROP COLUMN head;


/* Manipulate after create successfully */
SELECT student_id, house_id, house_name, head FROM assignment JOIN houses ON assignment.house_id = houses.id;

/* recreate the origin table by joining 3 new ones */
SELECT student_name, student_id, house_id, house_name, head FROM
students JOIN (SELECT student_id, house_id, house_name, head FROM assignment JOIN houses ON assignment.house_id = houses.id) AS houses1
ON students.id = houses1.student_id;


/********************************************************************/
/*                       Practice with Golang                       */
/********************************************************************/

/* import csv file to sqlite3 database */
/* add column to a table */
ALTER TABLE students ADD COLUMN house;
ALTER TABLE students ADD COLUMN head;
ALTER TABLE students RENAME COLUMN id TO id TEXT NOT NULL;


