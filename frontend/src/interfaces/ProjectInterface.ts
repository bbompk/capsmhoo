export interface ProjectInterface {
    id?: string;
    name: string;
    description: string;
    professor_id: string;
    team_id?: string;
    status: string;
    label: string;
}

export interface ProjectRequestInterface {
    project_request_id?: string;
    project_id: string;
    team_id: string;
    message: string;
    status: string;
}

export interface ProjectRequestWithTeamInfoInterface {
    project_request_id?: string;
    project_id: string;
    team_id: string;
    message: string;
    status: string;
    team_name: string;
    team_profile: string;
}