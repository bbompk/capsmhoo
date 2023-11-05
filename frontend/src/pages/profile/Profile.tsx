import Swal from 'sweetalert2'
import { useNavigate } from 'react-router-dom';
import { useState, useEffect } from 'react';

import { useUser } from '../../hooks/useUser';
// import { getUserProfile, updateUserProfile } from '../../service/UserService'; // These services need to be implemented

const MOCK_DATA = {
    professor: {
      id: 'PROF-001',
      name: 'Jane Doe',
      email: 'jane.doe@university.edu',
      password: 'password',
      profile: 'Computer Science Professor specializing in Artificial Intelligence.',
    },
    student: {
      id: 'STUD-001',
      name: 'John Smith',
      email: 'john.smith@student.university.edu',
      password: 'password',
      team_id: 'TEAM-001',
      profile: ''
    }
  };

  const Profile = () => {
    const { userId, role } = useUser();
    
    // Initialize state with mock data based on the role
    const [profileData, setProfileData] = useState(
      role === 'professor' ? MOCK_DATA.professor : MOCK_DATA.student
    );
  
    // Normally, you'd fetch the user profile data using userId and role here
    
    const handleChange = (e) => {
      const { name, value } = e.target;
      setProfileData((prevData) => ({
        ...prevData,
        [name]: value,
      }));
    };
  
    const handleUpdate = async () => {
      // Here you would send the update to your backend
      console.log('Updated profile data:', profileData);
      // Simulate a successful update
      alert('Profile updated successfully!');
    };
  
    return (
      <div className="profile-page-container">
        <div className="flex min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
            <div className="max-w-md w-full space-y-8">
                <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
                    Profile Page
                </h2>
            </div>
        </div>
        <input
          type="text"
          name="name"
          value={profileData.name}
          onChange={handleChange}
        />
        {/* ... other input fields ... */}
        {role === 'professor' && (
          <textarea
            name="profile"
            value={profileData.profile}
            onChange={handleChange}
          />
        )}
        <button onClick={handleUpdate}>Update</button>
      </div>
    );
  };
  
  export default Profile;

