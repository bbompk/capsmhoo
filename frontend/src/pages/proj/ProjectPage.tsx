import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { ProjectInterface } from "../../interfaces/ProjectInterface";

const ProjectPage = () => {
  const { id } = useParams();
  const [data, setData] = useState<ProjectInterface>();
  const fetchData = async () => {
    await fetch("http://localhost:8082/project/" + id).then(async (res) => {
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
        <h3>Professor ID: {data?.professor_id}</h3>
        <p>Description: {data?.description}</p>
      </div>
    </div>
  );
};

export default ProjectPage;
