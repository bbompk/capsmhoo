import { useState, useEffect } from "react";
import axios from "axios";

const ProjectList = () => {
    const [projectData, setProjectData] = useState();
    const allBlog = () => {
      axios
        .get('url')
        .then(function (response) {
          setProjectData(response?.data?.data);
          console.log(response?.data?.data);
        })
        .catch(function (error) {
          // handle error
          //setLoading(false);
          //   setMessage(error?.response?.data?.message);
          //   openSnackbar(error?.response?.data?.message);
          console.log(error);
        });
    };
    useEffect(() => {
      allBlog();
    }, []);
    return (
      <div>
      <div className="min-h-screen">
  
        <h1 className=" text-center text-3xl p-4">List of Project</h1>
        
        <div className="container my-12 mx-auto px-4 md:px-12">
        <div className="flex flex-wrap -mx-1 lg:-mx-4">
          {projectData?.map((project) => (
            <div className="my-1 px-1 w-full md:w-1/2 lg:my-4 lg:px-4 lg:w-1/3">
              <article className="overflow-hidden rounded-lg shadow-lg">
                <a href={`/detail/${project.id}`}>
                  <img
                    alt="Placeholder"
                    className="block h-72 w-full"
                    src='../../../public/default-group-image.webp'
                  />
                </a>

                <header className="flex items-center justify-between leading-tight p-2 md:p-4">
                  <h1 className="text-lg">
                    <a
                      className="no-underline hover:underline text-black"
                      href={`/detail/${project.id}`}
                    >
                      {project.name}
                    </a>
                  </h1>
                </header>

                <footer className="flex items-center justify-between leading-none p-2 md:p-4">
                    <p className="ml-2 text-sm">
                      {project.description}
                    </p>
                </footer>
              </article>
            </div>
          ))}
        </div>
      </div>
        
      </div>
      </div>
    );
  };
  
export default ProjectList;
  