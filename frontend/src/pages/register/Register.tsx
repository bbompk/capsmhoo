import { useState } from 'react';
import Swal from 'sweetalert2'
import { useNavigate } from 'react-router-dom';

import { createStudent } from "../../service/StudentService";
import { createProfessor } from "../../service/ProfessorService";


enum UserType {
    Student = 'student',
    Professor = 'professor',
  }
  
  const Register = () => {
    const [userType, setUserType] = useState<UserType>(UserType.Student);
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [profile, setProfile] = useState('');
  
    const navigate = useNavigate();
  
    const resetForm = () => {
      setName('');
      setEmail('');
      setPassword('');
      setProfile('');
    };
  
    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
      e.preventDefault();
  
      try {
        if (userType === UserType.Student) {
          await createStudent(name, email, password);
        } else {
          await createProfessor(name, email, password, profile);
        }
      } catch (error) {
        console.error(error);
        Swal.fire({
          icon: 'error',
          title: 'Registration Failed',
          text: 'Please try again',
        });
        resetForm();
        return;
      }
  
      Swal.fire({
        icon: 'success',
        title: 'Success',
        text: `Registered successfully as ${userType}`,
      });
  
      navigate('/'); // Navigate to the home page or dashboard
    };
  
    return (
      <>
        {/* Toggle buttons for UserType */}
        <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
            <div className="sm:mx-auto sm:w-full sm:max-w-sm">
                <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
                    Register your account
                </h2>
            </div>
            <div className="flex min-h-full flex-1 flex-row justify-center px-6 py-12 lg:px-8 space-x-8 user-type-toggle">
                <button onClick={() => setUserType(UserType.Student)}>Student</button>
                <button onClick={() => setUserType(UserType.Professor)}>Professor</button>
            </div>
            <div className="flex min-h-full flex-1 flex-col justify-center">
                <form onSubmit={handleSubmit} className="flex min-h-full flex-1 flex-col justify-center">
                    {/* Shared fields between Student and Professor */}
                    <input type="text" value={name} onChange={(e) => setName(e.target.value)} required placeholder="Name"/>
                    <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required placeholder="Email" />
                    <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required placeholder="Password" />
                        {/* Professor-specific field */}
                        {userType === UserType.Professor && (
                        <input type="text" value={profile} onChange={(e) => setProfile(e.target.value)} required placeholder="Profile" />
                    )}
                    <button type="submit"
                        className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                    >
                        Register as {userType}
                    </button>
                    {/* <button type="submit">Register as {userType}</button> */}
                </form>
            </div>
        </div>
      </>
    );
  };
  
  export default Register;
