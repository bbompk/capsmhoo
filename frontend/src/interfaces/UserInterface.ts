export interface UserInterface {
    id: string;
    email: string;
    password?: string;
    role: string;
}

export interface StudentInterface {
    id: string;
    name: string;
    team_id: string;
    user_id: string;
}

export interface ProfessorInterface {
    id: string;
    name: string;
    profile: string;
    user_id: string;
}