CREATE TABLE "User" (
    user_id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE Hunt (
    hunt_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    created_by INTEGER,
    FOREIGN KEY (created_by) REFERENCES "User" (user_id)
);

CREATE TABLE Clue (
    clue_id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    hunt_id INTEGER,
    category TEXT,
    score INTEGER,
    FOREIGN KEY (hunt_id) REFERENCES Hunt (hunt_id)
);

CREATE TABLE Team (
    team_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    hunt_id INTEGER,
    FOREIGN KEY (hunt_id) REFERENCES Hunt (hunt_id)
);

CREATE TABLE Submission (
    submission_id SERIAL PRIMARY KEY,
    team_id INTEGER,
    clue_id INTEGER,
    answer TEXT,
    score INTEGER,
    FOREIGN KEY (team_id) REFERENCES Team (team_id),
    FOREIGN KEY (clue_id) REFERENCES Clue (clue_id)
);
