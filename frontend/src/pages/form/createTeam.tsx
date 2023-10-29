import CreateTeamForm from "../../components/form/createTeamForm"

const createTeam = () => {
  return (
    <div>
    <div className="min-h-screen"
    style={{
      backgroundImage:'url(src/assets/unsplash1.jpg)',
      backgroundSize:'cover',
      backgroundPosition:'center',
    }}
    >
        <CreateTeamForm />
    </div>
    </div>
  )
}

export default createTeam