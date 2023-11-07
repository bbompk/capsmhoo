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
  const studentId = useStudent()
  const professorId = useProfessor()
  // const role = "Professor"
  // const studentId = "useStudent()"
  // const professorId = "PRF-f9b8730a-dddd-47ce-a4cd-4985f78c5223"
  useEffect(() => {
    const fetch = async (Id: string) => {
      const res = await getAllNotificationsByRoleId(Id!)
      if(res && res.data){setNotiList(res.data)}
    }
    if (role === "Student"){
      fetch(studentId);
    }
    if (role === "Professor"){
      fetch(professorId);
    }
  },[role,studentId,professorId])
  useEffect(()=>{
    setDummy('dummy')
  },[isShowing])
  const read = async () => {
    // if(role==="Student")readAllNotificationsByRoleId(studentId)
    if(role==="Professor")readAllNotificationsByRoleId(professorId)
  }
  useEffect(()=>{
    // 
  },[dummy])
  console.log(notiList)
  return (
      <div>
      <Menu as="div">
      <Menu.Button
        type="button" onClick={() => {setIsShowing((isShowing) => !isShowing);read()}} 
        className="rounded-full bg-gray-800 p-1 text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
      >
        <span className="sr-only">View notifications</span>
        <BellIcon className="h-6 w-6" aria-hidden="true"/>
        <h1 style={{color: "red", background: "white",border: "2px solid red", borderRadius: "50%",fontWeight: "bold"}}>{notiList?.length}</h1>
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
          <div>
          {
            notiList && dummy ?
              notiList.map((notification,index) => (
                <Menu.Item key={index}>
                  <div
                    style={index === 0 ? {
                      paddingTop: '5px',
                      paddingBottom: '5px'
                    } : {
                      borderTop: '2px solid black',
                      paddingTop: '5px',
                      paddingBottom: '5px'
                    }}
                    onMouseEnter={(e) => e.currentTarget.style.backgroundColor = 'rgb(243 244 246)'}
                    onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'transparent'}
                  >
                    <p>{notification.body}</p>
                  </div>
                </Menu.Item>
              ))
              : <Menu.Item>
                <p>there's no notification</p>
              </Menu.Item>
          }
          </div>
          </Transition>
        </Menu.Items>
      </Menu>
      </div>
    );
  };
  
export default Notification;
  