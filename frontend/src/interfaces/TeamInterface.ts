export interface TeamInterface {
    id: string;
    name: string;
    profile: string;
}

export interface TeamCreateInterface {
    id: string;
    name: string;
    profile: string;
    creator_id: string;
}

export interface TeamJoinRequestInterface {
    id: string;
    team_id: string;
    student_id: string;
}