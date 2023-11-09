import CreateProjectForm from "../../components/form/CreateProjectForm"

const CreateProject = () => {
  return (
    <div>
    <div className="min-h-screen"
    style={{
      backgroundSize:'cover',
      backgroundPosition:'center',
    }}
    >
        <CreateProjectForm />
    </div>
    </div>
  )
}

export default CreateProject