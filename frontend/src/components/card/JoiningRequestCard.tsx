import { useEffect, useState } from "react";
import Swal from "sweetalert2";
import { getStudentById } from "../../service/StudentService";
import 'bootstrap/dist/css/bootstrap.css';

interface Props {
  id: string;
  student_id: string;
  handleAccept: (id: string) => void;
  handleReject: (id: string) => void;
}

const JoiningRequestCard = ({
  id,
  student_id,
  handleAccept,
  handleReject,
}: Props) => {
  const [studentName, setStudentName] = useState("");

  useEffect(() => {
    const fetchData = async () => {
      const student = await getStudentById(student_id);
      if (!student.data) {
        Swal.fire("Failed to load the student data");
        return;
      }
      setStudentName(student.data.name);
    };
    fetchData();
  });

  return (
    <div className="list-group-item flex-column align-items-start">
      <div className="d-flex w-100 justify-content-between">
        <h5 className="mb-1">{studentName}</h5>
        <small>{student_id}</small>
      </div>
      <div className="btn-group" role="group">
        <div className="btn btn-success">
          <button type="button" onClick={() => handleAccept(id)}>
            Accept
          </button>
        </div>
        <div className="btn btn-danger">
          <button type="button" onClick={() => handleReject(id)}>
            Reject
          </button>
        </div>
      </div>
    </div>
  );
};

export default JoiningRequestCard;
