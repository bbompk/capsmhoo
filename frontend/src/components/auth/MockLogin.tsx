import { useNavigate } from "react-router-dom";
import { addHoursToDate } from "../../utils/DateTimeUtil";

const MockLogin = () => {
    let navigate = useNavigate();

    let login = async () => {
      sessionStorage.setItem("accessToken", "test");
      sessionStorage.setItem("token_expires", addHoursToDate(new Date(), 1).toString())
      navigate("/");
    };

    let logout = async () => {
      sessionStorage.removeItem('accessToken')
      sessionStorage.removeItem('token_expires')
      navigate("/login");
    }
  
    return (
      <div className="grid w-screen place-items-center gap-8 pt-4">
        <p>You must log in</p>
        <button className="bg-gray-200 p-2 rounded-md hover:bg-gray-400" onClick={login}>Log in</button>
        <button className="bg-gray-200 p-2 rounded-md hover:bg-gray-400" onClick={logout}>Log out</button>
      </div>
    );
  }

export default MockLogin;