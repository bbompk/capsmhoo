import { useEffect, useState } from "react";
import { useUser } from "../../hooks/useUser";
import Swal from "sweetalert2";
import { TeamJoinRequestInterface } from "../../interfaces/TeamInterface";
import { useNavigate } from "react-router-dom";
import {
  approveStudentIntoTeam,
  deleteTeamJoinRequestById,
  getAllTeamJoinRequestByTeamId,
} from "../../service/TeamJoinRequestService";
import JoiningRequestCard from "../../components/card/JoiningRequestCard";
import { getTeamByUserId } from "../../service/TeamService";

const TeamJoinRequest = () => {
  const { userId, role } = useUser();
  const [data, setData] = useState<TeamJoinRequestInterface[]>();
  const navigate = useNavigate();

  useEffect(() => {
    const fetchData = async () => {
      if (!userId) {
        Swal.fire("Please log in to view your team joining request.");
        navigate("/login");
        return;
      }

      if (role !== "Student") {
        Swal.fire("Only student can view team joining request.");
        navigate("/");
        return;
      }

      const team = await getTeamByUserId(userId);
      if (!team.data) {
        Swal.fire("This student is not in a team.");
        navigate("/view-team");
        return;
      }

      const teamJoinRequest = await getAllTeamJoinRequestByTeamId(team.data.id);
      if (!teamJoinRequest) {
        Swal.fire("Failed to load team joining request data");
        navigate("/");
        return;
      }
      setData(teamJoinRequest.data);
    };
    fetchData();
  },[]);

  const handleAccept = async (id: string) => {
    approveStudentIntoTeam(id);
    deleteTeamJoinRequestById(id);
    window.location.reload();
  };

  const handleReject = async (id: string) => {
    deleteTeamJoinRequestById(id);
    window.location.reload();
  };

  return (
    <div>
      <div className="min-h-screen">
        <h1 className=" text-center text-3xl p-4">
          Your Team Joining Requests
        </h1>
        <div className="container my-12 mx-auto px-4 md:px-12">
          <div className="list-group">
            {Array.isArray(data)
              ? data.map((item) => (
                  <JoiningRequestCard
                    key={item.id}
                    id={item.id}
                    student_id={item.student_id}
                    handleAccept={() => handleAccept(item.id)}
                    handleReject={() => handleReject(item.id)}
                  />
                ))
              : null}
          </div>
        </div>
      </div>
    </div>
  );
};

export default TeamJoinRequest;
