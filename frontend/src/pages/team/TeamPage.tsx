import { useState, useEffect } from "react";
import {
  TeamInterface,
  TeamJoinRequestInterface,
} from "../../interfaces/TeamInterface";
import { useNavigate, useParams } from "react-router-dom";
import { createTeamJoinRequest, getAllTeamJoinRequestByStudentId } from "../../service/TeamJoinRequestService";
import { useUser } from "../../hooks/useUser";
import Swal from "sweetalert2";
import 'bootstrap/dist/css/bootstrap.css';
import { getStudentByUserId } from "../../service/StudentService";
import { getTeamById } from "../../service/TeamService";

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

    try {
      const teamDetailRes = await getTeamById(id);
      if (!teamDetailRes.data) {
        Swal.fire("Failed to load team data");
        navigate("/view-team");
        return;
      }
      setData(teamDetailRes.data);
    }
    catch (err) {
      console.log(err);
      Swal.fire("Error", "Cannot get this team", 'error')
    }
  };

  useEffect(() => {
    fetchData();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (role !== "Student") {
      Swal.fire({
        icon: "error",
        title: "Only student can join a team.",
        text: "Only student can join a team. Professor can only view teams.",
      });
      return;
    }

    if (!data) {
      throw new Error("Failed to fetch team data.");
    }
    if (!userId) {
      throw new Error("Failed to fetch user data.");
    }

    const studentRes = await getStudentByUserId(userId);
    if (!studentRes.data) {
      throw new Error("Cannot find student ID.");
    }
    const student_id = studentRes.data.id;

    const teamJoinRes = await getAllTeamJoinRequestByStudentId(student_id)

    if (teamJoinRes.code !== '200') {
      throw new Error("Cannot retrieve Team Join Request DB data.");
    }
    if (teamJoinRes.data !== null) {
      Swal.fire({
        icon: "error",
        title: "You already requested to join another team.",
        text: "Only one team joining request is allowed for a student. Please contact the team you previously sent a request.",
      });
      return;
    }

    try {
      const teamJoinRequest: TeamJoinRequestInterface = {
        id: "",
        team_id: data.id,
        student_id: student_id
      };
      await createTeamJoinRequest(teamJoinRequest);
    }
    catch (error) {
      console.error(error);
      Swal.fire({
        icon: "error",
        title: "Request to Join Failed",
        text: "Please try again",
      });
      return;
    }
    Swal.fire({
      icon: "success",
      title: "Request Created",
      text: `You requested to join team ${data.name}`,
    });
    navigate("/"); // Navigate to the home page
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
