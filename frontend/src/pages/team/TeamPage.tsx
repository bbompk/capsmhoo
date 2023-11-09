import { useState, useEffect } from "react";
import {
  TeamInterface,
  TeamJoinRequestInterface,
} from "../../interfaces/TeamInterface";
import { useNavigate, useParams } from "react-router-dom";
import { createTeamJoinRequest } from "../../service/TeamJoinRequestService";
import { useUser } from "../../hooks/useUser";
import Swal from "sweetalert2";
import 'bootstrap/dist/css/bootstrap.css';

const TeamPage = () => {
  const { id } = useParams();
  const { userId, role } = useUser();
  const [data, setData] = useState<TeamInterface>();

  const navigate = useNavigate();

  const fetchData = async () => {
    if (!id) {
      Swal.fire("Cannot retrive this team data");
      navigate("/view-team");
      return;
    }

    await fetch("http://localhost:8082/team/" + id).then(async (res) => {
      const response = await res.json();
      setData(response.data);
    });
  };

  useEffect(() => {
    fetchData();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    if (role !== "Student") {
      Swal.fire({
        icon: "error",
        title: "Only student can join a team.",
        text: "Only student can join a team. Professor can only view teams.",
      });
    }

    if (!data) {
      throw new Error("Failed to fetch team data");
    }
    if (!userId) {
      throw new Error("Failed to fetch user data");
    }
    const teamJoinRequest: TeamJoinRequestInterface = {
      id: "",
      team_id: data.id,
      student_id: userId,
    };

    e.preventDefault();
    try {
      await createTeamJoinRequest(teamJoinRequest);
    } catch (error) {
      console.error(error);
      return;
    }
  };

  return (
    <div>
      <div className="min-h-screen">
        <h1 className=" text-center text-3xl p-4">Team Page</h1>
        <h2>Name: {data?.name}</h2>
        <p>Profile: {data?.profile}</p>
        <form onSubmit={handleSubmit}>
          <div className="btn btn-primary">
            <button type="submit">Request to Join</button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default TeamPage;
