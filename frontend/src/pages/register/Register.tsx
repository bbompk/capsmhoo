import { useState } from 'react';
import Swal from 'sweetalert2'
import { useNavigate } from 'react-router-dom';

import { UserInterface, StudentInterface, ProfessorInterface } from "../../interfaces/UserInterface";
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
        // Create the user first
        const userResponse = await createUser({ email, password, role: userType });
        if (!userResponse.ok) throw new Error('Failed to create user.');
  
        // Extract the user_id from the response
        const user = await userResponse.json();
        const userId = user.id; // Make sure this matches the actual key in your JSON response
  
        // Depending on the user type, create a student or professor profile
        if (userType === UserType.Student) {
          await createStudent({ name, user_id: userId });
        } else {
          await createProfessor({ name, profile, user_id: userId });
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
            <div className="user-type-toggle">
            <button onClick={() => setUserType(UserType.Student)}>Student</button>
            <button onClick={() => setUserType(UserType.Professor)}>Professor</button>
            </div>
            
            {/* Registration Form */}
            <form onSubmit={handleSubmit}>
            {/* Shared fields between Student and Professor */}
            <input type="text" value={name} onChange={(e) => setName(e.target.value)} required placeholder="Name" />
            <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required placeholder="Email" />
            <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required placeholder="Password" />
            
            {/* Professor-specific field */}
            {userType === UserType.Professor && (
                <input type="text" value={profile} onChange={(e) => setProfile(e.target.value)} required placeholder="Profile" />
            )}
            
            <button type="submit">Register as {userType}</button>
            </form>
        </div>
      </>
    );
  };
  
  export default Register;
