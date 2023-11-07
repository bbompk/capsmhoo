import "./Card.css";

interface Props {
  id: string;
  name: string;
  profile: string;
}

const Card = ({ id, name, profile }: Props) => {
  return (
    <div className="card-container">
      <div className="card-content">
        <div className="card-title">
          <h5>{name}</h5>
        </div>
        <div className="card-body">
          <div className="card-text">
            <p>{profile}</p>
          </div>
          <div className="card-botton">
            <a href={"http://localhost:5173/view-team/" + id} className="btn btn-primary">
              View More
            </a>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Card;