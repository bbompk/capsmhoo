import { useState } from "react";
import { useNavigate } from "react-router-dom";
import Swal from "sweetalert2";

import { useUser } from "../../hooks/useUser";
import { createTeam } from "../../service/TeamService";
import { TeamCreateInterface } from "../../interfaces/TeamInterface";
import { getStudentByUserId } from "../../service/StudentService";

export default function CreateTeamForm() {
  const [name, setName] = useState("");
  const [profile, setProfile] = useState("");

  const { userId, role } = useUser();
  const navigate = useNavigate();

  const resetForm = () => {
    setName("");
    setProfile("");
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!userId) {
      Swal.fire({
        icon: "error",
        title: "Creating Team Failed",
        text: "Please log in to create a team.",
      });
      navigate("/login");
      return;
    } else if (role !== "Student") {
      Swal.fire({
        icon: "error",
        title: "Creating Team Failed",
        text: "Only student role can create team.",
      });
      navigate("/");
      return;
    }

    const student = await getStudentByUserId(userId);
    if (!student.data) {
      throw new Error("Failed to fetch student data");
    }

    const teamCreate: TeamCreateInterface = {
      id: "",
      name: name,
      profile: profile,
      creator_id: student.data.id,
    };

    try {
      const createdTeam = await createTeam(teamCreate);
      const team_id = createdTeam.data?.id;
      navigate(`/team-detail/${team_id}`);
    } catch (error) {
      console.error(error);
      Swal.fire({
        icon: "error",
        title: "Creating Team Failed",
        text: "Please try again",
      });
      resetForm();
      return;
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div className="space-y-12">
        <div className="border-b border-gray-900/10 pb-12">
          <div className="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
            <div className="sm:col-span-4">
              <label
                htmlFor="name"
                className="block text-sm font-medium leading-6 text-gray-900"
              >
                Team Name
              </label>
              <div className="mt-2">
                <div className="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 focus-within:ring-2 focus-within:ring-inset focus-within:ring-indigo-600 sm:max-w-md">
                  <input
                    type="text"
                    name="name"
                    id="name"
                    required
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    className="block flex-1 border-0 bg-transparent py-1.5 pl-1 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6"
                    placeholder="team name"
                  />
                </div>
              </div>
            </div>

            <div className="col-span-full">
              <label
                htmlFor="profile"
                className="block text-sm font-medium leading-6 text-gray-900"
              >
                Profile
              </label>
              <div className="mt-2">
                <textarea
                  id="profile"
                  name="profile"
                  value={profile}
                  onChange={(e) => setProfile(e.target.value)}
                  rows={3}
                  className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  placeholder="Team Description"
                ></textarea>
              </div>
            </div>

            <div className="mt-6 flex items-center justify-end gap-x-6">
              <button
                type="submit"
                className="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
              >
                Submit
              </button>
            </div>
          </div>
        </div>
      </div>
    </form>
  );
}
