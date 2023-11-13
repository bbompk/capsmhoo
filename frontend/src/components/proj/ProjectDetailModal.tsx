import { useState, useEffect, useCallback } from "react"
import Swal from 'sweetalert2'
import { Modal, Tag } from "antd";

import { getProjectById } from "../../service/ProjectService";
import { ProjectRequestInterface, ProjectRequestWithTeamInfoInterface } from "../../interfaces/ProjectInterface";
import { createProjectRequest, getAllProjectRequestByProjectId, acceptProjectRequestById, rejectProjectRequestById } from "../../service/ProjectRequestService";
import { getStudentByUserId } from "../../service/StudentService";
import { getTeamById } from "../../service/TeamService";

const ProjectDetailModal = ({projectId, isModalVisible, setOpenModal}: {projectId:string, isModalVisible:boolean, setOpenModal:React.Dispatch<React.SetStateAction<boolean>>}) => {
    const [projectRequestData, setProjectRequestData] = useState<ProjectRequestInterface[]|null>([]);
    const [projectRequestWithTeamInfoData, setProjectRequestWithTeamInfoData] = useState<ProjectRequestWithTeamInfoInterface[]|null>([]);
    const [projectName, setProjectName] = useState<string>('');
    const [projectDescription, setProjectDescription] = useState<string>('');
    const [projectLabel, setProjectLabel] = useState<string>('Label');
    const [projectStatus, setProjectStatus] = useState<string>('open');
    const [projectTeamId, setProjectTeamId] = useState<string>('');
    const [teamId, setTeamId ] = useState<string>('');

    const role = sessionStorage.getItem("role")
    const userId = sessionStorage.getItem("userId")

    const fetchProjectData = useCallback(async () => {
        resetProjectModal();
        // get project info
        const projectRes = await getProjectById(projectId)
        console.log(projectRes.data)
        if(!projectRes.data) return;
        setProjectName(projectRes.data.name);
        setProjectDescription(projectRes.data.description);
        setProjectLabel(projectRes.data.label);
        setProjectStatus(projectRes.data.status);
        setProjectTeamId(projectRes.data.team_id || '');

        // get project requests for the project
        const projectRequestRes = await getAllProjectRequestByProjectId(projectId)
        if(projectRequestRes.code !== '200') return;
        setProjectRequestData(projectRequestRes.data||[]);

        // get team info for each project request
        if(!projectRequestRes.data)return;
        // if(projectStatus==='closed') return;
        let projectRequestWithTeamInfo:ProjectRequestWithTeamInfoInterface[] = [];
        await projectRequestRes.data.reduce(async (memo, projectRequest) => {
            await memo;
            if(!projectRequest.team_id) return;
            const teamRes = await getTeamById(projectRequest.team_id);
            if(!teamRes.data) return;
            projectRequestWithTeamInfo.push({...projectRequest, team_name: teamRes.data.name, team_profile: teamRes.data.profile});
        }, Promise.resolve());
        setProjectRequestWithTeamInfoData(projectRequestWithTeamInfo)
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

    const resetProjectModal = () => {
        setProjectName('');
        setProjectDescription('');
        setProjectLabel('');
        setProjectStatus('');
        setProjectTeamId('');
        setProjectRequestData([]);
        setProjectRequestWithTeamInfoData([]);
    }

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
            const projectRequestRes = await createProjectRequest({project_id: projectId, team_id: teamId, message: 'request to join the project', status: 'pending'})

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

            if(projectRequestRes.code!=='200'){
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
    const pendingProjectRequests = projectRequestWithTeamInfoData?.filter((projectRequest) => projectRequest.status === 'pending')
    const teamOfProject = projectRequestWithTeamInfoData?.find((projectRequest) => projectRequest.status === 'accepted') || {team_name: '', team_id: '', team_profile: '', status: ''};

    return (
        <Modal 
        title={
            <div className="flex justify-between items-center">
                <h1 className="font-bold text-2xl">Project Details</h1>
                <Tag className="text-grey-darker text-sm ml-2 mr-8" bordered={false} color={projectStatus==='open'?'success':'error'}>
                    {projectStatus}
                </Tag>
            </div>
        } 
        open={isModalVisible} footer={null} onCancel={handleClose}
        >
            <div className="pb-4">
                <div className="grid grid-cols-1 gap-x-6 gap-y-4 sm:grid-cols-6">
                <div className="sm:col-span-4">
                    <label htmlFor="name" className="block text-base font-bold leading-6 text-gray-900">
                        Project Name
                    </label>
                    <div className="text-base mt-2">
                        {projectName}
                    </div>
                </div>

                <div className="col-span-full">
                    <label htmlFor="description" className="block text-base font-bold leading-6 text-gray-900">
                        Description
                    </label>
                    <div className="text-base mt-2">
                        {projectDescription}
                    </div>
                </div>

                <div className="col-span-full">
                    <label htmlFor="label" className="block text-base font-bold leading-6 text-gray-900">
                        Label
                    </label>
                    <div className="text-base mt-2">
                        {projectLabel}
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
                <div className="border-b border-gray-900/10">
                    <div key={teamOfProject.team_id} className="my-1 px-1 w-full">
                        <h1 className="text-lg font-bold">Team</h1>
                        <article className="overflow-hidden rounded-lg shadow-lg md:py-3">
                            <header className="flex items-center justify-between leading-tight gap-x-2 p-2 md:px-4">
                                <div className="flex flex-col ">
                                    <h1 className="text-base font-bold">
                                        {teamOfProject.team_name}
                                    </h1>
                                    <p className="text-sm font-light">
                                        {teamOfProject.team_id}
                                    </p>
                                </div>
                            </header>

                            <div className="flex flex-col items-start justify-between leading-none px-2 md:px-4">
                            <p className="text-sm italic">
                                {teamOfProject.team_profile}
                            </p>
                            </div>
                        </article>
                        </div>
                </div>
            }
            {
                // Professor's project request list
                role === 'Professor' && projectStatus === 'open' &&
                <div className="border-b border-gray-900/10 py-4">
                    <h1 className="text-lg font-bold">Project Requests</h1>
                    <div className="flex flex-wrap -mx-1 lg:-mx-4">
                    {pendingProjectRequests && (pendingProjectRequests.length)>0 
                    ? pendingProjectRequests.map((projectRequest) => (
                        <div key={projectRequest.project_request_id} className="my-1 px-1 w-full">
                        <article className="overflow-hidden rounded-lg shadow-lg md:py-3">
                            <header className="flex items-center justify-between leading-tight gap-x-2 p-2 md:px-4">
                                <div className="flex flex-col ">
                                    <h1 className="text-base font-bold">
                                        {projectRequest.team_name}
                                    </h1>
                                    <p className="text-sm font-light">
                                        {projectRequest.team_id}
                                    </p>
                                </div>
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

                            <div className="flex flex-col items-start justify-between leading-none px-2 md:px-4">
                            <p className="text-sm italic">
                                {projectRequest.team_profile}
                            </p>
                            <p className="ml-2 text-sm pt-2">
                                message: {projectRequest.message}
                            </p>
                            </div>
                        </article>
                        </div>
                    ))
                    : <div className="my-1 px-1 w-full">
                            <header className="flex items-center justify-between leading-tight gap-x-2 pl-2 md:pl-4">
                            <h1 className="text-base">
                                No project requests
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