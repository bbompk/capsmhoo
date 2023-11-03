import React, { useContext, createContext, useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import Swal from 'sweetalert2'
import { UserInterface } from "../../interfaces/UserInterface";
import { addHoursToDate } from "../../utils/DateTimeUtil";

export interface AuthContextInterface {
  isAuthenticated?: boolean,
  user?: UserInterface,
  loading?: boolean
}

const authContext = createContext<AuthContextInterface>({});

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [loading, setLoading] = useState<boolean>(true)
  const [user, setUser] = useState<UserInterface | undefined>();
  const navigate = useNavigate();

  useEffect(() => {
    setLoading(true)
    let interval: string | number | NodeJS.Timeout | undefined;
    const accessToken = localStorage.getItem('accessToken')
    if(accessToken == null){
        navigate("/login");
        clearInterval(interval);
        Swal.fire("Error","Please log in", 'error')
        return;
    } else {
        setUser({ id: "test" } as UserInterface)
    }

    interval = setInterval(() => {
        const expiredAt = new Date(localStorage.getItem("token_expires") ?? addHoursToDate(new Date(), -1))
        const now = new Date();
        if (now > expiredAt) {
            ForceLogout();
        }
    }, 1000);
    
    function ForceLogout() {
        clearInterval(interval);
        localStorage.removeItem('accessToken')
        localStorage.removeItem('token_expires')
        navigate("/login");

        Swal.fire("Error","Please log in", 'error')
    }
    setLoading(false)
    return () => clearInterval(interval);
  }, [])

  return <authContext.Provider value={{ isAuthenticated: !!user, user, loading}}>{children}</authContext.Provider>;
}

export const useAuth = () => useContext(authContext);







