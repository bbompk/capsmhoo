import { useState, useEffect, Fragment } from "react";
import { Menu, Transition } from "@headlessui/react";
import { BellIcon } from "@heroicons/react/24/outline";
import { NotificationInterface } from "../../interfaces/NotiInterface";
import { useUser,useStudent,useProfessor } from "../../hooks/useUser"
import { getAllNotificationsByRoleId,readAllNotificationsByRoleId } from '../../service/NotiService'

const Notification = () => {
  const [isShowing, setIsShowing] = useState(false)
  const [notiList, setNotiList] = useState<NotificationInterface[]>()
  const [dummy, setDummy] = useState("deadDummy")
  const {userId,role} = useUser()
  // const studentId = useStudent()
  // const professorId = useProfessor()
  // const role = "Professor"
  const studentId = "useStudent()"
  const professorId = "PRF-f9b8730a-dddd-47ce-a4cd-4985f78c5223"
  useEffect(() => {
    if (role === "Student"){
      const fetch = async () => {
        const res = await getAllNotificationsByRoleId(studentId!)
        if(res && res.data){setNotiList(res.data)}
      }
      fetch();
    }
    if (role === "Professor"){
      const fetch = async () => {
        console.log("fetching")
        const res = await getAllNotificationsByRoleId(professorId!)
        if(res && res.data){console.log(res);setNotiList(res.data)}
        else {console.log("no res")}
      }
      fetch();
    }
  },[role,studentId,professorId])
  useEffect(()=>{
    setDummy('dummy')
  },[isShowing])
  useEffect(()=>{
    const read = async () => {
      if(role==="Student")readAllNotificationsByRoleId(studentId)
      if(role==="Professor")readAllNotificationsByRoleId(professorId)
    }
    read()
  },[dummy])
  return (
      // <div>
      <Menu as="div">
      <Menu.Button
        type="button" onClick={() => setIsShowing((isShowing) => !isShowing)} 
        className="rounded-full bg-gray-800 p-1 text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
      >
        <span className="sr-only">View notifications</span>
        <BellIcon className="h-6 w-6" aria-hidden="true"/>
      </Menu.Button>
      <Menu.Items className="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
        <Transition
          as={Fragment}
          show={isShowing}
          enter="transition ease-out duration-100"
          enterFrom="transform opacity-0 scale-95"
          enterTo="transform opacity-100 scale-100"
          leave="transition ease-in duration-75"
          leaveFrom="transform opacity-100 scale-100"
          leaveTo="transform opacity-0 scale-95"
        >
          <h1>hello</h1>
          </Transition>
        </Menu.Items>
      </Menu>
      // </div>
    );
  };
  
export default Notification;
  