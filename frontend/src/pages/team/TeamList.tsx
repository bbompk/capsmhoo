import { useState, useEffect } from "react";
import { TeamInterface } from "../../interfaces/TeamInterface";
import Card from "../../components/card/Card";
import { getAllTeams } from "../../service/TeamService";
import Swal from "sweetalert2";

const TeamList = () => {
  const [data, setData] = useState<TeamInterface[]>();
  const fetchData = async () => {
    try {
      const teamRes = await getAllTeams();
      if (!teamRes.data) {
        Swal.fire("Error", "Cannot retrieve team data.");
        return;
      }
      setData(teamRes.data);
    } catch (err) {
      console.log(err);
      Swal.fire("Error", "Cannot get teams", "error");
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div>
      <div className="min-h-screen">
        <h1 className=" text-center text-3xl p-4">List of Team</h1>

        <div className="container my-12 mx-auto px-4 md:px-12">
          <div className="flex flex-wrap -mx-1 lg:-mx-4">
            {Array.isArray(data)
              ? data.map((team) => (
                  <Card
                    key={team.id}
                    id={team.id}
                    title={team.name}
                    body={team.profile}
                    next_path="team-detail"
                  />
                ))
              : null}
          </div>
        </div>
      </div>
    </div>
  );
};

export default TeamList;
