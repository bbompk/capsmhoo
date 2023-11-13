import {
  UserInterface,
  StudentInterface,
  ProfessorInterface,
} from "./UserInterface";

export interface AuthContextInterface {
  isAuthenticated?: boolean;
  user?: UserInterface | StudentInterface | ProfessorInterface;
  loading?: boolean;
}
