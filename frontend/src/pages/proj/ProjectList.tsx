import { useState, useEffect } from "react";
import { ProjectInterface } from "../../interfaces/ProjectInterface";
import Card from "../../components/card/Card";

const ProjectList = () => {
    const [projectData, setProjectData] = useState<ProjectInterface>();
    const fetchData = async () => {
      await fetch("http://localhost:8082/project")
        .then(async (res) => {
          const response = await res.json();
          setProjectData(response.data)
        })
    };
  
    useEffect(() => {
      fetchData();
    }, []);

    return (
      <div>
      <div className="min-h-screen">
  
        <h1 className=" text-center text-3xl p-4">List of Project</h1>
        
        <div className="container my-12 mx-auto px-4 md:px-12">
        <div className="flex flex-wrap -mx-1 lg:-mx-4">
        {Array.isArray(projectData)? 
              projectData.map((project) => (
                  <Card key={project.id} id={project.id} title={project.name} body={project.description} next_path='project-detail' />
              ))
              : null
            }
        </div>
      </div>
        
      </div>
      </div>
    );
  };
  
export default ProjectList;
  