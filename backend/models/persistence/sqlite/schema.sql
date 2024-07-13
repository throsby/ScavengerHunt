CREATE TABLE User (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE Hunt (
    hunt_id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    created_by INTEGER,
    FOREIGN KEY (created_by) REFERENCES User (user_id)
);

CREATE TABLE Clue (
    clue_id INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT NOT NULL,
    hunt_id INTEGER,
    category TEXT,
    score INTEGER,
    FOREIGN KEY (hunt_id) REFERENCES Hunt (hunt_id)
);

CREATE TABLE Team (
    team_id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    hunt_id INTEGER,
    FOREIGN KEY (hunt_id) REFERENCES Hunt (hunt_id)
);

CREATE TABLE Submission (
    submission_id INTEGER PRIMARY KEY AUTOINCREMENT,
    team_id INTEGER,
    clue_id INTEGER,
    answer TEXT,
    score INTEGER,
    FOREIGN KEY (team_id) REFERENCES Team (team_id),
    FOREIGN KEY (clue_id) REFERENCES Clue (clue_id)
);
