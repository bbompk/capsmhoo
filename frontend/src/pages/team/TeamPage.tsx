import { useState, useEffect } from "react";
import { TeamInterface, TeamJoinRequestInterface } from "../../interfaces/TeamInterface";
import { useParams } from "react-router-dom";
import { createTeamJoinRequest } from "../../service/TeamJoinRequestService";

const TeamPage = () => {
  const { id } = useParams();
  const [data, setData] = useState<TeamInterface>();
  const fetchData = async () => {
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
    const teamJoinRequest : TeamJoinRequestInterface = {
      id: "",
      team_id: JSON.stringify(data?.id),
      student_id: ""
    }
    
    e.preventDefault();
    try {
        await createTeamJoinRequest(teamJoinRequest);
    } catch (error) {
      console.error(error);
      return;
    }
  }

  return (
    <div>
      <div className="min-h-screen">
        <h1 className=" text-center text-3xl p-4">Team Page</h1>
        <h2>Name: {data?.name}</h2>
        <p>Profile: {data?.profile}</p>
        <form onSubmit={handleSubmit}>
          <button type="submit" className="btn btn-primary">
            Request to Join
          </button>
        </form>
      </div>
    </div>
  );
};

export default TeamPage;
