import { useState, useEffect } from "react"
import Swal from 'sweetalert2'
import { Modal } from "antd";

import { getProjectById } from "../../service/ProjectService";

const ProjectDetailModal = ({projectId, isModalVisible, setOpenModal}: {projectId:string, isModalVisible:boolean, setOpenModal:React.Dispatch<React.SetStateAction<boolean>>}) => {
    const [projectName, setProjectName] = useState<string>('');
    const [projectDescription, setProjectDescription] = useState<string>('');
    const [projectLabel, setProjectLabel] = useState<string>('Label');
    const [projectStatus, setProjectStatus] = useState<string>('open');

    useEffect(() => {
        const fetchProjectById = async (projectId: string) => {
            const res = await getProjectById(projectId);
            if(!res.data) return;
            setProjectName(res.data.name);
            setProjectDescription(res.data.description);
            setProjectLabel(res.data.label);
            setProjectStatus(res.data.status);

        }
        try {
            fetchProjectById(projectId);
        } catch (err) {
            Swal.fire("Error", "Cannot get project", 'error')
        }
    }, [projectId])

    const handleClose = () => {
        setOpenModal(false);
    }

    return (
        <Modal title={projectName} open={isModalVisible} footer={null} onCancel={handleClose}>
            <h1>{projectName}</h1>
            <h3>{projectDescription}</h3>
            <h3>{projectLabel}</h3>
            <h3>{projectStatus}</h3>
        </Modal>
    )
}

export default ProjectDetailModal