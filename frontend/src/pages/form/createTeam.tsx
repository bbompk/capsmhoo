import CreateTeamForm from "../../components/form/CreateTeamForm"

const CreateTeam = () => {
  return (
    <div>
    <div className="min-h-screen"
    style={{
      backgroundSize:'cover',
      backgroundPosition:'center',
    }}
    >
        <CreateTeamForm />
    </div>
    </div>
  )
}

export default CreateTeam