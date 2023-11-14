import { useNavigate } from "react-router-dom";
import "./Card.css";

interface Props {
  id: string;
  title: string;
  body: string;
  next_path: string;
}

const Card = ({ id, title, body, next_path }: Props) => {
  const navigate = useNavigate();
  const handleClick = (id: string) => {
    const string_path = next_path + "/" + id;
    navigate(string_path);
    return;
  }

  return (
    <div className="card-container">
      <div className="card-content">
        <div className="card-title">
          <h5>{title}</h5>
        </div>
        <div className="card-body">
          <div className="card-text">
            <p>{body}</p>
          </div>
          <div className="card-botton btn btn-primary">
            <button type="button" onClick={() => handleClick(id)}>
              View More
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Card;
