-- Keep a log of any SQL queries you execute as you solve the mystery.

-- The THIEF is: Bruce
-- The city the thief ESCAPED TO: New York City
-- The ACCOMPLICE is: Kenny

select count(name) from people;                     -- 200
select count(id) from crime_scene_reports;          -- 301
select count(id) from interviews;                   -- 192
select count(id) from atm_transactions;             -- 1332
select count(person_id) from bank_accounts;         -- 137 account <=> person_id
select count(id) from airports;                     -- 12
select count(id) from flights;                      -- 58
select count(id) from phone_calls;                  -- 518
select count(id) from bakery_security_logs;         -- 468


/* See description of the scene */
SELECT description, day, month, year, street FROM crime_scene_reports where street = 'Humphrey Street';

+--------------------------------------------------------------+-----+-------+------+-----------------+
| Theft of the CS50 duck took place at 10:15am at the Humphrey | 28  | 7     | 2021 | Humphrey Street |
|  Street bakery. Interviews were conducted today with three w |     |       |      |                 |
| itnesses who were present at the time - each of their interv |     |       |      |                 |
| iew transcripts mentions the bakery.                         |     |       |      |                 |
+--------------------------------------------------------------+-----+-------+------+-----------------+

/* look for transcript of interviews */
SELECT name, transcript FROM interviews WHERE year = 2021 and month = 7 and day = 28;


/* look for ATM transaction + bank_accounts + name + phone number + passport number + lilcence-plate */

SELECT name, phone_number, passport_number, license_plate FROM people,
(    
    SELECT person_id FROM bank_accounts INNER JOIN
    (
        SELECT account_number, amount FROM atm_transactions
        WHERE atm_location = 'Leggett Street' and transaction_type = 'withdraw' and year = 2021 and month = 7 and day = 28
    ) AS ATM ON bank_accounts.account_number = ATM.account_number
) AS accounts
WHERE people.id = accounts.person_id;

+---------+----------------+-----------------+---------------+
|  name   |  phone_number  | passport_number | license_plate |
+---------+----------------+-----------------+---------------+
| Bruce   | (367) 555-5533 | 5773159633      | 94KL13X       |
| Diana   | (770) 555-1861 | 3592750733      | 322W7JE       |
| Brooke  | (122) 555-4581 | 4408372428      | QX4YZN3       |
| Kenny   | (826) 555-1652 | 9878712108      | 30G67EN       |
| Iman    | (829) 555-5269 | 7049073643      | L93JTIZ       |
| Luca    | (389) 555-5198 | 8496433585      | 4328GD8       |
| Taylor  | (286) 555-6063 | 1988161715      | 1106N58       |
| Benista | (338) 555-6650 | 9586786673      | 8X428L0       |
+---------+----------------+-----------------+---------------+

/* look for phone call from the thetf */
SELECT caller, receiver FROM phone_calls where year = 2021 and month = 7 and day = 28 and duration < 60;

/* match name and passport_number of CALLER vs the theft who withdrawed money at ATM */
SELECT name, passport_number, license_plate FROM phone_calls INNER JOIN
(    
    SELECT name, phone_number, passport_number, license_plate FROM people,
    (    
        SELECT person_id FROM bank_accounts INNER JOIN
        (
            SELECT account_number, amount FROM atm_transactions
            WHERE atm_location = 'Leggett Street' and transaction_type = 'withdraw' and year = 2021 and month = 7 and day = 28
        ) AS ATM ON bank_accounts.account_number = ATM.account_number
    ) AS accounts
    WHERE people.id = accounts.person_id
) AS theftInfors
ON phone_calls.caller = theftInfors.phone_number WHERE duration < 60;

/* match name and passport_number of RECEIVER */
SELECT name, passport_number, license_plate FROM phone_calls INNER JOIN
(    
    SELECT name, phone_number, passport_number, license_plate FROM people,
    (    
        SELECT person_id FROM bank_accounts INNER JOIN
        (
            SELECT account_number, amount FROM atm_transactions
            WHERE atm_location = 'Leggett Street' and transaction_type = 'withdraw' and year = 2021 and month = 7 and day = 28
        ) AS ATM ON bank_accounts.account_number = ATM.account_number
    ) AS accounts
    WHERE people.id = accounts.person_id
) AS theftInfors
ON phone_calls.receiver = theftInfors.phone_number WHERE duration < 60;

              /* caller list */                             /* receiver list */
+---------+-----------------+---------------+       +-------+-----------------+---------------+
|  name   | passport_number | license_plate |       | name  | passport_number | license_plate |
+---------+-----------------+---------------+       +-------+-----------------+---------------+
| Bruce   | 5773159633      | 94KL13X       |       | Kenny | 9878712108      | 30G67EN       |
| Bruce   | 5773159633      | 94KL13X       |       | Kenny | 9878712108      | 30G67EN       |
| Diana   | 3592750733      | 322W7JE       |       +-------+-----------------+---------------+
| Kenny   | 9878712108      | 30G67EN       |
| Taylor  | 1988161715      | 1106N58       |
| Benista | 9586786673      | 8X428L0       |
+---------+-----------------+---------------+


/* matching licence-plate of the theft at bakery parking lot */
SELECT callers.name, callers.passport_number, licensePlates.hour AS hour, licensePlates.minute AS minute, activity FROM
(SELECT activity,license_plate, hour, minute, activity FROM bakery_security_logs WHERE year = 2021 and month = 7 and day = 28) AS licensePlates,
(    
    SELECT name, passport_number, license_plate FROM phone_calls INNER JOIN
    (    
        SELECT name, phone_number, passport_number, license_plate FROM people,
        (    
            SELECT person_id FROM bank_accounts INNER JOIN
            (
                SELECT account_number, amount FROM atm_transactions
                WHERE atm_location = 'Leggett Street' and transaction_type = 'withdraw' and year = 2021 and month = 7 and day = 28
            ) AS ATM ON bank_accounts.account_number = ATM.account_number
        ) AS accounts
        WHERE people.id = accounts.person_id
    ) AS theftInfors
    ON phone_calls.caller = theftInfors.phone_number WHERE duration < 60 and phone_calls.year = 2021 and phone_calls.month = 7 and phone_calls.day = 28
) AS callers
WHERE callers.license_plate = licensePlates.license_plate;

/* Inspect the flitght of caller - thief */

SELECT abbreviation, full_name, city FROM airports,
(SELECT origin_airport_id, destination_airport_id FROM flights WHERE year = 2021 and month = 7 and day = 29 ORDER BY hour ASC LIMIT 1)
WHERE airports.id = destination_airport_id;

/* destination of the theft */
+--------------+-------------------+---------------+
| abbreviation |     full_name     |     city      |
+--------------+-------------------+---------------+
| LGA          | LaGuardia Airport | New York City |
+--------------+-------------------+---------------+


/* Flight_id & passengers + infor of the thief at bakery_securety_logs => the thief */
SELECT bakery.name, bakery.hour, bakery.minute FROM
(    
    SELECT passport_number FROM passengers,
    (SELECT id FROM flights WHERE year = 2021 and month = 7 and day = 29 ORDER BY hour ASC LIMIT 1)
    WHERE passengers.flight_id = id
) AS passengers,
(    
    SELECT callers.name, callers.passport_number, licensePlates.hour AS hour, licensePlates.minute AS minute, activity FROM
    (SELECT activity,license_plate, hour, minute, activity FROM bakery_security_logs WHERE year = 2021 and month = 7 and day = 28) AS licensePlates,
    (    
        SELECT name, passport_number, license_plate FROM phone_calls INNER JOIN
        (    
            SELECT name, phone_number, passport_number, license_plate FROM people,
            (    
                SELECT person_id FROM bank_accounts INNER JOIN
                (
                    SELECT account_number, amount FROM atm_transactions
                    WHERE atm_location = 'Leggett Street' and transaction_type = 'withdraw' and year = 2021 and month = 7 and day = 28
                ) AS ATM ON bank_accounts.account_number = ATM.account_number
            ) AS accounts
            WHERE people.id = accounts.person_id
        ) AS theftInfors
        ON phone_calls.caller = theftInfors.phone_number WHERE duration < 60 and phone_calls.year = 2021 and phone_calls.month = 7 and phone_calls.day = 28
    ) AS callers
    WHERE callers.license_plate = licensePlates.license_plate
) AS bakery
WHERE bakery.passport_number = passengers.passport_number;


SELECT name, passport_number, license_plate FROM phone_calls INNER JOIN
(    
    SELECT name, phone_number, passport_number, license_plate FROM people,
    (    
        SELECT person_id FROM bank_accounts INNER JOIN
        (
            SELECT account_number, amount FROM atm_transactions
            WHERE atm_location = 'Leggett Street' and transaction_type = 'withdraw' and year = 2021 and month = 7 and day = 28
        ) AS ATM ON bank_accounts.account_number = ATM.account_number
    ) AS accounts
    WHERE people.id = accounts.person_id
) AS theftInfors
ON phone_calls.caller = theftInfors.phone_number WHERE duration < 60;