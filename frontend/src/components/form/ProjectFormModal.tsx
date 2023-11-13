import { useEffect, useState, useCallback } from 'react'
import { Dropdown, Modal, Space, Button } from "antd"
import { DownOutlined } from '@ant-design/icons';
import type { MenuProps } from 'antd';
import Swal from 'sweetalert2'

import { getProjectById, createProject, updateProjectById, deleteProjectById } from '../../service/ProjectService';

const items: MenuProps['items'] = [
    {
        label: 'data science',
        key: 'data science',
    },
    {
        label: 'software development',
        key: 'software development',
    },
    {
        label: 'research',
        key: 'research',
    },
]

export type FormMode = 'create' | 'edit';

export const ProjectFormModal = ({isModalVisible, formMode, setOpenModal, projectId}:{ isModalVisible:boolean, formMode:FormMode ,setOpenModal:React.Dispatch<React.SetStateAction<boolean>>, projectId:string}) => {
    const [projectName, setProjectName] = useState<string>('');
    const [projectDescription, setProjectDescription] = useState<string>('');
    const [projectLabel, setProjectLabel] = useState<string>('Label');

    const professorId = sessionStorage.getItem("professorId")

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
        try{
            if(formMode == 'create'){
                if(!projectName || !projectDescription || !projectLabel || !professorId){
                    return;
                }
                await createProject({name: projectName, description: projectDescription, label: projectLabel, professor_id: professorId, team_id: '', status: 'open'})
                Swal.fire("Success","Project created", 'success')
            }else if(formMode == 'edit'){
                if(!projectName || !projectDescription || !projectLabel || !professorId){
                    console.log('missing field');
                    return;
                }
                await updateProjectById(projectId, {name: projectName, description: projectDescription, label: projectLabel, professor_id: professorId, team_id: '', status: 'open'})
                Swal.fire("Success","Project edited", 'success')
            }
        }catch(err){
            console.log(err);
            Swal.fire("Error","Cannot create/edit project", 'error')
        }finally{
            setOpenModal(false);
            resetForm();
        }
    }

    const handleDelete = async () => {
        try{
            const projectRes = await deleteProjectById(projectId)
            if(!projectRes.data)return;
            Swal.fire("Success","Project deleted", 'success')
        }catch(err){
            console.log(err);
            Swal.fire("Error","Cannot delete project", 'error')
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
    <Modal 
        title={
            formMode=='create'
            ? <h1>Create New Project</h1> 
            : <div className='flex justify-between mr-10'><h1>Edit Your Project</h1>
                <button className='flex justify-center items-center px-4 py-1 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-red-500 hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500'
                onClick={handleDelete}
                >
                    Delete
                </button>
            </div>
        } 
        open={isModalVisible} 
        onOk={handleOk} 
        onCancel={handleCancel}
        okButtonProps={{className: 'px-4 py-1 border-transparent text-base rounded-md shadow-sm text-white bg-blue-500 hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 cursor-pointer'}}
        >
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