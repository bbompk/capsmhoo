import { useState, useEffect } from "react";
import { TeamJoinRequestInterface } from "../../interfaces/TeamInterface";
import { getAllTeamJoinRequests } from "../../service/TeamJoinRequestService";

const TeamJoinRequestList = () => {
  console.log(getAllTeamJoinRequests);
  const [data, setData] = useState<TeamJoinRequestInterface[]>();
  const fetchData = async () => {
    await fetch("http://localhost:8082/team").then(async (res) => {
      const response = await res.json();
      setData(response.data);
    });
  };

  useEffect(() => {
    fetchData();
  }, []);

  console.log(data);

  return (
    <div>
      <div className="min-h-screen">
        <h1 className=" text-center text-3xl p-4">List of Team Joining Request</h1>

        <div className="container my-12 mx-auto px-4 md:px-12">
          <div className="flex flex-wrap -mx-1 lg:-mx-4">
            {Array.isArray(data)
              ? data.map((teamReq) => (
                  <div className="my-1 px-1 w-full md:w-1/2 lg:my-4 lg:px-4 lg:w-1/3">
                    <article className="overflow-hidden rounded-lg shadow-lg">
                      <a href={`/detail/${teamReq.id}`}>
                        <img
                          alt="Placeholder"
                          className="block h-72 w-full"
                          src="../../../public/default-group-image.webp"
                        />
                      </a>

                      <header className="flex items-center justify-between leading-tight p-2 md:p-4">
                        <h1 className="text-lg">
                          <a
                            className="no-underline hover:underline text-black"
                            href={`/detail/${team.id}`}
                          >
                            {team.name}
                          </a>
                        </h1>
                      </header>

                      <footer className="flex items-center justify-between leading-none p-2 md:p-4">
                        <p className="ml-2 text-sm">{team.profile}</p>
                      </footer>
                    </article>
                  </div>
                ))
              : null}
          </div>
        </div>
      </div>
    </div>
  );
};

export default TeamJoinRequestList;
