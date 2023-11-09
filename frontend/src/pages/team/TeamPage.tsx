import { useState, useEffect } from "react";
import { TeamInterface } from "../../interfaces/TeamInterface";
import { useParams } from "react-router-dom";

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

  return (
    <div>
      <div className="min-h-screen">
        <h1 className=" text-center text-3xl p-4">Team Page</h1>
        <h2>Name: {data?.name}</h2>
        <p>Profile: {data?.profile}</p>
        <form action="http://localhost:8082/team-join-request" method="POST">
          <input type="hidden" name="teamID" value={data?.id} />
          <button type="submit" className="btn btn-primary">
            Request to Join
          </button>
        </form>
      </div>
    </div>
  );
};

export default TeamPage;
