import CreateProjectForm from "../../components/form/createProjectForm"

const createProject = () => {
  return (
    <div>
    <div className="min-h-screen"
    style={{
      backgroundImage:'url(src/assets/unsplash1.jpg)',
      backgroundSize:'cover',
      backgroundPosition:'center',
    }}
    >
        <CreateProjectForm />
    </div>
    </div>
  )
}

export default createProject