import "./Card.css";

interface Props {
  id: string;
  title: string;
  body: string;
  next_path: string;
}

const Card = ({ id, title, body, next_path }: Props) => {
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
          <div className="card-botton">
            <a
              href={"http://localhost:5173/" + next_path + "/" + id}
              className="btn btn-primary"
            >
              View More
            </a>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Card;
