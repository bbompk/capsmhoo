import { UserInterface } from "../interfaces/UserInterface";

// Not Ready To Use
export const userUser = () => {
    const user = JSON.parse(localStorage.getItem('user') ?? '{}') as UserInterface;
    return user;
}