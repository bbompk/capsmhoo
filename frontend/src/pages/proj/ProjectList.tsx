import { useState, useEffect } from "react";
import Swal from 'sweetalert2'

import { FormMode, ProjectFormModal } from "../../components/form/ProjectFormModal";
import { ProjectInterface } from "../../interfaces/ProjectInterface";
import { getAllProjects } from "../../service/ProjectService";
import ProjectDetailModal from "../../components/proj/ProjectDetailModal";

const ProjectList = () => {
    const [projectData, setProjectData] = useState<ProjectInterface[]|null>(
      [
        {
          id: '1',
          team_id: '1',
          professor_id: '1',
          name: 'Project 1',
          description: 'Project 1 description',
          status: 'open',
          label: 'Project 1 label',
        },
        {
          id: '2',
          team_id: '2',
          professor_id: '2',
          name: 'Project 2',
          description: 'Project 2 description',
          status: 'open',
          label: 'Project 2 label',
        },
        {
          id: '3',
          team_id: '3',
          professor_id: '3',
          name: 'Project 3',
          description: 'Project 3 description',
          status: 'open',
          label: 'Project 3 label',
        },
        {
          id: '4',
          team_id: '4',
          professor_id: '4',
          name: 'Project 4',
          description: 'Project 4 description',
          status: 'open',
          label: 'Project 4 label',
        },
        {
          id: '5',
          team_id: '5',
          professor_id: '5',
          name: 'Project 5',
          description: 'Project 5 description',
          status: 'open',
          label: 'Project 5 label',
        },
      ]
    );
    const [isProjectModalOpen, setIsProjectModalOpen] = useState(false);
    const [formMode, setFormMode] = useState<FormMode>('create');
    const [projectId, setProjectId] = useState<string>('');
    const [isProjectDetailModalOpen, setIsProjectDetailModalOpen] = useState(false);
    
    const role = localStorage.getItem("role")

    useEffect(() => {
      fetchAllProjects();
    }, []);

    useEffect(() => {
      if(isProjectModalOpen) return;
      fetchAllProjects();
    }, [isProjectModalOpen]);

    const fetchAllProjects = () => {
      try{
        getAllProjects().then((res) => {
          console.log(res);
          if(res.data === undefined)return;
          setProjectData(res.data);
        });
      } catch(err){
        console.log(err);
        Swal.fire("Error","Cannot get projects", 'error')
      }
    }

    const showProjectCreateModal = () => {
      setIsProjectModalOpen(true);
      setFormMode('create');
    }

    const showProjectEditModal = (projectId:string) => {
      setIsProjectModalOpen(true);
      setFormMode('edit');
      setProjectId(projectId);
    }

    const hideProjectModal = () => {
      setIsProjectModalOpen(false);
    }

    const showProjectDetailModal = (projectId:string) => {
      setProjectId(projectId);
      setIsProjectDetailModalOpen(true);
    }

    return (
      <div>
      <div className="min-h-screen">
        <h1 className=" text-center text-3xl p-4">List of Project</h1>
        <div className="container my-12 mx-auto px-4 md:px-12">
          {
            role === "Professor" &&
            <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" onClick={showProjectCreateModal}>
              Create Project
            </button>
          }
        <div className="flex flex-wrap -mx-1 lg:-mx-4">
          {projectData && projectData.map((project) => (
            <div className="my-1 px-1 w-full md:w-1/2 lg:w-1/3 cursor-pointer" onClick={()=>{showProjectDetailModal(project.id!)}}>
              <article className="overflow-hidden rounded-lg shadow-lg">
                <header className="flex items-center justify-between leading-tight p-2 md:p-4">
                  <h1 className="text-lg">
                    <a
                      className="no-underline hover:underline text-black"
                      href={`/detail/${project.id}`}
                      >
                      {project.name}
                    </a>
                  </h1>
                  <p className="text-grey-darker text-sm">
                    {project.status}
                  </p>
                  {
                    role === "Professor" &&
                    <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" onClick={(e)=>{e.stopPropagation();showProjectEditModal(project.id!);}}>
                      Edit
                    </button>
                  }  
                </header>

                <div className="flex items-center justify-between leading-none p-2 md:p-4">
                  <p className="ml-2 text-sm">
                    {project.label}
                  </p>
                </div>

                <footer className="flex items-center justify-between leading-none p-2 md:p-4">
                    <p className="ml-2 text-sm">
                      {project.description}
                    </p>
                </footer>
              </article>
            </div>
          ))}
        </div>
        <ProjectFormModal projectId={projectId} isModalVisible={isProjectModalOpen} setOpenModal={setIsProjectModalOpen} formMode={formMode}/>
        <ProjectDetailModal projectId={projectId} isModalVisible={isProjectDetailModalOpen} setOpenModal={setIsProjectDetailModalOpen}/>
      </div>
      </div>
      </div>
    );
  };
  
export default ProjectList;
  