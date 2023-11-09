import { useState, useEffect } from "react";
import { getAllTeams } from "../../service/TeamService";
import { TeamInterface } from "../../interfaces/TeamInterface";
import Card from "../../components/card/Card";

const TeamList = () => {
  console.log(getAllTeams);
  const [data, setData] = useState<TeamInterface[]>();
  const fetchData = async () => {
    await fetch("http://localhost:8082/team")
      .then(async (res) => {
        const response = await res.json();
        setData(response.data)
      })
  };

  useEffect(() => {
    fetchData();
  }, []);

  console.log(data)

  return (
    <div>
      <div className="min-h-screen">
        <h1 className=" text-center text-3xl p-4">List of Team</h1>

        <div className="container my-12 mx-auto px-4 md:px-12">
          <div className="flex flex-wrap -mx-1 lg:-mx-4">
            {Array.isArray(data)? 
              data.map((team) => (
                  <Card key={team.id} id={team.id} title={team.name} body={team.profile} next_path='team-detail' />
              ))
              : null
            }
          </div>
        </div>
      </div>
    </div>
  );
};

export default TeamList;
