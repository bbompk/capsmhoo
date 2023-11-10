import { useState, useEffect } from "react";
import React, { CSSProperties } from 'react';
import { useUser,useStudent,useProfessor } from "../../hooks/useUser"
import { useNavigate } from "react-router-dom";
import Swal from "sweetalert2";
import { getAllStudentByTeamId, getStudentById, getStudentByUserId } from "../../service/StudentService";
import { getTeamById, updateTeamById } from "../../service/TeamService";
import { ProfessorInterface, StudentInterface } from "../../interfaces/UserInterface";
import { TeamInterface } from "../../interfaces/TeamInterface";
import { ProjectInterface } from "../../interfaces/ProjectInterface";
import { getProjectByTeamId } from "../../service/ProjectService";
import { getProfessorById } from "../../service/ProfessorService";

const MyTeam = () => {
  const { userId, role } = useUser();
  const [student, SetStudent] = useState<StudentInterface>();
  const [studentList, SetStudentList] = useState<StudentInterface[]>();
  const [team, SetTeam] = useState<TeamInterface>();
  const [professor, SetProfessor] = useState<ProfessorInterface>();
  const [project, SetProject] = useState<ProjectInterface>();
  const [teamName, setTeamName] = useState('');
  const [teamProfile, setTeamProfile] = useState('');
  const navigate = useNavigate();
  const styles: Record<string, CSSProperties> = {
 
    container: {
      display: 'flex',
      flexDirection: 'column',
      // alignItems: 'center',
      // justifyContent: 'center',
      marginTop: "5rem",
      minHeight: '100vh', // This will take up the full screen height
      maxWidth: "50vw",
      marginInline: "auto"
    },
    formGroup: {
      margin: '1rem 0',
      display: 'flex',
      flexDirection: 'row',
      // alignItems: 'center',
      width: '100%', // Or you could give it a max-width and set width to 'auto'
    },
    input: {
      padding: '0.5rem',
      margin: '0.5rem 0',
      border: '1px solid #ccc',
      borderRadius: '4px',
      width: '80%', // Adjust the width as needed
    },
    textArea: {
      padding: '0.5rem',
      margin: '0.5rem 0',
      border: '1px solid #ccc',
      borderRadius: '4px',
      width: '80%', // Adjust the width as needed
      minHeight: '100px', // You can adjust this as well
    },
    button: {
      padding: '0.5rem 1rem',
      border: 'none',
      borderRadius: '4px',
      cursor: 'pointer',
      backgroundColor: '#007bff',
      color: 'white',
      fontWeight: 'bold',
      fontSize: '1rem',
      marginTop: '1rem',
    },
  };
  
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
        setTeamName(resTeam.data.name)
        setTeamProfile(resTeam.data.profile)
        console.log(resTeam)
        const resStudentList = await getAllStudentByTeamId(resStudent.data!.team_id!)
        if(resStudentList.data){
          SetStudentList(resStudentList.data)
        }
        const resProject = await getProjectByTeamId(resTeam.data.id)
          // Should return []
          if(resProject.data){
            console.log(resProject.data)
            SetProject(resProject.data)
            const resProfessor = await getProfessorById(resProject.data.professor_id)
            if(resProfessor.data){
              SetProfessor(resProfessor.data)
            }
          }
        }
      }
      fetchUserData()
  }, [userId,navigate,role]);
  const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setTeamName(e.target.value);
  };

  const handleProfileChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setTeamProfile(e.target.value);
  };
  const updateTeamData = async () => {
    // Logic to update team data, possibly involving an API call
    const newTeam : TeamInterface = {
      id: team?.id!,
      name: teamName,
      profile: teamProfile
    }
    await updateTeamById(newTeam.id, newTeam)
  };
  return (
    <div style={styles.container}>
      <h1 style={{fontSize: "1.3rem", fontWeight: "normal"}}>My Team</h1>
    {team && (
      <>
        <div style={styles.formGroup}>
          <label htmlFor="teamName" style={{marginRight:"0.5rem",marginTop: "1rem"}}>Team Name:</label>
          <input
            id="teamName"
            style={styles.input}
            value={teamName}
            onChange={handleNameChange}
          />
        </div>
        <div style={styles.formGroup}>
          <label htmlFor="teamProfile" style={{marginRight:"0.5rem",marginTop: "1rem"}}>Team Profile:</label>
          <textarea
            id="teamProfile"
            style={styles.textArea}
            value={teamProfile}
            onChange={handleProfileChange}
          />
        </div>
        {studentList && (
          <div>
            <h1>Student List</h1>
            {studentList.map((student,studentnum)=>
              <h1>{studentnum+1+" "+student.name}</h1>
            )}
          </div>
        )}
        <button style={styles.button} onClick={updateTeamData}>
          Save Changes
        </button>
        </>
      )}
      {project && (
        <div style={{marginTop:"2rem"}}>
          <h1 style={{fontSize: "1.3rem", fontWeight: "normal"}}>My Project</h1>
          <h1>{"Project Name: "+project.name}</h1>
          <h1>{"Project Label: "+project.label}</h1>
          <h1>{"Project Detail: "+project.description}</h1>
          <h1>{"Professor Name: "+professor?.name}</h1>
        </div>
      )}
    </div>
  );
  };
  
export default MyTeam;
  