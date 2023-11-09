import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { createProject } from "../../service/ProjectService";
import { ProjectInterface } from "../../interfaces/ProjectInterface";

export default function CreateProjectForm() {
    const [professorId,setProfessorId] = useState("")
    const [name, setName] = useState("");
    const [description, setDescription] = useState("");
    
    const navigate = useNavigate();

  const resetForm = () => {
    setProfessorId("");
    setName("");
    setDescription("");
  };
  
  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    const project : ProjectInterface = {
        name: name,
        description: description,
        professor_id: professorId,
        project_id: "",
        team_id: "",
    }
    
    e.preventDefault();
    try {
        await createProject(project);
    } catch (error) {
      console.error(error);
      resetForm();
      return;
    }
    navigate("/view-project");
  }
  return (
    <form onSubmit={handleSubmit}>
        <div className="space-y-12">
            <div className="border-b border-gray-900/10 pb-12">
                <div className="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
                <div className="sm:col-span-4">
                    <label htmlFor="professorId" className="block text-sm font-medium leading-6 text-gray-900">
                        Professer ID
                    </label>
                    <div className="mt-2">
                        <div className="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 focus-within:ring-2 focus-within:ring-inset focus-within:ring-indigo-600 sm:max-w-md">
                            <input
                                type="text"
                                name="professorId"
                                id="professorId"
                                required
                                value={professorId}
                                onChange={(e) => setName(e.target.value)}
                                className="block flex-1 border-0 bg-transparent py-1.5 pl-1 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6"
                            />
                        </div>
                    </div>
                </div>

                <div className="sm:col-span-4">
                    <label htmlFor="name" className="block text-sm font-medium leading-6 text-gray-900">
                        Project Name
                    </label>
                    <div className="mt-2">
                        <div className="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 focus-within:ring-2 focus-within:ring-inset focus-within:ring-indigo-600 sm:max-w-md">
                        <input
                            type="text"
                            name="name"
                            id="name"
                            required
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                            className="block flex-1 border-0 bg-transparent py-1.5 pl-1 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6"
                            placeholder="project name"
                        />
                        </div>
                    </div>
                </div>

                <div className="col-span-full">
                    <label htmlFor="description" className="block text-sm font-medium leading-6 text-gray-900">
                        Description
                    </label>
                    <div className="mt-2">
                        <textarea
                            id="description"
                            name="description"
                            value={description}
                            onChange={(e) => setName(e.target.value)}
                            rows={3}
                            className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                            defaultValue={''}
                            placeholder="Project Description"
                        />
                    </div>
                </div>

                <div className="mt-6 flex items-center justify-end gap-x-6">
                    <button
                        type="submit"
                        className="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                    >
                    Submit
                    </button>
                </div>
                </div>
            </div>
        </div>
    </form>
  )
}
