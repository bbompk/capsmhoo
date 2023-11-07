import { useState, useEffect } from "react";
import { useUser,useStudent,useProfessor } from "../../hooks/useUser"
import { useNavigate } from "react-router-dom";
import Swal from "sweetalert2";
import { getStudentById, getStudentByUserId } from "../../service/StudentService";
import { getTeamById } from "../../service/TeamService";
import { StudentInterface } from "../../interfaces/UserInterface";
import { TeamInterface } from "../../interfaces/TeamInterface";
import { ProjectInterface } from "../../interfaces/ProjectInterface";
import { getProjectByTeamId } from "../../service/ProjectService";

const MyTeam = () => {
  const { userId, role } = useUser();
  const [student, SetStudent] = useState<StudentInterface>();
  const [team, SetTeam] = useState<TeamInterface>();
  const [project, SetProject] = useState<ProjectInterface>();
  const navigate = useNavigate();
    useEffect(() => {
      const fetchUserData = async () => {
        // Should be create as a hook
        if (!userId) {
          Swal.fire("Please log in to view your profile.");
          navigate("/login");
          return;
        }
        // Get Team and Project data
        const resStudent = await getStudentByUserId(userId)
        if(!resStudent.data){
          throw new Error("Failed to fetch student data")
        }
        SetStudent(resStudent.data)
        if(resStudent.data?.team_id!= ""){
          const resTeam = await getTeamById(resStudent.data!.team_id!)
          if(!resTeam.data){
            throw new Error("Failed to fetch team data")
          }
          SetTeam(resTeam.data)
          console.log(resTeam)
          const resProject = await getProjectByTeamId(resTeam.data.id)
            // Should return []
            if(resProject.data){
              console.log(resProject.data)
              SetProject(resProject.data)
            }
          }
        }
        fetchUserData()
    }, [userId,navigate,role]);
    return (
      <div>
        <h1>My Team</h1>
        <h1>{team?.name}</h1>
        <h1>{team?.profile}</h1>
        <h1>{project?.name}</h1>
      </div>
    );
  };
  
export default MyTeam;
  