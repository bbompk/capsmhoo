import { useNavigate } from "react-router-dom";
import { addHoursToDate } from "../../utils/DateTimeUtil";

const MockLogin = () => {
    let navigate = useNavigate();

    let login = async () => {
        localStorage.setItem("accessToken", "test");
        localStorage.setItem("token_expires", addHoursToDate(new Date(), 1).toString())
        navigate("/");
    };
  
    return (
      <div className="grid w-screen place-items-center gap-8 pt-4">
        <p>You must log in</p>
        <button className="bg-gray-200 p-2 rounded-md hover:bg-gray-400" onClick={login}>Log in</button>
      </div>
    );
  }

export default MockLogin;