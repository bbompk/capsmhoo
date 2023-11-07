import { useState } from "react";
import Swal from "sweetalert2";
import { useNavigate } from "react-router-dom";

import { createStudent } from "../../service/StudentService";
import { createProfessor } from "../../service/ProfessorService";

enum UserType {
  Student = "student",
  Professor = "professor",
}

const Register = () => {
  const [userType, setUserType] = useState<UserType>(UserType.Student);
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [profile, setProfile] = useState("");

  const navigate = useNavigate();

  const resetForm = () => {
    setName("");
    setEmail("");
    setPassword("");
    setProfile("");
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
        icon: "error",
        title: "Registration Failed",
        text: "Please try again",
      });
      resetForm();
      return;
    }

    Swal.fire({
      icon: "success",
      title: "Success",
      text: `Registered successfully as ${userType}`,
    });

    navigate("/"); // Navigate to the home page or dashboard
  };

  return (
    <>
      <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
        <div className="sm:mx-auto sm:w-full sm:max-w-md">
          <h2 className="mt-10 text-center text-3xl font-extrabold text-gray-900">
            Register your account
          </h2>
          <div className="mt-8 flex justify-center">
            <button
              type="button"
              onClick={() => setUserType(UserType.Student)}
              className={`px-4 py-2 rounded-md ${
                userType === UserType.Student
                  ? "bg-indigo-600 text-white"
                  : "bg-white text-indigo-600"
              }`}
            >
              Student
            </button>
            <button
              type="button"
              onClick={() => setUserType(UserType.Professor)}
              className={`px-4 py-2 rounded-md ml-4 ${
                userType === UserType.Professor
                  ? "bg-indigo-600 text-white"
                  : "bg-white text-indigo-600"
              }`}
            >
              Professor
            </button>
          </div>
          <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
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
                required
                placeholder="Your name"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                value={name}
                onChange={(e) => setName(e.target.value)}
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
                required
                placeholder="Your email"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
            </div>
            <div>
              <label
                htmlFor="password"
                className="block text-sm font-medium text-gray-700"
              >
                Password
              </label>
              <input
                type="password"
                name="password"
                id="password"
                required
                placeholder="Your password"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
            {userType === UserType.Professor && (
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
                  required
                  placeholder="Your profile description"
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                  value={profile}
                  onChange={(e) => setProfile(e.target.value)}
                />
              </div>
            )}
            <button
              type="submit"
              className="group relative flex w-full justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
            >
              Register as {userType}
            </button>
          </form>
        </div>
      </div>
    </>
  );
};

export default Register;
