export interface Hunt {
    hunt_id:  number;
    title: string;
    description: string;
    created_by: number;
    max_team_size: number;
}

export interface User {
    user_id: number;
    username: string;
    email: string;
    password: string;
}

export interface Clue {
    clue_id: number;
    name: string;
    text: string;
    // hunt_id: number;
    category: string;
    value: number;
    max_submissions: number;
}

export interface Team {
    team_id: number;
    name: string;
    hunt_id: number;
}

