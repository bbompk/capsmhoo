export interface ProjectInterface {
    project_id: string;
    name: string;
    description: string;
    professor_id: string;
    team_id: string;
}

export interface ProjectRequestInterface {
    project_request_id: string;
    project_id: string;
    team_id: string;
    message: string;
    status: string;
}