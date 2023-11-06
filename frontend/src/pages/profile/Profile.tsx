import { useState, useEffect } from "react";
import Swal from "sweetalert2";
import { useNavigate } from "react-router-dom";

import { useUser } from "../../hooks/useUser";
import { getUserById, updateUserById } from "../../service/UserService";
import {
  getStudentByUserId,
  updateStudentById,
} from "../../service/StudentService";
import {
  getProfessorByUserId,
  updateProfessorById,
} from "../../service/ProfessorService";

const Profile = () => {
  const { userId, role } = useUser();
  const navigate = useNavigate();
  const [profileData, setProfileData] = useState({
    name: "",
    email: "",
    password: "",
    profile: "",
  });
  const [initialData, setInitialData] = useState({
    roleId: "",
    email: "",
    name: "",
    profile: "",
  });

  useEffect(() => {
    const fetchUserData = async () => {
      if (!userId) {
        Swal.fire("Please log in to view your profile.");
        navigate("/login");
        return;
      }
      const user = await getUserById(userId);

      try {
        if (!user.data) {
          throw new Error("Failed to fetch user data");
        }
        setProfileData((prevData) => ({
          ...prevData,
          email: user.data?.email ?? "",
          // password: user.data?.password ?? '',
        }));
        if (role === "Student") {
          const student = await getStudentByUserId(userId);
          if (!student.data) {
            throw new Error("Failed to fetch student data");
          }
          setProfileData((prevData) => ({
            ...prevData,
            name: student?.data?.name ?? "",
          }));
          setInitialData({
            roleId: student?.data?.id ?? "",
            email: user.data.email ?? "",
            name: student?.data?.name ?? "",
            profile: "",
          });
        } else if (role === "Professor") {
          const professor = await getProfessorByUserId(userId);
          if (!professor.data) {
            throw new Error("Failed to fetch professor data");
          }
          setProfileData((prevData) => ({
            ...prevData,
            name: professor?.data?.name ?? "",
            profile: professor?.data?.profile ?? "",
          }));
          setInitialData({
            roleId: professor?.data?.id ?? "",
            email: user.data.email ?? "",
            name: professor?.data?.name ?? "",
            profile: professor?.data?.profile ?? "",
          });
        }
      } catch (error) {
        Swal.fire({
          icon: "error",
          title: "Failed to load profile",
          text: "Unable to fetch user data, please try again later.",
        });
      }
    };
    fetchUserData();
  }, [userId, navigate, role]);

  const handleInputChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target as HTMLInputElement | HTMLTextAreaElement;
    setProfileData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      if (role === "Student" && profileData.name !== initialData.name) {
        await updateStudentById(initialData.roleId, { name: profileData.name });
        setInitialData((prevData) => ({
          ...prevData,
          name: profileData.name,
        }));
      }

      if (
        role === "Professor" &&
        (profileData.profile !== initialData.profile ||
          profileData.name !== initialData.name)
      ) {
        await updateProfessorById(initialData.roleId, {
          name: profileData.name,
          profile: profileData.profile,
        });
        setInitialData((prevData) => ({
          ...prevData,
          name: profileData.name,
          profile: profileData.profile,
        }));
      }

      if (userId && profileData.email !== initialData.email) {
        await updateUserById(userId, { email: profileData.email });
        setInitialData((prevData) => ({
          ...prevData,
          email: profileData.email,
        }));
      }

      if (userId && profileData.password) {
        await updateUserById(userId, { password: profileData.password });
        // force logout
        localStorage.removeItem("accessToken");
        localStorage.removeItem("token_expires");
        localStorage.removeItem("userId");
        localStorage.removeItem("role");
        localStorage.removeItem("studentId");
        localStorage.removeItem("professorId");
        console.log("logout");
        Swal.fire({
          icon: "success",
          title: "Your password has been updated",
          text: "Please re-login with new password",
        });
        navigate("/login");
        return;
      }
      Swal.fire({
        icon: "success",
        title: "Your profile has been updated",
      });
      // Swal.fire('Success', 'Your profile has been updated.', 'success');
    } catch (error) {
      Swal.fire({
        icon: "error",
        title: "There was a problem updating your profile.",
        text: "Please try again later.",
      });
      // Swal.fire('Error', 'There was a problem updating your profile.', 'error');
    }
  };

  return (
    <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8 profile-page-container">
      <div className="sm:mx-auto sm:w-full sm:max-w-md">
        <h2 className="mt-10 text-center text-3xl font-extrabold text-gray-900">
          Edit Profile
        </h2>
        <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
          {" "}
          {/* Updated this line */}
          <div>
            <label
              htmlFor="name"
              className="block text-sm font-medium text-gray-700"
            >
              Name
            </label>
            <input
              type="text"
              name="name"
              id="name"
              value={profileData.name}
              onChange={handleInputChange}
              required
              placeholder="Your name"
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
            />
          </div>
          <div>
            <label
              htmlFor="email"
              className="block text-sm font-medium text-gray-700"
            >
              Email
            </label>
            <input
              type="email"
              name="email"
              id="email"
              value={profileData.email}
              onChange={handleInputChange}
              required
              placeholder="Your email"
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
            />
          </div>
          <div>
            <label
              htmlFor="password"
              className="block text-sm font-medium text-gray-700"
            >
              Password (leave blank to keep current password)
            </label>
            <input
              type="password"
              name="password"
              id="password"
              value={profileData.password}
              onChange={handleInputChange}
              placeholder="New password"
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
            />
          </div>
          {role === "Professor" && (
            <div>
              <label
                htmlFor="profile"
                className="block text-sm font-medium text-gray-700"
              >
                Profile
              </label>
              <textarea
                name="profile"
                id="profile"
                value={profileData.profile}
                onChange={handleInputChange}
                placeholder="Your profile description"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              />
            </div>
          )}
          <button
            type="submit"
            className="group relative flex w-full justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
          >
            Update Profile
          </button>
        </form>
      </div>
    </div>
  );
};

export default Profile;
