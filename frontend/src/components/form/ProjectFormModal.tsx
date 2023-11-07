import { useEffect, useState } from 'react'
import { Dropdown, Modal, Space, Button } from "antd"
import { DownOutlined } from '@ant-design/icons';
import type { MenuProps } from 'antd';
import Swal from 'sweetalert2'

import { getProjectById, createProject, updateProjectById } from '../../service/ProjectService';

const items: MenuProps['items'] = [
    {
        label: 'item 1',
        key: 'item 1',
    },
    {
        label: 'item 2',
        key: 'item 2',
    },
    {
        label: 'item 3',
        key: 'item 3',
    },
]

export type FormMode = 'create' | 'edit';

export const ProjectFormModal = ({isModalVisible, formMode, setOpenModal, projectId}:{ isModalVisible:boolean, formMode:FormMode ,setOpenModal:React.Dispatch<React.SetStateAction<boolean>>, projectId:string}) => {
    const [projectName, setProjectName] = useState<string>('');
    const [projectDescription, setProjectDescription] = useState<string>('');
    const [projectLabel, setProjectLabel] = useState<string>('Label');

    const professorId = localStorage.getItem("professorId")

    useEffect(() => {
        if(formMode == 'edit' && isModalVisible){
            getProjectById(projectId)
            .then((res) => {
                if(res.data){
                    setProjectName(res.data.name)
                    setProjectDescription(res.data.description)
                    setProjectLabel(res.data.label)
                }
            })
            .catch((err) => {
                console.log(err);
                Swal.fire("Error","Cannot get project", 'error')
                setOpenModal(false);
                resetForm();
            })
        }
    }, [formMode, projectId, isModalVisible])

    const resetForm = () => {
        setProjectName('');
        setProjectDescription('');
        setProjectLabel('Label');
    }
    
    const handleLabelClick: MenuProps['onClick'] = (e) => {
        console.log('click', e.key);
        setProjectLabel(e.key);
    };
    
    const menuProps = {
        items,
        onClick: handleLabelClick,
    };

    const handleOk = async () => {
        console.log(projectName, projectDescription, projectLabel, professorId);
        try{
            if(formMode == 'create'){
                if(!projectName || !projectDescription || !projectLabel || !professorId){
                    return;
                }
                await createProject({name: projectName, description: projectDescription, label: projectLabel, professor_id: professorId, team_id: '', status: 'open'})
            }else if(formMode == 'edit'){
                if(!projectName || !projectDescription || !projectLabel || !professorId){
                    console.log('missing field');
                    return;
                }
                await updateProjectById(projectId, {name: projectName, description: projectDescription, label: projectLabel, professor_id: professorId, team_id: '', status: 'open'})
    
            }
        }catch(err){
            console.log(err);
            Swal.fire("Error","Cannot create project", 'error')
        }finally{
            setOpenModal(false);
            resetForm();
        }
    }

    const handleCancel = () => {
        console.log('cancel');
        setOpenModal(false)
        resetForm();
    }

    return (
    <Modal title={formMode=='create'? 'Create New Project' : 'Edit Your Project'} open={isModalVisible} onOk={handleOk} onCancel={handleCancel}>
        <div className="space-y-12">
            <div className="border-b border-gray-900/10 pb-12">
                <div className="grid grid-cols-1 gap-x-6 gap-y-4 sm:grid-cols-6">
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
                            className="block flex-1 border-0 bg-transparent py-1.5 pl-1 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6"
                            placeholder="project name"
                            value={projectName}
                            onChange={(e) => setProjectName(e.target.value)}
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
                            rows={3}
                            className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                            placeholder="Project Description"
                            value={projectDescription}
                            onChange={(e) => setProjectDescription(e.target.value)}
                        />
                    </div>
                </div>

                <div className="col-span-full">
                    <label htmlFor="label" className="block text-sm font-medium leading-6 text-gray-900">
                        Label
                    </label>
                    <div className="mt-2">
                        <Dropdown menu={menuProps}>
                            <Button>
                                <Space>
                                    <span>{projectLabel}</span>
                                    <DownOutlined />
                                </Space>
                            </Button>
                        </Dropdown>
                    </div>
                </div>
                </div>
            </div>
        </div>
    </Modal>
    )
}