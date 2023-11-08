import { useState, useEffect, useCallback } from "react"
import Swal from 'sweetalert2'
import { Modal } from "antd";

import { getProjectById } from "../../service/ProjectService";
import { ProjectRequestInterface } from "../../interfaces/ProjectInterface";
import { createProjectRequest, getAllProjectRequestByProjectId, acceptProjectRequestById, rejectProjectRequestById } from "../../service/ProjectRequestService";
import { getStudentByUserId } from "../../service/StudentService";

const ProjectDetailModal = ({projectId, isModalVisible, setOpenModal}: {projectId:string, isModalVisible:boolean, setOpenModal:React.Dispatch<React.SetStateAction<boolean>>}) => {
    const [projectRequestData, setProjectRequestData] = useState<ProjectRequestInterface[]|null>([]);
    const [projectName, setProjectName] = useState<string>('');
    const [projectDescription, setProjectDescription] = useState<string>('');
    const [projectLabel, setProjectLabel] = useState<string>('Label');
    const [projectStatus, setProjectStatus] = useState<string>('open');
    const [projectTeamId, setProjectTeamId] = useState<string>('');
    const [teamId, setTeamId ] = useState<string>('');

    const role = localStorage.getItem("role")
    const userId = localStorage.getItem("userId")

    const fetchProjectData = useCallback(async () => {
        const projectRes = await getProjectById(projectId)
        console.log(projectRes.data)
        if(!projectRes.data) return;
        setProjectName(projectRes.data.name);
        setProjectDescription(projectRes.data.description);
        setProjectLabel(projectRes.data.label);
        setProjectStatus(projectRes.data.status);
        setProjectTeamId(projectRes.data.team_id || '');

        const projectRequestRes = await getAllProjectRequestByProjectId(projectId)
        if(projectRequestRes.code !== '200') return;
        setProjectRequestData(projectRequestRes.data||[]);
        console.log(projectRequestRes.data)
    }, [projectId])

    useEffect(() => {
        if(!isModalVisible) return;
        if(!userId) return;

        fetchProjectData()
            .catch((err) => {
                console.log(err);
                Swal.fire("Error","Cannot get project", 'error')
            })
        
        if(role === 'Student') {
            getStudentByUserId(userId)
                .then((res) => {
                    if(!res.data) return;
                    if(!res.data.team_id) return;
                    setTeamId(res.data.team_id);
                })
                .catch((err) => {
                    console.log(err);
                    Swal.fire("Error","Cannot get student", 'error')
                })
        }
    }, [fetchProjectData, isModalVisible])


    const handleClose = () => {
        setOpenModal(false);
    }

    const handleRequestProject = async () => {
        if(role !== 'Student') return;
        if(!userId) return;
        if(!teamId) {
            Swal.fire("Error","You must be in a team to request a project", 'error')
            return;
        }

        try {
            const projectRequestRes = await createProjectRequest({project_id: projectId, team_id: teamId, message: 'hello', status: 'pending'})

            if(projectRequestRes.code != '200'){
                Swal.fire("Error","Cannot request project", 'error')
                console.error(projectRequestRes);
                return;
            }
            await fetchProjectData();
            Swal.fire("Success","Project requested", 'success')
        } catch (error) {
            console.log(error);
            Swal.fire("Error","Cannot request project", 'error')
            return;
        }
    }

    async function handleAcceptProjectRequest(projectRequestId:string) {
        if(role !== 'Professor') return;
        if(!userId) return;
        if(!projectRequestId) return;

        try {
            const projectRequestRes = await acceptProjectRequestById(projectRequestId)

            if(projectRequestRes.code != '200'){
                Swal.fire("Error","Cannot accept project request", 'error')
                console.error(projectRequestRes);
                return;
            }
            await fetchProjectData();
            Swal.fire("Success","Project request accepted", 'success')
        } catch (error) {
            console.log(error);
            Swal.fire("Error","Cannot accept project request", 'error')
            return;
        }
    }

    async function handleRejectProjectRequest(projectRequestId:string) {
        if(role !== 'Professor') return;
        if(!userId) return;
        if(!projectRequestId) return;

        try {
            const projectRequestRes = await rejectProjectRequestById(projectRequestId)

            if(!projectRequestRes.data){
                Swal.fire("Error","Cannot reject project request", 'error')
                return;
            }
            await fetchProjectData();
            Swal.fire("Success","Project request rejected", 'success')
        } catch (error) {
            console.log(error);
            Swal.fire("Error","Cannot reject project request", 'error')
            return;
        }
    }

    const projectRequestStatus = projectRequestData?.find((projectRequest) => projectRequest.team_id === teamId)?.status
    const pendingProjectRequests = projectRequestData?.filter((projectRequest) => projectRequest.status === 'pending')

    return (
        <Modal title="Project Details" open={isModalVisible} footer={null} onCancel={handleClose}>
            <div className="space-y-12">
            <div className="pb-4">
                <div className="grid grid-cols-1 gap-x-6 gap-y-4 sm:grid-cols-6">
                <div className="sm:col-span-4">
                    <label htmlFor="name" className="block text-sm font-medium leading-6 text-gray-900">
                        Project Name
                    </label>
                    <div className="mt-2">
                        {projectName}
                    </div>
                </div>
                {projectStatus}

                <div className="col-span-full">
                    <label htmlFor="description" className="block text-sm font-medium leading-6 text-gray-900">
                        Description
                    </label>
                    <div className="mt-2">
                        {projectDescription}
                    </div>
                </div>

                <div className="col-span-full">
                    <label htmlFor="label" className="block text-sm font-medium leading-6 text-gray-900">
                        Label
                    </label>
                    <div className="mt-2">
                        {projectLabel}
                    </div>
                </div>
                </div>
            </div>
            {
                // Student request project button
                role === 'Student' && (
                    {
                        "pending" : <button type="button" className="flex justify-center items-center px-4 py-2 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-yellow-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 w-full" disabled>
                                        Pending...
                                    </button>,
                        "accepted" : <button type="button" className="flex justify-center items-center px-4 py-2 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-green-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 w-full" disabled>
                                        Accepted!
                                    </button>,
                        "rejected" : <button type="button" className="flex justify-center items-center px-4 py-2 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-red-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 w-full" disabled>
                                        Rejected!
                                    </button>
                    }[projectRequestStatus!] ||
                    <button type="button" className="flex justify-center items-center px-4 py-2 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-blue-500 hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 w-full" onClick={handleRequestProject}>
                        Request Project
                    </button>
                )
            }
            {
                projectStatus === 'closed' &&
                <div className="border-b border-gray-900/10 pb-4">
                    <h1 className="text-lg">Team Id</h1>
                    <div className="flex flex-wrap -mx-1 lg:-mx-4">
                        <div className="my-1 px-1 w-full">
                            <header className="flex items-center justify-between leading-tight gap-x-2 p-2 md:p-4 ">
                            <h1 className="text-lg">
                                {projectTeamId}
                            </h1>
                            </header>
                        </div>
                    </div>
                </div>
            }
            {
                // Professor's project request list
                role === 'Professor' && projectStatus === 'open' &&
                <div className="border-b border-gray-900/10 pb-4">
                    <h1 className="text-lg">Project Requests</h1>
                    <div className="flex flex-wrap -mx-1 lg:-mx-4">
                    {pendingProjectRequests && (pendingProjectRequests.length)>0 
                    ? pendingProjectRequests.map((projectRequest) => (
                        <div className="my-1 px-1 w-full">
                        <article className="overflow-hidden rounded-lg shadow-lg">
                            <header className="flex items-center justify-between leading-tight gap-x-2 p-2 md:p-4">
                            <h1 className="text-lg">
                                {projectRequest.team_id}
                            </h1>
                            {
                                projectRequest.status === 'pending' ?
                                <>
                                <button type="button" className="flex justify-center items-center px-4 py-2 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-blue-500 hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 " onClick={()=>handleAcceptProjectRequest(projectRequest.project_request_id!)}>
                                    Accept
                                </button>
                                <button type="button" className="flex justify-center items-center px-4 py-2 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-red-500 hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 " onClick={()=>handleRejectProjectRequest(projectRequest.project_request_id!)}>
                                    Reject
                                </button>
                                </> :
                                <div>
                                    projectRequest.status
                                </div>

                            }
                            </header>

                            <div className="flex items-center justify-between leading-none p-2 md:p-4">
                            <p className="ml-2 text-sm">
                                {projectRequest.message}
                            </p>
                            </div>
                        </article>
                        </div>
                    ))
                    : <div className="my-1 px-1 w-full">
                            <header className="flex items-center justify-between leading-tight gap-x-2 p-2 md:p-4 ">
                            <h1 className="text-lg">
                                No project request
                            </h1>
                            </header>
                        </div>
                    }
                    </div>
                </div>
            }
        </div>
        </Modal>
    )
}

export default ProjectDetailModal