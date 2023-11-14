import { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { ProjectInterface } from "../../interfaces/ProjectInterface";
import Swal from "sweetalert2";
import { getProjectById } from "../../service/ProjectService";

const ProjectPage = () => {
  const { id } = useParams();
  const [data, setData] = useState<ProjectInterface>();
  const navigate = useNavigate();
  const fetchData = async () => {
    if (!id) {
      Swal.fire("Cannot retrive this project data");
      navigate("/view-project");
      return;
    }

    try {
      const projDetailRes = await getProjectById(id);
      if (!projDetailRes.data) {
        Swal.fire("Failed to load project data");
        navigate("/view-project");
        return;
      }
      setData(projDetailRes.data);
    } catch (err) {
      console.log(err);
      Swal.fire("Error", "Cannot get this project", "error");
    }
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
