import { useState } from 'react';
import { login } from '../../service/AuthService'
import Swal from 'sweetalert2'
import { useNavigate } from 'react-router-dom';

const Profile = () => {
    return (
        <div className="flex min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
            <div className="max-w-md w-full space-y-8">
                <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
                    Profile Page
                </h2>
                {/* Profile content will go here */}
            </div>
        </div>
    );
};
export default Profile;
